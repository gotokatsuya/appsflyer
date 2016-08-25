package rawdata

import (
	"github.com/gotokatsuya/appsflyer/dispatcher"
	"github.com/gotokatsuya/appsflyer/model/rawdata"
	"github.com/gotokatsuya/appsflyer/util/csv"
)

func GetInstallsReports(client *dispatcher.Client) ([]rawdata.Report, error) {
	body, err := client.DispatchGetRequest("installs_report/v5")
	if err != nil {
		return nil, err
	}
	var entities []rawdata.Report
	if err := csv.Parse(string(body), rawdata.Report{}, func(v interface{}) {
		entities = append(entities, v.(rawdata.Report))
	}); err != nil {
		return nil, err
	}
	return entities, nil
}
