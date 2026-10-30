package core

import (
	"errors"
	"fmt"
)

func readlength(data []byte) (int, int) {
	pos, length := 0, 0
	for pos = range data {
		b := data[pos]
		if !(b >= '0' && b <= '9') {
			return length, pos + 2
		}
		length = length*10 + int(b-'0')

	}
	return 0, 0

}

func readstring(data []byte) (string, int, error) {
	pos := 1
	for ; data[pos] != '\r'; pos++ {
	}
	return string(data[1:pos]), pos + 2, nil

}

func readerror(data []byte) (string, int, error) {
	return readstring(data)
}

func readinteger(data []byte) (int64, int, error) {
	pos := 1
	var value int64 = 0
	for ; data[pos] != '\r'; pos++ {
		value = value*10 + int64(data[pos]-'0')
	}
	return value, pos + 2, nil
}

func readbulkstring(data []byte) (string, int, error) {
	pos := 1
	len, delta := readlength(data[pos:])
	pos += delta
	return string(data[pos : pos+len]), pos + len + 2, nil
}

func readarray(data []byte) (interface{}, int, error) {
	// first character *
	pos := 1

	// reading the length
	count, delta := readlength(data[pos:])
	pos += delta

	var elems []interface{} = make([]interface{}, count)
	for i := range elems {
		elem, delta, err := decodeone(data[pos:])
		if err != nil {
			return nil, 0, err
		}
		elems[i] = elem
		pos += delta
	}
	return elems, pos, nil
}

func decodeone(data []byte) (interface{}, int, error) {
	if len(data) == 0 {
		return nil, 0, errors.New("no data")
	}
	switch data[0] {
	case '+':
		return readstring(data)
	case '-':
		return readerror(data)
	case ':':
		return readinteger(data)
	case '$':
		return readbulkstring(data)
	case '*':
		return readarray(data)
	}
	return nil, 0, nil
}
func DecodeArrayString(data []byte) ([]string, error) {
	value, err := Decode(data)
	if err != nil {
		return nil, err
	}

	ts := value.([]interface{})
	tokens := make([]string, len(ts))
	for i := range tokens {
		tokens[i] = ts[i].(string)
	}

	return tokens, nil
}

func Decode(data []byte) (interface{}, error) {
	if len(data) == 0 {
		return nil, errors.New("no data")
	}
	value, _, err := decodeone(data)
	return value, err
}

func Encode(value interface{}, isSimple bool) []byte {
	switch v := value.(type) {
	case string:
		if isSimple {
			return []byte(fmt.Sprintf("+%s\r\n", v))
		}
		return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(v), v))
	}
	return []byte{}
}
