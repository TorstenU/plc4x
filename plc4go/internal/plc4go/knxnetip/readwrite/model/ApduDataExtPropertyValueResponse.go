//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ApduDataExtPropertyValueResponse struct {
	ObjectIndex uint8
	PropertyId  uint8
	Count       uint8
	Index       uint16
	Data        []uint8
	Parent      *ApduDataExt
}

// The corresponding interface
type IApduDataExtPropertyValueResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtPropertyValueResponse) ExtApciType() uint8 {
	return 0x16
}

func (m *ApduDataExtPropertyValueResponse) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtPropertyValueResponse(objectIndex uint8, propertyId uint8, count uint8, index uint16, data []uint8) *ApduDataExt {
	child := &ApduDataExtPropertyValueResponse{
		ObjectIndex: objectIndex,
		PropertyId:  propertyId,
		Count:       count,
		Index:       index,
		Data:        data,
		Parent:      NewApduDataExt(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataExtPropertyValueResponse(structType interface{}) *ApduDataExtPropertyValueResponse {
	castFunc := func(typ interface{}) *ApduDataExtPropertyValueResponse {
		if casted, ok := typ.(ApduDataExtPropertyValueResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExtPropertyValueResponse); ok {
			return casted
		}
		if casted, ok := typ.(ApduDataExt); ok {
			return CastApduDataExtPropertyValueResponse(casted.Child)
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return CastApduDataExtPropertyValueResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExtPropertyValueResponse) GetTypeName() string {
	return "ApduDataExtPropertyValueResponse"
}

func (m *ApduDataExtPropertyValueResponse) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectIndex)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (count)
	lengthInBits += 4

	// Simple field (index)
	lengthInBits += 12

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *ApduDataExtPropertyValueResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataExtPropertyValueResponseParse(io *utils.ReadBuffer, length uint8) (*ApduDataExt, error) {

	// Simple Field (objectIndex)
	objectIndex, _objectIndexErr := io.ReadUint8(8)
	if _objectIndexErr != nil {
		return nil, errors.Wrap(_objectIndexErr, "Error parsing 'objectIndex' field")
	}

	// Simple Field (propertyId)
	propertyId, _propertyIdErr := io.ReadUint8(8)
	if _propertyIdErr != nil {
		return nil, errors.Wrap(_propertyIdErr, "Error parsing 'propertyId' field")
	}

	// Simple Field (count)
	count, _countErr := io.ReadUint8(4)
	if _countErr != nil {
		return nil, errors.Wrap(_countErr, "Error parsing 'count' field")
	}

	// Simple Field (index)
	index, _indexErr := io.ReadUint16(12)
	if _indexErr != nil {
		return nil, errors.Wrap(_indexErr, "Error parsing 'index' field")
	}

	// Array field (data)
	// Count array
	data := make([]uint8, uint16(length)-uint16(uint16(5)))
	for curItem := uint16(0); curItem < uint16(uint16(length)-uint16(uint16(5))); curItem++ {
		_item, _err := io.ReadUint8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data[curItem] = _item
	}
	if len(data) == 0 {
		data = nil
	}

	// Create a partially initialized instance
	_child := &ApduDataExtPropertyValueResponse{
		ObjectIndex: objectIndex,
		PropertyId:  propertyId,
		Count:       count,
		Index:       index,
		Data:        data,
		Parent:      &ApduDataExt{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataExtPropertyValueResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (objectIndex)
		objectIndex := uint8(m.ObjectIndex)
		_objectIndexErr := io.WriteUint8(8, (objectIndex))
		if _objectIndexErr != nil {
			return errors.Wrap(_objectIndexErr, "Error serializing 'objectIndex' field")
		}

		// Simple Field (propertyId)
		propertyId := uint8(m.PropertyId)
		_propertyIdErr := io.WriteUint8(8, (propertyId))
		if _propertyIdErr != nil {
			return errors.Wrap(_propertyIdErr, "Error serializing 'propertyId' field")
		}

		// Simple Field (count)
		count := uint8(m.Count)
		_countErr := io.WriteUint8(4, (count))
		if _countErr != nil {
			return errors.Wrap(_countErr, "Error serializing 'count' field")
		}

		// Simple Field (index)
		index := uint16(m.Index)
		_indexErr := io.WriteUint16(12, (index))
		if _indexErr != nil {
			return errors.Wrap(_indexErr, "Error serializing 'index' field")
		}

		// Array Field (data)
		if m.Data != nil {
			for _, _element := range m.Data {
				_elementErr := io.WriteUint8(8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataExtPropertyValueResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "objectIndex":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ObjectIndex = data
			case "propertyId":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.PropertyId = data
			case "count":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Count = data
			case "index":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Index = data
			case "data":
				var data []uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Data = data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *ApduDataExtPropertyValueResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.ObjectIndex, xml.StartElement{Name: xml.Name{Local: "objectIndex"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.PropertyId, xml.StartElement{Name: xml.Name{Local: "propertyId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Count, xml.StartElement{Name: xml.Name{Local: "count"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Index, xml.StartElement{Name: xml.Name{Local: "index"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Data, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	return nil
}

func (m ApduDataExtPropertyValueResponse) String() string {
	return string(m.Box("ApduDataExtPropertyValueResponse", utils.DefaultWidth*2))
}

func (m ApduDataExtPropertyValueResponse) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "ApduDataExtPropertyValueResponse"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("ObjectIndex", m.ObjectIndex, width-2))
	boxes = append(boxes, utils.BoxAnything("PropertyId", m.PropertyId, width-2))
	boxes = append(boxes, utils.BoxAnything("Count", m.Count, width-2))
	boxes = append(boxes, utils.BoxAnything("Index", m.Index, width-2))
	boxes = append(boxes, utils.BoxAnything("Data", m.Data, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
