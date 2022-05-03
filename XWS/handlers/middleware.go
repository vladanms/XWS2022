package handlers

import (
	"context"
	"net/http"
	"xws_proj/data"
)

// MiddlewareValidateProduct validates the product in the request and calls next if ok
func (u *Users) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := &data.User{}
		u.l.Println("[DEBUG] entered middleware validation")
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
