package jpush

import (
	"errors"
	"reflect"
)

// Message is
type Message struct {
	Title       interface{}            `json:"title,omitempty,string"`
	Content     string                 `json:"msg_content,omitempty"`
	ContentType interface{}            `json:"content_type,omitempty,string"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

// SetTitle is
func (m *Message) SetTitle(title string) *Message {
	m.Title = title
	return m
}

// SetContentType is
func (m *Message) SetContentType(ctype string) *Message {
	m.ContentType = ctype
	return m
}

// AddExtra is
func (m *Message) AddExtra(key string, value interface{}) (*Message, error) {
	if m.Extras == nil {
		m.Extras = make(map[string]interface{})
	}

	var t = reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Int32:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.String:
		fallthrough
	case reflect.Bool:
		m.Extras[key] = value
	default:
		return m, errors.New("错误的数据类型")
	}

	return m, nil
}

// NewMessage is
func NewMessage(c string) *Message {
	p := new(Message)
	p.Title = nil
	p.Content = c
	p.ContentType = nil
	p.Extras = nil

	return p
}
