package model

type Book struct {
	Dict []Dictionary
}

type ActionFn func() string

type Dictionary struct {
	Response    []string
	Inputs      []string
	Action      string
	CommandType string
	Priority    int
}
