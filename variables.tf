variable "ddos_plan_name_strings" {
  type = list
  description = "This should be a list of strings which in conjunction with the seperator make up the resource name"
  default     = null
}

variable "ddos_plan_name_separator" {
  type = string
  description = "Used with name_strings to make up the resource name"
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

variable "ddos_plan_tags" {
  type = map
  description = "(Optional) A mapping of tags to assign to the resource."
  default = null
}

