package handler

import (
	"net/http"

	"github.com/go-api/dto"
	"github.com/go-api/util"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.JSON(w, http.StatusMethodNotAllowed, dto.ApiResponse{Message: "failed", Error: "Method not allowed"})
		return
	}

	util.JSON(w, http.StatusOK, dto.ApiResponse{Message: "success", Data: "ok"})
}
