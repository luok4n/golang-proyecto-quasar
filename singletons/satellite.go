package singleton

type Satellite struct {
	Name        string
	CoordinateX float32
	CoordinateY float32
	Distance    float32
	Message     []string
}

func (s *Satellite) GetCoordinates() (coordinateX, coordinateY float32) {
	return s.CoordinateX, s.CoordinateY
}

func (s *Satellite) GetName() (name string) {
	return s.Name
}

func (s *Satellite) SetName(name string) {
	s.Name = name
}

func (s *Satellite) SetCoordinates(coordinateX, coordinateY float32) {
	s.CoordinateX = coordinateX
	s.CoordinateY = coordinateY
}

func (s *Satellite) GetDistance() (distance float32) {
	return s.Distance
}

func (s *Satellite) SetDistance(distance float32) {
	s.Distance = distance
}

func (s *Satellite) SetMessage(message []string) {
	s.Message = message
}

func (s *Satellite) GetMessage() (message []string) {
	return s.Message
}
