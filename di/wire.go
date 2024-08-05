// wire.go
//go:build wireinject
// +build wireinject

package di

import (
	"time"

	"github.com/google/wire"
	"github.com/gupta-bhavesh/go-server/api/controller"
	"github.com/gupta-bhavesh/go-server/bootstrap"
	"github.com/gupta-bhavesh/go-server/mongo"
	"github.com/gupta-bhavesh/go-server/repository"
	"github.com/gupta-bhavesh/go-server/usecase"
)

func InitializeSignUpController(collection string, db mongo.Database, timeout time.Duration, env *bootstrap.Env) *controller.SignupController {
	wire.Build(repository.NewUserRepository, usecase.NewSignupUsecase, controller.NewSignupController)
	return nil
}
