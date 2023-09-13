package entity

// buat generic function provider
func NewEntity[T any]() *T {
	var entity T
	return &entity
}

func NewSubDistrict() *SubDistrict {
	return &SubDistrict{}
}
