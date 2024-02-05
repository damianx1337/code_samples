package healthcheck-controller

import (
	"log"
	"net/http"

	"go-skeleton/database"

	"github.com/gin-gonic/gin"
)

func Livez (c *gin.Context){
	log.Println(c.Request.Header.Get("Origin"), c.Request.Host, c.Request.RemoteAddr, c.Request.RequestURI)

	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Println(err)
	}

	pingResult := sqlDB.Ping()
	if pingResult != nil {
		sqlDB.Close()
		c.JSON(http.StatusOK, gin.H{"state": "died"})
	}
	c.JSON(http.StatusOK, gin.H{"state": "livez"})
}

func Ready (c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"state": "ready"})
}
