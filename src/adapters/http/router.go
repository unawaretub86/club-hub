package web

import (
	"github.com/gin-gonic/gin"
	"github.com/unawaretub86/club-hub/src/core/ports"
)

type ClubHubRouter struct {
	clubHubService ports.ClubHubPort
}

func NewRouter(clubHubServicePorts ports.ClubHubPort) *ClubHubRouter {
	return &ClubHubRouter{
		clubHubService: clubHubServicePorts,
	}
}

func (r *ClubHubRouter) SetRoutes(g *gin.Engine) {
	group := g.Group("/company")

	// read
	group.GET("/", r.GetCompany)
	group.GET("/franchise", r.GetCompanyByFranchise)
	group.GET("/owner", r.GetCompanyByOwner)
	group.GET("/information", r.GetCompanyByInformation)

	// write
	group.PATCH("/:id", r.UpdateCompany)
	group.POST("/", r.SaveCompany)
}
