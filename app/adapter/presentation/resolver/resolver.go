package resolver

//go:generate go run github.com/99designs/gqlgen generate
import (
	"github.com/yuorei/video-server/app/application"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	usecase *application.UseCase
}

func NewResolver(app *application.Application) *Resolver {
	return &Resolver{
		usecase: application.NewUseCase(app),
	}
}
