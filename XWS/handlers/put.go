package handlers

import (
	"net/http"
	"xws_proj/data"
)

func (u *Users) AddComment(content string, post *data.Post) (rw http.ResponseWriter, r *http.Request) {
	session, _ := data.Store.Get(r, "session")
	username, ok := session.Values["username"]
	if !ok {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}
	u.l.Println("[DEBUG] you are logged in")

	var comment *data.Comment
	comment.Author = username.(string)
	comment.Content = content

	data.AddCommentToPost(*post, *comment)

	return
}

func (u *Users) AddLike(content bool, post *data.Post) (rw http.ResponseWriter, r *http.Request) {
	session, _ := data.Store.Get(r, "session")
	username, ok := session.Values["username"]
	if !ok {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}
	u.l.Println("[DEBUG] you are logged in")

	var like *data.Like
	like.Author = username.(string)
	like.Content = content

	data.AddLikeToPost(*post, *like)

	return
}
