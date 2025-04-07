package userservice

import (
	"golang/pet_project/internal/tasksService"
)

type UserService struct {
	repo        UserRepository
	TaskService tasksService.TaskService
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) PostUser(user User) (User, error) {
	return s.repo.PostUser(user)
}

func (s *UserService) GetUser() ([]User, error) {
	return s.repo.GetUsers()
}

func (s *UserService) PatchUserByID(id uint, user User) (User, error) {
	return s.repo.PatchUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}

func (s *UserService) GetTaskByUserId(id uint) ([]tasksService.Task, error) {
	return s.TaskService.GetTaskByUserId(id)
}