package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-time-entry-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetTimeEntry(iD string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "TimeEntryCollection":
			func() {
				c.TimeEntryCollection(iD)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
//timeentrycollection
func (c *SAPAPICaller) TimeEntryCollection(iD string) {
	timeEntryCollectionData, err := c.callTimeEntrySrvAPIRequirementTimeEntryCollection("TimeEntryCollection", iD)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(timeEntryCollectionData)
}

func (c *SAPAPICaller) callTimeEntrySrvAPIRequirementTimeEntryCollection(api, iD string) ([]sap_api_output_formatter.TimeEntryCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithTimeEntryCollection(req, iD)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToTimeEntryCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}
//ここまで
func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithTimeEntryCollection(req *http.Request, iD string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ID eq '%s'", iD))
	req.URL.RawQuery = params.Encode()
}