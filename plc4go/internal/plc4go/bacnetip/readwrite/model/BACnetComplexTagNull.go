/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetComplexTagNull struct {
	*BACnetComplexTag
}

// The corresponding interface
type IBACnetComplexTagNull interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetComplexTagNull) DataType() BACnetDataType {
	return BACnetDataType_NULL
}

func (m *BACnetComplexTagNull) InitializeParent(parent *BACnetComplexTag, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, actualLength uint32) {
	m.TagNumber = tagNumber
	m.TagClass = tagClass
	m.LengthValueType = lengthValueType
	m.ExtTagNumber = extTagNumber
	m.ExtLength = extLength
	m.ExtExtLength = extExtLength
	m.ExtExtExtLength = extExtExtLength
}

func NewBACnetComplexTagNull(tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetComplexTag {
	child := &BACnetComplexTagNull{
		BACnetComplexTag: NewBACnetComplexTag(tagNumber, tagClass, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength),
	}
	child.Child = child
	return child.BACnetComplexTag
}

func CastBACnetComplexTagNull(structType interface{}) *BACnetComplexTagNull {
	castFunc := func(typ interface{}) *BACnetComplexTagNull {
		if casted, ok := typ.(BACnetComplexTagNull); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetComplexTagNull); ok {
			return casted
		}
		if casted, ok := typ.(BACnetComplexTag); ok {
			return CastBACnetComplexTagNull(casted.Child)
		}
		if casted, ok := typ.(*BACnetComplexTag); ok {
			return CastBACnetComplexTagNull(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetComplexTagNull) GetTypeName() string {
	return "BACnetComplexTagNull"
}

func (m *BACnetComplexTagNull) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetComplexTagNull) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetComplexTagNull) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetComplexTagNullParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (*BACnetComplexTag, error) {
	if pullErr := readBuffer.PullContext("BACnetComplexTagNull"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetComplexTagNull"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetComplexTagNull{
		BACnetComplexTag: &BACnetComplexTag{},
	}
	_child.BACnetComplexTag.Child = _child
	return _child.BACnetComplexTag, nil
}

func (m *BACnetComplexTagNull) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetComplexTagNull"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetComplexTagNull"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetComplexTagNull) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
