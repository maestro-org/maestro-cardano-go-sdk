package models

import "github.com/maestro-org/go-sdk/utils"

type BytesSize struct {
	Bytes int64 `json:"bytes"`
}

type ChainTip struct {
	Data        ChainTipData      `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}

type ChainTipData struct {
	BlockHash string `json:"block_hash"`
	Height    int64  `json:"height"`
	Slot      int64  `json:"slot"`
}

type Era struct {
	End        EraTimeStamp `json:"end"`
	Parameters EraParams    `json:"parameters"`
	Start      EraTimeStamp `json:"start"`
}

type EraParams struct {
	EpochLength int64         `json:"epoch_length"`
	SafeZone    int64         `json:"safe_zone"`
	SlotLength  EraSlotLength `json:"slot_length"`
}

type EraSummaries struct {
	Data        []Era             `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}

type EraSlotLength struct {
	Milliseconds int64 `json:"milliseconds"`
}

type EraTime struct {
	Seconds int64 `json:"seconds"`
}

type EraTimeStamp struct {
	Epoch int64   `json:"epoch"`
	Slot  int64   `json:"slot"`
	Time  EraTime `json:"time"`
}

type ExUnits struct {
	Memory int64 `json:"memory"`
	Steps  int64 `json:"steps"`
}

type LovelaceAmount struct {
	Lovelace int64 `json:"lovelace"`
}

type ProtocolParameters struct {
	Data        ProtocolParams    `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}

type ProtocolVersion struct {
	Major int64 `json:"major"`
	Minor int64 `json:"minor"`
}

// StringExUnits holds the script execution prices from the protocol
// parameters, expressed as fraction strings (for example "721/10000000").
// The Maestro API returns these under "memory" and "cpu"; the CPU/step
// price is the "cpu" field (not "steps").
type StringExUnits struct {
	Memory string `json:"memory"`
	Cpu    string `json:"cpu"`
}

type AdaAmount struct {
	LovelaceAmount LovelaceAmount `json:"ada"`
}

type ProtocolParams struct {
	CollateralPercentage            int64           `json:"collateral_percentage"`
	DesiredNumberOfStakePools       int64           `json:"desired_number_of_stake_pools"`
	MaxBlockBodySize                BytesSize       `json:"max_block_body_size"`
	MaxBlockHeaderSize              BytesSize       `json:"max_block_header_size"`
	MaxCollateralInputs             int64           `json:"max_collateral_inputs"`
	MaxExecutionUnitsPerBlock       ExUnits         `json:"max_execution_units_per_block"`
	MaxExecutionUnitsPerTransaction ExUnits         `json:"max_execution_units_per_transaction"`
	MaxTransactionSize              BytesSize       `json:"max_transaction_size"`
	MaxValueSize                    BytesSize       `json:"max_value_size"`
	MinFeeCoefficient               int64           `json:"min_fee_coefficient"`
	MinFeeConstant                  AdaAmount       `json:"min_fee_constant"`
	MinStakePoolCost                AdaAmount       `json:"min_stake_pool_cost"`
	MinUtxoDepositCoefficient       int64           `json:"min_utxo_deposit_coefficient"`
	MinUtxoDepositConstant          AdaAmount       `json:"min_utxo_deposit_constant"`
	MonetaryExpansion               string          `json:"monetary_expansion"`
	PlutusCostModels                any             `json:"plutus_cost_models"`
	ProtocolVersion                 ProtocolVersion `json:"version"`
	ScriptExecutionPrices           StringExUnits   `json:"script_execution_prices"`
	StakeCredentialDeposit          AdaAmount       `json:"stake_credential_deposit"`
	StakePoolDeposit                AdaAmount       `json:"stake_pool_deposit"`
	StakePoolPledgeInfluence        string          `json:"stake_pool_pledge_influence"`
	StakePoolRetirementEpochBound   int64           `json:"stake_pool_retirement_epoch_bound"`
	TreasuryExpansion               string          `json:"treasury_expansion"`
}
