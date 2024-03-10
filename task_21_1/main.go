package main

import "errors"

// Ожидаемый интерфейс, вызываевый клиентом
type Adaptee interface {
	SetPort(string)
}

// Адаптируемый класс, методы которого надо вызвать
type ServerInfo struct {
	url  string
	port string
}

func (s *ServerInfo) ChangePort(port string) bool {
	s.port = port
	return true
}

// "Оболочка" для адаптируемого класса, реализущая желаемый интерфейс
type Adapter struct {
	srv *ServerInfo
}

func (a *Adapter) SetPort(port string) {
	if !a.srv.ChangePort(port) {
		panic(errors.New("cannot change port"))
	}
}

func main() {
	var a Adaptee
	s := &ServerInfo{url: "localhost", port: "8080"}
	a = &Adapter{srv: s}
	a.SetPort("8081")
}
