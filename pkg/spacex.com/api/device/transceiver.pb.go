// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.5
// source: spacex/api/device/transceiver.proto

package device

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransceiverModulatorState int32

const (
	TransceiverModulatorState_MODSTATE_UNKNOWN  TransceiverModulatorState = 0
	TransceiverModulatorState_MODSTATE_ENABLED  TransceiverModulatorState = 1
	TransceiverModulatorState_MODSTATE_DISABLED TransceiverModulatorState = 2
)

// Enum value maps for TransceiverModulatorState.
var (
	TransceiverModulatorState_name = map[int32]string{
		0: "MODSTATE_UNKNOWN",
		1: "MODSTATE_ENABLED",
		2: "MODSTATE_DISABLED",
	}
	TransceiverModulatorState_value = map[string]int32{
		"MODSTATE_UNKNOWN":  0,
		"MODSTATE_ENABLED":  1,
		"MODSTATE_DISABLED": 2,
	}
)

func (x TransceiverModulatorState) Enum() *TransceiverModulatorState {
	p := new(TransceiverModulatorState)
	*p = x
	return p
}

func (x TransceiverModulatorState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransceiverModulatorState) Descriptor() protoreflect.EnumDescriptor {
	return file_spacex_api_device_transceiver_proto_enumTypes[0].Descriptor()
}

func (TransceiverModulatorState) Type() protoreflect.EnumType {
	return &file_spacex_api_device_transceiver_proto_enumTypes[0]
}

func (x TransceiverModulatorState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransceiverModulatorState.Descriptor instead.
func (TransceiverModulatorState) EnumDescriptor() ([]byte, []int) {
	return file_spacex_api_device_transceiver_proto_rawDescGZIP(), []int{0}
}

type TransceiverTxRxState int32

const (
	TransceiverTxRxState_TXRX_UNKNOWN  TransceiverTxRxState = 0
	TransceiverTxRxState_TXRX_ENABLED  TransceiverTxRxState = 1
	TransceiverTxRxState_TXRX_DISABLED TransceiverTxRxState = 2
)

// Enum value maps for TransceiverTxRxState.
var (
	TransceiverTxRxState_name = map[int32]string{
		0: "TXRX_UNKNOWN",
		1: "TXRX_ENABLED",
		2: "TXRX_DISABLED",
	}
	TransceiverTxRxState_value = map[string]int32{
		"TXRX_UNKNOWN":  0,
		"TXRX_ENABLED":  1,
		"TXRX_DISABLED": 2,
	}
)

func (x TransceiverTxRxState) Enum() *TransceiverTxRxState {
	p := new(TransceiverTxRxState)
	*p = x
	return p
}

func (x TransceiverTxRxState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransceiverTxRxState) Descriptor() protoreflect.EnumDescriptor {
	return file_spacex_api_device_transceiver_proto_enumTypes[1].Descriptor()
}

func (TransceiverTxRxState) Type() protoreflect.EnumType {
	return &file_spacex_api_device_transceiver_proto_enumTypes[1]
}

func (x TransceiverTxRxState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransceiverTxRxState.Descriptor instead.
func (TransceiverTxRxState) EnumDescriptor() ([]byte, []int) {
	return file_spacex_api_device_transceiver_proto_rawDescGZIP(), []int{1}
}

type TransceiverTransmitBlankingState int32

const (
	TransceiverTransmitBlankingState_TB_UNKNOWN  TransceiverTransmitBlankingState = 0
	TransceiverTransmitBlankingState_TB_ENABLED  TransceiverTransmitBlankingState = 1
	TransceiverTransmitBlankingState_TB_DISABLED TransceiverTransmitBlankingState = 2
)

// Enum value maps for TransceiverTransmitBlankingState.
var (
	TransceiverTransmitBlankingState_name = map[int32]string{
		0: "TB_UNKNOWN",
		1: "TB_ENABLED",
		2: "TB_DISABLED",
	}
	TransceiverTransmitBlankingState_value = map[string]int32{
		"TB_UNKNOWN":  0,
		"TB_ENABLED":  1,
		"TB_DISABLED": 2,
	}
)

func (x TransceiverTransmitBlankingState) Enum() *TransceiverTransmitBlankingState {
	p := new(TransceiverTransmitBlankingState)
	*p = x
	return p
}

func (x TransceiverTransmitBlankingState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransceiverTransmitBlankingState) Descriptor() protoreflect.EnumDescriptor {
	return file_spacex_api_device_transceiver_proto_enumTypes[2].Descriptor()
}

func (TransceiverTransmitBlankingState) Type() protoreflect.EnumType {
	return &file_spacex_api_device_transceiver_proto_enumTypes[2]
}

func (x TransceiverTransmitBlankingState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransceiverTransmitBlankingState.Descriptor instead.
func (TransceiverTransmitBlankingState) EnumDescriptor() ([]byte, []int) {
	return file_spacex_api_device_transceiver_proto_rawDescGZIP(), []int{2}
}

type TransceiverIFLoopbackTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnableIfLoopback bool `protobuf:"varint,1,opt,name=enable_if_loopback,json=enableIfLoopback,proto3" json:"enable_if_loopback,omitempty"`
}

func (x *TransceiverIFLoopbackTestRequest) Reset() {
	*x = TransceiverIFLoopbackTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacex_api_device_transceiver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransceiverIFLoopbackTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransceiverIFLoopbackTestRequest) ProtoMessage() {}

func (x *TransceiverIFLoopbackTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacex_api_device_transceiver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransceiverIFLoopbackTestRequest.ProtoReflect.Descriptor instead.
func (*TransceiverIFLoopbackTestRequest) Descriptor() ([]byte, []int) {
	return file_spacex_api_device_transceiver_proto_rawDescGZIP(), []int{0}
}

func (x *TransceiverIFLoopbackTestRequest) GetEnableIfLoopback() bool {
	if x != nil {
		return x.EnableIfLoopback
	}
	return false
}

type TransceiverIFLoopbackTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BerLoopbackTest  float32 `protobuf:"fixed32,1,opt,name=ber_loopback_test,json=berLoopbackTest,proto3" json:"ber_loopback_test,omitempty"`
	SnrLoopbackTest  float32 `protobuf:"fixed32,2,opt,name=snr_loopback_test,json=snrLoopbackTest,proto3" json:"snr_loopback_test,omitempty"`
	RssiLoopbackTest float32 `protobuf:"fixed32,3,opt,name=rssi_loopback_test,json=rssiLoopbackTest,proto3" json:"rssi_loopback_test,omitempty"`
	PllLock          bool    `protobuf:"varint,4,opt,name=pll_lock,json=pllLock,proto3" json:"pll_lock,omitempty"`
}

func (x *TransceiverIFLoopbackTestResponse) Reset() {
	*x = TransceiverIFLoopbackTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacex_api_device_transceiver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransceiverIFLoopbackTestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransceiverIFLoopbackTestResponse) ProtoMessage() {}

func (x *TransceiverIFLoopbackTestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacex_api_device_transceiver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransceiverIFLoopbackTestResponse.ProtoReflect.Descriptor instead.
func (*TransceiverIFLoopbackTestResponse) Descriptor() ([]byte, []int) {
	return file_spacex_api_device_transceiver_proto_rawDescGZIP(), []int{1}
}

func (x *TransceiverIFLoopbackTestResponse) GetBerLoopbackTest() float32 {
	if x != nil {
		return x.BerLoopbackTest
	}
	return 0
}

func (x *TransceiverIFLoopbackTestResponse) GetSnrLoopbackTest() float32 {
	if x != nil {
		return x.SnrLoopbackTest
	}
	return 0
}

func (x *TransceiverIFLoopbackTestResponse) GetRssiLoopbackTest() float32 {
	if x != nil {
		return x.RssiLoopbackTest
	}
	return 0
}

func (x *TransceiverIFLoopbackTestResponse) GetPllLock() bool {
	if x != nil {
		return x.PllLock
	}
	return false
}

type TransceiverGetStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TransceiverGetStatusRequest) Reset() {
	*x = TransceiverGetStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacex_api_device_transceiver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransceiverGetStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransceiverGetStatusRequest) ProtoMessage() {}

func (x *TransceiverGetStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacex_api_device_transceiver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransceiverGetStatusRequest.ProtoReflect.Descriptor instead.
func (*TransceiverGetStatusRequest) Descriptor() ([]byte, []int) {
	return file_spacex_api_device_transceiver_proto_rawDescGZIP(), []int{2}
}

type TransceiverGetStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModState              TransceiverModulatorState        `protobuf:"varint,1,opt,name=mod_state,json=modState,proto3,enum=SpaceX.API.Device.TransceiverModulatorState" json:"mod_state,omitempty"`
	DemodState            TransceiverModulatorState        `protobuf:"varint,2,opt,name=demod_state,json=demodState,proto3,enum=SpaceX.API.Device.TransceiverModulatorState" json:"demod_state,omitempty"`
	TxState               TransceiverTxRxState             `protobuf:"varint,3,opt,name=tx_state,json=txState,proto3,enum=SpaceX.API.Device.TransceiverTxRxState" json:"tx_state,omitempty"`
	RxState               TransceiverTxRxState             `protobuf:"varint,4,opt,name=rx_state,json=rxState,proto3,enum=SpaceX.API.Device.TransceiverTxRxState" json:"rx_state,omitempty"`
	State                 DishState                        `protobuf:"varint,1006,opt,name=state,proto3,enum=SpaceX.API.Device.DishState" json:"state,omitempty"`
	Faults                *TransceiverFaults               `protobuf:"bytes,1007,opt,name=faults,proto3" json:"faults,omitempty"`
	TransmitBlankingState TransceiverTransmitBlankingState `protobuf:"varint,1008,opt,name=transmit_blanking_state,json=transmitBlankingState,proto3,enum=SpaceX.API.Device.TransceiverTransmitBlankingState" json:"transmit_blanking_state,omitempty"`
	ModemAsicTemp         float32                          `protobuf:"fixed32,1009,opt,name=modem_asic_temp,json=modemAsicTemp,proto3" json:"modem_asic_temp,omitempty"`
	TxIfTemp              float32                          `protobuf:"fixed32,1010,opt,name=tx_if_temp,json=txIfTemp,proto3" json:"tx_if_temp,omitempty"`
}

func (x *TransceiverGetStatusResponse) Reset() {
	*x = TransceiverGetStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacex_api_device_transceiver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransceiverGetStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransceiverGetStatusResponse) ProtoMessage() {}

func (x *TransceiverGetStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacex_api_device_transceiver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransceiverGetStatusResponse.ProtoReflect.Descriptor instead.
func (*TransceiverGetStatusResponse) Descriptor() ([]byte, []int) {
	return file_spacex_api_device_transceiver_proto_rawDescGZIP(), []int{3}
}

func (x *TransceiverGetStatusResponse) GetModState() TransceiverModulatorState {
	if x != nil {
		return x.ModState
	}
	return TransceiverModulatorState_MODSTATE_UNKNOWN
}

func (x *TransceiverGetStatusResponse) GetDemodState() TransceiverModulatorState {
	if x != nil {
		return x.DemodState
	}
	return TransceiverModulatorState_MODSTATE_UNKNOWN
}

func (x *TransceiverGetStatusResponse) GetTxState() TransceiverTxRxState {
	if x != nil {
		return x.TxState
	}
	return TransceiverTxRxState_TXRX_UNKNOWN
}

func (x *TransceiverGetStatusResponse) GetRxState() TransceiverTxRxState {
	if x != nil {
		return x.RxState
	}
	return TransceiverTxRxState_TXRX_UNKNOWN
}

func (x *TransceiverGetStatusResponse) GetState() DishState {
	if x != nil {
		return x.State
	}
	return DishState_UNKNOWN
}

func (x *TransceiverGetStatusResponse) GetFaults() *TransceiverFaults {
	if x != nil {
		return x.Faults
	}
	return nil
}

func (x *TransceiverGetStatusResponse) GetTransmitBlankingState() TransceiverTransmitBlankingState {
	if x != nil {
		return x.TransmitBlankingState
	}
	return TransceiverTransmitBlankingState_TB_UNKNOWN
}

func (x *TransceiverGetStatusResponse) GetModemAsicTemp() float32 {
	if x != nil {
		return x.ModemAsicTemp
	}
	return 0
}

func (x *TransceiverGetStatusResponse) GetTxIfTemp() float32 {
	if x != nil {
		return x.TxIfTemp
	}
	return 0
}

type TransceiverFaults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OverTempModemAsicFault bool `protobuf:"varint,1,opt,name=over_temp_modem_asic_fault,json=overTempModemAsicFault,proto3" json:"over_temp_modem_asic_fault,omitempty"`
	OverTempPcbaFault      bool `protobuf:"varint,2,opt,name=over_temp_pcba_fault,json=overTempPcbaFault,proto3" json:"over_temp_pcba_fault,omitempty"`
	DcVoltageFault         bool `protobuf:"varint,3,opt,name=dc_voltage_fault,json=dcVoltageFault,proto3" json:"dc_voltage_fault,omitempty"`
}

func (x *TransceiverFaults) Reset() {
	*x = TransceiverFaults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacex_api_device_transceiver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransceiverFaults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransceiverFaults) ProtoMessage() {}

func (x *TransceiverFaults) ProtoReflect() protoreflect.Message {
	mi := &file_spacex_api_device_transceiver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransceiverFaults.ProtoReflect.Descriptor instead.
func (*TransceiverFaults) Descriptor() ([]byte, []int) {
	return file_spacex_api_device_transceiver_proto_rawDescGZIP(), []int{4}
}

func (x *TransceiverFaults) GetOverTempModemAsicFault() bool {
	if x != nil {
		return x.OverTempModemAsicFault
	}
	return false
}

func (x *TransceiverFaults) GetOverTempPcbaFault() bool {
	if x != nil {
		return x.OverTempPcbaFault
	}
	return false
}

func (x *TransceiverFaults) GetDcVoltageFault() bool {
	if x != nil {
		return x.DcVoltageFault
	}
	return false
}

var File_spacex_api_device_transceiver_proto protoreflect.FileDescriptor

var file_spacex_api_device_transceiver_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x70, 0x61, 0x63, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x53, 0x70, 0x61, 0x63, 0x65, 0x58, 0x2e, 0x41, 0x50,
	0x49, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x78,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x78,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x46, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x66, 0x5f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x66,
	0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x22, 0xc4, 0x01, 0x0a, 0x21, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x46, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61,
	0x63, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x62, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x62, 0x65, 0x72, 0x4c, 0x6f,
	0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x6e,
	0x72, 0x5f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x73, 0x6e, 0x72, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61,
	0x63, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x73, 0x73, 0x69, 0x5f, 0x6c,
	0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x10, 0x72, 0x73, 0x73, 0x69, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b,
	0x54, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x6c, 0x6c, 0x4c, 0x6f, 0x63, 0x6b, 0x22,
	0x1d, 0x0a, 0x1b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xea,
	0x04, 0x0a, 0x1c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x49, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x58, 0x2e, 0x41, 0x50, 0x49, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x08, 0x6d, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x4d, 0x0a, 0x0b, 0x64, 0x65,
	0x6d, 0x6f, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2c, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x58, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x64,
	0x65, 0x6d, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x74, 0x78, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x58, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x54, 0x78, 0x52, 0x78, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x07, 0x74, 0x78, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a,
	0x08, 0x72, 0x78, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x27, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x58, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x54,
	0x78, 0x52, 0x78, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x07, 0x72, 0x78, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0xee, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x58, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0xef, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x58,
	0x2e, 0x41, 0x50, 0x49, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x06, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x6c, 0x0a, 0x17, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69,
	0x74, 0x5f, 0x62, 0x6c, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0xf0, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x58,
	0x2e, 0x41, 0x50, 0x49, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x42,
	0x6c, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x15, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x6f, 0x64, 0x65, 0x6d, 0x5f, 0x61, 0x73, 0x69,
	0x63, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x18, 0xf1, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x6d,
	0x6f, 0x64, 0x65, 0x6d, 0x41, 0x73, 0x69, 0x63, 0x54, 0x65, 0x6d, 0x70, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x78, 0x5f, 0x69, 0x66, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x18, 0xf2, 0x07, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x74, 0x78, 0x49, 0x66, 0x54, 0x65, 0x6d, 0x70, 0x22, 0xaa, 0x01, 0x0a, 0x11,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x46, 0x61, 0x75, 0x6c, 0x74,
	0x73, 0x12, 0x3a, 0x0a, 0x1a, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6d, 0x5f, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x6f, 0x76, 0x65, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x4d,
	0x6f, 0x64, 0x65, 0x6d, 0x41, 0x73, 0x69, 0x63, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x2f, 0x0a,
	0x14, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x70, 0x63, 0x62, 0x61, 0x5f,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x6f, 0x76, 0x65,
	0x72, 0x54, 0x65, 0x6d, 0x70, 0x50, 0x63, 0x62, 0x61, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x28,
	0x0a, 0x10, 0x64, 0x63, 0x5f, 0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x64, 0x63, 0x56, 0x6f, 0x6c, 0x74,
	0x61, 0x67, 0x65, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x2a, 0x5e, 0x0a, 0x19, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x4f, 0x44, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4d,
	0x4f, 0x44, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x4f, 0x44, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x49,
	0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x4d, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x54, 0x78, 0x52, 0x78, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x10, 0x0a, 0x0c, 0x54, 0x58, 0x52, 0x58, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x58, 0x52, 0x58, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x58, 0x52, 0x58, 0x5f, 0x44, 0x49, 0x53,
	0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x53, 0x0a, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x42, 0x6c,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x54,
	0x42, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x54,
	0x42, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x54,
	0x42, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x42, 0x17, 0x5a, 0x15,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spacex_api_device_transceiver_proto_rawDescOnce sync.Once
	file_spacex_api_device_transceiver_proto_rawDescData = file_spacex_api_device_transceiver_proto_rawDesc
)

func file_spacex_api_device_transceiver_proto_rawDescGZIP() []byte {
	file_spacex_api_device_transceiver_proto_rawDescOnce.Do(func() {
		file_spacex_api_device_transceiver_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacex_api_device_transceiver_proto_rawDescData)
	})
	return file_spacex_api_device_transceiver_proto_rawDescData
}

var file_spacex_api_device_transceiver_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_spacex_api_device_transceiver_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spacex_api_device_transceiver_proto_goTypes = []interface{}{
	(TransceiverModulatorState)(0),            // 0: SpaceX.API.Device.TransceiverModulatorState
	(TransceiverTxRxState)(0),                 // 1: SpaceX.API.Device.TransceiverTxRxState
	(TransceiverTransmitBlankingState)(0),     // 2: SpaceX.API.Device.TransceiverTransmitBlankingState
	(*TransceiverIFLoopbackTestRequest)(nil),  // 3: SpaceX.API.Device.TransceiverIFLoopbackTestRequest
	(*TransceiverIFLoopbackTestResponse)(nil), // 4: SpaceX.API.Device.TransceiverIFLoopbackTestResponse
	(*TransceiverGetStatusRequest)(nil),       // 5: SpaceX.API.Device.TransceiverGetStatusRequest
	(*TransceiverGetStatusResponse)(nil),      // 6: SpaceX.API.Device.TransceiverGetStatusResponse
	(*TransceiverFaults)(nil),                 // 7: SpaceX.API.Device.TransceiverFaults
	(DishState)(0),                            // 8: SpaceX.API.Device.DishState
}
var file_spacex_api_device_transceiver_proto_depIdxs = []int32{
	0, // 0: SpaceX.API.Device.TransceiverGetStatusResponse.mod_state:type_name -> SpaceX.API.Device.TransceiverModulatorState
	0, // 1: SpaceX.API.Device.TransceiverGetStatusResponse.demod_state:type_name -> SpaceX.API.Device.TransceiverModulatorState
	1, // 2: SpaceX.API.Device.TransceiverGetStatusResponse.tx_state:type_name -> SpaceX.API.Device.TransceiverTxRxState
	1, // 3: SpaceX.API.Device.TransceiverGetStatusResponse.rx_state:type_name -> SpaceX.API.Device.TransceiverTxRxState
	8, // 4: SpaceX.API.Device.TransceiverGetStatusResponse.state:type_name -> SpaceX.API.Device.DishState
	7, // 5: SpaceX.API.Device.TransceiverGetStatusResponse.faults:type_name -> SpaceX.API.Device.TransceiverFaults
	2, // 6: SpaceX.API.Device.TransceiverGetStatusResponse.transmit_blanking_state:type_name -> SpaceX.API.Device.TransceiverTransmitBlankingState
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_spacex_api_device_transceiver_proto_init() }
func file_spacex_api_device_transceiver_proto_init() {
	if File_spacex_api_device_transceiver_proto != nil {
		return
	}
	file_spacex_api_device_common_proto_init()
	file_spacex_api_device_dish_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spacex_api_device_transceiver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransceiverIFLoopbackTestRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spacex_api_device_transceiver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransceiverIFLoopbackTestResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spacex_api_device_transceiver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransceiverGetStatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spacex_api_device_transceiver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransceiverGetStatusResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spacex_api_device_transceiver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransceiverFaults); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spacex_api_device_transceiver_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spacex_api_device_transceiver_proto_goTypes,
		DependencyIndexes: file_spacex_api_device_transceiver_proto_depIdxs,
		EnumInfos:         file_spacex_api_device_transceiver_proto_enumTypes,
		MessageInfos:      file_spacex_api_device_transceiver_proto_msgTypes,
	}.Build()
	File_spacex_api_device_transceiver_proto = out.File
	file_spacex_api_device_transceiver_proto_rawDesc = nil
	file_spacex_api_device_transceiver_proto_goTypes = nil
	file_spacex_api_device_transceiver_proto_depIdxs = nil
}
