package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	followProtos "follows_service/protos/follows"
	postProtos "posts_service/protos/posts"
	protos "users_service/protos/user"
	"xws_proj/data"

	"github.com/golang/protobuf/jsonpb"
	"github.com/gorilla/mux"
)

func (u *Users) GetRequests(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] getting requests")
	cookie, err := r.Cookie("username")
	if err != nil {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}
	req := followProtos.GetFollowRRequest{Username: cookie.Value}
	requests, err := u.fc.GetFollowRequests(r.Context(), &req)
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = no requests" {
			http.Error(rw, "no requests", http.StatusNoContent)
			return
		}
		u.l.Println("[ERROR] mongo error")
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}
	if requests == nil {
		u.l.Println("[ERROR] no requests")
		http.Error(rw, "no follow requests", http.StatusNoContent)
		return
	}
	data.ToJSON(requests.Results, rw)
}

// get a follow connection if one exists
func (u *Users) GetFollow(rw http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}
	usernameFollower := cookie.Value
	vars := mux.Vars(r)
	usernameFollowee := vars["username"]
	if usernameFollowee == usernameFollower {
		u.l.Println("[DEBUG] same usernames")
		emptyFollow := &data.Follow{}
		data.ToJSON(emptyFollow, rw)
		return
	}
	req := followProtos.Follow{Follower: usernameFollower, Followee: usernameFollowee}
	follow, err := u.fc.GetFollow(r.Context(), &req)
	if err != nil && err.Error() != fmt.Errorf("rpc error: code = Unknown desc = follow not found").Error() {
		http.Error(rw, "server error", http.StatusInternalServerError)
		return
	}
	if follow == nil {
		http.Error(rw, "not following user", http.StatusNoContent)
		return
	}
	data.ToJSON(follow, rw)
}

func (u *Users) GetImageFromPost(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] get image from post")
	vars := mux.Vars(r)
	filename := vars["fname"]
	u.l.Println(filename)
	file, err := os.Open("./PostImagesFS/" + filename)
	if err != nil {
		http.Error(rw, "file not found", http.StatusNotFound)
		return
	}
	http.ServeFile(rw, r, file.Name())

}

// ListAll handles GET requests and returns all current users
func (u *Users) ListAll(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] get all records")
	ur := &protos.UsersRequest{}
	resp, err := u.uc.GetUsers(context.Background(), ur)
	if err != nil {
		u.l.Println("[ERROR] getting users via microservice")
	}
	err = data.ToJSON(resp.Results, rw)
	if err != nil {
		// we should never be here but log the error just incase
		u.l.Println("[ERROR] serializing user", err)
	}
}
func (u *Users) ListAllPublic(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] enetered list all public")
	users := data.GetUsers()
	var publicUsers data.Users
	for i := 0; i < len(users); i++ {
		if users[i].Public {
			users[i].Password = nil
			publicUsers = append(publicUsers, users[i])
		}
	}
	if len(publicUsers) == 0 {
		http.Error(rw, "No public users", http.StatusNotFound)
		return
	}
	err := data.ToJSON(publicUsers, rw)
	if err != nil {
		// we should never be here but log the error just incase
		u.l.Println("[ERROR] serializing public users", err)
		http.Error(rw, "serializing error", http.StatusInternalServerError)
		return
	}
}

func (u *Users) ListSingle(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] get single record")

	vars := mux.Vars(r)
	id := vars["id"]
	uss, err := data.GetUserByID(id)

	switch err {
	case nil:
	case data.ErrUserNotFound:
		u.l.Println("[ERROR] fetching user", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		u.l.Println("[ERROR] fetching user", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	err = data.ToJSON(uss, rw)
	if err != nil {
		u.l.Println("[ERROR] serializing user", err)
	}

}
func (u *Users) GetPostsLen(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] get posts len")
	cookie, err := r.Cookie("username")
	if err == http.ErrNoCookie {
		http.Error(rw, "no cookie provided", http.StatusUnauthorized)
		return
	}
	username := cookie.Value
	posts, _ := data.GetPostsUser(username)
	data.ToJSON(posts, rw)
}
func (u *Users) GetUserByUsername(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] get user by username")
	vars := mux.Vars(r)
	username := vars["username"]
	req := protos.UserByUsernameRequest{Username: username}
	fmt.Println(req)
	user, err := u.uc.GetUserByUsername(r.Context(), &req)
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = mongo: no documents in result" {
			http.Error(rw, "no user with that username", http.StatusNoContent)
			return
		}
		u.l.Println("[ERROR] retrieving user from db")
		http.Error(rw, "database error", http.StatusInternalServerError)
		return
	}
	//err = data.ToJSON(user, rw)
	var marshaler jsonpb.Marshaler
	err = marshaler.Marshal(rw, user)
	if err != nil {
		fmt.Println(err)
	}
}

func (u *Users) GetAllPostsFromUser(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] get all posts")
	vars := mux.Vars(r)
	username := vars["username"]
	cookie, err := r.Cookie("username")
	if err != nil {
		//if there is no cookie, continue
		cookie = &http.Cookie{}
		cookie.Value = ""
	}
	userReq := protos.UserByUsernameRequest{Username: username}
	user, err := u.uc.GetUserByUsername(r.Context(), &userReq)
	//user, err := data.GetUserByUsername(username)
	if err != nil {
		u.l.Println("[ERROR] retrieving user from db")
		return
	}
	if username != cookie.Value {
		if !user.Public {
			//not logged in
			if cookie.Value == "" {
				u.l.Println("[ERROR] profile is private")
				http.Error(rw, "profile is private", http.StatusUnauthorized)
				return
			}
			req := followProtos.Follow{Follower: cookie.Value, Followee: username}
			_, err := u.fc.GetFollow(r.Context(), &req)
			//_, err := data.GetFollow(cookie.Value, username)
			//not following user
			if err != nil {
				u.l.Println("[ERROR] profile is private")
				http.Error(rw, "profile is private", http.StatusUnauthorized)
				return
			}
		}
	}
	req := postProtos.PostsRequest{Username: username}
	posts, err := u.pc.GetAllPostsFromUser(r.Context(), &req)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}
	err = data.ToJSON(posts.Results, rw)
	if err != nil {
		u.l.Println("[ERROR] marshaling to json")
		return
	}
	u.l.Println("[DEBUG] finished getting posts for user")
}

func (u *Users) GetNotificationPosts(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] entered getting notif posts")
	cookie, err := r.Cookie("username")
	if err == http.ErrNoCookie {
		http.Error(rw, "no cookie provided", http.StatusUnauthorized)
		return
	}
	username := cookie.Value
	u.l.Println("[DEBUG] you are logged in")
	req := postProtos.NotificationPostsRequest{Username: username}
	response, err := u.pc.GetNotificationPosts(r.Context(), &req)
	if err != nil {
		http.Error(rw, "no notifs", http.StatusNoContent)
		return
	}

	data.ToJSON(response.Results, rw)
	u.l.Println("[DEBUG] finished getting notif posts")
}

func (u *Users) GetRequest(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] get single request")
	cookie, err := r.Cookie("username")
	if err == http.ErrNoCookie {
		http.Error(rw, "no cookie provided", http.StatusUnauthorized)
		return
	}
	requester := cookie.Value
	vars := mux.Vars(r)
	requestee := vars["username"]
	followRequest, err := u.fc.GetFollowRequest(r.Context(), &followProtos.FollowRequest{Requester: requester, Requestee: requestee})
	if err != nil {
		if err.Error() != "rpc error: code = Unknown desc = follow request not found" {
			http.Error(rw, "server error", http.StatusInternalServerError)
			return
		}
		http.Error(rw, "no request", http.StatusNoContent)
		return
	}
	data.ToJSON(followRequest, rw)

}
