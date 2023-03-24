package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type ItsmIntegrationService service

type ResponseItsmIntegrationCreateItsmIntegrationSetting struct {
	ID                 string                                                                   `json:"id,omitempty"`                 // Id
	DypID              string                                                                   `json:"dypId,omitempty"`              // Dyp Id
	DypName            string                                                                   `json:"dypName,omitempty"`            // Dyp Name
	Name               string                                                                   `json:"name,omitempty"`               // Name
	UniqueKey          string                                                                   `json:"uniqueKey,omitempty"`          // Unique Key
	DypMajorVersion    *int                                                                     `json:"dypMajorVersion,omitempty"`    // Dyp Major Version
	Description        string                                                                   `json:"description,omitempty"`        // Description
	Data               *ResponseItsmIntegrationCreateItsmIntegrationSettingData                 `json:"data,omitempty"`               //
	CreatedDate        *int                                                                     `json:"createdDate,omitempty"`        // Created Date
	CreatedBy          string                                                                   `json:"createdBy,omitempty"`          // Created By
	UpdatedBy          string                                                                   `json:"updatedBy,omitempty"`          // Updated By
	SoftwareVersionLog *[]ResponseItsmIntegrationCreateItsmIntegrationSettingSoftwareVersionLog `json:"softwareVersionLog,omitempty"` // Software Version Log
	SchemaVersion      *float64                                                                 `json:"schemaVersion,omitempty"`      // Schema Version
	TenantID           string                                                                   `json:"tenantId,omitempty"`           // Tenant Id
}
type ResponseItsmIntegrationCreateItsmIntegrationSettingData struct {
	ConnectionSettings *ResponseItsmIntegrationCreateItsmIntegrationSettingDataConnectionSettings `json:"ConnectionSettings,omitempty"` //
}
type ResponseItsmIntegrationCreateItsmIntegrationSettingDataConnectionSettings struct {
	URL          string `json:"Url,omitempty"`           // Url
	AuthUserName string `json:"Auth_UserName,omitempty"` // Auth User Name
	AuthPassword string `json:"Auth_Password,omitempty"` // Auth Password
}
type ResponseItsmIntegrationCreateItsmIntegrationSettingSoftwareVersionLog interface{}
type ResponseItsmIntegrationUpdateItsmIntegrationSetting struct {
	TypeID          string                                                   `json:"_id,omitempty"`             // Id
	ID              string                                                   `json:"id,omitempty"`              // Id
	DypID           string                                                   `json:"dypId,omitempty"`           // Dyp Id
	DypName         string                                                   `json:"dypName,omitempty"`         // Dyp Name
	DypMajorVersion *int                                                     `json:"dypMajorVersion,omitempty"` // Dyp Major Version
	Name            string                                                   `json:"name,omitempty"`            // Name
	UniqueKey       string                                                   `json:"uniqueKey,omitempty"`       // Unique Key
	Description     string                                                   `json:"description,omitempty"`     // Description
	Data            *ResponseItsmIntegrationUpdateItsmIntegrationSettingData `json:"data,omitempty"`            //
	UpdatedDate     *int                                                     `json:"updatedDate,omitempty"`     // Updated Date
	UpdatedBy       string                                                   `json:"updatedBy,omitempty"`       // Updated By
	TenantID        string                                                   `json:"tenantId,omitempty"`        // Tenant Id
}
type ResponseItsmIntegrationUpdateItsmIntegrationSettingData struct {
	ConnectionSettings *ResponseItsmIntegrationUpdateItsmIntegrationSettingDataConnectionSettings `json:"ConnectionSettings,omitempty"` //
}
type ResponseItsmIntegrationUpdateItsmIntegrationSettingDataConnectionSettings struct {
	URL          string `json:"Url,omitempty"`           // Url
	AuthUserName string `json:"Auth_UserName,omitempty"` // Auth User Name
	AuthPassword string `json:"Auth_Password,omitempty"` // Auth Password
}
type ResponseItsmIntegrationGetItsmIntegrationSettingByID struct {
	TypeID          string                                                    `json:"_id,omitempty"`             // Id
	ID              string                                                    `json:"id,omitempty"`              // Id
	DypID           string                                                    `json:"dypId,omitempty"`           // Dyp Id
	DypName         string                                                    `json:"dypName,omitempty"`         // Dyp Name
	DypMajorVersion *int                                                      `json:"dypMajorVersion,omitempty"` // Dyp Major Version
	Name            string                                                    `json:"name,omitempty"`            // Name
	UniqueKey       string                                                    `json:"uniqueKey,omitempty"`       // Unique Key
	Description     string                                                    `json:"description,omitempty"`     // Description
	Data            *ResponseItsmIntegrationGetItsmIntegrationSettingByIDData `json:"data,omitempty"`            //
	UpdatedDate     *int                                                      `json:"updatedDate,omitempty"`     // Updated Date
	UpdatedBy       string                                                    `json:"updatedBy,omitempty"`       // Updated By
	TenantID        string                                                    `json:"tenantId,omitempty"`        // Tenant Id
}
type ResponseItsmIntegrationGetItsmIntegrationSettingByIDData struct {
	ConnectionSettings *ResponseItsmIntegrationGetItsmIntegrationSettingByIDDataConnectionSettings `json:"ConnectionSettings,omitempty"` //
}
type ResponseItsmIntegrationGetItsmIntegrationSettingByIDDataConnectionSettings struct {
	URL          string `json:"Url,omitempty"`           // Url
	AuthUserName string `json:"Auth_UserName,omitempty"` // Auth User Name
	AuthPassword string `json:"Auth_Password,omitempty"` // Auth Password
}
type ResponseItsmIntegrationGetAllItsmIntegrationSettings []ResponseItemItsmIntegrationGetAllItsmIntegrationSettings // Array of ResponseItsmIntegrationGetAllITSMIntegrationSettings
type ResponseItemItsmIntegrationGetAllItsmIntegrationSettings struct {
	ID                 string                                                                        `json:"id,omitempty"`                 // Id
	DypID              string                                                                        `json:"dypId,omitempty"`              // Dyp Id
	DypName            string                                                                        `json:"dypName,omitempty"`            // Dyp Name
	Name               string                                                                        `json:"name,omitempty"`               // Name
	UniqueKey          string                                                                        `json:"uniqueKey,omitempty"`          // Unique Key
	DypMajorVersion    *int                                                                          `json:"dypMajorVersion,omitempty"`    // Dyp Major Version
	Description        string                                                                        `json:"description,omitempty"`        // Description
	CreatedDate        *int                                                                          `json:"createdDate,omitempty"`        // Created Date
	CreatedBy          string                                                                        `json:"createdBy,omitempty"`          // Created By
	UpdatedBy          string                                                                        `json:"updatedBy,omitempty"`          // Updated By
	SoftwareVersionLog *[]ResponseItemItsmIntegrationGetAllItsmIntegrationSettingsSoftwareVersionLog `json:"softwareVersionLog,omitempty"` // Software Version Log
	SchemaVersion      *float64                                                                      `json:"schemaVersion,omitempty"`      // Schema Version
	TenantID           string                                                                        `json:"tenantId,omitempty"`           // Tenant Id
}
type ResponseItemItsmIntegrationGetAllItsmIntegrationSettingsSoftwareVersionLog interface{}
type RequestItsmIntegrationCreateItsmIntegrationSetting struct {
	Name        string                                                  `json:"name,omitempty"`        // Name of the setting instance
	Description string                                                  `json:"description,omitempty"` // Description of the setting instance
	Data        *RequestItsmIntegrationCreateItsmIntegrationSettingData `json:"data,omitempty"`        //
	DypName     string                                                  `json:"dypName,omitempty"`     // It should be ServiceNowConnection
}
type RequestItsmIntegrationCreateItsmIntegrationSettingData struct {
	ConnectionSettings *RequestItsmIntegrationCreateItsmIntegrationSettingDataConnectionSettings `json:"ConnectionSettings,omitempty"` //
}
type RequestItsmIntegrationCreateItsmIntegrationSettingDataConnectionSettings struct {
	URL          string `json:"Url,omitempty"`           // Url
	AuthUserName string `json:"Auth_UserName,omitempty"` // Auth User Name
	AuthPassword string `json:"Auth_Password,omitempty"` // Auth Password
}
type RequestItsmIntegrationUpdateItsmIntegrationSetting struct {
	Name        string                                                  `json:"name,omitempty"`        // Name of the setting instance
	Description string                                                  `json:"description,omitempty"` // Description of the setting instance
	Data        *RequestItsmIntegrationUpdateItsmIntegrationSettingData `json:"data,omitempty"`        //
	DypName     string                                                  `json:"dypName,omitempty"`     // It should be ServiceNowConnection
}
type RequestItsmIntegrationUpdateItsmIntegrationSettingData struct {
	ConnectionSettings *RequestItsmIntegrationUpdateItsmIntegrationSettingDataConnectionSettings `json:"ConnectionSettings,omitempty"` //
}
type RequestItsmIntegrationUpdateItsmIntegrationSettingDataConnectionSettings struct {
	URL          string `json:"Url,omitempty"`           // Url
	AuthUserName string `json:"Auth_UserName,omitempty"` // Auth User Name
	AuthPassword string `json:"Auth_Password,omitempty"` // Auth Password
}

//GetItsmIntegrationSettingByID Get ITSM Integration setting by Id - 1086-aa18-4cda-8471
/* Fetches ITSM Integration setting by ID


@param instanceID instanceId path parameter. Instance Id of the Integration setting instance

*/
func (s *ItsmIntegrationService) GetItsmIntegrationSettingByID(instanceID string) (*ResponseItsmIntegrationGetItsmIntegrationSettingByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/integration-settings/instances/itsm/{instanceId}"
	path = strings.Replace(path, "{instanceId}", fmt.Sprintf("%v", instanceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseItsmIntegrationGetItsmIntegrationSettingByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetItsmIntegrationSettingById")
	}

	result := response.Result().(*ResponseItsmIntegrationGetItsmIntegrationSettingByID)
	return result, response, err

}

//GetAllItsmIntegrationSettings Get all ITSM Integration settings - 468a-fb23-4379-86eb
/* Fetches all ITSM Integration settings


 */
func (s *ItsmIntegrationService) GetAllItsmIntegrationSettings() (*ResponseItsmIntegrationGetAllItsmIntegrationSettings, *resty.Response, error) {
	path := "/dna/intent/api/v1/integration-settings/itsm/instances"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseItsmIntegrationGetAllItsmIntegrationSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAllItsmIntegrationSettings")
	}

	result := response.Result().(*ResponseItsmIntegrationGetAllItsmIntegrationSettings)
	return result, response, err

}

//CreateItsmIntegrationSetting Create ITSM Integration setting - 0cb0-1a15-4d79-a440
/* Creates ITSM Integration setting


 */
func (s *ItsmIntegrationService) CreateItsmIntegrationSetting(requestItsmIntegrationCreateITSMIntegrationSetting *RequestItsmIntegrationCreateItsmIntegrationSetting) (*ResponseItsmIntegrationCreateItsmIntegrationSetting, *resty.Response, error) {
	path := "/dna/intent/api/v1/integration-settings/instances/itsm"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestItsmIntegrationCreateITSMIntegrationSetting).
		SetResult(&ResponseItsmIntegrationCreateItsmIntegrationSetting{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateItsmIntegrationSetting")
	}

	result := response.Result().(*ResponseItsmIntegrationCreateItsmIntegrationSetting)
	return result, response, err

}

//UpdateItsmIntegrationSetting Update ITSM Integration setting - 5fbe-680a-4208-92e6
/* Updates the ITSM Integration setting


@param instanceID instanceId path parameter. Instance Id of the Integration setting instance

*/
func (s *ItsmIntegrationService) UpdateItsmIntegrationSetting(instanceID string, requestItsmIntegrationUpdateITSMIntegrationSetting *RequestItsmIntegrationUpdateItsmIntegrationSetting) (*ResponseItsmIntegrationUpdateItsmIntegrationSetting, *resty.Response, error) {
	path := "/dna/intent/api/v1/integration-settings/instances/itsm/{instanceId}"
	path = strings.Replace(path, "{instanceId}", fmt.Sprintf("%v", instanceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestItsmIntegrationUpdateITSMIntegrationSetting).
		SetResult(&ResponseItsmIntegrationUpdateItsmIntegrationSetting{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateItsmIntegrationSetting")
	}

	result := response.Result().(*ResponseItsmIntegrationUpdateItsmIntegrationSetting)
	return result, response, err

}

//DeleteItsmIntegrationSetting Delete ITSM Integration setting - e8b9-ba8b-4c29-b637
/* Deletes the ITSM Integration setting


@param instanceID instanceId path parameter. Instance Id of the Integration setting instance

*/
func (s *ItsmIntegrationService) DeleteItsmIntegrationSetting(instanceID string) (*resty.Response, error) {
	//instanceID string
	path := "/dna/intent/api/v1/integration-settings/instances/itsm/{instanceId}"
	path = strings.Replace(path, "{instanceId}", fmt.Sprintf("%v", instanceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteItsmIntegrationSetting")
	}

	return response, err

}
