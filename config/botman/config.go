package botman

import (
	"github.com/crossplane/upjet/pkg/config"
)

const botmanGroup = "botman"

// Configure configures the botman category resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("akamai_botman_bot_detection_action", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "BotDetectionAction"
	})
	p.AddResourceConfigurator("akamai_botman_akamai_bot_category_action", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "AkamaiBotCategoryAction"
	})
	p.AddResourceConfigurator("akamai_botman_bot_category_exception", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "BotCategoryException"
	})
	p.AddResourceConfigurator("akamai_botman_bot_management_settings", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "BotManagementSettings"
	})
	p.AddResourceConfigurator("akamai_botman_custom_bot_category_item_sequence", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "CustomBotCategoryItemSequence"
	})
	p.AddResourceConfigurator("akamai_botman_challenge_action", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "ChallengeAction"
	})
	p.AddResourceConfigurator("akamai_botman_challenge_interception_rules", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "ChallengeInterceptionRules"
	})
	p.AddResourceConfigurator("akamai_botman_challenge_injection_rules", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "ChallengeInjectionRules"
	})
	p.AddResourceConfigurator("akamai_botman_conditional_action", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "ConditionalAction"
	})
	p.AddResourceConfigurator("akamai_botman_content_protection_javascript_injection_rule", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "ContentProtectionJavascriptInjectionRule"
	})
	p.AddResourceConfigurator("akamai_botman_content_protection_rule", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "ContentProtectionRule"
	})
	p.AddResourceConfigurator("akamai_botman_content_protection_rule_sequence", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "ContentProtectionRuleSequence"
	})
	p.AddResourceConfigurator("akamai_botman_client_side_security", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "ClientSideSecurity"
	})
	p.AddResourceConfigurator("akamai_botman_custom_bot_category", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "CustomBotCategory"
	})
	p.AddResourceConfigurator("akamai_botman_custom_bot_category_action", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "CustomBotCategoryAction"
	})
	p.AddResourceConfigurator("akamai_botman_custom_bot_category_sequence", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "CustomBotCategorySequence"
	})
	p.AddResourceConfigurator("akamai_botman_custom_client", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "CustomClient"
	})
	p.AddResourceConfigurator("akamai_botman_custom_client_sequence", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "CustomClientSequence"
	})
	p.AddResourceConfigurator("akamai_botman_custom_defined_bot", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "CustomDefinedBot"
	})
	p.AddResourceConfigurator("akamai_botman_custom_deny_action", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "CustomDenyAction"
	})
	p.AddResourceConfigurator("akamai_botman_javascript_injection", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "JavascriptInjection"
	})
	p.AddResourceConfigurator("akamai_botman_recategorized_akamai_defined_bot", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "RecategorizedAkamaiDefinedBot"
	})
	p.AddResourceConfigurator("akamai_botman_serve_alternate_action", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "ServeAlternateAction"
	})
	p.AddResourceConfigurator("akamai_botman_transactional_endpoint", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "TransactionalEndpoint"
	})
	p.AddResourceConfigurator("akamai_botman_transactional_endpoint_protection", func(r *config.Resource) {
		r.ShortGroup = botmanGroup
		r.Kind = "TransactionalEndpointProtection"
	})
}
