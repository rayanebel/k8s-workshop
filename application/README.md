# Create a Blog API

The goal of this application is to create a simple REST API to manage the articles of your blog.
This API will be written in Golang and it should be simple. No need to add a databases for now or to make the code thread-safe.

The articles will be stored in a variable which will be an object with the following properties:

- ID
- Title
- Content

This API should have 5 routes:
- `/` `[GET]`: Return just a message.
- `/articles` `[GET]`: Get a list of all the articles.
- `/articles` `[POST]`: Create a new article.
- `/articles/{id}` `[GET]`: Get an article by using its unique ID.
- `/articles/{id}` `[DELETE]`: Delete an article by using its unique ID.


To get a few items when you start the application, you can pre-fill your variables with the following content:

```golang
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
```

When the application is working _(not necessarily finished)_, you can start creating the Dockerfile and building a container image with your application.

## Cheatsheet

### Create a CHI router
```golang
r := chi.NewRouter()
```

### Create a route

```golang
r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my blog"))
})

// OR a group of routes

r.Route("/group", func(r chi.Router) {
    r.Get("/", getFunction)
    r.Post("/", PostFunction)
})
```

### Start the HTTP server
```golang
http.ListenAndServe(":3000", r)
```

### Create a handler function

```golang
func handlerfunc(w http.ResponseWriter, r *http.Request) {
    // Your code here
}
```

### Get url parameter with CHI
```golang
chi.URLParam(r, "id")
```

### Convert a string to a int
```golang
strconv.Atoi()
```