package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maestro-org/go-sdk/models"
	"github.com/maestro-org/go-sdk/utils"
)

func (c *Client) AccountAddresses(
	stakeAddr string,
	params *utils.Parameters,
) (*models.AccountAddresses, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/accounts/%s/addresses%s", stakeAddr, formattedParams)
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
	var accountAddresses models.AccountAddresses
	err = json.NewDecoder(resp.Body).Decode(&accountAddresses)
	if err != nil {
		return nil, err
	}
	return &accountAddresses, nil
}

func (c *Client) AccountAssets(
	stakeAddr string,
	params *utils.Parameters,
) (*models.AccountAssets, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/accounts/%s/assets%s", stakeAddr, formattedParams)
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
	var accountAssets models.AccountAssets
	err = json.NewDecoder(resp.Body).Decode(&accountAssets)
	if err != nil {
		return nil, err
	}
	return &accountAssets, nil
}

func (c *Client) StakeAccountHistory(
	stakeAddr string,
	params *utils.Parameters,
) (*models.StakeAccountHistory, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/accounts/%s/history%s", stakeAddr, formattedParams)
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
	var stakeAccountHistory models.StakeAccountHistory
	err = json.NewDecoder(resp.Body).Decode(&stakeAccountHistory)
	if err != nil {
		return nil, err
	}
	return &stakeAccountHistory, nil
}

func (c *Client) StakeAccountInformation(
	stakeAddr string,
) (*models.StakeAccountInformation, error) {
	url := fmt.Sprintf("/accounts/%s", stakeAddr)
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
	var stakeAccountInformation models.StakeAccountInformation
	err = json.NewDecoder(resp.Body).Decode(&stakeAccountInformation)
	if err != nil {
		return nil, err
	}
	return &stakeAccountInformation, nil
}

func (c *Client) StakeAccountRewards(
	stakeAddr string,
	params *utils.Parameters,
) (*models.StakeAccountRewards, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/accounts/%s/rewards%s", stakeAddr, formattedParams)
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
	var stakeAccountRewards models.StakeAccountRewards
	err = json.NewDecoder(resp.Body).Decode(&stakeAccountRewards)
	if err != nil {
		return nil, err
	}
	return &stakeAccountRewards, nil
}

func (c *Client) StakeAccountUpdates(
	stakeAddr string,
	params *utils.Parameters,
) (*models.StakeAccountUpdates, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/accounts/%s/updates%s", stakeAddr, formattedParams)
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
	var stakeAccountUpdates models.StakeAccountUpdates
	err = json.NewDecoder(resp.Body).Decode(&stakeAccountUpdates)
	if err != nil {
		return nil, err
	}
	return &stakeAccountUpdates, nil
}
