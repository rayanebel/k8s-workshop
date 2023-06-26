package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var articles = Articles{}

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Articles struct {
	Articles []Article `json:"articles"`
}

func main() {
	initArticles := []Article{
		{
			ID:      1,
			Title:   "Introduction to Go",
			Content: "Go is a statically typed, compiled programming language designed at Google. It is known for its simplicity, efficiency, and strong support for concurrency.",
		},
		{
			ID:      2,
			Title:   "Web Development with Go",
			Content: "Go has a robust standard library that makes it easy to build web applications. It provides a built-in HTTP server, a powerful template engine, and support for handling HTTP requests and responses.",
		},
		{
			ID:      3,
			Title:   "Concurrency in Go",
			Content: "Go has excellent support for concurrency with goroutines and channels. Goroutines allow you to run lightweight concurrent functions, while channels enable communication and synchronization between goroutines.",
		},
		{
			ID:      4,
			Title:   "Advanced Go Topics",
			Content: "Go offers advanced features such as reflection, interfaces, and composition that allow developers to write modular and extensible code. It also has built-in testing and benchmarking support.",
		},
		{
			ID:      5,
			Title:   "Go Best Practices",
			Content: "There are several best practices to follow when developing with Go, including writing clean and idiomatic code, handling errors properly, using the standard library effectively, and optimizing performance when needed.",
		},
	}

	articles = Articles{
		Articles: initArticles,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to my blog"))
	})

	r.Route("/articles", func(r chi.Router) {
		r.Get("/", listArticles)
		r.Post("/", createArticle)
		r.Get("/{id}", getArticle)
		r.Delete("/{id}", deleteProduct)
	})

	http.ListenAndServe(":3000", r)
}

func listArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(articles)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	article := Article{}
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	article.ID = len(articles.Articles) + 1
	articles.Articles = append(articles.Articles, article)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Article has been successfully created"))
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	articleID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var foundArticle *Article
	for _, article := range articles.Articles {
		if article.ID == articleID {
			foundArticle = &article
			break
		}
	}

	if foundArticle == nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(foundArticle)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	articleID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var foundArticle bool
	for index, article := range articles.Articles {
		if article.ID == articleID {
			foundArticle = true
			articles.Articles = append(articles.Articles[:index], articles.Articles[index+1:]...)
			break
		}
	}

	if !foundArticle {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.Write([]byte("OK"))
}
