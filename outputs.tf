output "ddos_plan_name" {
  value = azurerm_network_ddos_protection_plan.ddos_plan.name
}

output "ddos_plan_location" {
  value = azurerm_network_ddos_protection_plan.ddos_plan.location
}

output "ddos_plan_tags" {
  value = azurerm_network_ddos_protection_plan.ddos_plan.tags
}