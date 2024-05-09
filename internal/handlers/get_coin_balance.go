package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/mallett002/example-go-api/api"
	"github.com/mallett002/example-go-api/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(writer http.ResponseWriter, req *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, req.URL.Query()) // grab values off url query params and populate our params struct (username)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(writer)
		return
	}

	var coinDetails *tools.CoinDetails
	coinDetails = (*database).GetUserCoins(params.Username)
	if coinDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*&coinDetails).Coins,
		Code: http.StatusOK,
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}
}
