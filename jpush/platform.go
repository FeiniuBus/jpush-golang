package jpush

// Platform is
type Platform struct {
	DeviceTypes []string `json:"deviceTypes,omitempty"`
}

// Ios is
func (plat *Platform) Ios() *Platform {
	if plat.DeviceTypes == nil {
		plat.DeviceTypes = make([]string, 0)
	}

	plat.DeviceTypes = append(plat.DeviceTypes, DeviceType.String(Ios))
	return plat
}

// Android is
func (plat *Platform) Android() *Platform {
	if plat.DeviceTypes == nil {
		plat.DeviceTypes = make([]string, 0)
	}

	plat.DeviceTypes = append(plat.DeviceTypes, DeviceType.String(Android))
	return plat
}

// Winphone is
func (plat *Platform) Winphone() *Platform {
	if plat.DeviceTypes == nil {
		plat.DeviceTypes = make([]string, 0)
	}

	plat.DeviceTypes = append(plat.DeviceTypes, DeviceType.String(Winphone))
	return plat
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
