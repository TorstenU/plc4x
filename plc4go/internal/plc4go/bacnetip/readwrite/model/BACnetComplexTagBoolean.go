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
type BACnetComplexTagBoolean struct {
	Parent *BACnetComplexTag
}

// The corresponding interface
type IBACnetComplexTagBoolean interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetComplexTagBoolean) DataType() BACnetDataType {
	return BACnetDataType_BOOLEAN
}

func (m *BACnetComplexTagBoolean) InitializeParent(parent *BACnetComplexTag, tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, isPrimitiveAndNotBoolean bool, actualLength uint32) {
	m.Parent.TagNumber = tagNumber
	m.Parent.LengthValueType = lengthValueType
	m.Parent.ExtTagNumber = extTagNumber
	m.Parent.ExtLength = extLength
	m.Parent.ExtExtLength = extExtLength
	m.Parent.ExtExtExtLength = extExtExtLength
}

func NewBACnetComplexTagBoolean(tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetComplexTag {
	child := &BACnetComplexTagBoolean{
		Parent: NewBACnetComplexTag(tagNumber, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetComplexTagBoolean(structType interface{}) *BACnetComplexTagBoolean {
	castFunc := func(typ interface{}) *BACnetComplexTagBoolean {
		if casted, ok := typ.(BACnetComplexTagBoolean); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetComplexTagBoolean); ok {
			return casted
		}
		if casted, ok := typ.(BACnetComplexTag); ok {
			return CastBACnetComplexTagBoolean(casted.Child)
		}
		if casted, ok := typ.(*BACnetComplexTag); ok {
			return CastBACnetComplexTagBoolean(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetComplexTagBoolean) GetTypeName() string {
	return "BACnetComplexTagBoolean"
}

func (m *BACnetComplexTagBoolean) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetComplexTagBoolean) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetComplexTagBoolean) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetComplexTagBooleanParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (*BACnetComplexTag, error) {
	if pullErr := readBuffer.PullContext("BACnetComplexTagBoolean"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetComplexTagBoolean"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetComplexTagBoolean{
		Parent: &BACnetComplexTag{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetComplexTagBoolean) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetComplexTagBoolean"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetComplexTagBoolean"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetComplexTagBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}