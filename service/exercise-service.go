package service

import (
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/repository"
)

type ExerciseService interface {
	Save(entity.Exercise) error
	Update(exercise entity.Exercise) error
	Delete(exercise entity.Exercise) error
	FindAll() []entity.Exercise
}

type exerciseService struct {
	repository repository.ExerciseRepository
}

func New(exerciseRepository repository.ExerciseRepository) ExerciseService {
	return &exerciseService{
		repository: exerciseRepository,
	}
}

func (service *exerciseService) Save(exercise entity.Exercise) error {
	service.repository.Save(exercise)
	return nil
}

func (service *exerciseService) Update(exercise entity.Exercise) error {
	service.repository.Update(exercise)
	return nil
}

func (service *exerciseService) Delete(exercise entity.Exercise) error {
	service.repository.Delete(exercise)
	return nil
}

func (service *exerciseService) FindAll() []entity.Exercise {
	return service.repository.FindAll()
}
