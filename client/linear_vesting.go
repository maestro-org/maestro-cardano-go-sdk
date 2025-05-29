package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maestro-org/go-sdk/models"
)

type LockBody struct {
	Sender                   string `json:"sender"`
	Beneficiary              string `json:"beneficiary"`
	AssetPolicyId            string `json:"asset_policy_id"`
	AssetTokenName           string `json:"asset_token_name"`
	TotalVestingQuantity     int64  `json:"total_vesting_quantity"`
	VestingPeriodStart       int64  `json:"vesting_period_start"`
	VestingPeriodEnd         int64  `json:"vesting_period_end"`
	FirstUnlockPossibleAfter int64  `json:"first_unlock_possible_after"`
	TotalInstallments        int64  `json:"total_installments"`
}

func (c *Client) LockAssets(lockBody LockBody) (*models.LockTransaction, error) {
	url := "/contracts/vesting/lock"

	body := lockBody
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
	var lockTransaction models.LockTransaction
	err = json.NewDecoder(resp.Body).Decode(&lockTransaction)
	if err != nil {
		return nil, err
	}
	return &lockTransaction, nil
}

func (c *Client) StateOfVestingAssets(beneficiary string) (*[]models.VestingState, error) {
	url := fmt.Sprintf("/contracts/vesting/state/%s", beneficiary)
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
	var vestingStates []models.VestingState
	err = json.NewDecoder(resp.Body).Decode(&vestingStates)
	if err != nil {
		return nil, err
	}
	return &vestingStates, nil

}

func (c *Client) CollectAssets(beneficiary string) (*models.CollectTransaction, error) {
	url := fmt.Sprintf("/contracts/vesting/collect/%s", beneficiary)
	resp, err := c.post(url, nil)
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
	var collectTransaction models.CollectTransaction
	err = json.NewDecoder(resp.Body).Decode(&collectTransaction)
	if err != nil {
		return nil, err
	}
	return &collectTransaction, nil
}
