package algebra

type OrderedPair struct {
	X         int64
	Y         int64
	IsPlotted bool
}

func (o *OrderedPair) Quadrant() string {
	if o.X == 0 && o.Y == 0 {
		return "Origin"
	} else if o.X > 0 && o.Y == 0 {
		return "I and IV"
	} else if o.X < 0 && o.Y == 0 {
		return "II and III"
	} else if o.X == 0 && o.Y > 0 {
		return "I and II"
	} else if o.X == 0 && o.Y < 0 {
		return "III and IV"
	} else if o.X > 0 && o.Y > 0 {
		return "I"
	} else if o.X < 0 && o.Y > 0 {
		return "II"
	} else if o.X < 0 && o.Y < 0 {
		return "III"
	} else if o.X > 0 && o.Y < 0 {
		return "IV"
	}
	return "UNKNOWN"
}
