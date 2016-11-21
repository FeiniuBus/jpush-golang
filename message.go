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
	case reflect.Uint32:
	case reflect.Int16:
	case reflect.Uint16:
	case reflect.Int8:
	case reflect.Uint8:
	case reflect.Int64:
	case reflect.Uint64:
	case reflect.String:
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
