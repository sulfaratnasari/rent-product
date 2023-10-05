package config

type ConfigCredential struct {
	Data *DataCredential `json:"data"`
}

type DataCredential struct {
	Database *Database `json:"database"`
}

type Database struct {
	Master string `json:"master"`
	Slave  string `json:"slave"`
	Prefix string `json:"prefix"`
}
