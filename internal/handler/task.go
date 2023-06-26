package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
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
		return
	}

	lists := h.services.Task.GetAll(userId)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getTaskByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	list := h.services.Task.GetByID(userId, uint(id))

	c.JSON(http.StatusOK, list)
}

func (h *Handler) createTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input model.Task

	if err := c.BindJSON(&input); err != nil {
		return
	}

	id := h.services.Task.Create(userId, input)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		log.Fatal(err)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	h.services.Task.Delete(userId, uint(id))

	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})

}

func (h *Handler) updateTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		log.Fatal(err)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	var input model.UpdateTaskInput
	if err := c.BindJSON(&input); err != nil {
		log.Fatal(err)
	}

	err = h.services.Task.Update(userId, uint(id), input)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
