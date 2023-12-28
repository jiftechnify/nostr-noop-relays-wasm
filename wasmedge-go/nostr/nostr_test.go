package nostr

import (
	"bytes"
	"testing"
)

func TestMarshalR2C_EOSE(t *testing.T) {
	r2c := R2C_EOSE{SubID: "sub"}
	want := `["EOSE","sub"]`

	got, err := r2c.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(got, []byte(want)) {
		t.Errorf("got: %s, want: %s", string(got), want)
	}
}

func TestMarshalR2C_OK(t *testing.T) {
	r2c := R2C_OK{EventID: "eventid", OK: true, Reason: "reason"}
	want := `["OK","eventid",true,"reason"]`

	got, err := r2c.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(got, []byte(want)) {
		t.Errorf("got: %s, want: %s", string(got), want)
	}
}

func TestParseC2RMsg_REQ(t *testing.T) {
	got := ParseC2RMsg([]byte(`["REQ","subid",{}]`))
	if got == nil {
		t.Fatal("unexpected error")
	}
	req, ok := got.(*C2R_REQ)
	if !ok {
		t.Fatal("want REQ, but not")
	}
	if req.SubID != "subid" {
		t.Fatal("unexpected parse result")
	}
	if len(req.Filters) != 1 {
		t.Fatal("unexpected parse result")
	}
}

func TestParseC2RMsg_EVENT(t *testing.T) {
	got := ParseC2RMsg([]byte(`["EVENT",{"id":"eventid"}]`))
	if got == nil {
		t.Fatal("unexpected parse error")
	}
	event, ok := got.(*C2R_EVENT)
	if !ok {
		t.Fatal("want EVENT, but not")
	}
	if event.Event.ID != "eventid" {
		t.Fatal("unexpected parse result")
	}
}
