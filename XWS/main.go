package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"xws_proj/data"
	"xws_proj/handlers"

	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	env.Parse()

	l := log.New(os.Stdout, "users-api ", log.LstdFlags)
	v := data.NewValidation()

	// create the handlers
	ph := handlers.NewUsers(l, v)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// handlers for API
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/users", ph.ListAll)
	getR.HandleFunc("/users/{id:(?s).*}", ph.ListSingle)

	//putR := sm.Methods(http.MethodPut).Subrouter()
	//putR.HandleFunc("/products", ph.Update)
	//putR.Use(ph.MiddlewareValidateProduct)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/register", ph.Register)
	postR.Use(ph.MiddlewareValidateProduct)

	LogInR := sm.Methods(http.MethodPost).Subrouter()
	LogInR.HandleFunc("/login", ph.LogIn)

	LogOutR := sm.Methods(http.MethodPost).Subrouter()
	LogOutR.HandleFunc("/logout", ph.LogOut)

	PostImageR := sm.Methods(http.MethodPost).Subrouter()
	PostImageR.HandleFunc("/posts", ph.CreatePost)

	GetPostsR := sm.Methods(http.MethodGet).Subrouter()
	GetPostsR.HandleFunc("/posts/{username:(?s).*}", ph.GetAllPostsFromUser)

	//UpdateUserR := sm.Methods(http.MethodPatch).Subrouter()
	//UpdateUserR.HandleFunc("/users/{username:(?s).*}", ph.UpdateUser)

	getPublicUsersR := sm.Methods(http.MethodGet).Subrouter()
	getPublicUsersR.HandleFunc("/public-users", ph.ListAllPublic)

	FollowUserR := sm.Methods(http.MethodPost).Subrouter()
	FollowUserR.HandleFunc("/follow/{username:(?s).*}", ph.FollowUser)

	AcceptFollowRequestR := sm.Methods(http.MethodPost).Subrouter()
	AcceptFollowRequestR.HandleFunc("/follow-accept/{username:(?s).*}", ph.AcceptFollowRequest)

	DeclineFollowRequestR := sm.Methods(http.MethodPost).Subrouter()
	DeclineFollowRequestR.HandleFunc("/follow-decline/{username:(?s).*}", ph.DeclineFollowRequest)

	GetNotificationPostsR := sm.Methods(http.MethodGet).Subrouter()
	GetNotificationPostsR.HandleFunc("/notifications", ph.GetNotificationPosts)

	CommentR := sm.Methods(http.MethodPut).Subrouter()
	CommentR.HandleFunc("/comment", ph.AddComment)

	LikeR := sm.Methods(http.MethodPut).Subrouter()
	LikeR.HandleFunc("/like", ph.AddLike)

	UpdateProfileR := sm.Methods(http.MethodPatch).Subrouter()
	UpdateProfileR.HandleFunc("/user/{id:(?s).*}", ph.UpdateProfile)
	UpdateProfileR.Use(ph.MiddlewareValidateUpdate)

	// create a new server
	s := http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}
