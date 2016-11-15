package jpush

// Notification is
type Notification struct {
	Alert    interface{}           `json:"alert,omitempty,string"`
	Ios      *IosNotification      `json:"ios,omitempty"`
	Android  *AndroidNodification  `json:"android,omitempty"`
	Winphone *WinphoneNotification `json:"winphone,omitempty"`
}

// SetAlert is
func (n *Notification) SetAlert(alert string) *Notification {
	n.Alert = alert
	return n
}

// SetAndroid is
func (n *Notification) SetAndroid(alert string, title string) *Notification {
	p := NewAndroidNodification()
	p.SetAlert(alert)
	n.Android = p
	n.Alert = alert
	return n
}

// SetIos is
func (n *Notification) SetIos(alert string) *Notification {
	p := NewIosNotification()
	p.SetAlert(alert)
	n.Ios = p
	n.Alert = alert
	return n
}

// SetWinphone is
func (n *Notification) SetWinphone(alert string) *Notification {
	p := NewWinphoneNotification()
	p.SetAlert(alert)
	n.Winphone = p
	n.Alert = alert
	return n
}

// NewNotification is
func NewNotification() *Notification {
	p := new(Notification)
	p.Alert = nil
	p.Android = nil
	p.Ios = nil
	p.Winphone = nil
	return p
}
