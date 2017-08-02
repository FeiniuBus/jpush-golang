package jpush

// PushPayload is
type PushPayload struct {
	Platform     *Platform     `json:"platform,omitempty"`
	Audience     *Audience     `json:"audience,omitempty"`
	Notification *Notification `json:"notification,omitempty"`
	Message      *Message      `json:"message,omitempty"`
	Options      *Options      `json:"options,omitempty"`
}

// NewPushPayload is
func NewPushPayload() *PushPayload {
	p := new(PushPayload)
	p.Platform = nil
	p.Audience = nil
	p.Notification = nil
	p.Message = nil
	p.Options = NewOptions()

	return p
}
