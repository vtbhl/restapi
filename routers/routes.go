package routers

import (
	"fmt"
	"net/http"
	"github.com/ittrada/restapi/structs"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TutorialEdge REST API: v0.0.1")
}

var routes = Routes{
	Route{"Index", "GET", "/", Index},

	// All Tags Routes
	Route{"Alltags", "GET", "/tags", structs.AllTags},
	Route{"Gettag", "GET", "/tag/{id}", structs.GetTag},
	Route{"NewTag", "POST", "/tag", structs.InsertTag},
	Route{"EditTag", "POST", "/tag/{id}", structs.EditTag},
	Route{"DeleteTag", "DELETE", "/tag/{id}", structs.DeleteTag},
	// All Blog Posts Routes
	// ...
	Route{"AllPosts", "GET", "/posts", structs.AllPosts},
	Route{"GetPost", "GET", "/post/{id}", structs.GetPost},
	Route{"NewPost", "POST", "/post", structs.InsertPost},
	Route{"EditPost", "POST", "/post/{id}", structs.EditPost},
	Route{"DeletePost", "DELETE", "/post/{id}", structs.AllPosts},
	// All Tutorials Routes
	Route{"AllTutorials", "GET", "/tutorials", structs.AllTutorials},
	Route{"GetTutorial", "GET", "/tutorial/{id}", structs.GetTutorial},
	Route{"NewTutorial", "POST", "/tutorial", structs.InsertTutorial},
	Route{"EditTutorial", "POST", "/tutorial/{id}", structs.EditTutorial},
	Route{"DeleteTutorial", "DELETE", "/tutorial/{id}", structs.AllTutorials},
	// All Categories
	Route{"AllCategories", "GET", "/categories", structs.AllCategories},
	Route{"GetCategory", "GET", "/category/{id}", structs.GetCategory},
	Route{"NewCategory", "POST", "/category", structs.InsertCategory},
	Route{"EditCategory", "POST", "/category/{id}", structs.EditCategory},
	Route{"DeleteCategory", "DELETE", "/category/{id}", structs.AllCategories},
	// All Categories
	Route{"AllCourses", "GET", "/courses", structs.AllCourses},
	// Route{"GetCategory", "GET", "/category/{id}", GetCategory},
	// Route{"NewCategory", "POST", "/category", InsertCategory},
	// Route{"EditCategory", "POST", "/category/{id}", EditCategory},
	// Route{"DeleteCategory", "DELETE", "/category/{id}", AllCategories},
}
