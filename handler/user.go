package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-api/dto"
	"github.com/go-api/model"
	"github.com/go-api/util"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		DB: db,
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.JSON(w, http.StatusMethodNotAllowed, dto.ApiResponse{Message: "failed", Error: "Method not allowed"})
		return
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	if limit < 1 {
		limit = 10
	}
	log.Println("limit", limit, "offset", offset)

	var users []model.User
	h.DB.Limit(limit).Offset(offset).Find(&users)
	var total int64
	h.DB.Model(&users).Count(&total)
	util.JSON(w, http.StatusOK, dto.ApiResponse{Message: "success", Data: users, Limit: &limit, Offset: &offset, Total: &total})
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.JSON(w, http.StatusMethodNotAllowed, dto.ApiResponse{Message: "failed", Error: "Method not allowed"})
		return
	}

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		util.JSON(w, http.StatusBadRequest, dto.ApiResponse{Message: "failed", Error: err.Error()})
		return
	}

	if err := h.DB.Create(&user).Error; err != nil {
		util.JSON(w, http.StatusInternalServerError, dto.ApiResponse{Message: "failed", Error: err.Error()})
		return
	}

	util.JSON(w, http.StatusCreated, dto.ApiResponse{Message: "success", Data: user})
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.JSON(w, http.StatusMethodNotAllowed, dto.ApiResponse{Message: "failed", Error: "Method not allowed"})
		return
	}

	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var user model.User
	if err := h.DB.First(&user, id).Error; err != nil {
		util.JSON(w, http.StatusNotFound, dto.ApiResponse{Message: "failed", Error: "User not found"})
		return
	}

	util.JSON(w, http.StatusOK, dto.ApiResponse{Message: "success", Data: user})
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		util.JSON(w, http.StatusMethodNotAllowed, dto.ApiResponse{Message: "failed", Error: "Method not allowed"})
		return
	}

	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var user model.User
	if err := h.DB.First(&user, id).Error; err != nil {
		util.JSON(w, http.StatusNotFound, dto.ApiResponse{Message: "failed", Error: "User not found"})
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		util.JSON(w, http.StatusBadRequest, dto.ApiResponse{Message: "failed", Error: err.Error()})
		return
	}

	h.DB.Save(&user)
	util.JSON(w, http.StatusOK, dto.ApiResponse{Message: "success", Data: user})
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		util.JSON(w, http.StatusMethodNotAllowed, dto.ApiResponse{Message: "failed", Error: "Method not allowed"})
		return
	}

	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	if err := h.DB.Delete(&model.User{}, id).Error; err != nil {
		util.JSON(w, http.StatusInternalServerError, dto.ApiResponse{Message: "failed", Error: err.Error()})
		return
	}

	util.JSON(w, http.StatusOK, dto.ApiResponse{Message: "success", Data: "User deleted"})
}
