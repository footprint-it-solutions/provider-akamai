// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	aapselectedhostnames "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/aapselectedhostnames"
	activations "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/activations"
	advancedsettingsattackpayloadlogging "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/advancedsettingsattackpayloadlogging"
	advancedsettingsevasivepathmatch "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/advancedsettingsevasivepathmatch"
	advancedsettingslogging "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/advancedsettingslogging"
	advancedsettingspiilearning "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/advancedsettingspiilearning"
	advancedsettingspragmaheader "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/advancedsettingspragmaheader"
	advancedsettingsprefetch "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/advancedsettingsprefetch"
	advancedsettingsrequestbody "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/advancedsettingsrequestbody"
	apiconstraintsprotection "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/apiconstraintsprotection"
	apirequestconstraints "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/apirequestconstraints"
	attackgroup "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/attackgroup"
	bypassnetworklists "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/bypassnetworklists"
	configuration "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/configuration"
	configurationrename "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/configurationrename"
	customdeny "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/customdeny"
	customrule "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/customrule"
	customruleaction "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/customruleaction"
	eval "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/eval"
	evalgroup "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/evalgroup"
	evalpenaltybox "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/evalpenaltybox"
	evalpenaltyboxconditions "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/evalpenaltyboxconditions"
	evalrule "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/evalrule"
	ipgeo "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/ipgeo"
	ipgeoprotection "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/ipgeoprotection"
	malwarepolicy "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/malwarepolicy"
	malwarepolicyaction "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/malwarepolicyaction"
	malwarepolicyactions "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/malwarepolicyactions"
	malwareprotection "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/malwareprotection"
	matchtarget "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/matchtarget"
	matchtargetsequence "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/matchtargetsequence"
	penaltybox "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/penaltybox"
	penaltyboxconditions "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/penaltyboxconditions"
	rapidrules "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/rapidrules"
	ratepolicy "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/ratepolicy"
	ratepolicyaction "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/ratepolicyaction"
	rateprotection "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/rateprotection"
	reputationprofile "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/reputationprofile"
	reputationprofileaction "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/reputationprofileaction"
	reputationprofileanalysis "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/reputationprofileanalysis"
	reputationprotection "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/reputationprotection"
	rule "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/rule"
	ruleupgrade "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/ruleupgrade"
	securitypolicy "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/securitypolicy"
	securitypolicydefaultprotections "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/securitypolicydefaultprotections"
	securitypolicyrename "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/securitypolicyrename"
	siemsettings "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/siemsettings"
	slowpost "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/slowpost"
	slowpostprotection "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/slowpostprotection"
	threatintel "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/threatintel"
	versionnotes "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/versionnotes"
	wafmode "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/wafmode"
	wafprotection "github.com/footprint-it-solutions/provider-akamai/internal/controller/appsec/wafprotection"
	akamaibotcategoryaction "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/akamaibotcategoryaction"
	botanalyticscookie "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/botanalyticscookie"
	botcategoryexception "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/botcategoryexception"
	botdetectionaction "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/botdetectionaction"
	botmanagementsettings "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/botmanagementsettings"
	challengeaction "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/challengeaction"
	challengeinjectionrules "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/challengeinjectionrules"
	clientsidesecurity "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/clientsidesecurity"
	conditionalaction "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/conditionalaction"
	contentprotectionjavascriptinjectionrule "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/contentprotectionjavascriptinjectionrule"
	contentprotectionrule "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/contentprotectionrule"
	contentprotectionrulesequence "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/contentprotectionrulesequence"
	custombotcategory "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/custombotcategory"
	custombotcategoryaction "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/custombotcategoryaction"
	custombotcategoryitemsequence "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/custombotcategoryitemsequence"
	custombotcategorysequence "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/custombotcategorysequence"
	customclient "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/customclient"
	customclientsequence "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/customclientsequence"
	customcode "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/customcode"
	customdefinedbot "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/customdefinedbot"
	customdenyaction "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/customdenyaction"
	javascriptinjection "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/javascriptinjection"
	recategorizedakamaidefinedbot "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/recategorizedakamaidefinedbot"
	servealternateaction "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/servealternateaction"
	transactionalendpoint "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/transactionalendpoint"
	transactionalendpointprotection "github.com/footprint-it-solutions/provider-akamai/internal/controller/botman/transactionalendpointprotection"
	activation "github.com/footprint-it-solutions/provider-akamai/internal/controller/clientlist/activation"
	list "github.com/footprint-it-solutions/provider-akamai/internal/controller/clientlist/list"
	applicationloadbalancer "github.com/footprint-it-solutions/provider-akamai/internal/controller/cloudlets/applicationloadbalancer"
	applicationloadbalanceractivation "github.com/footprint-it-solutions/provider-akamai/internal/controller/cloudlets/applicationloadbalanceractivation"
	policy "github.com/footprint-it-solutions/provider-akamai/internal/controller/cloudlets/policy"
	policyactivation "github.com/footprint-it-solutions/provider-akamai/internal/controller/cloudlets/policyactivation"
	activationcloudwrapper "github.com/footprint-it-solutions/provider-akamai/internal/controller/cloudwrapper/activation"
	configurationcloudwrapper "github.com/footprint-it-solutions/provider-akamai/internal/controller/cloudwrapper/configuration"
	code "github.com/footprint-it-solutions/provider-akamai/internal/controller/cp/code"
	dvenrollment "github.com/footprint-it-solutions/provider-akamai/internal/controller/cps/dvenrollment"
	dvvalidation "github.com/footprint-it-solutions/provider-akamai/internal/controller/cps/dvvalidation"
	thirdpartyenrollment "github.com/footprint-it-solutions/provider-akamai/internal/controller/cps/thirdpartyenrollment"
	uploadcertificate "github.com/footprint-it-solutions/provider-akamai/internal/controller/cps/uploadcertificate"
	datastream "github.com/footprint-it-solutions/provider-akamai/internal/controller/datastream/datastream"
	record "github.com/footprint-it-solutions/provider-akamai/internal/controller/dns/record"
	zone "github.com/footprint-it-solutions/provider-akamai/internal/controller/dns/zone"
	hostname "github.com/footprint-it-solutions/provider-akamai/internal/controller/edge/hostname"
	edgekv "github.com/footprint-it-solutions/provider-akamai/internal/controller/edgekv/edgekv"
	groupitems "github.com/footprint-it-solutions/provider-akamai/internal/controller/edgekv/groupitems"
	activationedgeworker "github.com/footprint-it-solutions/provider-akamai/internal/controller/edgeworker/activation"
	edgeworker "github.com/footprint-it-solutions/provider-akamai/internal/controller/edgeworker/edgeworker"
	asmap "github.com/footprint-it-solutions/provider-akamai/internal/controller/gtm/asmap"
	cidrmap "github.com/footprint-it-solutions/provider-akamai/internal/controller/gtm/cidrmap"
	datacenter "github.com/footprint-it-solutions/provider-akamai/internal/controller/gtm/datacenter"
	domain "github.com/footprint-it-solutions/provider-akamai/internal/controller/gtm/domain"
	geomap "github.com/footprint-it-solutions/provider-akamai/internal/controller/gtm/geomap"
	property "github.com/footprint-it-solutions/provider-akamai/internal/controller/gtm/property"
	resource "github.com/footprint-it-solutions/provider-akamai/internal/controller/gtm/resource"
	blockeduserproperties "github.com/footprint-it-solutions/provider-akamai/internal/controller/iam/blockeduserproperties"
	group "github.com/footprint-it-solutions/provider-akamai/internal/controller/iam/group"
	ipallowlist "github.com/footprint-it-solutions/provider-akamai/internal/controller/iam/ipallowlist"
	role "github.com/footprint-it-solutions/provider-akamai/internal/controller/iam/role"
	user "github.com/footprint-it-solutions/provider-akamai/internal/controller/iam/user"
	policyimage "github.com/footprint-it-solutions/provider-akamai/internal/controller/imaging/policyimage"
	policyset "github.com/footprint-it-solutions/provider-akamai/internal/controller/imaging/policyset"
	policyvideo "github.com/footprint-it-solutions/provider-akamai/internal/controller/imaging/policyvideo"
	activationsnetworklist "github.com/footprint-it-solutions/provider-akamai/internal/controller/networklist/activations"
	description "github.com/footprint-it-solutions/provider-akamai/internal/controller/networklist/description"
	networklist "github.com/footprint-it-solutions/provider-akamai/internal/controller/networklist/networklist"
	subscription "github.com/footprint-it-solutions/provider-akamai/internal/controller/networklist/subscription"
	activationproperty "github.com/footprint-it-solutions/provider-akamai/internal/controller/property/activation"
	bootstrap "github.com/footprint-it-solutions/provider-akamai/internal/controller/property/bootstrap"
	include "github.com/footprint-it-solutions/provider-akamai/internal/controller/property/include"
	includeactivation "github.com/footprint-it-solutions/provider-akamai/internal/controller/property/includeactivation"
	propertyproperty "github.com/footprint-it-solutions/provider-akamai/internal/controller/property/property"
	providerconfig "github.com/footprint-it-solutions/provider-akamai/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aapselectedhostnames.Setup,
		activations.Setup,
		advancedsettingsattackpayloadlogging.Setup,
		advancedsettingsevasivepathmatch.Setup,
		advancedsettingslogging.Setup,
		advancedsettingspiilearning.Setup,
		advancedsettingspragmaheader.Setup,
		advancedsettingsprefetch.Setup,
		advancedsettingsrequestbody.Setup,
		apiconstraintsprotection.Setup,
		apirequestconstraints.Setup,
		attackgroup.Setup,
		bypassnetworklists.Setup,
		configuration.Setup,
		configurationrename.Setup,
		customdeny.Setup,
		customrule.Setup,
		customruleaction.Setup,
		eval.Setup,
		evalgroup.Setup,
		evalpenaltybox.Setup,
		evalpenaltyboxconditions.Setup,
		evalrule.Setup,
		ipgeo.Setup,
		ipgeoprotection.Setup,
		malwarepolicy.Setup,
		malwarepolicyaction.Setup,
		malwarepolicyactions.Setup,
		malwareprotection.Setup,
		matchtarget.Setup,
		matchtargetsequence.Setup,
		penaltybox.Setup,
		penaltyboxconditions.Setup,
		rapidrules.Setup,
		ratepolicy.Setup,
		ratepolicyaction.Setup,
		rateprotection.Setup,
		reputationprofile.Setup,
		reputationprofileaction.Setup,
		reputationprofileanalysis.Setup,
		reputationprotection.Setup,
		rule.Setup,
		ruleupgrade.Setup,
		securitypolicy.Setup,
		securitypolicydefaultprotections.Setup,
		securitypolicyrename.Setup,
		siemsettings.Setup,
		slowpost.Setup,
		slowpostprotection.Setup,
		threatintel.Setup,
		versionnotes.Setup,
		wafmode.Setup,
		wafprotection.Setup,
		akamaibotcategoryaction.Setup,
		botanalyticscookie.Setup,
		botcategoryexception.Setup,
		botdetectionaction.Setup,
		botmanagementsettings.Setup,
		challengeaction.Setup,
		challengeinjectionrules.Setup,
		clientsidesecurity.Setup,
		conditionalaction.Setup,
		contentprotectionjavascriptinjectionrule.Setup,
		contentprotectionrule.Setup,
		contentprotectionrulesequence.Setup,
		custombotcategory.Setup,
		custombotcategoryaction.Setup,
		custombotcategoryitemsequence.Setup,
		custombotcategorysequence.Setup,
		customclient.Setup,
		customclientsequence.Setup,
		customcode.Setup,
		customdefinedbot.Setup,
		customdenyaction.Setup,
		javascriptinjection.Setup,
		recategorizedakamaidefinedbot.Setup,
		servealternateaction.Setup,
		transactionalendpoint.Setup,
		transactionalendpointprotection.Setup,
		activation.Setup,
		list.Setup,
		applicationloadbalancer.Setup,
		applicationloadbalanceractivation.Setup,
		policy.Setup,
		policyactivation.Setup,
		activationcloudwrapper.Setup,
		configurationcloudwrapper.Setup,
		code.Setup,
		dvenrollment.Setup,
		dvvalidation.Setup,
		thirdpartyenrollment.Setup,
		uploadcertificate.Setup,
		datastream.Setup,
		record.Setup,
		zone.Setup,
		hostname.Setup,
		edgekv.Setup,
		groupitems.Setup,
		activationedgeworker.Setup,
		edgeworker.Setup,
		asmap.Setup,
		cidrmap.Setup,
		datacenter.Setup,
		domain.Setup,
		geomap.Setup,
		property.Setup,
		resource.Setup,
		blockeduserproperties.Setup,
		group.Setup,
		ipallowlist.Setup,
		role.Setup,
		user.Setup,
		policyimage.Setup,
		policyset.Setup,
		policyvideo.Setup,
		activationsnetworklist.Setup,
		description.Setup,
		networklist.Setup,
		subscription.Setup,
		activationproperty.Setup,
		bootstrap.Setup,
		include.Setup,
		includeactivation.Setup,
		propertyproperty.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
