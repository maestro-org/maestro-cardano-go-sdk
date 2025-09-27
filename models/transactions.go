package models

import "github.com/maestro-org/go-sdk/utils"

type Certificates struct {
	AuthCommitteeHotCerts    []any `json:"auth_committee_hot_certs"`
	MirTransfers             []any `json:"mir_transfers"`
	PoolRegistrations        []any `json:"pool_registrations"`
	PoolRetirements          []any `json:"pool_retirements"`
	RegCerts                 []any `json:"reg_certs"`
	RegDRepCerts             []any `json:"reg_drep_certs"`
	ResignCommitteeColdCerts []any `json:"resign_committee_cold_certs"`
	StakeDelegations         []any `json:"stake_delegations"`
	StakeDeregistrations     []any `json:"stake_deregistrations"`
	StakeRegDelegations      []any `json:"stake_reg_delegations"`
	StakeRegistrations       []any `json:"stake_registrations"`
	StakeVoteDelegations     []any `json:"stake_vote_delegations"`
	StakeVoteRegDelegations  []any `json:"stake_vote_reg_delegations"`
	UnRegCerts               []any `json:"unreg_certs"`
	UnRegDRepCerts           []any `json:"unreg_drep_certs"`
	UpdateDRepCerts          []any `json:"update_drep_certs"`
	VoteDelegations          []any `json:"vote_delegations"`
	VoteRegDelegations       []any `json:"vote_reg_delegations"`
}

type Redeemers struct {
	Certificates []any `json:"certificates"`
	Mints        []any `json:"mints"`
	Spends       []any `json:"spends"`
	Withdrawals  []any `json:"withdrawals"`
	Votes        []any `json:"votes"`
	Proposals    []any `json:"proposals"`
}

type MintAsset struct {
	Unit   string `json:"unit"`
	Amount any    `json:"amount"`
}

type TransactionDetail struct {
	AdditionalSigners []string     `json:"additional_signers"`
	BlockAbsoluteSlot int64        `json:"block_absolute_slot"`
	BlockEpoch        int64        `json:"block_epoch"`
	BlockHash         string       `json:"block_hash"`
	BlockHeight       int64        `json:"block_height"`
	BlockTimestamp    int64        `json:"block_timestamp"`
	BlockTxIndex      int32        `json:"block_tx_index"`
	Certificates      Certificates `json:"certificates"`
	CollateralInputs  []Utxo       `json:"collateral_inputs"`
	CollateralReturn  *Utxo        `json:"collateral_return,omitempty"`
	Deposit           int64        `json:"deposit"`
	Fee               int64        `json:"fee"`
	Inputs            []Utxo       `json:"inputs"`
	InvalidBefore     *int64       `json:"invalid_before,omitempty"`
	InvalidHereafter  *int64       `json:"invalid_hereafter,omitempty"`
	Metadata          interface{}  `json:"metadata,omitempty"`
	Mint              []MintAsset  `json:"mint"`
	Outputs           []Utxo       `json:"outputs"`
	Redeemers         Redeemers    `json:"redeemers"`
	ReferenceInputs   []Utxo       `json:"reference_inputs"`
	ScriptsExecuted   []Script     `json:"scripts_executed"`
	ScriptsSuccessful bool         `json:"scripts_successful"`
	Size              int64        `json:"size"`
	TxHash            string       `json:"tx_hash"`
	Withdrawals       []any        `json:"withdrawals"`
}

type TransactionDetails struct {
	Data        TransactionDetail `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}

type TransactionOutputFromReference struct {
	Data        Utxo              `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}

type TxoReference struct {
	TxHash string `json:"tx_hash"`
	Index  int    `json:"index"`
}

type TransactionOutputsFromReferences struct {
	Data        []Utxo            `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
	NextCursor  string            `json:"next_cursor"`
}

type EvaluateTx struct {
	Cbor            string           `json:"cbor"`
	AdditionalUtxos []AdditionalUtxo `json:"additional_utxos"`
}

type ExecutionUnits struct {
	Mem   int64 `json:"mem"`
	Steps int64 `json:"steps"`
}

type RedeemerEvaluation struct {
	ExUnits       ExecutionUnits `json:"ex_units"`
	RedeemerIndex int            `json:"redeemer_index"`
	RedeemerTag   string         `json:"redeemer_tag"`
}

type EvaluateTxResponse []RedeemerEvaluation

type AdditionalUtxo struct {
	TxHash    string `json:"tx_hash"`
	Index     int    `json:"index"`
	TxoutCbor string `json:"txout_cbor"`
}
