package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/view"
	"testing"
)

func TestAccResourceDnsView_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsViewBasicTestStep(),
		},
	})
}

func resourceDnsViewBasicTestStep() resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
					resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
						name = "tf_acc_test_dns_view"
					}`),
		Check: resource.ComposeAggregateTestCheckFunc(
			testAccDnsViewExists("b1ddi_dns_view.tf_acc_test_dns_view"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "comment", ""),
			resource.TestCheckResourceAttrSet("b1ddi_dns_view.tf_acc_test_dns_view", "created_at"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "custom_root_ns.%", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "custom_root_ns_enabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "disabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_enable_validation", "true"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_enabled", "true"),

			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_root_keys.0.algorithm", "8"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_root_keys.0.protocol_zone", "."),
			resource.TestCheckResourceAttrSet("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_root_keys.0.public_key"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_root_keys.0.sep", "true"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_root_keys.0.zone", "."),

			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_trust_anchors.%", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_validate_expiry", "true"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "ecs_enabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "ecs_forwarding", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "ecs_prefix_v4", "24"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "ecs_prefix_v6", "56"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "ecs_zones.%", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "edns_udp_size", "1232"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "forwarders.%", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "forwarders_only", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "gss_tsig_enabled", "false"),
			resource.TestCheckNoResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "inheritance_sources"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "ip_spaces.%", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "lame_ttl", "600"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "match_clients_acl.%", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "match_destinations_acl.%", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "match_recursive_only", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "max_cache_ttl", "604800"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "max_negative_ttl", "10800"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "max_udp_size", "1232"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "minimal_responses", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "name", "tf_acc_test_dns_view"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "notify", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "query_acl.%", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "recursion_acl.%", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "recursion_enabled", "true"),
			resource.TestCheckNoResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "tags"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "transfer_acl.%", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "update_acl.%", "0"),
			resource.TestCheckResourceAttrSet("b1ddi_dns_view.tf_acc_test_dns_view", "updated_at"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "use_forwarders_for_subzones", "true"),

			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.default_ttl", "28800"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.expire", "2419200"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.mname", "ns.b1ddi"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.negative_ttl", "900"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.protocol_mname", "ns.b1ddi"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.protocol_rname", "hostmaster"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.refresh", "10800"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.retry", "3600"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.rname", "hostmaster"),
			resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.use_default_mname", "true"),
		),
	}
}

func TestAccResourceDnsView_full_config(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
						comment = "This DNS View is created by the terraform provider acceptance test"

						#custom_root_ns {}
						#custom_root_ns_enabled = true

						disabled = true
						dnssec_enable_validation = false
						dnssec_enabled = false
						
						#dnssec_trust_anchors {}

						dnssec_validate_expiry = false
						ecs_enabled = false
						ecs_forwarding = true
						ecs_prefix_v4 = 12
						ecs_prefix_v6 = 28
						
						#ecs_zones {}
						
						edns_udp_size = 1024

						#forwarders {}
						#forwarders_only = true

						gss_tsig_enabled = true

						ip_spaces = [b1ddi_ip_space.tf_acc_test_space.id]

						lame_ttl = 1200

						#match_clients_acl {}
						#match_destinations_acl {}

						match_recursive_only = true
						max_cache_ttl = 302400
						max_negative_ttl = 302400
						max_udp_size = 1024
						minimal_responses = true
						name = "tf_acc_test_dns_view_full_config"
						notify = true

						#query_acl {}
						#recursion_acl {}

						recursion_enabled = false
						tags = {
							TestType = "Acceptance"
						}

						#transfer_acl {}
						#update_acl {}

						use_forwarders_for_subzones = false

						zone_authority {
							default_ttl = 14400
							expire = 1209600
							mname = "mname"
							negative_ttl = 700
							refresh = 5400
							retry = 1800
							rname = "rname"
							use_default_mname = false
						}
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDnsViewExists("b1ddi_dns_view.tf_acc_test_dns_view"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "comment", "This DNS View is created by the terraform provider acceptance test"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_view.tf_acc_test_dns_view", "created_at"),
					// ToDo add check for custom root ns
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "custom_root_ns.%", "0"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "custom_root_ns_enabled", "false"),

					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "disabled", "true"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_enable_validation", "false"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_enabled", "false"),

					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_root_keys.0.algorithm", "8"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_root_keys.0.protocol_zone", "."),
					resource.TestCheckResourceAttrSet("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_root_keys.0.public_key"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_root_keys.0.sep", "true"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_root_keys.0.zone", "."),

					// ToDo add check for custom dnssec trust anchors
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_trust_anchors.%", "0"),

					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_validate_expiry", "false"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "ecs_enabled", "false"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "ecs_forwarding", "true"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "ecs_prefix_v4", "12"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "ecs_prefix_v6", "28"),
					// ToDo add check for custom ecs_zones and ecs_enabled=true
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "ecs_zones.%", "0"),

					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "edns_udp_size", "1024"),

					// ToDo add check for custom forwarders list
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "forwarders.%", "0"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "forwarders_only", "false"),

					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "gss_tsig_enabled", "true"),
					resource.TestCheckNoResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "inheritance_sources"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_view.tf_acc_test_dns_view", "ip_spaces.0"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "lame_ttl", "1200"),

					// ToDo add check for custom match_clients_acl
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "match_clients_acl.%", "0"),
					// ToDo add check for custom match_destinations_acl
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "match_destinations_acl.%", "0"),

					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "match_recursive_only", "true"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "max_cache_ttl", "302400"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "max_negative_ttl", "302400"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "max_udp_size", "1024"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "minimal_responses", "true"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "name", "tf_acc_test_dns_view_full_config"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "notify", "true"),

					// ToDo add check for custom ACLs
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "query_acl.%", "0"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "recursion_acl.%", "0"),

					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "recursion_enabled", "false"),
					resource.TestCheckNoResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "tags"),

					// ToDo add check for custom ACLs
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "transfer_acl.%", "0"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "update_acl.%", "0"),

					resource.TestCheckResourceAttrSet("b1ddi_dns_view.tf_acc_test_dns_view", "updated_at"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "use_forwarders_for_subzones", "false"),

					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.default_ttl", "14400"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.expire", "1209600"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.mname", "mname"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.negative_ttl", "700"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.protocol_mname", "mname"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.protocol_rname", "rname"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.refresh", "5400"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.retry", "1800"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.rname", "rname"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.use_default_mname", "false"),
				),
			},
		},
	})
}

func TestAccResourceDnsView_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsViewBasicTestStep(),
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
						comment = "This DNS View is created by the terraform provider acceptance test"

						#custom_root_ns {}
						#custom_root_ns_enabled = true

						disabled = true
						dnssec_enable_validation = false
						dnssec_enabled = false
						
						#dnssec_trust_anchors {}

						dnssec_validate_expiry = false
						ecs_enabled = false
						ecs_forwarding = true
						ecs_prefix_v4 = 12
						ecs_prefix_v6 = 28
						
						#ecs_zones {}
						
						edns_udp_size = 1024

						#forwarders {}
						#forwarders_only = true

						gss_tsig_enabled = true

						ip_spaces = [b1ddi_ip_space.tf_acc_test_space.id]

						lame_ttl = 1200

						#match_clients_acl {}
						#match_destinations_acl {}

						match_recursive_only = true
						max_cache_ttl = 302400
						max_negative_ttl = 302400
						max_udp_size = 1024
						minimal_responses = true
						name = "tf_acc_test_dns_view"
						notify = true

						#query_acl {}
						#recursion_acl {}

						recursion_enabled = false
						tags = {
							TestType = "Acceptance"
						}

						#transfer_acl {}
						#update_acl {}

						use_forwarders_for_subzones = false

						zone_authority {
							default_ttl = 14400
							expire = 1209600
							mname = "mname"
							negative_ttl = 700
							refresh = 5400
							retry = 1800
							rname = "rname"
							use_default_mname = false
						}
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDnsViewExists("b1ddi_dns_view.tf_acc_test_dns_view"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "comment", "This DNS View is created by the terraform provider acceptance test"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_view.tf_acc_test_dns_view", "created_at"),
					// ToDo add check for custom root ns
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "custom_root_ns.%", "0"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "custom_root_ns_enabled", "false"),

					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "disabled", "true"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_enable_validation", "false"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_enabled", "false"),

					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_root_keys.0.algorithm", "8"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_root_keys.0.protocol_zone", "."),
					resource.TestCheckResourceAttrSet("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_root_keys.0.public_key"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_root_keys.0.sep", "true"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_root_keys.0.zone", "."),

					// ToDo add check for custom dnssec trust anchors
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_trust_anchors.%", "0"),

					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "dnssec_validate_expiry", "false"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "ecs_enabled", "false"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "ecs_forwarding", "true"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "ecs_prefix_v4", "12"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "ecs_prefix_v6", "28"),
					// ToDo add check for custom ecs_zones and ecs_enabled=true
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "ecs_zones.%", "0"),

					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "edns_udp_size", "1024"),

					// ToDo add check for custom forwarders list
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "forwarders.%", "0"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "forwarders_only", "false"),

					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "gss_tsig_enabled", "true"),
					resource.TestCheckNoResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "inheritance_sources"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_view.tf_acc_test_dns_view", "ip_spaces.0"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "lame_ttl", "1200"),

					// ToDo add check for custom match_clients_acl
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "match_clients_acl.%", "0"),
					// ToDo add check for custom match_destinations_acl
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "match_destinations_acl.%", "0"),

					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "match_recursive_only", "true"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "max_cache_ttl", "302400"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "max_negative_ttl", "302400"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "max_udp_size", "1024"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "minimal_responses", "true"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "name", "tf_acc_test_dns_view"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "notify", "true"),

					// ToDo add check for custom ACLs
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "query_acl.%", "0"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "recursion_acl.%", "0"),

					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "recursion_enabled", "false"),
					resource.TestCheckNoResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "tags"),

					// ToDo add check for custom ACLs
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "transfer_acl.%", "0"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "update_acl.%", "0"),

					resource.TestCheckResourceAttrSet("b1ddi_dns_view.tf_acc_test_dns_view", "updated_at"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "use_forwarders_for_subzones", "false"),

					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.default_ttl", "14400"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.expire", "1209600"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.mname", "mname"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.negative_ttl", "700"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.protocol_mname", "mname"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.protocol_rname", "rname"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.refresh", "5400"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.retry", "1800"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.rname", "rname"),
					resource.TestCheckResourceAttr("b1ddi_dns_view.tf_acc_test_dns_view", "zone_authority.0.use_default_mname", "false"),
				),
			},
		},
	})
}

func testAccDnsViewExists(resPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.DNSConfigurationAPI.View.ViewRead(
			&view.ViewReadParams{ID: res.Primary.ID, Context: context.TODO()},
			nil,
		)
		if err != nil {
			return err
		}
		if resp.Payload.Result.ID != res.Primary.ID {
			return fmt.Errorf(
				"'id' does not match: \n got: '%s', \nexpected: '%s'",
				resp.Payload.Result.ID,
				res.Primary.ID)
		}
		return nil
	}
}
