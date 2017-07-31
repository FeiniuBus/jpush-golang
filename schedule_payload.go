package jpush

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type SchedulePayload struct {
	Name    string       `json:"name"`
	Enabled bool         `json:"enabled"`
	Trigger *TriggerNode `json:"trigger"`
	Push    *PushNode    `json:"push"`
}
type TriggerNode struct {
	Single     *TriggerSingleNode
	Periodical *TriggerPeriodicalNode
}
type TriggerSingleNode struct {
	Time string `json:"time"`
}
type TriggerPeriodicalNode struct {
	Start     string `json:"start"`
	End       string `json:"end"`
	Time      string `json:"time"`
	TimeUnit  string `json:"time_unit"`
	Frequency int    `json:"frequency"`
	Point     string `json:"point"`
}

type PushNode struct {
	Platform     string                `json:"platform"`
	Audience     string                `json:"audience"`
	Notification *PushNotificationNode `json:"notification"`
	Message      *PushMessageNode      `json:"message"`
	Options      *PushOptionsNode      `json:"options"`
}

type PushNotificationNode struct {
	Alert string `json:"alert"`
}

type PushMessageNode struct {
	MsgContent string `json:"msg_content"`
}

type PushOptionsNode struct {
	TimeToLive int `json:"time_to_live"`
}
type ScheduleUpdateRequest SchedulePayload

func (s ScheduleUpdateRequest) MarshalJSON() ([]byte, error) {
	var buff bytes.Buffer
	buff.WriteString("{")
	if len(s.Name) > 0 {
		buff.WriteString("\"name\":\"" + s.Name + "\",")
	}
	buff.WriteString("\"enabled\":" + strconv.FormatBool(s.Enabled) + ",")
	if s.Trigger != nil {
		trigger, err := json.Marshal(*s.Trigger)
		if err != nil {
			return nil, err
		}
		buff.WriteString("\"trigger\":" + string(trigger) + ",")
	}
	if s.Push != nil {
		push, err := json.Marshal(*s.Push)
		if err != nil {
			return nil, err
		}
		buff.WriteString("\"push\":" + string(push) + ",")
	}
	buff.WriteString("{},}")
	return buff.Bytes(), nil
}

const (
	Single = iota
	Periodical
)
const (
	N = iota
	M
)

func NewScheduleTrigger(tp int) *TriggerNode {
	t := new(TriggerNode)
	if tp == Single {
		t.Single = new(TriggerSingleNode)
	} else if tp == Periodical {
		t.Periodical = new(TriggerPeriodicalNode)
	}
	return t
}

func NewSchedulePush(tp int, content, audience, platform string) *PushNode {
	p := new(PushNode)
	p.Audience = audience
	p.Platform = platform
	p.Options = &PushOptionsNode{3600}
	if tp == M {
		p.Message = &PushMessageNode{
			MsgContent: content,
		}
	} else if tp == N {
		p.Notification = &PushNotificationNode{
			Alert: content,
		}
	}
	return p
}

func NewSchedulePayload(scheduleType, pushType int, content, audience, platform string) *SchedulePayload {
	s := new(SchedulePayload)
	s.Enabled = true
	s.Trigger = NewScheduleTrigger(scheduleType)
	s.Push = NewSchedulePush(pushType, content, audience, platform)
	return s
}
