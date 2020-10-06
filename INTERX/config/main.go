package config

import (
	tsukiapp "github.com/TsukiCore/tsuki/simapp"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/sr25519"
)

// GenPrivKey is a function to generate a privKey
func GenPrivKey() crypto.PrivKey {
	return sr25519.GenPrivKey()
}

var (
	// EncodingCg is a configuration for Amino Encoding
	EncodingCg = tsukiapp.MakeEncodingConfig()
	// PrivKey is a private key for interx server
	PrivKey = GenPrivKey()
)
