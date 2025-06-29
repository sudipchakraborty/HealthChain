package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		RecordList: []Record{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in record
	recordIdMap := make(map[uint64]bool)
	recordCount := gs.GetRecordCount()
	for _, elem := range gs.RecordList {
		if _, ok := recordIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for record")
		}
		if elem.Id >= recordCount {
			return fmt.Errorf("record id should be lower or equal than the last id")
		}
		recordIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
