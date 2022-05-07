package handlers

import (
	"net/http"
	"xws_proj/data"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type getComment struct {
	idPost  primitive.ObjectID
	content string
	//username string
}

type getLike struct {
	idPost  primitive.ObjectID
	content bool
}

func (u *Users) AddComment(rw http.ResponseWriter, r *http.Request) {
	session, _ := data.Store.Get(r, "session")
	username, ok := session.Values["username"]

	if !ok {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}
	u.l.Println("[DEBUG] you are logged in")

	var post getComment
	data.FromJSON(&post, r.Body)

	var comment *data.Comment
	comment.Author = username.(string)
	comment.Content = post.content

	data.AddCommentToPost(post.idPost.String(), *comment)

	return
}

func (u *Users) AddLike(rw http.ResponseWriter, r *http.Request) {
	session, _ := data.Store.Get(r, "session")
	username, ok := session.Values["username"]
	if !ok {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}
	u.l.Println("[DEBUG] you are logged in")

	var post getLike
	data.FromJSON(&post, r.Body)

	var like *data.Like
	like.Author = username.(string)
	like.Content = post.content

	data.AddLikeToPost(post.idPost.String(), *like)

	return
}
