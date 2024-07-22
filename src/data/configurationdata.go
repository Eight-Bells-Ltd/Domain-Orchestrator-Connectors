package data

type AppConfiguration struct {
	Orchestrator string `mapstructure:"orchestrator"`
	NorthBoundPort int `mapstructure:"northBoundPort"`
}
