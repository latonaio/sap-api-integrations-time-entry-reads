package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema      string `json:"api_schema"`
	MaterialCode   string `json:"material_code"`
	Plant_Supplier string `json:"plant/supplier"`
	Stock          string `json:"stock"`
	DocumentType   string `json:"document_type"`
	DocumentNo     string `json:"document_no"`
	PlannedDate    string `json:"planned_date"`
	ValidatedDate  string `json:"validated_date"`
	Deleted        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey       string `json:"connection_key"`
	Result              bool   `json:"result"`
	RedisKey            string `json:"redis_key"`
	Filepath            string `json:"filepath"`
	TimeEntryCollection struct {
		ObjectID                            string `json:"ObjectID"`
		UUID                                string `json:"UUID"`
		InternalID                          string `json:"InternalID"`
		BusinessPartnerFormattedName        string `json:"BusinessPartnerFormattedName"`
		ServiceRequestID                    string `json:"ServiceRequestID"`
		AutoTimeRecordingIndicator          bool   `json:"AutoTimeRecordingIndicator"`
		BillingRelevanceIndicator           bool   `json:"BillingRelevanceIndicator"`
		Date                                string `json:"Date"`
		Duration                            string `json:"Duration"`
		EmployeeUUID                        string `json:"EmployeeUUID"`
		Description                         string `json:"Description"`
		LanguageCode                        string `json:"languageCode"`
		EndDateTime                         string `json:"EndDateTime"`
		HeaderReferenceUUID                 string `json:"HeaderReferenceUUID"`
		ID                                  string `json:"ID"`
		InformationLifeCycleStatusCode      string `json:"InformationLifeCycleStatusCode"`
		ItemReferenceUUID                   string `json:"ItemReferenceUUID"`
		BusinessTransactionDocumentTypeCode string `json:"BusinessTransactionDocumentTypeCode"`
		StartDateTime                       string `json:"StartDateTime"`
		LifeCycleStatusCode                 string `json:"LifeCycleStatusCode"`
		LastChangeDateTime                  string `json:"LastChangeDateTime"`
		TimeReportUUID                      string `json:"TimeReportUUID"`
		TimeTypeCode                        string `json:"TimeTypeCode"`
		TimeZone                            string `json:"TimeZone"`
		CreationDateTime                    string `json:"CreationDateTime"`
		LastChangeIdentityUUID              string `json:"LastChangeIdentityUUID"`
		CreationIdentityUUID                string `json:"CreationIdentityUUID"`
		EmployeeID                          string `json:"EmployeeID"`
		EntityLastChangedOn                 string `json:"EntityLastChangedOn"`
		ETag                                string `json:"ETag"`
	} `json:"TimeEntryCollection"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	TimeEntryCode string   `json:"time_entry_code"`
	Deleted       bool     `json:"deleted"`
}
