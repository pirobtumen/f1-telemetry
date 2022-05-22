package pkg

type TelemetryPacket struct {
	Type uint8
	Data interface{}
}

type TelemetryCarData struct {
	Speed    uint16
	Throttle float32
	Brake    float32
	Steering float32
}

type TelemetryLapData struct {
	LastTime uint32
}
