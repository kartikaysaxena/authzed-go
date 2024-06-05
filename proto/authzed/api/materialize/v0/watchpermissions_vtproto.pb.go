// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: authzed/api/materialize/v0/watchpermissions.proto

package v0

import (
	fmt "fmt"
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *WatchPermissionsRequest) CloneVT() *WatchPermissionsRequest {
	if m == nil {
		return (*WatchPermissionsRequest)(nil)
	}
	r := new(WatchPermissionsRequest)
	if rhs := m.Permissions; rhs != nil {
		tmpContainer := make([]*WatchedPermission, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Permissions = tmpContainer
	}
	if rhs := m.OptionalStartingAfter; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface{ CloneVT() *v1.ZedToken }); ok {
			r.OptionalStartingAfter = vtpb.CloneVT()
		} else {
			r.OptionalStartingAfter = proto.Clone(rhs).(*v1.ZedToken)
		}
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *WatchPermissionsRequest) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *WatchedPermission) CloneVT() *WatchedPermission {
	if m == nil {
		return (*WatchedPermission)(nil)
	}
	r := new(WatchedPermission)
	r.ResourceType = m.ResourceType
	r.Permission = m.Permission
	r.SubjectType = m.SubjectType
	r.OptionalSubjectRelation = m.OptionalSubjectRelation
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *WatchedPermission) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *WatchPermissionsResponse) CloneVT() *WatchPermissionsResponse {
	if m == nil {
		return (*WatchPermissionsResponse)(nil)
	}
	r := new(WatchPermissionsResponse)
	if m.Response != nil {
		r.Response = m.Response.(interface {
			CloneVT() isWatchPermissionsResponse_Response
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *WatchPermissionsResponse) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *WatchPermissionsResponse_Change) CloneVT() isWatchPermissionsResponse_Response {
	if m == nil {
		return (*WatchPermissionsResponse_Change)(nil)
	}
	r := new(WatchPermissionsResponse_Change)
	r.Change = m.Change.CloneVT()
	return r
}

func (m *WatchPermissionsResponse_CompletedRevision) CloneVT() isWatchPermissionsResponse_Response {
	if m == nil {
		return (*WatchPermissionsResponse_CompletedRevision)(nil)
	}
	r := new(WatchPermissionsResponse_CompletedRevision)
	if rhs := m.CompletedRevision; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface{ CloneVT() *v1.ZedToken }); ok {
			r.CompletedRevision = vtpb.CloneVT()
		} else {
			r.CompletedRevision = proto.Clone(rhs).(*v1.ZedToken)
		}
	}
	return r
}

func (m *PermissionChange) CloneVT() *PermissionChange {
	if m == nil {
		return (*PermissionChange)(nil)
	}
	r := new(PermissionChange)
	r.Permission = m.Permission
	r.Permissionship = m.Permissionship
	if rhs := m.Revision; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface{ CloneVT() *v1.ZedToken }); ok {
			r.Revision = vtpb.CloneVT()
		} else {
			r.Revision = proto.Clone(rhs).(*v1.ZedToken)
		}
	}
	if rhs := m.Resource; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface{ CloneVT() *v1.ObjectReference }); ok {
			r.Resource = vtpb.CloneVT()
		} else {
			r.Resource = proto.Clone(rhs).(*v1.ObjectReference)
		}
	}
	if rhs := m.Subject; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface{ CloneVT() *v1.SubjectReference }); ok {
			r.Subject = vtpb.CloneVT()
		} else {
			r.Subject = proto.Clone(rhs).(*v1.SubjectReference)
		}
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *PermissionChange) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *WatchPermissionsRequest) EqualVT(that *WatchPermissionsRequest) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Permissions) != len(that.Permissions) {
		return false
	}
	for i, vx := range this.Permissions {
		vy := that.Permissions[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &WatchedPermission{}
			}
			if q == nil {
				q = &WatchedPermission{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if equal, ok := interface{}(this.OptionalStartingAfter).(interface{ EqualVT(*v1.ZedToken) bool }); ok {
		if !equal.EqualVT(that.OptionalStartingAfter) {
			return false
		}
	} else if !proto.Equal(this.OptionalStartingAfter, that.OptionalStartingAfter) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *WatchPermissionsRequest) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*WatchPermissionsRequest)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *WatchedPermission) EqualVT(that *WatchedPermission) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ResourceType != that.ResourceType {
		return false
	}
	if this.Permission != that.Permission {
		return false
	}
	if this.SubjectType != that.SubjectType {
		return false
	}
	if this.OptionalSubjectRelation != that.OptionalSubjectRelation {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *WatchedPermission) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*WatchedPermission)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *WatchPermissionsResponse) EqualVT(that *WatchPermissionsResponse) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Response == nil && that.Response != nil {
		return false
	} else if this.Response != nil {
		if that.Response == nil {
			return false
		}
		if !this.Response.(interface {
			EqualVT(isWatchPermissionsResponse_Response) bool
		}).EqualVT(that.Response) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *WatchPermissionsResponse) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*WatchPermissionsResponse)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *WatchPermissionsResponse_Change) EqualVT(thatIface isWatchPermissionsResponse_Response) bool {
	that, ok := thatIface.(*WatchPermissionsResponse_Change)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Change, that.Change; p != q {
		if p == nil {
			p = &PermissionChange{}
		}
		if q == nil {
			q = &PermissionChange{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *WatchPermissionsResponse_CompletedRevision) EqualVT(thatIface isWatchPermissionsResponse_Response) bool {
	that, ok := thatIface.(*WatchPermissionsResponse_CompletedRevision)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.CompletedRevision, that.CompletedRevision; p != q {
		if p == nil {
			p = &v1.ZedToken{}
		}
		if q == nil {
			q = &v1.ZedToken{}
		}
		if equal, ok := interface{}(p).(interface{ EqualVT(*v1.ZedToken) bool }); ok {
			if !equal.EqualVT(q) {
				return false
			}
		} else if !proto.Equal(p, q) {
			return false
		}
	}
	return true
}

func (this *PermissionChange) EqualVT(that *PermissionChange) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if equal, ok := interface{}(this.Revision).(interface{ EqualVT(*v1.ZedToken) bool }); ok {
		if !equal.EqualVT(that.Revision) {
			return false
		}
	} else if !proto.Equal(this.Revision, that.Revision) {
		return false
	}
	if equal, ok := interface{}(this.Resource).(interface {
		EqualVT(*v1.ObjectReference) bool
	}); ok {
		if !equal.EqualVT(that.Resource) {
			return false
		}
	} else if !proto.Equal(this.Resource, that.Resource) {
		return false
	}
	if this.Permission != that.Permission {
		return false
	}
	if equal, ok := interface{}(this.Subject).(interface {
		EqualVT(*v1.SubjectReference) bool
	}); ok {
		if !equal.EqualVT(that.Subject) {
			return false
		}
	} else if !proto.Equal(this.Subject, that.Subject) {
		return false
	}
	if this.Permissionship != that.Permissionship {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *PermissionChange) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*PermissionChange)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *WatchPermissionsRequest) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WatchPermissionsRequest) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *WatchPermissionsRequest) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.OptionalStartingAfter != nil {
		if vtmsg, ok := interface{}(m.OptionalStartingAfter).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.OptionalStartingAfter)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Permissions) > 0 {
		for iNdEx := len(m.Permissions) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Permissions[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *WatchedPermission) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WatchedPermission) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *WatchedPermission) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.OptionalSubjectRelation) > 0 {
		i -= len(m.OptionalSubjectRelation)
		copy(dAtA[i:], m.OptionalSubjectRelation)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.OptionalSubjectRelation)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SubjectType) > 0 {
		i -= len(m.SubjectType)
		copy(dAtA[i:], m.SubjectType)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.SubjectType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Permission) > 0 {
		i -= len(m.Permission)
		copy(dAtA[i:], m.Permission)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Permission)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ResourceType) > 0 {
		i -= len(m.ResourceType)
		copy(dAtA[i:], m.ResourceType)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ResourceType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WatchPermissionsResponse) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WatchPermissionsResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *WatchPermissionsResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if vtmsg, ok := m.Response.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *WatchPermissionsResponse_Change) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *WatchPermissionsResponse_Change) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Change != nil {
		size, err := m.Change.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *WatchPermissionsResponse_CompletedRevision) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *WatchPermissionsResponse_CompletedRevision) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CompletedRevision != nil {
		if vtmsg, ok := interface{}(m.CompletedRevision).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.CompletedRevision)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *PermissionChange) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PermissionChange) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *PermissionChange) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.Permissionship != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Permissionship))
		i--
		dAtA[i] = 0x28
	}
	if m.Subject != nil {
		if vtmsg, ok := interface{}(m.Subject).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Subject)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Permission) > 0 {
		i -= len(m.Permission)
		copy(dAtA[i:], m.Permission)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Permission)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Resource != nil {
		if vtmsg, ok := interface{}(m.Resource).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Resource)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Revision != nil {
		if vtmsg, ok := interface{}(m.Revision).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Revision)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WatchPermissionsRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Permissions) > 0 {
		for _, e := range m.Permissions {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.OptionalStartingAfter != nil {
		if size, ok := interface{}(m.OptionalStartingAfter).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.OptionalStartingAfter)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *WatchedPermission) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ResourceType)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.Permission)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.SubjectType)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.OptionalSubjectRelation)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *WatchPermissionsResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Response.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *WatchPermissionsResponse_Change) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Change != nil {
		l = m.Change.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	return n
}
func (m *WatchPermissionsResponse_CompletedRevision) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CompletedRevision != nil {
		if size, ok := interface{}(m.CompletedRevision).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.CompletedRevision)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	return n
}
func (m *PermissionChange) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Revision != nil {
		if size, ok := interface{}(m.Revision).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Revision)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Resource != nil {
		if size, ok := interface{}(m.Resource).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Resource)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.Permission)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Subject != nil {
		if size, ok := interface{}(m.Subject).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Subject)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Permissionship != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Permissionship))
	}
	n += len(m.unknownFields)
	return n
}

func (m *WatchPermissionsRequest) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
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
			return fmt.Errorf("proto: WatchPermissionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchPermissionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permissions = append(m.Permissions, &WatchedPermission{})
			if err := m.Permissions[len(m.Permissions)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalStartingAfter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OptionalStartingAfter == nil {
				m.OptionalStartingAfter = &v1.ZedToken{}
			}
			if unmarshal, ok := interface{}(m.OptionalStartingAfter).(interface {
				UnmarshalVT([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.OptionalStartingAfter); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WatchedPermission) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
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
			return fmt.Errorf("proto: WatchedPermission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchedPermission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permission", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permission = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalSubjectRelation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OptionalSubjectRelation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WatchPermissionsResponse) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
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
			return fmt.Errorf("proto: WatchPermissionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchPermissionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Change", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.Response.(*WatchPermissionsResponse_Change); ok {
				if err := oneof.Change.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &PermissionChange{}
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Response = &WatchPermissionsResponse_Change{Change: v}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletedRevision", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.Response.(*WatchPermissionsResponse_CompletedRevision); ok {
				if unmarshal, ok := interface{}(oneof.CompletedRevision).(interface {
					UnmarshalVT([]byte) error
				}); ok {
					if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
						return err
					}
				} else {
					if err := proto.Unmarshal(dAtA[iNdEx:postIndex], oneof.CompletedRevision); err != nil {
						return err
					}
				}
			} else {
				v := &v1.ZedToken{}
				if unmarshal, ok := interface{}(v).(interface {
					UnmarshalVT([]byte) error
				}); ok {
					if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
						return err
					}
				} else {
					if err := proto.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
						return err
					}
				}
				m.Response = &WatchPermissionsResponse_CompletedRevision{CompletedRevision: v}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PermissionChange) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
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
			return fmt.Errorf("proto: PermissionChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PermissionChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Revision == nil {
				m.Revision = &v1.ZedToken{}
			}
			if unmarshal, ok := interface{}(m.Revision).(interface {
				UnmarshalVT([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.Revision); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Resource == nil {
				m.Resource = &v1.ObjectReference{}
			}
			if unmarshal, ok := interface{}(m.Resource).(interface {
				UnmarshalVT([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.Resource); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permission", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permission = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Subject == nil {
				m.Subject = &v1.SubjectReference{}
			}
			if unmarshal, ok := interface{}(m.Subject).(interface {
				UnmarshalVT([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.Subject); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissionship", wireType)
			}
			m.Permissionship = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Permissionship |= PermissionChange_Permissionship(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
