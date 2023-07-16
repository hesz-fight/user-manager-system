package simrpc

import (
	"encoding/json"
	"fmt"
	"learn/user-manager-system/rpcsvr/global"
	"learn/user-manager-system/rpcsvr/pkg/errcode"
	"net"
	"reflect"
)

type serverObj struct {
	Trans *transfer
}

type Request struct {
	ServiceMethod string
	MethodObj     *methodObj
}

type methodObj struct {
	Rcvr     reflect.Value
	Method   reflect.Method
	ParamTyp reflect.Type
	RspTyp   reflect.Type
}

var RequestMap = make(map[string]*Request)

func getServiceMethod(service string, method string) string {
	return fmt.Sprintf("%s.%s", service, method)
}

func RegisterRequest(svr interface{}, serviceName string, methodName string) error {
	serviceMethod := getServiceMethod(serviceName, methodName)
	if _, ok := RequestMap[serviceMethod]; ok {
		return ErrorMethodHasExist
	}
	method, exist := reflect.TypeOf(svr).MethodByName(methodName)
	if !exist {
		global.LogLogger.Errorf("method %s doen not exist.", methodName)
		return ErrorMehodNotFound
	}
	request := &Request{
		ServiceMethod: serviceMethod,
		MethodObj: &methodObj{
			Rcvr:     reflect.ValueOf(svr),
			Method:   method,
			ParamTyp: method.Type.In(1),
			RspTyp:   method.Type.In(2),
		},
	}
	RequestMap[serviceMethod] = request

	return nil
}

func ListenAndServe(address string) error {
	lst, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	for {
		conn, err := lst.Accept()
		if err != nil {
			return err
		}
		server := serverObj{
			Trans: &transfer{
				Buf:  make([]byte, MaxBufSize),
				Conn: conn,
			},
		}

		go server.receiveRequest()
	}
}

func (s *serverObj) receiveRequest() {
	defer s.Trans.Conn.Close()
	for {
		serviceMethod, data, err := s.Trans.Read()
		if err != nil {
			global.LogLogger.Errorf("handleRequest error: %s", err)
			return
		}
		if err := s.handleRequest(serviceMethod, data); err != nil {
			global.LogLogger.Errorf("handleRequest error: %s", err)
			return
		}
	}
}

func (s *serverObj) handleRequest(serviceMethod string, data []byte) error {
	request, ok := RequestMap[serviceMethod]
	if !ok {
		return errcode.InvalidParams
	}
	method := request.MethodObj
	f := method.Method.Func
	req := reflect.New(method.ParamTyp).Interface()
	if err := json.Unmarshal(data, req); err != nil {
		return err
	}
	rsp := reflect.New(method.RspTyp)
	in := []reflect.Value{method.Rcvr, reflect.ValueOf(req), rsp}
	out := f.Call(in)
	if err := out[0].Interface().(error); err != nil {
		return err
	}
	// response data
	if err := s.responseRequest(rsp.Interface()); err != nil {
		return err
	}

	return nil
}

func (s *serverObj) responseRequest(rsp interface{}) error {
	data, err := json.Marshal(rsp)
	if err != nil {
		return err
	}
	if err := s.Trans.Write("", data); err != nil {
		return err
	}
	return nil
}
