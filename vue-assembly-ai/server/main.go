package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	var err error = runServer()
	if err != nil {
		log.Println("server Error")
		os.Exit(2)
	}
}

func runServer() error {
	var ln, lnErr = net.Listen("tcp", ":42069")

	if lnErr != nil {
		return errors.New("tcp Listen Error")
	}

	for {
		var con, conErr = ln.Accept()
		if conErr != nil {
			continue
		}
		var _, dataErr = ioutil.ReadAll(con)
		if dataErr != nil {
			log.Println("data Parse Err")
			continue
		}

	}

	return nil
}
