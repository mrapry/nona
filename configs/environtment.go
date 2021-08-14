package configs

import "os"

type Environment struct {
	Kafka struct{
		TopicSample string
	}
}

var env Environment

// GetEnv get global additional environment
func GetEnv() Environment {
	return env
}

func loadAdditionalEnv() {
	var ok bool

	env.Kafka.TopicSample, ok = os.LookupEnv("app.kafka.topic.sample")
	if !ok {
		panic("missing app.kafka.topic.sample environment")
	}

}