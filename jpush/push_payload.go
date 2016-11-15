package jpush

// PushPayload is
type PushPayload struct {
	Platform     Platform     `json:"platform,omitempty"`
	Audience     Audience     `json:"audience,omitempty"`
	Notification Notification `json:"notification,omitempty"`
	Message      Message      `json:"message,omitempty"`
	Options      Options      `json:"options,omitempty"`
}
