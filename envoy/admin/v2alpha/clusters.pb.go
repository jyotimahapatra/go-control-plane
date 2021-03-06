// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/clusters.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Clusters struct {
	ClusterStatuses      []*ClusterStatus `protobuf:"bytes,1,rep,name=cluster_statuses,json=clusterStatuses,proto3" json:"cluster_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Clusters) Reset()         { *m = Clusters{} }
func (m *Clusters) String() string { return proto.CompactTextString(m) }
func (*Clusters) ProtoMessage()    {}
func (*Clusters) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6251a3a957f478b, []int{0}
}

func (m *Clusters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Clusters.Unmarshal(m, b)
}
func (m *Clusters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Clusters.Marshal(b, m, deterministic)
}
func (m *Clusters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clusters.Merge(m, src)
}
func (m *Clusters) XXX_Size() int {
	return xxx_messageInfo_Clusters.Size(m)
}
func (m *Clusters) XXX_DiscardUnknown() {
	xxx_messageInfo_Clusters.DiscardUnknown(m)
}

var xxx_messageInfo_Clusters proto.InternalMessageInfo

func (m *Clusters) GetClusterStatuses() []*ClusterStatus {
	if m != nil {
		return m.ClusterStatuses
	}
	return nil
}

type ClusterStatus struct {
	Name                                    string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AddedViaApi                             bool           `protobuf:"varint,2,opt,name=added_via_api,json=addedViaApi,proto3" json:"added_via_api,omitempty"`
	SuccessRateEjectionThreshold            *_type.Percent `protobuf:"bytes,3,opt,name=success_rate_ejection_threshold,json=successRateEjectionThreshold,proto3" json:"success_rate_ejection_threshold,omitempty"`
	HostStatuses                            []*HostStatus  `protobuf:"bytes,4,rep,name=host_statuses,json=hostStatuses,proto3" json:"host_statuses,omitempty"`
	LocalOriginSuccessRateEjectionThreshold *_type.Percent `protobuf:"bytes,5,opt,name=local_origin_success_rate_ejection_threshold,json=localOriginSuccessRateEjectionThreshold,proto3" json:"local_origin_success_rate_ejection_threshold,omitempty"`
	XXX_NoUnkeyedLiteral                    struct{}       `json:"-"`
	XXX_unrecognized                        []byte         `json:"-"`
	XXX_sizecache                           int32          `json:"-"`
}

func (m *ClusterStatus) Reset()         { *m = ClusterStatus{} }
func (m *ClusterStatus) String() string { return proto.CompactTextString(m) }
func (*ClusterStatus) ProtoMessage()    {}
func (*ClusterStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6251a3a957f478b, []int{1}
}

func (m *ClusterStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterStatus.Unmarshal(m, b)
}
func (m *ClusterStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterStatus.Marshal(b, m, deterministic)
}
func (m *ClusterStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterStatus.Merge(m, src)
}
func (m *ClusterStatus) XXX_Size() int {
	return xxx_messageInfo_ClusterStatus.Size(m)
}
func (m *ClusterStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterStatus proto.InternalMessageInfo

func (m *ClusterStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterStatus) GetAddedViaApi() bool {
	if m != nil {
		return m.AddedViaApi
	}
	return false
}

func (m *ClusterStatus) GetSuccessRateEjectionThreshold() *_type.Percent {
	if m != nil {
		return m.SuccessRateEjectionThreshold
	}
	return nil
}

func (m *ClusterStatus) GetHostStatuses() []*HostStatus {
	if m != nil {
		return m.HostStatuses
	}
	return nil
}

func (m *ClusterStatus) GetLocalOriginSuccessRateEjectionThreshold() *_type.Percent {
	if m != nil {
		return m.LocalOriginSuccessRateEjectionThreshold
	}
	return nil
}

type HostStatus struct {
	Address                *core.Address     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Stats                  []*SimpleMetric   `protobuf:"bytes,2,rep,name=stats,proto3" json:"stats,omitempty"`
	HealthStatus           *HostHealthStatus `protobuf:"bytes,3,opt,name=health_status,json=healthStatus,proto3" json:"health_status,omitempty"`
	SuccessRate            *_type.Percent    `protobuf:"bytes,4,opt,name=success_rate,json=successRate,proto3" json:"success_rate,omitempty"`
	Weight                 uint32            `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	Hostname               string            `protobuf:"bytes,6,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Priority               uint32            `protobuf:"varint,7,opt,name=priority,proto3" json:"priority,omitempty"`
	LocalOriginSuccessRate *_type.Percent    `protobuf:"bytes,8,opt,name=local_origin_success_rate,json=localOriginSuccessRate,proto3" json:"local_origin_success_rate,omitempty"`
	Locality               *core.Locality    `protobuf:"bytes,9,opt,name=locality,proto3" json:"locality,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}          `json:"-"`
	XXX_unrecognized       []byte            `json:"-"`
	XXX_sizecache          int32             `json:"-"`
}

func (m *HostStatus) Reset()         { *m = HostStatus{} }
func (m *HostStatus) String() string { return proto.CompactTextString(m) }
func (*HostStatus) ProtoMessage()    {}
func (*HostStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6251a3a957f478b, []int{2}
}

func (m *HostStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostStatus.Unmarshal(m, b)
}
func (m *HostStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostStatus.Marshal(b, m, deterministic)
}
func (m *HostStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostStatus.Merge(m, src)
}
func (m *HostStatus) XXX_Size() int {
	return xxx_messageInfo_HostStatus.Size(m)
}
func (m *HostStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HostStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HostStatus proto.InternalMessageInfo

func (m *HostStatus) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *HostStatus) GetStats() []*SimpleMetric {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *HostStatus) GetHealthStatus() *HostHealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return nil
}

func (m *HostStatus) GetSuccessRate() *_type.Percent {
	if m != nil {
		return m.SuccessRate
	}
	return nil
}

func (m *HostStatus) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *HostStatus) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *HostStatus) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *HostStatus) GetLocalOriginSuccessRate() *_type.Percent {
	if m != nil {
		return m.LocalOriginSuccessRate
	}
	return nil
}

func (m *HostStatus) GetLocality() *core.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

type HostHealthStatus struct {
	FailedActiveHealthCheck   bool              `protobuf:"varint,1,opt,name=failed_active_health_check,json=failedActiveHealthCheck,proto3" json:"failed_active_health_check,omitempty"`
	FailedOutlierCheck        bool              `protobuf:"varint,2,opt,name=failed_outlier_check,json=failedOutlierCheck,proto3" json:"failed_outlier_check,omitempty"`
	FailedActiveDegradedCheck bool              `protobuf:"varint,4,opt,name=failed_active_degraded_check,json=failedActiveDegradedCheck,proto3" json:"failed_active_degraded_check,omitempty"`
	PendingDynamicRemoval     bool              `protobuf:"varint,5,opt,name=pending_dynamic_removal,json=pendingDynamicRemoval,proto3" json:"pending_dynamic_removal,omitempty"`
	PendingActiveHc           bool              `protobuf:"varint,6,opt,name=pending_active_hc,json=pendingActiveHc,proto3" json:"pending_active_hc,omitempty"`
	EdsHealthStatus           core.HealthStatus `protobuf:"varint,3,opt,name=eds_health_status,json=edsHealthStatus,proto3,enum=envoy.api.v2.core.HealthStatus" json:"eds_health_status,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}          `json:"-"`
	XXX_unrecognized          []byte            `json:"-"`
	XXX_sizecache             int32             `json:"-"`
}

func (m *HostHealthStatus) Reset()         { *m = HostHealthStatus{} }
func (m *HostHealthStatus) String() string { return proto.CompactTextString(m) }
func (*HostHealthStatus) ProtoMessage()    {}
func (*HostHealthStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6251a3a957f478b, []int{3}
}

func (m *HostHealthStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostHealthStatus.Unmarshal(m, b)
}
func (m *HostHealthStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostHealthStatus.Marshal(b, m, deterministic)
}
func (m *HostHealthStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostHealthStatus.Merge(m, src)
}
func (m *HostHealthStatus) XXX_Size() int {
	return xxx_messageInfo_HostHealthStatus.Size(m)
}
func (m *HostHealthStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HostHealthStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HostHealthStatus proto.InternalMessageInfo

func (m *HostHealthStatus) GetFailedActiveHealthCheck() bool {
	if m != nil {
		return m.FailedActiveHealthCheck
	}
	return false
}

func (m *HostHealthStatus) GetFailedOutlierCheck() bool {
	if m != nil {
		return m.FailedOutlierCheck
	}
	return false
}

func (m *HostHealthStatus) GetFailedActiveDegradedCheck() bool {
	if m != nil {
		return m.FailedActiveDegradedCheck
	}
	return false
}

func (m *HostHealthStatus) GetPendingDynamicRemoval() bool {
	if m != nil {
		return m.PendingDynamicRemoval
	}
	return false
}

func (m *HostHealthStatus) GetPendingActiveHc() bool {
	if m != nil {
		return m.PendingActiveHc
	}
	return false
}

func (m *HostHealthStatus) GetEdsHealthStatus() core.HealthStatus {
	if m != nil {
		return m.EdsHealthStatus
	}
	return core.HealthStatus_UNKNOWN
}

func init() {
	proto.RegisterType((*Clusters)(nil), "envoy.admin.v2alpha.Clusters")
	proto.RegisterType((*ClusterStatus)(nil), "envoy.admin.v2alpha.ClusterStatus")
	proto.RegisterType((*HostStatus)(nil), "envoy.admin.v2alpha.HostStatus")
	proto.RegisterType((*HostHealthStatus)(nil), "envoy.admin.v2alpha.HostHealthStatus")
}

func init() { proto.RegisterFile("envoy/admin/v2alpha/clusters.proto", fileDescriptor_c6251a3a957f478b) }

var fileDescriptor_c6251a3a957f478b = []byte{
	// 692 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdf, 0x6f, 0xd3, 0x3a,
	0x18, 0x55, 0xbb, 0x6e, 0xcb, 0xdc, 0xf5, 0x6e, 0xf3, 0xee, 0xdd, 0xb2, 0xde, 0x49, 0xeb, 0x2a,
	0x10, 0x15, 0x42, 0x29, 0x2a, 0xd3, 0xf6, 0xc0, 0x03, 0xda, 0x0f, 0xa4, 0x09, 0x18, 0x9b, 0x32,
	0x84, 0x04, 0x2f, 0x96, 0xe7, 0x7c, 0x34, 0x86, 0x34, 0x8e, 0x6c, 0xb7, 0xd0, 0xff, 0x94, 0x17,
	0xfe, 0x0a, 0xfe, 0x01, 0x14, 0xdb, 0xe9, 0x32, 0x96, 0xc2, 0x5b, 0x9c, 0x73, 0x8e, 0xfd, 0x9d,
	0xef, 0x3b, 0x36, 0xea, 0x42, 0x3a, 0x11, 0xd3, 0x3e, 0x8d, 0x46, 0x3c, 0xed, 0x4f, 0x06, 0x34,
	0xc9, 0x62, 0xda, 0x67, 0xc9, 0x58, 0x69, 0x90, 0x2a, 0xc8, 0xa4, 0xd0, 0x02, 0x6f, 0x1a, 0x4e,
	0x60, 0x38, 0x81, 0xe3, 0xb4, 0xf7, 0xab, 0x84, 0x23, 0xd0, 0x92, 0x33, 0xa7, 0x6b, 0xef, 0x39,
	0x4a, 0xc6, 0xfb, 0x93, 0x41, 0x9f, 0x09, 0x09, 0x7d, 0x1a, 0x45, 0x12, 0x54, 0x41, 0xd8, 0xbd,
	0x4f, 0xb8, 0xa1, 0x0a, 0x1c, 0xfa, 0xe0, 0x3e, 0x1a, 0x03, 0x4d, 0x74, 0x4c, 0x58, 0x0c, 0xec,
	0x8b, 0x63, 0xf9, 0x96, 0xa5, 0xa7, 0x19, 0xf4, 0x33, 0x90, 0x0c, 0x52, 0x6d, 0x91, 0xee, 0x07,
	0xe4, 0x9d, 0x3a, 0x23, 0xf8, 0x02, 0xad, 0x3b, 0x53, 0x44, 0x69, 0xaa, 0xc7, 0x0a, 0x94, 0x5f,
	0xeb, 0x2c, 0xf4, 0x9a, 0x83, 0x6e, 0x50, 0xe1, 0x2e, 0x70, 0xc2, 0x6b, 0xc3, 0x0d, 0xd7, 0x58,
	0x79, 0x09, 0xaa, 0xfb, 0xa3, 0x8e, 0x5a, 0x77, 0x28, 0x18, 0xa3, 0x46, 0x4a, 0x47, 0xe0, 0xd7,
	0x3a, 0xb5, 0xde, 0x4a, 0x68, 0xbe, 0x71, 0x17, 0xb5, 0x68, 0x14, 0x41, 0x44, 0x26, 0x9c, 0x12,
	0x9a, 0x71, 0xbf, 0xde, 0xa9, 0xf5, 0xbc, 0xb0, 0x69, 0x7e, 0xbe, 0xe7, 0xf4, 0x38, 0xe3, 0xf8,
	0x23, 0xda, 0x53, 0x63, 0xc6, 0x40, 0x29, 0x22, 0xa9, 0x06, 0x02, 0x9f, 0x81, 0x69, 0x2e, 0x52,
	0xa2, 0x63, 0x09, 0x2a, 0x16, 0x49, 0xe4, 0x2f, 0x74, 0x6a, 0xbd, 0xe6, 0x60, 0xd3, 0xd5, 0x99,
	0x1b, 0x0d, 0xae, 0xac, 0xd1, 0x70, 0xd7, 0x69, 0x43, 0xaa, 0xe1, 0xa5, 0x53, 0xbe, 0x2b, 0x84,
	0xf8, 0x0c, 0xb5, 0x62, 0xa1, 0xf4, 0xad, 0xe3, 0x86, 0x71, 0xbc, 0x57, 0xe9, 0xf8, 0x5c, 0x28,
	0xed, 0xec, 0xae, 0xc6, 0xb3, 0x6f, 0x50, 0x58, 0xa2, 0x27, 0x89, 0x60, 0x34, 0x21, 0x42, 0xf2,
	0x21, 0x4f, 0xc9, 0xdf, 0xca, 0x5d, 0x9c, 0x5f, 0xee, 0x23, 0xb3, 0xd1, 0xa5, 0xd9, 0xe7, 0xfa,
	0x0f, 0x95, 0x77, 0xbf, 0x2f, 0x20, 0x74, 0x5b, 0x10, 0x3e, 0x40, 0xcb, 0x2e, 0x38, 0xa6, 0xbf,
	0xcd, 0x41, 0xbb, 0xb0, 0x90, 0xf1, 0x60, 0x32, 0x08, 0xf2, 0x6c, 0x04, 0xc7, 0x96, 0x11, 0x16,
	0x54, 0x7c, 0x84, 0x16, 0x73, 0xe7, 0xca, 0xaf, 0x1b, 0xdb, 0xfb, 0x95, 0xb6, 0xaf, 0xf9, 0x28,
	0x4b, 0xe0, 0xc2, 0xe4, 0x36, 0xb4, 0x7c, 0xfc, 0x0a, 0xb5, 0x5c, 0xd0, 0x6c, 0xe7, 0xdc, 0x04,
	0x1e, 0xce, 0xed, 0xdb, 0xb9, 0x61, 0xcf, 0xba, 0x57, 0x5a, 0xe1, 0x43, 0xb4, 0x5a, 0x6e, 0x98,
	0xdf, 0x98, 0xdf, 0x9d, 0x66, 0x69, 0x98, 0x78, 0x0b, 0x2d, 0x7d, 0x05, 0x3e, 0x8c, 0xb5, 0xe9,
	0x67, 0x2b, 0x74, 0x2b, 0xdc, 0x46, 0x5e, 0x3e, 0x1d, 0x93, 0xb5, 0x25, 0x93, 0xb5, 0xd9, 0x3a,
	0xc7, 0x32, 0xc9, 0x85, 0xe4, 0x7a, 0xea, 0x2f, 0x1b, 0xd5, 0x6c, 0x8d, 0xdf, 0xa2, 0x9d, 0xb9,
	0x53, 0xf4, 0xbd, 0xf9, 0x45, 0x6d, 0x55, 0x8f, 0x0c, 0x1f, 0x21, 0xcf, 0x20, 0xf9, 0x59, 0x2b,
	0x46, 0xfe, 0x7f, 0xc5, 0x4c, 0xde, 0x38, 0x4a, 0x38, 0x23, 0x77, 0x7f, 0xd6, 0xd1, 0xfa, 0xef,
	0x3d, 0xc3, 0xcf, 0x51, 0xfb, 0x13, 0xe5, 0x09, 0x44, 0x84, 0x32, 0xcd, 0x27, 0x40, 0xca, 0x17,
	0xdd, 0xcc, 0xdc, 0x0b, 0xb7, 0x2d, 0xe3, 0xd8, 0x10, 0xac, 0xfa, 0x34, 0x87, 0xf1, 0x53, 0xf4,
	0xaf, 0x13, 0x8b, 0xb1, 0x4e, 0x38, 0x48, 0x27, 0xb3, 0xb7, 0x0d, 0x5b, 0xec, 0xd2, 0x42, 0x56,
	0xf1, 0x02, 0xed, 0xde, 0x3d, 0x2e, 0x82, 0xa1, 0xa4, 0xf9, 0x4d, 0xb5, 0xca, 0x86, 0x51, 0xee,
	0x94, 0x0f, 0x3c, 0x73, 0x0c, 0xbb, 0xc1, 0x21, 0xda, 0xce, 0x20, 0x8d, 0x78, 0x3a, 0x24, 0xd1,
	0x34, 0xa5, 0x23, 0xce, 0x88, 0x84, 0x91, 0x98, 0xd0, 0xc4, 0x8c, 0xcb, 0x0b, 0xff, 0x73, 0xf0,
	0x99, 0x45, 0x43, 0x0b, 0xe2, 0xc7, 0x68, 0xa3, 0xd0, 0x15, 0x46, 0x99, 0x19, 0xa3, 0x17, 0xae,
	0x39, 0xc0, 0xf9, 0x63, 0xf8, 0x35, 0xda, 0x80, 0x48, 0x91, 0xfb, 0x49, 0xfc, 0xe7, 0xf6, 0x06,
	0x97, 0x5a, 0x7d, 0x27, 0x83, 0x6b, 0x10, 0xa9, 0xf2, 0x8f, 0x93, 0x03, 0xb4, 0xcf, 0x85, 0x55,
	0x65, 0x52, 0x7c, 0x9b, 0x56, 0x45, 0xf9, 0xa4, 0x78, 0xd2, 0xd4, 0x55, 0xfe, 0x7e, 0x5e, 0xd5,
	0x6e, 0x96, 0xcc, 0x43, 0xfa, 0xec, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x18, 0x85, 0xa0, 0xdd,
	0x25, 0x06, 0x00, 0x00,
}
