// Code generated by protoc-gen-gogo.
// source: openstack/image.proto
// DO NOT EDIT!

package echopb_openstack

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Image model
// Does not include the literal image data; just metadata.
// returned by listing images, and by fetching a specific image.
type Image struct {
	// ID is the image UUID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name is the human-readable display name for the image.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Status is the image status. It can be "queued" or "active"
	// See imageservice/v2/images/type.go
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	// Tags is a list of image tags. Tags are arbitrarily defined strings
	// attached to an image.
	Tags []string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	// ContainerFormat is the format of the container.
	// Valid values are ami, ari, aki, bare, and ovf.
	ContainerFormat string `protobuf:"bytes,5,opt,name=container_format,json=containerFormat,proto3" json:"container_format,omitempty"`
	// DiskFormat is the format of the disk.
	// If set, valid values are ami, ari, aki, vhd, vmdk, raw, qcow2, vdi, and iso.
	DiskFormat string `protobuf:"bytes,6,opt,name=disk_format,json=diskFormat,proto3" json:"disk_format,omitempty"`
	// MinDiskGigabytes is the amount of disk space in GB that is required to boot the image.
	MinDisk int32 `protobuf:"varint,7,opt,name=min_disk,json=minDisk,proto3" json:"min_disk,omitempty"`
	// MinRAMMegabytes [optional] is the amount of RAM in MB that is required to boot the image.
	MinRam int32 `protobuf:"varint,8,opt,name=min_ram,json=minRam,proto3" json:"min_ram,omitempty"`
	// Owner is the tenant the image belongs to.
	Owner string `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	// Protected is whether the image is deletable or not.
	Protected bool `protobuf:"varint,10,opt,name=protected,proto3" json:"protected,omitempty"`
	// Visibility defines who can see/use the image.
	Visibility string `protobuf:"bytes,11,opt,name=visibility,proto3" json:"visibility,omitempty"`
	// Checksum is the checksum of the data that's associated with the image
	Checksum string `protobuf:"bytes,12,opt,name=checksum,proto3" json:"checksum,omitempty"`
	// SizeBytes is the size of the data that's associated with the image.
	Size_ int64 `protobuf:"varint,13,opt,name=size,proto3" json:"size,omitempty"`
	// Metadata is a set of metadata associated with the image.
	// Image metadata allow for meaningfully define the image properties
	// and tags. See http://docs.openstack.org/developer/glance/metadefs-concepts.html.
	Metadata map[string]string `protobuf:"bytes,14,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Properties is a set of key-value pairs, if any, that are associated with the image.
	Properties map[string]string `protobuf:"bytes,15,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// CreatedAt is the date when the image has been created.
	// google.protobuf.Timestamp created_at = 16;
	CreateAt string `protobuf:"bytes,16,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	// UpdatedAt is the date when the last change has been made to the image or it's properties.
	// google.protobuf.Timestamp updated_at = 17;
	UpdatedAt string `protobuf:"bytes,17,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// File is the trailing path after the glance endpoint that represent the location
	// of the image or the path to retrieve it.
	File string `protobuf:"bytes,18,opt,name=file,proto3" json:"file,omitempty"`
	// Schema is the path to the JSON-schema that represent the image or image entity.
	Schema string `protobuf:"bytes,19,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptorImage, []int{0} }

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Image) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Image) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Image) GetContainerFormat() string {
	if m != nil {
		return m.ContainerFormat
	}
	return ""
}

func (m *Image) GetDiskFormat() string {
	if m != nil {
		return m.DiskFormat
	}
	return ""
}

func (m *Image) GetMinDisk() int32 {
	if m != nil {
		return m.MinDisk
	}
	return 0
}

func (m *Image) GetMinRam() int32 {
	if m != nil {
		return m.MinRam
	}
	return 0
}

func (m *Image) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Image) GetProtected() bool {
	if m != nil {
		return m.Protected
	}
	return false
}

func (m *Image) GetVisibility() string {
	if m != nil {
		return m.Visibility
	}
	return ""
}

func (m *Image) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *Image) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *Image) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Image) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Image) GetCreateAt() string {
	if m != nil {
		return m.CreateAt
	}
	return ""
}

func (m *Image) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Image) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *Image) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func init() {
	proto.RegisterType((*Image)(nil), "echopb.openstack.Image")
}
func (m *Image) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Image) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintImage(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintImage(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Status) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintImage(dAtA, i, uint64(len(m.Status)))
		i += copy(dAtA[i:], m.Status)
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.ContainerFormat) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintImage(dAtA, i, uint64(len(m.ContainerFormat)))
		i += copy(dAtA[i:], m.ContainerFormat)
	}
	if len(m.DiskFormat) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintImage(dAtA, i, uint64(len(m.DiskFormat)))
		i += copy(dAtA[i:], m.DiskFormat)
	}
	if m.MinDisk != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintImage(dAtA, i, uint64(m.MinDisk))
	}
	if m.MinRam != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintImage(dAtA, i, uint64(m.MinRam))
	}
	if len(m.Owner) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintImage(dAtA, i, uint64(len(m.Owner)))
		i += copy(dAtA[i:], m.Owner)
	}
	if m.Protected {
		dAtA[i] = 0x50
		i++
		if m.Protected {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Visibility) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintImage(dAtA, i, uint64(len(m.Visibility)))
		i += copy(dAtA[i:], m.Visibility)
	}
	if len(m.Checksum) > 0 {
		dAtA[i] = 0x62
		i++
		i = encodeVarintImage(dAtA, i, uint64(len(m.Checksum)))
		i += copy(dAtA[i:], m.Checksum)
	}
	if m.Size_ != 0 {
		dAtA[i] = 0x68
		i++
		i = encodeVarintImage(dAtA, i, uint64(m.Size_))
	}
	if len(m.Metadata) > 0 {
		for k, _ := range m.Metadata {
			dAtA[i] = 0x72
			i++
			v := m.Metadata[k]
			mapSize := 1 + len(k) + sovImage(uint64(len(k))) + 1 + len(v) + sovImage(uint64(len(v)))
			i = encodeVarintImage(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintImage(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintImage(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.Properties) > 0 {
		for k, _ := range m.Properties {
			dAtA[i] = 0x7a
			i++
			v := m.Properties[k]
			mapSize := 1 + len(k) + sovImage(uint64(len(k))) + 1 + len(v) + sovImage(uint64(len(v)))
			i = encodeVarintImage(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintImage(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintImage(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.CreateAt) > 0 {
		dAtA[i] = 0x82
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintImage(dAtA, i, uint64(len(m.CreateAt)))
		i += copy(dAtA[i:], m.CreateAt)
	}
	if len(m.UpdatedAt) > 0 {
		dAtA[i] = 0x8a
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintImage(dAtA, i, uint64(len(m.UpdatedAt)))
		i += copy(dAtA[i:], m.UpdatedAt)
	}
	if len(m.File) > 0 {
		dAtA[i] = 0x92
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintImage(dAtA, i, uint64(len(m.File)))
		i += copy(dAtA[i:], m.File)
	}
	if len(m.Schema) > 0 {
		dAtA[i] = 0x9a
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintImage(dAtA, i, uint64(len(m.Schema)))
		i += copy(dAtA[i:], m.Schema)
	}
	return i, nil
}

func encodeFixed64Image(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Image(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintImage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Image) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovImage(uint64(l))
		}
	}
	l = len(m.ContainerFormat)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	l = len(m.DiskFormat)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	if m.MinDisk != 0 {
		n += 1 + sovImage(uint64(m.MinDisk))
	}
	if m.MinRam != 0 {
		n += 1 + sovImage(uint64(m.MinRam))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	if m.Protected {
		n += 2
	}
	l = len(m.Visibility)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	l = len(m.Checksum)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	if m.Size_ != 0 {
		n += 1 + sovImage(uint64(m.Size_))
	}
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovImage(uint64(len(k))) + 1 + len(v) + sovImage(uint64(len(v)))
			n += mapEntrySize + 1 + sovImage(uint64(mapEntrySize))
		}
	}
	if len(m.Properties) > 0 {
		for k, v := range m.Properties {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovImage(uint64(len(k))) + 1 + len(v) + sovImage(uint64(len(v)))
			n += mapEntrySize + 1 + sovImage(uint64(mapEntrySize))
		}
	}
	l = len(m.CreateAt)
	if l > 0 {
		n += 2 + l + sovImage(uint64(l))
	}
	l = len(m.UpdatedAt)
	if l > 0 {
		n += 2 + l + sovImage(uint64(l))
	}
	l = len(m.File)
	if l > 0 {
		n += 2 + l + sovImage(uint64(l))
	}
	l = len(m.Schema)
	if l > 0 {
		n += 2 + l + sovImage(uint64(l))
	}
	return n
}

func sovImage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozImage(x uint64) (n int) {
	return sovImage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Image) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImage
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
			return fmt.Errorf("proto: Image: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Image: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerFormat", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerFormat = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DiskFormat", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DiskFormat = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDisk", wireType)
			}
			m.MinDisk = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinDisk |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinRam", wireType)
			}
			m.MinRam = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinRam |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protected", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Protected = bool(v != 0)
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Visibility", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Visibility = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthImage
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Metadata == nil {
				m.Metadata = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowImage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowImage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthImage
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.Metadata[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.Metadata[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthImage
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Properties == nil {
				m.Properties = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowImage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowImage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthImage
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.Properties[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.Properties[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreateAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field File", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.File = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schema", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schema = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImage
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
func skipImage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowImage
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
					return 0, ErrIntOverflowImage
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
					return 0, ErrIntOverflowImage
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
				return 0, ErrInvalidLengthImage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowImage
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
				next, err := skipImage(dAtA[start:])
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
	ErrInvalidLengthImage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowImage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("openstack/image.proto", fileDescriptorImage) }

var fileDescriptorImage = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x8e, 0xd3, 0x4c,
	0x10, 0xfd, 0x9c, 0x4c, 0x32, 0x49, 0xe5, 0x9b, 0x89, 0x69, 0xfe, 0x9a, 0x00, 0x1e, 0x0b, 0x09,
	0x61, 0x36, 0x8e, 0x04, 0x1b, 0x04, 0x62, 0x11, 0xc4, 0x8f, 0x58, 0x20, 0x21, 0x5f, 0x20, 0xea,
	0xd8, 0x95, 0xa4, 0xe5, 0xb4, 0xdb, 0x72, 0x97, 0x07, 0x85, 0x93, 0x70, 0x24, 0x24, 0x36, 0x1c,
	0x01, 0x85, 0x8b, 0xa0, 0x6e, 0x7b, 0x3c, 0x30, 0x12, 0x0b, 0x76, 0xaf, 0xde, 0x7b, 0x55, 0xea,
	0xea, 0x7a, 0x70, 0x53, 0x97, 0x58, 0x18, 0x12, 0x69, 0x3e, 0x97, 0x4a, 0x6c, 0x30, 0x2e, 0x2b,
	0x4d, 0x9a, 0xf9, 0x98, 0x6e, 0x75, 0xb9, 0x8a, 0x3b, 0x75, 0x76, 0xb6, 0xd1, 0x7a, 0xb3, 0xc3,
	0xb9, 0xd3, 0x57, 0xf5, 0x7a, 0x4e, 0x52, 0xa1, 0x21, 0xa1, 0xca, 0xa6, 0xe5, 0xc1, 0xb7, 0x01,
	0x0c, 0xde, 0xdb, 0x11, 0xec, 0x14, 0x7a, 0x32, 0xe3, 0x5e, 0xe8, 0x45, 0xe3, 0xa4, 0x27, 0x33,
	0xc6, 0xe0, 0xa8, 0x10, 0x0a, 0x79, 0xcf, 0x31, 0x0e, 0xb3, 0x5b, 0x30, 0x34, 0x24, 0xa8, 0x36,
	0xbc, 0xef, 0xd8, 0xb6, 0xb2, 0x5e, 0x12, 0x1b, 0xc3, 0x8f, 0xc2, 0xbe, 0xf5, 0x5a, 0xcc, 0x1e,
	0x83, 0x9f, 0xea, 0x82, 0x84, 0x2c, 0xb0, 0x5a, 0xae, 0x75, 0xa5, 0x04, 0xf1, 0x81, 0xeb, 0x9a,
	0x76, 0xfc, 0x5b, 0x47, 0xb3, 0x33, 0x98, 0x64, 0xd2, 0xe4, 0x17, 0xae, 0xa1, 0x73, 0x81, 0xa5,
	0x5a, 0xc3, 0x1d, 0x18, 0x29, 0x59, 0x2c, 0x2d, 0xc3, 0x8f, 0x43, 0x2f, 0x1a, 0x24, 0xc7, 0x4a,
	0x16, 0xaf, 0xa5, 0xc9, 0xd9, 0x6d, 0xb0, 0x70, 0x59, 0x09, 0xc5, 0x47, 0x4e, 0x19, 0x2a, 0x59,
	0x24, 0x42, 0xb1, 0x1b, 0x30, 0xd0, 0x9f, 0x0a, 0xac, 0xf8, 0xd8, 0x8d, 0x6b, 0x0a, 0x76, 0x0f,
	0xc6, 0x76, 0x71, 0x4c, 0x09, 0x33, 0x0e, 0xa1, 0x17, 0x8d, 0x92, 0x4b, 0x82, 0x05, 0x00, 0xe7,
	0xd2, 0xc8, 0x95, 0xdc, 0x49, 0xda, 0xf3, 0x49, 0xf3, 0x8e, 0x4b, 0x86, 0xcd, 0x60, 0x94, 0x6e,
	0x31, 0xcd, 0x4d, 0xad, 0xf8, 0xff, 0x4e, 0xed, 0x6a, 0xfb, 0x07, 0x46, 0x7e, 0x46, 0x7e, 0x12,
	0x7a, 0x51, 0x3f, 0x71, 0x98, 0x2d, 0x60, 0xa4, 0x90, 0x44, 0x26, 0x48, 0xf0, 0xd3, 0xb0, 0x1f,
	0x4d, 0x9e, 0x3c, 0x8c, 0xaf, 0xde, 0x28, 0x76, 0xdf, 0x1f, 0x7f, 0x68, 0x7d, 0x6f, 0x0a, 0xaa,
	0xf6, 0x49, 0xd7, 0xc6, 0xde, 0x01, 0x94, 0x95, 0x2e, 0xb1, 0x22, 0x89, 0x86, 0x4f, 0xdd, 0x90,
	0x47, 0x7f, 0x1b, 0xf2, 0xb1, 0x73, 0x36, 0x63, 0x7e, 0x6b, 0x65, 0x77, 0x61, 0x9c, 0x56, 0x28,
	0x08, 0x97, 0x82, 0xb8, 0xdf, 0x3e, 0xde, 0x11, 0x0b, 0x62, 0xf7, 0x01, 0xea, 0x32, 0x13, 0x84,
	0x99, 0x55, 0xaf, 0x39, 0x75, 0xdc, 0x32, 0x0b, 0xb2, 0xbb, 0xad, 0xe5, 0x0e, 0x39, 0x6b, 0xb2,
	0x60, 0xb1, 0xcb, 0x42, 0xba, 0x45, 0x25, 0xf8, 0xf5, 0x36, 0x0b, 0xae, 0x9a, 0xbd, 0x80, 0x93,
	0x3f, 0x76, 0x61, 0x3e, 0xf4, 0x73, 0xdc, 0xb7, 0xc9, 0xb2, 0xd0, 0x9e, 0xe6, 0x5c, 0xec, 0xea,
	0x8b, 0x6c, 0x35, 0xc5, 0xf3, 0xde, 0x33, 0x6f, 0xf6, 0x12, 0xa6, 0x57, 0x76, 0xf8, 0x97, 0xf6,
	0x57, 0xfe, 0xd7, 0x43, 0xe0, 0x7d, 0x3f, 0x04, 0xde, 0x8f, 0x43, 0xe0, 0x7d, 0xf9, 0x19, 0xfc,
	0xb7, 0x1a, 0xba, 0x98, 0x3f, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0xed, 0x00, 0xbc, 0x9b, 0x32,
	0x03, 0x00, 0x00,
}