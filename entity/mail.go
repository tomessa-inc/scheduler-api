package entity

type SESSendTemplatedEmailArgs struct {
	TOs          []string
	CCs          []string
	TemplateName string
	TemplateData map[string]string
}

type SESSendEmailArgs struct {
	TOs      []string
	CCs      []string
	Subject  string
	HtmlBody string
	TextBody string
}
