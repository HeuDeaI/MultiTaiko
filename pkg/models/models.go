package models

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

type TransactionData struct {
	Nonce    string
	GasPrice string
	GasLimit string
	To       string
	Value    string
	Data     string
}
