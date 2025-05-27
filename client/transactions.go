package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maestro-org/go-sdk/models"
	"github.com/maestro-org/go-sdk/utils"
)

func (c *Client) AddressByOutputReference(txHash string, index int) (*models.BasicResponse, error) {
	url := fmt.Sprintf("/transactions/%s/outputs/%d/address", txHash, index)
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("empty response")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected error: %d", resp.Body)
	}
	defer resp.Body.Close() //nolint:errcheck
	var addressByOutputReference models.BasicResponse
	err = json.NewDecoder(resp.Body).Decode(&addressByOutputReference)
	if err != nil {
		return nil, err
	}
	return &addressByOutputReference, nil

}

func (c *Client) SubmitTx(cbor string) (models.BasicResponse, error) {
	url := "/submitmodels.BasicResponse{}/tx"
	resp, err := c.post(url, cbor)
	if err != nil {
		return models.BasicResponse{}, err
	}
	if resp == nil {
		return models.BasicResponse{}, fmt.Errorf("empty response")
	}
	if resp.StatusCode != http.StatusOK {
		return models.BasicResponse{}, fmt.Errorf("unexpected error: %d", resp.Body)
	}
	defer resp.Body.Close() //nolint:errcheck
	var submitTx models.BasicResponse

	err = json.NewDecoder(resp.Body).Decode(&submitTx)
	if err != nil {
		return models.BasicResponse{}, err
	}
	return submitTx, nil
}

func (c *Client) TransactionCbor(txHash string) (*models.BasicResponse, error) {
	url := fmt.Sprintf("/transactions/%s/cbor", txHash)
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("empty response")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected error: %d", resp.Body)
	}
	defer resp.Body.Close() //nolint:errcheck
	var transactionCbor models.BasicResponse
	err = json.NewDecoder(resp.Body).Decode(&transactionCbor)
	if err != nil {
		return nil, err
	}
	return &transactionCbor, nil

}

func (c *Client) TransactionDetails(txHash string) (*models.TransactionDetails, error) {
	url := fmt.Sprintf("/transactions/%s", txHash)
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("empty response")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected error: %d", resp.Body)
	}
	defer resp.Body.Close() //nolint:errcheck
	var transactionDetails models.TransactionDetails
	err = json.NewDecoder(resp.Body).Decode(&transactionDetails)
	if err != nil {
		return nil, err
	}
	return &transactionDetails, nil

}

func (c *Client) TransactionOutputFromReference(
	txHash string,
	index int,
	params *utils.Parameters,
) (*models.TransactionOutputFromReference, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/transactions/%s/outputs/%d/txo%s", txHash, index, formattedParams)
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("empty response")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected error: %d", resp.Body)
	}
	defer resp.Body.Close() //nolint:errcheck
	var transactionOutputFromReference models.TransactionOutputFromReference
	err = json.NewDecoder(resp.Body).Decode(&transactionOutputFromReference)
	if err != nil {
		return nil, err
	}
	return &transactionOutputFromReference, nil
}

func (c *Client) TransactionOutputsFromReferences(
	references []models.TxoReference,
	params *utils.Parameters,
) (*models.TransactionOutputsFromReferences, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}

	url := "/transactions/outputs" + formattedParams
	resp, err := c.post(url, references)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("empty response")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected error: %d", resp.Body)
	}
	defer resp.Body.Close() //nolint:errcheck
	var transactionOutputsFromReferences models.TransactionOutputsFromReferences
	err = json.NewDecoder(resp.Body).Decode(&transactionOutputsFromReferences)
	if err != nil {
		return nil, err
	}
	return &transactionOutputsFromReferences, nil

}

func (c *Client) EvaluateTx(
	txCbor string,
	AdditionalUtxos ...string,
) (models.EvaluateTxResponse, error) {
	url := "/transactions/evaluate"
	body := models.EvaluateTx{
		Cbor:            txCbor,
		AdditionalUtxos: AdditionalUtxos,
	}
	resp, err := c.post(url, body)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("empty response")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected error: %d", resp.Body)
	}
	defer resp.Body.Close() //nolint:errcheck
	var redeemerEvals models.EvaluateTxResponse
	err = json.NewDecoder(resp.Body).Decode(&redeemerEvals)
	if err != nil {
		return nil, err
	}
	return redeemerEvals, nil
}
