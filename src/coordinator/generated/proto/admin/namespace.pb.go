// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3db/src/coordinator/generated/proto/admin/namespace.proto

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package admin

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import namespace "github.com/m3db/m3db/src/dbnode/generated/proto/namespace"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NamespaceGetResponse struct {
	Registry *namespace.Registry `protobuf:"bytes,1,opt,name=registry" json:"registry,omitempty"`
}

func (m *NamespaceGetResponse) Reset()                    { *m = NamespaceGetResponse{} }
func (m *NamespaceGetResponse) String() string            { return proto.CompactTextString(m) }
func (*NamespaceGetResponse) ProtoMessage()               {}
func (*NamespaceGetResponse) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{0} }

func (m *NamespaceGetResponse) GetRegistry() *namespace.Registry {
	if m != nil {
		return m.Registry
	}
	return nil
}

type NamespaceAddRequest struct {
	Name    string                      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Options *namespace.NamespaceOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
}

func (m *NamespaceAddRequest) Reset()                    { *m = NamespaceAddRequest{} }
func (m *NamespaceAddRequest) String() string            { return proto.CompactTextString(m) }
func (*NamespaceAddRequest) ProtoMessage()               {}
func (*NamespaceAddRequest) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{1} }

func (m *NamespaceAddRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamespaceAddRequest) GetOptions() *namespace.NamespaceOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func init() {
	proto.RegisterType((*NamespaceGetResponse)(nil), "admin.NamespaceGetResponse")
	proto.RegisterType((*NamespaceAddRequest)(nil), "admin.NamespaceAddRequest")
}
func (m *NamespaceGetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceGetResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Registry != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.Registry.Size()))
		n1, err := m.Registry.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *NamespaceAddRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceAddRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Options != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.Options.Size()))
		n2, err := m.Options.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeVarintNamespace(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *NamespaceGetResponse) Size() (n int) {
	var l int
	_ = l
	if m.Registry != nil {
		l = m.Registry.Size()
		n += 1 + l + sovNamespace(uint64(l))
	}
	return n
}

func (m *NamespaceAddRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	if m.Options != nil {
		l = m.Options.Size()
		n += 1 + l + sovNamespace(uint64(l))
	}
	return n
}

func sovNamespace(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNamespace(x uint64) (n int) {
	return sovNamespace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NamespaceGetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NamespaceGetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceGetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Registry == nil {
				m.Registry = &namespace.Registry{}
			}
			if err := m.Registry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
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
func (m *NamespaceAddRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NamespaceAddRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceAddRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Options == nil {
				m.Options = &namespace.NamespaceOptions{}
			}
			if err := m.Options.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
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
func skipNamespace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNamespace
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
					return 0, ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNamespace
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthNamespace
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNamespace
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipNamespace(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthNamespace = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNamespace   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3db/src/coordinator/generated/proto/admin/namespace.proto", fileDescriptorNamespace)
}

var fileDescriptorNamespace = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xcf, 0x3f, 0x4a, 0x04, 0x31,
	0x14, 0x06, 0x70, 0x23, 0xfe, 0x8d, 0x8d, 0x64, 0x2d, 0x16, 0x85, 0x41, 0xb6, 0xb2, 0x9a, 0x80,
	0x8b, 0x07, 0xd0, 0x66, 0xd1, 0x42, 0x21, 0x27, 0x30, 0x33, 0xef, 0x31, 0xa6, 0x48, 0x5e, 0x4c,
	0xde, 0x16, 0xde, 0xc2, 0x63, 0x59, 0x7a, 0x04, 0x19, 0x2f, 0x22, 0x66, 0xdd, 0x28, 0x8a, 0x4d,
	0x08, 0x5f, 0xbe, 0xef, 0x07, 0x91, 0x37, 0x83, 0xe3, 0x87, 0x65, 0xd7, 0xf6, 0xe4, 0xb5, 0x9f,
	0x43, 0xb7, 0x3a, 0x72, 0xea, 0x75, 0x4f, 0x94, 0xc0, 0x05, 0xcb, 0x94, 0xf4, 0x80, 0x01, 0x93,
	0x65, 0x04, 0x1d, 0x13, 0x31, 0x69, 0x0b, 0xde, 0x05, 0x1d, 0xac, 0xc7, 0x1c, 0x6d, 0x8f, 0x6d,
	0x49, 0xd5, 0x76, 0x89, 0x8f, 0xaf, 0xff, 0x25, 0xa1, 0x0b, 0x04, 0xf8, 0x47, 0xab, 0xce, 0x6f,
	0x71, 0xb6, 0x90, 0x47, 0xb7, 0xeb, 0x68, 0x81, 0x6c, 0x30, 0x47, 0x0a, 0x19, 0x95, 0x96, 0x7b,
	0x09, 0x07, 0x97, 0x39, 0x3d, 0x4d, 0xc5, 0xa9, 0x38, 0x3b, 0x38, 0x9f, 0xb4, 0xdf, 0x5b, 0xf3,
	0xf5, 0x64, 0x6a, 0x69, 0x76, 0x2f, 0x27, 0x15, 0xba, 0x04, 0x30, 0xf8, 0xb8, 0xc4, 0xcc, 0x4a,
	0xc9, 0xad, 0xcf, 0x59, 0x31, 0xf6, 0x4d, 0xb9, 0xab, 0x0b, 0xb9, 0x4b, 0x91, 0x1d, 0x85, 0x3c,
	0xdd, 0x2c, 0xf4, 0xc9, 0x0f, 0xba, 0x22, 0x77, 0xab, 0x8a, 0x59, 0x77, 0xaf, 0x0e, 0x5f, 0xc6,
	0x46, 0xbc, 0x8e, 0x8d, 0x78, 0x1b, 0x1b, 0xf1, 0xfc, 0xde, 0x6c, 0x74, 0x3b, 0xe5, 0x0f, 0xf3,
	0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x22, 0x2c, 0x00, 0x63, 0x01, 0x00, 0x00,
}