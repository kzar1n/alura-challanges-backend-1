package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/kzar1n/alura-challanges-backend-1/internal/controller"
)

const path = "/videos"

func HandleRequests() {
	r := gin.Default()
	r.POST(path, controllers.CreateVideo)
	r.GET(path, controllers.GetAllVideos)
	r.GET(path+"/:id_video", controllers.GetVideoById)
	r.PUT(path+"/:id_video", controllers.UpdateVideo)
	r.DELETE(path+"/:id_video", controllers.DeleteVideo)
	r.Run()
}
