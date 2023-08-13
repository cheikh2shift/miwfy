package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Add[T any](db DatabaseWrapper) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T

		err := c.BindJSON(&req)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		err = db.Create(req)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"result": "Data created",
			},
		)
	}
}

func Update[T any](db DatabaseWrapper) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T

		id := c.Param("id")

		if id == "" {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": "URL parameter `id` is required",
				},
			)
			return
		}

		err := c.BindJSON(&req)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		intId, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()},
			)
			return
		}

		err = db.Update(intId, req)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"result": "Row updated",
			},
		)
	}
}

func Remove[T any](db DatabaseWrapper) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		if id == "" {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": "URL parameter `id` is required",
				},
			)
			return
		}

		idInt, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()},
			)
			return
		}

		err = db.Delete(idInt)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"result": "row removed",
			},
		)
	}
}

func Read[T any](db DatabaseWrapper) gin.HandlerFunc {
	return func(c *gin.Context) {

		var r []T

		err := db.ReadAll(0, &r)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"result": r,
			},
		)

	}
}

func ReadOne[T any](db DatabaseWrapper) gin.HandlerFunc {
	return func(c *gin.Context) {

		query := c.Param("id")
		var r T

		idInt, err := strconv.Atoi(query)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		err = db.Read(idInt, &r)

		if err != nil {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"result": r,
			},
		)

	}
}
