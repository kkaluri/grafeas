// Copyright 2018 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: attestation.proto

package attestation_go_proto

import (
	proto "github.com/golang/protobuf/proto"
	common_go_proto "github.com/grafeas/grafeas/proto/v1beta1/common_go_proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Type (for example schema) of the attestation payload that was signed.
type PgpSignedAttestation_ContentType int32

const (
	// `ContentType` is not set.
	PgpSignedAttestation_CONTENT_TYPE_UNSPECIFIED PgpSignedAttestation_ContentType = 0
	// Atomic format attestation signature. See
	// https://github.com/containers/image/blob/8a5d2f82a6e3263290c8e0276c3e0f64e77723e7/docs/atomic-signature.md
	// The payload extracted from `signature` is a JSON blob conforming to the
	// linked schema.
	PgpSignedAttestation_SIMPLE_SIGNING_JSON PgpSignedAttestation_ContentType = 1
)

// Enum value maps for PgpSignedAttestation_ContentType.
var (
	PgpSignedAttestation_ContentType_name = map[int32]string{
		0: "CONTENT_TYPE_UNSPECIFIED",
		1: "SIMPLE_SIGNING_JSON",
	}
	PgpSignedAttestation_ContentType_value = map[string]int32{
		"CONTENT_TYPE_UNSPECIFIED": 0,
		"SIMPLE_SIGNING_JSON":      1,
	}
)

func (x PgpSignedAttestation_ContentType) Enum() *PgpSignedAttestation_ContentType {
	p := new(PgpSignedAttestation_ContentType)
	*p = x
	return p
}

func (x PgpSignedAttestation_ContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PgpSignedAttestation_ContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_attestation_proto_enumTypes[0].Descriptor()
}

func (PgpSignedAttestation_ContentType) Type() protoreflect.EnumType {
	return &file_attestation_proto_enumTypes[0]
}

func (x PgpSignedAttestation_ContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PgpSignedAttestation_ContentType.Descriptor instead.
func (PgpSignedAttestation_ContentType) EnumDescriptor() ([]byte, []int) {
	return file_attestation_proto_rawDescGZIP(), []int{0, 0}
}

// Type of the attestation plaintext that was signed.
type GenericSignedAttestation_ContentType int32

const (
	// `ContentType` is not set.
	GenericSignedAttestation_CONTENT_TYPE_UNSPECIFIED GenericSignedAttestation_ContentType = 0
	// Atomic format attestation signature. See
	// https://github.com/containers/image/blob/8a5d2f82a6e3263290c8e0276c3e0f64e77723e7/docs/atomic-signature.md
	// The payload extracted in `plaintext` is a JSON blob conforming to the
	// linked schema.
	GenericSignedAttestation_SIMPLE_SIGNING_JSON GenericSignedAttestation_ContentType = 1
)

// Enum value maps for GenericSignedAttestation_ContentType.
var (
	GenericSignedAttestation_ContentType_name = map[int32]string{
		0: "CONTENT_TYPE_UNSPECIFIED",
		1: "SIMPLE_SIGNING_JSON",
	}
	GenericSignedAttestation_ContentType_value = map[string]int32{
		"CONTENT_TYPE_UNSPECIFIED": 0,
		"SIMPLE_SIGNING_JSON":      1,
	}
)

func (x GenericSignedAttestation_ContentType) Enum() *GenericSignedAttestation_ContentType {
	p := new(GenericSignedAttestation_ContentType)
	*p = x
	return p
}

func (x GenericSignedAttestation_ContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GenericSignedAttestation_ContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_attestation_proto_enumTypes[1].Descriptor()
}

func (GenericSignedAttestation_ContentType) Type() protoreflect.EnumType {
	return &file_attestation_proto_enumTypes[1]
}

func (x GenericSignedAttestation_ContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GenericSignedAttestation_ContentType.Descriptor instead.
func (GenericSignedAttestation_ContentType) EnumDescriptor() ([]byte, []int) {
	return file_attestation_proto_rawDescGZIP(), []int{1, 0}
}

// An attestation wrapper with a PGP-compatible signature. This message only
// supports `ATTACHED` signatures, where the payload that is signed is included
// alongside the signature itself in the same file.
type PgpSignedAttestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The raw content of the signature, as output by GNU Privacy Guard
	// (GPG) or equivalent. Since this message only supports attached signatures,
	// the payload that was signed must be attached. While the signature format
	// supported is dependent on the verification implementation, currently only
	// ASCII-armored (`--armor` to gpg), non-clearsigned (`--sign` rather than
	// `--clearsign` to gpg) are supported. Concretely, `gpg --sign --armor
	// --output=signature.gpg payload.json` will create the signature content
	// expected in this field in `signature.gpg` for the `payload.json`
	// attestation payload.
	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Type (for example schema) of the attestation payload that was signed.
	// The verifier must ensure that the provided type is one that the verifier
	// supports, and that the attestation payload is a valid instantiation of that
	// type (for example by validating a JSON schema).
	ContentType PgpSignedAttestation_ContentType `protobuf:"varint,3,opt,name=content_type,json=contentType,proto3,enum=grafeas.v1beta1.attestation.PgpSignedAttestation_ContentType" json:"content_type,omitempty"`
	// This field is used by verifiers to select the public key used to validate
	// the signature. Note that the policy of the verifier ultimately determines
	// which public keys verify a signature based on the context of the
	// verification. There is no guarantee validation will succeed if the
	// verifier has no key matching this ID, even if it has a key under a
	// different ID that would verify the signature. Note that this ID should also
	// be present in the signature content above, but that is not expected to be
	// used by the verifier.
	//
	// Types that are assignable to KeyId:
	//	*PgpSignedAttestation_PgpKeyId
	KeyId isPgpSignedAttestation_KeyId `protobuf_oneof:"key_id"`
}

func (x *PgpSignedAttestation) Reset() {
	*x = PgpSignedAttestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attestation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PgpSignedAttestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PgpSignedAttestation) ProtoMessage() {}

func (x *PgpSignedAttestation) ProtoReflect() protoreflect.Message {
	mi := &file_attestation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PgpSignedAttestation.ProtoReflect.Descriptor instead.
func (*PgpSignedAttestation) Descriptor() ([]byte, []int) {
	return file_attestation_proto_rawDescGZIP(), []int{0}
}

func (x *PgpSignedAttestation) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *PgpSignedAttestation) GetContentType() PgpSignedAttestation_ContentType {
	if x != nil {
		return x.ContentType
	}
	return PgpSignedAttestation_CONTENT_TYPE_UNSPECIFIED
}

func (m *PgpSignedAttestation) GetKeyId() isPgpSignedAttestation_KeyId {
	if m != nil {
		return m.KeyId
	}
	return nil
}

func (x *PgpSignedAttestation) GetPgpKeyId() string {
	if x, ok := x.GetKeyId().(*PgpSignedAttestation_PgpKeyId); ok {
		return x.PgpKeyId
	}
	return ""
}

type isPgpSignedAttestation_KeyId interface {
	isPgpSignedAttestation_KeyId()
}

type PgpSignedAttestation_PgpKeyId struct {
	// The cryptographic fingerprint of the key used to generate the signature,
	// as output by, e.g. `gpg --list-keys`. This should be the version 4, full
	// 160-bit fingerprint, expressed as a 40 character hexidecimal string. See
	// https://tools.ietf.org/html/rfc4880#section-12.2 for details.
	// Implementations may choose to acknowledge "LONG", "SHORT", or other
	// abbreviated key IDs, but only the full fingerprint is guaranteed to work.
	// In gpg, the full fingerprint can be retrieved from the `fpr` field
	// returned when calling --list-keys with --with-colons.  For example:
	// ```
	// gpg --with-colons --with-fingerprint --force-v4-certs \
	//     --list-keys attester@example.com
	// tru::1:1513631572:0:3:1:5
	// pub:...<SNIP>...
	// fpr:::::::::24FF6481B76AC91E66A00AC657A93A81EF3AE6FB:
	// ```
	// Above, the fingerprint is `24FF6481B76AC91E66A00AC657A93A81EF3AE6FB`.
	PgpKeyId string `protobuf:"bytes,2,opt,name=pgp_key_id,json=pgpKeyId,proto3,oneof"`
}

func (*PgpSignedAttestation_PgpKeyId) isPgpSignedAttestation_KeyId() {}

// An attestation wrapper that uses the Grafeas `Signature` message.
// This attestation must define the `serialized_payload` that the `signatures` verify
// and any metadata necessary to interpret that plaintext.  The signatures
// should always be over the `serialized_payload` bytestring.
type GenericSignedAttestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type (for example schema) of the attestation payload that was signed.
	// The verifier must ensure that the provided type is one that the verifier
	// supports, and that the attestation payload is a valid instantiation of that
	// type (for example by validating a JSON schema).
	ContentType GenericSignedAttestation_ContentType `protobuf:"varint,1,opt,name=content_type,json=contentType,proto3,enum=grafeas.v1beta1.attestation.GenericSignedAttestation_ContentType" json:"content_type,omitempty"`
	// The serialized payload that is verified by one or more `signatures`.
	// The encoding and semantic meaning of this payload must match what is set in
	// `content_type`.
	SerializedPayload []byte `protobuf:"bytes,2,opt,name=serialized_payload,json=serializedPayload,proto3" json:"serialized_payload,omitempty"`
	// One or more signatures over `serialized_payload`.  Verifier implementations
	// should consider this attestation message verified if at least one
	// `signature` verifies `serialized_payload`.  See `Signature` in common.proto
	// for more details on signature structure and verification.
	Signatures []*common_go_proto.Signature `protobuf:"bytes,3,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *GenericSignedAttestation) Reset() {
	*x = GenericSignedAttestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attestation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericSignedAttestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericSignedAttestation) ProtoMessage() {}

func (x *GenericSignedAttestation) ProtoReflect() protoreflect.Message {
	mi := &file_attestation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericSignedAttestation.ProtoReflect.Descriptor instead.
func (*GenericSignedAttestation) Descriptor() ([]byte, []int) {
	return file_attestation_proto_rawDescGZIP(), []int{1}
}

func (x *GenericSignedAttestation) GetContentType() GenericSignedAttestation_ContentType {
	if x != nil {
		return x.ContentType
	}
	return GenericSignedAttestation_CONTENT_TYPE_UNSPECIFIED
}

func (x *GenericSignedAttestation) GetSerializedPayload() []byte {
	if x != nil {
		return x.SerializedPayload
	}
	return nil
}

func (x *GenericSignedAttestation) GetSignatures() []*common_go_proto.Signature {
	if x != nil {
		return x.Signatures
	}
	return nil
}

// Note kind that represents a logical attestation "role" or "authority". For
// example, an organization might have one `Authority` for "QA" and one for
// "build". This note is intended to act strictly as a grouping mechanism for
// the attached occurrences (Attestations). This grouping mechanism also
// provides a security boundary, since IAM ACLs gate the ability for a principle
// to attach an occurrence to a given note. It also provides a single point of
// lookup to find all attached attestation occurrences, even if they don't all
// live in the same project.
type Authority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hint hints at the purpose of the attestation authority.
	Hint *Authority_Hint `protobuf:"bytes,1,opt,name=hint,proto3" json:"hint,omitempty"`
}

func (x *Authority) Reset() {
	*x = Authority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attestation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authority) ProtoMessage() {}

func (x *Authority) ProtoReflect() protoreflect.Message {
	mi := &file_attestation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authority.ProtoReflect.Descriptor instead.
func (*Authority) Descriptor() ([]byte, []int) {
	return file_attestation_proto_rawDescGZIP(), []int{2}
}

func (x *Authority) GetHint() *Authority_Hint {
	if x != nil {
		return x.Hint
	}
	return nil
}

// Details of an attestation occurrence.
type Details struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Attestation for the resource.
	Attestation *Attestation `protobuf:"bytes,1,opt,name=attestation,proto3" json:"attestation,omitempty"`
}

func (x *Details) Reset() {
	*x = Details{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attestation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Details) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Details) ProtoMessage() {}

func (x *Details) ProtoReflect() protoreflect.Message {
	mi := &file_attestation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Details.ProtoReflect.Descriptor instead.
func (*Details) Descriptor() ([]byte, []int) {
	return file_attestation_proto_rawDescGZIP(), []int{3}
}

func (x *Details) GetAttestation() *Attestation {
	if x != nil {
		return x.Attestation
	}
	return nil
}

// Occurrence that represents a single "attestation". The authenticity of an
// attestation can be verified using the attached signature. If the verifier
// trusts the public key of the signer, then verifying the signature is
// sufficient to establish trust. In this circumstance, the authority to which
// this attestation is attached is primarily useful for look-up (how to find
// this attestation if you already know the authority and artifact to be
// verified) and intent (which authority was this attestation intended to sign
// for).
type Attestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The signature, generally over the `resource_url`, that verifies
	// this attestation. The semantics of the signature veracity are ultimately
	// determined by the verification engine.
	//
	// Types that are assignable to Signature:
	//	*Attestation_PgpSignedAttestation
	//	*Attestation_GenericSignedAttestation
	Signature isAttestation_Signature `protobuf_oneof:"signature"`
}

func (x *Attestation) Reset() {
	*x = Attestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attestation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attestation) ProtoMessage() {}

func (x *Attestation) ProtoReflect() protoreflect.Message {
	mi := &file_attestation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attestation.ProtoReflect.Descriptor instead.
func (*Attestation) Descriptor() ([]byte, []int) {
	return file_attestation_proto_rawDescGZIP(), []int{4}
}

func (m *Attestation) GetSignature() isAttestation_Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (x *Attestation) GetPgpSignedAttestation() *PgpSignedAttestation {
	if x, ok := x.GetSignature().(*Attestation_PgpSignedAttestation); ok {
		return x.PgpSignedAttestation
	}
	return nil
}

func (x *Attestation) GetGenericSignedAttestation() *GenericSignedAttestation {
	if x, ok := x.GetSignature().(*Attestation_GenericSignedAttestation); ok {
		return x.GenericSignedAttestation
	}
	return nil
}

type isAttestation_Signature interface {
	isAttestation_Signature()
}

type Attestation_PgpSignedAttestation struct {
	// A PGP signed attestation.
	PgpSignedAttestation *PgpSignedAttestation `protobuf:"bytes,1,opt,name=pgp_signed_attestation,json=pgpSignedAttestation,proto3,oneof"`
}

type Attestation_GenericSignedAttestation struct {
	// An attestation that supports multiple `Signature`s
	// over the same attestation payload. The signatures
	// (defined in common.proto) support a superset of
	// public key types and IDs compared to PgpSignedAttestation.
	GenericSignedAttestation *GenericSignedAttestation `protobuf:"bytes,2,opt,name=generic_signed_attestation,json=genericSignedAttestation,proto3,oneof"`
}

func (*Attestation_PgpSignedAttestation) isAttestation_Signature() {}

func (*Attestation_GenericSignedAttestation) isAttestation_Signature() {}

// This submessage provides human-readable hints about the purpose of the
// authority. Because the name of a note acts as its resource reference, it is
// important to disambiguate the canonical name of the Note (which might be a
// UUID for security purposes) from "readable" names more suitable for debug
// output. Note that these hints should not be used to look up authorities in
// security sensitive contexts, such as when looking up attestations to
// verify.
type Authority_Hint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The human readable name of this attestation authority, for
	// example "qa".
	HumanReadableName string `protobuf:"bytes,1,opt,name=human_readable_name,json=humanReadableName,proto3" json:"human_readable_name,omitempty"`
}

func (x *Authority_Hint) Reset() {
	*x = Authority_Hint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attestation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authority_Hint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authority_Hint) ProtoMessage() {}

func (x *Authority_Hint) ProtoReflect() protoreflect.Message {
	mi := &file_attestation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authority_Hint.ProtoReflect.Descriptor instead.
func (*Authority_Hint) Descriptor() ([]byte, []int) {
	return file_attestation_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Authority_Hint) GetHumanReadableName() string {
	if x != nil {
		return x.HumanReadableName
	}
	return ""
}

var File_attestation_proto protoreflect.FileDescriptor

var file_attestation_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x02, 0x0a,
	0x14, 0x50, 0x67, 0x70, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x60, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d, 0x2e, 0x67, 0x72, 0x61, 0x66,
	0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x61, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x67, 0x70, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x67, 0x70, 0x5f, 0x6b, 0x65, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x67, 0x70,
	0x4b, 0x65, 0x79, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x53, 0x49, 0x47,
	0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x6b,
	0x65, 0x79, 0x5f, 0x69, 0x64, 0x22, 0xb1, 0x02, 0x0a, 0x18, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x64, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65,
	0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72,
	0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x22, 0x44, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x17, 0x0a, 0x13, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x49,
	0x4e, 0x47, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x01, 0x22, 0x84, 0x01, 0x0a, 0x09, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x48, 0x69,
	0x6e, 0x74, 0x52, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x1a, 0x36, 0x0a, 0x04, 0x48, 0x69, 0x6e, 0x74,
	0x12, 0x2e, 0x0a, 0x13, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x68,
	0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x55, 0x0a, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x4a, 0x0a, 0x0b, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xfc, 0x01, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x69, 0x0a, 0x16, 0x70, 0x67, 0x70, 0x5f, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x67, 0x70, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x14, 0x70, 0x67,
	0x70, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x75, 0x0a, 0x1a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x53, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52,
	0x18, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x67, 0x0a, 0x1e, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x61,
	0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x61, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x67,
	0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x47, 0x52, 0x41, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attestation_proto_rawDescOnce sync.Once
	file_attestation_proto_rawDescData = file_attestation_proto_rawDesc
)

func file_attestation_proto_rawDescGZIP() []byte {
	file_attestation_proto_rawDescOnce.Do(func() {
		file_attestation_proto_rawDescData = protoimpl.X.CompressGZIP(file_attestation_proto_rawDescData)
	})
	return file_attestation_proto_rawDescData
}

var file_attestation_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_attestation_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_attestation_proto_goTypes = []interface{}{
	(PgpSignedAttestation_ContentType)(0),     // 0: grafeas.v1beta1.attestation.PgpSignedAttestation.ContentType
	(GenericSignedAttestation_ContentType)(0), // 1: grafeas.v1beta1.attestation.GenericSignedAttestation.ContentType
	(*PgpSignedAttestation)(nil),              // 2: grafeas.v1beta1.attestation.PgpSignedAttestation
	(*GenericSignedAttestation)(nil),          // 3: grafeas.v1beta1.attestation.GenericSignedAttestation
	(*Authority)(nil),                         // 4: grafeas.v1beta1.attestation.Authority
	(*Details)(nil),                           // 5: grafeas.v1beta1.attestation.Details
	(*Attestation)(nil),                       // 6: grafeas.v1beta1.attestation.Attestation
	(*Authority_Hint)(nil),                    // 7: grafeas.v1beta1.attestation.Authority.Hint
	(*common_go_proto.Signature)(nil),         // 8: grafeas.v1beta1.Signature
}
var file_attestation_proto_depIdxs = []int32{
	0, // 0: grafeas.v1beta1.attestation.PgpSignedAttestation.content_type:type_name -> grafeas.v1beta1.attestation.PgpSignedAttestation.ContentType
	1, // 1: grafeas.v1beta1.attestation.GenericSignedAttestation.content_type:type_name -> grafeas.v1beta1.attestation.GenericSignedAttestation.ContentType
	8, // 2: grafeas.v1beta1.attestation.GenericSignedAttestation.signatures:type_name -> grafeas.v1beta1.Signature
	7, // 3: grafeas.v1beta1.attestation.Authority.hint:type_name -> grafeas.v1beta1.attestation.Authority.Hint
	6, // 4: grafeas.v1beta1.attestation.Details.attestation:type_name -> grafeas.v1beta1.attestation.Attestation
	2, // 5: grafeas.v1beta1.attestation.Attestation.pgp_signed_attestation:type_name -> grafeas.v1beta1.attestation.PgpSignedAttestation
	3, // 6: grafeas.v1beta1.attestation.Attestation.generic_signed_attestation:type_name -> grafeas.v1beta1.attestation.GenericSignedAttestation
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_attestation_proto_init() }
func file_attestation_proto_init() {
	if File_attestation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attestation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PgpSignedAttestation); i {
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
		file_attestation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericSignedAttestation); i {
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
		file_attestation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authority); i {
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
		file_attestation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Details); i {
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
		file_attestation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attestation); i {
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
		file_attestation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authority_Hint); i {
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
	file_attestation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PgpSignedAttestation_PgpKeyId)(nil),
	}
	file_attestation_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Attestation_PgpSignedAttestation)(nil),
		(*Attestation_GenericSignedAttestation)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_attestation_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_attestation_proto_goTypes,
		DependencyIndexes: file_attestation_proto_depIdxs,
		EnumInfos:         file_attestation_proto_enumTypes,
		MessageInfos:      file_attestation_proto_msgTypes,
	}.Build()
	File_attestation_proto = out.File
	file_attestation_proto_rawDesc = nil
	file_attestation_proto_goTypes = nil
	file_attestation_proto_depIdxs = nil
}
