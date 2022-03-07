package main

import (
	"fmt"

	"gofish/gofish"
)

func main() {
	config := gofish.ClientConfig{
		Endpoint: "https://10.10.1.102:443",
		Username: "root",
		Password: "calvin",
		Insecure: true,
	}
	// config := gofish.ClientConfig{
	// 	Endpoint: "https://10.10.1.47:443",
	// 	Username: "admin",
	// 	Password: "hpinvent",
	// 	Insecure: true,
	// }

	c, err := gofish.Connect(config)
	if err != nil {
		panic(err)
	}
	defer c.Logout()
	fmt.Printf("connect: %#v\n\n", c)

	service := c.Service

	/*chassis session token을 사용해 셰시 데이터를 쿼리한다.*/
	chassis, err := service.Chassis()
	if err != nil {
		panic(err)
	}

	fmt.Printf("chassis: %#v 개\n\n", len(chassis))

	cs, err := chassis[0].ComputerSystems()
	if err != nil {
		panic(err)
	}

	fmt.Printf("cs : %#v\n\n", len(cs))

	storage, err := cs[0].Storage()
	if err != nil {
		panic(err)
	}

	fmt.Printf("storage %#v개\n\n", len(storage))

	for _, st := range storage {
		fmt.Printf("%#v\n", st)
	}

	// fmt.Printf("storage: %#v\n\n", storage)
	// fmt.Printf("storage-controller: %#v\n\n", storage[0].StorageControllers)

	simplestorage, err := cs[0].SimpleStorages()
	if err != nil {
		panic(err)
	}

	fmt.Printf("simplestorage: %#v 개\n\n", len(simplestorage))

}
