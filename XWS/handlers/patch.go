package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"xws_proj/data"

	userProtos "users_service/protos/user"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/golang/protobuf/jsonpb"
)

func (u *Users) UpdateProfile(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] updating user profile")
	cookie, err := r.Cookie("username")
	if err != nil {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}

	username := cookie.Value
	usr, err := u.uc.GetUserByUsername(r.Context(), &userProtos.UserByUsernameRequest{Username: username})
	fmt.Println(usr.DateOfBirth)
	//usr, err := data.GetUserByID(id)
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
	var marshaler jsonpb.Marshaler
	original, err := marshaler.MarshalToString(usr)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	originalBytes := []byte(original)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	bodyString := string(body)
	fmt.Println(bodyString)
	patchJSON := []byte(`[` + bodyString + `]`)
	patch, err := jsonpatch.DecodePatch(patchJSON)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	modified, err := patch.Apply(originalBytes)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	var user userProtos.UserResponse
	err = jsonpb.Unmarshal(bytes.NewReader(modified), &user)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(user.DateOfBirth.AsTime())
	_, err = u.uc.UpdateUser(r.Context(), &user)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, "server error", http.StatusInternalServerError)
	}

	//data.UpdateUser(*user, id)
}
