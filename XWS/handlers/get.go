package handlers

import (
	"fmt"
	"net/http"

	"xws_proj/data"

	"github.com/gorilla/mux"
)

// ListAll handles GET requests and returns all current products
func (u *Users) ListAll(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] get all records")
	//Check if the person requesting is logged in
	// they have a coockie
	session, _ := data.Store.Get(r, "session")
	_, ok := session.Values["username"]
	fmt.Println("ok: ", ok)
	if !ok {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}
	u.l.Println("[DEBUG] you are logged in")

	uss := data.GetUsers()

	err := data.ToJSON(uss, rw)
	if err != nil {
		// we should never be here but log the error just incase
		u.l.Println("[ERROR] serializing product", err)
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
		u.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		u.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	err = data.ToJSON(uss, rw)
	if err != nil {
		u.l.Println("[ERROR] serializing product", err)
	}

}
