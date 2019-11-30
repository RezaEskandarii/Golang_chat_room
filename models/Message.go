package models

// TODo: Create ChatMessage and Drawing struct
type Message struct {
	// chat - drawing
	Type string `json:"type"`

	// chat fields .
	ID       uint   `json:"id"`
	Body     string `json:"body"`
	Nickname string `json:"nickname"`
	Image    string


	// drawing fields
	PointX    float64 `json:"point_x"`
	PointY    float64 `json:"point_y"`
	PointLeft float64 `json:"point_left"`
	PointTop  float64 `json:"point_top"`
}
