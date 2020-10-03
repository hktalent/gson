package gson

import (
	"encoding/json"
	"io"
)

// New JSON from string, []byte, or io.Reader.
func New(raw interface{}) (j JSON) {
	switch v := raw.(type) {
	case string:
		_ = j.UnmarshalJSON([]byte(v))
	case []byte:
		_ = j.UnmarshalJSON(v)
	case io.Reader:
		_ = json.NewDecoder(v).Decode(&j)
	default:
		j.Sets(raw)
	}
	return
}

// UnmarshalJSON interface
func (j *JSON) UnmarshalJSON(b []byte) error {
	var v interface{}
	err := json.Unmarshal(b, &v)
	*j = JSON{v}
	return err
}

// Set by json path. It's a shortcut for Sets.
func (j *JSON) Set(path string, val interface{}) *JSON {
	return j.Sets(val, Path(path)...)
}

// Sets element by path sections. If a section is not string or int, it will be ignored.
func (j *JSON) Sets(value interface{}, sections ...interface{}) *JSON {
	last := len(sections) - 1
	val := j.val
	var override func(interface{})

	if last == -1 {
		j.val = toJSONVal(value)
		return j
	}

	for i, sect := range sections {
		switch k := sect.(type) {
		case int:
			arr, ok := val.([]interface{})
			if !ok || k >= len(arr) {
				nArr := make([]interface{}, k+1)
				copy(nArr, arr)
				arr = nArr
				override(arr)
			}
			if i == last {
				arr[k] = toJSONVal(value)
				return j
			}
			val = arr[k]

			override = func(val interface{}) {
				arr[k] = val
			}
		case string:
			obj, ok := val.(map[string]interface{})
			if !ok {
				obj = map[string]interface{}{}
				override(obj)
			}
			if i == last {
				obj[k] = toJSONVal(value)
			}
			val = obj[k]

			override = func(val interface{}) {
				obj[k] = val
			}
		}
	}
	return j
}
