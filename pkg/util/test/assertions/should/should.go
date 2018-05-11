// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
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

package should

import "go.thethings.network/lorawan-stack/pkg/util/test/assertions"

var (
	// Describe asserts that the errors.Error was generated by a given errors.ErrDescriptor.
	Describe = assertions.ShouldDescribe
	// NotDescribe asserts that the errors.Error was not generated by a given errors.ErrDescriptor.
	NotDescribe = assertions.ShouldNotDescribe
)
