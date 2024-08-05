package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gupta-bhavesh/go-server/bootstrap"
	"github.com/gupta-bhavesh/go-server/di"
	"github.com/gupta-bhavesh/go-server/domain"
	"github.com/gupta-bhavesh/go-server/mongo"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	sc := di.InitializeSignUpController(domain.CollectionUser, db, timeout, env)
	group.POST("/signup", sc.Signup)
}
