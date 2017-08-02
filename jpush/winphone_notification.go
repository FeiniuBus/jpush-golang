package jpush

// WinphoneNotification is
type WinphoneNotification struct {
	platformNotification
	Title    interface{} `json:"title,omitempty"`
	OpenPage string      `json:"_open_page"`
}

// SetOpenPage is
func (wn *WinphoneNotification) SetOpenPage(openPage string) {
	wn.OpenPage = openPage
}

// NewWinphoneNotification is
func NewWinphoneNotification() *WinphoneNotification {
	p := new(WinphoneNotification)
	p.Alert = ""
	p.Title = nil
	p.OpenPage = ""
	p.Extras = make(map[string]interface{})

	return p
}
