package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"xws_proj/data"

	"github.com/gorilla/mux"
)

// Create handles POST requests to add new users
func (u *Users) Register(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] entered create")
	user := r.Context().Value(KeyUser{}).(*data.User)
	u.l.Println("[DEBUG] retreived the user")
	var err error
	*user.Password, err = data.HashPassword(*user.Password)
	user.Role = data.RegUser
	if err != nil {
		fmt.Println("[ERROR] hashing password")
		return
	}
	//u.l.Printf("[DEBUG] Inserting user:\n %#v\n", user)
	data.AddUser(*user)
}

func (u *Users) LogIn(rw http.ResponseWriter, r *http.Request) {
	var user *data.User
	err := data.FromJSON(&user, r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	//check if there is a user with that username
	userFromDB, err := data.GetUserByUsername(user.Username)
	if err != nil {
		http.Error(rw, "User not registered", http.StatusNoContent)
	}
	//check if the user is already logged in
	session, _ := data.Store.Get(r, "session")
	if session.Values["username"] == user.Username {
		http.Error(rw, "Already logged in", http.StatusOK)
		return
	}

	if data.CheckPasswordHash(*user.Password, *userFromDB.Password) {
		//logged in successfully
		fmt.Printf("User %s successfully logged in\n", user.Username)

		//session, _ := data.Store.Get(r, "session")
		session.Values["username"] = user.Username
		session.IsNew = false
		session.Save(r, rw)
		return
	}
	//password is incorrect
	u.l.Println("Password is incorrect")
}

func (u *Users) LogOut(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] Logging out")
	session, _ := data.Store.Get(r, "session")
	delete(session.Values, "username")
	session.Save(r, rw)
}

func (u *Users) CreatePost(rw http.ResponseWriter, r *http.Request) {

	session, _ := data.Store.Get(r, "session")
	username, ok := session.Values["username"]
	if !ok {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}
	u.l.Println("[DEBUG] you are logged in")

	var post *data.Post
	jsonFile, h, err := r.FormFile("json")
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err = data.FromJSON(&post, jsonFile)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	post.Username = username.(string)
	postID := data.AddPostToDB(*post)

	//notify followers about the post
	followers := data.GetAllFollowers(post.Username)
	fmt.Println(followers)
	var postNotifications data.PostNotifications
	for i := 0; i < len(followers); i++ {
		var postNotification data.PostNotification
		postNotification.PostID = postID
		postNotification.Recipient = followers[i]
		postNotifications = append(postNotifications, &postNotification)
	}
	fmt.Println(len(postNotifications))
	data.AddPostNotificationsToDB(postNotifications)

	err = r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		u.l.Println("[ERROR] parsing request body")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	//Access the photo key - First Approach
	file, h, err := r.FormFile("image")
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	tmpfile, err := os.Create("./" + h.Filename)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(tmpfile, file)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	data.StoreImageToDB(tmpfile.Name(), postID)
	tmpfile.Close()
	os.Remove(tmpfile.Name())
}

// handles request for following a user
func (u *Users) FollowUser(rw http.ResponseWriter, r *http.Request) {
	session, _ := data.Store.Get(r, "session")
	usernameFollower, ok := session.Values["username"]
	if !ok {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}
	u.l.Println("[DEBUG] you are logged in")

	vars := mux.Vars(r)
	usernameFollowee := vars["username"]
	userFollowee, err := data.GetUserByUsername(usernameFollowee)
	if err != nil {
		http.Error(rw, "user was not found", http.StatusNotFound)
		return
	}
	if usernameFollowee == usernameFollower.(string) {
		http.Error(rw, "can't follow yourself", http.StatusForbidden)
		return
	}
	followTest, err := data.GetFollow(usernameFollower.(string), usernameFollowee)
	if err != nil && err.Error() != fmt.Errorf("follow not found").Error() {
		fmt.Println("DESI BATE")
		http.Error(rw, "server error", http.StatusInternalServerError)
		return
	}
	if followTest != nil {
		http.Error(rw, "already following that user", http.StatusConflict)
		return
	}
	if !userFollowee.Public {
		//generate request
		fmt.Println("[DEBUG] user is private, request sent")
		var followRequest data.FollowRequest
		followRequest.Requester = usernameFollower.(string)
		followRequest.Requestee = usernameFollowee
		data.AddFollowRequestToDB(followRequest)
		return
	}

	var follow data.Follow
	follow.Follower = usernameFollower.(string)
	follow.Followee = usernameFollowee
	data.AddFollowToDB(follow)

}

func (u *Users) AcceptFollowRequest(rw http.ResponseWriter, r *http.Request) {
	session, _ := data.Store.Get(r, "session")
	username, ok := session.Values["username"]
	if !ok {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}
	u.l.Println("[DEBUG] you are logged in")

	vars := mux.Vars(r)
	usernameRequest := vars["username"]
	followRequest, err := data.GetFollowRequest(usernameRequest, username.(string))
	if err != nil {
		http.Error(rw, "request not found", http.StatusNotFound)
		return
	}
	err = data.DeleteFollowRequest(followRequest.FollowRequestID)
	if err != nil {
		http.Error(rw, "deleting follow request", http.StatusInternalServerError)
	}
	var follow data.Follow
	follow.Followee = username.(string)
	follow.Follower = usernameRequest
	data.AddFollowToDB(follow)

}

func (u *Users) DeclineFollowRequest(rw http.ResponseWriter, r *http.Request) {
	session, _ := data.Store.Get(r, "session")
	username, ok := session.Values["username"]
	if !ok {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}
	u.l.Println("[DEBUG] you are logged in")

	vars := mux.Vars(r)
	usernameRequest := vars["username"]
	followRequest, err := data.GetFollowRequest(usernameRequest, username.(string))
	if err != nil {
		http.Error(rw, "request not found", http.StatusNotFound)
		return
	}
	err = data.DeleteFollowRequest(followRequest.FollowRequestID)
	if err != nil {
		http.Error(rw, "deleting follow request", http.StatusInternalServerError)
	}

}
