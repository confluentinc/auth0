package management

import (
	"encoding/json"
	flowv1 "github.com/confluentinc/cc-structs/kafka/flow/v1"
	"errors"
)

type ConnectionManagerInterface interface {
	Create(c *Connection) error
	Read(id string, opts ...ReqOption) (*Connection, error)
	List(opts ...ReqOption) ([]*Connection, error)
	Update(id string, c *Connection) (err error)
	Delete(id string) (err error)
	GetConnectionID(connectionName string) (*string, error)
}

type Connection struct {
	// A generated string identifying the connection.
	ID *string `json:"id,omitempty"`

	// The name of the connection. Must start and end with an alphanumeric
	// character and can only contain alphanumeric characters and '-'. Max
	// length 128.
	Name *string `json:"name,omitempty"`

	// The identity provider identifier for the connection. Can be any of the
	// following:
	//
	// "ad", "adfs", "amazon", "dropbox", "bitbucket", "aol", "auth0-adldap",
	// "auth0-oidc", "auth0", "baidu", "bitly", "box", "custom", "daccount",
	// "dwolla", "email", "evernote-sandbox", "evernote", "exact", "facebook",
	// "fitbit", "flickr", "github", "google-apps", "google-oauth2", "guardian",
	//  "instagram", "ip", "linkedin", "miicard", "oauth1", "oauth2",
	// "office365", "paypal", "paypal-sandbox", "pingfederate",
	// "planningcenter", "renren", "salesforce-community", "salesforce-sandbox",
	//  "salesforce", "samlp", "sharepoint", "shopify", "sms", "soundcloud",
	// "thecity-sandbox", "thecity", "thirtysevensignals", "twitter", "untappd",
	//  "vkontakte", "waad", "weibo", "windowslive", "wordpress", "yahoo",
	// "yammer" or "yandex".
	Strategy *string `json:"strategy,omitempty"`

	// True if the connection is domain level
	IsDomainConnection *bool `json:"is_domain_connection,omitempty"`

	// Options for validation.
	// Options map[string]interface{} `json:"options,omitempty"`
	Options *flowv1.ConnectionOptions `json:"options,omitempty"`

	// The identifiers of the clients for which the connection is to be
	// enabled. If the array is empty or the property is not specified, no
	// clients are enabled.
	EnabledClients []string `json:"enabled_clients,omitempty"`

	// Defines the realms for which the connection will be used (ie: email
	// domains). If the array is empty or the property is not specified, the
	// connection name will be added as realm.
	Realms []interface{} `json:"realms,omitempty"`

	Metadata *interface{} `json:"metadata,omitempty"`
}

func (c *Connection) String() string {
	b, _ := json.MarshalIndent(c, "", "  ")
	return string(b)
}

type ConnectionOptions struct {
	SigningCert string `json:"signingCert,omitempty"`
	SignInEndpoint string `json:"signInEndpoint,omitempty"`
	SignOutEndpoint string  `json:"signOutEndpoint,omitempty"`
	TenantDomain string  `json:"tenant_domain,omitempty"`
	DomainAliases []string `json:"domain_aliases,omitempty"`
	FieldsMap interface{} `json:"fieldsMap,omitempty"`
}

type ConnectionManager struct {
	m *Management
}

func NewConnectionManager(m *Management) ConnectionManagerInterface {
	return &ConnectionManager{m}
}

func (cm *ConnectionManager) Create(c *Connection) error {
	return cm.m.post(cm.m.uri("connections"), c)
}

func (cm *ConnectionManager) Read(id string, opts ...ReqOption) (*Connection, error) {
	c := new(Connection)
	err := cm.m.get(cm.m.uri("connections", id)+cm.m.q(opts), c)
	return c, err
}

func (cm *ConnectionManager) List(opts ...ReqOption) ([]*Connection, error) {
	var c []*Connection
	err := cm.m.get(cm.m.uri("connections")+cm.m.q(opts), &c)
	return c, err
}

func (cm *ConnectionManager) Update(id string, c *Connection) (err error) {
	return cm.m.patch(cm.m.uri("connections", id), c)
}

func (cm *ConnectionManager) Delete(id string) (err error) {
	return cm.m.delete(cm.m.uri("connections", id))
}

func (cm *ConnectionManager) GetConnectionID(connectionName string) (*string, error) {
	connections, err := cm.m.Connection.List(Parameter("name", connectionName), Parameter("fields", "id"))
	if len(connections) == 1 {
		return connections[0].ID, nil
	} else if err == nil {
		err = errors.New(connectionName + " connection does not exist.")
	}
	return nil, err
}
