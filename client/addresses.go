package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maestro-org/go-sdk/models"
	"github.com/maestro-org/go-sdk/utils"
)

func (c *Client) DecodeAddress(address string) (*models.DecodedAddress, error) {
	url := fmt.Sprintf("/addresses/%s/decode/", address)
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
	var decodedAddress models.DecodedAddress
	err = json.NewDecoder(resp.Body).Decode(&decodedAddress)
	if err != nil {
		return nil, err
	}
	return &decodedAddress, nil
}

func (c *Client) AddressTransactionCount(address string) (*models.AddressTransactionCount, error) {
	url := fmt.Sprintf("/addresses/%s/transactions/count", address)
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
	var addressTransactionCount models.AddressTransactionCount
	err = json.NewDecoder(resp.Body).Decode(&addressTransactionCount)
	if err != nil {
		return nil, err
	}
	return &addressTransactionCount, nil
}

func (c *Client) AddressTransactions(
	address string,
	params *utils.Parameters,
) (*models.AddressTransactions, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/addresses/%s/transactions%s", address, formattedParams)
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
	var addressTransactions models.AddressTransactions
	err = json.NewDecoder(resp.Body).Decode(&addressTransactions)
	if err != nil {
		return nil, err
	}
	return &addressTransactions, nil
}

func (c *Client) PaymentCredentialTransactions(
	paymentCredential string,
	params *utils.Parameters,
) (*models.AddressTransactions, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/addresses/cred/%s/transactions%s", paymentCredential, formattedParams)
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
	var addressTransactions models.AddressTransactions
	err = json.NewDecoder(resp.Body).Decode(&addressTransactions)
	if err != nil {
		return nil, err
	}
	return &addressTransactions, nil
}

func (c *Client) UtxoReferencesAtAddress(
	address string,
	params *utils.Parameters,
) (*models.UtxoReferencesAtAddress, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/addresses/%s/utxo_refs%s", address, formattedParams)
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
	var utxoReferencesAtAddress models.UtxoReferencesAtAddress
	err = json.NewDecoder(resp.Body).Decode(&utxoReferencesAtAddress)
	if err != nil {
		return nil, err
	}
	return &utxoReferencesAtAddress, nil
}

func (c *Client) UtxosAtAddress(
	address string,
	params *utils.Parameters,
) (*models.UtxosAtAddress, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/addresses/%s/utxos%s", address, formattedParams)
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
	var utxosAtAddress models.UtxosAtAddress
	err = json.NewDecoder(resp.Body).Decode(&utxosAtAddress)
	if err != nil {
		return nil, err
	}
	return &utxosAtAddress, nil
}

func (c *Client) UtxosAtAddresses(
	addresses []string,
	params *utils.Parameters,
) (*models.UtxosAtAddress, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/addresses/utxos%s", formattedParams)
	body := addresses
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
	var utxosAtAddress models.UtxosAtAddress
	err = json.NewDecoder(resp.Body).Decode(&utxosAtAddress)
	if err != nil {
		return nil, err
	}
	return &utxosAtAddress, nil
}

func (c *Client) UtxosByPaymentCredential(
	paymentCredential string,
	params *utils.Parameters,
) (*models.UtxosAtAddress, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/addresses/cred/%s/utxos%s", paymentCredential, formattedParams)
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
	var utxosAtAddress models.UtxosAtAddress
	err = json.NewDecoder(resp.Body).Decode(&utxosAtAddress)
	if err != nil {
		return nil, err
	}
	return &utxosAtAddress, nil
}
