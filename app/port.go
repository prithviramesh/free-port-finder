package app

import (
	"net"
)


func FindPort() (port int, err error) {

	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return -1, err
	}

	defer l.Close()
	//Addr.Port cast found in Golang Nuts google group: https://groups.google.com/forum/#!topic/golang-nuts/JLzchxXm5Vs
	return l.Addr().(*net.TCPAddr).Port, nil
}

func CheckPort(port string) (bound bool, err error) {

	l, err := net.Listen("tcp", ":"+port)

	if err != nil {
		//port is probably taken
		return false, err
	}

	defer l.Close()
	return true, nil
}
