package helper

import (
	"encoding/json"
	"fmt"
	"syscall/js"
	"time"
)

// GetStructVal handles and return the JS value for struct
func GetStructVal(x interface{}) js.Value {	
    p, err := json.Marshal(x)
    if err != nil {
			fmt.Println(err)
			return js.Null()
    }
		obj := js.Global().Get("JSON").Call("parse", string(p))
		return obj
}

func GetCurrentISOTime() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05-0700")
}