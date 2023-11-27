package server

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"github.com/unawaretub86/club-hub/config"
	web "github.com/unawaretub86/club-hub/src/adapters/http"
	"github.com/unawaretub86/club-hub/src/adapters/repository"
	"github.com/unawaretub86/club-hub/src/adapters/rest"
	"github.com/unawaretub86/club-hub/src/core/ports"
	"github.com/unawaretub86/club-hub/src/core/services"
)

type Service struct {
	clubHub ports.ClubHubPort
	rest ports.ScrapperPort
}

type Initiator struct {
	db      *gorm.DB
	service Service
	router  *gin.Engine
}

func NewInitiator() *Initiator {
	return &Initiator{}
}

func (initiator *Initiator) InitDB() {
	dbConfig := config.GetDatabaseConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.DBName,
		// dbConfig.DatabasePort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connect error:" + err.Error())
	}

	initiator.db = db
}

func (initiator *Initiator) InitService() {
	clubHubRepo := repository.NewClubRepository(initiator.db)
	clubHubRest := rest.NewScrapper(initiator.service.rest)
	clubHubService := services.NewClubHubService(clubHubRepo, clubHubRest)

	initiator.service = Service{
		clubHubService,
		clubHubRest,
	}
}

func (initiator *Initiator) InitRouter() {
	initiator.router = gin.Default()

	clubRouter := web.NewRouter(initiator.service.clubHub)
	clubRouter.SetRoutes(initiator.router)
}

func (initiator *Initiator) GetRouter() *gin.Engine {
	return initiator.router
}

func (initiator *Initiator) GetDB() *gorm.DB {
	return initiator.db
}
