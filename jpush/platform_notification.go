package jpush

import (
	"errors"
	"fmt"
	"reflect"
)

type platformNotification struct {
	Alert  interface{}            `json:"alert,omitempty"`
	Extras map[string]interface{} `json:"extras,omitempty"`
}

func (n *platformNotification) SetAlert(alert string) {
	n.Alert = alert
}

func (n *platformNotification) AddExtra(key string, value interface{}) error {
	if n.Extras == nil {
		n.Extras = make(map[string]interface{})
	}

	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Int32:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.String:
		fallthrough
	case reflect.Int:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		fallthrough
	case reflect.Bool:
		n.Extras[key] = value
	default:
		m := fmt.Sprintf("错误的数据类型, %s", t.Kind())
		return errors.New(m)
	}

	return nil
}
