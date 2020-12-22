package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(context *gin.Context) entity.Video
}
type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
func (c *controller) Save(context *gin.Context) entity.Video {
	var video entity.Video
	context.BindJSON(&video)
	c.service.Save(video)
	return video
}
