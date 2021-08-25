package desgin_pattern

const (
	defaultMaxTotal = 10
	defaultMaxIdle  = 9
	defaultMinIdle  = 1
)

type ResourceConfig struct {
	name     string
	maxTotal string
}
