package resolver

//go:generate go run github.com/99designs/gqlgen generate
import (
	"github.com/yuorei/video-server/app/application"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	app *application.Application
}

func NewResolver(app *application.Application) *Resolver {
	return &Resolver{
		app: app,
	}
}
