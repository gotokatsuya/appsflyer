package rawdata

import (
	"testing"

	"github.com/gotokatsuya/appsflyer/dispatcher"
)

func TestGetInAppEventsReports(t *testing.T) {
	var (
		appID    = "xxx-xxx-xxx"
		fromDate = "2016-08-20"
		toDate   = "2016-08-21"
	)
	client := dispatcher.NewClient(appID, fromDate, toDate)
	client.SetOptionalParameter(dispatcher.OptionalParameter{
		Reattr: "true",
	})
	if v, err := GetInAppEventsReports(client); err == nil {
		t.Log("Passed", v)
	} else {
		t.Error("Failed", err)
	}
}
