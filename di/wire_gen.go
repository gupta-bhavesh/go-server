// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/gupta-bhavesh/go-server/api/controller"
	"github.com/gupta-bhavesh/go-server/bootstrap"
	"github.com/gupta-bhavesh/go-server/mongo"
	"github.com/gupta-bhavesh/go-server/repository"
	"github.com/gupta-bhavesh/go-server/usecase"
	"time"
)

// Injectors from wire.go:

func InitializeSignUpController(collection string, db mongo.Database, timeout time.Duration, env *bootstrap.Env) *controller.SignupController {
	userRepository := repository.NewUserRepository(db, collection)
	signupUsecase := usecase.NewSignupUsecase(userRepository, timeout)
	signupController := controller.NewSignupController(env, signupUsecase)
	return signupController
}
