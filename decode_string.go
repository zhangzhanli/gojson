package gojson

func (d *Decoder) ReadString() ([]byte, error) {
	if c := d.NextChar(); c == 'n' {
		return nil, d.AssetNull()
	} else if c != '"' {
		return nil, NewParseError(d.data[d.cursor], d.cursor)
	}

	data := d.data[d.cursor:]

	for i, c := range data {
		if c == '"' && i != 0 {
			if prev := i - 1; prev >= 0 && data[prev] == '\\' {
				continue
			}

			d.cursor = d.cursor + i + 1
			return data[:i+1], nil
		}
	}

	return nil, NewParseError(d.data[d.length-1], d.length-1)
}

func (d *Decoder) DecodeString() (string, error) {
	if c := d.NextChar(); c == 'n' {
		return "", d.AssetNull()
	} else if c != '"' {
		return "", NewParseError(d.data[d.cursor], d.cursor)
	}

	d.cursor++
	src := d.data[d.cursor:]
	data := make([]byte, len(src))

	k := 0
	for i := 0; i < len(src); i++ {
		c := src[i]
		if c == '"' {
			d.cursor = d.cursor + i + 1
			return UnsafeConvertBytesToString(data[:k]), nil
		} else if c == '\\' {
			i++
			if i < len(src) {
				c = src[i]
				switch c {
				case '\\':
					data[k] = '\\'
				case '"':
					data[k] = '"'
				case 'n':
					data[k] = '\n'
				case 't':
					data[k] = '\t'
				case 'f':
					data[k] = '\f'
				case 'b':
					data[k] = '\b'
				case 'r':
					data[k] = '\r'
				default:
					data[k] = '\\'
					k++
					data[k] = c
				}
			}
		} else {
			data[k] = c
		}
		k++
	}

	return "", NewParseError(d.data[d.length-1], d.length-1)
}

func (d *Decoder) SkipString() error {
	d.cursor++
	data := d.data[d.cursor:]

	for i, c := range data {
		if c == '"' {
			if prev := i - 1; prev >= 0 && data[prev] == '\\' {
				continue
			}

			d.cursor = d.cursor + i + 1
			return nil
		}
	}

	return NewParseError(d.data[d.length-1], d.length-1)
}
