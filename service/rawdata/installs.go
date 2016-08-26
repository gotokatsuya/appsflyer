package rawdata

import (
	"github.com/gotokatsuya/appsflyer/dispatcher"
	"github.com/gotokatsuya/appsflyer/model/rawdata"
	"github.com/gotokatsuya/appsflyer/util/csv"
)

const endpointInstallReport = "installs_report/v5"

func GetInstallsReports(client *dispatcher.Client) ([]rawdata.Report, error) {
	body, err := client.DispatchGetRequest(endpointInstallReport)
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

func GetEachInstallsReport(client *dispatcher.Client, f func(report rawdata.Report)) error {
	body, err := client.DispatchGetRequest(endpointInstallReport)
	if err != nil {
		return err
	}
	if err := csv.Parse(string(body), rawdata.Report{}, func(v interface{}) {
		f(v.(rawdata.Report))
	}); err != nil {
		return err
	}
	return nil
}
