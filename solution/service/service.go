package service

//introduce this package to break the import cycle present in the problem
//but most importantly using interfaces to pass logic

type Service struct {
	Client Client
}

type Client interface {
	Execute(uri string, method string, payload []byte)
}