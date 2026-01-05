package testcase

type IOType string

const (
	CACHEDATA IOType = "cache_data"
	EVENT     IOType = "event"
	REST      IOType = "rest"
	ROWSET    IOType = "rowset"
)

type IO struct {
	Target string `yaml:"target"`
	Type   IOType `yaml:"type"`
}

type TestCase struct {
	Trigger IO   `yaml:"trigger"`
	Inputs  []IO `yaml:"inputs"`
	Outputs []IO `yaml:"outputs"`
}
