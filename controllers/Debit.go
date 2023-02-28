package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/HasnathJami/go-myWallet-mux-mongoDb/models"
	"github.com/HasnathJami/go-myWallet-mux-mongoDb/utils"
)

func init(){
   isExecutedOnce = false
}

func DebitMoney(w http.ResponseWriter, r *http.Request) {
    var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)
	defer cancel()

	var debit models.Debit
	_ = json.NewDecoder(r.Body).Decode(&debit)

	
	if totalMoney <=0 || totalMoney < debit.Amount {
		//totalMoney = 0
		json.NewEncoder(w).Encode("Insufficient Balance")
		return

	} else{
        totalMoney -= debit.Amount
	}
	var  report models.Report
	report.TotalAmount = totalMoney
	//_, updateError := models.UpdateReport(ctx, &report)
	var createError error
	var updateError error

	if !isExecutedOnce{
       _ , createError = models.CreateReport(ctx, &report)
	   utils.CheckError(createError, "Money Has Not Debited Successfully", http.StatusInternalServerError, w)
	   isExecutedOnce = true
	}else{
		_, updateError = models.UpdateReport(ctx, &report)
		utils.CheckError(updateError, "Money Has Not Debited Successfully", http.StatusInternalServerError, w)
	}
	//debit.DebitId = primitive.NewObjectID()
	//_, insertError := models.DebitMoney(ctx, &debit)
	utils.ProcessSuccessful(http.StatusOK, "Money Has Debited Successfully", w, nil)
}