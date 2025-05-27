package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maestro-org/go-sdk/models"
	"github.com/maestro-org/go-sdk/utils"
)

func (c *Client) ListOfRegisteredPools(params *utils.Parameters) (*models.RegisteredPools, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/pools%s", formattedParams)
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
	var registeredPools models.RegisteredPools
	err = json.NewDecoder(resp.Body).Decode(&registeredPools)
	if err != nil {
		return nil, err
	}
	return &registeredPools, nil
}

func (c *Client) StakePoolMintedBlocks(
	poolId string,
	params *utils.Parameters,
) (*models.PoolMintedBlocks, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/pools/%s/blocks%s", poolId, formattedParams)
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
	var poolMintedBlocks models.PoolMintedBlocks
	err = json.NewDecoder(resp.Body).Decode(&poolMintedBlocks)
	if err != nil {
		return nil, err
	}
	return &poolMintedBlocks, nil
}

func (c *Client) StakePoolDelegators(
	poolId string,
	params *utils.Parameters,
) (*models.StakePoolDelegators, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/pools/%s/delegators%s", poolId, formattedParams)
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
	var stakePoolDelegators models.StakePoolDelegators
	err = json.NewDecoder(resp.Body).Decode(&stakePoolDelegators)
	if err != nil {
		return nil, err
	}
	return &stakePoolDelegators, nil
}

func (c *Client) StakePoolHistory(
	poolId string,
	params *utils.Parameters,
) (*models.StakePoolHistory, error) {
	formattedParams := ""
	if params != nil {
		formattedParams = params.Format()
	}
	url := fmt.Sprintf("/pools/%s/history%s", poolId, formattedParams)
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
	var stakePoolHistory models.StakePoolHistory
	err = json.NewDecoder(resp.Body).Decode(&stakePoolHistory)
	if err != nil {
		return nil, err
	}
	return &stakePoolHistory, nil
}

func (c *Client) StakePoolInformation(poolId string) (*models.StakePoolInformation, error) {
	url := fmt.Sprintf("/pools/%s/info", poolId)
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
	var stakePoolInformation models.StakePoolInformation
	err = json.NewDecoder(resp.Body).Decode(&stakePoolInformation)
	if err != nil {
		return nil, err
	}
	return &stakePoolInformation, nil
}

func (c *Client) StakePoolMetadata(poolId string) (*models.StakePoolMetadata, error) {
	url := fmt.Sprintf("/pools/%s/metadata", poolId)
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
	var stakePoolMetadata models.StakePoolMetadata
	err = json.NewDecoder(resp.Body).Decode(&stakePoolMetadata)
	if err != nil {
		return nil, err
	}
	return &stakePoolMetadata, nil
}

func (c *Client) StakePoolRelays(poolId string) (*models.StakePoolRelays, error) {
	url := fmt.Sprintf("/pools/%s/relays", poolId)
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
	var stakePoolRelays models.StakePoolRelays
	err = json.NewDecoder(resp.Body).Decode(&stakePoolRelays)
	if err != nil {
		return nil, err
	}
	return &stakePoolRelays, nil

}

func (c *Client) StakePoolUpdates(poolId string) (*models.StakePoolUpdates, error) {
	url := fmt.Sprintf("/pools/%s/updates", poolId)
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
	var stakePoolUpdates models.StakePoolUpdates
	err = json.NewDecoder(resp.Body).Decode(&stakePoolUpdates)
	if err != nil {
		return nil, err
	}
	return &stakePoolUpdates, nil

}
