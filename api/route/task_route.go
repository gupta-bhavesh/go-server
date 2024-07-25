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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
