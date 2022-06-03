package section

type Section struct {
	id                  int
	section_number      int
	current_temperature float64
	minimum_temperature float64
	current_capacity    int
	minimum_capacity    int
	maximum_capacity    int
	warehouse_id        int
	product_type_id     int
}
