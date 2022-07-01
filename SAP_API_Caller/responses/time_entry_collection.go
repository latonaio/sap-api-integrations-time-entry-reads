package responses

type TimeEntryCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
			TimeEntryItem                       struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"TimeEntryItem"`
		} `json:"results"`
	} `json:"d"`
}
