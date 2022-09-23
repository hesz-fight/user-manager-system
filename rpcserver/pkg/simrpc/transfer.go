package simrpc

import (
	"encoding/binary"
	"encoding/json"
	"net"
)

type header struct {
	ParamType int `json:"param_type"`
	Length    int `json:"length"`
}

const MaxBufSize = 4096

type transfer struct {
	Buf  []byte
	Conn net.Conn
}

func (t *transfer) Write(ParamType int, data []byte) error {
	headerObj := &header{
		ParamType: ParamType,
		Length:    len(data),
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

func (t *transfer) Read() (int, []byte, error) {
	if _, err := t.Conn.Read(t.Buf[:4]); err != nil {
		return 0, []byte{}, err
	}
	lenHeader := binary.BigEndian.Uint32(t.Buf[:4])
	if _, err := t.Conn.Read(t.Buf[:lenHeader]); err != nil {
		return 0, []byte{}, err
	}
	h := &header{}
	if err := json.Unmarshal(t.Buf[:lenHeader], h); err != nil {
		return 0, []byte{}, err
	}
	if _, err := t.Conn.Read(t.Buf[:h.Length]); err != nil {
		return 0, []byte{}, err
	}

	return h.ParamType, t.Buf[:h.Length], nil
}
