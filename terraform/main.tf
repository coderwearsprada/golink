module "dynamodb" {
  source              = "./dynamodb"
  enabled             = "true"
  environment         = "dev"
  table_name          = "go_link"
  hash_key            = "Short"
  tags                = map("env", "dev")

  attributes = [
    {
      name = "Short"
      type = "S"
    },
    {
      name = "Link"
      type = "S"
    },
    {
      name = "Owner"
      type = "S"
    }
  ]

  global_secondary_index_map = [
    {
      name               = "OwnerIndex"
      hash_key           = "Owner"
      range_key          = "Owner"
      write_capacity     = 5
      read_capacity      = 5
      projection_type    = "INCLUDE"
      non_key_attributes = ["Short", "Link"]
    },
    {
      name               = "LinkIndex"
      hash_key           = "Link"
      range_key          = "Link"
      write_capacity     = 5
      read_capacity      = 5
      projection_type    = "INCLUDE"
      non_key_attributes = ["Short", "Owner"]
    }
  ]
}