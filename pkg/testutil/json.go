package testutil

import "encoding/json"

func StringJSON(obj interface{}) string {
	jsonObj, _ := json.Marshal(obj)
	return string(jsonObj)
}
