module "resourcegroup" {
  source = "bytejunkie/resource-group/azurerm"

  name_strings   = var.resource_group_name_strings
  name_separator = var.resource_group_name_separator
}

module "ddos_plan" {
  source = "../"
# insert the required variables here
  ddos_plan_name_strings        = var.ddos_plan_name_strings
  ddos_plan_name_separator      = var.ddos_plan_name_separator
  ddos_plan_location            = var.ddos_plan_location
  resource_group_name = module.resourcegroup.resource_group_name
# insert the optional variables here
  ddos_plan_tags = var.ddos_plan_tags
}