package v1

import (
	"net/http"

	"github.com/Seunghoon-Oh/cloud-ml-manager/service"
	"github.com/gin-gonic/gin"
)

func GetAllMLServervices(c *gin.Context) {

	notebooks := service.GetNotebooks()
	studios := service.GetStudios()
	pipelines := service.GetPipelines()

	var data []string
	data = append(data, notebooks...)
	data = append(data, studios...)
	data = append(data, pipelines...)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func CreateMLNotebook(c *gin.Context) {
	result := service.PublishMsg("notebook", "create")
	c.String(200, result)
}

func CreateMLStudio(c *gin.Context) {
	result := service.PublishMsg("studio", "create")
	c.String(200, result)
}

func CreateMLPipeline(c *gin.Context) {
	result := service.PublishMsg("pipeline", "create")
	c.String(200, result)
}
