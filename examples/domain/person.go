package domain

import "github.com/audio35444/kibanalogger"

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname" kibana:"last_name"`
	Age     int32  `json:"age" kibana:"age"`
}

func (p Person) String() string {
	return kibanalogger.KibanaLogs(&p)
}
