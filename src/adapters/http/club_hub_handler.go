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

	company, err := r.clubHubService.GetCompany(filterFields)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixCompany, err)
		return
	}

	utils.EndWithStatus(c, http.StatusOK, suffixCompany, company)
}

func (r *ClubHubRouter) GetCompanyByFranchise(c *gin.Context) {
	filterFields := utils.ExtractAllParams(c)

	company, err := r.clubHubService.GetCompanyByFranchise(filterFields)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixCompany, err)
		return
	}

	utils.EndWithStatus(c, http.StatusOK, suffixCompany, company)
}

func (r *ClubHubRouter) GetCompanyByInformation(c *gin.Context) {
	filterFields := utils.ExtractAllParams(c)

	company, err := r.clubHubService.GetCompanyByInformation(filterFields)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixCompany, err)
		return
	}

	utils.EndWithStatus(c, http.StatusOK, suffixCompany, company)
}

func (r *ClubHubRouter) GetCompanyByOwner(c *gin.Context) {
	filterFields := utils.ExtractAllParams(c)

	company, err := r.clubHubService.GetCompanyByOwner(filterFields)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixCompany, err)
		return
	}

	utils.EndWithStatus(c, http.StatusOK, suffixCompany, company)
}

func (r *ClubHubRouter) SaveCompany(c *gin.Context) {
	companyData := &domain.ReqData{}

	if err := c.ShouldBindJSON(&companyData); err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixCompany, err)
		return
	}

	company, err := r.clubHubService.SaveCompany(companyData.Company)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixCompany, err)
		return
	}

	utils.EndWithStatus(c, http.StatusOK, suffixCompany, company)
}

func (r *ClubHubRouter) UpdateCompany(c *gin.Context) {
	companyData := &domain.ReqData{}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	if err := c.ShouldBindJSON(&companyData); err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixCompany, err)
		return
	}

	company, err := r.clubHubService.UpdateCompany(uint(id), companyData.Company)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	utils.EndWithStatus(c, http.StatusOK, suffixCompany, company)
}
