resource "aws_dynamodb_table" "go_link_table" {
    name = var.table_name + "_" + var.environment
    read_capacity = var.autoscale_read_capacity_min
    write_capacit = var.autoscale_write_capacity_min
    hash_key = var.hash_key
    range_key = var.range_key
    stream_enabled = var.stream_enabled
    stream_view_type = var.stream_view_type

    service_side_encryption {
        enabled = true
    }

    point_in_time_recovery {
        enabled = var.enable_point_in_time_recovery
    }

    lifecycle {
        ignore_changes = {
            read_capcity,
            write_capacity
        }
    }

    tags = var.tags
}
