package models

type Collections struct {
	MachineId int `json:"machine_id"`
	IPAddress string `json:"ip_address"`
	MachineTag string `json:"machine_tag"`
	WorkStatus string `json:"work_status"`
	CPUUsage string `json:"cpu_usage"`
	MemUsage string `json:"mem_usage"`
	MinerCount string `json:"miner_count"`
	NodeCount string `json:"node_count"`
	NodeGraffi string `json:"node_graffi"`
	NodeHeight string `json:"node_height"`
	NodeVersion string `json:"node_version"`
	UpdateTime string `json:"update_time"`
}
