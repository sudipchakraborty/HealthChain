// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: healthchain/healthchain/record.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Record struct {
	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PatientId uint64 `protobuf:"varint,2,opt,name=patientId,proto3" json:"patientId,omitempty"`
	Sys       uint64 `protobuf:"varint,3,opt,name=sys,proto3" json:"sys,omitempty"`
	Dia       uint64 `protobuf:"varint,4,opt,name=dia,proto3" json:"dia,omitempty"`
	Creator   string `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff0346de41887112, []int{0}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return m.Size()
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Record) GetPatientId() uint64 {
	if m != nil {
		return m.PatientId
	}
	return 0
}

func (m *Record) GetSys() uint64 {
	if m != nil {
		return m.Sys
	}
	return 0
}

func (m *Record) GetDia() uint64 {
	if m != nil {
		return m.Dia
	}
	return 0
}

func (m *Record) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Record)(nil), "healthchain.healthchain.Record")
}

func init() {
	proto.RegisterFile("healthchain/healthchain/record.proto", fileDescriptor_ff0346de41887112)
}

var fileDescriptor_ff0346de41887112 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc9, 0x48, 0x4d, 0xcc,
	0x29, 0xc9, 0x48, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x47, 0x66, 0x17, 0xa5, 0x26, 0xe7, 0x17, 0xa5,
	0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x23, 0xc9, 0xe8, 0x21, 0xb1, 0x95, 0x4a, 0xb8,
	0xd8, 0x82, 0xc0, 0x0a, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58,
	0x82, 0x98, 0x32, 0x53, 0x84, 0x64, 0xb8, 0x38, 0x0b, 0x12, 0x4b, 0x32, 0x53, 0xf3, 0x4a, 0x3c,
	0x53, 0x24, 0x98, 0xc0, 0xc2, 0x08, 0x01, 0x21, 0x01, 0x2e, 0xe6, 0xe2, 0xca, 0x62, 0x09, 0x66,
	0xb0, 0x38, 0x88, 0x09, 0x12, 0x49, 0xc9, 0x4c, 0x94, 0x60, 0x81, 0x88, 0xa4, 0x64, 0x26, 0x0a,
	0x49, 0x70, 0xb1, 0x27, 0x17, 0xa5, 0x26, 0x96, 0xe4, 0x17, 0x49, 0xb0, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0xc1, 0xb8, 0x4e, 0x96, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25,
	0x8f, 0xec, 0x85, 0x0a, 0x14, 0x0f, 0x95, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x3d, 0x64,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xda, 0xfd, 0x1d, 0xf8, 0x00, 0x00, 0x00,
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Dia != 0 {
		i = encodeVarintRecord(dAtA, i, uint64(m.Dia))
		i--
		dAtA[i] = 0x20
	}
	if m.Sys != 0 {
		i = encodeVarintRecord(dAtA, i, uint64(m.Sys))
		i--
		dAtA[i] = 0x18
	}
	if m.PatientId != 0 {
		i = encodeVarintRecord(dAtA, i, uint64(m.PatientId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintRecord(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Record) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovRecord(uint64(m.Id))
	}
	if m.PatientId != 0 {
		n += 1 + sovRecord(uint64(m.PatientId))
	}
	if m.Sys != 0 {
		n += 1 + sovRecord(uint64(m.Sys))
	}
	if m.Dia != 0 {
		n += 1 + sovRecord(uint64(m.Dia))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	return n
}

func sovRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRecord(x uint64) (n int) {
	return sovRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecord
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
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PatientId", wireType)
			}
			m.PatientId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PatientId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sys", wireType)
			}
			m.Sys = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sys |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dia", wireType)
			}
			m.Dia = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Dia |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecord
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
func skipRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRecord
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
					return 0, ErrIntOverflowRecord
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
					return 0, ErrIntOverflowRecord
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
				return 0, ErrInvalidLengthRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRecord = fmt.Errorf("proto: unexpected end of group")
)
