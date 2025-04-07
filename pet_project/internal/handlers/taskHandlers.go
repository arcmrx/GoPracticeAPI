package handlers

import (
	"context"
	"golang/pet_project/internal/tasksService"
	"golang/pet_project/internal/web/tasks"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	Service *tasksService.TaskService
}

func (h *TaskHandler) GetUsersIdTasks(ctx context.Context, request tasks.GetUsersIdTasksRequestObject) (tasks.GetUsersIdTasksResponseObject, error) {
	tasksUser, err := h.Service.GetTaskByUserId(request.Id)
	if err != nil {
		return nil, err
	}
	response := tasks.GetUsersIdTasks202JSONResponse{}

	for _, tsk := range tasksUser {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

// Нужна для создания структуры TaskHandler на этапе инициализации приложения
func NewTaskHandler(service *tasksService.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
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
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body

	// Проверяем, что все обязательные поля присутствуют
	if taskRequest.Task == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Поле 'Task' отсутствует")
	}
	if taskRequest.IsDone == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Поле 'IsDone' отсутствует")
	}
	if taskRequest.UserId == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Поле 'UserId' отсутствует")
	}

	// Обращаемся к сервису и создаем задачу
	taskToCreate := tasksService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserID: *taskRequest.UserId,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskRequest := request.Body
	taskID := request.Id

	taskToPatch := tasksService.Task{}
	if taskRequest.Task != nil {
		taskToPatch.Task = *taskRequest.Task
	}
	if taskRequest.IsDone != nil {
		taskToPatch.IsDone = *taskRequest.IsDone
	}
	if taskRequest.UserId != nil {
		taskToPatch.UserID = *taskRequest.UserId
	}

	patchedTask, err := h.Service.UpdateTaskByID(taskID, taskToPatch)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &patchedTask.ID,
		Task:   &patchedTask.Task,
		IsDone: &patchedTask.IsDone,
		UserId: &patchedTask.UserID,
	}

	return response, nil
}

func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id
	err := h.Service.DeleteTaskByID(taskID)
	if err != nil {
		return nil, err
	}

	response := tasks.DeleteTasksId204Response{}
	return response, nil
}
