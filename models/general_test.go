package models

import (
	"encoding/json"
	"testing"
)

// TestExUnitsParsesCpuAsSteps locks the fix for the Maestro protocol-parameters
// quirk: the per-tx/per-block max execution-units "steps" budget is returned
// under the JSON key "cpu" (not "steps"). A wrong tag silently parses Steps as
// 0, which zeroes the local Plutus evaluation budget downstream.
func TestExUnitsParsesCpuAsSteps(t *testing.T) {
	const js = `{"memory":14000000,"cpu":10000000000}`
	var e ExUnits
	if err := json.Unmarshal([]byte(js), &e); err != nil {
		t.Fatalf("unmarshal ExUnits: %v", err)
	}
	if e.Memory != 14000000 {
		t.Fatalf("Memory = %d, want 14000000", e.Memory)
	}
	if e.Steps != 10000000000 {
		t.Fatalf("Steps = %d, want 10000000000 (Maestro returns it under \"cpu\")", e.Steps)
	}
}

// TestProtocolParamsMaxExecutionUnits parses the execution-units fields exactly
// as Maestro's /protocol-parameters returns them and asserts the step budgets
// are non-zero for both per-transaction and per-block.
func TestProtocolParamsMaxExecutionUnits(t *testing.T) {
	const js = `{
		"max_execution_units_per_transaction": {"memory": 16500000, "cpu": 10000000000},
		"max_execution_units_per_block": {"memory": 62000000, "cpu": 20000000000}
	}`
	var p ProtocolParams
	if err := json.Unmarshal([]byte(js), &p); err != nil {
		t.Fatalf("unmarshal ProtocolParams: %v", err)
	}
	if got := p.MaxExecutionUnitsPerTransaction.Steps; got != 10000000000 {
		t.Fatalf("per-tx Steps = %d, want 10000000000 (regression: cpu vs steps)", got)
	}
	if got := p.MaxExecutionUnitsPerTransaction.Memory; got != 16500000 {
		t.Fatalf("per-tx Memory = %d, want 16500000", got)
	}
	if got := p.MaxExecutionUnitsPerBlock.Steps; got != 20000000000 {
		t.Fatalf("per-block Steps = %d, want 20000000000", got)
	}
}
