package rules

// Brought from https://github.com/ent/ent/blob/4a26cab7347abe01b77107560a4117ced3fad9eb/entc/gen/func.go#L150

import (
	"strings"
	"unicode"

	"github.com/go-openapi/inflect"
)

var (
	ent_acronyms = map[string]struct{}{}
	ent_rules    = inflect.NewDefaultRuleset()
)

func isSeparator(r rune) bool {
	return r == '_' || r == '-' || unicode.IsSpace(r)
}

func pascalWords(words []string) string {
	for i, w := range words {
		upper := strings.ToUpper(w)
		if _, ok := ent_acronyms[upper]; ok {
			words[i] = upper
		} else {
			words[i] = ent_rules.Capitalize(w)
		}
	}
	return strings.Join(words, "")
}

// pascal converts the given name into a PascalCase.
//
//	user_info 	=> UserInfo
//	full_name 	=> FullName
//	user_id   	=> UserID
//	full-admin	=> FullAdmin
func EntPascal(s string) string {
	words := strings.FieldsFunc(s, isSeparator)
	return pascalWords(words)
}

func EntSingularize(s string) string {
	return ent_rules.Singularize(s)
}

func EntPluralize(s string) string {
	return ent_rules.Pluralize(s)
}

func init() {
	for _, w := range []string{
		"ACL", "API", "ASCII", "AWS", "CPU", "CSS", "DNS", "EOF", "GB", "GUID",
		"HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "KB", "LHS", "MAC", "MB",
		"QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SQL", "SSH", "SSO", "TCP",
		"TLS", "TTL", "UDP", "UI", "UID", "URI", "URL", "UTF8", "UUID", "VM",
		"XML", "XMPP", "XSRF", "XSS",
	} {
		ent_acronyms[w] = struct{}{}
		ent_rules.AddAcronym(w)
		ent_rules.AddIrregular("Child", "Children")
	}
}
