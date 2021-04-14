variable "slack_webhook_url" {
  type        = string
  description = "full url to the slack webhook"
}

variable "tags" {
  default     = {}
  type        = map(any)
  description = "Map of different tags which are applied to all resources"
}
