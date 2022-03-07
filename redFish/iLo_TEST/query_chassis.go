/*
	해당 소스는 섀시 정보 조회 입니다.
*/

package main

import (
	"fmt"

	"github.com/stmcginnis/gofish"
)

func main() {
	config := gofish.ClientConfig{
		Endpoint: "https://10.10.1.47:443",
		Username: "admin",
		Password: "hpinvent",
		Insecure: true,
	}

	c, err := gofish.Connect(config)
	if err != nil {
		panic(err)
	}
	defer c.Logout()
	fmt.Printf("connect: %#v\n\n", c)

	service := c.Service

	// fmt.Printf("Oem: %#v\n\n", service.Oem)

	/*chassis session token을 사용해 셰시 데이터를 쿼리한다.*/
	chassis, err := service.Chassis()
	if err != nil {
		panic(err)
	}

	//조회된 정보들 for문으로 돌면서 print
	for _, chass := range chassis {
		fmt.Printf("Chassis: %#v\n\n", chass)
	}

	//fmt.Println("````")
	/*power: 전력 관련 자원 출력*/
	// power, err := chassis[0].Power()
	// if power != nil {
	// 	fmt.Printf("전력 정보: %#v\n\n", power)
	// 	fmt.Printf("전력 제어기: %#v\n\n", power.PowerControl)
	// 	fmt.Printf("전력 제어기 수: %#v\n\n", power.PowerControlCount)
	// 	fmt.Printf("전력 공급량: %#v\n\n", power.PowerSupplies)
	// 	fmt.Printf("전력 공급량 이름: %#v\n\n", power.PowerSupplies[0].Name)
	// 	fmt.Printf("전력 공급량 제품 일련 번호: %#v\n\n", power.PowerSupplies[0].SerialNumber)
	// }
	// fmt.Println("````")

	// computersystem, err := chassis[0].ComputerSystems()
	// if err != nil {
	// 	panic(err)
	// }

	// for _, cs := range computersystem {
	// 	//컴퓨터 시스템안에 있는 정보들: ODataType, ODataContext, networkInterfaces, memory, pcidevices 등
	// 	fmt.Printf("컴퓨터 시스템 정보 전체 출력: %#v\n\n", cs)
	// }

	/*위에 cs를 갖고 정보를 갖고와보자*/
	// memory, err := computersystem[0].Memory()
	// if err != nil {
	// 	panic(err)
	// }

	// /*메모리*/
	// for _, mem := range memory {
	// 	fmt.Printf("컴퓨터 시스템 메모리: %#v\n\n", memory)
	// }

	// process, err := computersystem[0].Processors()
	// if err != nil {
	// 	panic(err)
	// }

	// for _, po := range process {
	// 	/*
	// 		프로세스는 총 2개이다. 정보: pcieVirtualFunctions, Model, ProcessType, Status
	// 	*/
	// 	fmt.Printf("프로세스 전체 출력: %#v\n\n", po)
	// }

	// fmt.Printf("프로세스 타입: %#v\n\n", process[0].ProcessorType)
	// fmt.Printf("소켓: %#v\n\n", process[0].Socket)
	// fmt.Printf("총 코어 개수: %#v\n\n", process[0].TotalCores)

	nas, _ := chassis[0].NetworkAdapters()
	fmt.Printf("NetworkAdapters : %#v\n\n", len(nas))

	for i, na := range nas {
		fmt.Printf("NetworkAdapter %#v", i+1)
		fmt.Println("{")
		fmt.Printf("\t이름 : %#v\n", na.Name)
		fmt.Printf("\t모델명 : %#v\n", na.Model)
		fmt.Printf("\t제조사명 : %#v\n", na.Manufacturer)
		fmt.Printf("\t시리얼넘버 : %#v\n", na.SerialNumber)
		fmt.Printf("\t설명 : %#v\n", na.Description)
		fmt.Printf("\t상태 : %#v\n", na.Status)
		fmt.Println("}")
	}

}
