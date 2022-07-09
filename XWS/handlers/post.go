package handlers

import (
	"fmt"
	followsProtos "follows_service/protos/follows"
	"io"
	"net/http"
	"os"
	"path/filepath"
	postsProtos "posts_service/protos/posts"
	userProtos "users_service/protos/user"
	"xws_proj/data"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/encoding/protojson"
)

// Create handles POST requests to add new users
func (u *Users) Register(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] entered create")
	user := r.Context().Value(KeyUser{}).(*data.User)
	u.l.Println("[DEBUG] retrieved the user")
	userProto := userProtos.UserResponse{Username: user.Username, Password: *user.Password, Email: user.Email}
	var err error
	userProto.Password, err = data.HashPassword(userProto.Password)
	userProto.Role = 2
	if err != nil {
		fmt.Println("[ERROR] hashing password")
		return
	}
	//u.l.Printf("[DEBUG] Inserting user:\n %#v\n", user)
	u.uc.CreateUser(r.Context(), &userProto)
}

func (u *Users) LogIn(rw http.ResponseWriter, r *http.Request) {
	var user userProtos.UserResponse
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	err = protojson.Unmarshal(bytes, &user)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	//check if there is a user with that username
	req := userProtos.UserByUsernameRequest{Username: user.Username}
	userFromDB, err := u.uc.GetUserByUsername(r.Context(), &req)
	//userFromDB, err := data.GetUserByUsername(user.Username)
	if err != nil {
		http.Error(rw, "User not registered", http.StatusNoContent)
	}
	_, err = r.Cookie("username")
	if err == nil {
		http.Error(rw, "already logged in", http.StatusBadRequest)
		return
	}

	if data.CheckPasswordHash(user.Password, userFromDB.Password) {
		//logged in successfully
		fmt.Printf("User %s successfully logged in\n", user.Username)
		return
	}
	//password is incorrect
	u.l.Println("Password is incorrect")
	http.Error(rw, "incorrect password", http.StatusUnauthorized)
}

func (u *Users) LogOut(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] Logging out")
	session, _ := data.Store.Get(r, "session")
	delete(session.Values, "username")
	session.Save(r, rw)
}

func (u *Users) CreatePost(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] create post")
	post := postsProtos.CreateRequest{}
	err := r.ParseMultipartForm(128 * 1024) // maxMemory 32MB
	if err != nil {
		u.l.Println("[ERROR] parsing request body")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	if username == "" {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}
	u.l.Println("[DEBUG] you are logged in")

	post.TxtContent = r.FormValue("txtContent")
	post.Hyperlink = r.FormValue("hyperlink")
	post.Username = username
	post.Comments = make([]*postsProtos.Comment, 0)
	post.Likes = make([]*postsProtos.Like, 0)
	emptyImage := ""
	file, h, err := r.FormFile("image")
	if err != nil {
		emptyImage = r.FormValue("image")
		if emptyImage == "" {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	if emptyImage == "" {
		tmpfile, err := os.Create("./PostImagesFS/" + h.Filename)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = io.Copy(tmpfile, file)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		tmpfile.Close()
		post.ImagePath = "./PostImagesFS/"
		post.ImageName = filepath.Base(tmpfile.Name())
	}

	CreateResponse, err := u.pc.CreatePost(r.Context(), &post)
	if err != nil {
		fmt.Println(err)
		return
	}
	postID, err := primitive.ObjectIDFromHex(CreateResponse.PostID)
	if err != nil {
		fmt.Println(err)
		return
	}

	//notify followers about the post
	followers := data.GetAllFollowers(post.Username)
	var postNotifications data.PostNotifications
	for i := 0; i < len(followers); i++ {
		var postNotification data.PostNotification
		postNotification.PostID = postID
		postNotification.Recipient = followers[i]
		postNotifications = append(postNotifications, &postNotification)
	}
	if len(postNotifications) != 0 {
		data.AddPostNotificationsToDB(postNotifications)
	}
	u.l.Println("[DEBUG] finished create post")

}

// handles request for following a user
func (u *Users) FollowUser(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] entered follow user")
	cookie, err := r.Cookie("username")
	if err != nil {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}
	usernameFollower := cookie.Value
	u.l.Println("[DEBUG] you are logged in")

	vars := mux.Vars(r)
	usernameFollowee := vars["username"]
	userFollowee, err := u.uc.GetUserByUsername(r.Context(), &userProtos.UserByUsernameRequest{Username: usernameFollowee})

	if err != nil {
		http.Error(rw, "user was not found", http.StatusNotFound)
		return
	}
	if usernameFollowee == usernameFollower {
		http.Error(rw, "can't follow yourself", http.StatusForbidden)
		return
	}
	followTest, err := u.fc.GetFollow(r.Context(), &followsProtos.Follow{Followee: usernameFollowee, Follower: usernameFollower})
	if err != nil {
		if err.Error() != "rpc error: code = Unknown desc = follow not found" {
			fmt.Println(err.Error())
			http.Error(rw, "server error", http.StatusInternalServerError)
			return
		}
	}
	if followTest != nil {
		http.Error(rw, "already following that user", http.StatusConflict)
		return
	}
	if !userFollowee.Public {
		//generate request
		fmt.Println("[DEBUG] user is private, request sent")
		_, err = u.fc.AddFollowRequestToDB(r.Context(), &followsProtos.FollowRequest{Requester: usernameFollower, Requestee: usernameFollowee})
		if err != nil {
			fmt.Println(err)
			http.Error(rw, "server error", http.StatusInternalServerError)
		}
		return
	}

	//if user is public just add follow immediately
	u.fc.AddFollowToDB(r.Context(), &followsProtos.Follow{Follower: usernameFollower, Followee: usernameFollowee})

}

func (u *Users) AcceptFollowRequest(rw http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}
	username := cookie.Value
	u.l.Println("[DEBUG] you are logged in")

	vars := mux.Vars(r)
	usernameRequest := vars["username"]
	req := followsProtos.FollowRequest{Requester: usernameRequest, Requestee: username}
	followRequest, err := u.fc.GetFollowRequest(r.Context(), &req)
	if err != nil {
		http.Error(rw, "request not found", http.StatusNotFound)
		return
	}
	delReq := followsProtos.DeleteFollowRRequest{FollowRequestID: followRequest.RequestID}
	_, err = u.fc.DeleteFollowRequest(r.Context(), &delReq)
	if err != nil {
		http.Error(rw, "deleting follow request", http.StatusInternalServerError)
	}

	follow := followsProtos.Follow{}
	follow.Followee = username
	follow.Follower = usernameRequest
	u.fc.AddFollowToDB(r.Context(), &follow)

}

func (u *Users) DeclineFollowRequest(rw http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}
	username := cookie.Value
	u.l.Println("[DEBUG] you are logged in")

	vars := mux.Vars(r)
	usernameRequest := vars["username"]
	followRequest, err := data.GetFollowRequest(usernameRequest, username)
	if err != nil {
		http.Error(rw, "request not found", http.StatusNotFound)
		return
	}
	err = data.DeleteFollowRequest(followRequest.FollowRequestID)
	if err != nil {
		http.Error(rw, "deleting follow request", http.StatusInternalServerError)
	}

}
