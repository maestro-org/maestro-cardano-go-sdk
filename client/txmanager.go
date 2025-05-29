package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/maestro-org/go-sdk/models"
)

func (c *Client) TxManagerHistory() (*[]models.TxManagerState, error) {
	url := "/txmanager/history"
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
	var txManagerStates []models.TxManagerState
	err = json.NewDecoder(resp.Body).Decode(&txManagerStates)
	if err != nil {
		return nil, err
	}
	return &txManagerStates, nil
}

func (c *Client) TxManagerSubmit(txHex string) (string, error) {
	url := "/txmanager"
	txBuffer, err := hex.DecodeString(txHex)
	if err != nil {
		return "", err
	}
	resp, err := c.postBuffer(url, txBuffer)
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", fmt.Errorf("empty response")
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return "", fmt.Errorf("unexpected error: %d", resp.Body)
	}
	defer resp.Body.Close() //nolint:errcheck
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	txHash := string(bodyBytes)
	return txHash, nil
}

func (c *Client) TxManagerSubmitTurbo(txHex string) (string, error) {
	url := "/txmanager/turbosubmit"
	txBuffer, err := hex.DecodeString(txHex)
	if err != nil {
		return "", err
	}
	resp, err := c.postBuffer(url, txBuffer)
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", fmt.Errorf("empty response")
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return "", fmt.Errorf("unexpected error: %d", resp.Body)
	}
	defer resp.Body.Close() //nolint:errcheck
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	txHash := string(bodyBytes)
	return txHash, nil
}

func (c *Client) TxManagerState(txHash string) (*models.TxManagerState, error) {
	url := fmt.Sprintf("/txmanager/%s/state", txHash)
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
	var txManagerState models.TxManagerState
	err = json.NewDecoder(resp.Body).Decode(&txManagerState)
	if err != nil {
		return nil, err
	}
	return &txManagerState, nil
}
