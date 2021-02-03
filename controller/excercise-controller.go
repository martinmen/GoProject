package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
	"gitlab.com/pragmaticreviews/golang-gin-poc/validators"
)

type ExcerciseController interface {
	FindAll() []entity.Exercise
	Save(context *gin.Context) error
	Update(context *gin.Context) error
	Delete(context *gin.Context) error
	ShowAll(context *gin.Context)
}
type controller struct {
	service service.ExerciseService
}

var validate *validator.Validate

func New(service service.ExerciseService) ExcerciseController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)

	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Exercise {
	return c.service.FindAll()
}

func (c *controller) Save(context *gin.Context) error {
	var exercise entity.Exercise
	err := context.ShouldBindJSON(&exercise)
	if err != nil {
		return err
	}
	err = validate.Struct(exercise)
	if err != nil {
		return err
	}
	c.service.Save(exercise)
	return nil
}

func (c *controller) Update(context *gin.Context) error {

	var exercise entity.Exercise
	err := context.ShouldBindJSON(&exercise)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	exercise.ID = id

	err = validate.Struct(exercise)
	if err != nil {
		return err
	}
	c.service.Update(exercise)
	return nil

}
func (c *controller) Delete(context *gin.Context) error {
	var exercise entity.Exercise
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	exercise.ID = id
	c.service.Delete(exercise)
	return nil
}

func (c *controller) ShowAll(context *gin.Context) {
	exercise := c.service.FindAll()
	data := gin.H{
		"title":    "Exercise Page",
		"exercise": exercise,
	}
	context.HTML(http.StatusOK, "index.html", data)
}
