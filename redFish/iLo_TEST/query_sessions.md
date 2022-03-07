/*
	해당 소스 파일은 gofish의 커넥션 내용에 대한 것이다.
*/

package main

import (
	"fmt"

	"github.com/stmcginnis/gofish"
)

func main() {
	// Create a new instance of gofish client, ignoring self-signed certs
	config := gofish.ClientConfig{
		Endpoint: "https://10.10.1.47:443",
		Username: "hrgomp",
		Password: "hpinvent",
		Insecure: true,
	}
	c, err := gofish.Connect(config)
	if err != nil {
		panic(err)
	}
	defer c.Logout()

	// Retrieve the service root
	service := c.Service

	// Query the active sessions using the session token
	sessions, err := service.Sessions()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", sessions)

	for _, session := range sessions {
		fmt.Printf("Sessions: %#v\n\n", session)
	}

}
