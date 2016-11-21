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
func (audience *Audience) SetAll() *Audience {
	audience.isAll = true
	audience.Audiences = nil
	return audience
}

// SetTag is
func (audience *Audience) SetTag(value ...string) *Audience {
	target := new(AudienceTarget)
	target.SetTag(value)
	audience.addWithAudienceTarget(target)
	return audience
}

// SetTagWithArray is
func (audience *Audience) SetTagWithArray(values []string) *Audience {
	target := new(AudienceTarget)
	target.SetTag(values)
	audience.addWithAudienceTarget(target)
	return audience
}

// SetTagAnd is
func (audience *Audience) SetTagAnd(value ...string) *Audience {
	target := new(AudienceTarget)
	target.SetTagAnd(value)
	audience.addWithAudienceTarget(target)
	return audience
}

// SetTagAndWithArray is
func (audience *Audience) SetTagAndWithArray(values []string) *Audience {
	target := new(AudienceTarget)
	target.SetTagAnd(values)
	audience.addWithAudienceTarget(target)
	return audience
}

// SetAlias is
func (audience *Audience) SetAlias(value ...string) *Audience {
	target := new(AudienceTarget)
	target.SetAlias(value)
	audience.addWithAudienceTarget(target)
	return audience
}

// SetAliasWithArray is
func (audience *Audience) SetAliasWithArray(values []string) *Audience {
	target := new(AudienceTarget)
	target.SetAlias(values)
	audience.addWithAudienceTarget(target)
	return audience
}

// SetSegment is
func (audience *Audience) SetSegment(value ...string) *Audience {
	target := new(AudienceTarget)
	target.SetSegment(value)
	audience.addWithAudienceTarget(target)
	return audience
}

// SetRegistrationID is
func (audience *Audience) SetRegistrationID(value ...string) *Audience {
	target := new(AudienceTarget)
	target.SetRegistrationID(value)
	audience.addWithAudienceTarget(target)
	return audience
}

// MarshalJSON is
func (audience Audience) MarshalJSON() ([]byte, error) {
	if audience.isAll {
		buffer := bytes.NewBufferString("all")
		return buffer.Bytes(), nil
	}
	return json.Marshal(audience.Audiences)
}

func (audience *Audience) addWithAudienceTarget(target *AudienceTarget) *Audience {
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
	return audience
}

// AudienceTarget is
type AudienceTarget struct {
	Type         string   `json:"andienceType,omitempty"`
	ValueBuilder []string `json:"valueBuilder,omitempty"`
}

// SetTag is
func (at *AudienceTarget) SetTag(values []string) *AudienceTarget {
	at.Type = Tag.String()
	at.ValueBuilder = values

	return at
}

// SetTagAnd is
func (at *AudienceTarget) SetTagAnd(values []string) *AudienceTarget {
	at.Type = TagAnd.String()
	at.ValueBuilder = values

	return at
}

// SetAlias is
func (at *AudienceTarget) SetAlias(values []string) *AudienceTarget {
	at.Type = Alias.String()
	at.ValueBuilder = values

	return at
}

// SetSegment is
func (at *AudienceTarget) SetSegment(values []string) *AudienceTarget {
	at.Type = Segment.String()
	at.ValueBuilder = values

	return at
}

// SetRegistrationID is
func (at *AudienceTarget) SetRegistrationID(values []string) *AudienceTarget {
	at.Type = RegistrationID.String()
	at.ValueBuilder = values

	return at
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
