package jpush

import (
	"bytes"
	"encoding/json"
)

// Platform is
type Platform struct {
	DeviceTypes []string `json:"deviceTypes,omitempty"`
	isAll       bool
}

// SetAll is
func (plat *Platform) SetAll() *Platform {
	plat.isAll = true
	return plat
}

// Ios is
func (plat *Platform) Ios() *Platform {
	if plat.isAll {
		plat.isAll = false
	}
	if plat.DeviceTypes == nil {
		plat.DeviceTypes = make([]string, 0)
	}

	plat.DeviceTypes = append(plat.DeviceTypes, DeviceType.String(Ios))
	return plat
}

// Android is
func (plat *Platform) Android() *Platform {
	if plat.isAll {
		plat.isAll = false
	}
	if plat.DeviceTypes == nil {
		plat.DeviceTypes = make([]string, 0)
	}

	plat.DeviceTypes = append(plat.DeviceTypes, DeviceType.String(Android))
	return plat
}

// Winphone is
func (plat *Platform) Winphone() *Platform {
	if plat.isAll {
		plat.isAll = false
	}
	if plat.DeviceTypes == nil {
		plat.DeviceTypes = make([]string, 0)
	}

	plat.DeviceTypes = append(plat.DeviceTypes, DeviceType.String(Winphone))
	return plat
}

// MarshalJSON is
func (plat Platform) MarshalJSON() ([]byte, error) {
	if plat.isAll {
		buffer := bytes.NewBufferString("all")
		return buffer.Bytes(), nil
	}
	return json.Marshal(plat.DeviceTypes)
}

// DeviceType is
type DeviceType int

func (dt DeviceType) String() string {
	switch dt {
	case Android:
		return "android"
	case Ios:
		return "ios"
	case Winphone:
		return "winphone"
	}

	return ""
}

const (
	// Android is
	Android DeviceType = iota
	// Ios is
	Ios
	// Winphone is
	Winphone
)
