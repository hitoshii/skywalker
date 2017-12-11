// Code generated by protoc-gen-go.
// source: src/skywalker/rpc/response.proto
// DO NOT EDIT!

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AuthResponse_Status int32

const (
	AuthResponse_SUCCESS AuthResponse_Status = 1
	AuthResponse_FAILURE AuthResponse_Status = 2
)

var AuthResponse_Status_name = map[int32]string{
	1: "SUCCESS",
	2: "FAILURE",
}
var AuthResponse_Status_value = map[string]int32{
	"SUCCESS": 1,
	"FAILURE": 2,
}

func (x AuthResponse_Status) Enum() *AuthResponse_Status {
	p := new(AuthResponse_Status)
	*p = x
	return p
}
func (x AuthResponse_Status) String() string {
	return proto.EnumName(AuthResponse_Status_name, int32(x))
}
func (x *AuthResponse_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AuthResponse_Status_value, data, "AuthResponse_Status")
	if err != nil {
		return err
	}
	*x = AuthResponse_Status(value)
	return nil
}
func (AuthResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

type StatusResponse_Status int32

const (
	StatusResponse_STOPPED StatusResponse_Status = 1
	StatusResponse_RUNNING StatusResponse_Status = 2
	StatusResponse_ERROR   StatusResponse_Status = 3
)

var StatusResponse_Status_name = map[int32]string{
	1: "STOPPED",
	2: "RUNNING",
	3: "ERROR",
}
var StatusResponse_Status_value = map[string]int32{
	"STOPPED": 1,
	"RUNNING": 2,
	"ERROR":   3,
}

func (x StatusResponse_Status) Enum() *StatusResponse_Status {
	p := new(StatusResponse_Status)
	*p = x
	return p
}
func (x StatusResponse_Status) String() string {
	return proto.EnumName(StatusResponse_Status_name, int32(x))
}
func (x *StatusResponse_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(StatusResponse_Status_value, data, "StatusResponse_Status")
	if err != nil {
		return err
	}
	*x = StatusResponse_Status(value)
	return nil
}
func (StatusResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

type StartResponse_Status int32

const (
	StartResponse_STARTED StartResponse_Status = 1
	StartResponse_ERROR   StartResponse_Status = 2
	StartResponse_RUNNING StartResponse_Status = 3
)

var StartResponse_Status_name = map[int32]string{
	1: "STARTED",
	2: "ERROR",
	3: "RUNNING",
}
var StartResponse_Status_value = map[string]int32{
	"STARTED": 1,
	"ERROR":   2,
	"RUNNING": 3,
}

func (x StartResponse_Status) Enum() *StartResponse_Status {
	p := new(StartResponse_Status)
	*p = x
	return p
}
func (x StartResponse_Status) String() string {
	return proto.EnumName(StartResponse_Status_name, int32(x))
}
func (x *StartResponse_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(StartResponse_Status_value, data, "StartResponse_Status")
	if err != nil {
		return err
	}
	*x = StartResponse_Status(value)
	return nil
}
func (StartResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3, 0} }

type StopResponse_Status int32

const (
	StopResponse_STOPPED   StopResponse_Status = 1
	StopResponse_ERROR     StopResponse_Status = 2
	StopResponse_UNRUNNING StopResponse_Status = 3
)

var StopResponse_Status_name = map[int32]string{
	1: "STOPPED",
	2: "ERROR",
	3: "UNRUNNING",
}
var StopResponse_Status_value = map[string]int32{
	"STOPPED":   1,
	"ERROR":     2,
	"UNRUNNING": 3,
}

func (x StopResponse_Status) Enum() *StopResponse_Status {
	p := new(StopResponse_Status)
	*p = x
	return p
}
func (x StopResponse_Status) String() string {
	return proto.EnumName(StopResponse_Status_name, int32(x))
}
func (x *StopResponse_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(StopResponse_Status_value, data, "StopResponse_Status")
	if err != nil {
		return err
	}
	*x = StopResponse_Status(value)
	return nil
}
func (StopResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{4, 0} }

type InfoResponse_Status int32

const (
	InfoResponse_STOPPED InfoResponse_Status = 1
	InfoResponse_RUNNING InfoResponse_Status = 2
	InfoResponse_ERROR   InfoResponse_Status = 3
)

var InfoResponse_Status_name = map[int32]string{
	1: "STOPPED",
	2: "RUNNING",
	3: "ERROR",
}
var InfoResponse_Status_value = map[string]int32{
	"STOPPED": 1,
	"RUNNING": 2,
	"ERROR":   3,
}

func (x InfoResponse_Status) Enum() *InfoResponse_Status {
	p := new(InfoResponse_Status)
	*p = x
	return p
}
func (x InfoResponse_Status) String() string {
	return proto.EnumName(InfoResponse_Status_name, int32(x))
}
func (x *InfoResponse_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(InfoResponse_Status_value, data, "InfoResponse_Status")
	if err != nil {
		return err
	}
	*x = InfoResponse_Status(value)
	return nil
}
func (InfoResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{5, 0} }

type QuitResponse_Status int32

const (
	QuitResponse_QUITED QuitResponse_Status = 1
	QuitResponse_ERROR  QuitResponse_Status = 2
)

var QuitResponse_Status_name = map[int32]string{
	1: "QUITED",
	2: "ERROR",
}
var QuitResponse_Status_value = map[string]int32{
	"QUITED": 1,
	"ERROR":  2,
}

func (x QuitResponse_Status) Enum() *QuitResponse_Status {
	p := new(QuitResponse_Status)
	*p = x
	return p
}
func (x QuitResponse_Status) String() string {
	return proto.EnumName(QuitResponse_Status_name, int32(x))
}
func (x *QuitResponse_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(QuitResponse_Status_value, data, "QuitResponse_Status")
	if err != nil {
		return err
	}
	*x = QuitResponse_Status(value)
	return nil
}
func (QuitResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{7, 0} }

type ClearCacheResponse_Status int32

const (
	ClearCacheResponse_SUCCESS ClearCacheResponse_Status = 1
)

var ClearCacheResponse_Status_name = map[int32]string{
	1: "SUCCESS",
}
var ClearCacheResponse_Status_value = map[string]int32{
	"SUCCESS": 1,
}

func (x ClearCacheResponse_Status) Enum() *ClearCacheResponse_Status {
	p := new(ClearCacheResponse_Status)
	*p = x
	return p
}
func (x ClearCacheResponse_Status) String() string {
	return proto.EnumName(ClearCacheResponse_Status_name, int32(x))
}
func (x *ClearCacheResponse_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClearCacheResponse_Status_value, data, "ClearCacheResponse_Status")
	if err != nil {
		return err
	}
	*x = ClearCacheResponse_Status(value)
	return nil
}
func (ClearCacheResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{8, 0} }

// 出错
type Error struct {
	Msg              *string `protobuf:"bytes,1,req,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Error) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

type AuthResponse struct {
	Status           *AuthResponse_Status `protobuf:"varint,1,req,name=status,enum=rpc.AuthResponse_Status" json:"status,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *AuthResponse) Reset()                    { *m = AuthResponse{} }
func (m *AuthResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()               {}
func (*AuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *AuthResponse) GetStatus() AuthResponse_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return AuthResponse_SUCCESS
}

// status 命令的返回结果
type StatusResponse struct {
	Data             []*StatusResponse_Data `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *StatusResponse) GetData() []*StatusResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type StatusResponse_Data struct {
	Name             *string                `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Cname            *string                `protobuf:"bytes,2,opt,name=cname,def=" json:"cname,omitempty"`
	Sname            *string                `protobuf:"bytes,3,opt,name=sname,def=" json:"sname,omitempty"`
	Status           *StatusResponse_Status `protobuf:"varint,4,req,name=status,enum=rpc.StatusResponse_Status" json:"status,omitempty"`
	BindAddr         *string                `protobuf:"bytes,5,opt,name=bindAddr,def=" json:"bindAddr,omitempty"`
	BindPort         *int32                 `protobuf:"varint,6,opt,name=bindPort,def=0" json:"bindPort,omitempty"`
	StartTime        *int64                 `protobuf:"varint,7,opt,name=startTime,def=0" json:"startTime,omitempty"`
	Err              *string                `protobuf:"bytes,8,opt,name=err,def=" json:"err,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *StatusResponse_Data) Reset()                    { *m = StatusResponse_Data{} }
func (m *StatusResponse_Data) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse_Data) ProtoMessage()               {}
func (*StatusResponse_Data) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

const Default_StatusResponse_Data_BindPort int32 = 0
const Default_StatusResponse_Data_StartTime int64 = 0

func (m *StatusResponse_Data) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *StatusResponse_Data) GetCname() string {
	if m != nil && m.Cname != nil {
		return *m.Cname
	}
	return ""
}

func (m *StatusResponse_Data) GetSname() string {
	if m != nil && m.Sname != nil {
		return *m.Sname
	}
	return ""
}

func (m *StatusResponse_Data) GetStatus() StatusResponse_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return StatusResponse_STOPPED
}

func (m *StatusResponse_Data) GetBindAddr() string {
	if m != nil && m.BindAddr != nil {
		return *m.BindAddr
	}
	return ""
}

func (m *StatusResponse_Data) GetBindPort() int32 {
	if m != nil && m.BindPort != nil {
		return *m.BindPort
	}
	return Default_StatusResponse_Data_BindPort
}

func (m *StatusResponse_Data) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return Default_StatusResponse_Data_StartTime
}

func (m *StatusResponse_Data) GetErr() string {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return ""
}

// start 命令的返回结果
type StartResponse struct {
	Data             []*StartResponse_Data `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *StartResponse) Reset()                    { *m = StartResponse{} }
func (m *StartResponse) String() string            { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()               {}
func (*StartResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *StartResponse) GetData() []*StartResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type StartResponse_Data struct {
	Name             *string               `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Status           *StartResponse_Status `protobuf:"varint,2,req,name=status,enum=rpc.StartResponse_Status" json:"status,omitempty"`
	Err              *string               `protobuf:"bytes,3,opt,name=err,def=" json:"err,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *StartResponse_Data) Reset()                    { *m = StartResponse_Data{} }
func (m *StartResponse_Data) String() string            { return proto.CompactTextString(m) }
func (*StartResponse_Data) ProtoMessage()               {}
func (*StartResponse_Data) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3, 0} }

func (m *StartResponse_Data) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *StartResponse_Data) GetStatus() StartResponse_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return StartResponse_STARTED
}

func (m *StartResponse_Data) GetErr() string {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return ""
}

type StopResponse struct {
	Data             []*StopResponse_Data `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *StopResponse) Reset()                    { *m = StopResponse{} }
func (m *StopResponse) String() string            { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()               {}
func (*StopResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *StopResponse) GetData() []*StopResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type StopResponse_Data struct {
	Name             *string              `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Status           *StopResponse_Status `protobuf:"varint,2,req,name=status,enum=rpc.StopResponse_Status" json:"status,omitempty"`
	Err              *string              `protobuf:"bytes,3,opt,name=err,def=" json:"err,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *StopResponse_Data) Reset()                    { *m = StopResponse_Data{} }
func (m *StopResponse_Data) String() string            { return proto.CompactTextString(m) }
func (*StopResponse_Data) ProtoMessage()               {}
func (*StopResponse_Data) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4, 0} }

func (m *StopResponse_Data) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *StopResponse_Data) GetStatus() StopResponse_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return StopResponse_STOPPED
}

func (m *StopResponse_Data) GetErr() string {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return ""
}

type InfoResponse struct {
	Data             []*InfoResponse_Data `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *InfoResponse) Reset()                    { *m = InfoResponse{} }
func (m *InfoResponse) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()               {}
func (*InfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *InfoResponse) GetData() []*InfoResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type InfoResponse_Info struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InfoResponse_Info) Reset()                    { *m = InfoResponse_Info{} }
func (m *InfoResponse_Info) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse_Info) ProtoMessage()               {}
func (*InfoResponse_Info) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5, 0} }

func (m *InfoResponse_Info) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *InfoResponse_Info) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type InfoResponse_Data struct {
	Name             *string              `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Cname            *string              `protobuf:"bytes,2,opt,name=cname,def=" json:"cname,omitempty"`
	Sname            *string              `protobuf:"bytes,3,opt,name=sname,def=" json:"sname,omitempty"`
	Status           *InfoResponse_Status `protobuf:"varint,4,req,name=status,enum=rpc.InfoResponse_Status" json:"status,omitempty"`
	BindAddr         *string              `protobuf:"bytes,5,opt,name=bindAddr,def=" json:"bindAddr,omitempty"`
	BindPort         *int32               `protobuf:"varint,6,opt,name=bindPort,def=0" json:"bindPort,omitempty"`
	StartTime        *int64               `protobuf:"varint,7,opt,name=startTime,def=0" json:"startTime,omitempty"`
	Sent             *int64               `protobuf:"varint,8,opt,name=sent,def=0" json:"sent,omitempty"`
	Received         *int64               `protobuf:"varint,9,opt,name=received,def=0" json:"received,omitempty"`
	SentRate         *int64               `protobuf:"varint,10,opt,name=sentRate,def=0" json:"sentRate,omitempty"`
	ReceivedRate     *int64               `protobuf:"varint,11,opt,name=receivedRate,def=0" json:"receivedRate,omitempty"`
	CaInfo           []*InfoResponse_Info `protobuf:"bytes,12,rep,name=caInfo" json:"caInfo,omitempty"`
	SaInfo           []*InfoResponse_Info `protobuf:"bytes,13,rep,name=saInfo" json:"saInfo,omitempty"`
	Err              *string              `protobuf:"bytes,14,opt,name=err,def=" json:"err,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *InfoResponse_Data) Reset()                    { *m = InfoResponse_Data{} }
func (m *InfoResponse_Data) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse_Data) ProtoMessage()               {}
func (*InfoResponse_Data) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5, 1} }

const Default_InfoResponse_Data_BindPort int32 = 0
const Default_InfoResponse_Data_StartTime int64 = 0
const Default_InfoResponse_Data_Sent int64 = 0
const Default_InfoResponse_Data_Received int64 = 0
const Default_InfoResponse_Data_SentRate int64 = 0
const Default_InfoResponse_Data_ReceivedRate int64 = 0

func (m *InfoResponse_Data) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *InfoResponse_Data) GetCname() string {
	if m != nil && m.Cname != nil {
		return *m.Cname
	}
	return ""
}

func (m *InfoResponse_Data) GetSname() string {
	if m != nil && m.Sname != nil {
		return *m.Sname
	}
	return ""
}

func (m *InfoResponse_Data) GetStatus() InfoResponse_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return InfoResponse_STOPPED
}

func (m *InfoResponse_Data) GetBindAddr() string {
	if m != nil && m.BindAddr != nil {
		return *m.BindAddr
	}
	return ""
}

func (m *InfoResponse_Data) GetBindPort() int32 {
	if m != nil && m.BindPort != nil {
		return *m.BindPort
	}
	return Default_InfoResponse_Data_BindPort
}

func (m *InfoResponse_Data) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return Default_InfoResponse_Data_StartTime
}

func (m *InfoResponse_Data) GetSent() int64 {
	if m != nil && m.Sent != nil {
		return *m.Sent
	}
	return Default_InfoResponse_Data_Sent
}

func (m *InfoResponse_Data) GetReceived() int64 {
	if m != nil && m.Received != nil {
		return *m.Received
	}
	return Default_InfoResponse_Data_Received
}

func (m *InfoResponse_Data) GetSentRate() int64 {
	if m != nil && m.SentRate != nil {
		return *m.SentRate
	}
	return Default_InfoResponse_Data_SentRate
}

func (m *InfoResponse_Data) GetReceivedRate() int64 {
	if m != nil && m.ReceivedRate != nil {
		return *m.ReceivedRate
	}
	return Default_InfoResponse_Data_ReceivedRate
}

func (m *InfoResponse_Data) GetCaInfo() []*InfoResponse_Info {
	if m != nil {
		return m.CaInfo
	}
	return nil
}

func (m *InfoResponse_Data) GetSaInfo() []*InfoResponse_Info {
	if m != nil {
		return m.SaInfo
	}
	return nil
}

func (m *InfoResponse_Data) GetErr() string {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return ""
}

type ReloadResponse struct {
	Unchanged        []string `protobuf:"bytes,1,rep,name=unchanged" json:"unchanged,omitempty"`
	Added            []string `protobuf:"bytes,2,rep,name=added" json:"added,omitempty"`
	Deleted          []string `protobuf:"bytes,3,rep,name=deleted" json:"deleted,omitempty"`
	Updated          []string `protobuf:"bytes,4,rep,name=updated" json:"updated,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ReloadResponse) Reset()                    { *m = ReloadResponse{} }
func (m *ReloadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReloadResponse) ProtoMessage()               {}
func (*ReloadResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ReloadResponse) GetUnchanged() []string {
	if m != nil {
		return m.Unchanged
	}
	return nil
}

func (m *ReloadResponse) GetAdded() []string {
	if m != nil {
		return m.Added
	}
	return nil
}

func (m *ReloadResponse) GetDeleted() []string {
	if m != nil {
		return m.Deleted
	}
	return nil
}

func (m *ReloadResponse) GetUpdated() []string {
	if m != nil {
		return m.Updated
	}
	return nil
}

type QuitResponse struct {
	Status           *QuitResponse_Status `protobuf:"varint,1,req,name=status,enum=rpc.QuitResponse_Status" json:"status,omitempty"`
	Pid              *uint32              `protobuf:"varint,2,req,name=pid" json:"pid,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *QuitResponse) Reset()                    { *m = QuitResponse{} }
func (m *QuitResponse) String() string            { return proto.CompactTextString(m) }
func (*QuitResponse) ProtoMessage()               {}
func (*QuitResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *QuitResponse) GetStatus() QuitResponse_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return QuitResponse_QUITED
}

func (m *QuitResponse) GetPid() uint32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

type ClearCacheResponse struct {
	Status           *ClearCacheResponse_Status `protobuf:"varint,1,req,name=status,enum=rpc.ClearCacheResponse_Status" json:"status,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *ClearCacheResponse) Reset()                    { *m = ClearCacheResponse{} }
func (m *ClearCacheResponse) String() string            { return proto.CompactTextString(m) }
func (*ClearCacheResponse) ProtoMessage()               {}
func (*ClearCacheResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *ClearCacheResponse) GetStatus() ClearCacheResponse_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ClearCacheResponse_SUCCESS
}

type Response struct {
	Type             *RequestType        `protobuf:"varint,1,req,name=type,enum=rpc.RequestType" json:"type,omitempty"`
	Err              *Error              `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
	Auth             *AuthResponse       `protobuf:"bytes,3,opt,name=auth" json:"auth,omitempty"`
	Status           *StatusResponse     `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	Start            *StartResponse      `protobuf:"bytes,5,opt,name=start" json:"start,omitempty"`
	Stop             *StopResponse       `protobuf:"bytes,6,opt,name=stop" json:"stop,omitempty"`
	Info             *InfoResponse       `protobuf:"bytes,7,opt,name=info" json:"info,omitempty"`
	Reload           *ReloadResponse     `protobuf:"bytes,8,opt,name=reload" json:"reload,omitempty"`
	Quit             *QuitResponse       `protobuf:"bytes,9,opt,name=quit" json:"quit,omitempty"`
	Clear            *ClearCacheResponse `protobuf:"bytes,10,opt,name=clear" json:"clear,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *Response) GetType() RequestType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return RequestType_AUTH
}

func (m *Response) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *Response) GetAuth() *AuthResponse {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *Response) GetStatus() *StatusResponse {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Response) GetStart() *StartResponse {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Response) GetStop() *StopResponse {
	if m != nil {
		return m.Stop
	}
	return nil
}

func (m *Response) GetInfo() *InfoResponse {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Response) GetReload() *ReloadResponse {
	if m != nil {
		return m.Reload
	}
	return nil
}

func (m *Response) GetQuit() *QuitResponse {
	if m != nil {
		return m.Quit
	}
	return nil
}

func (m *Response) GetClear() *ClearCacheResponse {
	if m != nil {
		return m.Clear
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "rpc.Error")
	proto.RegisterType((*AuthResponse)(nil), "rpc.AuthResponse")
	proto.RegisterType((*StatusResponse)(nil), "rpc.StatusResponse")
	proto.RegisterType((*StatusResponse_Data)(nil), "rpc.StatusResponse.Data")
	proto.RegisterType((*StartResponse)(nil), "rpc.StartResponse")
	proto.RegisterType((*StartResponse_Data)(nil), "rpc.StartResponse.Data")
	proto.RegisterType((*StopResponse)(nil), "rpc.StopResponse")
	proto.RegisterType((*StopResponse_Data)(nil), "rpc.StopResponse.Data")
	proto.RegisterType((*InfoResponse)(nil), "rpc.InfoResponse")
	proto.RegisterType((*InfoResponse_Info)(nil), "rpc.InfoResponse.Info")
	proto.RegisterType((*InfoResponse_Data)(nil), "rpc.InfoResponse.Data")
	proto.RegisterType((*ReloadResponse)(nil), "rpc.ReloadResponse")
	proto.RegisterType((*QuitResponse)(nil), "rpc.QuitResponse")
	proto.RegisterType((*ClearCacheResponse)(nil), "rpc.ClearCacheResponse")
	proto.RegisterType((*Response)(nil), "rpc.Response")
	proto.RegisterEnum("rpc.AuthResponse_Status", AuthResponse_Status_name, AuthResponse_Status_value)
	proto.RegisterEnum("rpc.StatusResponse_Status", StatusResponse_Status_name, StatusResponse_Status_value)
	proto.RegisterEnum("rpc.StartResponse_Status", StartResponse_Status_name, StartResponse_Status_value)
	proto.RegisterEnum("rpc.StopResponse_Status", StopResponse_Status_name, StopResponse_Status_value)
	proto.RegisterEnum("rpc.InfoResponse_Status", InfoResponse_Status_name, InfoResponse_Status_value)
	proto.RegisterEnum("rpc.QuitResponse_Status", QuitResponse_Status_name, QuitResponse_Status_value)
	proto.RegisterEnum("rpc.ClearCacheResponse_Status", ClearCacheResponse_Status_name, ClearCacheResponse_Status_value)
}

func init() { proto.RegisterFile("src/skywalker/rpc/response.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 847 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xdd, 0x8e, 0x2a, 0x45,
	0x10, 0x76, 0xfe, 0x60, 0xa7, 0xf8, 0x09, 0xb6, 0x9e, 0xe3, 0x1c, 0xb2, 0xba, 0x64, 0xe2, 0x49,
	0x88, 0xc7, 0x05, 0xe4, 0xc2, 0x8b, 0xbd, 0x23, 0x2c, 0x1a, 0x12, 0xc3, 0xee, 0x36, 0x70, 0x6f,
	0x3b, 0xd3, 0x2e, 0x04, 0x96, 0x99, 0xed, 0xe9, 0x59, 0xe5, 0x61, 0x7c, 0x21, 0x13, 0x7d, 0x06,
	0xaf, 0xbc, 0xf6, 0x11, 0x4c, 0xd7, 0xfc, 0x30, 0xfc, 0x88, 0x26, 0x1a, 0xef, 0xa6, 0xea, 0xfb,
	0xa6, 0xaa, 0xbf, 0xea, 0xaa, 0x2e, 0x68, 0x45, 0xc2, 0xeb, 0x46, 0xab, 0xed, 0x0f, 0x6c, 0xbd,
	0xe2, 0xa2, 0x2b, 0x42, 0xaf, 0x2b, 0x78, 0x14, 0x06, 0x9b, 0x88, 0x77, 0x42, 0x11, 0xc8, 0x80,
	0x18, 0x22, 0xf4, 0x9a, 0x57, 0xa7, 0x68, 0xcf, 0x31, 0x8f, 0x64, 0xc2, 0x72, 0xdf, 0x80, 0x35,
	0x12, 0x22, 0x10, 0xa4, 0x01, 0xc6, 0x53, 0xf4, 0xe8, 0x68, 0x2d, 0xbd, 0x6d, 0x53, 0xf5, 0xe9,
	0xfa, 0x50, 0x1d, 0xc4, 0x72, 0x41, 0xd3, 0xb0, 0xa4, 0x07, 0xa5, 0x48, 0x32, 0x19, 0x47, 0x48,
	0xaa, 0xf7, 0x9d, 0x8e, 0x08, 0xbd, 0x4e, 0x91, 0xd2, 0x99, 0x22, 0x4e, 0x53, 0x9e, 0xeb, 0x42,
	0x29, 0xf1, 0x90, 0x0a, 0x94, 0xa7, 0xf3, 0xe1, 0x70, 0x34, 0x9d, 0x36, 0x34, 0x65, 0x7c, 0x35,
	0x18, 0x7f, 0x33, 0xa7, 0xa3, 0x86, 0xee, 0xfe, 0xa6, 0x43, 0x3d, 0xfd, 0x2d, 0x4b, 0xf4, 0x39,
	0x98, 0x3e, 0x93, 0xcc, 0xd1, 0x5a, 0x46, 0xbb, 0x92, 0xa6, 0xd9, 0xa7, 0x74, 0x6e, 0x99, 0x64,
	0x14, 0x59, 0xcd, 0x3f, 0x34, 0x30, 0x95, 0x49, 0x08, 0x98, 0x1b, 0xf6, 0xc4, 0x53, 0x09, 0xf8,
	0x4d, 0x5e, 0x83, 0xe5, 0xa1, 0x53, 0x6f, 0x69, 0x6d, 0xfb, 0xe6, 0x3d, 0x9a, 0x98, 0xca, 0x1f,
	0xa1, 0xdf, 0xc8, 0xfc, 0x68, 0x92, 0x7e, 0xae, 0xd1, 0x44, 0x8d, 0xcd, 0x53, 0xc9, 0xf7, 0x55,
	0x92, 0x4b, 0xb8, 0xf8, 0x6e, 0xb9, 0xf1, 0x07, 0xbe, 0x2f, 0x1c, 0x2b, 0x0d, 0x97, 0x7b, 0xc8,
	0xc7, 0x09, 0x7a, 0x1f, 0x08, 0xe9, 0x94, 0x5a, 0x5a, 0xdb, 0xba, 0xd1, 0x7a, 0x34, 0x77, 0x91,
	0x2b, 0xb0, 0x23, 0xc9, 0x84, 0x9c, 0x2d, 0x9f, 0xb8, 0x53, 0x6e, 0x69, 0x6d, 0x43, 0xe1, 0x3b,
	0x1f, 0x21, 0x60, 0x70, 0x21, 0x9c, 0x8b, 0x34, 0xb0, 0x32, 0xdc, 0xeb, 0xbd, 0xba, 0xce, 0xee,
	0xee, 0xef, 0x47, 0xb7, 0x49, 0x5d, 0xe9, 0x7c, 0x32, 0x19, 0x4f, 0xbe, 0x6e, 0xe8, 0xc4, 0x06,
	0x6b, 0x44, 0xe9, 0x1d, 0x6d, 0x18, 0xee, 0xaf, 0x1a, 0xd4, 0xa6, 0x2a, 0x60, 0x5e, 0xe1, 0x77,
	0x7b, 0x15, 0xfe, 0x28, 0x13, 0xb9, 0x63, 0x14, 0x0b, 0xcc, 0xce, 0xd4, 0xf7, 0x8b, 0xbc, 0x5e,
	0x3a, 0xd6, 0xeb, 0xcd, 0x89, 0x50, 0x07, 0xe5, 0x4a, 0x05, 0x19, 0x67, 0x04, 0x0d, 0xe8, 0x0c,
	0x05, 0xe5, 0x1a, 0xf4, 0xa2, 0x36, 0xc3, 0xfd, 0x45, 0x83, 0xea, 0x54, 0x06, 0x61, 0xae, 0xe7,
	0xb3, 0x3d, 0x3d, 0xaf, 0xd3, 0x43, 0xec, 0x08, 0x45, 0x39, 0xdf, 0x9e, 0x91, 0xd3, 0x3b, 0x90,
	0xe3, 0x1c, 0x47, 0xfa, 0x07, 0x6a, 0xba, 0xa7, 0xaf, 0xa7, 0xa0, 0xa6, 0x06, 0xf6, 0x7c, 0xb2,
	0xd3, 0xf3, 0xbb, 0x09, 0xd5, 0xf1, 0xe6, 0xfb, 0xe0, 0xac, 0x9e, 0x22, 0xa1, 0xa8, 0xa7, 0x03,
	0xa6, 0x82, 0xd4, 0x00, 0xaf, 0xf8, 0x36, 0x1b, 0xe0, 0x15, 0xdf, 0x92, 0x0f, 0xc1, 0x7a, 0x61,
	0xeb, 0x98, 0xa3, 0x18, 0x9b, 0x26, 0x46, 0xf3, 0x67, 0xe3, 0x3f, 0x9c, 0x97, 0xde, 0xc1, 0xbc,
	0x38, 0xc7, 0x47, 0xfd, 0x5f, 0xa7, 0xe5, 0x15, 0x98, 0x11, 0xdf, 0x48, 0x1c, 0x17, 0xc4, 0xd0,
	0x54, 0x61, 0x05, 0xf7, 0xf8, 0xf2, 0x85, 0xfb, 0x8e, 0x9d, 0x41, 0xb9, 0x4b, 0xc1, 0x8a, 0x46,
	0x99, 0xe4, 0x0e, 0xe4, 0x70, 0xe6, 0x22, 0x6f, 0xa1, 0x9a, 0x51, 0x91, 0x52, 0xc9, 0x28, 0x7b,
	0x6e, 0xd2, 0x81, 0x92, 0xc7, 0x94, 0x74, 0xa7, 0xfa, 0x57, 0xd7, 0x86, 0x46, 0xca, 0x52, 0xfc,
	0x28, 0xe1, 0xd7, 0xce, 0xf3, 0x13, 0x56, 0xd6, 0x6a, 0xf5, 0x7f, 0xf1, 0x12, 0xfc, 0x08, 0x75,
	0xca, 0xd7, 0x01, 0xf3, 0xf3, 0x4e, 0xbb, 0x04, 0x3b, 0xde, 0x78, 0x0b, 0xb6, 0x79, 0xe4, 0x3e,
	0xb6, 0x9b, 0x4d, 0x77, 0x0e, 0xd5, 0x41, 0xcc, 0xf7, 0xb9, 0xef, 0xe8, 0x88, 0x24, 0x06, 0x71,
	0xa0, 0xec, 0xf3, 0x35, 0x97, 0xdc, 0x77, 0x0c, 0xf4, 0x67, 0xa6, 0x42, 0xe2, 0xd0, 0x67, 0x0a,
	0x31, 0x13, 0x24, 0x35, 0xdd, 0x08, 0xaa, 0x0f, 0xf1, 0x52, 0xfe, 0xcd, 0x32, 0x29, 0x52, 0x0e,
	0x1b, 0xa7, 0x01, 0x46, 0xb8, 0xf4, 0xb1, 0x97, 0x6b, 0x54, 0x7d, 0xba, 0x57, 0xb9, 0x78, 0x80,
	0xd2, 0xc3, 0x7c, 0x7c, 0xf8, 0x68, 0xb8, 0x1e, 0x90, 0xe1, 0x9a, 0x33, 0x31, 0x64, 0xde, 0x82,
	0xe7, 0xa9, 0xbf, 0x3c, 0x48, 0xfd, 0x09, 0xa6, 0x3e, 0x26, 0x1e, 0x6e, 0xb3, 0x57, 0x27, 0xb7,
	0x99, 0xfb, 0x93, 0x01, 0x17, 0x79, 0xec, 0x4f, 0xc1, 0x94, 0xdb, 0x90, 0xa7, 0x91, 0x1b, 0x18,
	0x99, 0x26, 0x0b, 0x77, 0xb6, 0x0d, 0x39, 0x45, 0x94, 0x5c, 0x26, 0x37, 0xa9, 0x66, 0xac, 0xd2,
	0x07, 0x24, 0xe1, 0x12, 0xc6, 0x3b, 0x25, 0x6f, 0xc1, 0x64, 0xb1, 0x5c, 0xe0, 0xa8, 0x55, 0xfa,
	0xef, 0x1f, 0x6d, 0x59, 0x8a, 0x30, 0x79, 0x57, 0x18, 0x3d, 0x45, 0xfc, 0xe0, 0xc4, 0xaa, 0xca,
	0x8b, 0xd7, 0x06, 0x0b, 0x87, 0x04, 0x47, 0xae, 0xd2, 0x27, 0xc7, 0xcf, 0x34, 0x4d, 0x08, 0x2a,
	0x7b, 0x24, 0x83, 0x10, 0xa7, 0x2f, 0xcb, 0x5e, 0x7c, 0x00, 0x29, 0xc2, 0x8a, 0xb6, 0x54, 0xad,
	0x5b, 0x2e, 0xd0, 0x8a, 0xad, 0x4b, 0x11, 0x56, 0x87, 0x14, 0xd8, 0x70, 0x38, 0x91, 0xd9, 0x21,
	0xf7, 0x7b, 0x90, 0xa6, 0x14, 0x15, 0xf3, 0x39, 0x5e, 0x4a, 0x9c, 0xd0, 0x2c, 0x66, 0xb1, 0x23,
	0x28, 0xc2, 0xe4, 0x1a, 0x2c, 0x4f, 0x5d, 0x16, 0x8e, 0x6a, 0xb6, 0xbd, 0x8e, 0xaf, 0x8f, 0x26,
	0xac, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x79, 0x6e, 0x33, 0x2a, 0x09, 0x00, 0x00,
}
