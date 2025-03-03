package routes

import (
	"golearn/first-api/logger"
	"golearn/first-api/model/category"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getCategory(c *gin.Context) {
	categoryID, err := strconv.ParseInt(c.Param("categoryID"), 10, 64)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	category, err := category.Get(categoryID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.JSON(http.StatusOK, category)
}

func getCategoriesOfGame(c *gin.Context) {
	gameID := c.GetInt64("gameID")
	categories, err := category.GetByGameID(gameID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.JSON(http.StatusOK, categories)
}

func postCategory(c *gin.Context) {
	var category category.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	userID := c.GetInt64("userID")
	gameID := c.GetInt64("gameID")

	err = category.Save(userID, gameID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.Status(http.StatusCreated)
}

func putCategory(c *gin.Context) {
	var reqCategory category.Category
	categoryID, err := strconv.ParseInt(c.Param("categoryID"), 10, 64)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	err = c.ShouldBindJSON(&reqCategory)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	dbCategory, err := category.Get(categoryID)
	if dbCategory == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The 'categoryID' parameter could not be resolved to a resource."})
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	err = dbCategory.UpdateWith(reqCategory)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.Status(http.StatusOK)
}

func deleteCategory(c *gin.Context) {
	categoryID, err := strconv.ParseInt(c.Param("categoryID"), 10, 64)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	category, err := category.Get(categoryID)
	if category == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The 'categoryID' parameter could not be resolved to a resource."})
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	err = category.Delete()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.Status(http.StatusOK)
}
