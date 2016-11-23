package jpush

import (
	"errors"
	"reflect"
)

// AndroidNodification is
type AndroidNodification struct {
	Alert     string                 `json:"alert"`
	Extras    map[string]interface{} `json:"extras"`
	Title     string                 `json:"title"`
	BuilderID int                    `json:"builder_id"`
}

// SetTitle is
func (an *AndroidNodification) SetTitle(title string) *AndroidNodification {
	an.Title = title
	return an
}

// SetAlert is
func (an *AndroidNodification) SetAlert(alert string) *AndroidNodification {
	an.Alert = alert
	return an
}

// SetBuilderID is
func (an *AndroidNodification) SetBuilderID(id int) *AndroidNodification {
	an.BuilderID = id
	return an
}

// AddExtra is
func (an *AndroidNodification) AddExtra(key string, value interface{}) (*AndroidNodification, error) {
	if an.Extras == nil {
		an.Extras = make(map[string]interface{})
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
	case reflect.Bool:
		an.Extras[key] = value
	default:
		return an, errors.New("错误的数据类型")
	}

	return an, nil
}

// NewAndroidNodification is
func NewAndroidNodification() *AndroidNodification {
	p := new(AndroidNodification)
	p.Alert = ""
	p.Extras = make(map[string]interface{})
	p.Title = ""
	p.BuilderID = 0

	return p
}
