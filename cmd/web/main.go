package main

import (
	"fmt"
	"log"
	"time"

	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/mrbilll/bookings/pkg/config"
	"github.com/mrbilll/bookings/pkg/handlers"
	"github.com/mrbilll/bookings/pkg/render"
)

const portNumber = "localhost:8080"

var (
	app     config.AppConfig
	session *scs.SessionManager
)

// main is the main application function
func main() {

	// set this to true in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app) // give handlers access to config
	handlers.NewHandlers(repo)

	render.NewTemplates(&app) // give render access to config

	fmt.Println("Starting application on port number ", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
