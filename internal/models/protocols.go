package models

// TProtocolsFromMeta is the structure of data for the protocols metadata stored in data/meta/protocols
type TProtocolsFromMeta struct {
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	ChainID      uint64         `json:"chainID"`
	Localization *TLocalization `json:"localization,omitempty"`
}
