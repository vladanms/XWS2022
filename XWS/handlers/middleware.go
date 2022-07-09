package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"xws_proj/data"

	jsonpatch "github.com/evanphx/json-patch"
)

// MiddlewareValidateProduct validates the product in the request and calls next if ok
func (u *Users) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := &data.User{}
		u.l.Println("[DEBUG] entered middleware validation")
		/*bytes, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			http.Error(rw, "server error", http.StatusInternalServerError)
			return
		}
		err = protojson.Unmarshal(bytes, user)*/
		err := data.FromJSON(user, r.Body)
		if err != nil {
			u.l.Println("[ERROR] deserializing user", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}
		// validate the product
		errs := u.v.Validate(user)
		if len(errs) != 0 {
			u.l.Println("[ERROR] validating user", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		// add the user to the context
		ctx := context.WithValue(r.Context(), KeyUser{}, user)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}

func (u *Users) MiddlewareValidateUpdate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := &data.User{}

		u.l.Println("[DEBUG] entered middleware validation")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		bodyString := string(body)
		if strings.Contains(bodyString, "op") && strings.Contains(bodyString, "path") && strings.Contains(bodyString, "value") {
			if strings.Contains(bodyString, "/id") {
				u.l.Println("[ERROR] invalid field update")

				rw.WriteHeader(http.StatusBadRequest)
				data.ToJSON(&GenericError{Message: err.Error()}, rw)
				return
			}
		} else {
			u.l.Println("[ERROR] json patch error")

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		patchJSON := []byte(string(body))

		patch, err := jsonpatch.DecodePatch(patchJSON)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		original, err := json.Marshal(user)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		originalString := string(original)
		originalString = `{
		"Username": "paul567",
		"Email": "333@box.com",
		"Password": "AGf1431!_",
		"Role": "",
		"Gender": 1,
		"Name": "",
		"PhoneNumber": "",
		"DateOfBirth": "2022-05-07T20:53:54.3985265+02:00",
		"Biography": "",
		"Experience": "",
		"Skills": "",
		"Interests": "",
		"Education": 1}`

		modified, err := patch.Apply([]byte(originalString))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		json.Unmarshal(modified, user)

		// validate the product
		errs := u.v.Validate(user)
		if len(errs) != 0 {
			u.l.Println("[ERROR] validating user", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		next.ServeHTTP(rw, r)
	})
}
