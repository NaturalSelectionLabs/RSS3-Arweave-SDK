package bundle

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/hamba/avro/v2"
	"github.com/naturalselectionlabs/RSS3-Arweave-SDK/signature"
	"github.com/samber/lo"
)

const SchemaDataTags = `{"type":"array","items":{"type":"record","name":"Tag","fields":[{"name":"name","type":"bytes"},{"name":"value","type":"bytes"}]}}`

type Decoder struct {
	reader, buffer io.Reader
	dataInfos      []DataInfo
	dataTagsSchema avro.Schema
	cursor         int
}

func (d *Decoder) DecodeDataInfos() ([]DataInfo, error) {
	// Read the number of the data infos.
	buffer := make([]byte, 32)

	if _, err := io.ReadFull(d.reader, buffer); err != nil {
		return nil, fmt.Errorf("read number of infos: %w", err)
	}

	// Pre-allocate memory for the data info slice.
	d.dataInfos = make([]DataInfo, binary.LittleEndian.Uint64(buffer))

	for index := 0; index < len(d.dataInfos); index++ {
		var dataInfo DataInfo

		// Read the size of the data info.
		if _, err := io.ReadFull(d.reader, buffer); err != nil {
			return nil, fmt.Errorf("read size of info %d: %w", index, err)
		}

		dataInfo.Size = binary.LittleEndian.Uint64(buffer)

		// Read the id of the data info.
		if _, err := io.ReadFull(d.reader, buffer); err != nil {
			return nil, fmt.Errorf("read id of info %d: %w", index, err)
		}

		dataInfo.ID = base64.RawURLEncoding.EncodeToString(buffer)

		d.dataInfos[index] = dataInfo
	}

	return d.dataInfos, nil
}

func (d *Decoder) DecodeDataItem() (*DataItem, error) {
	dataItemInfo := d.dataInfos[d.cursor]

	d.buffer = io.LimitReader(d.reader, int64(dataItemInfo.Size))

	var dataItem DataItem

	if err := d.decodeDataItemSignature(&dataItem); err != nil {
		return nil, fmt.Errorf("decode signature: %w", err)
	}

	if err := d.decodeDataItemOwner(&dataItem); err != nil {
		return nil, fmt.Errorf("decode owner: %w", err)
	}

	if err := d.decodeDataItemTarget(&dataItem); err != nil {
		return nil, fmt.Errorf("decode target: %w", err)
	}

	if err := d.decodeDataItemAnchor(&dataItem); err != nil {
		return nil, fmt.Errorf("decode anchor: %w", err)
	}

	if err := d.decodeDataItemTags(&dataItem); err != nil {
		return nil, fmt.Errorf("decode tags: %w", err)
	}

	dataItem.Reader = Reader{
		reader: d.buffer,
	}

	d.cursor++

	return &dataItem, nil
}

func (d *Decoder) DecodeDataItemTags(dataItem DataItem) ([]DataTag, error) {
	var tags []DataTag

	if err := avro.Unmarshal(d.dataTagsSchema, dataItem.Tags, &tags); err != nil {
		return nil, err
	}

	return tags, nil
}

func (d *Decoder) decodeDataItemSignature(dataItem *DataItem) error {
	// Read the signature type for the data item.
	buffer := make([]byte, 2)
	if _, err := io.ReadFull(d.buffer, buffer); err != nil {
		return err
	}

	dataItem.SignatureType = signature.Type(binary.LittleEndian.Uint16(buffer))

	// Read the signature for the data item.
	signatureLength, err := dataItem.SignatureType.SignatureLength()
	if err != nil {
		return err
	}

	dataItem.Signature = make([]byte, signatureLength)
	if _, err := io.ReadFull(d.buffer, dataItem.Signature); err != nil {
		return err
	}

	return nil
}

func (d *Decoder) decodeDataItemOwner(dataItem *DataItem) error {
	publicKeyLength, err := dataItem.SignatureType.PublicKeyLength()
	if err != nil {
		return err
	}

	dataItem.Owner = make([]byte, publicKeyLength)
	if _, err := io.ReadFull(d.buffer, dataItem.Owner); err != nil {
		return err
	}

	return nil
}

func (d *Decoder) decodeDataItemTarget(dataItem *DataItem) error {
	buffer := make([]byte, 32+1)

	presence, err := d.decodeDataItemOptionField(buffer)
	if err != nil {
		return err
	}

	if presence {
		dataItem.Target = buffer[1:]
	}

	return nil
}

func (d *Decoder) decodeDataItemAnchor(dataItem *DataItem) error {
	buffer := make([]byte, 32+1)

	presence, err := d.decodeDataItemOptionField(buffer)
	if err != nil {
		return err
	}

	if presence {
		dataItem.Anchor = buffer[1:]
	}

	return nil
}

func (d *Decoder) decodeDataItemOptionField(buffer []byte) (bool, error) {
	if _, err := io.ReadFull(d.buffer, buffer[:1]); err != nil {
		return false, err
	}

	switch flag := buffer[0]; flag {
	case 0x0:
		return false, nil
	case 0x1:
		_, err := io.ReadFull(d.buffer, buffer[1:])

		return true, err
	default:
		return false, fmt.Errorf("invalid presence byte: %d", flag)
	}
}

func (d *Decoder) decodeDataItemTags(dataItem *DataItem) error {
	buffer := make([]byte, 8)

	// Read the count, but discard it
	if _, err := io.ReadFull(d.buffer, buffer); err != nil {
		return err
	}

	// Read the size, but discard it
	if _, err := io.ReadFull(d.buffer, buffer); err != nil {
		return err
	}

	dataItem.Tags = make([]byte, binary.LittleEndian.Uint16(buffer))

	if _, err := io.ReadFull(d.buffer, dataItem.Tags); err != nil {
		return err
	}

	return nil
}

func (d *Decoder) Next() bool {
	return d.cursor < len(d.dataInfos)
}

func NewDecoder(reader io.Reader) *Decoder {
	decoder := Decoder{
		reader:         reader,
		dataTagsSchema: lo.Must(avro.Parse(SchemaDataTags)),
	}

	return &decoder
}
