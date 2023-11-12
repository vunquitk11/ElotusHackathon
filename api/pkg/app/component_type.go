package app

const (
	// ComponentTypeAPI means an API component of the Project
	ComponentTypeAPI = ComponentType("api")
	// ComponentTypeJob means an Job component of the Project
	ComponentTypeJob = ComponentType("job")
	// ComponentTypeConsumer means a Kafka/RMQ/MQTT Consumer component of the Project
	ComponentTypeConsumer = ComponentType("consumer")
)

// ComponentType denotes the component type of the Project
type ComponentType string

// Valid checks if the env is valid or not
func (e ComponentType) Valid() bool {
	switch e {
	case ComponentTypeAPI, ComponentTypeJob, ComponentTypeConsumer:
		return true
	default:
		return false
	}
}

// String returns the string representation of ComponentType
func (e ComponentType) String() string {
	return string(e)
}
