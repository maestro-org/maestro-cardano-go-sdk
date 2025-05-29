package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maestro-org/go-sdk/models"
	"github.com/maestro-org/go-sdk/utils"
)

func (c *Client) AccountsHoldingPolicy(
	policy string,
	params *utils.Parameters,
) (*models.AccountsHoldingPolicy, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/assets/policy/%s/accounts?%s", policy, formattedParams)
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
	var response models.AccountsHoldingPolicy
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (c *Client) AddressesHoldingPolicy(
	policy string,
	params *utils.Parameters,
) (*models.AddressesHoldingPolicy, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/assets/policy/%s/addresses?%s", policy, formattedParams)
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
	var response models.AddressesHoldingPolicy
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (c *Client) SpecificPolicyInformations(
	policy string,
	params *utils.Parameters,
) (*models.PolicyInformation, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/assets/policy/%s?%s", policy, formattedParams)
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
	var response models.PolicyInformation
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (c *Client) TransactionsMovingPolicy(
	policy string,
	params *utils.Parameters,
) (*models.PolicyTransactions, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/assets/policy/%s/txs?%s", policy, formattedParams)
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
	var response models.PolicyTransactions
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (c *Client) UtxosContainingPolicy(
	policy string,
	params *utils.Parameters,
) (*models.UtxosContainingPolicy, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/assets/policy/%s/utxos?%s", policy, formattedParams)
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
	var response models.UtxosContainingPolicy
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, err
}
