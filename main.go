package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"strconv"

	"github.com/koodinikkarit/arttu/arttu_service"
	"google.golang.org/grpc"
)

func main() {
	var (
		serverAddr = flag.String("server_addr", "localhost:3112", "Serverin ip ja portti")
	)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		fmt.Println("Virhe yhteydessa")
	}

	client := ArttuService.NewArttuClient(conn)

	var arttuId int

	if _, err := os.Stat("a"); err == nil {
		c, _ := ioutil.ReadFile("a")
		arttuId, _ = strconv.Atoi(string(c))
		fmt.Println("arttuId", arttuId)
	}

	if arttuId == 0 {
		res, _ := client.Register(context.Background(), &ArttuService.RegisterRequest{})
		arttuId := int(res.ArttuId)
		ioutil.WriteFile("a", []byte(strconv.Itoa(arttuId)), 0644)

		ifaces, _ := net.Interfaces()

		addInterfaceClient, _ := client.AddInterface(context.Background())
		for _, iface := range ifaces {
			fmt.Println(iface)
			addInterfaceClient.Send(&ArttuService.AddInterfaceRequest{
				ArttuId: uint32(arttuId),
				Interface: &ArttuService.Interface{
					Name:  iface.Name,
					Mac:   iface.HardwareAddr,
					Index: uint32(iface.Index),
					Mtu:   uint32(iface.MTU),
					Flags: uint32(iface.Flags),
				},
			})
		}
	}

	go shutdown(client, arttuId)

	for {
	}
	//fmt.Println("lopetus", res)
}

func shutdown(client ArttuService.ArttuClient, aruttuId int) {
	res, err := client.ShutdownComputer(context.Background(), &ArttuService.ShutdownComputerRequest{ArttuId: 1})
	if err == nil {
		fmt.Println("halt")
		if res.Delay == 1 {
			fmt.Println("halt")
			//cmd := exec.Command("halt")
			//cmd.Run()
		}
	}
}

func update(client ArttuService.ArttuClient, arttuId int) {

}
