locals {
  attributes = concat(
    [
      {
        name = var.range_key
        type = var.range_key_type
      },
      {
        name = var.hash_key
        type = var.hash_key_type
      }
    ],
    var.attributes
  )

  # Remove the first map from the list if no `range_key` is provided
  from_index = length(var.range_key) > 0 ? 0 : 1

  attributes_final = slice(local.attributes, local.from_index, length(local.attributes))
}

resource "aws_dynamodb_table" "go_link_table" {
    name = join("_", [var.table_name, var.environment])
    read_capacity = var.autoscale_read_capacity_min
    write_capacity = var.autoscale_write_capacity_min
    hash_key = var.hash_key
    range_key = var.range_key
    stream_enabled = var.stream_enabled
    stream_view_type = var.stream_view_type

    server_side_encryption {
        enabled = true
    }

    point_in_time_recovery {
        enabled = var.enable_point_in_time_recovery
    }

    lifecycle {
        ignore_changes = [
            read_capacity,
            write_capacity
        ]
    }

    dynamic "attribute" {
      for_each = local.attributes_final
      content {
        name = attribute.value.name
        type = attribute.value.type
      }
    }

    dynamic "global_secondary_index" {
      for_each = var.global_secondary_index_map
      content {
        hash_key           = global_secondary_index.value.hash_key
        name               = global_secondary_index.value.name
        non_key_attributes = lookup(global_secondary_index.value, "non_key_attributes", null)
        projection_type    = global_secondary_index.value.projection_type
        range_key          = lookup(global_secondary_index.value, "range_key", null)
        read_capacity      = lookup(global_secondary_index.value, "read_capacity", null)
        write_capacity     = lookup(global_secondary_index.value, "write_capacity", null)
      }
    }
/*
    dynamic "local_secondary_index" {
      for_each = var.local_secondary_index_map
      content {
        name               = local_secondary_index.value.name
        non_key_attributes = lookup(local_secondary_index.value, "non_key_attributes", null)
        projection_type    = local_secondary_index.value.projection_type
        range_key          = local_secondary_index.value.range_key
      }
    }

    ttl {
      attribute_name = var.ttl_attribute
      enabled        = var.ttl_attribute != "" && var.ttl_attribute != null ? true : false
    }
*/
    tags = var.tags
}
