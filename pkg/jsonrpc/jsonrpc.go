package jsonrpc

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type JsonData struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

type JsonResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      string `json:"id"`
	Result  string `json:"result"`
}

func setData(method string, params []interface{}, id int) JsonData {
	return JsonData{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		Id:      id,
	}
}

func sendRequest(method string, params []interface{}) JsonResponse {
	const URL = "https://rpc.taiko.xyz"

	data := setData(method, params, 1)

	jsonData, _ := json.Marshal(data)

	resp, _ := http.Post(URL, "application/json", bytes.NewBuffer(jsonData))

	defer resp.Body.Close()

	var response JsonResponse
	json.NewDecoder(resp.Body).Decode(&response)

	return response
}

func getNonce(address string) string {
	const method = "eth_getTransactionCount"
	params := []interface{}{address, "latest"}

	response := sendRequest(method, params)

	return response.Result
}

func getGasPrice() string {
	const method = "eth_gasPrice"
	params := []interface{}{}

	response := sendRequest(method, params)

	return response.Result
}

func getGasLimit(address string, value string) string {
	const ethWrapContact = "0xa51894664a773981c6c112c43ce576f315d5b1b6"
	const method = "eth_estimateGas"
	params := []interface{}{
		map[string]interface{}{
			"to":    ethWrapContact,
			"from":  address,
			"value": value,
		},
	}

	response := sendRequest(method, params)

	return response.Result
}
