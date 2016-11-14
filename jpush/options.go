package jpush

// Options is
type Options struct {
	SendNo          int   `json:"sendno,omitempty"`
	OverrideMsgID   int64 `json:"override_msg_id,omitempty"`
	TimeToLive      int64 `json:"time_to_live,omitempty"`
	BigPushDuration int64 `json:"big_push_duration,omitempty"`
	ApnsProduction  bool  `json:"apns_production,omitempty"`
}

// NewOptions is
func NewOptions() *Options {
	p := new(Options)
	p.SendNo = 0
	p.OverrideMsgID = 0
	p.TimeToLive = -1
	p.BigPushDuration = 0
	p.ApnsProduction = false

	return p
}
