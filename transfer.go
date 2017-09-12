package yapool

import "encoding/json"

type Msg struct {
	Level Level         `json:"level"`
	Msg   []interface{} `json:"msg"`
}

func ConvertMsg(level Level, msg ...interface{}) *Msg {
	return &Msg{
		Level: level,
		Msg:   msg,
	}
}

func Decode(data []byte) (*Msg, error) {
	msg := &Msg{}
	err := json.Unmarshal(data, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (m *Msg) Encode() ([]byte, error) {
	byt, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return byt, nil
}
