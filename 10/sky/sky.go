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
	minX  int
	minY  int
	maxX  int
	maxY  int
}

// PassTime moves each vector based on how many seconds have passed
func (s *Sky) PassTime(seconds int) {
	for _, v := range s.Field {
		if v.PositionX = seconds * v.VelocityX; v.PositionX < s.minX {
			s.minX = v.PositionX
		} else if v.PositionX > s.maxX {
			s.maxX = v.PositionX
		}
		if v.PositionY = seconds * v.VelocityY; v.PositionY < s.minY {
			s.minY = v.PositionY
		} else if v.PositionY > s.maxY {
			s.maxY = v.PositionY
		}
	}
}

// AddVector will add a vector to the sky, and adjust min/max values as necessary
func (s *Sky) AddVector(v Vector) {
	// add to list
	s.Field = append(s.Field, v)

	// modify sky min/max values as necessary
	if v.PositionX < s.minX {
		s.minX = v.PositionX
	} else if v.PositionX > s.maxX {
		s.maxX = v.PositionX
	}
	if v.PositionY < s.minY {
		s.minY = v.PositionY
	} else if v.PositionY > s.maxY {
		s.maxY = v.PositionY
	}
}

// ToString returns a string displaying the current sky
func (s *Sky) ToString() (str string) {
	//strArr := make([]string, (s.maxX-s.minX)*(s.maxY-s.minY))
	strArr := make([][]string, s.maxY-s.minY)
	for i := 0; i < s.maxY-s.minY; i++ {
		strArr[i] = make([]string, s.maxX-s.minX)
	}

	for _, v := range s.Field {
		strArr[v.PositionX+s.minX][v.PositionY+s.minY] = "*"
	}

	for i := 0; i < s.maxY-s.minY; i++ {
		str += strings.Join(strArr[i], "")
	}

	return
}
