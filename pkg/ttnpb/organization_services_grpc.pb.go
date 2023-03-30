// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: lorawan-stack/api/organization_services.proto

package ttnpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrganizationRegistryClient is the client API for OrganizationRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationRegistryClient interface {
	// Create a new organization. This also sets the given user as
	// first collaborator with all possible rights.
	Create(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
	// Get the organization with the given identifiers, selecting the fields specified
	// in the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	Get(ctx context.Context, in *GetOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
	// List organizations where the given user or organization is a direct collaborator.
	// If no user or organization is given, this returns the organizations the caller
	// has access to.
	// Similar to Get, this selects the fields given by the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	List(ctx context.Context, in *ListOrganizationsRequest, opts ...grpc.CallOption) (*Organizations, error)
	// Update the organization, changing the fields specified by the field mask to the provided values.
	Update(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
	// Delete the organization. This may not release the organization ID for reuse.
	Delete(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Restore a recently deleted organization.
	//
	// Deployment configuration may specify if, and for how long after deletion,
	// entities can be restored.
	Restore(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Purge the organization. This will release the organization ID for reuse.
	// The user is responsible for clearing data from any (external) integrations
	// that may store and expose data by user or organization ID.
	Purge(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type organizationRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationRegistryClient(cc grpc.ClientConnInterface) OrganizationRegistryClient {
	return &organizationRegistryClient{cc}
}

func (c *organizationRegistryClient) Create(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error) {
	out := new(Organization)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) Get(ctx context.Context, in *GetOrganizationRequest, opts ...grpc.CallOption) (*Organization, error) {
	out := new(Organization)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) List(ctx context.Context, in *ListOrganizationsRequest, opts ...grpc.CallOption) (*Organizations, error) {
	out := new(Organizations)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) Update(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error) {
	out := new(Organization)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) Delete(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) Restore(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Restore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) Purge(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Purge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationRegistryServer is the server API for OrganizationRegistry service.
// All implementations must embed UnimplementedOrganizationRegistryServer
// for forward compatibility
type OrganizationRegistryServer interface {
	// Create a new organization. This also sets the given user as
	// first collaborator with all possible rights.
	Create(context.Context, *CreateOrganizationRequest) (*Organization, error)
	// Get the organization with the given identifiers, selecting the fields specified
	// in the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	Get(context.Context, *GetOrganizationRequest) (*Organization, error)
	// List organizations where the given user or organization is a direct collaborator.
	// If no user or organization is given, this returns the organizations the caller
	// has access to.
	// Similar to Get, this selects the fields given by the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	List(context.Context, *ListOrganizationsRequest) (*Organizations, error)
	// Update the organization, changing the fields specified by the field mask to the provided values.
	Update(context.Context, *UpdateOrganizationRequest) (*Organization, error)
	// Delete the organization. This may not release the organization ID for reuse.
	Delete(context.Context, *OrganizationIdentifiers) (*emptypb.Empty, error)
	// Restore a recently deleted organization.
	//
	// Deployment configuration may specify if, and for how long after deletion,
	// entities can be restored.
	Restore(context.Context, *OrganizationIdentifiers) (*emptypb.Empty, error)
	// Purge the organization. This will release the organization ID for reuse.
	// The user is responsible for clearing data from any (external) integrations
	// that may store and expose data by user or organization ID.
	Purge(context.Context, *OrganizationIdentifiers) (*emptypb.Empty, error)
	mustEmbedUnimplementedOrganizationRegistryServer()
}

// UnimplementedOrganizationRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedOrganizationRegistryServer struct {
}

func (UnimplementedOrganizationRegistryServer) Create(context.Context, *CreateOrganizationRequest) (*Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrganizationRegistryServer) Get(context.Context, *GetOrganizationRequest) (*Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedOrganizationRegistryServer) List(context.Context, *ListOrganizationsRequest) (*Organizations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOrganizationRegistryServer) Update(context.Context, *UpdateOrganizationRequest) (*Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOrganizationRegistryServer) Delete(context.Context, *OrganizationIdentifiers) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedOrganizationRegistryServer) Restore(context.Context, *OrganizationIdentifiers) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedOrganizationRegistryServer) Purge(context.Context, *OrganizationIdentifiers) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purge not implemented")
}
func (UnimplementedOrganizationRegistryServer) mustEmbedUnimplementedOrganizationRegistryServer() {}

// UnsafeOrganizationRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationRegistryServer will
// result in compilation errors.
type UnsafeOrganizationRegistryServer interface {
	mustEmbedUnimplementedOrganizationRegistryServer()
}

func RegisterOrganizationRegistryServer(s grpc.ServiceRegistrar, srv OrganizationRegistryServer) {
	s.RegisterService(&OrganizationRegistry_ServiceDesc, srv)
}

func _OrganizationRegistry_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Create(ctx, req.(*CreateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Get(ctx, req.(*GetOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).List(ctx, req.(*ListOrganizationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Update(ctx, req.(*UpdateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Delete(ctx, req.(*OrganizationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Restore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Restore(ctx, req.(*OrganizationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_Purge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Purge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Purge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Purge(ctx, req.(*OrganizationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

// OrganizationRegistry_ServiceDesc is the grpc.ServiceDesc for OrganizationRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrganizationRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.OrganizationRegistry",
	HandlerType: (*OrganizationRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrganizationRegistry_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _OrganizationRegistry_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _OrganizationRegistry_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OrganizationRegistry_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _OrganizationRegistry_Delete_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _OrganizationRegistry_Restore_Handler,
		},
		{
			MethodName: "Purge",
			Handler:    _OrganizationRegistry_Purge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/organization_services.proto",
}

// OrganizationAccessClient is the client API for OrganizationAccess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationAccessClient interface {
	// List the rights the caller has on this organization.
	ListRights(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*Rights, error)
	// Create an API key scoped to this organization.
	// Organization API keys can give access to the organization itself, as well as
	// any application, gateway and OAuth client this organization is a collaborator of.
	CreateAPIKey(ctx context.Context, in *CreateOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// List the API keys for this organization.
	ListAPIKeys(ctx context.Context, in *ListOrganizationAPIKeysRequest, opts ...grpc.CallOption) (*APIKeys, error)
	// Get a single API key of this organization.
	GetAPIKey(ctx context.Context, in *GetOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// Update the rights of an API key of the organization.
	// The caller is required to have all assigned or/and removed rights.
	UpdateAPIKey(ctx context.Context, in *UpdateOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// DeteleAPIKey removes the API key of an organization.
	DeleteAPIKey(ctx context.Context, in *DeleteOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// Get the rights of a collaborator (member) of the organization.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(ctx context.Context, in *GetOrganizationCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the organization.
	// Organization collaborators can get access to the organization itself, as well as
	// any application, gateway and OAuth client this organization is a collaborator of.
	// The caller is required to have all assigned or/and removed rights.
	SetCollaborator(ctx context.Context, in *SetOrganizationCollaboratorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DeleteCollaborator removes the rights of a user in the organization.
	DeleteCollaborator(ctx context.Context, in *DeleteOrganizationCollaboratorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// List the collaborators on this organization.
	ListCollaborators(ctx context.Context, in *ListOrganizationCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error)
}

type organizationAccessClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationAccessClient(cc grpc.ClientConnInterface) OrganizationAccessClient {
	return &organizationAccessClient{cc}
}

func (c *organizationAccessClient) ListRights(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*Rights, error) {
	out := new(Rights)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/ListRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) CreateAPIKey(ctx context.Context, in *CreateOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/CreateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) ListAPIKeys(ctx context.Context, in *ListOrganizationAPIKeysRequest, opts ...grpc.CallOption) (*APIKeys, error) {
	out := new(APIKeys)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/ListAPIKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) GetAPIKey(ctx context.Context, in *GetOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/GetAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) UpdateAPIKey(ctx context.Context, in *UpdateOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/UpdateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) DeleteAPIKey(ctx context.Context, in *DeleteOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/DeleteAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) GetCollaborator(ctx context.Context, in *GetOrganizationCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error) {
	out := new(GetCollaboratorResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/GetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) SetCollaborator(ctx context.Context, in *SetOrganizationCollaboratorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/SetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) DeleteCollaborator(ctx context.Context, in *DeleteOrganizationCollaboratorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/DeleteCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) ListCollaborators(ctx context.Context, in *ListOrganizationCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error) {
	out := new(Collaborators)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/ListCollaborators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationAccessServer is the server API for OrganizationAccess service.
// All implementations must embed UnimplementedOrganizationAccessServer
// for forward compatibility
type OrganizationAccessServer interface {
	// List the rights the caller has on this organization.
	ListRights(context.Context, *OrganizationIdentifiers) (*Rights, error)
	// Create an API key scoped to this organization.
	// Organization API keys can give access to the organization itself, as well as
	// any application, gateway and OAuth client this organization is a collaborator of.
	CreateAPIKey(context.Context, *CreateOrganizationAPIKeyRequest) (*APIKey, error)
	// List the API keys for this organization.
	ListAPIKeys(context.Context, *ListOrganizationAPIKeysRequest) (*APIKeys, error)
	// Get a single API key of this organization.
	GetAPIKey(context.Context, *GetOrganizationAPIKeyRequest) (*APIKey, error)
	// Update the rights of an API key of the organization.
	// The caller is required to have all assigned or/and removed rights.
	UpdateAPIKey(context.Context, *UpdateOrganizationAPIKeyRequest) (*APIKey, error)
	// DeteleAPIKey removes the API key of an organization.
	DeleteAPIKey(context.Context, *DeleteOrganizationAPIKeyRequest) (*APIKey, error)
	// Get the rights of a collaborator (member) of the organization.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(context.Context, *GetOrganizationCollaboratorRequest) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the organization.
	// Organization collaborators can get access to the organization itself, as well as
	// any application, gateway and OAuth client this organization is a collaborator of.
	// The caller is required to have all assigned or/and removed rights.
	SetCollaborator(context.Context, *SetOrganizationCollaboratorRequest) (*emptypb.Empty, error)
	// DeleteCollaborator removes the rights of a user in the organization.
	DeleteCollaborator(context.Context, *DeleteOrganizationCollaboratorRequest) (*emptypb.Empty, error)
	// List the collaborators on this organization.
	ListCollaborators(context.Context, *ListOrganizationCollaboratorsRequest) (*Collaborators, error)
	mustEmbedUnimplementedOrganizationAccessServer()
}

// UnimplementedOrganizationAccessServer must be embedded to have forward compatible implementations.
type UnimplementedOrganizationAccessServer struct {
}

func (UnimplementedOrganizationAccessServer) ListRights(context.Context, *OrganizationIdentifiers) (*Rights, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRights not implemented")
}
func (UnimplementedOrganizationAccessServer) CreateAPIKey(context.Context, *CreateOrganizationAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAPIKey not implemented")
}
func (UnimplementedOrganizationAccessServer) ListAPIKeys(context.Context, *ListOrganizationAPIKeysRequest) (*APIKeys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAPIKeys not implemented")
}
func (UnimplementedOrganizationAccessServer) GetAPIKey(context.Context, *GetOrganizationAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIKey not implemented")
}
func (UnimplementedOrganizationAccessServer) UpdateAPIKey(context.Context, *UpdateOrganizationAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAPIKey not implemented")
}
func (UnimplementedOrganizationAccessServer) DeleteAPIKey(context.Context, *DeleteOrganizationAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAPIKey not implemented")
}
func (UnimplementedOrganizationAccessServer) GetCollaborator(context.Context, *GetOrganizationCollaboratorRequest) (*GetCollaboratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollaborator not implemented")
}
func (UnimplementedOrganizationAccessServer) SetCollaborator(context.Context, *SetOrganizationCollaboratorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCollaborator not implemented")
}
func (UnimplementedOrganizationAccessServer) DeleteCollaborator(context.Context, *DeleteOrganizationCollaboratorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollaborator not implemented")
}
func (UnimplementedOrganizationAccessServer) ListCollaborators(context.Context, *ListOrganizationCollaboratorsRequest) (*Collaborators, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollaborators not implemented")
}
func (UnimplementedOrganizationAccessServer) mustEmbedUnimplementedOrganizationAccessServer() {}

// UnsafeOrganizationAccessServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationAccessServer will
// result in compilation errors.
type UnsafeOrganizationAccessServer interface {
	mustEmbedUnimplementedOrganizationAccessServer()
}

func RegisterOrganizationAccessServer(s grpc.ServiceRegistrar, srv OrganizationAccessServer) {
	s.RegisterService(&OrganizationAccess_ServiceDesc, srv)
}

func _OrganizationAccess_ListRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).ListRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/ListRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).ListRights(ctx, req.(*OrganizationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_CreateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrganizationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).CreateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/CreateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).CreateAPIKey(ctx, req.(*CreateOrganizationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_ListAPIKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationAPIKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).ListAPIKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/ListAPIKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).ListAPIKeys(ctx, req.(*ListOrganizationAPIKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_GetAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).GetAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/GetAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).GetAPIKey(ctx, req.(*GetOrganizationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_UpdateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrganizationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).UpdateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/UpdateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).UpdateAPIKey(ctx, req.(*UpdateOrganizationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_DeleteAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrganizationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).DeleteAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/DeleteAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).DeleteAPIKey(ctx, req.(*DeleteOrganizationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_GetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).GetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/GetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).GetCollaborator(ctx, req.(*GetOrganizationCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_SetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetOrganizationCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).SetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/SetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).SetCollaborator(ctx, req.(*SetOrganizationCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_DeleteCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrganizationCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).DeleteCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/DeleteCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).DeleteCollaborator(ctx, req.(*DeleteOrganizationCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_ListCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationCollaboratorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).ListCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/ListCollaborators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).ListCollaborators(ctx, req.(*ListOrganizationCollaboratorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrganizationAccess_ServiceDesc is the grpc.ServiceDesc for OrganizationAccess service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrganizationAccess_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.OrganizationAccess",
	HandlerType: (*OrganizationAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRights",
			Handler:    _OrganizationAccess_ListRights_Handler,
		},
		{
			MethodName: "CreateAPIKey",
			Handler:    _OrganizationAccess_CreateAPIKey_Handler,
		},
		{
			MethodName: "ListAPIKeys",
			Handler:    _OrganizationAccess_ListAPIKeys_Handler,
		},
		{
			MethodName: "GetAPIKey",
			Handler:    _OrganizationAccess_GetAPIKey_Handler,
		},
		{
			MethodName: "UpdateAPIKey",
			Handler:    _OrganizationAccess_UpdateAPIKey_Handler,
		},
		{
			MethodName: "DeleteAPIKey",
			Handler:    _OrganizationAccess_DeleteAPIKey_Handler,
		},
		{
			MethodName: "GetCollaborator",
			Handler:    _OrganizationAccess_GetCollaborator_Handler,
		},
		{
			MethodName: "SetCollaborator",
			Handler:    _OrganizationAccess_SetCollaborator_Handler,
		},
		{
			MethodName: "DeleteCollaborator",
			Handler:    _OrganizationAccess_DeleteCollaborator_Handler,
		},
		{
			MethodName: "ListCollaborators",
			Handler:    _OrganizationAccess_ListCollaborators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/organization_services.proto",
}
