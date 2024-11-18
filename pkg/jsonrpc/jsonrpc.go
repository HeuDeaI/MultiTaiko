package jsonrpc

import (
	"MultiTaiko/pkg/data"
	"bytes"
	"encoding/json"
	"net/http"
)

func setData(method string, params []interface{}, id int) data.JsonData {
	return data.JsonData{
		Jsonrpc: data.JsonRpcVersion,
		Method:  method,
		Params:  params,
		Id:      id,
	}
}

func sendRequest(method string, params []interface{}) data.JsonResponse {
	const URL = "https://rpc.taiko.xyz"

	requestData := setData(method, params, 1)

	jsonData, _ := json.Marshal(requestData)

	resp, _ := http.Post(URL, "application/json", bytes.NewBuffer(jsonData))
	defer resp.Body.Close()

	var response data.JsonResponse
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
	const method = "eth_estimateGas"
	params := []interface{}{
		map[string]interface{}{
			"to":    data.WethContractAddress,
			"from":  address,
			"value": value,
		},
	}

	response := sendRequest(method, params)

	return response.Result
}
