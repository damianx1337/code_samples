package controllers

import (
	"net/http"
  "fmt"

	"psgres/models"

	"github.com/gin-gonic/gin"
)


type CreateUserInput struct {
	Name      string `json:"name" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, FirstName: input.FirstName}
  //users = [5000]User{{Name: "jinzhu", Pets: []Pet{pet1, pet2, pet3}}...}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}


func FindPosts(c *gin.Context) {
	var users []models.User
  models.DB.Debug().Find(&users)

  var proposals = make(map[string]int)
  for k, v := range users { 
    fmt.Printf("key[%s] value[%s]\n", k, v.Name)
    proposals[v.Name]=1 
  }
  fmt.Println(proposals)

	c.JSON(http.StatusOK, gin.H{"items": users})
}

/*
func FindPost(c *gin.Context) {
	var post models.Post

	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

type UpdatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	var input UpdatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPost := models.Post{Title: input.Title, Content: input.Content}

	models.DB.Model(&post).Updates(&updatedPost)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	models.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}
*/
