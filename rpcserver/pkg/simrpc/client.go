package simrpc

import (
	"encoding/json"
	"errors"
	"net"
)

type ClientPool struct {
	pool chan *clientObj
}

type clientObj struct {
	Trans *transfer
}

func NewClientPool(connections int, network string, adress string) (*ClientPool, error) {
	clientPool := &ClientPool{
		pool: make(chan *clientObj),
	}

	for i := 0; i < connections; i++ {
		conn, err := net.Dial(network, adress)
		if err != nil {
			continue
		}
		client := &clientObj{
			Trans: &transfer{
				Buf:  make([]byte, MaxBufSize),
				Conn: conn,
			},
		}
		clientPool.pool <- client
	}

	return clientPool, nil
}

func (cp *ClientPool) Call(paramType int, pspType int, req interface{}, rsp interface{}) error {
	client := cp.getClient()
	defer cp.releaseClient(client)
	if err := client.sendRequest(paramType, req); err != nil {
		return err
	}
	if err := client.receiveRequest(pspType, rsp); err != nil {
		return err
	}

	return nil
}

func (cp *ClientPool) getClient() *clientObj {
	select {
	case clientObj := <-cp.pool:
		return clientObj
	}
}

func (cp *ClientPool) releaseClient(client *clientObj) {
	select {
	case cp.pool <- client:
		return
	}
}

func (c *clientObj) sendRequest(paramType int, req interface{}) error {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	if err := c.Trans.Write(paramType, data); err != nil {
		return err
	}
	return nil
}

func (c *clientObj) receiveRequest(rspType int, rsp interface{}) error {
	rcvType, data, err := c.Trans.Read()
	if err != nil {
		return err
	}
	if rspType != rcvType {
		return errors.New("wrong package type")
	}
	if err := json.Unmarshal(data, rsp); err != nil {
		return err
	}

	return nil
}
