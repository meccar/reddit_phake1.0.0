package web

// import (
// 	"html/template"
// 	"net/http"

// 	api "api"
// )

// func handleNews(w http.ResponseWriter, r *http.Request) {
// 	// Retrieve posts from the database
// 	posts, err := api.Handlers.GetPosts(r.Context())
// 	if err != nil {
// 		// Handle error (e.g., return HTTP 500 or render an error page)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	// Parse the HTML template
// 	tmpl, err := template.New("news").ParseFiles("news.html")
// 	if err != nil {
// 		// Handle template parsing error
// 		http.Error(w, "Error parsing template", http.StatusInternalServerError)
// 		return
// 	}

// 	// Execute the template with the posts data
// 	if err := tmpl.Execute(w, posts); err != nil {
// 		// Handle template execution error
// 		http.Error(w, "Error executing template", http.StatusInternalServerError)
// 		return
// 	}
// }
