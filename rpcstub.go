package jsonrpc

import (
	"context"
	"io"
	"reflect"
)

// Export some things to lotus/rpc.

func (s *RPCServer) Handle(ctx context.Context, req request, w func(func(io.Writer)), rpcError rpcErrFunc, done func(keepCtx bool), chOut chanOut) {
	s.handle(ctx, req, w, rpcError, done, chOut)
}
func (s *RPCServer) HandleReader(ctx context.Context, r io.Reader, w io.Writer, rpcError rpcErrFunc) {
	s.handleReader(ctx, r, w, rpcError)
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

func DefaultServerConfig() ServerConfig {
	return defaultServerConfig()
}

func (c *ServerConfig) ParamDecoders() map[reflect.Type]ParamDecoder {
	return c.paramDecoders
}

func NewServerExt(methods map[string]rpcHandler, paramDecoders map[reflect.Type]ParamDecoder) *RPCServer {
	return &RPCServer{
		methods:       methods,
		paramDecoders: paramDecoders,
	}
}
