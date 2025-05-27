package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maestro-org/go-sdk/models"
	"github.com/maestro-org/go-sdk/utils"
)

func (c *Client) AccountsHoldingAsset(
	assetId string,
	params *utils.Parameters,
) (*models.AccountsHoldingAsset, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/assets/%s/accounts%s", assetId, formattedParams)
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
	var accountsHoldingAsset models.AccountsHoldingAsset
	err = json.NewDecoder(resp.Body).Decode(&accountsHoldingAsset)
	if err != nil {
		return nil, err
	}
	return &accountsHoldingAsset, nil
}

func (c *Client) AddressHoldingAsset(
	assetId string,
	params *utils.Parameters,
) (*models.AddressesHoldingAsset, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/assets/%s/addresses%s", assetId, formattedParams)
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
	var addressesHoldingAsset models.AddressesHoldingAsset
	err = json.NewDecoder(resp.Body).Decode(&addressesHoldingAsset)
	if err != nil {
		return nil, err
	}
	return &addressesHoldingAsset, nil
}

func (c *Client) Asset(assetId string) (*models.AssetInformations, error) {
	url := fmt.Sprintf("/assets/%s", assetId)
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
	var assetInformation models.AssetInformations
	err = json.NewDecoder(resp.Body).Decode(&assetInformation)
	if err != nil {
		return nil, err
	}
	return &assetInformation, nil
}

func (c *Client) AssetTransactions(
	assetId string,
	params *utils.Parameters,
) (*models.AssetTransactions, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/assets/%s/txs%s", assetId, formattedParams)
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
	var assetTransactions models.AssetTransactions
	err = json.NewDecoder(resp.Body).Decode(&assetTransactions)
	if err != nil {
		return nil, err
	}
	return &assetTransactions, nil
}

func (c *Client) AssetUpdates(
	assetId string,
	params *utils.Parameters,
) (*models.AssetUpdates, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/assets/%s/updates%s", assetId, formattedParams)
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
	var assetUpdates models.AssetUpdates
	err = json.NewDecoder(resp.Body).Decode(&assetUpdates)
	if err != nil {
		return nil, err
	}
	return &assetUpdates, nil

}

func (c *Client) AssetUtxos(assetId string, params *utils.Parameters) (*models.AssetUtxos, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/assets/%s/utxos%s", assetId, formattedParams)
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
	var assetUtxos models.AssetUtxos
	err = json.NewDecoder(resp.Body).Decode(&assetUtxos)
	if err != nil {
		return nil, err
	}
	return &assetUtxos, nil
}
