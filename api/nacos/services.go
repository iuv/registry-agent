package nacos

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
)

var serviceMap = make(map[string][]Service)

func addService(s Service) {
	// add or update
	services := serviceMap[s.ServiceName]
	if len(services) > 0 {
		for i, service := range services {
			if s.InstanceId == service.InstanceId {
				services[i] = s
				return
			}
		}
		serviceMap[s.ServiceName] = append(services, s)
	} else {
		serviceMap[s.ServiceName] = []Service{s}
	}
}

func delService(s Service) {
	// add or update
	services := serviceMap[s.ServiceName]
	if len(services) > 0 {
		for i, service := range services {
			if s.InstanceId == service.InstanceId {
				serviceMap[s.ServiceName] = append(services[:i], services[i+1:]...)
				return
			}
		}
	}
}

// get service list
func ServiceList(c *gin.Context) {
	serviceName := c.Query("serviceName")
	log.Printf(serviceName)
	if serviceName == "" {
		c.String(http.StatusOK, "caused: Param 'serviceName' is required.;")
		return
	}
	var respService RespService
	respService = GetDefaultRespService(serviceName)
	c.JSON(http.StatusOK, respService)
}

func getNacosUrl(path string) string {
	return "http://" + viper.GetString("nacos.url") + "/nacos/v1" + path
}
func Register(c *gin.Context) {
	var s Service
	if err := c.ShouldBindQuery(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.InstanceId = s.Ip + "#" + strconv.Itoa(s.Port) + "#" + s.ClusterName + "#" + s.ServiceName
	addService(s)
	c.String(http.StatusOK, "ok")
}

// service health
func Beat(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
func DeRegister(c *gin.Context) {
	var s Service
	if err := c.ShouldBindQuery(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.InstanceId = s.Ip + "#" + strconv.Itoa(s.Port) + "#" + s.ClusterName + "#" + s.ServiceName
	delService(s)
	c.String(http.StatusOK, "ok")
}
