package queue

type (
	// The Option type is a function that modifies the queue
	// configuration.
	Option func(c *config)

	config struct {
		maxWorkers int
	}
)

// MaxWorkers specifies the maximum amount of concurrent jobs that the queue
// can process.
func MaxWorkers(max int) Option {
	return func(c *config) {
		c.maxWorkers = max
	}
}
