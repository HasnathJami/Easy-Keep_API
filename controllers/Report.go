package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/HasnathJami/go-myWallet-mux-mongoDb/models"
	"github.com/HasnathJami/go-myWallet-mux-mongoDb/utils"
)


func UserReport(w http.ResponseWriter, r *http.Request) {
	var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)
	defer cancel()
	
    reports , _, _ := models.GetReport(ctx)
	utils.ProcessSuccessful(http.StatusOK, "Reports Fetched Successfully", w, reports)
}