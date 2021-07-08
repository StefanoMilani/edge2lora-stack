// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/cluster.proto

package ttnpb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// PeerInfo
type PeerInfo struct {
	// Port on which the gRPC server is exposed.
	GRPCPort uint32 `protobuf:"varint,1,opt,name=grpc_port,json=grpcPort,proto3" json:"grpc_port,omitempty"`
	// Indicates whether the gRPC server uses TLS.
	TLS bool `protobuf:"varint,2,opt,name=tls,proto3" json:"tls,omitempty"`
	// Roles of the peer.
	Roles []ClusterRole `protobuf:"varint,3,rep,packed,name=roles,proto3,enum=ttn.lorawan.v3.ClusterRole" json:"roles,omitempty"`
	// Tags of the peer
	Tags                 map[string]string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PeerInfo) Reset()      { *m = PeerInfo{} }
func (*PeerInfo) ProtoMessage() {}
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5716c3fcd711eefd, []int{0}
}
func (m *PeerInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeerInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerInfo.Merge(m, src)
}
func (m *PeerInfo) XXX_Size() int {
	return m.Size()
}
func (m *PeerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PeerInfo proto.InternalMessageInfo

func (m *PeerInfo) GetGRPCPort() uint32 {
	if m != nil {
		return m.GRPCPort
	}
	return 0
}

func (m *PeerInfo) GetTLS() bool {
	if m != nil {
		return m.TLS
	}
	return false
}

func (m *PeerInfo) GetRoles() []ClusterRole {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *PeerInfo) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*PeerInfo)(nil), "ttn.lorawan.v3.PeerInfo")
	golang_proto.RegisterType((*PeerInfo)(nil), "ttn.lorawan.v3.PeerInfo")
	proto.RegisterMapType((map[string]string)(nil), "ttn.lorawan.v3.PeerInfo.TagsEntry")
	golang_proto.RegisterMapType((map[string]string)(nil), "ttn.lorawan.v3.PeerInfo.TagsEntry")
}

func init() { proto.RegisterFile("lorawan-stack/api/cluster.proto", fileDescriptor_5716c3fcd711eefd) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/cluster.proto", fileDescriptor_5716c3fcd711eefd)
}

var fileDescriptor_5716c3fcd711eefd = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x31, 0x6f, 0xd4, 0x30,
	0x14, 0xc7, 0xe3, 0xa6, 0x85, 0xc4, 0x40, 0x85, 0x22, 0x86, 0x70, 0x08, 0x5f, 0xd4, 0x29, 0x0c,
	0xb5, 0x45, 0x4f, 0x02, 0xc4, 0x78, 0x15, 0x42, 0x48, 0x0c, 0x27, 0xd3, 0x89, 0x05, 0x39, 0x91,
	0xeb, 0x3b, 0x25, 0xb5, 0x23, 0xfb, 0xe5, 0xaa, 0x6e, 0x7c, 0x04, 0x3e, 0x02, 0x23, 0x1f, 0x83,
	0x91, 0x91, 0x91, 0xa9, 0xa2, 0xbe, 0x85, 0x91, 0x81, 0x81, 0x11, 0xc5, 0x01, 0xa4, 0x83, 0xed,
	0xff, 0x7f, 0xef, 0xf7, 0x9e, 0xfe, 0xef, 0xe1, 0x69, 0x6b, 0xac, 0x38, 0x17, 0xfa, 0xd0, 0x81,
	0xa8, 0x1b, 0x26, 0xba, 0x15, 0xab, 0xdb, 0xde, 0x81, 0xb4, 0xb4, 0xb3, 0x06, 0x4c, 0xb6, 0x0f,
	0xa0, 0xe9, 0x6f, 0x88, 0xae, 0x67, 0x93, 0x43, 0xb5, 0x82, 0x65, 0x5f, 0xd1, 0xda, 0x9c, 0x31,
	0x65, 0x94, 0x61, 0x01, 0xab, 0xfa, 0xd3, 0xe0, 0x82, 0x09, 0x6a, 0x1c, 0x9f, 0xdc, 0xff, 0x7f,
	0xbf, 0xd4, 0xfd, 0x99, 0x1b, 0xdb, 0x07, 0x3f, 0x10, 0x4e, 0x16, 0x52, 0xda, 0x17, 0xfa, 0xd4,
	0x64, 0x0f, 0x70, 0xaa, 0x6c, 0x57, 0xbf, 0xe9, 0x8c, 0x85, 0x1c, 0x15, 0xa8, 0xbc, 0x35, 0xbf,
	0xe9, 0x2f, 0xa7, 0xc9, 0x73, 0xbe, 0x38, 0x5e, 0x18, 0x0b, 0x3c, 0x19, 0xda, 0x83, 0xca, 0xee,
	0xe2, 0x18, 0x5a, 0x97, 0xef, 0x14, 0xa8, 0x4c, 0xe6, 0xd7, 0xfd, 0xe5, 0x34, 0x3e, 0x79, 0xf9,
	0x8a, 0x0f, 0xb5, 0xec, 0x21, 0xde, 0xb3, 0xa6, 0x95, 0x2e, 0x8f, 0x8b, 0xb8, 0xdc, 0x3f, 0xba,
	0x47, 0xb7, 0x0f, 0xa0, 0xc7, 0xe3, 0x79, 0xdc, 0xb4, 0x92, 0x8f, 0x64, 0xf6, 0x08, 0xef, 0x82,
	0x50, 0x2e, 0xdf, 0x2d, 0xe2, 0xf2, 0xc6, 0xd1, 0xc1, 0xbf, 0x13, 0x7f, 0x02, 0xd2, 0x13, 0xa1,
	0xdc, 0x33, 0x0d, 0xf6, 0x82, 0x07, 0x7e, 0xf2, 0x18, 0xa7, 0x7f, 0x4b, 0xd9, 0x6d, 0x1c, 0x37,
	0xf2, 0x22, 0xe4, 0x4e, 0xf9, 0x20, 0xb3, 0x3b, 0x78, 0x6f, 0x2d, 0xda, 0x5e, 0x86, 0x98, 0x29,
	0x1f, 0xcd, 0xd3, 0x9d, 0x27, 0x68, 0x2e, 0xbe, 0x5c, 0x91, 0xe8, 0xe7, 0x15, 0x41, 0x6f, 0x3d,
	0x41, 0x1f, 0x3c, 0x41, 0x9f, 0x3c, 0x41, 0x9f, 0x3d, 0x41, 0x5f, 0x3d, 0x41, 0xdf, 0x3c, 0x89,
	0xbe, 0x7b, 0x82, 0xde, 0x6d, 0x48, 0xf4, 0x7e, 0x43, 0xa2, 0x8f, 0x1b, 0x82, 0x5e, 0x33, 0x65,
	0x28, 0x2c, 0x25, 0x2c, 0x57, 0x5a, 0x39, 0xaa, 0x25, 0x9c, 0x1b, 0xdb, 0xb0, 0xed, 0xe7, 0xae,
	0x67, 0xac, 0x6b, 0x14, 0x03, 0xd0, 0x5d, 0x55, 0x5d, 0x0b, 0x0f, 0x9e, 0xfd, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0xc8, 0x54, 0xc2, 0x96, 0xe1, 0x01, 0x00, 0x00,
}

func (this *PeerInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PeerInfo)
	if !ok {
		that2, ok := that.(PeerInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.GRPCPort != that1.GRPCPort {
		return false
	}
	if this.TLS != that1.TLS {
		return false
	}
	if len(this.Roles) != len(that1.Roles) {
		return false
	}
	for i := range this.Roles {
		if this.Roles[i] != that1.Roles[i] {
			return false
		}
	}
	if len(this.Tags) != len(that1.Tags) {
		return false
	}
	for i := range this.Tags {
		if this.Tags[i] != that1.Tags[i] {
			return false
		}
	}
	return true
}
func (m *PeerInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeerInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeerInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tags) > 0 {
		for k := range m.Tags {
			v := m.Tags[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintCluster(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintCluster(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintCluster(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Roles) > 0 {
		dAtA2 := make([]byte, len(m.Roles)*10)
		var j1 int
		for _, num := range m.Roles {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintCluster(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x1a
	}
	if m.TLS {
		i--
		if m.TLS {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.GRPCPort != 0 {
		i = encodeVarintCluster(dAtA, i, uint64(m.GRPCPort))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCluster(dAtA []byte, offset int, v uint64) int {
	offset -= sovCluster(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedPeerInfo(r randyCluster, easy bool) *PeerInfo {
	this := &PeerInfo{}
	this.GRPCPort = uint32(r.Uint32())
	this.TLS = bool(bool(r.Intn(2) == 0))
	v1 := r.Intn(10)
	this.Roles = make([]ClusterRole, v1)
	for i := 0; i < v1; i++ {
		this.Roles[i] = ClusterRole([]int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}[r.Intn(14)])
	}
	if r.Intn(5) != 0 {
		v2 := r.Intn(10)
		this.Tags = make(map[string]string)
		for i := 0; i < v2; i++ {
			this.Tags[randStringCluster(r)] = randStringCluster(r)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyCluster interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneCluster(r randyCluster) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringCluster(r randyCluster) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneCluster(r)
	}
	return string(tmps)
}
func randUnrecognizedCluster(r randyCluster, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldCluster(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldCluster(dAtA []byte, r randyCluster, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateCluster(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateCluster(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateCluster(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateCluster(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateCluster(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateCluster(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateCluster(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *PeerInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GRPCPort != 0 {
		n += 1 + sovCluster(uint64(m.GRPCPort))
	}
	if m.TLS {
		n += 2
	}
	if len(m.Roles) > 0 {
		l = 0
		for _, e := range m.Roles {
			l += sovCluster(uint64(e))
		}
		n += 1 + sovCluster(uint64(l)) + l
	}
	if len(m.Tags) > 0 {
		for k, v := range m.Tags {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovCluster(uint64(len(k))) + 1 + len(v) + sovCluster(uint64(len(v)))
			n += mapEntrySize + 1 + sovCluster(uint64(mapEntrySize))
		}
	}
	return n
}

func sovCluster(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCluster(x uint64) (n int) {
	return sovCluster(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PeerInfo) String() string {
	if this == nil {
		return "nil"
	}
	keysForTags := make([]string, 0, len(this.Tags))
	for k := range this.Tags {
		keysForTags = append(keysForTags, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForTags)
	mapStringForTags := "map[string]string{"
	for _, k := range keysForTags {
		mapStringForTags += fmt.Sprintf("%v: %v,", k, this.Tags[k])
	}
	mapStringForTags += "}"
	s := strings.Join([]string{`&PeerInfo{`,
		`GRPCPort:` + fmt.Sprintf("%v", this.GRPCPort) + `,`,
		`TLS:` + fmt.Sprintf("%v", this.TLS) + `,`,
		`Roles:` + fmt.Sprintf("%v", this.Roles) + `,`,
		`Tags:` + mapStringForTags + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCluster(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PeerInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCluster
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PeerInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeerInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GRPCPort", wireType)
			}
			m.GRPCPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GRPCPort |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TLS", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TLS = bool(v != 0)
		case 3:
			if wireType == 0 {
				var v ClusterRole
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCluster
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= ClusterRole(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Roles = append(m.Roles, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCluster
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthCluster
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthCluster
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Roles) == 0 {
					m.Roles = make([]ClusterRole, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v ClusterRole
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCluster
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= ClusterRole(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Roles = append(m.Roles, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tags == nil {
				m.Tags = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCluster
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCluster
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthCluster
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthCluster
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCluster
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthCluster
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthCluster
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipCluster(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthCluster
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Tags[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCluster(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCluster
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCluster
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCluster
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCluster
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCluster        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCluster          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCluster = fmt.Errorf("proto: unexpected end of group")
)
