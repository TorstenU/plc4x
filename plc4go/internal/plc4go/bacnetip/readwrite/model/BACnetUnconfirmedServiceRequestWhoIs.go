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
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// Constant values.
const BACnetUnconfirmedServiceRequestWhoIs_DEVICEINSTANCERANGELOWLIMITHEADER uint8 = 0x01
const BACnetUnconfirmedServiceRequestWhoIs_DEVICEINSTANCERANGEHIGHLIMITHEADER uint8 = 0x03

// The data-structure of this message
type BACnetUnconfirmedServiceRequestWhoIs struct {
	DeviceInstanceRangeLowLimitLength  uint8
	DeviceInstanceRangeLowLimit        []int8
	DeviceInstanceRangeHighLimitLength uint8
	DeviceInstanceRangeHighLimit       []int8
	Parent                             *BACnetUnconfirmedServiceRequest
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestWhoIs interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetUnconfirmedServiceRequestWhoIs) ServiceChoice() uint8 {
	return 0x08
}

func (m *BACnetUnconfirmedServiceRequestWhoIs) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func NewBACnetUnconfirmedServiceRequestWhoIs(deviceInstanceRangeLowLimitLength uint8, deviceInstanceRangeLowLimit []int8, deviceInstanceRangeHighLimitLength uint8, deviceInstanceRangeHighLimit []int8) *BACnetUnconfirmedServiceRequest {
	child := &BACnetUnconfirmedServiceRequestWhoIs{
		DeviceInstanceRangeLowLimitLength:  deviceInstanceRangeLowLimitLength,
		DeviceInstanceRangeLowLimit:        deviceInstanceRangeLowLimit,
		DeviceInstanceRangeHighLimitLength: deviceInstanceRangeHighLimitLength,
		DeviceInstanceRangeHighLimit:       deviceInstanceRangeHighLimit,
		Parent:                             NewBACnetUnconfirmedServiceRequest(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetUnconfirmedServiceRequestWhoIs(structType interface{}) *BACnetUnconfirmedServiceRequestWhoIs {
	castFunc := func(typ interface{}) *BACnetUnconfirmedServiceRequestWhoIs {
		if casted, ok := typ.(BACnetUnconfirmedServiceRequestWhoIs); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequestWhoIs); ok {
			return casted
		}
		if casted, ok := typ.(BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestWhoIs(casted.Child)
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestWhoIs(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetUnconfirmedServiceRequestWhoIs) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestWhoIs"
}

func (m *BACnetUnconfirmedServiceRequestWhoIs) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Const Field (deviceInstanceRangeLowLimitHeader)
	lengthInBits += 5

	// Simple field (deviceInstanceRangeLowLimitLength)
	lengthInBits += 3

	// Array field
	if len(m.DeviceInstanceRangeLowLimit) > 0 {
		lengthInBits += 8 * uint16(len(m.DeviceInstanceRangeLowLimit))
	}

	// Const Field (deviceInstanceRangeHighLimitHeader)
	lengthInBits += 5

	// Simple field (deviceInstanceRangeHighLimitLength)
	lengthInBits += 3

	// Array field
	if len(m.DeviceInstanceRangeHighLimit) > 0 {
		lengthInBits += 8 * uint16(len(m.DeviceInstanceRangeHighLimit))
	}

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestWhoIs) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestWhoIsParse(io *utils.ReadBuffer) (*BACnetUnconfirmedServiceRequest, error) {

	// Const Field (deviceInstanceRangeLowLimitHeader)
	deviceInstanceRangeLowLimitHeader, _deviceInstanceRangeLowLimitHeaderErr := io.ReadUint8(5)
	if _deviceInstanceRangeLowLimitHeaderErr != nil {
		return nil, errors.Wrap(_deviceInstanceRangeLowLimitHeaderErr, "Error parsing 'deviceInstanceRangeLowLimitHeader' field")
	}
	if deviceInstanceRangeLowLimitHeader != BACnetUnconfirmedServiceRequestWhoIs_DEVICEINSTANCERANGELOWLIMITHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestWhoIs_DEVICEINSTANCERANGELOWLIMITHEADER) + " but got " + fmt.Sprintf("%d", deviceInstanceRangeLowLimitHeader))
	}

	// Simple Field (deviceInstanceRangeLowLimitLength)
	deviceInstanceRangeLowLimitLength, _deviceInstanceRangeLowLimitLengthErr := io.ReadUint8(3)
	if _deviceInstanceRangeLowLimitLengthErr != nil {
		return nil, errors.Wrap(_deviceInstanceRangeLowLimitLengthErr, "Error parsing 'deviceInstanceRangeLowLimitLength' field")
	}

	// Array field (deviceInstanceRangeLowLimit)
	// Count array
	deviceInstanceRangeLowLimit := make([]int8, deviceInstanceRangeLowLimitLength)
	for curItem := uint16(0); curItem < uint16(deviceInstanceRangeLowLimitLength); curItem++ {
		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'deviceInstanceRangeLowLimit' field")
		}
		deviceInstanceRangeLowLimit[curItem] = _item
	}
	if len(deviceInstanceRangeLowLimit) == 0 {
		deviceInstanceRangeLowLimit = nil
	}

	// Const Field (deviceInstanceRangeHighLimitHeader)
	deviceInstanceRangeHighLimitHeader, _deviceInstanceRangeHighLimitHeaderErr := io.ReadUint8(5)
	if _deviceInstanceRangeHighLimitHeaderErr != nil {
		return nil, errors.Wrap(_deviceInstanceRangeHighLimitHeaderErr, "Error parsing 'deviceInstanceRangeHighLimitHeader' field")
	}
	if deviceInstanceRangeHighLimitHeader != BACnetUnconfirmedServiceRequestWhoIs_DEVICEINSTANCERANGEHIGHLIMITHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestWhoIs_DEVICEINSTANCERANGEHIGHLIMITHEADER) + " but got " + fmt.Sprintf("%d", deviceInstanceRangeHighLimitHeader))
	}

	// Simple Field (deviceInstanceRangeHighLimitLength)
	deviceInstanceRangeHighLimitLength, _deviceInstanceRangeHighLimitLengthErr := io.ReadUint8(3)
	if _deviceInstanceRangeHighLimitLengthErr != nil {
		return nil, errors.Wrap(_deviceInstanceRangeHighLimitLengthErr, "Error parsing 'deviceInstanceRangeHighLimitLength' field")
	}

	// Array field (deviceInstanceRangeHighLimit)
	// Count array
	deviceInstanceRangeHighLimit := make([]int8, deviceInstanceRangeHighLimitLength)
	for curItem := uint16(0); curItem < uint16(deviceInstanceRangeHighLimitLength); curItem++ {
		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'deviceInstanceRangeHighLimit' field")
		}
		deviceInstanceRangeHighLimit[curItem] = _item
	}
	if len(deviceInstanceRangeHighLimit) == 0 {
		deviceInstanceRangeHighLimit = nil
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestWhoIs{
		DeviceInstanceRangeLowLimitLength:  deviceInstanceRangeLowLimitLength,
		DeviceInstanceRangeLowLimit:        deviceInstanceRangeLowLimit,
		DeviceInstanceRangeHighLimitLength: deviceInstanceRangeHighLimitLength,
		DeviceInstanceRangeHighLimit:       deviceInstanceRangeHighLimit,
		Parent:                             &BACnetUnconfirmedServiceRequest{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetUnconfirmedServiceRequestWhoIs) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Const Field (deviceInstanceRangeLowLimitHeader)
		_deviceInstanceRangeLowLimitHeaderErr := io.WriteUint8(5, 0x01)
		if _deviceInstanceRangeLowLimitHeaderErr != nil {
			return errors.Wrap(_deviceInstanceRangeLowLimitHeaderErr, "Error serializing 'deviceInstanceRangeLowLimitHeader' field")
		}

		// Simple Field (deviceInstanceRangeLowLimitLength)
		deviceInstanceRangeLowLimitLength := uint8(m.DeviceInstanceRangeLowLimitLength)
		_deviceInstanceRangeLowLimitLengthErr := io.WriteUint8(3, (deviceInstanceRangeLowLimitLength))
		if _deviceInstanceRangeLowLimitLengthErr != nil {
			return errors.Wrap(_deviceInstanceRangeLowLimitLengthErr, "Error serializing 'deviceInstanceRangeLowLimitLength' field")
		}

		// Array Field (deviceInstanceRangeLowLimit)
		if m.DeviceInstanceRangeLowLimit != nil {
			for _, _element := range m.DeviceInstanceRangeLowLimit {
				_elementErr := io.WriteInt8(8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'deviceInstanceRangeLowLimit' field")
				}
			}
		}

		// Const Field (deviceInstanceRangeHighLimitHeader)
		_deviceInstanceRangeHighLimitHeaderErr := io.WriteUint8(5, 0x03)
		if _deviceInstanceRangeHighLimitHeaderErr != nil {
			return errors.Wrap(_deviceInstanceRangeHighLimitHeaderErr, "Error serializing 'deviceInstanceRangeHighLimitHeader' field")
		}

		// Simple Field (deviceInstanceRangeHighLimitLength)
		deviceInstanceRangeHighLimitLength := uint8(m.DeviceInstanceRangeHighLimitLength)
		_deviceInstanceRangeHighLimitLengthErr := io.WriteUint8(3, (deviceInstanceRangeHighLimitLength))
		if _deviceInstanceRangeHighLimitLengthErr != nil {
			return errors.Wrap(_deviceInstanceRangeHighLimitLengthErr, "Error serializing 'deviceInstanceRangeHighLimitLength' field")
		}

		// Array Field (deviceInstanceRangeHighLimit)
		if m.DeviceInstanceRangeHighLimit != nil {
			for _, _element := range m.DeviceInstanceRangeHighLimit {
				_elementErr := io.WriteInt8(8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'deviceInstanceRangeHighLimit' field")
				}
			}
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestWhoIs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "deviceInstanceRangeLowLimitLength":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DeviceInstanceRangeLowLimitLength = data
			case "deviceInstanceRangeLowLimit":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.DeviceInstanceRangeLowLimit = utils.ByteArrayToInt8Array(_decoded[0:_len])
			case "deviceInstanceRangeHighLimitLength":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DeviceInstanceRangeHighLimitLength = data
			case "deviceInstanceRangeHighLimit":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.DeviceInstanceRangeHighLimit = utils.ByteArrayToInt8Array(_decoded[0:_len])
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

func (m *BACnetUnconfirmedServiceRequestWhoIs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.DeviceInstanceRangeLowLimitLength, xml.StartElement{Name: xml.Name{Local: "deviceInstanceRangeLowLimitLength"}}); err != nil {
		return err
	}
	_encodedDeviceInstanceRangeLowLimit := hex.EncodeToString(utils.Int8ArrayToByteArray(m.DeviceInstanceRangeLowLimit))
	_encodedDeviceInstanceRangeLowLimit = strings.ToUpper(_encodedDeviceInstanceRangeLowLimit)
	if err := e.EncodeElement(_encodedDeviceInstanceRangeLowLimit, xml.StartElement{Name: xml.Name{Local: "deviceInstanceRangeLowLimit"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.DeviceInstanceRangeHighLimitLength, xml.StartElement{Name: xml.Name{Local: "deviceInstanceRangeHighLimitLength"}}); err != nil {
		return err
	}
	_encodedDeviceInstanceRangeHighLimit := hex.EncodeToString(utils.Int8ArrayToByteArray(m.DeviceInstanceRangeHighLimit))
	_encodedDeviceInstanceRangeHighLimit = strings.ToUpper(_encodedDeviceInstanceRangeHighLimit)
	if err := e.EncodeElement(_encodedDeviceInstanceRangeHighLimit, xml.StartElement{Name: xml.Name{Local: "deviceInstanceRangeHighLimit"}}); err != nil {
		return err
	}
	return nil
}

func (m BACnetUnconfirmedServiceRequestWhoIs) String() string {
	return string(m.Box("BACnetUnconfirmedServiceRequestWhoIs", utils.DefaultWidth*2))
}

func (m BACnetUnconfirmedServiceRequestWhoIs) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "BACnetUnconfirmedServiceRequestWhoIs"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("DeviceInstanceRangeLowLimitLength", m.DeviceInstanceRangeLowLimitLength, width-2))
	boxes = append(boxes, utils.BoxAnything("DeviceInstanceRangeLowLimit", m.DeviceInstanceRangeLowLimit, width-2))
	boxes = append(boxes, utils.BoxAnything("DeviceInstanceRangeHighLimitLength", m.DeviceInstanceRangeHighLimitLength, width-2))
	boxes = append(boxes, utils.BoxAnything("DeviceInstanceRangeHighLimit", m.DeviceInstanceRangeHighLimit, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
