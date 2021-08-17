package gojson

func (d *Decoder) DecodeInt8() (int8, error) {
	var negative bool

	if c := d.NextChar(); c == 'n' {
		return 0, d.AssetNull()
	} else if c == '-' {
		negative = true
		d.cursor++
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
	res, err := ConvertBytesToInt8(d.data[begin:d.cursor], begin)
	if err != nil {
		return 0, err
	}

	if !negative {
		return res, nil
	}

	return -res, nil
}

func (d *Decoder) DecodeInt16() (int16, error) {
	var negative bool

	if c := d.NextChar(); c == 'n' {
		return 0, d.AssetNull()
	} else if c == '-' {
		negative = true
		d.cursor++
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
	res, err := ConvertBytesToInt16(d.data[begin:d.cursor], begin)
	if err != nil {
		return 0, err
	}

	if !negative {
		return res, nil
	}

	return -res, nil
}

func (d *Decoder) DecodeInt32() (int32, error) {
	var negative bool

	if c := d.NextChar(); c == 'n' {
		return 0, d.AssetNull()
	} else if c == '-' {
		negative = true
		d.cursor++
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
	res, err := ConvertBytesToInt32(d.data[begin:d.cursor], begin)
	if err != nil {
		return 0, err
	}

	if !negative {
		return res, nil
	}

	return -res, nil
}


func (d *Decoder) DecodeNumber() (interface{}, error) {
	cursor := d.cursor
	if d.data[cursor] == '-'{
		cursor++
	}
	data := d.data[cursor:]
	for _, c := range data {
		if c == '.' {
			return d.DecodeFloat64()
		}

		if	c == ']' ||
			c == '}' ||
			c == ',' {
			break
		} else if IsSkip(c) {
			continue
		} else if !IsNumber(c) {
			return 0, NewParseError(d.data[d.cursor], d.cursor)
		}
	}
	return d.DecodeInt64()
}

func (d *Decoder) DecodeInt64() (int64, error) {
	var negative bool

	if c := d.NextChar(); c == 'n' {
		return 0, d.AssetNull()
	} else if c == '-' {
		negative = true
		d.cursor++
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
	res, err := ConvertBytesToInt64(d.data[begin:d.cursor], begin)
	if err != nil {
		return 0, err
	}

	if !negative {
		return res, nil
	}

	return -res, nil
}

func (d *Decoder) DecodeInt() (int, error) {
	v, err := d.DecodeInt64()
	if err != nil {
		return -1, err
	}

	return int(v), nil
}

func (d *Decoder) SkipInt8() error {
	_, err := d.DecodeInt8()
	return err
}

func (d *Decoder) SkipInt16() error {
	_, err := d.DecodeInt16()
	return err
}

func (d *Decoder) SkipInt32() error {
	_, err := d.DecodeInt32()
	return err
}

func (d *Decoder) SkipInt64() error {
	_, err := d.DecodeInt64()
	return err
}

func (d *Decoder) SkipInt() error {
	_, err := d.DecodeInt()
	return err
}
