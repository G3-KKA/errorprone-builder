package main

import "context"

// Step 1. Declare struct
type doneHTTPServer struct {
	IdleTime  string
	ReadTime  string
	WriteTime string
	MaxHeader int

	port    int
	timeout string
	host    string
	ctx     context.Context
}

// Step 2. add //go:generate:builder
// Step 3. go generate ./file.go:HTTPServer
// Will generate phase per every private field
// Result:
func BuildHTTPServer() phase0_HASH23b4bcj3 {
	return phase0_HASH23b4bcj3{}
}

type phase0_HASH23b4bcj3 struct {
}

func (phase phase0_HASH23b4bcj3) Set_ctx(ctx context.Context) phase1_HASH23b4bcj3 {
	return phase1_HASH23b4bcj3{
		ctx: ctx,
	}
}

// =>
type phase1_HASH23b4bcj3 struct {
	ctx context.Context
}

func (phase phase1_HASH23b4bcj3) Set_host(host string) phase2_HASH23b4bcj3 {
	return phase2_HASH23b4bcj3{
		host: host,
		ctx:  phase.ctx,
	}
}

// =>
type phase2_HASH23b4bcj3 struct {
	host string
	ctx  context.Context
}

func (phase phase2_HASH23b4bcj3) Set_timeout(timeout string) phase3_HASH23b4bcj3 {
	return phase3_HASH23b4bcj3{
		timeout: timeout,
		host:    phase.host,
		ctx:     phase.ctx,
	}
}

// =>
type phase3_HASH23b4bcj3 struct {
	timeout string
	host    string
	ctx     context.Context
}

func (phase phase3_HASH23b4bcj3) Set_port(port int) phase4_HASH23b4bcj3 {
	return phase4_HASH23b4bcj3{
		port:    port,
		timeout: phase.timeout,
		host:    phase.host,
		ctx:     phase.ctx,
	}
}

// =>
type phase4_HASH23b4bcj3 struct {
	port    int
	timeout string
	host    string
	ctx     context.Context
}

func (phase phase4_HASH23b4bcj3) Assemble() doneHTTPServer {
	return doneHTTPServer{
		port:    phase.port,
		timeout: phase.timeout,
		host:    phase.host,
		ctx:     phase.ctx,
	}
}

// code :
// Step 4. use it in code
func main() {
	// ...
	ctx := context.Background()
	host := "localhost"
	httpServer := BuildHTTPServer().
		Set_ctx(ctx).
		Set_host(host).
		Set_timeout("5s").
		Set_port(8080).
		Assemble()
	// ...
}

// BONUS  TODO  Step 5.
// run go generate ./file.go:HTTPServer?refine=true
// and swap
// func BuildHTTPServer() phase0_HASH23b4bcj3 {
//	return phase0_HASH23b4bcj3{}
//	}
// TO  =>
//	func BuildHTTPServer() doneHTTPServer {
//	return // ... LOGIC
//	}

// This should drasticly reduce copy-overhead for the generated code
