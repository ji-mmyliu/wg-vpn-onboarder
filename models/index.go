package models

type Network struct {
	Server  Server
	Clients []Client

	InterfaceName string
	ID            uint
	Address       string
}

type Server struct {
	Creds      Creds
	ListenPort int
}

type Client struct {
	Creds   Creds
	ID      uint
	Address string
}

type Creds struct {
	PublicKey  string
	PrivateKey string
}
