package dto

import (
	"errors"
	"strconv"
	"strings"
	"web3.go/web3/complex/types"
	"web3.go/web3/constants"

	"encoding/json"
	"fmt"
	"math/big"
)

type RequestResult struct {
	// ID      int         `json:"id"`
	// Version string      `json:"jsonrpc"`
	Result interface{} `json:"result"`
	Error  *Error      `json:"error,omitempty"`
	Data   string      `json:"data,omitempty"`
}

type BalanceResult struct {
	Address     string        `json:"address"`
	Nonce       int           `json:"nonce"`
	Balance     types.Uint128 `json:"balance"`
	StorageRoot string        `json:"storageRoot"`
	CodeHash    string        `json:"codeHash"`
}

type SendTxResult struct {
	TXhash string `json:"TXhash,omitempty"`
	ErrMsg string `json:"ErrMsg,omitempty"`
}

type TransactionResult struct {
	ChainId int    `json:"chainId"`
	From    string `json:"from"`
	To      string `json:"to"`
	Nonce   int    `json:"nonce"`
	Value   int    `json:"value"`
	Input   string `json:"input"`
}

type TxResult struct {
	Transaction     TransactionResult
	Root            string `json:"root"`
	Status          int    `json:"status"`
	Logs            string `json:"logs"`
	TransactionHash string `json:"transactionHash"`
	ContractAddress string `json:"contractAddress"`
	Out             string `json:"out"`
	BlockHeight     int    `json:"blockHeight"`
	ErrMsg          string `json:"ErrMsg,omitempty"`
}

type GetBlockResult struct {
	Hash         string `json:"hash"`
	Previoushash string `json:"previoushash"`
	Chainid      int    `json:"chainid"`
	Height       int    `json:"height"`
	Mergeroot    string `json:"mergeroot"`
	Deltaroot    string `json:"deltaroot"`
	Stateroot    string `json:"stateroot"`
	Txcount      int    `json:"txcount"`
	Timestamp    int64  `json:"timestamp"`
	ErrMsg       string `json:"ErrMsg,Omitempty"`
}

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (pointer *RequestResult) ToStringArray() ([]string, error) {

	if err := pointer.checkResponse(); err != nil {
		return nil, err
	}

	result := (pointer).Result.([]interface{})

	new := make([]string, len(result))
	for i, v := range result {
		new[i] = v.(string)
	}

	return new, nil

}

func (pointer *RequestResult) ToComplexString() (types.ComplexString, error) {

	if err := pointer.checkResponse(); err != nil {
		return "", err
	}

	result := (pointer).Result.(interface{})

	return types.ComplexString(result.(string)), nil

}

func (pointer *RequestResult) ToString() (string, error) {

	if err := pointer.checkResponse(); err != nil {
		return "", err
	}

	result := (pointer).Result.(interface{})

	return result.(string), nil

}

func (pointer *RequestResult) ToInt() (int64, error) {

	if err := pointer.checkResponse(); err != nil {
		return 0, err
	}

	result := (pointer).Result.(interface{})

	hex := result.(string)

	numericResult, err := strconv.ParseInt(hex, 16, 64)

	return numericResult, err

}

func (pointer *RequestResult) ToBigInt() (*big.Int, error) {

	if err := pointer.checkResponse(); err != nil {
		return nil, err
	}

	res := (pointer).Result.(interface{})

	ret, success := big.NewInt(0).SetString(res.(string)[2:], 16)

	if !success {
		return nil, errors.New(fmt.Sprintf("Failed to convert %s to BigInt", res.(string)))
	}

	return ret, nil
}

func (pointer *RequestResult) ToComplexIntResponse() (types.ComplexIntResponse, error) {

	if err := pointer.checkResponse(); err != nil {
		return types.ComplexIntResponse(0), err
	}

	result := (pointer).Result.(interface{})

	var hex string

	switch v := result.(type) {
	// Testrpc returns a float64
	case float64:
		hex = strconv.FormatFloat(v, 'E', 16, 64)
		break
	default:
		hex = result.(string)
	}

	cleaned := strings.TrimPrefix(hex, "0x")

	return types.ComplexIntResponse(cleaned), nil

}

func (pointer *RequestResult) ToBoolean() (bool, error) {

	if err := pointer.checkResponse(); err != nil {
		return false, err
	}

	result := (pointer).Result.(interface{})

	return result.(bool), nil

}

func (pointer *RequestResult) ToSignTransactionResponse() (*SignTransactionResponse, error) {
	if err := pointer.checkResponse(); err != nil {
		return nil, err
	}

	result := (pointer).Result.(map[string]interface{})

	if len(result) == 0 {
		return nil, customerror.EMPTYRESPONSE
	}

	signTransactionResponse := &SignTransactionResponse{}

	marshal, err := json.Marshal(result)

	if err != nil {
		return nil, customerror.UNPARSEABLEINTERFACE
	}

	err = json.Unmarshal([]byte(marshal), signTransactionResponse)

	return signTransactionResponse, err
}

func (pointer *RequestResult) ToTransactionResponse() (*TransactionResponse, error) {

	if err := pointer.checkResponse(); err != nil {
		return nil, err
	}

	result := (pointer).Result.(map[string]interface{})

	if len(result) == 0 {
		return nil, customerror.EMPTYRESPONSE
	}

	transactionResponse := &TransactionResponse{}

	marshal, err := json.Marshal(result)

	if err != nil {
		return nil, customerror.UNPARSEABLEINTERFACE
	}

	err = json.Unmarshal([]byte(marshal), transactionResponse)

	return transactionResponse, err

}

func (pointer *RequestResult) ToTransactionReceipt() (*TransactionReceipt, error) {

	if err := pointer.checkResponse(); err != nil {
		return nil, err
	}

	result := (pointer).Result.(map[string]interface{})

	if len(result) == 0 {
		return nil, customerror.EMPTYRESPONSE
	}

	transactionReceipt := &TransactionReceipt{}

	marshal, err := json.Marshal(result)

	if err != nil {
		return nil, customerror.UNPARSEABLEINTERFACE
	}

	err = json.Unmarshal([]byte(marshal), transactionReceipt)

	return transactionReceipt, err

}

func (pointer *RequestResult) ToBlock() (*Block, error) {

	if err := pointer.checkResponse(); err != nil {
		return nil, err
	}

	result := (pointer).Result.(map[string]interface{})

	if len(result) == 0 {
		return nil, customerror.EMPTYRESPONSE
	}

	block := &Block{}

	marshal, err := json.Marshal(result)

	if err != nil {
		return nil, customerror.UNPARSEABLEINTERFACE
	}

	err = json.Unmarshal([]byte(marshal), block)

	return block, err

}

func (pointer *RequestResult) ToSyncingResponse() (*SyncingResponse, error) {

	if err := pointer.checkResponse(); err != nil {
		return nil, err
	}

	var result map[string]interface{}

	switch (pointer).Result.(type) {
	case bool:
		return &SyncingResponse{}, nil
	case map[string]interface{}:
		result = (pointer).Result.(map[string]interface{})
	default:
		return nil, customerror.UNPARSEABLEINTERFACE
	}

	if len(result) == 0 {
		return nil, customerror.EMPTYRESPONSE
	}

	syncingResponse := &SyncingResponse{}

	marshal, err := json.Marshal(result)

	if err != nil {
		return nil, customerror.UNPARSEABLEINTERFACE
	}

	json.Unmarshal([]byte(marshal), syncingResponse)

	return syncingResponse, nil

}

// To avoid a conversion of a nil interface
func (pointer *RequestResult) checkResponse() error {

	if pointer.Error != nil {
		return errors.New(pointer.Error.Message)
	}

	if pointer.Result == nil {
		return customerror.EMPTYRESPONSE
	}

	return nil

}
