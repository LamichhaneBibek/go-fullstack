package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/LamichhaneBibek/familytree/internal/models"
	"github.com/LamichhaneBibek/familytree/internal/models/sqlite"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	users *sqlite.UserModel
}

func main() {
    db, err := sql.Open("sqlite3", "./database.db")
    if err != nil {
        log.Fatalf("failed to open database: %v", err)
    }
    defer db.Close()

    log.Println("database connection established")

    app := &app{
        users: &sqlite.UserModel{DB: db},
    }



    srv := http.Server{
        Addr:    ":8080",
        Handler: app.routes(),
    }

    log.Println("server is listening on port 8080")

    err = srv.ListenAndServe()
    if err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.getHome)
	mux.HandleFunc("GET /users", app.createUser)
	mux.HandleFunc("POST /users", app.StoreUser)
	return mux
}


func (app *app)getHome(w http.ResponseWriter, r *http.Request) {
	users, err := app.users.All()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("./assets/templates/home.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, map[string]any{
		"users": users,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *app) createUser(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./assets/templates/register.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *app) StoreUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user := models.User{
		ID: string(uuid.New().String()[:8]),
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		Created: time.Now().String(),
		Updated: time.Now().String(),
	}
	err = app.users.Insert(&user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("user created successfully")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}