package globals

import "encoding/json"

func MapToJson(m map[string]interface{}) string {
	if m == nil {
		return "{}"
	}
	bs, _ := json.Marshal(m)
	return Bytes2str(bs)
}

func JsonToMap(s string) map[string]interface{} {
	m := make(map[string]interface{})
	bs := Str2bytes(s)

	json.Unmarshal(bs, &m)
	return m
}
