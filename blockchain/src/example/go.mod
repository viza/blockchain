module main

go 1.19

replace blockchain/lorry => ../lorry

require blockchain/lorry v0.0.0-00010101000000-000000000000

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/ethereum/go-ethereum v1.10.26 // indirect
	golang.org/x/crypto v0.3.0 // indirect
	golang.org/x/sys v0.2.0 // indirect
)
