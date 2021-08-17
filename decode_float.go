package gojson

func (d *Decoder) ReadFloat() ([]byte, error) {
	if d.IsNull() {
		return nil, nil
	}

	data := d.data[d.cursor:]
	begin := d.cursor

	for _, c := range data {
		if c == ']' ||
			c == '}' ||
			c == ',' {
			goto End
		} else if IsSkip(c) {
			continue
		} else if !IsFloat(c) {
			return nil, NewParseError(d.data[d.cursor], d.cursor)
		}

		d.cursor++
	}

End:
	return d.data[begin:d.cursor], nil
}

func (d *Decoder) DecodeFloat32() (float32, error) {
	var negative bool

	if d.IsNull() {
		return 0, nil
	}

	if d.Char() == '-' {
		negative = true
		d.cursor++
	}

	data := d.data[d.cursor:]
	begin := d.cursor

	for _, c := range data {
		if c == ']' ||
			c == '}' ||
			c == ',' {
			goto End
		} else if IsSkip(c) {
			continue
		} else if !IsFloat(c) {
			return 0, NewParseError(d.data[d.cursor], d.cursor)
		}

		d.cursor++
	}

End:
	v, err := ConvertBytesToFloat32(d.data[begin:d.cursor], begin)
	if err != nil {
		return 0, err
	}

	if !negative {
		return v, nil
	}

	return -v, nil
}

func (d *Decoder) DecodeFloat64() (float64, error) {
	var negative bool

	if d.IsNull() {
		return 0, nil
	}

	if d.Char() == '-' {
		negative = true
		d.cursor++
	}

	data := d.data[d.cursor:]
	begin := d.cursor

	for _, c := range data {
		if c == ']' ||
			c == '}' ||
			c == ',' {
			goto End
		} else if IsSkip(c) {
			continue
		} else if !IsFloat(c) {
			return 0, NewParseError(d.data[d.cursor], d.cursor)
		}

		d.cursor++
	}

End:
	v, err := ConvertBytesToFloat64(d.data[begin:d.cursor], begin)
	if err != nil {
		return 0, err
	}

	if !negative {
		return v, nil
	}

	return -v, nil
}

func (d *Decoder) SkipFloat32() error {
	_, err := d.DecodeFloat32()
	return err
}

func (d *Decoder) SkipFloat64() error {
	if d.Char() == '-' {
		d.cursor++
	}

	data := d.data[d.cursor:]

	for _, c := range data {
		if c == ']' ||
			c == '}' ||
			c == ',' {
			return nil
		} else if IsSkip(c) {
			continue
		} else if !IsFloat(c) {
			return NewParseError(d.data[d.cursor], d.cursor)
		}

		d.cursor++
	}

	return nil
}
