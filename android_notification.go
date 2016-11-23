package jpush

// AndroidNodification is
type AndroidNodification struct {
	platformNotification
	Title     interface{} `json:"title,omitempty"`
	BuilderID int         `json:"builder_id"`
}

// SetTitle is
func (an *AndroidNodification) SetTitle(title string) {
	an.Title = title
}

// SetBuilderID is
func (an *AndroidNodification) SetBuilderID(id int) {
	an.BuilderID = id
}

// NewAndroidNodification is
func NewAndroidNodification() *AndroidNodification {
	p := new(AndroidNodification)
	p.Alert = ""
	p.Extras = make(map[string]interface{})
	p.Title = nil
	p.BuilderID = 0

	return p
}
