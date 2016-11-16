package jpush

import (
	"errors"
	"math"
	"reflect"
)

// IosNotification is
type IosNotification struct {
	Alert            interface{}            `json:"alert,omitempty"`
	Extras           map[string]interface{} `json:"extras,omitempty"`
	Sound            interface{}            `json:"sound,omitempty"`
	Badge            interface{}            `json:"badge,omitempty"`
	ContentAvailable bool                   `json:"content-available"`
	Category         interface{}            `json:"category,omitempty"`
}

//SetSound is
func (ios *IosNotification) SetSound(sound string) *IosNotification {
	ios.Sound = sound
	return ios
}

// AutoBadge is
func (ios *IosNotification) AutoBadge() (*IosNotification, error) {
	return ios.IncBadge(1)
}

// SetBadge is
func (ios *IosNotification) SetBadge(badge int) (*IosNotification, error) {
	if math.Abs(float64(badge)) > 99999 {
		return ios, errors.New("错误的badge值")
	}

	ios.Badge = string(badge)
	return ios, nil
}

// IncBadge is
func (ios *IosNotification) IncBadge(badge int) (*IosNotification, error) {
	if math.Abs(float64(badge)) > 99999 {
		return ios, errors.New("错误的badge值")
	}

	if badge >= 0 {
		ios.Badge = "+" + string(badge)
	} else {
		ios.Badge = string(badge)
	}

	return ios, nil
}

// SetAlert is
func (ios *IosNotification) SetAlert(alert string) *IosNotification {
	ios.Alert = alert
	return ios
}

// SetContentAvailable is
func (ios *IosNotification) SetContentAvailable(available bool) *IosNotification {
	ios.ContentAvailable = available
	return ios
}

// SetCategory is
func (ios *IosNotification) SetCategory(category string) *IosNotification {
	ios.Category = category
	return ios
}

// AddExtra is
func (ios *IosNotification) AddExtra(key string, value interface{}) (*IosNotification, error) {
	if ios.Extras == nil {
		ios.Extras = make(map[string]interface{})
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
		ios.Extras[key] = value
	default:
		return ios, errors.New("错误的数据类型")
	}

	return ios, nil
}

// NewIosNotification is
func NewIosNotification() *IosNotification {
	p := new(IosNotification)
	p.Alert = nil
	p.Extras = make(map[string]interface{})
	p.ContentAvailable = false
	p.Category = nil
	p.Badge = "+1"
	p.Sound = ""

	return p
}
