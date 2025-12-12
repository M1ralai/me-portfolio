package main

import (
	"log"
	"net/http"

	"github.com/M1ralai/me-portfolio/internal/infrasturacture/logger"
	"github.com/M1ralai/me-portfolio/internal/infrasturacture/logger/db"

	postHandler "github.com/M1ralai/me-portfolio/internal/modules/post/handler"
	postRepo "github.com/M1ralai/me-portfolio/internal/modules/post/repository"
	postService "github.com/M1ralai/me-portfolio/internal/modules/post/service"

	contactHandler "github.com/M1ralai/me-portfolio/internal/modules/contact/handler"
	contactRepo "github.com/M1ralai/me-portfolio/internal/modules/contact/repository"
	contactService "github.com/M1ralai/me-portfolio/internal/modules/contact/service"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

func main() {
	db, err := db.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	validate := validator.New()
	logger := logger.NewLogger(db)

	/*
		POST MODÜLLERİ
	*/
	repoPost := postRepo.NewPostRepository(db)
	servicePost := postService.NewService(repoPost)
	handlerPost := postHandler.NewPostHandler(servicePost, logger, validate)

	/*
		CONTACT MODULLERİ
	*/

	repoContact := contactRepo.NewContactRepository(db)
	serviceContact := contactService.NewService(repoContact)
	handlerContact := contactHandler.NewContactHandler(serviceContact, logger, validate)

	router := mux.NewRouter()
	/*		middlewareler de eklenecek buraya

	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.TimeoutMiddleware)
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	bu sekilde olacak
	*/

	router.HandleFunc("/api/post/list", handlerPost.List).Methods("GET")
	router.HandleFunc("/api/post", handlerPost.GetById).Methods("GET")
	router.HandleFunc("/api/post", handlerPost.CreatePost).Methods("POST")
	router.HandleFunc("/api/post", handlerPost.UpdatePost).Methods("PUT")
	router.HandleFunc("/api/post", handlerPost.DeletePost).Methods("DELETE")

	router.HandleFunc("/api/contact", handlerContact.Delete).Methods("DELETE")
	router.HandleFunc("/api/contact", handlerContact.Create).Methods("POST")
	router.HandleFunc("/api/contact", handlerContact.List).Methods("GET")

	log.Println("server baslatiliyor")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
