// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
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

//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"text/template"

	_ "github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

func init() {
	log.SetOutput(os.Stderr)
	log.SetFlags(0)
}

var bytesTyp = reflect.TypeOf([]byte{})

func typeString(typ reflect.Type) string {
	if typ.Kind() == reflect.Ptr {
		return "*" + typeString(typ.Elem())
	}
	if typ.PkgPath() == "github.com/gogo/protobuf/types" {
		return "pbtypes." + typ.Name()
	}
	return typ.String()
}

func main() {
	buf := &bytes.Buffer{}

	type MethodContext struct {
		FieldName     string
		ParameterName string
		ParameterType string
	}
	if err := template.Must(template.New("struct").Funcs(template.FuncMap{
		"typeOf":     reflect.TypeOf,
		"typeString": typeString,
		"fields": func(typ reflect.Type) (fs []MethodContext) {
			numFields := typ.NumField()
			for i := 0; i < numFields; i++ {
				f := typ.Field(i)
				if f.Name == "XXX_NoUnkeyedLiteral" || f.Name == "XXX_sizecache" {
					// internal protobuf fields
					continue
				}

				var pName, pType string
				switch {
				case f.Type == bytesTyp:
					pName = "v"
					pType = "[]byte"
				case f.Type.Kind() == reflect.Slice:
					pName = "vs"
					pType = "..." + typeString(f.Type.Elem())
				default:
					pName = "v"
					pType = typeString(f.Type)
				}
				fs = append(fs, MethodContext{
					FieldName:     f.Name,
					ParameterName: pName,
					ParameterType: pType,
				})
			}
			return fs
		},
	}).Parse(
		`// Code generated by generate_constructors.go. DO NOT EDIT.

package test

import (
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	pbtypes "github.com/gogo/protobuf/types"
)
{{ range . }}
{{ with $type := typeOf . -}}
{{ with $typeString := typeString $type -}}
{{ with $optionType := printf "%sOption" $type.Name -}}
{{ with $optionsType := printf "%sOptionNamespace" $type.Name -}}
type (
	// {{ $optionType }} transforms {{ $typeString }} and returns it.
	// Implemetations must be pure functions with no side-effects.
	{{ $optionType }} func({{ $typeString }}) {{ $typeString }}

	// {{ $optionsType }} represents the namespace, on which various {{ $optionType }} are defined.
	{{ $optionsType }} struct{}
)
{{ range fields $type }}
// With{{ .FieldName }} returns a {{ $optionType }}, which returns a copy of {{ $typeString }} with {{ .FieldName }} set to {{ .ParameterName }}.
func ({{ $optionsType }}) With{{ .FieldName }}({{ .ParameterName }} {{ .ParameterType }}) {{ $optionType }} {
	return func(x {{ $typeString }}) {{ $typeString }} {
		x.{{ .FieldName }} = {{ .ParameterName }}
		return x
	}
}
{{ end }}
// Compose returns a functional composition of opts as a singular {{ $optionType }}.
func ({{ $optionsType }}) Compose(opts ...{{ $optionType }}) {{ $optionType }} {
	return func(x {{ $typeString }}) {{ $typeString }} {
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// Compose returns a functional composition of f and opts as a singular {{ $optionType }}.
func (f {{ $optionType }}) Compose(opts ...{{ $optionType }}) {{ $optionType }} {
	return func(x {{ $typeString }}) {{ $typeString }} {
		x = f(x)
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

{{ with $optionsVar := printf "%sOptions" $type.Name -}}
// {{ $optionsVar }} is namespace containing {{ $typeString }} options.
var {{ $optionsVar }} {{ $optionsType }}

// Make{{ $type.Name }} constructs a new {{ $typeString }}.
func Make{{ $type.Name }}(opts ...{{ $optionType }}) *{{ $typeString }} {
	v := {{ $optionsVar }}.Compose(opts...)(base{{ $type.Name }})
	return &v
}
{{ end -}}
{{ end -}}
{{ end -}}
{{ end -}}
{{ end -}}
{{ end }}`)).Execute(buf, []interface{}{
		ttnpb.RootKeys{},
		ttnpb.SessionKeys{},
		ttnpb.Session{},
		ttnpb.MACState{},
		ttnpb.EndDeviceIdentifiers{},
		ttnpb.EndDevice{},
	}); err != nil {
		log.Fatalf("Failed to execute template: %s", err)
	}
	b, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("Failed to format source: %s", err)
	}
	if err := ioutil.WriteFile("constructors_generated.go", b, 0o644); err != nil {
		log.Fatalf("Failed to write output: %s", err)
	}
}
