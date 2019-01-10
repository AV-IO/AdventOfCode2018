package sky

import (
	"strings"
)

// Vector contains X/Y positions and velocities in the form of ints
type Vector struct {
	PositionX int
	PositionY int
	VelocityX int
	VelocityY int
}

// Sky contains a slice of vectors, along with min/max positions for the slice's initial state
type Sky struct {
	Field []Vector
	MinX  int
	MinY  int
	MaxX  int
	MaxY  int
}

// PassTime moves each vector based on how many seconds have passed
func (s *Sky) PassTime(seconds int) {
	for _, v := range s.Field {
		if v.PositionX = seconds * v.VelocityX; v.PositionX < s.MinX {
			s.MinX = v.PositionX
		} else if v.PositionX > s.MaxX {
			s.MaxX = v.PositionX
		}
		if v.PositionY = seconds * v.VelocityY; v.PositionY < s.MinY {
			s.MinY = v.PositionY
		} else if v.PositionY > s.MaxY {
			s.MaxY = v.PositionY
		}
	}
}

// AddVector will add a vector to the sky, and adjust min/max values as necessary
func (s *Sky) AddVector(v Vector) {
	// add to list
	s.Field = append(s.Field, v)

	// modify sky min/max values as necessary
	if v.PositionX < s.MinX {
		s.MinX = v.PositionX
	} else if v.PositionX > s.MaxX {
		s.MaxX = v.PositionX
	}
	if v.PositionY < s.MinY {
		s.MinY = v.PositionY
	} else if v.PositionY > s.MaxY {
		s.MaxY = v.PositionY
	}
}

// ToString returns a string displaying the current sky
func (s *Sky) ToString() (str string) {
	//strArr := make([]string, (s.MaxX-s.MinX)*(s.MaxY-s.MinY))
	strArr := make([][]string, s.MaxY-s.MinY)
	for i := 0; i < s.MaxY-s.MinY; i++ {
		strArr[i] = make([]string, s.MaxX-s.MinX)
	}

	for _, v := range s.Field {
		strArr[v.PositionX+s.MinX][v.PositionY+s.MinY] = "*"
	}

	for i := 0; i < s.MaxY-s.MinY; i++ {
		str += strings.Join(strArr[i], "")
	}

	return
}
