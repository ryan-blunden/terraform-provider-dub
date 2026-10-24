data "dub_tag" "my_tag" {
  array_of_str = [
    "..."
  ]
  page       = 1
  page_size  = 50
  search     = "...my_search..."
  sort_by    = "name"
  sort_order = "asc"
  str        = "...my_str..."
}