package pkg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

const (
	PACKET_CAR_TELEMETRY uint8 = 6
	PACKET_LAP_DATA      uint8 = 2
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

func ReadPacket(buf []byte) (TelemetryPacket, error) {
	header := PacketHeader{}
	if err := readPacketHeader(buf, header); err != nil {
		return TelemetryPacket{}, err
	}

	var packet interface{}
	packetType := header.PacketID

	switch header.PacketID {
	case PACKET_LAP_DATA:
		packet = ReadLapDataPacket(buf)
	case PACKET_CAR_TELEMETRY:
		packet = ReadCarTelemetryPacket(buf)
	default:
		return TelemetryPacket{}, fmt.Errorf("unknown packet %d", header.PacketID)
	}

	telemetry := TelemetryPacket{
		Type: packetType,
		Data: packet,
	}
	return telemetry, nil
}

func readPacketHeader(buf []byte, pack interface{}) error {
	reader := bytes.NewReader(buf)
	if err := binary.Read(reader, binary.LittleEndian, pack); err != nil {
		return err
	}

	return nil
}

func ReadCarTelemetryPacket(buf []byte) TelemetryCarData {
	speed := binary.LittleEndian.Uint16(buf[24:26])
	throttle := math.Float32frombits(binary.LittleEndian.Uint32(buf[26:30]))
	steering := math.Float32frombits(binary.LittleEndian.Uint32(buf[30:34]))
	brake := math.Float32frombits(binary.LittleEndian.Uint32(buf[34:38]))

	return TelemetryCarData{
		Speed:    speed,
		Throttle: throttle,
		Brake:    brake,
		Steering: steering,
	}
}

func ReadLapDataPacket(buf []byte) TelemetryLapData {
	lastTime := binary.LittleEndian.Uint32(buf[24:28])

	return TelemetryLapData{
		LastTime: lastTime,
	}
}
