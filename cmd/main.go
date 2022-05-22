package main

import (
	"f1-telemetry/pkg"
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Loading env config")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := pkg.GetEnv()

	fmt.Println("Setting up connections")
	client, api := pkg.NewDbConnection(env)
	defer client.Close()

	sAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:20777")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	sConn, err := net.ListenUDP("udp", sAddr)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer sConn.Close()

	buf := make([]byte, 1289)

	for {
		_, _, err := sConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Errorf("error: %s", err)
		}

		telemetry, _ := pkg.ReadPacket(buf)
		// fmt.Println("NEW PACKET")

		switch telemetry.Type {
		case pkg.PACKET_CAR_TELEMETRY:
			// fmt.Println("CAR DATA")
			pkg.SaveTelemetryCarData(api, telemetry.Data.(pkg.TelemetryCarData))
		case pkg.PACKET_LAP_DATA:
			// fmt.Println("LAP DATA")
			pkg.SaveTelemetryLapData(api, telemetry.Data.(pkg.TelemetryLapData))
			// default:
			// 	fmt.Println("OTHER")
		}
	}
}
