package utils

import (
	"strings"

	"github.com/Decred-Next/dcrnd/chaincfg/v3"
	cfg "github.com/decred/dcrd/chaincfg/v2"
	"github.com/decred/dcrwallet/errors"
	"github.com/jinzhu/copier"
)

var (
	mainnetParams = chaincfg.MainNetParams()
	testnetParams = chaincfg.TestNet3Params()
)

func ChainParams(netType string) (*cfg.Params, error) {
	param := cfg.Params{}
	switch strings.ToLower(netType) {
	case strings.ToLower(mainnetParams.Name):
		copier.Copy(&param, &mainnetParams)
		return &param, nil
	case strings.ToLower(testnetParams.Name):
		copier.Copy(&param, &testnetParams)
		return &param, nil
	default:
		return nil, errors.New("invalid net type")
	}
}
