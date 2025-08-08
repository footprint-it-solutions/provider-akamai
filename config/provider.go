/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/footprint-it-solutions/provider-akamai/config/appsecurity"
	"github.com/footprint-it-solutions/provider-akamai/config/botman"
	"github.com/footprint-it-solutions/provider-akamai/config/clientlist"
	"github.com/footprint-it-solutions/provider-akamai/config/cloudlets"
	"github.com/footprint-it-solutions/provider-akamai/config/cloudwrapper"
	"github.com/footprint-it-solutions/provider-akamai/config/cps"
	"github.com/footprint-it-solutions/provider-akamai/config/datastream"
	"github.com/footprint-it-solutions/provider-akamai/config/dns"
	"github.com/footprint-it-solutions/provider-akamai/config/edge"
	"github.com/footprint-it-solutions/provider-akamai/config/gtm"
	"github.com/footprint-it-solutions/provider-akamai/config/iam"
	"github.com/footprint-it-solutions/provider-akamai/config/imagingpolicy"
	"github.com/footprint-it-solutions/provider-akamai/config/networklist"
	"github.com/footprint-it-solutions/provider-akamai/config/property"
)

const (
	resourcePrefix = "akamai"
	modulePath     = "github.com/footprint-it-solutions/provider-akamai"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("akamai.footprintit.net"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		appsecurity.Configure,
		botman.Configure,
		clientlist.Configure,
		cloudlets.Configure,
		cloudwrapper.Configure,
		cps.Configure,
		datastream.Configure,
		dns.Configure,
		edge.Configure,
		gtm.Configure,
		iam.Configure,
		imagingpolicy.Configure,
		networklist.Configure,
		property.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
