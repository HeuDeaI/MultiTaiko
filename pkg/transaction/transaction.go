package transaction

import (
	"MultiTaiko/pkg/data"
	"MultiTaiko/pkg/jsonrpc"
	"MultiTaiko/pkg/rawencoding"
	"encoding/hex"
)

func setTransactionData(address string, value string) data.TransactionData {

	dataBytes := []byte{0xd0, 0xe3, 0x0d, 0xb0}
	dataString := hex.EncodeToString(dataBytes)

	return data.TransactionData{
		Nonce:    jsonrpc.GetNonce(address),
		GasPrice: jsonrpc.GetGasPrice(),
		GasLimit: jsonrpc.GetGasLimit(address, value),
		To:       data.WethContractAddress,
		Value:    value,
		Data:     dataString,
	}
}
