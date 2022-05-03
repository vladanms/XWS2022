package handlers

import (
	"fmt"
	"net/http"
	"xws_proj/data"
)

// Create handles POST requests to add new users
func (u *Users) Register(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] entered create")
	user := r.Context().Value(KeyUser{}).(*data.User)
	u.l.Println("[DEBUG] retreived the user")
	var err error
	user.Password, err = data.HashPassword(user.Password)
	if err != nil {
		fmt.Println("[ERROR] hashing password")
		return
	}
	//u.l.Printf("[DEBUG] Inserting user:\n %#v\n", user)
	data.AddUser(*user)
}

func (u *Users) LogIn(rw http.ResponseWriter, r *http.Request) {
	user := data.User{}
	err := data.FromJSON(user, r.Body)
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

	if data.CheckPasswordHash(user.Password, userFromDB.Password) {
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
