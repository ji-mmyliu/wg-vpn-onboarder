package models

type Network struct {
	Server  Server
	Clients []Client

	InterfaceName string
	ID            uint
	AddressPrefix string
	Address       string
	DnsServer     string
}

type Server struct {
	Creds      Creds
	ListenPort uint
	Address    string
	Endpoint   string
}

type Client struct {
	Creds   Creds
	ID      uint
	Address string

	PeerPublicKey  string
	PeerEndpoint   string
	PeerListenPort uint
	NetworkAddress string
	DnsServer      string
}

type Creds struct {
	PublicKey  string
	PrivateKey string
}
