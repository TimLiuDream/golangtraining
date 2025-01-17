// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package model

import (
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
)

// weaver.InstanceOf checks.

// weaver.Router checks.

// Local stub implementations.

// Client stub implementations.

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.21.2 (codegen
version v0.20.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

// Reflect stub implementations.

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*Golang)(nil)

type __is_Golang[T ~struct {
	weaver.AutoMarshal
	Channel   string
	Goroutine string
}] struct{}

var _ __is_Golang[Golang]

func (x *Golang) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Golang.WeaverMarshal: nil receiver"))
	}
	enc.String(x.Channel)
	enc.String(x.Goroutine)
}

func (x *Golang) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Golang.WeaverUnmarshal: nil receiver"))
	}
	x.Channel = dec.String()
	x.Goroutine = dec.String()
}
