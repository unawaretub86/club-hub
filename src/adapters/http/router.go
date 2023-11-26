package web

import (
	"github.com/gin-gonic/gin"
	"github.com/unawaretub86/club-hub/src/core/ports"
)

type ClubHubRouter struct {
	clubHubService ports.WebPort
}

func NewRouter(clubHubServicePorts ports.WebPort) *ClubHubRouter {
	return &ClubHubRouter{
		clubHubService: clubHubServicePorts,
	}
}

func (r *ClubHubRouter) SetRoutes(g *gin.Engine) {
	group := g.Group("/v1")

	// read
	group.GET("/company", r.GetCompany)

	// write
	group.PATCH("/:id", r.UpdateCompany)
	group.POST("/", r.SaveCompany)
}
