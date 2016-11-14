package jpush

import (
	"errors"
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
	case reflect.Uint32:
	case reflect.Int16:
	case reflect.Uint16:
	case reflect.Int8:
	case reflect.Uint8:
	case reflect.Int64:
	case reflect.Uint64:
	case reflect.String:
	case reflect.Bool:
		wn.Extras[key] = value
	default:
		return wn, errors.New("错误的数据类型")
	}

	return wn, nil
}

// NewWinphoneNotification is
func NewWinphoneNotification() *WinphoneNotification {
	p := new(WinphoneNotification)
	p.Alert = ""
	p.Title = ""
	p.OpenPage = ""
	p.Extras = nil

	return p
}
