package gtm

import (
	"github.com/crossplane/upjet/pkg/config"
)

const gtmGroup = "gtm"

// Configure configures the GTM category resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("akamai_gtm_asmap", func(r *config.Resource) {
		r.ShortGroup = gtmGroup
		r.Kind = "ASMap"
	})
	p.AddResourceConfigurator("akamai_gtm_cidrmap", func(r *config.Resource) {
		r.ShortGroup = gtmGroup
		r.Kind = "CIDRMap"
	})
	p.AddResourceConfigurator("akamai_gtm_datacenter", func(r *config.Resource) {
		r.ShortGroup = gtmGroup
		r.Kind = "Datacenter"
	})
	p.AddResourceConfigurator("akamai_gtm_domain", func(r *config.Resource) {
		r.ShortGroup = gtmGroup
		r.Kind = "Domain"
	})
	p.AddResourceConfigurator("akamai_gtm_geomap", func(r *config.Resource) {
		r.ShortGroup = gtmGroup
		r.Kind = "GeoMap"
	})
	p.AddResourceConfigurator("akamai_gtm_property", func(r *config.Resource) {
		r.ShortGroup = gtmGroup
		r.Kind = "Property"
	})
	p.AddResourceConfigurator("akamai_gtm_resource", func(r *config.Resource) {
		r.ShortGroup = gtmGroup
		r.Kind = "Resource"
	})
}
