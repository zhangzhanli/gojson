package gojson

func Unmarshal(data []byte) (interface{}, error) {
	dec := NewDecoder()
	dec.SetData(data)
	defer dec.Release()
	return dec.DecodeValue()
}

func Marshal(v interface{}) ([]byte, error) {
	enc := NewEncoder()
	defer enc.Release()

	err := enc.EncodeValue(v)
	if err != nil {
		return nil, err
	}
	return enc.Bytes(), nil
}
