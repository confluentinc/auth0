package management

import (
	flowv1 "github.com/confluentinc/cc-structs/kafka/flow/v1"
	"errors"
)

type ConnectionManagerInterface interface {
	Create(c *flowv1.Connection) error
	Read(id string, opts ...ReqOption) (*flowv1.Connection, error)
	List(opts ...ReqOption) ([]*flowv1.Connection, error)
	Update(id string, c *flowv1.Connection) (err error)
	Delete(id string) (err error)
	GetConnectionID(connectionName string) (string, error)
}

type ConnectionManager struct {
	m *Management
}

func NewConnectionManager(m *Management) ConnectionManagerInterface {
	return &ConnectionManager{m}
}

func (cm *ConnectionManager) Create(c *flowv1.Connection) error {
	return cm.m.post(cm.m.uri("connections"), c)
}

func (cm *ConnectionManager) Read(id string, opts ...ReqOption) (*flowv1.Connection, error) {
	c := new(flowv1.Connection)
	err := cm.m.get(cm.m.uri("connections", id)+cm.m.q(opts), c)
	return c, err
}

func (cm *ConnectionManager) List(opts ...ReqOption) ([]*flowv1.Connection, error) {
	var c []*flowv1.Connection
	err := cm.m.get(cm.m.uri("connections")+cm.m.q(opts), &c)
	return c, err
}

func (cm *ConnectionManager) Update(id string, c *flowv1.Connection) (err error) {
	return cm.m.patch(cm.m.uri("connections", id), c)
}

func (cm *ConnectionManager) Delete(id string) (err error) {
	return cm.m.delete(cm.m.uri("connections", id))
}

func (cm *ConnectionManager) GetConnectionID(connectionName string) (string, error) {
	connections, err := cm.m.Connection.List(Parameter("name", connectionName), Parameter("fields", "id"))
	if len(connections) == 1 {
		return connections[0].Id, nil
	} else if err == nil {
		err = errors.New(connectionName + " connection does not exist.")
	}
	return "", err
}
