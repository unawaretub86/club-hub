package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/unawaretub86/club-hub/src/core/domain"
	"github.com/unawaretub86/club-hub/utils"
)

const (
	suffixErr     = "Error"
	suffixCompany = "Company"
)

func (r *ClubHubRouter) GetCompany(c *gin.Context) {
	filterFields := utils.ExtractAllParams(c)

	company, err := r.clubHubService.Get(filterFields)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixCompany, err)
		return
	}

	utils.EndWithStatus(c, http.StatusOK, suffixCompany, company)
}

func (r *ClubHubRouter) CreateCompany(c *gin.Context) {
	companyData := &domain.Company{}

	err := c.ShouldBindJSON(&companyData)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixCompany, err)
		return
	}

	company, err := r.clubHubService.Save(*companyData)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixCompany, err)
		return
	}

	utils.EndWithStatus(c, http.StatusOK, suffixCompany, company)
}

func (r *ClubHubRouter) UpdateCompany(c *gin.Context) {
	companyData := &domain.Company{}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	err = c.ShouldBindJSON(&companyData)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	company, err := r.clubHubService.Update(uint(id), *companyData)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	utils.EndWithStatus(c, http.StatusOK, suffixCompany, company)
}
