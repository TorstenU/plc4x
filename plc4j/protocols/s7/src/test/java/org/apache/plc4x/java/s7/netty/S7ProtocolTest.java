/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.s7.netty;

import io.netty.channel.ChannelFuture;
import io.netty.channel.embedded.EmbeddedChannel;
import org.apache.plc4x.java.api.exceptions.PlcProtocolException;
import org.apache.plc4x.java.isotp.protocol.model.tpdus.DataTpdu;
import org.apache.plc4x.java.s7.netty.model.messages.S7RequestMessage;
import org.apache.plc4x.java.s7.netty.model.messages.SetupCommunicationRequestMessage;
import org.apache.plc4x.java.s7.netty.model.params.CpuServicesRequestParameter;
import org.apache.plc4x.java.s7.netty.model.params.VarParameter;
import org.apache.plc4x.java.s7.netty.model.params.items.S7AnyVarParameterItem;
import org.apache.plc4x.java.s7.netty.model.params.items.VarParameterItem;
import org.apache.plc4x.java.s7.netty.model.payloads.CpuServicesPayload;
import org.apache.plc4x.java.s7.netty.model.payloads.S7Payload;
import org.apache.plc4x.java.s7.netty.model.payloads.VarPayload;
import org.apache.plc4x.java.s7.netty.model.payloads.items.VarPayloadItem;
import org.apache.plc4x.java.s7.netty.model.payloads.ssls.SslModuleIdentificationDataRecord;
import org.apache.plc4x.java.s7.netty.model.types.*;
import org.junit.Before;
import org.junit.Test;

import java.util.Arrays;
import java.util.Collections;

import static org.hamcrest.Matchers.instanceOf;
import static org.hamcrest.core.IsEqual.equalTo;
import static org.hamcrest.core.IsNull.notNullValue;
import static org.hamcrest.core.IsNull.nullValue;
import static org.junit.Assert.assertThat;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

public class S7ProtocolTest {

    private EmbeddedChannel SUT;

    @Before
    public void setUp() {
        SUT = new EmbeddedChannel(new S7Protocol((short) 10, (short) 10, (short) 50, null));
    }

    /**
     * When not receiving S7Message objects, an exception should be thrown.
     */
    @Test
    public void testWriteJunk() {
        ChannelFuture channelFuture = SUT.writeOneOutbound("Hurz");
        Object outbound = SUT.readOutbound();

        assertThat("The protocol layer should not have output anything", outbound, nullValue());
    }

    @Test
    public void testSetupCommunication() {
        SUT.writeOneOutbound(
            new SetupCommunicationRequestMessage((short) 0x03, (short) 0x04, (short) 0x05, (short) 0x06, null)
        );
        Object outbound = SUT.readOutbound();

        assertThat("The protocol layer should have output something", outbound, notNullValue());
        assertThat("The output should have been of type DataTpdu", outbound, instanceOf(DataTpdu.class));

        DataTpdu dataTpdu = (DataTpdu) outbound;
        assertThat("The DataTpdu shouldn't have any parameters", dataTpdu.getParameters().isEmpty(), equalTo(true));

        byte[] actUserData = new byte[dataTpdu.getUserData().readableBytes()];
        dataTpdu.getUserData().readBytes(actUserData);

        byte[] refUserData = toByteArray(new int[] {
            // Protocol Id: 0x32 => S7Comm
            0x32,
            // MessageType.JOB
            0x01,
            0x00, 0x00,
            // Pdu Reference = 3
            0x00, 0x03,
            // Parameter Length = 8
            0x00, 0x08,
            // Payload Length = 0
            0x00, 0x00,
            // ParameterType.SETUP_COMMUNICATION
            0xf0,
            // Reserved
            0x00,
            // Max AMQ Calling = 0x04
            0x00, 0x04,
            // Max AMQ Callee = 0x05
            0x00, 0x05,
            // PLU Size = 0x06
            0x00, 0x06});

        assertThat("Output generated by the current layer doesn't match the expected output",
            Arrays.equals(actUserData, refUserData), equalTo(true));
    }

    @Test
    public void testCpuServices() {
        SUT.writeOneOutbound(
            new S7RequestMessage(MessageType.USER_DATA, (short) 2,
                Collections.singletonList(new CpuServicesRequestParameter(
                    CpuServicesParameterFunctionGroup.CPU_FUNCTIONS,
                    CpuServicesParameterSubFunctionGroup.READ_SSL, (byte) 0)),
                Collections.singletonList(new CpuServicesPayload(DataTransportErrorCode.OK, SslId.MODULE_IDENTIFICATION,
                    (short) 0x0000)), null));
        Object outbound = SUT.readOutbound();

        assertThat("The protocol layer should have output something", outbound, notNullValue());
        assertThat("The output should have been of type DataTpdu", outbound, instanceOf(DataTpdu.class));

        DataTpdu dataTpdu = (DataTpdu) outbound;
        assertThat("The DataTpdu shouldn't have any parameters", dataTpdu.getParameters().isEmpty(), equalTo(true));

        byte[] actUserData = new byte[dataTpdu.getUserData().readableBytes()];
        dataTpdu.getUserData().readBytes(actUserData);
        outputArray(actUserData);
        byte[] refUserData = toByteArray(new int[] {
            // Protocol Id: 0x32 => S7Comm
            0x32,
            // MessageType.USER_DATA
            0x07,
            // Reserved
            0x00, 0x00,
            // Pdu Reference = 2
            0x00, 0x02,
            // Parameter Length = 8
            0x00, 0x08,
            // Payload Length = 8
            0x00, 0x08,
            // ParameterType.CPU_SERVICES
            0x00,
            // ???
            0x01, 0x12,
            // Parameter Length
            0x04,
            // Type Request
            0x11,
            // Type: Request = 0x4 & Subtype: CPU functions = 0x4
            0x44,
            // Sub-function: Read SZL
            0x01,
            // Sequence Number: 0x00
            0x00,
            //////// Payload
            // Return code: Success
            0xFF,
            // Transport Size
            0x09,
            // Length
            0x00, 0x04,
            // SZL Id
            0x00, 0x11,
            // SZL Index
            0x00, 0x00
        });

        assertThat("Output generated by the current layer doesn't match the expected output",
            Arrays.equals(actUserData, refUserData), equalTo(true));
    }

    @Test
    public void testReadVar() {
        SUT.writeOneOutbound(
            new S7RequestMessage(MessageType.JOB, (short) 1, Collections.singletonList(
                new VarParameter(ParameterType.READ_VAR, Collections.singletonList(
                    new S7AnyVarParameterItem(SpecificationType.VARIABLE_SPECIFICATION, MemoryArea.DATA_BLOCKS,
                        TransportSize.BYTE, 1, (short) 2, (short) 3, (byte) 0)))),
                null, null));
        Object outbound = SUT.readOutbound();

        assertThat("The protocol layer should have output something", outbound, notNullValue());
        assertThat("The output should have been of type DataTpdu", outbound, instanceOf(DataTpdu.class));

        DataTpdu dataTpdu = (DataTpdu) outbound;
        assertThat("The DataTpdu shouldn't have any parameters", dataTpdu.getParameters().isEmpty(), equalTo(true));

        byte[] actUserData = new byte[dataTpdu.getUserData().readableBytes()];
        dataTpdu.getUserData().readBytes(actUserData);
//        outputArray(actUserData);
        byte[] refUserData = toByteArray(new int[] {
            // Protocol Id: 0x32 => S7Comm
            0x32,
            // MessageType.JOB
            0x01,
            0x00, 0x00,
            // Pdu Reference = 1
            0x00, 0x01,
            // Parameter Length = 14
            0x00, 0x0e,
            // Payload Length = 0
            0x00, 0x00,
            // ParameterType.READ_VAR
            0x04,
            // Number of items = 1
            0x01,
                // SpecificationType.VARIABLE_SPECIFICATION
                0x12,
                // Variable specification length = 10
                0x0a,
                // S7Any type of item
                0x10,
                // TransportSize.BYTE = 0x02
                0x02,
                // Number of items = 1
                0x00, 0x01,
                // DB Number = 2
                0x00, 0x02,
                // MemoryArea.DATA_BLOCKS = 0x84
                0x84,
                // Address: 00000000 00000000 00010000
                //          -----------------------...
                //          byte address           bit
                0x00, 0x00, 0x18});

        assertThat("Output generated by the current layer doesn't match the expected output",
            Arrays.equals(actUserData, refUserData), equalTo(true));
    }

    @Test
    public void testWriteVar() {
        SUT.writeOneOutbound(
            new S7RequestMessage(MessageType.JOB, (short) 1, Collections.singletonList(
                new VarParameter(ParameterType.WRITE_VAR, Collections.singletonList(
                    new S7AnyVarParameterItem(SpecificationType.VARIABLE_SPECIFICATION, MemoryArea.DATA_BLOCKS,
                        TransportSize.BYTE, 1, (short) 2, (short) 3, (byte) 0)))),
                Collections.singletonList(
                    new VarPayload(ParameterType.WRITE_VAR, Collections.singletonList(
                        new VarPayloadItem(DataTransportErrorCode.OK, DataTransportSize.BYTE_WORD_DWORD, new byte[]{(byte) 0x0A})))
                ), null));
        Object outbound = SUT.readOutbound();

        assertThat("The protocol layer should have output something", outbound, notNullValue());
        assertThat("The output should have been of type DataTpdu", outbound, instanceOf(DataTpdu.class));

        DataTpdu dataTpdu = (DataTpdu) outbound;
        assertThat("The DataTpdu shouldn't have any parameters", dataTpdu.getParameters().isEmpty(), equalTo(true));

        byte[] actUserData = new byte[dataTpdu.getUserData().readableBytes()];
        dataTpdu.getUserData().readBytes(actUserData);

        byte[] refUserData = toByteArray(new int[] {
            // Protocol Id: 0x32 => S7Comm
            0x32,
            // MessageType.JOB
            0x01,
            0x00, 0x00,
            // Pdu Reference = 1
            0x00, 0x01,
            // Parameter Length = 14
            0x00, 0x0e,
            // Payload Length = 5
            0x00, 0x05,
            // ParameterType.WRITE_VAR
            0x05,
            // Number of items = 1
            0x01,
            /////////////////////////////////////////////////////
            // Parameters ...
            /////////////////////////////////////////////////////
            // SpecificationType.VARIABLE_SPECIFICATION
            0x12,
            // Variable specification length = 10
            0x0a,
            // S7Any type of item
            0x10,
            // TransportSize.BYTE = 0x02
            0x02,
            // Number of items = 1
            0x00, 0x01,
            // DB Number = 2
            0x00, 0x02,
            // MemoryArea.DATA_BLOCKS = 0x84
            0x84,
            // Address: 00000000 00000000 00010000
            //          -----------------------...
            //          byte address           bit
            0x00, 0x00, 0x18,
            /////////////////////////////////////////////////////
            // Payloads
            /////////////////////////////////////////////////////
            // DataTransportErrorCode.OK
            0xff,
            // DataTransportSize.BYTE_WORD_DWOR
            0x04,
            // Length = 1
            0x00, 0x01,
            // Data: 0x0A
            0x0a
        });

        assertThat("Output generated by the current layer doesn't match the expected output",
            Arrays.equals(actUserData, refUserData), equalTo(true));
    }

    @Test
    public void testTooBigTpdu() {
        ChannelFuture channelFuture = SUT.writeOneOutbound(
            new S7RequestMessage(MessageType.JOB, (short) 1, Collections.singletonList(
                new VarParameter(ParameterType.READ_VAR, Arrays.asList(
                    new S7AnyVarParameterItem(SpecificationType.VARIABLE_SPECIFICATION, MemoryArea.DATA_BLOCKS,
                        TransportSize.BYTE, 1, (short) 2, (short) 3, (byte) 0),
                    new S7AnyVarParameterItem(SpecificationType.VARIABLE_SPECIFICATION, MemoryArea.DATA_BLOCKS,
                        TransportSize.BYTE, 1, (short) 2, (short) 3, (byte) 0),
                    new S7AnyVarParameterItem(SpecificationType.VARIABLE_SPECIFICATION, MemoryArea.DATA_BLOCKS,
                        TransportSize.BYTE, 1, (short) 2, (short) 3, (byte) 0),
                    new S7AnyVarParameterItem(SpecificationType.VARIABLE_SPECIFICATION, MemoryArea.DATA_BLOCKS,
                        TransportSize.BYTE, 1, (short) 2, (short) 3, (byte) 0),
                    new S7AnyVarParameterItem(SpecificationType.VARIABLE_SPECIFICATION, MemoryArea.DATA_BLOCKS,
                        TransportSize.BYTE, 1, (short) 2, (short) 3, (byte) 0)))),
                null, null));
        Object outbound = SUT.readOutbound();
        Throwable exception = channelFuture.cause();

        assertThat("The protocol layer should not have output anything", outbound, nullValue());
        assertThat("The protocol layer should have thrown an exception", exception, notNullValue());
        assertThat("The protocol layer should have thrown an exception", exception, instanceOf(PlcProtocolException.class));
    }

    @Test
    public void testNotImplementedParameter() {
        ChannelFuture channelFuture = SUT.writeOneOutbound(
            new S7RequestMessage(MessageType.JOB, (short) 1, Collections.singletonList(
                new VarParameter(ParameterType.UPLOAD, Collections.emptyList())),
                null, null));
        Object outbound = SUT.readOutbound();
        Throwable exception = channelFuture.cause();

        assertThat("The protocol layer should not have output anything", outbound, nullValue());
        assertThat("The protocol layer should have thrown an exception", exception, notNullValue());
        assertThat("The protocol layer should have thrown an exception", exception, instanceOf(PlcProtocolException.class));
    }

    @Test
    public void testNotImplementedPayload() {
        S7Payload payload  = mock(S7Payload.class);
        when(payload.getType()).thenReturn(ParameterType.UPLOAD);
        ChannelFuture channelFuture = SUT.writeOneOutbound(
            new S7RequestMessage(MessageType.JOB, (short) 1,
                Collections.emptyList(),
                Collections.singletonList(payload),
                null));
        Object outbound = SUT.readOutbound();
        Throwable exception = channelFuture.cause();

        assertThat("The protocol layer should not have output anything", outbound, nullValue());
        assertThat("The protocol layer should have thrown an exception", exception, notNullValue());
        assertThat("The protocol layer should have thrown an exception", exception, instanceOf(PlcProtocolException.class));
    }

    @Test
    public void testNotImplementedAddressingType() {
        VarParameterItem varParameterItem = mock(VarParameterItem.class);
        when(varParameterItem.getAddressingMode()).thenReturn(VariableAddressingMode.ALARM_QUERYREQ);
        ChannelFuture channelFuture = SUT.writeOneOutbound(
            new S7RequestMessage(MessageType.JOB, (short) 1, Collections.singletonList(
                new VarParameter(ParameterType.READ_VAR, Collections.singletonList(varParameterItem))),
                null, null));
        Object outbound = SUT.readOutbound();
        Throwable exception = channelFuture.cause();

        assertThat("The protocol layer should not have output anything", outbound, nullValue());
        assertThat("The protocol layer should have thrown an exception", exception, notNullValue());
        assertThat("The protocol layer should have thrown an exception", exception, instanceOf(PlcProtocolException.class));
    }

    @Test
    public void cpuServicesRequestWithSzlRecords() {
        ChannelFuture channelFuture = SUT.writeOneOutbound(
            new S7RequestMessage(MessageType.USER_DATA, (short) 2,
                Collections.singletonList(new CpuServicesRequestParameter(
                    CpuServicesParameterFunctionGroup.CPU_FUNCTIONS,
                    CpuServicesParameterSubFunctionGroup.READ_SSL, (byte) 0)),
                Collections.singletonList(new CpuServicesPayload(DataTransportErrorCode.OK, SslId.MODULE_IDENTIFICATION,
                    (short) 0x0000, Collections.singletonList(new SslModuleIdentificationDataRecord((short) 0, "hurz", (short) 0, (short) 0, (short) 0)))),
                null));
        Object outbound = SUT.readOutbound();
        Throwable exception = channelFuture.cause();

        assertThat("The protocol layer should not have output anything", outbound, nullValue());
        assertThat("The protocol layer should have thrown an exception", exception, notNullValue());
        assertThat("The protocol layer should have thrown an exception", exception, instanceOf(PlcProtocolException.class));
    }

    private static byte[] toByteArray(int[] input) {
        byte[] output = new byte[input.length];
        for (int i = 0; i < input.length; i++) {
            if((input[i] > 0xFF) || (input[i] < 0)) {
                throw new IllegalArgumentException("Int values passed in to 'toByteArray' should be vaild byte values.");
            }
            output[i] = (byte) input[i];
        }
        return output;
    }

    private static void outputArray(byte[] input) {
        StringBuilder sb = new StringBuilder();
        for (byte anInput : input) {
            if (sb.length() > 0) {
                sb.append(", ");
            }
            sb.append(String.format("0x%02x", anInput));
        }
        System.out.println(sb.toString());
    }

}