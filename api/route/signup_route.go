package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gupta-bhavesh/go-server/api/controller"
	"github.com/gupta-bhavesh/go-server/bootstrap"
	"github.com/gupta-bhavesh/go-server/domain"
	"github.com/gupta-bhavesh/go-server/mongo"
	"github.com/gupta-bhavesh/go-server/repository"
	"github.com/gupta-bhavesh/go-server/usecase"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
