package iam

import (
	"github.com/crossplane/upjet/pkg/config"
)

const iamGroup = "iam"

// Configure configures the IAM category resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("akamai_iam_blocked_user_properties", func(r *config.Resource) {
		r.ShortGroup = iamGroup
		r.Kind = "BlockedUserProperties"
	})
	p.AddResourceConfigurator("akamai_iam_group", func(r *config.Resource) {
		r.ShortGroup = iamGroup
		r.Kind = "Group"
	})
	p.AddResourceConfigurator("akamai_iam_ip_allowlist", func(r *config.Resource) {
		r.ShortGroup = iamGroup
		r.Kind = "IPAllowlist"
	})
	p.AddResourceConfigurator("akamai_iam_role", func(r *config.Resource) {
		r.ShortGroup = iamGroup
		r.Kind = "Role"
	})
	p.AddResourceConfigurator("akamai_iam_user", func(r *config.Resource) {
		r.ShortGroup = iamGroup
		r.Kind = "User"
	})
}
