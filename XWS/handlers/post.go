package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"xws_proj/data"
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
	u.l.Println(username)
	post.Username = username.(string)
	postID := data.AddPostToDB(*post)
	u.l.Println("[DEBUG] returned from adding post to db")

	//r.Body = http.MaxBytesReader(rw, r.Body, 32<<20+512)
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
