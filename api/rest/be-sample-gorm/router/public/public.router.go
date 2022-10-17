package public

import (
	"github.com/HiBang15/sample-gorm.git/internal/module/user/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func SetRouter(router *gin.RouterGroup) {
	log.Print("Start init public router  BE SAMPLE GORM.....")

	users := router.Group("users")
	{
		users.POST("/", controller.CreateUser)
	}

	log.Print("Finish init public router BE SAMPLE GORM ....")
}
