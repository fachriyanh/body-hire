package routes

import (
	controllers "test_3/TEST_3/controller"

	"github.com/gin-gonic/gin"
)

type BookRouteController struct {
	postController controllers.BookController
}

func NewRouteBookController(postController controllers.BookController) BookRouteController {
	return BookRouteController{postController}
}

func (pc *BookRouteController) BookRoute(rg *gin.RouterGroup) {

	router := rg.Group("posts")
	router.POST("/", pc.postController.CreateBook)
	router.GET("/:name", pc.postController.GetBookByName)
	router.PUT("/:idBook", pc.postController.UpdateBookByID)
	router.DELETE("/:idBook", pc.postController.DeleteBook)
}
