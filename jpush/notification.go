package jpush

// Notification is
type Notification struct {
	Alert    string               `json:"alert,omitempty"`
	Ios      IosNotification      `json:"ios,omitempty"`
	Android  AndroidNodification  `json:"android,omitempty"`
	Winphone WinphoneNotification `json:"winphone,omitempty"`
}
