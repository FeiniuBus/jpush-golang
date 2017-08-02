package jpush

import (
	"bytes"
	"encoding/json"
)

// Audience is
type Audience struct {
	Audiences map[string][]string `json:"dictionary,omitempty"`
	isAll     bool
}

// SetAll is
func (audience *Audience) SetAll() {
	audience.isAll = true
	audience.Audiences = nil
}

// SetTag is
func (audience *Audience) SetTag(value ...string) {
	target := new(AudienceTarget)
	target.SetTag(value)
	audience.addWithAudienceTarget(target)
}

// SetTagWithArray is
func (audience *Audience) SetTagWithArray(values []string) {
	target := new(AudienceTarget)
	target.SetTag(values)
	audience.addWithAudienceTarget(target)
}

// SetTagAnd is
func (audience *Audience) SetTagAnd(value ...string) {
	target := new(AudienceTarget)
	target.SetTagAnd(value)
	audience.addWithAudienceTarget(target)
}

// SetTagAndWithArray is
func (audience *Audience) SetTagAndWithArray(values []string) {
	target := new(AudienceTarget)
	target.SetTagAnd(values)
	audience.addWithAudienceTarget(target)
}

// SetAlias is
func (audience *Audience) SetAlias(value ...string) {
	target := new(AudienceTarget)
	target.SetAlias(value)
	audience.addWithAudienceTarget(target)
}

// SetAliasWithArray is
func (audience *Audience) SetAliasWithArray(values []string) {
	target := new(AudienceTarget)
	target.SetAlias(values)
	audience.addWithAudienceTarget(target)
}

// SetSegment is
func (audience *Audience) SetSegment(value ...string) {
	target := new(AudienceTarget)
	target.SetSegment(value)
	audience.addWithAudienceTarget(target)
}

// SetRegistrationID is
func (audience *Audience) SetRegistrationID(value ...string) {
	target := new(AudienceTarget)
	target.SetRegistrationID(value)
	audience.addWithAudienceTarget(target)
}

// MarshalJSON is
func (audience Audience) MarshalJSON() ([]byte, error) {
	if audience.isAll {
		buffer := bytes.NewBufferString("all")
		return buffer.Bytes(), nil
	}
	return json.Marshal(audience.Audiences)
}

// UnmarshalJSON is
func (audience *Audience) UnmarshalJSON(data []byte) error {
	if string(data) == "all" {
		audience.isAll = true
	} else {
		return json.Unmarshal(data, &audience.Audiences)
	}
	return nil
}

func (audience *Audience) addWithAudienceTarget(target *AudienceTarget) {
	if target.ValueBuilder != nil {
		if audience.isAll {
			audience.isAll = false
		}

		if audience.Audiences == nil {
			audience.Audiences = make(map[string][]string)
		}

		if val, ok := audience.Audiences[target.Type]; ok {
			val = append(val, target.ValueBuilder...)
			audience.Audiences[target.Type] = val
		} else {
			audience.Audiences[target.Type] = target.ValueBuilder
		}
	}
}

// AudienceTarget is
type AudienceTarget struct {
	Type         string   `json:"andienceType,omitempty"`
	ValueBuilder []string `json:"valueBuilder,omitempty"`
}

// SetTag is
func (at *AudienceTarget) SetTag(values []string) {
	at.Type = Tag.String()
	at.ValueBuilder = values
}

// SetTagAnd is
func (at *AudienceTarget) SetTagAnd(values []string) {
	at.Type = TagAnd.String()
	at.ValueBuilder = values
}

// SetAlias is
func (at *AudienceTarget) SetAlias(values []string) {
	at.Type = Alias.String()
	at.ValueBuilder = values
}

// SetSegment is
func (at *AudienceTarget) SetSegment(values []string) {
	at.Type = Segment.String()
	at.ValueBuilder = values
}

// SetRegistrationID is
func (at *AudienceTarget) SetRegistrationID(values []string) {
	at.Type = RegistrationID.String()
	at.ValueBuilder = values
}

// AudienceType is
type AudienceType int

// String is
func (at AudienceType) String() string {
	switch at {
	case Tag:
		return "tag"
	case TagAnd:
		return "tag_and"
	case Alias:
		return "alias"
	case Segment:
		return "segment"
	case RegistrationID:
		return "registration_id"
	}

	return ""
}

const (
	// Tag is
	Tag AudienceType = iota
	// TagAnd is
	TagAnd
	// Alias is
	Alias
	// Segment is
	Segment
	// RegistrationID is
	RegistrationID
)
