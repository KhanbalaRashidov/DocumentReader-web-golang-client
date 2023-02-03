package models

type AreaContainer struct {
	Count  int
	List   []RectangleCoordinates
	Points []PointsContainer
}

func NewAreaContainer(count int, list []RectangleCoordinates, points []PointsContainer) *AreaContainer {
	return &AreaContainer{
		Count:  count,
		List:   list,
		Points: points,
	}
}
