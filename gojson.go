package gojson

import (
	"encoding/json"
	"fmt"
)

func Unmarshal(data []byte, v interface{}) error {
	switch t := v.(type) {
	case *string:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeString()
		if err != nil {
			return err
		}

		*t = v

	case *int:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeInt()
		if err != nil {
			return err
		}

		*t = v

	case *int8:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeInt8()
		if err != nil {
			return err
		}

		*t = v

	case *int16:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeInt16()
		if err != nil {
			return err
		}

		*t = v

	case *int32:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeInt32()
		if err != nil {
			return err
		}

		*t = v

	case *int64:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeInt64()
		if err != nil {
			return err
		}

		*t = v

	case *uint:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeUint()
		if err != nil {
			return err
		}

		*t = v

	case *uint8:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeUint8()
		if err != nil {
			return err
		}

		*t = v

	case *uint16:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeUint16()
		if err != nil {
			return err
		}

		*t = v

	case *uint32:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeUint32()
		if err != nil {
			return err
		}

		*t = v

	case *uint64:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeUint64()
		if err != nil {
			return err
		}

		*t = v

	case *[]byte:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeBytes()
		if err != nil {
			return err
		}

		*t = v

	case []byte:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeBytes()
		if err != nil {
			return err
		}

		t = v

	case *map[string]interface{}:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeObject()
		if err != nil {
			return err
		}

		*t = v

	case map[string]interface{}:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeObject()
		if err != nil {
			return err
		}

		t = v

	case *[]interface{}:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeArray()
		if err != nil {
			return err
		}

		*t = v

	case []interface{}:
		dec := NewDecoder()
		dec.SetData(data)
		defer dec.Release()

		v, err := dec.DecodeArray()
		if err != nil {
			return err
		}

		t = v

	case json.Unmarshaler:
		return t.UnmarshalJSON(data)

	default:
		return fmt.Errorf("Unsupported type %T in Unmarshal", v)
	}

	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	switch t := v.(type) {
	case string:
		enc := NewEncoder()
		defer enc.Release()

		enc.EncodeString(t)
		return enc.Bytes(), nil

	case int:
		enc := NewEncoder()
		defer enc.Release()

		enc.EncodeInt(t)
		return enc.Bytes(), nil

	case int8:
		enc := NewEncoder()
		defer enc.Release()

		enc.EncodeInt8(t)
		return enc.Bytes(), nil

	case int16:
		enc := NewEncoder()
		defer enc.Release()

		enc.EncodeInt16(t)
		return enc.Bytes(), nil

	case int32:
		enc := NewEncoder()
		defer enc.Release()

		enc.EncodeInt32(t)
		return enc.Bytes(), nil

	case int64:
		enc := NewEncoder()
		defer enc.Release()

		enc.EncodeInt64(t)
		return enc.Bytes(), nil

	case uint:
		enc := NewEncoder()
		defer enc.Release()

		enc.EncodeUint(t)
		return enc.Bytes(), nil

	case uint8:
		enc := NewEncoder()
		defer enc.Release()

		enc.EncodeUint8(t)
		return enc.Bytes(), nil

	case uint16:
		enc := NewEncoder()
		defer enc.Release()

		enc.EncodeUint16(t)
		return enc.Bytes(), nil

	case uint32:
		enc := NewEncoder()
		defer enc.Release()

		enc.EncodeUint32(t)
		return enc.Bytes(), nil

	case uint64:
		enc := NewEncoder()
		defer enc.Release()

		enc.EncodeUint64(t)
		return enc.Bytes(), nil

	case []byte:
		enc := NewEncoder()
		defer enc.Release()

		enc.EncodeBytes(t)
		return enc.Bytes(), nil

	case map[string]interface{}:
		enc := NewEncoder()
		defer enc.Release()

		enc.EncodeObject(t)
		return enc.Bytes(), nil

	case []interface{}:
		enc := NewEncoder()
		defer enc.Release()

		enc.EncodeArray(t)
		return enc.Bytes(), nil

	case json.Marshaler:
		return t.MarshalJSON()

	default:
		return nil, fmt.Errorf("Unsupported type %T in Unmarshal", v)
	}
}
