package jsonrpc

import (
	"MultiTaiko/pkg/models"
	"bytes"
	"encoding/json"
	"net/http"
)

func setData(method string, params []interface{}, id int) models.JsonData {
	return models.JsonData{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		Id:      id,
	}
}

func sendRequest(method string, params []interface{}) models.JsonResponse {
	const URL = "https://rpc.taiko.xyz"

	data := setData(method, params, 1)

	jsonData, _ := json.Marshal(data)

	resp, _ := http.Post(URL, "application/json", bytes.NewBuffer(jsonData))

	defer resp.Body.Close()

	var response models.JsonResponse
	json.NewDecoder(resp.Body).Decode(&response)

	return response
}

func GetNonce(address string) string {
	const method = "eth_getTransactionCount"
	params := []interface{}{address, "latest"}

	response := sendRequest(method, params)

	return response.Result
}

func GetGasPrice() string {
	const method = "eth_gasPrice"
	params := []interface{}{}

	response := sendRequest(method, params)

	return response.Result
}

func GetGasLimit(address string, value string) string {
	const wethContractAddress = "0xa51894664a773981c6c112c43ce576f315d5b1b6"
	const method = "eth_estimateGas"
	params := []interface{}{
		map[string]interface{}{
			"to":    wethContractAddress,
			"from":  address,
			"value": value,
		},
	}

	response := sendRequest(method, params)

	return response.Result
}
