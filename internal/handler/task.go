package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todolist/internal/model"
)

func getUserId(c *gin.Context) (uint, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idUint, ok := id.(uint)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idUint, nil
}

type getAllListsResponse struct {
	Data []model.Task
}

func (h *Handler) getAllTasks(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	lists := h.Services.Task.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getTaskByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list := h.Services.Task.GetByID(userId, uint(id))

	c.JSON(http.StatusOK, list)
}

func (h *Handler) createTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	var input model.Task

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id := h.Services.Task.Create(userId, input)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":      id,
		"status":  "ok",
		"message": "task created",
	})
}

func (h *Handler) deleteTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	h.Services.Task.Delete(userId, uint(id))

	c.JSON(http.StatusOK, statusResponse{
		Status:  "ok",
		Message: "task deleted",
	})

}

func (h *Handler) updateTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input model.UpdateTaskInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Services.Task.Update(userId, uint(id), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status:  "ok",
		Message: "task updated",
	})
}
