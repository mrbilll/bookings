package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool                          // set to false for developer mode in main
	TemplateCache map[string]*template.Template // created in main: tc
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
