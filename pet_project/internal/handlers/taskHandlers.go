package handlers

import (
	"context"
	"golang/pet_project/internal/tasksService" // Импортируем наш сервис
	"golang/pet_project/internal/web/tasks"
)

type Handler struct {
	Service *tasksService.TaskService
}

// Нужна для создания структуры Handler на этапе инициализации приложения
func NewHandler(service *tasksService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Text,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := tasksService.Task{
		Text:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}

// PatchTasks implements tasks.StrictServerInterface.
func (h *Handler) PatchTasks(ctx context.Context, request tasks.PatchTasksRequestObject) (tasks.PatchTasksResponseObject, error) {
	taskRequest := request.Body
	taskID := request.Body.Id

	taskToPatch := tasksService.Task{
		Text:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	patchedTask, err := h.Service.UpdateTaskByID(taskID, taskToPatch)

	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasks200JSONResponse{
		Id:     &patchedTask.ID,
		Task:   &patchedTask.Text,
		IsDone: &patchedTask.IsDone,
	}

	return response, nil
}

// DeleteTasks implements tasks.StrictServerInterface.
func (h *Handler) DeleteTasks(ctx context.Context, request tasks.DeleteTasksRequestObject) (tasks.DeleteTasksResponseObject, error) {
	taskID := request.Body.Id
	err := h.Service.DeleteTaskByID(taskID)
	if err != nil {
		return nil, err
	}

	response := tasks.DeleteTasks204Response{}
	return response, nil

}
