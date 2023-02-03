package models

type AreaArray struct {
	List   []RectangleCoordinates
	Points []PointArray
}

func NewAreaArray(list []RectangleCoordinates, points []PointArray) *AreaArray {
	return &AreaArray{
		List:   list,
		Points: points,
	}
}
