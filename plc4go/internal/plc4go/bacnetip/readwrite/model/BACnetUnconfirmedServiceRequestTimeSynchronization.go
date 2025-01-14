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
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetUnconfirmedServiceRequestTimeSynchronization struct {
	*BACnetUnconfirmedServiceRequest
	SynchronizedDate *BACnetTagApplicationDate
	SynchronizedTime *BACnetTagApplicationTime
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestTimeSynchronization interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) ServiceChoice() uint8 {
	return 0x06
}

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func NewBACnetUnconfirmedServiceRequestTimeSynchronization(synchronizedDate *BACnetTagApplicationDate, synchronizedTime *BACnetTagApplicationTime) *BACnetUnconfirmedServiceRequest {
	child := &BACnetUnconfirmedServiceRequestTimeSynchronization{
		SynchronizedDate:                synchronizedDate,
		SynchronizedTime:                synchronizedTime,
		BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(),
	}
	child.Child = child
	return child.BACnetUnconfirmedServiceRequest
}

func CastBACnetUnconfirmedServiceRequestTimeSynchronization(structType interface{}) *BACnetUnconfirmedServiceRequestTimeSynchronization {
	castFunc := func(typ interface{}) *BACnetUnconfirmedServiceRequestTimeSynchronization {
		if casted, ok := typ.(BACnetUnconfirmedServiceRequestTimeSynchronization); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequestTimeSynchronization); ok {
			return casted
		}
		if casted, ok := typ.(BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestTimeSynchronization(casted.Child)
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestTimeSynchronization(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestTimeSynchronization"
}

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (synchronizedDate)
	lengthInBits += m.SynchronizedDate.LengthInBits()

	// Simple field (synchronizedTime)
	lengthInBits += m.SynchronizedTime.LengthInBits()

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestTimeSynchronizationParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetUnconfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestTimeSynchronization"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (synchronizedDate)
	if pullErr := readBuffer.PullContext("synchronizedDate"); pullErr != nil {
		return nil, pullErr
	}
	_synchronizedDate, _synchronizedDateErr := BACnetTagParse(readBuffer)
	if _synchronizedDateErr != nil {
		return nil, errors.Wrap(_synchronizedDateErr, "Error parsing 'synchronizedDate' field")
	}
	synchronizedDate := CastBACnetTagApplicationDate(_synchronizedDate)
	if closeErr := readBuffer.CloseContext("synchronizedDate"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (synchronizedTime)
	if pullErr := readBuffer.PullContext("synchronizedTime"); pullErr != nil {
		return nil, pullErr
	}
	_synchronizedTime, _synchronizedTimeErr := BACnetTagParse(readBuffer)
	if _synchronizedTimeErr != nil {
		return nil, errors.Wrap(_synchronizedTimeErr, "Error parsing 'synchronizedTime' field")
	}
	synchronizedTime := CastBACnetTagApplicationTime(_synchronizedTime)
	if closeErr := readBuffer.CloseContext("synchronizedTime"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestTimeSynchronization"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestTimeSynchronization{
		SynchronizedDate:                CastBACnetTagApplicationDate(synchronizedDate),
		SynchronizedTime:                CastBACnetTagApplicationTime(synchronizedTime),
		BACnetUnconfirmedServiceRequest: &BACnetUnconfirmedServiceRequest{},
	}
	_child.BACnetUnconfirmedServiceRequest.Child = _child
	return _child.BACnetUnconfirmedServiceRequest, nil
}

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestTimeSynchronization"); pushErr != nil {
			return pushErr
		}

		// Simple Field (synchronizedDate)
		if pushErr := writeBuffer.PushContext("synchronizedDate"); pushErr != nil {
			return pushErr
		}
		_synchronizedDateErr := m.SynchronizedDate.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("synchronizedDate"); popErr != nil {
			return popErr
		}
		if _synchronizedDateErr != nil {
			return errors.Wrap(_synchronizedDateErr, "Error serializing 'synchronizedDate' field")
		}

		// Simple Field (synchronizedTime)
		if pushErr := writeBuffer.PushContext("synchronizedTime"); pushErr != nil {
			return pushErr
		}
		_synchronizedTimeErr := m.SynchronizedTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("synchronizedTime"); popErr != nil {
			return popErr
		}
		if _synchronizedTimeErr != nil {
			return errors.Wrap(_synchronizedTimeErr, "Error serializing 'synchronizedTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestTimeSynchronization"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
