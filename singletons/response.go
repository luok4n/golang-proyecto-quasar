package singleton

type ResultPosition struct {
	Position Positions `json:"position"`
	Message  string    `json:"message"`
}

type Positions struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func (r *ResultPosition) SetPositions(position Positions) {
	r.Position = position
}

func (p *Positions) SetXY(x, y float32) {
	p.X = x
	p.Y = y
}

func (r *ResultPosition) SetMessage(message string) {
	r.Message = message
}

type SatellitesBody struct {
	Satellites []SatelliteBody `json:"satellites"`
}

type SatelliteBody struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

func (r *SatellitesBody) GetSatellites() (satellites []SatelliteBody) {
	return r.Satellites
}
