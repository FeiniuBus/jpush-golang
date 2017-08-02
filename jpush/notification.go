package jpush

// Notification is
type Notification struct {
	Alert    interface{}           `json:"alert,omitempty,string"`
	Ios      *IosNotification      `json:"ios,omitempty"`
	Android  *AndroidNodification  `json:"android,omitempty"`
	Winphone *WinphoneNotification `json:"winphone,omitempty"`
}

// SetAlert is
func (n *Notification) SetAlert(alert string) {
	n.Alert = alert
}

// SetAndroid is
func (n *Notification) SetAndroid(alert string, title string) {
	p := NewAndroidNodification()
	p.SetAlert(alert)
	p.SetTitle(title)
	n.Android = p
	n.Alert = alert
}

// SetIos is
func (n *Notification) SetIos(alert string) {
	p := NewIosNotification()
	p.SetAlert(alert)
	n.Ios = p
	n.Alert = alert
}

// SetWinphone is
func (n *Notification) SetWinphone(alert string) {
	p := NewWinphoneNotification()
	p.SetAlert(alert)
	n.Winphone = p
	n.Alert = alert
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
