package main

import (
	"bytes"
	"encoding/binary"
	"f1-telemetry/infrastructure"
	"fmt"
	"log"
	"math"
	"net"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/joho/godotenv"
)

type PacketHeader struct {
	PacketFormat            uint16  // 2020
	GameMajorVersion        uint8   // Game major version - "X.00"
	GameMinorVersion        uint8   // Game minor version - "1.XX"
	PacketVersion           uint8   // Version of this packet type, all start from 1
	PacketID                uint8   // Identifier for the packet type, see below
	SessionUID              uint64  // Unique identifier for the session
	SessionTime             float32 // Session timestamp
	FrameIdentifier         uint32  // Identifier for the frame the data was retrieved on
	PlayerCarIndex          uint8   // Index of player's car in the array
	SecondaryPlayerCarIndex uint8   // Index of secondary player's car in the array (split screen)
}

type PacketCarTelemetry struct {
	Speed    uint16
	Throttle float32
	Brake    float32
	Steering float32
}

type PacketLapData struct {
	LastTime uint32
}

func ReadPacket(buf []byte, pack interface{}) error {
	reader := bytes.NewReader(buf)
	if err := binary.Read(reader, binary.LittleEndian, pack); err != nil {
		return err
	}

	return nil
}

func ReadCarTelemetryPacket(buf []byte) PacketCarTelemetry {
	speed := binary.LittleEndian.Uint16(buf[24:26])
	throttle := math.Float32frombits(binary.LittleEndian.Uint32(buf[26:30]))
	steering := math.Float32frombits(binary.LittleEndian.Uint32(buf[30:34]))
	brake := math.Float32frombits(binary.LittleEndian.Uint32(buf[34:38]))

	return PacketCarTelemetry{
		Speed:    speed,
		Throttle: throttle,
		Brake:    brake,
		Steering: steering,
	}
}

func ReadLapDataPacket(buf []byte) PacketLapData {
	lastTime := binary.LittleEndian.Uint32(buf[24:28])

	return PacketLapData{
		LastTime: lastTime,
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var PACKET_CAR_TELEMETRY uint8 = 6
	var PACKET_LAP_DATA uint8 = 2
	env := infrastructure.GetEnv()

	client := influxdb2.NewClientWithOptions(
		env.DbHost,
		env.DbToken,
		influxdb2.DefaultOptions().SetBatchSize(20))
	writeAPI := client.WriteAPI(env.DbOrganization, env.DbBucket)

	sAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:20777")
	if err != nil {
		fmt.Errorf("Error: ", err)
	}

	sConn, err := net.ListenUDP("udp", sAddr)
	if err != nil {
		fmt.Errorf("Error: ", err)
	}
	defer sConn.Close()

	buf := make([]byte, 1289)

	for {
		_, _, err := sConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Errorf("Error: ", err)
		}

		header := new(PacketHeader)
		if err = ReadPacket(buf, header); err != nil {
			fmt.Errorf("Error: ", err)
		}

		switch header.PacketID {
		case PACKET_LAP_DATA:
			packet := ReadLapDataPacket(buf)
			p := influxdb2.NewPoint(
				"system",
				map[string]string{
					"game": "f1-2021",
					"type": "lap-data",
				},
				map[string]interface{}{
					"lastTime": packet.LastTime,
				},
				time.Now())

			writeAPI.WritePoint(p)
		case PACKET_CAR_TELEMETRY:
			packet := ReadCarTelemetryPacket(buf)

			p := influxdb2.NewPoint(
				"system",
				map[string]string{
					"game": "f1-2021",
					"type": "car-telemetry",
				},
				map[string]interface{}{
					"speed":    packet.Speed,
					"throttle": packet.Throttle,
					"brake":    packet.Brake,
					"steering": packet.Steering,
				},
				time.Now())

			writeAPI.WritePoint(p)
		}
	}
}
