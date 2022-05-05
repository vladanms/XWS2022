package handlers

import (
	"net/http"

	"xws_proj/data"

	"github.com/gorilla/mux"
)

// ListAll handles GET requests and returns all current users
func (u *Users) ListAll(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] get all records")
	uss := data.GetUsers()
	for i := 0; i < len(uss); i++ {
		uss[i].Password = nil
	}
	err := data.ToJSON(uss, rw)
	if err != nil {
		// we should never be here but log the error just incase
		u.l.Println("[ERROR] serializing user", err)
	}
}
func (u *Users) ListAllPublic(rw http.ResponseWriter, r *http.Request) {}

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
