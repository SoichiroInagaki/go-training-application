package handlers

import (
	"go-training-application/database"
	"go-training-application/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := task.Create(database.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {

	tasks, err := models.GetAll(database.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {

	db := database.DB

	//パスパラメーター(ID)の文字列をint型に変換
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	//IDに紐づくレコードを取得
	task, err := models.GetById(db, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	//リクエストボディを構造体にバインド
	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//更新後の構造体に修正
	task.Title = input.Title
	task.Description = input.Description

	//更新処理を実行
	if err := task.Update(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}
