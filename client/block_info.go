package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/maestro-org/go-sdk/utils"
)

type OperationalCertificate struct {
	HotVkey        string `json:"hot_vkey"`
	KesPeriod      int64  `json:"kes_period"`
	KesSignature   string `json:"kes_signature"`
	SequenceNumber int64  `json:"sequence_number"`
}

type TotalExUnits struct {
	Mem   int64 `json:"mem"`
	Steps int64 `json:"steps"`
}

type BlockInfo struct {
	Data struct {
		AbsoluteSlot           int64                  `json:"absolute_slot"`
		BlockProducer          string                 `json:"block_producer"`
		Confirmations          int64                  `json:"confirmations"`
		Epoch                  int64                  `json:"epoch"`
		EpochSlot              int64                  `json:"epoch_slot"`
		Era                    string                 `json:"era"`
		Hash                   string                 `json:"hash"`
		Height                 int64                  `json:"height"`
		OperationalCertificate OperationalCertificate `json:"operational_certificate"`
		PreviousBlock          string                 `json:"previous_block"`
		ProtocolVersion        []int                  `json:"protocol_version"`
		ScriptInvocations      int32                  `json:"script_invocations"`
		Size                   int32                  `json:"size"`
		Timestamp              string                 `json:"timestamp"`
		TotalExUnits           TotalExUnits           `json:"total_ex_units"`
		TotalFees              int64                  `json:"total_fees"`
		TotalOutputLovelace    string                 `json:"total_output_lovelace"`
		TxHashes               []string               `json:"tx_hashes"`
		VrfKey                 string                 `json:"vrf_key"`
	} `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}

func (c *Client) BlockInfo(hashOrHeight string) (*BlockInfo, error) {
	return c.BlockInfoWithOptions(hashOrHeight, false)
}

func (c *Client) BlockInfoWithOptions(hashOrHeight string, amountsAsStrings bool) (*BlockInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blocks/%s", c.BaseUrl, hashOrHeight), nil)
	if err != nil {
		return nil, err
	}
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}

	req.Header.Set("Content-Type", "application/json")
	if amountsAsStrings {
		req.Header.Set("amounts-as-strings", "true")
	} else {
		req.Header.Set("amounts-as-strings", "false")
	}

	var responseBody string
	blockInfo := BlockInfo{}
	if err := c.sendRequest(req, &responseBody); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(responseBody), &blockInfo); err != nil {
		fmt.Println("Cannot unmarshal JSON: ", err)
		os.Exit(1)
	}

	return &blockInfo, nil
}

func (c *Client) LatestBlock() (*BlockInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blocks/latest", c.BaseUrl), nil)
	if err != nil {
		return nil, err
	}
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}

	req.Header.Set("Content-Type", "application/json")

	var responseBody string
	blockInfo := BlockInfo{}
	if err := c.sendRequest(req, &responseBody); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(responseBody), &blockInfo); err != nil {
		fmt.Println("Cannot unmarshal JSON: ", err)
		os.Exit(1)
	}

	return &blockInfo, nil
}
