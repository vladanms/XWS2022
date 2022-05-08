package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"xws_proj/data"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/gorilla/mux"
)

func (u *Users) UpdateProfile(rw http.ResponseWriter, r *http.Request) {

	session, _ := data.Store.Get(r, "session")
	_, ok := session.Values["username"]
	if !ok {
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	usr, err := data.GetUserByID(id)
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

	if session.Values["username"] != usr.Username {
		http.Error(rw, "must log in as the user", http.StatusUnauthorized)
		return
	}

	original, err := json.Marshal(usr)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	bodyString := string(body)
	patchJSON := []byte(bodyString)

	patch, err := jsonpatch.DecodePatch(patchJSON)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	modified, err := patch.Apply(original)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var user *data.User
	json.Unmarshal(modified, &user)

	session.Values["username"] = user.Username
	session.Save(r, rw)

	data.UpdateUser(*user, id)
}
