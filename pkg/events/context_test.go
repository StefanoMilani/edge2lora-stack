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

package events_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

type testContextKeyType string

var testContextKey = testContextKeyType("ctx-test")

type testContextMarshaler struct{}

func (testContextMarshaler) MarshalContext(ctx context.Context) []byte {
	if val, ok := ctx.Value(testContextKey).(string); ok {
		return []byte(val)
	}
	return nil
}

func (testContextMarshaler) UnmarshalContext(ctx context.Context, b []byte) (context.Context, error) {
	return context.WithValue(ctx, testContextKey, string(b)), nil
}

func TestContextMarshaler(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	var m testContextMarshaler
	events.RegisterContextMarshaler("test", m)

	ctx := context.WithValue(context.Background(), testContextKey, "foo")

	evt := events.New(ctx, "test", "test")

	b, err := json.Marshal(evt)
	a.So(err, should.BeNil)

	unmarshaled, err := events.UnmarshalJSON(b)
	a.So(err, should.BeNil)
	a.So(unmarshaled, should.Resemble, evt)

	val, ok := unmarshaled.Context().Value(testContextKey).(string)
	if a.So(ok, should.BeTrue) {
		a.So(val, should.Equal, "foo")
	}
}
