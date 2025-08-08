package property

import (
	"github.com/crossplane/upjet/pkg/config"
)

const propertyGroup = "property"

// Configure configures the property category resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("akamai_property", func(r *config.Resource) {
		r.ShortGroup = propertyGroup
		r.Kind = "Property"
	})
	p.AddResourceConfigurator("akamai_property_activation", func(r *config.Resource) {
		r.ShortGroup = propertyGroup
		r.Kind = "Activation"
	})
	p.AddResourceConfigurator("akamai_property_bootstrap", func(r *config.Resource) {
		r.ShortGroup = propertyGroup
		r.Kind = "Bootstrap"
	})
	p.AddResourceConfigurator("akamai_property_include", func(r *config.Resource) {
		r.ShortGroup = propertyGroup
		r.Kind = "Include"
	})
	p.AddResourceConfigurator("akamai_property_include_activation", func(r *config.Resource) {
		r.ShortGroup = propertyGroup
		r.Kind = "IncludeActivation"
	})
}
