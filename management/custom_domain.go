package management

import "encoding/json"

type CustomDomain struct {

	// The id of the custom domain
	ID *string `json:"custom_domain_id,omitempty"`

	// The custom domain.
	Domain *string `json:"domain,omitempty"`

	// The custom domain provisioning type. Can be either "auth0_managed_certs"
	// or "self_managed_certs"
	Type *string `json:"type,omitempty"`

	// Primary is true if the domain was marked as "primary", false otherwise.
	Primary *bool `json:"primary,omitempty"`

	// The custom domain configuration status. Can be any of the following:
	//
	// "disabled", "pending", "pending_verification" or "ready"
	Status *string `json:"status,omitempty"`

	// The custom domain verification method. The only allowed value is "txt".
	VerificationMethod *string `json:"verification_method,omitempty"`

	Verification *CustomDomainVerification `json:"verification,omitempty"`
}

func (c *CustomDomain) String() string {
	b, _ := json.MarshalIndent(c, "", "  ")
	return string(b)
}

type CustomDomainVerification struct {

	// The custom domain verification methods.
	Methods []map[string]interface{} `json:"methods,omitempty"`
}

type CustomDomainManager struct {
	m *Management
}

func NewCustomDomainManager(m *Management) *CustomDomainManager {
	return &CustomDomainManager{m}
}

func (cm *CustomDomainManager) Create(c *CustomDomain) (err error) {
	return cm.m.post(cm.m.uri("custom-domains"), c)
}

func (cm *CustomDomainManager) Read(id string, opts ...ReqOption) (*CustomDomain, error) {
	c := new(CustomDomain)
	err := cm.m.get(cm.m.uri("custom-domains", id)+cm.m.q(opts), c)
	return c, err
}

func (cm *CustomDomainManager) Update(id string, c *CustomDomain) (err error) {
	return cm.m.patch(cm.m.uri("custom-domains", id), c)
}

func (cm *CustomDomainManager) Delete(id string) (err error) {
	return cm.m.delete(cm.m.uri("custom-domains", id))
}
