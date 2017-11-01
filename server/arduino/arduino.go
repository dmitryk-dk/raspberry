package arduino

import (
	"github.com/tarm/serial"
)

const (
	ArduinoSerialName = "/dev/ttyACM0"
	ArduinoBaud       = 9600
)

type Arduino struct {
	Port  *serial.Port
	Bytes int
}

func (a *Arduino) Connect() error {
	var err error

	c := &serial.Config{
		Name: ArduinoSerialName,
		Baud: ArduinoBaud,
	}

	a.Port, err = serial.OpenPort(c)

	if err != nil {
		return err
	}

	return nil
}

func (a *Arduino) SendCommand() (int, error) {
	var err error

	a.Bytes, err = a.Port.Write([]byte("s"))
	if err != nil {
		return 0, err
	}

	return a.Bytes, nil
}

func (a *Arduino) GetData() (int, []byte, error) {
	var n int
	var err error

	buf := make([]byte, 16)
	n, err = a.Port.Read(buf)
	if err != nil {
		return 0, nil, err
	}

	return n, buf, nil
}