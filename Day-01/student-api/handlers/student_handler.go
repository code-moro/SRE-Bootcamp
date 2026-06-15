package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"student-api/models"
	"student-api/repository"
)

func CreateStudent(c *gin.Context) {

	var student models.Student

	if err := c.BindJSON(&student); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := repository.AddStudent(student)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "student created",
		},
	)
}

func GetStudents(c *gin.Context) {

	students, err := repository.GetStudents()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(
		http.StatusOK,
		students,
	)
}

func GetStudent(c *gin.Context) {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})

		return
	}

	student, err := repository.GetStudent(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "student not found",
		})

		return
	}

	c.JSON(
		http.StatusOK,
		student,
	)
}

func UpdateStudent(c *gin.Context) {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})

		return
	}

	var student models.Student

	if err := c.BindJSON(&student); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err = repository.UpdateStudent(
		id,
		student,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "student updated",
		},
	)
}

func DeleteStudent(c *gin.Context) {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})

		return
	}

	err = repository.DeleteStudent(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.Status(
		http.StatusNoContent,
	)
}

func HealthCheck(c *gin.Context) {

	c.JSON(
		http.StatusOK,
		gin.H{
			"status": "healthy",
		},
	)
}