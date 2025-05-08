package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/marcelochb/training-go-payment-gateway/internal/dto"
	"github.com/marcelochb/training-go-payment-gateway/internal/service"
)

type AccountHandler struct {
	service *service.AccountService
}

func NewAccountHandler(service *service.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

func (handler *AccountHandler) Create(response http.ResponseWriter, request *http.Request) {
	var input dto.AccountDtoInput
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := handler.service.CreateAccount(input)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(output)

}

func (handler *AccountHandler) Get(response http.ResponseWriter, request *http.Request) {
	apikey := request.Header.Get("X-API-KEY")
	if apikey == "" {
		http.Error(response, "API Key is required", http.StatusUnauthorized)
		return
	}

	output, err := handler.service.FindByAPIKey(apikey)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(output)

}
