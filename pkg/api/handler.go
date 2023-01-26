package api

import (
	"context"
	"net/http"
	"orm_gorm/pkg/config"
	"orm_gorm/pkg/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func NewDB(db *gorm.DB) handler {
	return handler{DB: db}
}

const appTimeout = time.Second * 10

func (h *handler) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload models.User
		defer cancel()

		config.ValidateBody(ctx, &payload)

		newData := models.User{
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Title:     payload.Title,
		}

		data := h.DB.Create(&newData)
		if data.Error != nil {
			config.ErrorJSON(ctx, data.Error)
			return
		}

		config.WriteJSON(ctx, http.StatusOK, data)
	}
}

func (h *handler) GetAUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		userId := ctx.Param("userId")
		var payload models.User
		defer cancel()

		data := h.DB.First(&payload, userId)
		if data.Error != nil {
			config.ErrorJSON(ctx, data.Error)
			return
		}

		config.WriteJSON(ctx, http.StatusOK, data)
	}
}

func (h *handler) EditAUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		userId := ctx.Param("userId")
		var payload models.User
		defer cancel()

		config.ValidateBody(ctx, &payload)
		
		//find user
		if result := h.DB.First(&payload, userId); result.Error != nil {
			config.ErrorJSON(ctx, result.Error)
			return
		}

		//save user
		updatedData := models.User{
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Title:     payload.Title,
		}
		h.DB.Save(&updatedData)

		//get updated user
		data := h.DB.First(&updatedData, userId)
		if data.Error != nil {
			config.ErrorJSON(ctx, data.Error)
			return
		}

		config.WriteJSON(ctx, http.StatusOK, data)
	}
}