// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package loan

import (
	fmt "fmt"
	io "io"
	reflect "reflect"
	sync "sync"

	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

var (
	md_Loan            protoreflect.MessageDescriptor
	fd_Loan_id         protoreflect.FieldDescriptor
	fd_Loan_amount     protoreflect.FieldDescriptor
	fd_Loan_fee        protoreflect.FieldDescriptor
	fd_Loan_collateral protoreflect.FieldDescriptor
	fd_Loan_deadline   protoreflect.FieldDescriptor
	fd_Loan_state      protoreflect.FieldDescriptor
	fd_Loan_borrower   protoreflect.FieldDescriptor
	fd_Loan_lender     protoreflect.FieldDescriptor
)

func init() {
	file_loan_loan_loan_proto_init()
	md_Loan = File_loan_loan_loan_proto.Messages().ByName("Loan")
	fd_Loan_id = md_Loan.Fields().ByName("id")
	fd_Loan_amount = md_Loan.Fields().ByName("amount")
	fd_Loan_fee = md_Loan.Fields().ByName("fee")
	fd_Loan_collateral = md_Loan.Fields().ByName("collateral")
	fd_Loan_deadline = md_Loan.Fields().ByName("deadline")
	fd_Loan_state = md_Loan.Fields().ByName("state")
	fd_Loan_borrower = md_Loan.Fields().ByName("borrower")
	fd_Loan_lender = md_Loan.Fields().ByName("lender")
}

var _ protoreflect.Message = (*fastReflection_Loan)(nil)

type fastReflection_Loan Loan

func (x *Loan) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Loan)(x)
}

func (x *Loan) slowProtoReflect() protoreflect.Message {
	mi := &file_loan_loan_loan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Loan_messageType fastReflection_Loan_messageType
var _ protoreflect.MessageType = fastReflection_Loan_messageType{}

type fastReflection_Loan_messageType struct{}

func (x fastReflection_Loan_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Loan)(nil)
}
func (x fastReflection_Loan_messageType) New() protoreflect.Message {
	return new(fastReflection_Loan)
}
func (x fastReflection_Loan_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Loan
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Loan) Descriptor() protoreflect.MessageDescriptor {
	return md_Loan
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Loan) Type() protoreflect.MessageType {
	return _fastReflection_Loan_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Loan) New() protoreflect.Message {
	return new(fastReflection_Loan)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Loan) Interface() protoreflect.ProtoMessage {
	return (*Loan)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Loan) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_Loan_id, value) {
			return
		}
	}
	if x.Amount != "" {
		value := protoreflect.ValueOfString(x.Amount)
		if !f(fd_Loan_amount, value) {
			return
		}
	}
	if x.Fee != "" {
		value := protoreflect.ValueOfString(x.Fee)
		if !f(fd_Loan_fee, value) {
			return
		}
	}
	if x.Collateral != "" {
		value := protoreflect.ValueOfString(x.Collateral)
		if !f(fd_Loan_collateral, value) {
			return
		}
	}
	if x.Deadline != "" {
		value := protoreflect.ValueOfString(x.Deadline)
		if !f(fd_Loan_deadline, value) {
			return
		}
	}
	if x.State != "" {
		value := protoreflect.ValueOfString(x.State)
		if !f(fd_Loan_state, value) {
			return
		}
	}
	if x.Borrower != "" {
		value := protoreflect.ValueOfString(x.Borrower)
		if !f(fd_Loan_borrower, value) {
			return
		}
	}
	if x.Lender != "" {
		value := protoreflect.ValueOfString(x.Lender)
		if !f(fd_Loan_lender, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Loan) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "loan.loan.Loan.id":
		return x.Id != uint64(0)
	case "loan.loan.Loan.amount":
		return x.Amount != ""
	case "loan.loan.Loan.fee":
		return x.Fee != ""
	case "loan.loan.Loan.collateral":
		return x.Collateral != ""
	case "loan.loan.Loan.deadline":
		return x.Deadline != ""
	case "loan.loan.Loan.state":
		return x.State != ""
	case "loan.loan.Loan.borrower":
		return x.Borrower != ""
	case "loan.loan.Loan.lender":
		return x.Lender != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: loan.loan.Loan"))
		}
		panic(fmt.Errorf("message loan.loan.Loan does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Loan) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "loan.loan.Loan.id":
		x.Id = uint64(0)
	case "loan.loan.Loan.amount":
		x.Amount = ""
	case "loan.loan.Loan.fee":
		x.Fee = ""
	case "loan.loan.Loan.collateral":
		x.Collateral = ""
	case "loan.loan.Loan.deadline":
		x.Deadline = ""
	case "loan.loan.Loan.state":
		x.State = ""
	case "loan.loan.Loan.borrower":
		x.Borrower = ""
	case "loan.loan.Loan.lender":
		x.Lender = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: loan.loan.Loan"))
		}
		panic(fmt.Errorf("message loan.loan.Loan does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Loan) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "loan.loan.Loan.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "loan.loan.Loan.amount":
		value := x.Amount
		return protoreflect.ValueOfString(value)
	case "loan.loan.Loan.fee":
		value := x.Fee
		return protoreflect.ValueOfString(value)
	case "loan.loan.Loan.collateral":
		value := x.Collateral
		return protoreflect.ValueOfString(value)
	case "loan.loan.Loan.deadline":
		value := x.Deadline
		return protoreflect.ValueOfString(value)
	case "loan.loan.Loan.state":
		value := x.State
		return protoreflect.ValueOfString(value)
	case "loan.loan.Loan.borrower":
		value := x.Borrower
		return protoreflect.ValueOfString(value)
	case "loan.loan.Loan.lender":
		value := x.Lender
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: loan.loan.Loan"))
		}
		panic(fmt.Errorf("message loan.loan.Loan does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Loan) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "loan.loan.Loan.id":
		x.Id = value.Uint()
	case "loan.loan.Loan.amount":
		x.Amount = value.Interface().(string)
	case "loan.loan.Loan.fee":
		x.Fee = value.Interface().(string)
	case "loan.loan.Loan.collateral":
		x.Collateral = value.Interface().(string)
	case "loan.loan.Loan.deadline":
		x.Deadline = value.Interface().(string)
	case "loan.loan.Loan.state":
		x.State = value.Interface().(string)
	case "loan.loan.Loan.borrower":
		x.Borrower = value.Interface().(string)
	case "loan.loan.Loan.lender":
		x.Lender = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: loan.loan.Loan"))
		}
		panic(fmt.Errorf("message loan.loan.Loan does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Loan) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "loan.loan.Loan.id":
		panic(fmt.Errorf("field id of message loan.loan.Loan is not mutable"))
	case "loan.loan.Loan.amount":
		panic(fmt.Errorf("field amount of message loan.loan.Loan is not mutable"))
	case "loan.loan.Loan.fee":
		panic(fmt.Errorf("field fee of message loan.loan.Loan is not mutable"))
	case "loan.loan.Loan.collateral":
		panic(fmt.Errorf("field collateral of message loan.loan.Loan is not mutable"))
	case "loan.loan.Loan.deadline":
		panic(fmt.Errorf("field deadline of message loan.loan.Loan is not mutable"))
	case "loan.loan.Loan.state":
		panic(fmt.Errorf("field state of message loan.loan.Loan is not mutable"))
	case "loan.loan.Loan.borrower":
		panic(fmt.Errorf("field borrower of message loan.loan.Loan is not mutable"))
	case "loan.loan.Loan.lender":
		panic(fmt.Errorf("field lender of message loan.loan.Loan is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: loan.loan.Loan"))
		}
		panic(fmt.Errorf("message loan.loan.Loan does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Loan) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "loan.loan.Loan.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "loan.loan.Loan.amount":
		return protoreflect.ValueOfString("")
	case "loan.loan.Loan.fee":
		return protoreflect.ValueOfString("")
	case "loan.loan.Loan.collateral":
		return protoreflect.ValueOfString("")
	case "loan.loan.Loan.deadline":
		return protoreflect.ValueOfString("")
	case "loan.loan.Loan.state":
		return protoreflect.ValueOfString("")
	case "loan.loan.Loan.borrower":
		return protoreflect.ValueOfString("")
	case "loan.loan.Loan.lender":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: loan.loan.Loan"))
		}
		panic(fmt.Errorf("message loan.loan.Loan does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Loan) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in loan.loan.Loan", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Loan) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Loan) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Loan) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Loan) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Loan)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		l = len(x.Amount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Fee)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Collateral)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Deadline)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.State)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Borrower)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Lender)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Loan)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Lender) > 0 {
			i -= len(x.Lender)
			copy(dAtA[i:], x.Lender)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Lender)))
			i--
			dAtA[i] = 0x42
		}
		if len(x.Borrower) > 0 {
			i -= len(x.Borrower)
			copy(dAtA[i:], x.Borrower)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Borrower)))
			i--
			dAtA[i] = 0x3a
		}
		if len(x.State) > 0 {
			i -= len(x.State)
			copy(dAtA[i:], x.State)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.State)))
			i--
			dAtA[i] = 0x32
		}
		if len(x.Deadline) > 0 {
			i -= len(x.Deadline)
			copy(dAtA[i:], x.Deadline)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Deadline)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.Collateral) > 0 {
			i -= len(x.Collateral)
			copy(dAtA[i:], x.Collateral)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Collateral)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Fee) > 0 {
			i -= len(x.Fee)
			copy(dAtA[i:], x.Fee)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Fee)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Amount) > 0 {
			i -= len(x.Amount)
			copy(dAtA[i:], x.Amount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Amount)))
			i--
			dAtA[i] = 0x12
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Loan)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Loan: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Loan: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Amount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Fee = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Collateral = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Deadline = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.State = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Borrower", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Borrower = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Lender", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Lender = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: loan/loan/loan.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Loan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount     string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee        string `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
	Collateral string `protobuf:"bytes,4,opt,name=collateral,proto3" json:"collateral,omitempty"`
	Deadline   string `protobuf:"bytes,5,opt,name=deadline,proto3" json:"deadline,omitempty"`
	State      string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	Borrower   string `protobuf:"bytes,7,opt,name=borrower,proto3" json:"borrower,omitempty"`
	Lender     string `protobuf:"bytes,8,opt,name=lender,proto3" json:"lender,omitempty"`
}

func (x *Loan) Reset() {
	*x = Loan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loan_loan_loan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Loan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Loan) ProtoMessage() {}

// Deprecated: Use Loan.ProtoReflect.Descriptor instead.
func (*Loan) Descriptor() ([]byte, []int) {
	return file_loan_loan_loan_proto_rawDescGZIP(), []int{0}
}

func (x *Loan) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Loan) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Loan) GetFee() string {
	if x != nil {
		return x.Fee
	}
	return ""
}

func (x *Loan) GetCollateral() string {
	if x != nil {
		return x.Collateral
	}
	return ""
}

func (x *Loan) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *Loan) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Loan) GetBorrower() string {
	if x != nil {
		return x.Borrower
	}
	return ""
}

func (x *Loan) GetLender() string {
	if x != nil {
		return x.Lender
	}
	return ""
}

var File_loan_loan_loan_proto protoreflect.FileDescriptor

var file_loan_loan_loan_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x6f, 0x61, 0x6e, 0x2f, 0x6c, 0x6f, 0x61, 0x6e, 0x2f, 0x6c, 0x6f, 0x61, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x6c, 0x6f, 0x61,
	0x6e, 0x22, 0xc6, 0x01, 0x0a, 0x04, 0x4c, 0x6f, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x66, 0x65, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x74, 0x65, 0x72,
	0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x74,
	0x65, 0x72, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x7b, 0x0a, 0x0d, 0x63, 0x6f,
	0x6d, 0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x42, 0x09, 0x4c, 0x6f, 0x61,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x61, 0x6e, 0x2f,
	0x6c, 0x6f, 0x61, 0x6e, 0xa2, 0x02, 0x03, 0x4c, 0x4c, 0x58, 0xaa, 0x02, 0x09, 0x4c, 0x6f, 0x61,
	0x6e, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0xca, 0x02, 0x09, 0x4c, 0x6f, 0x61, 0x6e, 0x5c, 0x4c, 0x6f,
	0x61, 0x6e, 0xe2, 0x02, 0x15, 0x4c, 0x6f, 0x61, 0x6e, 0x5c, 0x4c, 0x6f, 0x61, 0x6e, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x4c, 0x6f, 0x61,
	0x6e, 0x3a, 0x3a, 0x4c, 0x6f, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loan_loan_loan_proto_rawDescOnce sync.Once
	file_loan_loan_loan_proto_rawDescData = file_loan_loan_loan_proto_rawDesc
)

func file_loan_loan_loan_proto_rawDescGZIP() []byte {
	file_loan_loan_loan_proto_rawDescOnce.Do(func() {
		file_loan_loan_loan_proto_rawDescData = protoimpl.X.CompressGZIP(file_loan_loan_loan_proto_rawDescData)
	})
	return file_loan_loan_loan_proto_rawDescData
}

var file_loan_loan_loan_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_loan_loan_loan_proto_goTypes = []interface{}{
	(*Loan)(nil), // 0: loan.loan.Loan
}
var file_loan_loan_loan_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_loan_loan_loan_proto_init() }
func file_loan_loan_loan_proto_init() {
	if File_loan_loan_loan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loan_loan_loan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Loan); i {
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
			RawDescriptor: file_loan_loan_loan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_loan_loan_loan_proto_goTypes,
		DependencyIndexes: file_loan_loan_loan_proto_depIdxs,
		MessageInfos:      file_loan_loan_loan_proto_msgTypes,
	}.Build()
	File_loan_loan_loan_proto = out.File
	file_loan_loan_loan_proto_rawDesc = nil
	file_loan_loan_loan_proto_goTypes = nil
	file_loan_loan_loan_proto_depIdxs = nil
}
