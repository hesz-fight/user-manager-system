package simrpc

import "reflect"

type serverObj struct {
	Trans *transfer
}

type Request struct {
	Method    reflect.Value
	ParamType int
	RspType   int
}

func RegisterRequest(sr interface{}, methodName string, paramType int, rspType int) {

}

func ListenAndServe(address string) error {
	return nil
}

func (s *serverObj) receiveRequest() {

}

func (server *serverObj) handleRequest(paramType int, param interface{}) error {
	return nil
}

func (server *serverObj) responseRequest(paramType int, response interface{}) error {
	return nil
}
