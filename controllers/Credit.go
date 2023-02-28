package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/HasnathJami/go-myWallet-mux-mongoDb/models"
	"github.com/HasnathJami/go-myWallet-mux-mongoDb/utils"
)

var (
	totalMoney float64 = 0.0
	isExecutedOnce bool = true
)

func init(){
   isExecutedOnce = false
}

func CreditMoney(w http.ResponseWriter, r *http.Request) {
    var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)
	defer cancel()

	var credit models.Credit
	_ = json.NewDecoder(r.Body).Decode(&credit)

	totalMoney += credit.Amount
	//credit.CreditId = primitive.NewObjectID()
	//_, insertError := models.CreditMoney(ctx, &credit)
	var report models.Report
	report.TotalAmount = totalMoney
	fmt.Println(report)
	//_, updateError := models.UpdateReport(ctx, &report)
	var createError error
	var updateError error
	if !isExecutedOnce{
       _ , createError = models.CreateReport(ctx, &report)
	   utils.CheckError(createError, "Money Has Not Credited Successfully", http.StatusInternalServerError, w)
	   isExecutedOnce = true
	}else{
		_, updateError = models.UpdateReport(ctx, &report)
		utils.CheckError(updateError, "Money Has Not Credited Successfully", http.StatusInternalServerError, w)
	}
	utils.ProcessSuccessful(http.StatusOK, "Money Has Credited Successfully", w, nil)
}