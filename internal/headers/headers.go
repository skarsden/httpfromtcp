package headers

import (
	"bytes"
	"fmt"
	"strings"
)

const crlf = "\r\n"
const specialChars = "1#$%&'*+-.^_`|~ "

type Headers map[string]string

func NewHeaders() Headers {
	return map[string]string{}
}

func (h Headers) Get(key string) (string, bool) {
	val, ok := h[strings.ToLower(key)]
	return val, ok
}

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	idx := bytes.Index(data, []byte(crlf))
	if idx == -1 {
		return 0, false, nil
	}
	if idx == 0 {
		//empty line. headers done, consume crlf
		return 2, true, nil
	}

	parts := bytes.SplitN(data[:idx], []byte(":"), 2)
	key := strings.ToLower(string(parts[0]))

	if key != strings.TrimRight(key, " ") || len(strings.Split(strings.Trim(key, " "), " ")) > 1 {
		return 0, false, fmt.Errorf("invalid header name: %s", key)
	}
	for _, c := range key {
		if !(c >= 'A' && c <= 'Z') && !(c >= 'a' && c <= 'z') && !(c >= 1 && c <= 9) && !strings.Contains(specialChars, string(c)) {
			return 0, false, fmt.Errorf("invalid character in header: %v", c)
		}
	}

	value := bytes.TrimSpace(parts[1])
	key = strings.TrimSpace(key)

	h.Set(strings.ToLower(key), string(value))
	return idx + 2, false, nil
}

func (h Headers) Set(key, value string) {
	key = strings.ToLower(key)
	v, ok := h[key]
	if ok {
		value = strings.Join([]string{v, value}, ", ")
	}
	h[key] = value
}
