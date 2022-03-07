// package main

// import (
// 	"fmt"
// 	"gofish/gofish"
// )

// func main() {

// 	config := gofish.ClientConfig{
// 		Endpoint: "https://10.10.1.102:443",
// 		Username: "root ",
// 		Password: "calvin",
// 		Insecure: true,
// 	}

// 	c, err := gofish.Connect(config)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer c.Logout()
// 	fmt.Printf("connect: %#v\n\n", c)

// 	service := c.Service

// 	/*chassis session token을 사용해 셰시 데이터를 쿼리한다.*/
// 	chassis, err := service.Chassis()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("chassis: %#v\n\n", chassis)

// 	cs, err := chassis[0].ComputerSystems()
// 	if err != nil {
// 		panic(err)
// 	}

// 	storage, err := cs[0].Storage()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("storage: %#v\n\n", storage)
// 	fmt.Printf("storage-controller: %#v\n\n", storage[0].StorageControllers)

// 	simplestorage, err := cs[0].SimpleStorages()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("simplestorage: %#v\n\n", simplestorage[0].Status.Health)

// }
