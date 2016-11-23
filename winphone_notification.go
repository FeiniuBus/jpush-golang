package jpush

import (
	"errors"
	"fmt"
	"reflect"
)

// WinphoneNotification is
type WinphoneNotification struct {
	Alert    string                 `json:"alert"`
	Extras   map[string]interface{} `json:"extras"`
	Title    string                 `json:"title"`
	OpenPage string                 `json:"_open_page"`
}

// SetAlert is
func (wn *WinphoneNotification) SetAlert(alert string) *WinphoneNotification {
	wn.Alert = alert
	return wn
}

// SetOpenPage is
func (wn *WinphoneNotification) SetOpenPage(openPage string) *WinphoneNotification {
	wn.OpenPage = openPage
	return wn
}

// AddExtra is
func (wn *WinphoneNotification) AddExtra(key string, value interface{}) (*WinphoneNotification, error) {
	if wn.Extras == nil {
		wn.Extras = make(map[string]interface{})
	}

	var t = reflect.TypeOf(value)
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
		wn.Extras[key] = value
	default:
		m := fmt.Sprintf("错误的数据类型, %s", t.Kind())
		return wn, errors.New(m)
	}

	return wn, nil
}

// NewWinphoneNotification is
func NewWinphoneNotification() *WinphoneNotification {
	p := new(WinphoneNotification)
	p.Alert = ""
	p.Title = ""
	p.OpenPage = ""
	p.Extras = make(map[string]interface{})

	return p
}
