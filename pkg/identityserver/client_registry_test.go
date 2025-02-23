// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package identityserver

import (
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/storetest"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestClientsPermissionDenied(t *testing.T) {
	p := &storetest.Population{}
	usr1 := p.NewUser()
	cli1 := p.NewClient(usr1.GetOrganizationOrUserIdentifiers())

	t.Parallel()
	a, ctx := test.New(t)

	testWithIdentityServer(t, func(_ *IdentityServer, cc *grpc.ClientConn) {
		reg := ttnpb.NewClientRegistryClient(cc)

		_, err := reg.Create(ctx, &ttnpb.CreateClientRequest{
			Client: &ttnpb.Client{
				Ids: &ttnpb.ClientIdentifiers{ClientId: "foo-cli"},
			},
			Collaborator: usr1.GetOrganizationOrUserIdentifiers(),
		})
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		_, err = reg.Get(ctx, &ttnpb.GetClientRequest{
			ClientIds: cli1.GetIds(),
			FieldMask: ttnpb.FieldMask("name"),
		})
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsUnauthenticated(err), should.BeTrue)
		}

		listRes, err := reg.List(ctx, &ttnpb.ListClientsRequest{
			FieldMask: ttnpb.FieldMask("name"),
		})
		a.So(err, should.BeNil)
		if a.So(listRes, should.NotBeNil) {
			a.So(listRes.Clients, should.BeEmpty)
		}

		_, err = reg.List(ctx, &ttnpb.ListClientsRequest{
			Collaborator: usr1.GetOrganizationOrUserIdentifiers(),
			FieldMask:    ttnpb.FieldMask("name"),
		})
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		_, err = reg.Update(ctx, &ttnpb.UpdateClientRequest{
			Client: &ttnpb.Client{
				Ids:  cli1.GetIds(),
				Name: "Updated Name",
			},
			FieldMask: ttnpb.FieldMask("name"),
		})
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		_, err = reg.Delete(ctx, cli1.GetIds())
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}
	}, withPrivateTestDatabase(p))
}

func TestClientsCRUD(t *testing.T) {
	p := &storetest.Population{}

	adminUsr := p.NewUser()
	adminUsr.Admin = true
	adminKey, _ := p.NewAPIKey(adminUsr.GetEntityIdentifiers(), ttnpb.Right_RIGHT_ALL)
	adminCreds := rpcCreds(adminKey)

	usr1 := p.NewUser()
	for i := 0; i < 5; i++ {
		p.NewClient(usr1.GetOrganizationOrUserIdentifiers())
	}

	usr2 := p.NewUser()
	for i := 0; i < 5; i++ {
		p.NewClient(usr2.GetOrganizationOrUserIdentifiers())
	}

	key, _ := p.NewAPIKey(usr1.GetEntityIdentifiers(), ttnpb.Right_RIGHT_ALL)
	creds := rpcCreds(key)
	keyWithoutRights, _ := p.NewAPIKey(usr1.GetEntityIdentifiers())
	credsWithoutRights := rpcCreds(keyWithoutRights)

	t.Parallel()
	a, ctx := test.New(t)

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		reg := ttnpb.NewClientRegistryClient(cc)

		is.config.UserRights.CreateClients = false

		_, err := reg.Create(ctx, &ttnpb.CreateClientRequest{
			Client: &ttnpb.Client{
				Ids:  &ttnpb.ClientIdentifiers{ClientId: "foo"},
				Name: "Foo Client",
			},
			Collaborator: usr1.GetOrganizationOrUserIdentifiers(),
		}, creds)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		is.config.UserRights.CreateClients = true

		created, err := reg.Create(ctx, &ttnpb.CreateClientRequest{
			Client: &ttnpb.Client{
				Ids:  &ttnpb.ClientIdentifiers{ClientId: "foo"},
				Name: "Foo Client",
			},
			Collaborator: usr1.GetOrganizationOrUserIdentifiers(),
		}, creds)
		if a.So(err, should.BeNil) && a.So(created, should.NotBeNil) {
			a.So(created.Name, should.Equal, "Foo Client")
		}

		got, err := reg.Get(ctx, &ttnpb.GetClientRequest{
			ClientIds: created.GetIds(),
			FieldMask: ttnpb.FieldMask("name"),
		}, creds)
		if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) {
			a.So(got.Name, should.Equal, created.Name)
		}

		got, err = reg.Get(ctx, &ttnpb.GetClientRequest{
			ClientIds: created.GetIds(),
			FieldMask: ttnpb.FieldMask("ids"),
		}, credsWithoutRights)
		a.So(err, should.BeNil)

		got, err = reg.Get(ctx, &ttnpb.GetClientRequest{
			ClientIds: created.GetIds(),
			FieldMask: ttnpb.FieldMask("attributes"),
		}, credsWithoutRights)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		updated, err := reg.Update(ctx, &ttnpb.UpdateClientRequest{
			Client: &ttnpb.Client{
				Ids:  created.GetIds(),
				Name: "Updated Name",
			},
			FieldMask: ttnpb.FieldMask("name"),
		}, creds)
		if a.So(err, should.BeNil) && a.So(updated, should.NotBeNil) {
			a.So(updated.Name, should.Equal, "Updated Name")
		}

		updated, err = reg.Update(ctx, &ttnpb.UpdateClientRequest{
			Client: &ttnpb.Client{
				Ids:              created.GetIds(),
				State:            ttnpb.State_STATE_FLAGGED,
				StateDescription: "something is wrong",
			},
			FieldMask: ttnpb.FieldMask("state", "state_description"),
		}, adminCreds)
		if a.So(err, should.BeNil) && a.So(updated, should.NotBeNil) {
			a.So(updated.State, should.Equal, ttnpb.State_STATE_FLAGGED)
			a.So(updated.StateDescription, should.Equal, "something is wrong")
		}

		updated, err = reg.Update(ctx, &ttnpb.UpdateClientRequest{
			Client: &ttnpb.Client{
				Ids:   created.GetIds(),
				State: ttnpb.State_STATE_APPROVED,
			},
			FieldMask: ttnpb.FieldMask("state"),
		}, adminCreds)
		if a.So(err, should.BeNil) && a.So(updated, should.NotBeNil) {
			a.So(updated.State, should.Equal, ttnpb.State_STATE_APPROVED)
		}

		got, err = reg.Get(ctx, &ttnpb.GetClientRequest{
			ClientIds: created.GetIds(),
			FieldMask: ttnpb.FieldMask("state", "state_description"),
		}, creds)
		if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) {
			a.So(got.State, should.Equal, ttnpb.State_STATE_APPROVED)
			a.So(got.StateDescription, should.Equal, "")
		}

		for _, collaborator := range []*ttnpb.OrganizationOrUserIdentifiers{nil, usr1.GetOrganizationOrUserIdentifiers()} {
			list, err := reg.List(ctx, &ttnpb.ListClientsRequest{
				FieldMask:    ttnpb.FieldMask("name"),
				Collaborator: collaborator,
			}, creds)
			if a.So(err, should.BeNil) && a.So(list, should.NotBeNil) && a.So(list.Clients, should.HaveLength, 6) {
				var found bool
				for _, item := range list.Clients {
					if item.GetIds().GetClientId() == created.GetIds().GetClientId() {
						found = true
						a.So(item.Name, should.Equal, "Updated Name")
					}
				}
				a.So(found, should.BeTrue)
			}
		}

		_, err = reg.Delete(ctx, created.GetIds(), creds)
		a.So(err, should.BeNil)

		_, err = reg.Purge(ctx, created.GetIds(), creds)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		_, err = reg.Purge(ctx, created.GetIds(), adminCreds)
		a.So(err, should.BeNil)
	}, withPrivateTestDatabase(p))
}

func TestClientsPagination(t *testing.T) {
	p := &storetest.Population{}

	usr1 := p.NewUser()
	for i := 0; i < 3; i++ {
		p.NewClient(usr1.GetOrganizationOrUserIdentifiers())
	}

	key, _ := p.NewAPIKey(usr1.GetEntityIdentifiers(), ttnpb.Right_RIGHT_ALL)
	creds := rpcCreds(key)

	t.Parallel()
	a, ctx := test.New(t)

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		reg := ttnpb.NewClientRegistryClient(cc)

		var md metadata.MD

		list, err := reg.List(ctx, &ttnpb.ListClientsRequest{
			FieldMask:    ttnpb.FieldMask("name"),
			Collaborator: usr1.GetOrganizationOrUserIdentifiers(),
			Limit:        2,
			Page:         1,
		}, creds, grpc.Header(&md))
		if a.So(err, should.BeNil) && a.So(list, should.NotBeNil) {
			a.So(list.Clients, should.HaveLength, 2)
			a.So(md.Get("x-total-count"), should.Resemble, []string{"3"})
		}

		list, err = reg.List(ctx, &ttnpb.ListClientsRequest{
			FieldMask:    ttnpb.FieldMask("name"),
			Collaborator: usr1.GetOrganizationOrUserIdentifiers(),
			Limit:        2,
			Page:         2,
		}, creds)
		if a.So(err, should.BeNil) && a.So(list, should.NotBeNil) {
			a.So(list.Clients, should.HaveLength, 1)
		}

		list, err = reg.List(ctx, &ttnpb.ListClientsRequest{
			FieldMask:    ttnpb.FieldMask("name"),
			Collaborator: usr1.GetOrganizationOrUserIdentifiers(),
			Limit:        2,
			Page:         3,
		}, creds)
		if a.So(err, should.BeNil) && a.So(list, should.NotBeNil) {
			a.So(list.Clients, should.BeEmpty)
		}
	}, withPrivateTestDatabase(p))
}
