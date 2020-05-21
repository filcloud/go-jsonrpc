package jsonrpc

import (
	"context"
	"io"
	"reflect"
)

// Export some things to lotus/rpc.

type Handlers = handlers

func (h Handlers) Handle(ctx context.Context, req request, w func(func(io.Writer)), rpcError rpcErrFunc, done func(keepCtx bool), chOut chanOut) {
	h.handle(ctx, req, w, rpcError, done, chOut)
}
func (h Handlers) HandleReader(ctx context.Context, r io.Reader, w io.Writer, rpcError rpcErrFunc) {
	h.handleReader(ctx, r, w, rpcError)
}
func (h handlers) Register(namespace string, r interface{}) {
	h.register(namespace, r)
}

type RPCRequest = request
type RPCHandler = rpcHandler
type ClientResponse = clientResponse
type Param = param

var RPCError = rpcError
var RPCParseError = rpcParseError
var ErrorType = errorType
var ContextType = contextType
var ProcessFuncOut = processFuncOut

func NewParam(v reflect.Value) Param {
	return Param{
		v: v,
	}
}
