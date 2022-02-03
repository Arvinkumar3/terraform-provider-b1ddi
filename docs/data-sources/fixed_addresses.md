---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "b1ddi_fixed_addresses Data Source - terraform-provider-b1ddi"
subcategory: ""
description: |-
  
---

# b1ddi_fixed_addresses (Data Source)



## Example Usage

```terraform
terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "github.com/infobloxopen/b1ddi"
    }
  }
}

# Select specific fixed address
data "b1ddi_fixed_addresses" "tf_example_fixed_address" {
  filters = {
    "address" = "192.168.1.15"
  }
}

# Get all Fixed Addresses
data "b1ddi_fixed_addresses" "all_fixed_addresses" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **filters** (Map of String) Configure a map of filters to be applied on the search result.
- **id** (String) The ID of this resource.

### Read-Only

- **results** (List of Object) List of Fixed Addresses matching filters. The schema of each element is identical to the b1ddi_fixed_address resource schema. (see [nested schema](../resources/fixed_address.md))