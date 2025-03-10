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

package log_test

import (
	"context"
	"testing"

	"github.com/smartystreets/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestContext(t *testing.T) {
	a := assertions.New(t)
	ctx := context.Background()

	logger := NewLogger(NoopHandler)

	a.So(FromContext(ctx), should.NotBeNil)
	a.So(FromContext(ctx), should.Equal, Noop)

	ctx = NewContext(ctx, logger)

	a.So(FromContext(ctx), should.Equal, logger)

	t.Run("NewContextWithField", func(t *testing.T) {
		a := assertions.New(t)
		withKV := FromContext(NewContextWithField(ctx, "key", "value")).(Entry)
		fields := withKV.Fields().Fields()
		v, ok := fields["key"]
		a.So(ok, should.BeTrue)
		a.So(v, should.Equal, "value")
	})

	t.Run("NewContextWithFields", func(t *testing.T) {
		a := assertions.New(t)
		withKV := FromContext(NewContextWithFields(ctx, Fields("key", "value"))).(Entry)
		fields := withKV.Fields().Fields()
		v, ok := fields["key"]
		a.So(ok, should.BeTrue)
		a.So(v, should.Equal, "value")
	})
}
