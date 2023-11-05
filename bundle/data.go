package bundle

import "github.com/naturalselectionlabs/RSS3-Arweave-SDK/signature"

const (
	MaxTags          = 128
	MaxTagKeyBytes   = 1024
	MaxTagValueBytes = 3072
)

type DataInfo struct {
	Size uint64
	ID   string
}

type DataItem struct {
	SignatureType signature.Type
	Signature     []byte
	Owner         []byte
	Target        []byte
	Anchor        []byte
	Tags          []byte

	Reader Reader
}

type DataTag struct {
	Name  []byte `avro:"name"`
	Value []byte `avro:"value"`
}
