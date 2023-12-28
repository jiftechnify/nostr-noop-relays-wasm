package nostr

import (
	"bytes"
	"encoding/json"
	"errors"
)

type R2C interface {
	json.Marshaler
	r2c()
}

type R2C_EOSE struct {
	SubID string
}

func (e R2C_EOSE) MarshalJSON() ([]byte, error) {
	a := []any{
		"EOSE",
		e.SubID,
	}
	return json.Marshal(a)
}
func (e R2C_EOSE) r2c() {}

type R2C_OK struct {
	EventID string
	OK      bool
	Reason  string
}

func (e R2C_OK) MarshalJSON() ([]byte, error) {
	a := []any{
		"OK",
		e.EventID,
		e.OK,
		e.Reason,
	}
	return json.Marshal(a)
}
func (e R2C_OK) r2c() {}

type C2R interface {
	json.Unmarshaler
	c2r()
}

func ParseC2RMsg(data []byte) C2R {
	var c2r C2R
	switch {
	case bytes.Contains(data, []byte(`"REQ"`)):
		c2r = &C2R_REQ{}
	case bytes.Contains(data, []byte(`"EVENT"`)):
		c2r = &C2R_EVENT{}
	default:
		return nil
	}

	if err := json.Unmarshal(data, c2r); err != nil {
		return nil
	}
	return c2r
}

type C2R_REQ struct {
	SubID   string
	Filters []map[string]any
}

var invalidREQ = errors.New("invalid REQ")

func (e *C2R_REQ) UnmarshalJSON(data []byte) error {
	a := []*json.RawMessage{}
	if err := json.Unmarshal(data, &a); err != nil {
		return invalidREQ
	}
	if len(a) < 3 {
		return invalidREQ
	}
	if !bytes.Equal(*a[0], []byte(`"REQ"`)) {
		return invalidREQ
	}

	var subID string
	if err := json.Unmarshal(*a[1], &subID); err != nil {
		return invalidREQ
	}

	filters := make([]map[string]any, 0)
	for _, v := range a[2:] {
		var f map[string]any
		if err := json.Unmarshal(*v, &f); err != nil {
			return invalidREQ
		}
		filters = append(filters, f)
	}

	*e = C2R_REQ{
		SubID:   subID,
		Filters: filters,
	}
	return nil
}
func (e *C2R_REQ) c2r() {}

type Event struct {
	ID string `json:"id"`
}

type C2R_EVENT struct {
	Event Event
}

var invalidEVENT = errors.New("invalid EVENT")

func (e *C2R_EVENT) UnmarshalJSON(data []byte) error {
	a := []*json.RawMessage{}
	if err := json.Unmarshal(data, &a); err != nil {
		return invalidEVENT
	}
	if len(a) < 2 {
		return invalidEVENT
	}
	if !bytes.Equal(*a[0], []byte(`"EVENT"`)) {
		return invalidEVENT
	}

	var ev Event
	if err := json.Unmarshal(*a[1], &ev); err != nil {
		return invalidEVENT
	}

	*e = C2R_EVENT{
		Event: ev,
	}
	return nil
}
func (e C2R_EVENT) c2r() {}
