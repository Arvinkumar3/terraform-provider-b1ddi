package b1ddi

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceIpamsvcFixedAddress(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceFixedAddressBasicTestStep(),
			{
				Config: fmt.Sprintf(`
					data "b1ddi_fixed_addresses" "tf_acc_fixed_addresses" {
						filters = {
							"name" = "tf_acc_test_fixed_address"
						}
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_fixed_addresses.tf_acc_fixed_addresses", "results.#", "1"),
					resource.TestCheckResourceAttr("data.b1ddi_fixed_addresses.tf_acc_fixed_addresses", "results.0.comment", "This Fixed Address is created by terraform provider acceptance test"),
				),
			},
		},
	})
}