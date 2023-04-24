package api

import (
	"div11/util"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEvents(t *testing.T) {
	config := util.NewConfig(":8088")
	api := NewServer(config)
	router := api.NewRouter()
	ts := httptest.NewServer(router)

	defer ts.Close()

	newreq := func(method, url string, body io.Reader) *http.Request {
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		return r
	}

	tests := []struct {
		name string
		r    *http.Request
	}{
		{name: "1: testing create_event", r: newreq("POST", ts.URL+"/create_event?user_id=6&date=2019-09-09&event_name=daily", nil)},
		{name: "2: testing events_for_day", r: newreq("GET", ts.URL+"/events_for_day?user_id=6&date=2019-09-09", nil)},
		{name: "3: testing events_for_week", r: newreq("GET", ts.URL+"/events_for_week?user_id=6&date=2019-09-09", nil)},
		{name: "4: testing events_for_month", r: newreq("GET", ts.URL+"/events_for_month?user_id=6&date=2019-09-09", nil)},
		{name: "5: testing update_event", r: newreq("POST", ts.URL+"/update_event?user_id=6&date=2021-09-09&event_name=daily", nil)},
		{name: "6: testing delete_event", r: newreq("POST", ts.URL+"/delete_event?user_id=6&date=2021-09-09&event_name=daily", nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tt.r)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Status code is not OK: %v", resp.StatusCode)
			}
		})
	}
}
