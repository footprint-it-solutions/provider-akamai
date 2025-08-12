/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"os"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/footprint-it-solutions/provider-akamai/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal akamai credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials as environment variables for Terraform provider
		if err := os.Setenv("AKAMAI_CLIENT_SECRET", creds["default"]["client_secret"]); err != nil {
			return ps, errors.Wrap(err, "cannot set AKAMAI_CLIENT_SECRET environment variable")
		}
		if err := os.Setenv("AKAMAI_HOST", creds["default"]["host"]); err != nil {
			return ps, errors.Wrap(err, "cannot set AKAMAI_HOST environment variable")
		}
		if err := os.Setenv("AKAMAI_ACCESS_TOKEN", creds["default"]["access_token"]); err != nil {
			return ps, errors.Wrap(err, "cannot set AKAMAI_ACCESS_TOKEN environment variable")
		}
		if err := os.Setenv("AKAMAI_CLIENT_TOKEN", creds["default"]["client_token"]); err != nil {
			return ps, errors.Wrap(err, "cannot set AKAMAI_CLIENT_TOKEN environment variable")
		}

		return ps, nil
	}
}
