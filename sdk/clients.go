package dnac

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// ClientsService is the service to communicate with the Clients API endpoint
type ClientsService service

// GetClientDetailResponse is the GetClientDetailResponse definition
type GetClientDetailResponse struct {
	ConnectionInfo struct {
		Band          string `json:"band,omitempty"`          //
		Channel       string `json:"channel,omitempty"`       //
		ChannelWidth  string `json:"channelWidth,omitempty"`  //
		HostType      string `json:"hostType,omitempty"`      //
		NwDeviceMac   string `json:"nwDeviceMac,omitempty"`   //
		NwDeviceName  string `json:"nwDeviceName,omitempty"`  //
		Protocol      string `json:"protocol,omitempty"`      //
		SpatialStream string `json:"spatialStream,omitempty"` //
		Timestamp     int    `json:"timestamp,omitempty"`     //
		Uapsd         string `json:"uapsd,omitempty"`         //
		Wmm           string `json:"wmm,omitempty"`           //
	} `json:"connectionInfo,omitempty"` //
	Detail struct {
		ApGroup          string   `json:"apGroup,omitempty"`          //
		AuthType         string   `json:"authType,omitempty"`         //
		AvgRssi          string   `json:"avgRssi,omitempty"`          //
		AvgSnr           string   `json:"avgSnr,omitempty"`           //
		Channel          string   `json:"channel,omitempty"`          //
		ClientConnection string   `json:"clientConnection,omitempty"` //
		ClientType       string   `json:"clientType,omitempty"`       //
		ConnectedDevice  []string `json:"connectedDevice,omitempty"`  //
		ConnectionStatus string   `json:"connectionStatus,omitempty"` //
		DataRate         string   `json:"dataRate,omitempty"`         //
		DNSFailure       string   `json:"dnsFailure,omitempty"`       //
		DNSSuccess       string   `json:"dnsSuccess,omitempty"`       //
		Frequency        string   `json:"frequency,omitempty"`        //
		HealthScore      []struct {
			HealthType string `json:"healthType,omitempty"` //
			Reason     string `json:"reason,omitempty"`     //
			Score      int    `json:"score,omitempty"`      //
		} `json:"healthScore,omitempty"` //
		HostIPV4    string   `json:"hostIpV4,omitempty"`    //
		HostIPV6    []string `json:"hostIpV6,omitempty"`    //
		HostMac     string   `json:"hostMac,omitempty"`     //
		HostName    string   `json:"hostName,omitempty"`    //
		HostOs      string   `json:"hostOs,omitempty"`      //
		HostType    string   `json:"hostType,omitempty"`    //
		HostVersion string   `json:"hostVersion,omitempty"` //
		ID          string   `json:"id,omitempty"`          //
		IosCapable  bool     `json:"iosCapable,omitempty"`  //
		IssueCount  int      `json:"issueCount,omitempty"`  //
		LastUpdated int      `json:"lastUpdated,omitempty"` //
		Location    string   `json:"location,omitempty"`    //
		Onboarding  struct {
			AaaRootcauseList     []string `json:"aaaRootcauseList,omitempty"`     //
			AaaServerIP          string   `json:"aaaServerIp,omitempty"`          //
			AssocDoneTime        string   `json:"assocDoneTime,omitempty"`        //
			AssocRootcauseList   []string `json:"assocRootcauseList,omitempty"`   //
			AuthDoneTime         string   `json:"authDoneTime,omitempty"`         //
			AverageAssocDuration string   `json:"averageAssocDuration,omitempty"` //
			AverageAuthDuration  string   `json:"averageAuthDuration,omitempty"`  //
			AverageDhcpDuration  string   `json:"averageDhcpDuration,omitempty"`  //
			AverageRunDuration   string   `json:"averageRunDuration,omitempty"`   //
			DhcpDoneTime         string   `json:"dhcpDoneTime,omitempty"`         //
			DhcpRootcauseList    []string `json:"dhcpRootcauseList,omitempty"`    //
			DhcpServerIP         string   `json:"dhcpServerIp,omitempty"`         //
			MaxAssocDuration     string   `json:"maxAssocDuration,omitempty"`     //
			MaxAuthDuration      string   `json:"maxAuthDuration,omitempty"`      //
			MaxDhcpDuration      string   `json:"maxDhcpDuration,omitempty"`      //
			MaxRunDuration       string   `json:"maxRunDuration,omitempty"`       //
			OtherRootcauseList   []string `json:"otherRootcauseList,omitempty"`   //
		} `json:"onboarding,omitempty"` //
		OnboardingTime string `json:"onboardingTime,omitempty"` //
		Port           string `json:"port,omitempty"`           //
		Rssi           string `json:"rssi,omitempty"`           //
		RxBytes        string `json:"rxBytes,omitempty"`        //
		Snr            string `json:"snr,omitempty"`            //
		SSID           string `json:"ssid,omitempty"`           //
		SubType        string `json:"subType,omitempty"`        //
		TxBytes        string `json:"txBytes,omitempty"`        //
		UserID         string `json:"userId,omitempty"`         //
		VLANID         string `json:"vlanId,omitempty"`         //
		Vnid           string `json:"vnid,omitempty"`           //
	} `json:"detail,omitempty"` //
	Topology struct {
		Links []struct {
			ID              string   `json:"id,omitempty"`              //
			Label           []string `json:"label,omitempty"`           //
			LinkStatus      string   `json:"linkStatus,omitempty"`      //
			PortUtilization string   `json:"portUtilization,omitempty"` //
			Source          string   `json:"source,omitempty"`          //
			Target          string   `json:"target,omitempty"`          //
		} `json:"links,omitempty"` //
		Nodes []struct {
			Clients         string `json:"clients,omitempty"`         //
			ConnectedDevice string `json:"connectedDevice,omitempty"` //
			Count           string `json:"count,omitempty"`           //
			Description     string `json:"description,omitempty"`     //
			DeviceType      string `json:"deviceType,omitempty"`      //
			FabricGroup     string `json:"fabricGroup,omitempty"`     //
			Family          string `json:"family,omitempty"`          //
			HealthScore     int    `json:"healthScore,omitempty"`     //
			ID              string `json:"id,omitempty"`              //
			IP              string `json:"ip,omitempty"`              //
			Level           int    `json:"level,omitempty"`           //
			Name            string `json:"name,omitempty"`            //
			NodeType        string `json:"nodeType,omitempty"`        //
			PlatformID      string `json:"platformId,omitempty"`      //
			RadioFrequency  string `json:"radioFrequency,omitempty"`  //
			Role            string `json:"role,omitempty"`            //
			SoftwareVersion string `json:"softwareVersion,omitempty"` //
			UserID          string `json:"userId,omitempty"`          //
		} `json:"nodes,omitempty"` //
	} `json:"topology,omitempty"` //
}

// GetClientEnrichmentDetailsResponse is the GetClientEnrichmentDetailsResponse definition
type GetClientEnrichmentDetailsResponse struct {
	ConnectedDevice []struct {
		DeviceDetails struct {
			ApManagerInterfaceIP  string `json:"apManagerInterfaceIp,omitempty"`  //
			AssociatedWlcIP       string `json:"associatedWlcIp,omitempty"`       //
			BootDateTime          string `json:"bootDateTime,omitempty"`          //
			Cisco360view          string `json:"cisco360view,omitempty"`          //
			CollectionInterval    string `json:"collectionInterval,omitempty"`    //
			CollectionStatus      string `json:"collectionStatus,omitempty"`      //
			ErrorCode             string `json:"errorCode,omitempty"`             //
			ErrorDescription      string `json:"errorDescription,omitempty"`      //
			Family                string `json:"family,omitempty"`                //
			Hostname              string `json:"hostname,omitempty"`              //
			ID                    string `json:"id,omitempty"`                    //
			InstanceUUID          string `json:"instanceUuid,omitempty"`          //
			InterfaceCount        string `json:"interfaceCount,omitempty"`        //
			InventoryStatusDetail string `json:"inventoryStatusDetail,omitempty"` //
			LastUpdateTime        int    `json:"lastUpdateTime,omitempty"`        //
			LastUpdated           string `json:"lastUpdated,omitempty"`           //
			LineCardCount         string `json:"lineCardCount,omitempty"`         //
			LineCardID            string `json:"lineCardId,omitempty"`            //
			Location              string `json:"location,omitempty"`              //
			LocationName          string `json:"locationName,omitempty"`          //
			MacAddress            string `json:"macAddress,omitempty"`            //
			ManagementIPAddress   string `json:"managementIpAddress,omitempty"`   //
			MemorySize            string `json:"memorySize,omitempty"`            //
			NeighborTopology      []struct {
				Links []struct {
					ID              string   `json:"id,omitempty"`              //
					Label           []string `json:"label,omitempty"`           //
					LinkStatus      string   `json:"linkStatus,omitempty"`      //
					PortUtilization string   `json:"portUtilization,omitempty"` //
					Source          string   `json:"source,omitempty"`          //
					Target          string   `json:"target,omitempty"`          //
				} `json:"links,omitempty"` //
				Nodes []struct {
					Clients         int    `json:"clients,omitempty"`         //
					Count           string `json:"count,omitempty"`           //
					Description     string `json:"description,omitempty"`     //
					DeviceType      string `json:"deviceType,omitempty"`      //
					FabricGroup     string `json:"fabricGroup,omitempty"`     //
					Family          string `json:"family,omitempty"`          //
					HealthScore     string `json:"healthScore,omitempty"`     //
					ID              string `json:"id,omitempty"`              //
					IP              string `json:"ip,omitempty"`              //
					Level           int    `json:"level,omitempty"`           //
					Name            string `json:"name,omitempty"`            //
					NodeType        string `json:"nodeType,omitempty"`        //
					PlatformID      string `json:"platformId,omitempty"`      //
					RadioFrequency  string `json:"radioFrequency,omitempty"`  //
					Role            string `json:"role,omitempty"`            //
					SoftwareVersion string `json:"softwareVersion,omitempty"` //
					UserID          string `json:"userId,omitempty"`          //
				} `json:"nodes,omitempty"` //
			} `json:"neighborTopology,omitempty"` //
			PlatformID                string `json:"platformId,omitempty"`                //
			ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` //
			ReachabilityStatus        string `json:"reachabilityStatus,omitempty"`        //
			Role                      string `json:"role,omitempty"`                      //
			RoleSource                string `json:"roleSource,omitempty"`                //
			SerialNumber              string `json:"serialNumber,omitempty"`              //
			Series                    string `json:"series,omitempty"`                    //
			SNMPContact               string `json:"snmpContact,omitempty"`               //
			SNMPLocation              string `json:"snmpLocation,omitempty"`              //
			SoftwareVersion           string `json:"softwareVersion,omitempty"`           //
			TagCount                  string `json:"tagCount,omitempty"`                  //
			TunnelUDPPort             string `json:"tunnelUdpPort,omitempty"`             //
			Type                      string `json:"type,omitempty"`                      //
			UpTime                    string `json:"upTime,omitempty"`                    //
			WaasDeviceMode            string `json:"waasDeviceMode,omitempty"`            //
		} `json:"deviceDetails,omitempty"` //
	} `json:"connectedDevice,omitempty"` //
	IssueDetails struct {
		Issue []struct {
			ImpactedHosts []struct {
				ConnectedInterface string `json:"connectedInterface,omitempty"` //
				FailedAttempts     int    `json:"failedAttempts,omitempty"`     //
				HostName           string `json:"hostName,omitempty"`           //
				HostOs             string `json:"hostOs,omitempty"`             //
				HostType           string `json:"hostType,omitempty"`           //
				Location           struct {
					ApsImpacted []string `json:"apsImpacted,omitempty"` //
					Area        string   `json:"area,omitempty"`        //
					Building    string   `json:"building,omitempty"`    //
					Floor       string   `json:"floor,omitempty"`       //
					SiteID      string   `json:"siteId,omitempty"`      //
					SiteType    string   `json:"siteType,omitempty"`    //
				} `json:"location,omitempty"` //
				MacAddress string `json:"macAddress,omitempty"` //
				SSID       string `json:"ssid,omitempty"`       //
				Timestamp  int    `json:"timestamp,omitempty"`  //
			} `json:"impactedHosts,omitempty"` //
			IssueCategory    string `json:"issueCategory,omitempty"`    //
			IssueDescription string `json:"issueDescription,omitempty"` //
			IssueEntity      string `json:"issueEntity,omitempty"`      //
			IssueEntityValue string `json:"issueEntityValue,omitempty"` //
			IssueID          string `json:"issueId,omitempty"`          //
			IssueName        string `json:"issueName,omitempty"`        //
			IssuePriority    string `json:"issuePriority,omitempty"`    //
			IssueSeverity    string `json:"issueSeverity,omitempty"`    //
			IssueSource      string `json:"issueSource,omitempty"`      //
			IssueSummary     string `json:"issueSummary,omitempty"`     //
			IssueTimestamp   int    `json:"issueTimestamp,omitempty"`   //
			SuggestedActions []struct {
				Message string   `json:"message,omitempty"` //
				Steps   []string `json:"steps,omitempty"`   //
			} `json:"suggestedActions,omitempty"` //
		} `json:"issue,omitempty"` //
	} `json:"issueDetails,omitempty"` //
	UserDetails struct {
		AuthType         string   `json:"authType,omitempty"`         //
		ClientConnection string   `json:"clientConnection,omitempty"` //
		ConnectedDevice  []string `json:"connectedDevice,omitempty"`  //
		ConnectionStatus string   `json:"connectionStatus,omitempty"` //
		DataRate         string   `json:"dataRate,omitempty"`         //
		HealthScore      []struct {
			HealthType string `json:"healthType,omitempty"` //
			Reason     string `json:"reason,omitempty"`     //
			Score      int    `json:"score,omitempty"`      //
		} `json:"healthScore,omitempty"` //
		HostIPV4    string   `json:"hostIpV4,omitempty"`    //
		HostIPV6    []string `json:"hostIpV6,omitempty"`    //
		HostMac     string   `json:"hostMac,omitempty"`     //
		HostName    string   `json:"hostName,omitempty"`    //
		HostOs      string   `json:"hostOs,omitempty"`      //
		HostType    string   `json:"hostType,omitempty"`    //
		HostVersion string   `json:"hostVersion,omitempty"` //
		ID          string   `json:"id,omitempty"`          //
		IssueCount  int      `json:"issueCount,omitempty"`  //
		LastUpdated int      `json:"lastUpdated,omitempty"` //
		Location    string   `json:"location,omitempty"`    //
		Port        string   `json:"port,omitempty"`        //
		Rssi        string   `json:"rssi,omitempty"`        //
		Snr         string   `json:"snr,omitempty"`         //
		SSID        string   `json:"ssid,omitempty"`        //
		SubType     string   `json:"subType,omitempty"`     //
		UserID      string   `json:"userId,omitempty"`      //
		VLANID      string   `json:"vlanId,omitempty"`      //
	} `json:"userDetails,omitempty"` //
}

// GetOverallClientHealthResponse is the GetOverallClientHealthResponse definition
type GetOverallClientHealthResponse struct {
	Response []struct {
		ScoreDetail []struct {
			ClientCount       int `json:"clientCount,omitempty"`       //
			ClientUniqueCount int `json:"clientUniqueCount,omitempty"` //
			Endtime           int `json:"endtime,omitempty"`           //
			ScoreCategory     struct {
				ScoreCategory string `json:"scoreCategory,omitempty"` //
				Value         string `json:"value,omitempty"`         //
			} `json:"scoreCategory,omitempty"` //
			ScoreList []struct {
				ClientCount       int `json:"clientCount,omitempty"`       //
				ClientUniqueCount int `json:"clientUniqueCount,omitempty"` //
				Endtime           int `json:"endtime,omitempty"`           //
				ScoreCategory     struct {
					ScoreCategory string `json:"scoreCategory,omitempty"` //
					Value         string `json:"value,omitempty"`         //
				} `json:"scoreCategory,omitempty"` //
				ScoreList []struct {
					ClientCount       int    `json:"clientCount,omitempty"`       //
					ClientUniqueCount string `json:"clientUniqueCount,omitempty"` //
					Endtime           int    `json:"endtime,omitempty"`           //
					ScoreCategory     struct {
						ScoreCategory string `json:"scoreCategory,omitempty"` //
						Value         string `json:"value,omitempty"`         //
					} `json:"scoreCategory,omitempty"` //
					ScoreValue int `json:"scoreValue,omitempty"` //
					Starttime  int `json:"starttime,omitempty"`  //
				} `json:"scoreList,omitempty"` //
				ScoreValue int `json:"scoreValue,omitempty"` //
				Starttime  int `json:"starttime,omitempty"`  //
			} `json:"scoreList,omitempty"` //
			ScoreValue int `json:"scoreValue,omitempty"` //
			Starttime  int `json:"starttime,omitempty"`  //
		} `json:"scoreDetail,omitempty"` //
		SiteID string `json:"siteId,omitempty"` //
	} `json:"response,omitempty"` //
}

// GetClientDetailQueryParams defines the query parameters for this request
type GetClientDetailQueryParams struct {
	Timestamp  string `url:"timestamp,omitempty"`  // Epoch time(in milliseconds) when the Client health data is required
	MacAddress string `url:"macAddress,omitempty"` // MAC Address of the client
}

// GetClientDetail getClientDetail
/* Returns detailed Client information retrieved by Mac Address for any given point of time.
@param timestamp Epoch time(in milliseconds) when the Client health data is required
@param macAddress MAC Address of the client
*/
func (s *ClientsService) GetClientDetail(getClientDetailQueryParams *GetClientDetailQueryParams) (*GetClientDetailResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/client-detail"

	queryString, _ := query.Values(getClientDetailQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetClientDetailResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetClientDetailResponse)
	return result, response, err
}

// GetClientEnrichmentDetails getClientEnrichmentDetails
/* Enriches a given network End User context (a network user-id or end user’s device Mac Address) with details about the user, the devices that the user is connected to and the assurance issues that the user is impacted by
@param entity_type Client enrichment details can be fetched based on either User ID or Client MAC address. This parameter value must either be network_user_id/mac_address
@param entity_value Contains the actual value for the entity type that has been defined
@param issueCategory The category of the DNA event based on which the underlying issues need to be fetched
*/
func (s *ClientsService) GetClientEnrichmentDetails() (*GetClientEnrichmentDetailsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/client-enrichment-details"

	response, err := RestyClient.R().
		SetResult(&GetClientEnrichmentDetailsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetClientEnrichmentDetailsResponse)
	return result, response, err
}

// GetOverallClientHealthQueryParams defines the query parameters for this request
type GetOverallClientHealthQueryParams struct {
	Timestamp string `url:"timestamp,omitempty"` // Epoch time(in milliseconds) when the Client health data is required
}

// GetOverallClientHealth getOverallClientHealth
/* Returns Overall Client Health information by Client type (Wired and Wireless) for any given point of time
@param timestamp Epoch time(in milliseconds) when the Client health data is required
*/
func (s *ClientsService) GetOverallClientHealth(getOverallClientHealthQueryParams *GetOverallClientHealthQueryParams) (*GetOverallClientHealthResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/client-health"

	queryString, _ := query.Values(getOverallClientHealthQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetOverallClientHealthResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetOverallClientHealthResponse)
	return result, response, err
}
