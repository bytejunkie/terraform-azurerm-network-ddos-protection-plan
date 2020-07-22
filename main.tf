resource "azurerm_network_ddos_protection_plan" "ddos_plan" {
  name     = join( var.ddos_plan_name_separator, var.ddos_plan_name_strings)
  location = var.ddos_plan_location
  resource_group_name = var.resource_group_name

  tags = var.ddos_plan_tags != "" ? var.ddos_plan_tags : null
}
