package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-time-entry-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToTimeEntryCollection(raw []byte, l *logger.Logger) ([]TimeEntryCollection, error) {
	pm := &responses.TimeEntryCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to TimeEntryCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	timeEntryCollection := make([]TimeEntryCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		timeEntryCollection = append(timeEntryCollection, TimeEntryCollection{
			ObjectID:                            data.ObjectID,
			UUID:                                data.UUID,
			InternalID:                          data.InternalID,
			BusinessPartnerFormattedName:        data.BusinessPartnerFormattedName,
			ServiceRequestID:                    data.ServiceRequestID,
			AutoTimeRecordingIndicator:          data.AutoTimeRecordingIndicator,
			BillingRelevanceIndicator:           data.BillingRelevanceIndicator,
			Date:                                data.Date,
			Duration:                            data.Duration,
			EmployeeUUID:                        data.EmployeeUUID,
			Description:                         data.Description,
			LanguageCode:                        data.LanguageCode,
			EndDateTime:                         data.EndDateTime,
			HeaderReferenceUUID:                 data.HeaderReferenceUUID,
			ID:                                  data.ID,
			InformationLifeCycleStatusCode:      data.InformationLifeCycleStatusCode,
			ItemReferenceUUID:                   data.ItemReferenceUUID,
			BusinessTransactionDocumentTypeCode: data.BusinessTransactionDocumentTypeCode,
			StartDateTime:                       data.StartDateTime,
			LifeCycleStatusCode:                 data.LifeCycleStatusCode,
			LastChangeDateTime:                  data.LastChangeDateTime,
			TimeReportUUID:                      data.TimeReportUUID,
			TimeTypeCode:                        data.TimeTypeCode,
			TimeZone:                            data.TimeZone,
			CreationDateTime:                    data.CreationDateTime,
			LastChangeIdentityUUID:              data.LastChangeIdentityUUID,
			CreationIdentityUUID:                data.CreationIdentityUUID,
			EmployeeID:                          data.EmployeeID,
			EntityLastChangedOn:                 data.EntityLastChangedOn,
			ETag:                                data.ETag,                    
		})
	}

	return timeEntryCollection, nil
}