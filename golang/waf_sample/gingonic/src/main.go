package main

import (
//		"log"
		"net/http"

		"github.com/gin-gonic/gin"
		
    "github.com/corazawaf/coraza/v2"
		"github.com/corazawaf/coraza/v2/seclang"
    corazagin "github.com/jptosso/coraza-gin"
)
func main() {
    r := gin.New()

    waf := coraza.NewWaf()
		parser, _ := seclang.NewParser(waf)
   	parser.FromString(`SecDebugLogLevel 9`)
   	parser.FromString(`SecDebugLog /dev/stdout`)
   	parser.FromString(`SecRule ARGS:id "@eq 0" "id:1, phase:1,deny, status:403,msg:'Invalid id',log,auditlog"`)
   	parser.FromString(`SecRequestBodyAccess On`)
   	parser.FromString(`SecRule REQUEST_BODY "@contains password" "id:100, phase:2,deny, status:403,msg:'Invalid request body',log,auditlog"`)

    r.Use(corazagin.Coraza(waf))

		r.GET("/", func(c *gin.Context) {
   		c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  	})

    r.Run(":8080")
}
