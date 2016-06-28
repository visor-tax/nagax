// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	Error
	ErrorResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Error struct {
	Code             *string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Title            *string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Detail           *string `protobuf:"bytes,3,opt,name=detail" json:"detail,omitempty"`
	Field            *string `protobuf:"bytes,4,opt,name=field" json:"field,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Error) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *Error) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *Error) GetDetail() string {
	if m != nil && m.Detail != nil {
		return *m.Detail
	}
	return ""
}

func (m *Error) GetField() string {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return ""
}

type ErrorResponse struct {
	Errors           []*Error `protobuf:"bytes,1,rep,name=errors" json:"errors,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ErrorResponse) Reset()                    { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string            { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()               {}
func (*ErrorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ErrorResponse) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "nagax.router.api.Error")
	proto.RegisterType((*ErrorResponse)(nil), "nagax.router.api.ErrorResponse")
}
func (m *Error) SetCode(v *string) {
	m.Code = v
}

func (m *Error) SetTitle(v *string) {
	m.Title = v
}

func (m *Error) SetDetail(v *string) {
	m.Detail = v
}

func (m *Error) SetField(v *string) {
	m.Field = v
}

func (m *ErrorResponse) SetErrors(v []*Error) {
	m.Errors = v
}

var fileDescriptor0 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc8, 0x4b, 0x4c, 0x4f, 0xac, 0xd0, 0x2b, 0xca, 0x2f,
	0x2d, 0x49, 0x2d, 0xd2, 0x03, 0x8a, 0x2b, 0xc5, 0x73, 0xb1, 0xba, 0x16, 0x15, 0xe5, 0x17, 0x09,
	0x09, 0x71, 0xb1, 0x24, 0xe7, 0xa7, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9,
	0x42, 0x22, 0x5c, 0xac, 0x25, 0x99, 0x25, 0x39, 0xa9, 0x12, 0x4c, 0x60, 0x41, 0x08, 0x47, 0x48,
	0x8c, 0x8b, 0x2d, 0x25, 0xb5, 0x24, 0x31, 0x33, 0x47, 0x82, 0x19, 0x2c, 0x0c, 0xe5, 0x81, 0x54,
	0xa7, 0x65, 0xa6, 0xe6, 0xa4, 0x48, 0xb0, 0x40, 0x54, 0x83, 0x39, 0x4a, 0x0e, 0x5c, 0xbc, 0x60,
	0x0b, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xf4, 0xb9, 0xd8, 0x52, 0x41, 0x02,
	0xc5, 0x40, 0xab, 0x98, 0x35, 0xb8, 0x8d, 0xc4, 0xf5, 0xd0, 0x1d, 0xa5, 0x07, 0xd1, 0x00, 0x55,
	0x06, 0x08, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xea, 0x5d, 0x91, 0xc0, 0x00, 0x00, 0x00,
}
