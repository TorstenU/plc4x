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
type ApduDataExtAuthorizeRequest struct {
	Level  uint8
	Data   []uint8
	Parent *ApduDataExt
}

// The corresponding interface
type IApduDataExtAuthorizeRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtAuthorizeRequest) ExtApciType() uint8 {
	return 0x11
}

func (m *ApduDataExtAuthorizeRequest) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtAuthorizeRequest(level uint8, data []uint8) *ApduDataExt {
	child := &ApduDataExtAuthorizeRequest{
		Level:  level,
		Data:   data,
		Parent: NewApduDataExt(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataExtAuthorizeRequest(structType interface{}) *ApduDataExtAuthorizeRequest {
	castFunc := func(typ interface{}) *ApduDataExtAuthorizeRequest {
		if casted, ok := typ.(ApduDataExtAuthorizeRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExtAuthorizeRequest); ok {
			return casted
		}
		if casted, ok := typ.(ApduDataExt); ok {
			return CastApduDataExtAuthorizeRequest(casted.Child)
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return CastApduDataExtAuthorizeRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExtAuthorizeRequest) GetTypeName() string {
	return "ApduDataExtAuthorizeRequest"
}

func (m *ApduDataExtAuthorizeRequest) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (level)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *ApduDataExtAuthorizeRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataExtAuthorizeRequestParse(io *utils.ReadBuffer) (*ApduDataExt, error) {

	// Simple Field (level)
	level, _levelErr := io.ReadUint8(8)
	if _levelErr != nil {
		return nil, errors.Wrap(_levelErr, "Error parsing 'level' field")
	}

	// Array field (data)
	// Count array
	data := make([]uint8, uint16(4))
	for curItem := uint16(0); curItem < uint16(uint16(4)); curItem++ {
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
	_child := &ApduDataExtAuthorizeRequest{
		Level:  level,
		Data:   data,
		Parent: &ApduDataExt{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataExtAuthorizeRequest) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (level)
		level := uint8(m.Level)
		_levelErr := io.WriteUint8(8, (level))
		if _levelErr != nil {
			return errors.Wrap(_levelErr, "Error serializing 'level' field")
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

func (m *ApduDataExtAuthorizeRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "level":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Level = data
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

func (m *ApduDataExtAuthorizeRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Level, xml.StartElement{Name: xml.Name{Local: "level"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Data, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	return nil
}

func (m ApduDataExtAuthorizeRequest) String() string {
	return string(m.Box("ApduDataExtAuthorizeRequest", utils.DefaultWidth*2))
}

func (m ApduDataExtAuthorizeRequest) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "ApduDataExtAuthorizeRequest"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("Level", m.Level, width-2))
	boxes = append(boxes, utils.BoxAnything("Data", m.Data, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
