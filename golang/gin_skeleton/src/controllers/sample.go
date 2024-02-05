package sample-controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// @Summary Ping
// @Description ping the API endpoint
// @ID ping
// @Accept x-www-form-urlencoded
// @Produce json
// @Param name formData string true "input = name"
// @Success 200 {object} inputs.TextInput
// @Router /v1/ping [post]
func Ping (c *gin.Context) {
	var input inputs.TextInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, input.Name)
}
