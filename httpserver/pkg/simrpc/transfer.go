package simrpc

import (
	"encoding/binary"
	"encoding/json"
	"net"
)

type header struct {
	ServiceMethod string `json:"service_method"`
	Length        int    `json:"length"`
}

const MaxBufSize = 4096

type transfer struct {
	Buf  []byte
	Conn net.Conn
}

func (t *transfer) Write(serviceMethod string, data []byte) error {
	headerObj := &header{
		ServiceMethod: serviceMethod,
		Length:        len(data),
	}
	headerByte, err := json.Marshal(headerObj)
	if err != nil {
		return err
	}

	lenHeader := make([]byte, 4)
	binary.BigEndian.PutUint32(lenHeader, uint32(len(headerByte)))
	if _, err = t.Conn.Write(lenHeader); err != nil {
		return err
	}
	transData := make([]byte, 0)
	transData = append(transData, headerByte...)
	transData = append(transData, data...)
	if _, err = t.Conn.Write(transData); err != nil {
		return err
	}

	return nil
}

func (t *transfer) Read() (string, []byte, error) {
	if _, err := t.Conn.Read(t.Buf[:4]); err != nil {
		return "", []byte{}, err
	}
	lenHeader := binary.BigEndian.Uint32(t.Buf[:4])
	if _, err := t.Conn.Read(t.Buf[:lenHeader]); err != nil {
		return "", []byte{}, err
	}
	h := &header{}
	if err := json.Unmarshal(t.Buf[:lenHeader], h); err != nil {
		return "", []byte{}, err
	}
	if _, err := t.Conn.Read(t.Buf[:h.Length]); err != nil {
		return "", []byte{}, err
	}

	return h.ServiceMethod, t.Buf[:h.Length], nil
}
