package handlers

import (
	"fmt"
	"net/http"
	postProtos "posts_service/protos/posts"
	"xws_proj/data"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type getComment struct {
	IdPost  primitive.ObjectID `json:"idPost,omitempty"`
	Content string             `json:"content,omitempty"`
	//username string
}

type getLike struct {
	IdPost  primitive.ObjectID `json:"idPost,omitempty"`
	Content bool               `json:"content,omitempty"`
}

func (u *Users) AddComment(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] entered adding comment")
	cookie, err := r.Cookie("username")

	if err != nil {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "You must be logged in before this action", http.StatusUnauthorized)
		return
	}
	username := cookie.Value

	post := getComment{}
	data.FromJSON(&post, r.Body)
	u.l.Println(post)
	fmt.Println(post.IdPost)
	commentReq := postProtos.CommentRequest{Author: username, Content: post.Content, PostID: post.IdPost.Hex()}
	if post.Content == "" {
		http.Error(rw, "Comment must not be empty", http.StatusExpectationFailed)
		return
	}
	u.pc.AddCommentToPost(r.Context(), &commentReq)
}

func (u *Users) AddLike(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] entered adding like")
	cookie, err := r.Cookie("username")
	if err != nil {
		http.Error(rw, "You must be logged in before this action", http.StatusUnauthorized)
		return
	}
	username := cookie.Value
	post := getLike{}
	data.FromJSON(&post, r.Body)
	u.l.Println(post)
	like := postProtos.LikeRequest{Author: username, Content: post.Content, PostID: post.IdPost.Hex()}
	_, err = u.pc.AddLikeToPost(r.Context(), &like)
	if err != nil {
		fmt.Println(err)
	}
}
