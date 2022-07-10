package handlers

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"xws_proj/data"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (u *Users) ListAllPublic(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] enetered list all public")
	users := data.GetUsers()
	var publicUsers data.Users
	for i := 0; i < len(users); i++ {
		if users[i].Public {
			users[i].Password = nil
			publicUsers = append(publicUsers, users[i])
		}
	}
	if len(publicUsers) == 0 {
		http.Error(rw, "No public users", http.StatusNotFound)
		return
	}
	err := data.ToJSON(publicUsers, rw)
	if err != nil {
		// we should never be here but log the error just incase
		u.l.Println("[ERROR] serializing public users", err)
		http.Error(rw, "serializing error", http.StatusInternalServerError)
		return
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
func (u *Users) GetAllPostsFromUser(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("[DEBUG] get all posts")

	vars := mux.Vars(r)
	username := vars["username"]

	user, err := data.GetUserByUsername(username)
	if err != nil {
		u.l.Println("[ERROR] retrieving user from db")
		return
	}
	if !user.Public {
		session, _ := data.Store.Get(r, "session")
		usernameLog, ok := session.Values["username"]
		//not logged in
		if !ok {
			u.l.Println("[ERROR] profile is private")
			http.Error(rw, "profile is private", http.StatusForbidden)
			return
		}
		_, err := data.GetFollow(usernameLog.(string), username)
		//not following user
		if err != nil {
			u.l.Println("[ERROR] profile is private")
			http.Error(rw, "profile is private", http.StatusForbidden)
			return
		}
	}
	posts, postIDs := data.GetPostsUser(username)
	images := data.GetImageByPostIDs(postIDs)
	writeFilesPosts(posts, images)
	os.Remove("test.zip")
}

func (u *Users) GetNotificationPosts(rw http.ResponseWriter, r *http.Request) {
	session, _ := data.Store.Get(r, "session")
	username, ok := session.Values["username"]
	if !ok {
		u.l.Println("[DEBUG] not logged in")
		http.Error(rw, "must log in first", http.StatusUnauthorized)
		return
	}
	u.l.Println("[DEBUG] you are logged in")
	postNotifications := data.GetPostNotificationsForUser(username.(string))
	if postNotifications == nil {
		http.Error(rw, "no notifications", http.StatusNoContent)
		return
	}
	var posts data.Posts
	postIDs := make([]primitive.ObjectID, len(postNotifications))
	for i := 0; i < len(postNotifications); i++ {
		posts = append(posts, data.GetSinglePost(postNotifications[i].PostID))
		postIDs[i] = postNotifications[i].PostID
	}

	images := data.GetImageByPostIDs(postIDs)
	writeFilesPosts(posts, images)
	rw.Header().Set("Content-Type", "application/zip")
	rw.Header().Set("Content-Disposition", "attachment; filename='test.zip'")
	http.ServeFile(rw, r, "test.zip")
	os.Remove("test.zip")

}

//function writes json files with posts and zips them with image files
func writeFilesPosts(posts data.Posts, images data.Images) {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	//create zip file for delivering all posts and images
	file, err := os.OpenFile("test.zip", flags, 0644)
	if err != nil {
		log.Fatalf("Failed to open zip for writing: %s", err)
	}
	defer file.Close()

	zipw := zip.NewWriter(file)
	defer zipw.Close()
	postsFiles := make([]string, 0)

	//create all json files for storing post data and write data in
	for i := 0; i < len(posts); i++ {
		filename := fmt.Sprintf("post%d.json", i)
		jsonFile, err := os.Create(filename)
		if err != nil {
			fmt.Println("[ERROR] creating file")
		}
		postsFiles = append(postsFiles, filename)
		file, _ := json.MarshalIndent(posts[i], "", " ")
		ioutil.WriteFile(filename, file, 0644)
		jsonFile.Close()
	}

	//appends files and writes them to zip
	for i := 0; i < len(posts); i++ {
		if err := data.AppendFiles(postsFiles[i], zipw); err != nil {
			log.Fatalf("Failed to add file %s to zip: %s", postsFiles[i], err)
		}
	}
	for i := 0; i < len(images); i++ {
		if err := data.AppendFiles(images[i].Filename, zipw); err != nil {
			log.Fatalf("Failed to add file %s to zip: %s", postsFiles[i], err)
		}
	}

	//delete all local files
	file.Close()
	zipw.Close()
	for i := 0; i < len(images); i++ {
		os.Remove(images[i].Filename)
	}
	for i := 0; i < len(posts); i++ {
		os.Remove(postsFiles[i])
	}
}
