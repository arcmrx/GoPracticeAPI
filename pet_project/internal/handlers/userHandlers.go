package handlers

import (
	"context"
	"golang/pet_project/internal/userService"
	"golang/pet_project/internal/web/users"
)

type UserHandler struct {
	Service *userservice.UserService
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userID := request.Id
	err := h.Service.DeleteUserByID(userID)
	if err != nil {
		return nil, err
	}

	response := users.DeleteUsersId204Response{}
	return response, nil
}

func (h *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userRequest := request.Body
	userID := request.Id

	userToPatch := userservice.User{}
	if userRequest.Email != nil {
		userToPatch.Email = *userRequest.Email
	}
	if userRequest.Password != nil {
		userToPatch.Password = *userRequest.Password
	}

	patchedUser, err := h.Service.PatchUserByID(userID, userToPatch)
	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &patchedUser.ID,
		Email:    &patchedUser.Email,
		Password: &patchedUser.Password,
	}

	return response, nil
}

func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userservice.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	createdUser, err := h.Service.PostUser(userToCreate)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	return response, nil
}

func NewUserHandler(service *userservice.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetUser()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	return response, nil
}


