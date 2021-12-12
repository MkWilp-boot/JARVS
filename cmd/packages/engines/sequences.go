package engines

type Client int

const (
	Wikipedia Client = iota
	Google
	Youtube
	ExternalAPI
)
