package jpush

import (
	"errors"
	"math"
)

// IosNotification is
type IosNotification struct {
	platformNotification
	Sound            interface{} `json:"sound,omitempty"`
	Badge            interface{} `json:"badge,omitempty"`
	ContentAvailable bool        `json:"content-available"`
	Category         interface{} `json:"category,omitempty"`
}

//SetSound is
func (ios *IosNotification) SetSound(sound string) {
	ios.Sound = sound
}

// AutoBadge is
func (ios *IosNotification) AutoBadge() error {
	return ios.IncBadge(1)
}

// SetBadge is
func (ios *IosNotification) SetBadge(badge int) error {
	if math.Abs(float64(badge)) > 99999 {
		return errors.New("错误的badge值")
	}

	ios.Badge = string(badge)
	return nil
}

// IncBadge is
func (ios *IosNotification) IncBadge(badge int) error {
	if math.Abs(float64(badge)) > 99999 {
		return errors.New("错误的badge值")
	}

	if badge >= 0 {
		ios.Badge = "+" + string(badge)
	} else {
		ios.Badge = string(badge)
	}

	return nil
}

// SetContentAvailable is
func (ios *IosNotification) SetContentAvailable(available bool) {
	ios.ContentAvailable = available
}

// SetCategory is
func (ios *IosNotification) SetCategory(category string) {
	ios.Category = category
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
