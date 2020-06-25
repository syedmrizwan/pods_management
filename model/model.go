package model

type ConfigurationInfo struct {
	VcenterId     int64  `json:"vcenter_id"`
	IpAddress     string `json:"ip_address"`
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
	DatastoreID   int64  `json:"datastore_id"`
	DatastoreName string `json:"datastore_name"`
	ClusterID     int64  `json:"cluster_id"`
	ClusterName   string `json:"cluster_name"`
	TemplateName  string `json:"template_name"`
	TypeName      string `json:"type_name"`
}

type PodConfiguration struct {
	PodID         int64              `json:"pod_id"`
	PodName       string             `json:"pod_name"`
	Configuration *ConfigurationInfo `json:"configuration"`
	Status        string             `json:"status"`
}

type PodBody struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	IpAddress string `json:"ip_address"`
}

type Response struct {
	Message string `json:"message"`
}