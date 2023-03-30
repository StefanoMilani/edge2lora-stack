// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: lorawan-stack/api/gateway_services.proto

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

// GatewayRegistryClient is the client API for GatewayRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayRegistryClient interface {
	// Create a new gateway. This also sets the given organization or user as
	// first collaborator with all possible rights.
	Create(ctx context.Context, in *CreateGatewayRequest, opts ...grpc.CallOption) (*Gateway, error)
	// Get the gateway with the given identifiers, selecting the fields specified
	// in the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	Get(ctx context.Context, in *GetGatewayRequest, opts ...grpc.CallOption) (*Gateway, error)
	// Get the identifiers of the gateway that has the given EUI registered.
	GetIdentifiersForEUI(ctx context.Context, in *GetGatewayIdentifiersForEUIRequest, opts ...grpc.CallOption) (*GatewayIdentifiers, error)
	// List gateways where the given user or organization is a direct collaborator.
	// If no user or organization is given, this returns the gateways the caller
	// has access to.
	// Similar to Get, this selects the fields given by the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	List(ctx context.Context, in *ListGatewaysRequest, opts ...grpc.CallOption) (*Gateways, error)
	// Update the gateway, changing the fields specified by the field mask to the provided values.
	Update(ctx context.Context, in *UpdateGatewayRequest, opts ...grpc.CallOption) (*Gateway, error)
	// Delete the gateway. This may not release the gateway ID for reuse, but it does release the EUI.
	Delete(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Restore a recently deleted gateway. This does not restore the EUI,
	// as that was released when deleting the gateway.
	//
	// Deployment configuration may specify if, and for how long after deletion,
	// entities can be restored.
	Restore(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Purge the gateway. This will release both gateway ID and EUI for reuse.
	// The gateway owner is responsible for clearing data from any (external) integrations
	// that may store and expose data by gateway ID.
	Purge(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type gatewayRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayRegistryClient(cc grpc.ClientConnInterface) GatewayRegistryClient {
	return &gatewayRegistryClient{cc}
}

func (c *gatewayRegistryClient) Create(ctx context.Context, in *CreateGatewayRequest, opts ...grpc.CallOption) (*Gateway, error) {
	out := new(Gateway)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayRegistry/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayRegistryClient) Get(ctx context.Context, in *GetGatewayRequest, opts ...grpc.CallOption) (*Gateway, error) {
	out := new(Gateway)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayRegistryClient) GetIdentifiersForEUI(ctx context.Context, in *GetGatewayIdentifiersForEUIRequest, opts ...grpc.CallOption) (*GatewayIdentifiers, error) {
	out := new(GatewayIdentifiers)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayRegistry/GetIdentifiersForEUI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayRegistryClient) List(ctx context.Context, in *ListGatewaysRequest, opts ...grpc.CallOption) (*Gateways, error) {
	out := new(Gateways)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayRegistryClient) Update(ctx context.Context, in *UpdateGatewayRequest, opts ...grpc.CallOption) (*Gateway, error) {
	out := new(Gateway)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayRegistry/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayRegistryClient) Delete(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayRegistryClient) Restore(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayRegistry/Restore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayRegistryClient) Purge(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayRegistry/Purge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayRegistryServer is the server API for GatewayRegistry service.
// All implementations must embed UnimplementedGatewayRegistryServer
// for forward compatibility
type GatewayRegistryServer interface {
	// Create a new gateway. This also sets the given organization or user as
	// first collaborator with all possible rights.
	Create(context.Context, *CreateGatewayRequest) (*Gateway, error)
	// Get the gateway with the given identifiers, selecting the fields specified
	// in the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	Get(context.Context, *GetGatewayRequest) (*Gateway, error)
	// Get the identifiers of the gateway that has the given EUI registered.
	GetIdentifiersForEUI(context.Context, *GetGatewayIdentifiersForEUIRequest) (*GatewayIdentifiers, error)
	// List gateways where the given user or organization is a direct collaborator.
	// If no user or organization is given, this returns the gateways the caller
	// has access to.
	// Similar to Get, this selects the fields given by the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	List(context.Context, *ListGatewaysRequest) (*Gateways, error)
	// Update the gateway, changing the fields specified by the field mask to the provided values.
	Update(context.Context, *UpdateGatewayRequest) (*Gateway, error)
	// Delete the gateway. This may not release the gateway ID for reuse, but it does release the EUI.
	Delete(context.Context, *GatewayIdentifiers) (*emptypb.Empty, error)
	// Restore a recently deleted gateway. This does not restore the EUI,
	// as that was released when deleting the gateway.
	//
	// Deployment configuration may specify if, and for how long after deletion,
	// entities can be restored.
	Restore(context.Context, *GatewayIdentifiers) (*emptypb.Empty, error)
	// Purge the gateway. This will release both gateway ID and EUI for reuse.
	// The gateway owner is responsible for clearing data from any (external) integrations
	// that may store and expose data by gateway ID.
	Purge(context.Context, *GatewayIdentifiers) (*emptypb.Empty, error)
	mustEmbedUnimplementedGatewayRegistryServer()
}

// UnimplementedGatewayRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayRegistryServer struct {
}

func (UnimplementedGatewayRegistryServer) Create(context.Context, *CreateGatewayRequest) (*Gateway, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGatewayRegistryServer) Get(context.Context, *GetGatewayRequest) (*Gateway, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGatewayRegistryServer) GetIdentifiersForEUI(context.Context, *GetGatewayIdentifiersForEUIRequest) (*GatewayIdentifiers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdentifiersForEUI not implemented")
}
func (UnimplementedGatewayRegistryServer) List(context.Context, *ListGatewaysRequest) (*Gateways, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedGatewayRegistryServer) Update(context.Context, *UpdateGatewayRequest) (*Gateway, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGatewayRegistryServer) Delete(context.Context, *GatewayIdentifiers) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGatewayRegistryServer) Restore(context.Context, *GatewayIdentifiers) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedGatewayRegistryServer) Purge(context.Context, *GatewayIdentifiers) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purge not implemented")
}
func (UnimplementedGatewayRegistryServer) mustEmbedUnimplementedGatewayRegistryServer() {}

// UnsafeGatewayRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayRegistryServer will
// result in compilation errors.
type UnsafeGatewayRegistryServer interface {
	mustEmbedUnimplementedGatewayRegistryServer()
}

func RegisterGatewayRegistryServer(s grpc.ServiceRegistrar, srv GatewayRegistryServer) {
	s.RegisterService(&GatewayRegistry_ServiceDesc, srv)
}

func _GatewayRegistry_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayRegistryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayRegistry/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayRegistryServer).Create(ctx, req.(*CreateGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayRegistryServer).Get(ctx, req.(*GetGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayRegistry_GetIdentifiersForEUI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayIdentifiersForEUIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayRegistryServer).GetIdentifiersForEUI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayRegistry/GetIdentifiersForEUI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayRegistryServer).GetIdentifiersForEUI(ctx, req.(*GetGatewayIdentifiersForEUIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGatewaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayRegistryServer).List(ctx, req.(*ListGatewaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayRegistry_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayRegistryServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayRegistry/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayRegistryServer).Update(ctx, req.(*UpdateGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayRegistryServer).Delete(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayRegistry_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayRegistryServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayRegistry/Restore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayRegistryServer).Restore(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayRegistry_Purge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayRegistryServer).Purge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayRegistry/Purge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayRegistryServer).Purge(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayRegistry_ServiceDesc is the grpc.ServiceDesc for GatewayRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.GatewayRegistry",
	HandlerType: (*GatewayRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GatewayRegistry_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GatewayRegistry_Get_Handler,
		},
		{
			MethodName: "GetIdentifiersForEUI",
			Handler:    _GatewayRegistry_GetIdentifiersForEUI_Handler,
		},
		{
			MethodName: "List",
			Handler:    _GatewayRegistry_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GatewayRegistry_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GatewayRegistry_Delete_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _GatewayRegistry_Restore_Handler,
		},
		{
			MethodName: "Purge",
			Handler:    _GatewayRegistry_Purge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/gateway_services.proto",
}

// GatewayAccessClient is the client API for GatewayAccess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayAccessClient interface {
	// List the rights the caller has on this gateway.
	ListRights(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*Rights, error)
	// Create an API key scoped to this gateway.
	CreateAPIKey(ctx context.Context, in *CreateGatewayAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// List the API keys for this gateway.
	ListAPIKeys(ctx context.Context, in *ListGatewayAPIKeysRequest, opts ...grpc.CallOption) (*APIKeys, error)
	// Get a single API key of this gateway.
	GetAPIKey(ctx context.Context, in *GetGatewayAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// Update the rights of an API key of the gateway.
	// This method can also be used to delete the API key, by giving it no rights.
	// The caller is required to have all assigned or/and removed rights.
	UpdateAPIKey(ctx context.Context, in *UpdateGatewayAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// DeteleAPIKey removes the API key of an gateway.
	DeleteAPIKey(ctx context.Context, in *DeleteGatewayAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// Get the rights of a collaborator (member) of the gateway.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(ctx context.Context, in *GetGatewayCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the gateway.
	// This method can also be used to delete the collaborator, by giving them no rights.
	// The caller is required to have all assigned or/and removed rights.
	SetCollaborator(ctx context.Context, in *SetGatewayCollaboratorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// List the collaborators on this gateway.
	ListCollaborators(ctx context.Context, in *ListGatewayCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error)
}

type gatewayAccessClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayAccessClient(cc grpc.ClientConnInterface) GatewayAccessClient {
	return &gatewayAccessClient{cc}
}

func (c *gatewayAccessClient) ListRights(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*Rights, error) {
	out := new(Rights)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayAccess/ListRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayAccessClient) CreateAPIKey(ctx context.Context, in *CreateGatewayAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayAccess/CreateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayAccessClient) ListAPIKeys(ctx context.Context, in *ListGatewayAPIKeysRequest, opts ...grpc.CallOption) (*APIKeys, error) {
	out := new(APIKeys)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayAccess/ListAPIKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayAccessClient) GetAPIKey(ctx context.Context, in *GetGatewayAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayAccess/GetAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayAccessClient) UpdateAPIKey(ctx context.Context, in *UpdateGatewayAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayAccess/UpdateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayAccessClient) DeleteAPIKey(ctx context.Context, in *DeleteGatewayAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayAccess/DeleteAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayAccessClient) GetCollaborator(ctx context.Context, in *GetGatewayCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error) {
	out := new(GetCollaboratorResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayAccess/GetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayAccessClient) SetCollaborator(ctx context.Context, in *SetGatewayCollaboratorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayAccess/SetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayAccessClient) ListCollaborators(ctx context.Context, in *ListGatewayCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error) {
	out := new(Collaborators)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayAccess/ListCollaborators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayAccessServer is the server API for GatewayAccess service.
// All implementations must embed UnimplementedGatewayAccessServer
// for forward compatibility
type GatewayAccessServer interface {
	// List the rights the caller has on this gateway.
	ListRights(context.Context, *GatewayIdentifiers) (*Rights, error)
	// Create an API key scoped to this gateway.
	CreateAPIKey(context.Context, *CreateGatewayAPIKeyRequest) (*APIKey, error)
	// List the API keys for this gateway.
	ListAPIKeys(context.Context, *ListGatewayAPIKeysRequest) (*APIKeys, error)
	// Get a single API key of this gateway.
	GetAPIKey(context.Context, *GetGatewayAPIKeyRequest) (*APIKey, error)
	// Update the rights of an API key of the gateway.
	// This method can also be used to delete the API key, by giving it no rights.
	// The caller is required to have all assigned or/and removed rights.
	UpdateAPIKey(context.Context, *UpdateGatewayAPIKeyRequest) (*APIKey, error)
	// DeteleAPIKey removes the API key of an gateway.
	DeleteAPIKey(context.Context, *DeleteGatewayAPIKeyRequest) (*APIKey, error)
	// Get the rights of a collaborator (member) of the gateway.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(context.Context, *GetGatewayCollaboratorRequest) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the gateway.
	// This method can also be used to delete the collaborator, by giving them no rights.
	// The caller is required to have all assigned or/and removed rights.
	SetCollaborator(context.Context, *SetGatewayCollaboratorRequest) (*emptypb.Empty, error)
	// List the collaborators on this gateway.
	ListCollaborators(context.Context, *ListGatewayCollaboratorsRequest) (*Collaborators, error)
	mustEmbedUnimplementedGatewayAccessServer()
}

// UnimplementedGatewayAccessServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayAccessServer struct {
}

func (UnimplementedGatewayAccessServer) ListRights(context.Context, *GatewayIdentifiers) (*Rights, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRights not implemented")
}
func (UnimplementedGatewayAccessServer) CreateAPIKey(context.Context, *CreateGatewayAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAPIKey not implemented")
}
func (UnimplementedGatewayAccessServer) ListAPIKeys(context.Context, *ListGatewayAPIKeysRequest) (*APIKeys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAPIKeys not implemented")
}
func (UnimplementedGatewayAccessServer) GetAPIKey(context.Context, *GetGatewayAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIKey not implemented")
}
func (UnimplementedGatewayAccessServer) UpdateAPIKey(context.Context, *UpdateGatewayAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAPIKey not implemented")
}
func (UnimplementedGatewayAccessServer) DeleteAPIKey(context.Context, *DeleteGatewayAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAPIKey not implemented")
}
func (UnimplementedGatewayAccessServer) GetCollaborator(context.Context, *GetGatewayCollaboratorRequest) (*GetCollaboratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollaborator not implemented")
}
func (UnimplementedGatewayAccessServer) SetCollaborator(context.Context, *SetGatewayCollaboratorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCollaborator not implemented")
}
func (UnimplementedGatewayAccessServer) ListCollaborators(context.Context, *ListGatewayCollaboratorsRequest) (*Collaborators, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollaborators not implemented")
}
func (UnimplementedGatewayAccessServer) mustEmbedUnimplementedGatewayAccessServer() {}

// UnsafeGatewayAccessServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayAccessServer will
// result in compilation errors.
type UnsafeGatewayAccessServer interface {
	mustEmbedUnimplementedGatewayAccessServer()
}

func RegisterGatewayAccessServer(s grpc.ServiceRegistrar, srv GatewayAccessServer) {
	s.RegisterService(&GatewayAccess_ServiceDesc, srv)
}

func _GatewayAccess_ListRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAccessServer).ListRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayAccess/ListRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAccessServer).ListRights(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayAccess_CreateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGatewayAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAccessServer).CreateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayAccess/CreateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAccessServer).CreateAPIKey(ctx, req.(*CreateGatewayAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayAccess_ListAPIKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGatewayAPIKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAccessServer).ListAPIKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayAccess/ListAPIKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAccessServer).ListAPIKeys(ctx, req.(*ListGatewayAPIKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayAccess_GetAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAccessServer).GetAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayAccess/GetAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAccessServer).GetAPIKey(ctx, req.(*GetGatewayAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayAccess_UpdateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGatewayAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAccessServer).UpdateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayAccess/UpdateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAccessServer).UpdateAPIKey(ctx, req.(*UpdateGatewayAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayAccess_DeleteAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGatewayAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAccessServer).DeleteAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayAccess/DeleteAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAccessServer).DeleteAPIKey(ctx, req.(*DeleteGatewayAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayAccess_GetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAccessServer).GetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayAccess/GetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAccessServer).GetCollaborator(ctx, req.(*GetGatewayCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayAccess_SetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGatewayCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAccessServer).SetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayAccess/SetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAccessServer).SetCollaborator(ctx, req.(*SetGatewayCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayAccess_ListCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGatewayCollaboratorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAccessServer).ListCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayAccess/ListCollaborators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAccessServer).ListCollaborators(ctx, req.(*ListGatewayCollaboratorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayAccess_ServiceDesc is the grpc.ServiceDesc for GatewayAccess service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayAccess_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.GatewayAccess",
	HandlerType: (*GatewayAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRights",
			Handler:    _GatewayAccess_ListRights_Handler,
		},
		{
			MethodName: "CreateAPIKey",
			Handler:    _GatewayAccess_CreateAPIKey_Handler,
		},
		{
			MethodName: "ListAPIKeys",
			Handler:    _GatewayAccess_ListAPIKeys_Handler,
		},
		{
			MethodName: "GetAPIKey",
			Handler:    _GatewayAccess_GetAPIKey_Handler,
		},
		{
			MethodName: "UpdateAPIKey",
			Handler:    _GatewayAccess_UpdateAPIKey_Handler,
		},
		{
			MethodName: "DeleteAPIKey",
			Handler:    _GatewayAccess_DeleteAPIKey_Handler,
		},
		{
			MethodName: "GetCollaborator",
			Handler:    _GatewayAccess_GetCollaborator_Handler,
		},
		{
			MethodName: "SetCollaborator",
			Handler:    _GatewayAccess_SetCollaborator_Handler,
		},
		{
			MethodName: "ListCollaborators",
			Handler:    _GatewayAccess_ListCollaborators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/gateway_services.proto",
}

// GatewayConfiguratorClient is the client API for GatewayConfigurator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayConfiguratorClient interface {
	PullConfiguration(ctx context.Context, in *PullGatewayConfigurationRequest, opts ...grpc.CallOption) (GatewayConfigurator_PullConfigurationClient, error)
}

type gatewayConfiguratorClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayConfiguratorClient(cc grpc.ClientConnInterface) GatewayConfiguratorClient {
	return &gatewayConfiguratorClient{cc}
}

func (c *gatewayConfiguratorClient) PullConfiguration(ctx context.Context, in *PullGatewayConfigurationRequest, opts ...grpc.CallOption) (GatewayConfigurator_PullConfigurationClient, error) {
	stream, err := c.cc.NewStream(ctx, &GatewayConfigurator_ServiceDesc.Streams[0], "/ttn.lorawan.v3.GatewayConfigurator/PullConfiguration", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewayConfiguratorPullConfigurationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GatewayConfigurator_PullConfigurationClient interface {
	Recv() (*Gateway, error)
	grpc.ClientStream
}

type gatewayConfiguratorPullConfigurationClient struct {
	grpc.ClientStream
}

func (x *gatewayConfiguratorPullConfigurationClient) Recv() (*Gateway, error) {
	m := new(Gateway)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GatewayConfiguratorServer is the server API for GatewayConfigurator service.
// All implementations must embed UnimplementedGatewayConfiguratorServer
// for forward compatibility
type GatewayConfiguratorServer interface {
	PullConfiguration(*PullGatewayConfigurationRequest, GatewayConfigurator_PullConfigurationServer) error
	mustEmbedUnimplementedGatewayConfiguratorServer()
}

// UnimplementedGatewayConfiguratorServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayConfiguratorServer struct {
}

func (UnimplementedGatewayConfiguratorServer) PullConfiguration(*PullGatewayConfigurationRequest, GatewayConfigurator_PullConfigurationServer) error {
	return status.Errorf(codes.Unimplemented, "method PullConfiguration not implemented")
}
func (UnimplementedGatewayConfiguratorServer) mustEmbedUnimplementedGatewayConfiguratorServer() {}

// UnsafeGatewayConfiguratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayConfiguratorServer will
// result in compilation errors.
type UnsafeGatewayConfiguratorServer interface {
	mustEmbedUnimplementedGatewayConfiguratorServer()
}

func RegisterGatewayConfiguratorServer(s grpc.ServiceRegistrar, srv GatewayConfiguratorServer) {
	s.RegisterService(&GatewayConfigurator_ServiceDesc, srv)
}

func _GatewayConfigurator_PullConfiguration_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullGatewayConfigurationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GatewayConfiguratorServer).PullConfiguration(m, &gatewayConfiguratorPullConfigurationServer{stream})
}

type GatewayConfigurator_PullConfigurationServer interface {
	Send(*Gateway) error
	grpc.ServerStream
}

type gatewayConfiguratorPullConfigurationServer struct {
	grpc.ServerStream
}

func (x *gatewayConfiguratorPullConfigurationServer) Send(m *Gateway) error {
	return x.ServerStream.SendMsg(m)
}

// GatewayConfigurator_ServiceDesc is the grpc.ServiceDesc for GatewayConfigurator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayConfigurator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.GatewayConfigurator",
	HandlerType: (*GatewayConfiguratorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullConfiguration",
			Handler:       _GatewayConfigurator_PullConfiguration_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lorawan-stack/api/gateway_services.proto",
}
