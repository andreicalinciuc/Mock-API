package model

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestRequest(t *testing.T) {
	data, err := ioutil.ReadFile("request.json")
	if err != nil {
		t.Errorf("read %s", err)
	}

	var req Request

	err = json.Unmarshal(data, &req)
	if err != nil {
		t.Errorf("Unmarshal %s", err)
	}
	t.Logf("req %#v \n %s", req.Headers, req.Payload)
	t.Logf("Headers %s", req.Headers.Get("X-Powered-By"))
	t.Logf("payload %s ", req.Payload["userId"])

}
