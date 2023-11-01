package signature

import (
	"errors"
	"fmt"
)

var ErrorUnsupportedType = errors.New("unsupported type")

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=Type --linecomment --output type_string.go
type Type uint16

const (
	TypeArweave       Type = iota + 1 // arweave
	TypeED25519                       // ed25519
	TypeEthereum                      // ethereum
	TypeSolana                        // solana
	TypeAptos                         // aptos
	TypeMultiAptos                    // multi_aptos
	TypeTypedEthereum                 // typed_ethereum
)

func (t Type) SignatureLength() (int, error) {
	switch t {
	case TypeArweave:
		return 512, nil
	case TypeED25519:
		return 64, nil
	case TypeEthereum:
		return 64 + 1, nil
	case TypeSolana:
		return 64, nil
	case TypeAptos:
		return 64, nil
	case TypeMultiAptos:
		return 64*32 + 4, nil
	case TypeTypedEthereum:
		return 64 + 1, nil
	default:
		return 0, fmt.Errorf("%w: %d", ErrorUnsupportedType, t)
	}
}

func (t Type) PublicKeyLength() (int, error) {
	switch t {
	case TypeArweave:
		return 512, nil
	case TypeED25519:
		return 32, nil
	case TypeEthereum:
		return 65, nil
	case TypeSolana:
		return 32, nil
	case TypeAptos:
		return 32, nil
	case TypeMultiAptos:
		return 32*32 + 1, nil
	case TypeTypedEthereum:
		return 42, nil
	default:
		return 0, fmt.Errorf("%w: %d", ErrorUnsupportedType, t)
	}
}
