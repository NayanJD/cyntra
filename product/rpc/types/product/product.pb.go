// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: product.proto

package product

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price           uint64                 `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	Description     string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Size            string                 `protobuf:"bytes,4,opt,name=size,proto3" json:"size,omitempty"`
	Color           string                 `protobuf:"bytes,5,opt,name=color,proto3" json:"color,omitempty"`
	Brand           string                 `protobuf:"bytes,6,opt,name=brand,proto3" json:"brand,omitempty"`
	Shade           string                 `protobuf:"bytes,7,opt,name=shade,proto3" json:"shade,omitempty"`
	WashCare        string                 `protobuf:"bytes,8,opt,name=washCare,proto3" json:"washCare,omitempty"`
	Stretchable     string                 `protobuf:"bytes,9,opt,name=stretchable,proto3" json:"stretchable,omitempty"`
	Distress        string                 `protobuf:"bytes,10,opt,name=distress,proto3" json:"distress,omitempty"`
	Features        string                 `protobuf:"bytes,11,opt,name=features,proto3" json:"features,omitempty"`
	Fade            string                 `protobuf:"bytes,12,opt,name=fade,proto3" json:"fade,omitempty"`
	Fabric          string                 `protobuf:"bytes,13,opt,name=fabric,proto3" json:"fabric,omitempty"`
	Category        string                 `protobuf:"bytes,14,opt,name=category,proto3" json:"category,omitempty"`
	CountryOfOrigin string                 `protobuf:"bytes,15,opt,name=countryOfOrigin,proto3" json:"countryOfOrigin,omitempty"`
	Discount        uint32                 `protobuf:"varint,16,opt,name=discount,proto3" json:"discount,omitempty"`
	Quantity        uint64                 `protobuf:"varint,17,opt,name=quantity,proto3" json:"quantity,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,19,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	ArchivedAt      *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=archivedAt,proto3" json:"archivedAt,omitempty"`
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetPrice() uint64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductRequest) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *CreateProductRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *CreateProductRequest) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *CreateProductRequest) GetShade() string {
	if x != nil {
		return x.Shade
	}
	return ""
}

func (x *CreateProductRequest) GetWashCare() string {
	if x != nil {
		return x.WashCare
	}
	return ""
}

func (x *CreateProductRequest) GetStretchable() string {
	if x != nil {
		return x.Stretchable
	}
	return ""
}

func (x *CreateProductRequest) GetDistress() string {
	if x != nil {
		return x.Distress
	}
	return ""
}

func (x *CreateProductRequest) GetFeatures() string {
	if x != nil {
		return x.Features
	}
	return ""
}

func (x *CreateProductRequest) GetFade() string {
	if x != nil {
		return x.Fade
	}
	return ""
}

func (x *CreateProductRequest) GetFabric() string {
	if x != nil {
		return x.Fabric
	}
	return ""
}

func (x *CreateProductRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CreateProductRequest) GetCountryOfOrigin() string {
	if x != nil {
		return x.CountryOfOrigin
	}
	return ""
}

func (x *CreateProductRequest) GetDiscount() uint32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *CreateProductRequest) GetQuantity() uint64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CreateProductRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CreateProductRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *CreateProductRequest) GetArchivedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ArchivedAt
	}
	return nil
}

type ProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price           uint64                 `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Description     string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Size            string                 `protobuf:"bytes,5,opt,name=size,proto3" json:"size,omitempty"`
	Color           string                 `protobuf:"bytes,6,opt,name=color,proto3" json:"color,omitempty"`
	Brand           string                 `protobuf:"bytes,7,opt,name=brand,proto3" json:"brand,omitempty"`
	Shade           string                 `protobuf:"bytes,8,opt,name=shade,proto3" json:"shade,omitempty"`
	WashCare        string                 `protobuf:"bytes,9,opt,name=washCare,proto3" json:"washCare,omitempty"`
	Stretchable     string                 `protobuf:"bytes,10,opt,name=stretchable,proto3" json:"stretchable,omitempty"`
	Distress        string                 `protobuf:"bytes,11,opt,name=distress,proto3" json:"distress,omitempty"`
	Features        string                 `protobuf:"bytes,12,opt,name=features,proto3" json:"features,omitempty"`
	Fade            string                 `protobuf:"bytes,13,opt,name=fade,proto3" json:"fade,omitempty"`
	Fabric          string                 `protobuf:"bytes,14,opt,name=fabric,proto3" json:"fabric,omitempty"`
	Category        string                 `protobuf:"bytes,15,opt,name=category,proto3" json:"category,omitempty"`
	CountryOfOrigin string                 `protobuf:"bytes,16,opt,name=countryOfOrigin,proto3" json:"countryOfOrigin,omitempty"`
	Discount        uint32                 `protobuf:"varint,17,opt,name=discount,proto3" json:"discount,omitempty"`
	Quantity        uint64                 `protobuf:"varint,18,opt,name=quantity,proto3" json:"quantity,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,19,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	ArchivedAt      *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=archivedAt,proto3" json:"archivedAt,omitempty"`
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductResponse) GetPrice() uint64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductResponse) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *ProductResponse) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *ProductResponse) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *ProductResponse) GetShade() string {
	if x != nil {
		return x.Shade
	}
	return ""
}

func (x *ProductResponse) GetWashCare() string {
	if x != nil {
		return x.WashCare
	}
	return ""
}

func (x *ProductResponse) GetStretchable() string {
	if x != nil {
		return x.Stretchable
	}
	return ""
}

func (x *ProductResponse) GetDistress() string {
	if x != nil {
		return x.Distress
	}
	return ""
}

func (x *ProductResponse) GetFeatures() string {
	if x != nil {
		return x.Features
	}
	return ""
}

func (x *ProductResponse) GetFade() string {
	if x != nil {
		return x.Fade
	}
	return ""
}

func (x *ProductResponse) GetFabric() string {
	if x != nil {
		return x.Fabric
	}
	return ""
}

func (x *ProductResponse) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ProductResponse) GetCountryOfOrigin() string {
	if x != nil {
		return x.CountryOfOrigin
	}
	return ""
}

func (x *ProductResponse) GetDiscount() uint32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *ProductResponse) GetQuantity() uint64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ProductResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ProductResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ProductResponse) GetArchivedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ArchivedAt
	}
	return nil
}

type GetProductByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetProductByIdRequest) Reset() {
	*x = GetProductByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductByIdRequest) ProtoMessage() {}

func (x *GetProductByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductByIdRequest.ProtoReflect.Descriptor instead.
func (*GetProductByIdRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Price           *uint64 `protobuf:"varint,2,opt,name=price,proto3,oneof" json:"price,omitempty"`
	Description     *string `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Size            *string `protobuf:"bytes,4,opt,name=size,proto3,oneof" json:"size,omitempty"`
	Color           *string `protobuf:"bytes,5,opt,name=color,proto3,oneof" json:"color,omitempty"`
	Brand           *string `protobuf:"bytes,6,opt,name=brand,proto3,oneof" json:"brand,omitempty"`
	Shade           *string `protobuf:"bytes,7,opt,name=shade,proto3,oneof" json:"shade,omitempty"`
	WashCare        *string `protobuf:"bytes,8,opt,name=washCare,proto3,oneof" json:"washCare,omitempty"`
	Stretchable     *string `protobuf:"bytes,9,opt,name=stretchable,proto3,oneof" json:"stretchable,omitempty"`
	Distress        *string `protobuf:"bytes,10,opt,name=distress,proto3,oneof" json:"distress,omitempty"`
	Features        *string `protobuf:"bytes,11,opt,name=features,proto3,oneof" json:"features,omitempty"`
	Fade            *string `protobuf:"bytes,12,opt,name=fade,proto3,oneof" json:"fade,omitempty"`
	Fabric          *string `protobuf:"bytes,13,opt,name=fabric,proto3,oneof" json:"fabric,omitempty"`
	Category        *string `protobuf:"bytes,14,opt,name=category,proto3,oneof" json:"category,omitempty"`
	CountryOfOrigin *string `protobuf:"bytes,15,opt,name=countryOfOrigin,proto3,oneof" json:"countryOfOrigin,omitempty"`
	Discount        *uint32 `protobuf:"varint,16,opt,name=discount,proto3,oneof" json:"discount,omitempty"`
	Quantity        *uint64 `protobuf:"varint,17,opt,name=quantity,proto3,oneof" json:"quantity,omitempty"`
}

func (x *UpdateProductRequest) Reset() {
	*x = UpdateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductRequest) ProtoMessage() {}

func (x *UpdateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProductRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateProductRequest) GetPrice() uint64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *UpdateProductRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdateProductRequest) GetSize() string {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return ""
}

func (x *UpdateProductRequest) GetColor() string {
	if x != nil && x.Color != nil {
		return *x.Color
	}
	return ""
}

func (x *UpdateProductRequest) GetBrand() string {
	if x != nil && x.Brand != nil {
		return *x.Brand
	}
	return ""
}

func (x *UpdateProductRequest) GetShade() string {
	if x != nil && x.Shade != nil {
		return *x.Shade
	}
	return ""
}

func (x *UpdateProductRequest) GetWashCare() string {
	if x != nil && x.WashCare != nil {
		return *x.WashCare
	}
	return ""
}

func (x *UpdateProductRequest) GetStretchable() string {
	if x != nil && x.Stretchable != nil {
		return *x.Stretchable
	}
	return ""
}

func (x *UpdateProductRequest) GetDistress() string {
	if x != nil && x.Distress != nil {
		return *x.Distress
	}
	return ""
}

func (x *UpdateProductRequest) GetFeatures() string {
	if x != nil && x.Features != nil {
		return *x.Features
	}
	return ""
}

func (x *UpdateProductRequest) GetFade() string {
	if x != nil && x.Fade != nil {
		return *x.Fade
	}
	return ""
}

func (x *UpdateProductRequest) GetFabric() string {
	if x != nil && x.Fabric != nil {
		return *x.Fabric
	}
	return ""
}

func (x *UpdateProductRequest) GetCategory() string {
	if x != nil && x.Category != nil {
		return *x.Category
	}
	return ""
}

func (x *UpdateProductRequest) GetCountryOfOrigin() string {
	if x != nil && x.CountryOfOrigin != nil {
		return *x.CountryOfOrigin
	}
	return ""
}

func (x *UpdateProductRequest) GetDiscount() uint32 {
	if x != nil && x.Discount != nil {
		return *x.Discount
	}
	return 0
}

func (x *UpdateProductRequest) GetQuantity() uint64 {
	if x != nil && x.Quantity != nil {
		return *x.Quantity
	}
	return 0
}

type DeleteProductByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteProductByIdRequest) Reset() {
	*x = DeleteProductByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductByIdRequest) ProtoMessage() {}

func (x *DeleteProductByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductByIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductByIdRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteProductByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x05, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x68, 0x61, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x68, 0x61, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x73, 0x68, 0x43, 0x61, 0x72, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x73, 0x68, 0x43, 0x61, 0x72, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x74, 0x63, 0x68, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x61,
	0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x61, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x66, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x4f, 0x66, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x93, 0x05, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x64, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x61, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61,
	0x73, 0x68, 0x43, 0x61, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61,
	0x73, 0x68, 0x43, 0x61, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x74, 0x63,
	0x68, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72,
	0x65, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x61, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x61, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x4f, 0x66, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x66, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3a,
	0x0a, 0x0a, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xfd, 0x05, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x48, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x19, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x04, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x05, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x64, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x05, 0x73, 0x68, 0x61, 0x64, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1f, 0x0a, 0x08, 0x77, 0x61, 0x73, 0x68, 0x43, 0x61, 0x72, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x08, 0x77, 0x61, 0x73, 0x68, 0x43, 0x61, 0x72, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x74,
	0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0a, 0x52, 0x08,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x66,
	0x61, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0b, 0x52, 0x04, 0x66, 0x61, 0x64,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x0c, 0x52, 0x06, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x88, 0x01,
	0x01, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x0d, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x88,
	0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x66, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0e, 0x52, 0x0f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x66, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0d, 0x48, 0x0f, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x04, 0x48, 0x10, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x68, 0x61, 0x64, 0x65, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x43, 0x61, 0x72, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73,
	0x74, 0x72, 0x65, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x66, 0x61, 0x64, 0x65, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x4f, 0x66, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x2a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32,
	0xb3, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x46, 0x0a, 0x0a, 0x67,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a,
	0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_product_proto_goTypes = []interface{}{
	(*CreateProductRequest)(nil),     // 0: product.CreateProductRequest
	(*ProductResponse)(nil),          // 1: product.ProductResponse
	(*GetProductByIdRequest)(nil),    // 2: product.GetProductByIdRequest
	(*UpdateProductRequest)(nil),     // 3: product.UpdateProductRequest
	(*DeleteProductByIdRequest)(nil), // 4: product.DeleteProductByIdRequest
	(*timestamppb.Timestamp)(nil),    // 5: google.protobuf.Timestamp
}
var file_product_proto_depIdxs = []int32{
	5,  // 0: product.CreateProductRequest.createdAt:type_name -> google.protobuf.Timestamp
	5,  // 1: product.CreateProductRequest.updatedAt:type_name -> google.protobuf.Timestamp
	5,  // 2: product.CreateProductRequest.archivedAt:type_name -> google.protobuf.Timestamp
	5,  // 3: product.ProductResponse.createdAt:type_name -> google.protobuf.Timestamp
	5,  // 4: product.ProductResponse.updatedAt:type_name -> google.protobuf.Timestamp
	5,  // 5: product.ProductResponse.archivedAt:type_name -> google.protobuf.Timestamp
	2,  // 6: product.Product.getProduct:input_type -> product.GetProductByIdRequest
	0,  // 7: product.Product.createProduct:input_type -> product.CreateProductRequest
	3,  // 8: product.Product.updateProduct:input_type -> product.UpdateProductRequest
	4,  // 9: product.Product.deleteProduct:input_type -> product.DeleteProductByIdRequest
	1,  // 10: product.Product.getProduct:output_type -> product.ProductResponse
	1,  // 11: product.Product.createProduct:output_type -> product.ProductResponse
	1,  // 12: product.Product.updateProduct:output_type -> product.ProductResponse
	1,  // 13: product.Product.deleteProduct:output_type -> product.ProductResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductRequest); i {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductResponse); i {
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
		file_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductByIdRequest); i {
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
		file_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductRequest); i {
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
		file_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductByIdRequest); i {
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
	file_product_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
