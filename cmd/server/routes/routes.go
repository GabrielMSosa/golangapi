package routes

import (
	"GabrielMSosa/crud-api/cmd/server/handler/seller"
	"GabrielMSosa/crud-api/cmd/server/middleware"
	"GabrielMSosa/crud-api/internal/repository"
	"GabrielMSosa/crud-api/internal/service"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *sql.DB
}

func NewRouter(eng *gin.Engine, db *sql.DB) Router {
	return &router{
		eng: eng,
		db:  db,
	}
}
func (r *router) MapRoutes() {
	r.setGroup()
	r.buildSellerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.eng.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {
	repo := repository.NewRepository(r.db)
	service := service.NewService(repo)
	handler := seller.NewSeller(service)
	gr := r.rg.Group("/sellers", middleware.LoggerMIddleware())
	gr.POST("", handler.Create())
	gr.GET("/:id", handler.GetById())
	gr.GET("", handler.GetAll())
}
