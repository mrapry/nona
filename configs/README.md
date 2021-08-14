#Configs System

We have 2 file in here.
* config
* environment

##Config

Use for config system like using dependency, database, etc

##Environment

Use for load Environment OS additional need for system. Sample you have config for topic kafka. you just set like this

```go
import "os"

type Environment struct {
	Kafka struct{
		TopicSample string // add variable in here for define call struct env in another file
	}
}

var env Environment

// GetEnv get global additional environment
func GetEnv() Environment {
	return env
}

func loadAdditionalEnv() {
	var ok bool

	env.Kafka.TopicSample, ok = os.LookupEnv("app.kafka.topic.sample") // load env additional in here
	if !ok {
		panic("missing app.kafka.topic.sample environment") // if error variable is mandatory request you can set error to panic
	}

}
```