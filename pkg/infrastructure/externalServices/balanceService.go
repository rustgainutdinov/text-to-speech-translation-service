package externalServices

import (
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"strconv"
	"text-to-speech-translation-service/pkg/app/service"
)

const canWriteOffUrl = "/api/v1/balance/canWriteOff"

type balanceService struct {
	balanceServiceAddress string
}

func (b *balanceService) CanWriteOf(userID uuid.UUID, score int) (bool, error) {
	resp, err := http.Get("http://" + b.balanceServiceAddress + canWriteOffUrl + "?userID=" + userID.String() + "&score=" + strconv.Itoa(score))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var response canWriteOffResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false, err
	}
	return response.Result, nil
}

func NewBalanceService(balanceServiceAddress string) service.BalanceService {
	return &balanceService{balanceServiceAddress: balanceServiceAddress}
}

type canWriteOffResponse struct {
	Result bool `json:"result"`
}
