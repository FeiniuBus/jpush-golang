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
func (plat *Platform) SetAll() {
	plat.DeviceTypes = nil
	plat.isAll = true
}

// Ios is
func (plat *Platform) Ios() {
	if plat.isAll {
		plat.isAll = false
	}
	if plat.DeviceTypes == nil {
		plat.DeviceTypes = make([]string, 0)
	}

	plat.DeviceTypes = append(plat.DeviceTypes, DeviceType.String(Ios))
}

// Android is
func (plat *Platform) Android() {
	if plat.isAll {
		plat.isAll = false
	}
	if plat.DeviceTypes == nil {
		plat.DeviceTypes = make([]string, 0)
	}

	plat.DeviceTypes = append(plat.DeviceTypes, DeviceType.String(Android))
}

// Winphone is
func (plat *Platform) Winphone() {
	if plat.isAll {
		plat.isAll = false
	}
	if plat.DeviceTypes == nil {
		plat.DeviceTypes = make([]string, 0)
	}

	plat.DeviceTypes = append(plat.DeviceTypes, DeviceType.String(Winphone))
}

// MarshalJSON is
func (plat Platform) MarshalJSON() ([]byte, error) {
	if plat.isAll {
		buffer := bytes.NewBufferString("\"all\"")
		return buffer.Bytes(), nil
	}
	return json.Marshal(plat.DeviceTypes)
}

// UnmarshalJSON is
func (plat *Platform) UnmarshalJSON(data []byte) error {
	if string(data) == "all" {
		plat.isAll = true
	} else {
		return json.Unmarshal(data, &plat.DeviceTypes)
	}
	return nil
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
