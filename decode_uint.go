package gojson

func (d *Decoder) DecodeUint8() (uint8, error) {
	if d.IsNull() {
		return 0, nil
	}

	data := d.data[d.cursor:]
	begin := d.cursor

	for _, c := range data {
		if c == '.' ||
			c == ']' ||
			c == '}' ||
			c == ',' {
			goto End
		} else if IsSkip(c) {
			continue
		} else if !IsNumber(c) {
			return 0, NewParseError(d.data[d.cursor], d.cursor)
		}

		d.cursor++
	}

End:
	res, err := ConvertBytesToUint8(d.data[begin:d.cursor], begin)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (d *Decoder) DecodeUint16() (uint16, error) {
	if d.IsNull() {
		return 0, nil
	}

	data := d.data[d.cursor:]
	begin := d.cursor

	for _, c := range data {
		if c == '.' ||
			c == ']' ||
			c == '}' ||
			c == ',' {
			goto End
		} else if IsSkip(c) {
			continue
		} else if !IsNumber(c) {
			return 0, NewParseError(d.data[d.cursor], d.cursor)
		}

		d.cursor++
	}

End:
	res, err := ConvertBytesToUint16(d.data[begin:d.cursor], begin)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (d *Decoder) DecodeUint32() (uint32, error) {
	if d.IsNull() {
		return 0, nil
	}

	data := d.data[d.cursor:]
	begin := d.cursor

	for _, c := range data {
		if c == '.' ||
			c == ']' ||
			c == '}' ||
			c == ',' {
			goto End
		} else if IsSkip(c) {
			continue
		} else if !IsNumber(c) {
			return 0, NewParseError(d.data[d.cursor], d.cursor)
		}

		d.cursor++
	}

End:
	res, err := ConvertBytesToUint32(d.data[begin:d.cursor], begin)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (d *Decoder) DecodeUint64() (uint64, error) {
	if d.IsNull() {
		return 0, nil
	}

	data := d.data[d.cursor:]
	begin := d.cursor

	for _, c := range data {
		if c == '.' ||
			c == ']' ||
			c == '}' ||
			c == ',' {
			goto End
		} else if IsSkip(c) {
			continue
		} else if !IsNumber(c) {
			return 0, NewParseError(d.data[d.cursor], d.cursor)
		}

		d.cursor++
	}

End:
	res, err := ConvertBytesToUint64(d.data[begin:d.cursor], begin)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (d *Decoder) DecodeUint() (uint, error) {
	v, err := d.DecodeUint64()
	if err != nil {
		return 0, err
	}

	return uint(v), nil
}

func (d *Decoder) SkipUint8() error {
	_, err := d.DecodeUint8()
	return err
}

func (d *Decoder) SkipUint16() error {
	_, err := d.DecodeUint16()
	return err
}

func (d *Decoder) SkipUint32() error {
	_, err := d.DecodeUint32()
	return err
}

func (d *Decoder) SkipUint64() error {
	_, err := d.DecodeUint64()
	return err
}

func (d *Decoder) SkipUint() error {
	_, err := d.DecodeUint()
	return err
}
