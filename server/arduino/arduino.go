package arduino

import (
	"github.com/tarm/serial"
)

type Arduino struct {
	Port  *serial.Port
	Bytes int
}

func (a *Arduino) Connect(arduinoSerial string, arduinoBaud int) error {
	var err error

	c := &serial.Config{
		Name: arduinoSerial,
		Baud: arduinoBaud,
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
