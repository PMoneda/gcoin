package hub

type apiResponse struct {
	Message string `json:"message"`
}
type serviceDescription struct {
	DNS       string `json:"dns"`
	Port      uint   `json:"port"`
	Protocol  string `json:"protocol"`
	Version   string `json:"version"`
	Type      string `json:"type"`
	Region    string `json:"region"`
	PublicKey string `json:"public_key"`
}
