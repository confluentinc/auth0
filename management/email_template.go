package management

import "encoding/json"

type EmailTemplate struct {

	// The template name. Can be one of "verify_email", "reset_email",
	// "welcome_email", "blocked_account", "stolen_credentials",
	// "enrollment_email", "change_password", "password_reset" or
	// "mfa_oob_code".
	Template *string `json:"template"`

	// The body of the template.
	Body *string `json:"body,omitempty"`

	// The sender of the email.
	From *string `json:"from,omitempty"`

	// The URL to redirect the user to after a successful action.
	ResultURL *string `json:"resultUrl,omitempty"`

	// The subject of the email.
	Subject *string `json:"subject,omitempty"`

	// The syntax of the template body.
	Syntax *string `json:"syntax,omitempty"`

	// The lifetime in seconds that the link within the email will be valid for.
	URLLifetimeInSecoonds *int `json:"urlLifetimeInSeconds,omitempty"`

	// Whether or not the template is enabled.
	Enabled *bool `json:"enabled"`
}

func (e *EmailTemplate) String() string {
	b, _ := json.MarshalIndent(e, "", "  ")
	return string(b)
}

type EmailTemplateManager struct {
	m *Management
}

func NewEmailTemplateManager(m *Management) *EmailTemplateManager {
	return &EmailTemplateManager{m}
}

func (em *EmailTemplateManager) Create(e *EmailTemplate) error {
	return em.m.post(em.m.uri("email-templates"), e)
}

func (em *EmailTemplateManager) Read(template string, opts ...ReqOption) (*EmailTemplate, error) {
	e := new(EmailTemplate)
	err := em.m.get(em.m.uri("email-templates", template)+em.m.q(opts), e)
	return e, err
}

func (em *EmailTemplateManager) Update(template string, e *EmailTemplate) (err error) {
	return em.m.patch(em.m.uri("email-templates", template), e)
}

func (em *EmailTemplateManager) Replace(template string, e *EmailTemplate) (err error) {
	return em.m.put(em.m.uri("email-templates", template), e)
}
