package jpush

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

//SchedulePayload is
type SchedulePayload struct {
	Name    string       `json:"name"`
	Enabled bool         `json:"enabled"`
	Trigger *TriggerNode `json:"trigger"`
	Push    *PushPayload `json:"push"`
}

//TriggerNode is
type TriggerNode struct {
	Single     *TriggerSingleNode     `json:"single"`
	Periodical *TriggerPeriodicalNode `json:"periodical"`
}

//MarshalJSON is
func (t TriggerNode) MarshalJSON() ([]byte, error) {
	var j []byte
	var err error
	buff := bytes.NewBufferString("{")
	if t.Single != nil {
		j, err = json.Marshal(t.Single)
		buff.WriteString("\"single\":")
		buff.Write(j)
	}
	if t.Single != nil && t.Periodical != nil {
		buff.WriteString(",")
	}
	if t.Periodical != nil {
		buff.WriteString("\"periodical\":")
		j, err = json.Marshal(t.Periodical)
		buff.Write(j)
	}
	if err != nil {
		return nil, err
	}
	buff.WriteString("}")
	return buff.Bytes(), nil
}

const (
	scheduleDateTimeFormt = "2006-01-02 15:04:05"
	scheduleTimeFormat    = "15:04:05"
)

//ScheduleDateTime is
type ScheduleDateTime struct {
	Time *time.Time
}

//MarshalJSON is
func (t ScheduleDateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(scheduleDateTimeFormt)+2)
	b = append(b, '"')
	if t.Time != nil {
		b = t.Time.AppendFormat(b, scheduleDateTimeFormt)
	}
	b = append(b, '"')
	return b, nil
}

//UnmarshalJSON is
func (t *ScheduleDateTime) UnmarshalJSON(data []byte) error {
	s := strings.Replace(string(data), "\"", "", 2)
	if s == "" {
		return nil
	}
	time, err := time.Parse(scheduleDateTimeFormt, s)
	if err == nil {
		t.Time = &time
	}
	return err
}

//ScheduleTime is
type ScheduleTime struct {
	Time *time.Time
}

//MarshalJSON is
func (t ScheduleTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(scheduleTimeFormat)+2)
	b = append(b, '"')
	if t.Time != nil {
		b = t.Time.AppendFormat(b, scheduleTimeFormat)
	}
	b = append(b, '"')
	return b, nil
}

//UnmarshalJSON is
func (t *ScheduleTime) UnmarshalJSON(data []byte) error {
	s := strings.Replace(string(data), "\"", "", 2)
	if s == "" {
		return nil
	}
	time, err := time.Parse(scheduleTimeFormat, s)
	if err == nil {
		t.Time = &time
	}
	return err
}

//TriggerSingleNode is
type TriggerSingleNode struct {
	Time ScheduleDateTime `json:"time"`
}

//TriggerPeriodicalNode is
type TriggerPeriodicalNode struct {
	Start     ScheduleDateTime `json:"start"`
	End       ScheduleDateTime `json:"end"`
	Time      ScheduleTime     `json:"time"`
	TimeUnit  string           `json:"time_unit"`
	Frequency int              `json:"frequency"`
	Point     *string          `json:"point"`
}

//ScheduleResponse is
type ScheduleResponse struct {
	ScheduleID string      `json:"schedule_id"`
	RichPush   interface{} `json:"richpush"`
	SendSource interface{} `json:"sendsource"`
	SchedulePayload
}

//ScheduleUpdateRequest is
type ScheduleUpdateRequest SchedulePayload

//MarshalJSON is
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
	buff.Truncate(buff.Len() - 1)
	buff.WriteString("}")
	return buff.Bytes(), nil
}

const (
	//Single is
	Single = iota
	//Periodical is
	Periodical
)

//NewScheduleTrigger is
func NewScheduleTrigger(tp int) *TriggerNode {
	t := new(TriggerNode)
	if tp == Single {
		t.Single = new(TriggerSingleNode)
	} else if tp == Periodical {
		t.Periodical = new(TriggerPeriodicalNode)
	}
	return t
}

//NewSchedulePayload is
func NewSchedulePayload(name string, scheduleType int, push *PushPayload) *SchedulePayload {
	s := new(SchedulePayload)
	s.Name = name
	s.Enabled = true
	s.Trigger = NewScheduleTrigger(scheduleType)
	s.Push = push
	return s
}

//NewSchedulePayloadWithSingle is
func NewSchedulePayloadWithSingle(name string, single *TriggerSingleNode, push *PushPayload) *SchedulePayload {
	s := NewSchedulePayload(name, Single, push)
	s.Trigger.Single = single
	return s
}

//NewSchedulePayloadWithPeriodical is
func NewSchedulePayloadWithPeriodical(name string, periodical *TriggerPeriodicalNode, push *PushPayload) *SchedulePayload {
	s := NewSchedulePayload(name, Periodical, push)
	s.Trigger.Periodical = periodical
	return s
}
