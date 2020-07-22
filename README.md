#  Azure Virtual Network DDOS Protection Plan

## Background
This module deploys the Virtual Network DDOS Protection Plan

## usage

Its quite easy to use the module, just by supplying the required parameters. There is (as with most Azure Modules) a pre-requisite for a resource group, so this either needs to be a data resource or another modular deployment, which is how the example code below was written.

```
module "resourcegroup" {
  source = "bytejunkie/resource-group/azurerm"

  name_strings   = var.resource_group_name_strings
  name_separator = var.resource_group_name_separator
}

module "ddos_plan" {
  source  = "bytejunkie/network-ddos-protection-plan/azurerm"
# insert the required variables here
  ddos_plan_name_strings        = var.ddos_plan_name_strings
  ddos_plan_name_separator      = var.ddos_plan_name_separator
  ddos_plan_location            = var.ddos_plan_location
  resource_group_name = module.resourcegroup.resource_group_name
# insert the optional variables here
  ddos_plan_tags = var.ddos_plan_tags
}
```

## Required Parameters

``` hcl
variable "ddos_plan_name_strings" {
  type = list
  description = "(Required)This should be a list of strings which in conjunction with the seperator make up the resource group name"
  default     = null
}

variable "ddos_plan_name_separator" {
  type = string
  description = "(Required)Used with name_strings to make up the resource group name"
  default     = null
}

variable "ddos_plan_location" {
  description = "(Required) Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created."
  default = null
}

variable "resource_group_name" {
  description = "(Required) The name of the resource group in which to create the resource. Changing this forces a new resource to be created."
  default = null
}
```

## Optional Parameters

```hcl
variable "ddos_plan_tags" {
  type = map
  description = "(Optional) A mapping of tags to assign to the resource."
  default = null
}
```
