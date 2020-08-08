variable "table_name" {
    type = string
    default = ""
    description = "Name of the dynamdb table"
}
variable "environment" {
  type        = string
  default     = ""
  description = "Environment, e.g. 'prod', 'qa', 'dev'"
}

variable "enabled" {
  type        = bool
  default     = true
  description = "Set to false to prevent the module from creating any resources"
}

variable "tags" {
  type        = map(string)
  default     = {}
  description = "Additional tags (e.g. `map('BusinessUnit','XYZ')`"
}

variable "autoscale_write_target" {
  type        = number
  default     = 50
  description = "The target value (in %) for DynamoDB write autoscaling"
}

variable "autoscale_read_target" {
  type        = number
  default     = 50
  description = "The target value (in %) for DynamoDB read autoscaling"
}

variable "autoscale_read_capacity_min" {
  type        = number
  default     = 5
  description = "DynamoDB autoscaling min read capacity"
}

variable "autoscale_read_capacity_max" {
  type        = number
  default     = 20
  description = "DynamoDB autoscaling max read capacity"
}

variable "autoscale_write_capacity_min" {
  type        = number
  default     = 5
  description = "DynamoDB autoscaling min write capacity"
}

variable "autoscale_write_capacity_max" {
  type        = number
  default     = 20
  description = "DynamoDB autoscaling max write capacity"
}

variable "billing_mode" {
  type        = string
  default     = "PROVISIONED"
  description = "DynamoDB Billing mode. Can be PROVISIONED or PAY_PER_REQUEST"
}

variable "stream_enabled" {
  type        = bool
  default     = false
  description = "Enable DynamoDB stream"
}

variable "stream_view_type" {
  type        = string
  default     = ""
  description = "When an item in the table is modified, what information is written to the stream"
}

variable "enable_point_in_time_recovery" {
  type        = bool
  default     = true
  description = "Enable DynamoDB point in time recovery"
}

variable "hash_key" {
  type        = string
  description = "DynamoDB table Hash Key"
}

variable "hash_key_type" {
  type        = string
  default     = "S"
  description = "Hash Key type, which must be a scalar type: `S`, `N`, or `B` for (S)tring, (N)umber or (B)inary data"
}

variable "range_key" {
  type        = string
  default     = ""
  description = "DynamoDB table Range Key"
}

variable "range_key_type" {
  type        = string
  default     = "S"
  description = "Range Key type, which must be a scalar type: `S`, `N`, or `B` for (S)tring, (N)umber or (B)inary data"
}

variable "ttl_attribute" {
  type        = string
  default     = "Expires"
  description = "DynamoDB table TTL attribute"
}

variable "enable_autoscaler" {
  type        = bool
  default     = false
  description = "Enable or disable DynamoDB autoscaling"
}