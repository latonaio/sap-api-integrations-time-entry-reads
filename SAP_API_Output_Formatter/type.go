package sap_api_output_formatter

type TimeEntry struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	APISchema       string `json:"api_schema"`
	TimeEntryCode   string `json:"time_entry_code"`
	Deleted         bool   `json:"deleted"`
}

type TimeEntryCollection struct {
	ObjectID                            string      `json:"ObjectID"`
	UUID                                string      `json:"UUID"`
	InternalID                          string      `json:"InternalID"`
	BusinessPartnerFormattedName        string      `json:"BusinessPartnerFormattedName"`
	ServiceRequestID                    string      `json:"ServiceRequestID"`
	AutoTimeRecordingIndicator          bool        `json:"AutoTimeRecordingIndicator"`
	BillingRelevanceIndicator           bool        `json:"BillingRelevanceIndicator"`
	Date                                string      `json:"Date"`
	Duration                            string      `json:"Duration"`
	EmployeeUUID                        string      `json:"EmployeeUUID"`
	Description                         string      `json:"Description"`
	LanguageCode                        string      `json:"languageCode"`
	EndDateTime                         string      `json:"EndDateTime"`
	HeaderReferenceUUID                 string      `json:"HeaderReferenceUUID"`
	ID                                  string      `json:"ID"`
	InformationLifeCycleStatusCode      string      `json:"InformationLifeCycleStatusCode"`
	ItemReferenceUUID                   string      `json:"ItemReferenceUUID"`
	BusinessTransactionDocumentTypeCode string      `json:"BusinessTransactionDocumentTypeCode"`
	StartDateTime                       string      `json:"StartDateTime"`
	LifeCycleStatusCode                 string      `json:"LifeCycleStatusCode"`
	LastChangeDateTime                  string      `json:"LastChangeDateTime"`
	TimeReportUUID                      string      `json:"TimeReportUUID"`
	TimeTypeCode                        string      `json:"TimeTypeCode"`
	TimeZone                            string      `json:"TimeZone"`
	CreationDateTime                    string      `json:"CreationDateTime"`
	LastChangeIdentityUUID              string      `json:"LastChangeIdentityUUID"`
	CreationIdentityUUID                string      `json:"CreationIdentityUUID"`
	EmployeeID                          string      `json:"EmployeeID"`
	EntityLastChangedOn                 string      `json:"EntityLastChangedOn"`
	ETag                                string      `json:"ETag"`
}