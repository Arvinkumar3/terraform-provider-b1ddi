terraform {
  required_providers {
    b1ddi = {
      source  = "infobloxopen/b1ddi"
    }
  }
}

resource "b1ddi_ip_space" "example_tf_space" {
  name    = "example_tf_space"
  comment = "Example IP space created by the terraform provider"
}

resource "b1ddi_address_block" "tf_example_address_block" {
  address = "192.168.0.0"
  cidr = 16
  name = "tf_example_address_block"
  space = b1ddi_ip_space.example_tf_space.id
}

resource "b1ddi_subnet" "example_tf_subnet" {
  name    = "example_tf_subnet"
  space   = b1ddi_ip_space.example_tf_space.id
  address = "192.168.1.0"
  cidr    = 24
  comment = "Example Subnet created by the terraform provider"
}

resource "b1ddi_range" "tf_acc_test_range" {
  start = "192.168.1.15"
  end = "192.168.1.30"
  name = "example_tf_range"
  space = b1ddi_ip_space.example_tf_space.id
  comment = "Example Range created by the terraform provider"
  depends_on = [b1ddi_subnet.example_tf_subnet]
}

# Address can be specified manually via the address field.
# This is a static assignment
resource "b1ddi_address" "example_tf_address" {
  address    = "192.168.1.45"
  comment    = "Example Address created by the terraform provider"
  space      = b1ddi_ip_space.example_tf_space.id
  depends_on = [b1ddi_subnet.example_tf_subnet]
}

# Address can be allocated from a Range via the parent field
# Dynamically getting the Next available IP in the Range defined by the address field
resource "b1ddi_address" "next_available_range_ip" {
  address =  b1ddi_range.tf_acc_test_range.id
  comment = "Example Address automatically allocated from the Range"
  space   = b1ddi_ip_space.example_tf_space.id
  lifecycle {
    ignore_changes = [address, range]
  }
}

# Address can be allocated from a Address block via the parent field
resource "b1ddi_address" "next_available_address_block_ip" {
  address = b1ddi_address_block.tf_example_address_block.id
  comment = "Example Address automatically allocated from the address_block"
  space   = b1ddi_ip_space.example_tf_space.id
  lifecycle {
    ignore_changes = [address]
  }
}

# Address can be allocated from a subnet via the parent field
# Dynamically getting the Next available IP in the Subnet defined by the address field
resource "b1ddi_address" "next_available_subnet_ip" {
  address  = b1ddi_subnet.example_tf_subnet.id
  comment = "Example Address automatically allocated from the subnet"
  space   = b1ddi_ip_space.example_tf_space.id
  lifecycle {
    ignore_changes = [address]
  }
  depends_on = [b1ddi_address.next_available_address_block_ip]
}

