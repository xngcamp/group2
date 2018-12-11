package movie

type Movie interface {
	GetPrice(int) float64
	GetPoints(int) int
	GetName() string
}

type MovieCommon struct {
	Name string
}

func (m MovieCommon) GetName() string {
	return m.Name
}

type CommonMovie struct {
	MovieCommon
}

func (m *CommonMovie) GetPrice(days int) float64 {
	result := 2.0
	if days > 2 {
		result += float64(days - 2) * 1.5
	}
	return result
}

func (m *CommonMovie) GetPoints(days int) int {
	return 1
}

type ChildrenMovie struct {
	MovieCommon
}

func (m *ChildrenMovie) GetPrice(days int) float64 {
	result := 1.5
	if days > 3 {
		result += float64(days - 3) * 1.5
	}
	return result
}

func (m *ChildrenMovie) GetPoints(days int) int {
	return 1
}

type NewReleaseMovie struct {
	MovieCommon
}

func (m *NewReleaseMovie) GetPrice(days int) float64 {
	return 2 * float64(days)
}

func (m *NewReleaseMovie) GetPoints(days int) int {
	result := 1
	if days > 1 {
		result += 1
	}
	return result
}