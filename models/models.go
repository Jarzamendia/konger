package models

//RESPONSES

// ServiceInfoResponse type details.
type ServiceInfoResponse struct {
	err      bool
	infoType string
	result   []ServiceInfo
}

//RouteInfoResponse type details.
type RouteInfoResponse struct {
	err      bool
	infoType string
	result   []RouteInfo
}

//ConsumersInfoResponse type details.
type ConsumersInfoResponse struct {
	err      bool
	infoType string
	result   []ConsumersInfo
}

//PluginInfoResponse type details.
type PluginInfoResponse struct {
	err      bool
	infoType string
	result   []PluginInfo
}

//TYPES

// ServiceInfo type details.
type ServiceInfo struct {
	ID       string // ID do serviço.
	Name     string // Nome do serviço.
	Protocol string // Protocolo usado (HTTP, HTTPS).
	Hosts    string // Host usado (xxx.com.br).
	Port     int    // Porta usada (80, 8080).
	Path     string // Path usado (URI).
	Routes   []RouteInfo
}

//RouteInfo type details.
type RouteInfo struct {
	ID           string   //ID do Route
	Hosts        []string //Hosts configurados
	Paths        []string //Paths acessiveis.
	Protocols    string
	Methods      string
	StripPath    bool
	PreserveHost bool
	ServiceID    string //ID do service vinculado.
}

//PluginInfo type details.
type PluginInfo struct {
	ID         string
	Name       string
	Enabled    bool
	ServiceID  string
	RouteID    string
	ConsumerID string
	Config     map[string]interface{}
}

//ConsumersInfo type details.
type ConsumersInfo struct {
	ID       string // ID do consumer.
	Username string // Username do consumer.
}
