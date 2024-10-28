package transaction

import (
	"MultiTaiko/pkg/jsonrpc"
	"MultiTaiko/pkg/models"
	"MultiTaiko/pkg/rawencoding"
	"encoding/hex"
)

func setTransactionData(address string, value string) models.TransactionData {
	const wethContractAddress = "0xa51894664a773981c6c112c43ce576f315d5b1b6"

	dataBytes := []byte{0xd0, 0xe3, 0x0d, 0xb0}
	dataString := hex.EncodeToString(dataBytes)

	return models.TransactionData{
		Nonce:    jsonrpc.GetNonce(address),
		GasPrice: jsonrpc.GetGasPrice(),
		GasLimit: jsonrpc.GetGasLimit(address, value),
		To:       wethContractAddress,
		Value:    value,
		Data:     dataString,
	}
}

func prepareTransactionWithoutSign(address string, value string) {
	txData := setTransactionData(address, value)
	rawencoding.HashTransaction(txData)
}
