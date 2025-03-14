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
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	_ "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestHTTP(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	errDef := errors.DefineInvalidArgument("test_http_conversion_err_def", "HTTP Conversion Error", "foo")
	a.So(errors.FromGRPCStatus(errDef.GRPCStatus()).Definition, should.EqualErrorOrDefinition, errDef)

	errHello := errDef.WithAttributes("foo", "bar", "baz", "qux")
	errHelloExpected := errDef.WithAttributes("foo", "bar")

	handler := func(w http.ResponseWriter, _ *http.Request) {
		err := errors.ToHTTP(errHello, w)
		a.So(err, should.BeNil)
	}

	req := httptest.NewRequest("GET", "http://example.com/err", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	a.So(w.Result().StatusCode, should.Equal, http.StatusBadRequest)
	a.So(errors.FromHTTP(resp), should.EqualErrorOrDefinition, errHelloExpected)
}
