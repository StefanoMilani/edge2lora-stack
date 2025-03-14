// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

package errors_test

import (
	"testing"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestAttributesError(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	errInvalidFoo := errors.DefineInvalidArgument("test_attributes_invalid_foo", "Invalid Foo: {foo}", "foo")

	a.So(errInvalidFoo.Attributes(), should.BeEmpty)

	a.So(func() { errInvalidFoo.WithAttributes("only_key") }, should.Panic) //nolint:errcheck

	err1 := errInvalidFoo.WithAttributes("foo", "bar")
	err2 := err1.WithAttributes("bar", "baz")

	a.So(err1, should.HaveSameErrorDefinitionAs, errInvalidFoo)
	a.So(errors.Attributes(err1), should.Resemble, map[string]any{"foo": "bar"})
	a.So(errors.PublicAttributes(err1), should.Resemble, map[string]any{"foo": "bar"})

	a.So(err2, should.HaveSameErrorDefinitionAs, err1)
	a.So(err2, should.HaveSameErrorDefinitionAs, errInvalidFoo)
	a.So(errors.Attributes(err2), should.Resemble, map[string]any{"foo": "bar", "bar": "baz"})
	a.So(errors.PublicAttributes(err2), should.Resemble, map[string]any{"foo": "bar"})
}

func TestAttributes(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		Name   string
		V      any
		Expect any
	}{
		{"int", int(42), int(42)},
		{"float64", float64(42), float64(42)},
		{"string", "foo", "foo"},
		{"nil", nil, "<nil>"},
		{"complex64", complex(42, 42), "(42+42i)"},
		{"uint32", uint32(42), "42"},
		{"array", [5]int{1, 2, 3}, "[1 2 3 0 0]"},
	}

	for _, tc := range tcs {
		tc := tc // shadow range variable.
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			assertions.New(t).So(errors.Supported(tc.V), should.Equal, tc.Expect)
		})
	}
}
