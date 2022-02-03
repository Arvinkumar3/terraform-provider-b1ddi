---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "b1ddi_dns_views Data Source - terraform-provider-b1ddi"
subcategory: ""
description: |-
  
---

# b1ddi_dns_views (Data Source)



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

# Select DNS View by name
data "b1ddi_dns_views" "example_dns_view" {
  filters = {
    "name" = "example_tf_dns_view"
  }
}

# Get all DNS Views
data "b1ddi_dns_views" "all_dns_views" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **filters** (Map of String) Configure a map of filters to be applied on the search result.
- **id** (String) The ID of this resource.

### Read-Only

- **results** (List of Object) List of DNS Views matching filters. The schema of each element is identical to the b1ddi_dns_view resource schema. (see [nested schema](../resources/dns_view.md))