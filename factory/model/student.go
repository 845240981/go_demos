package model


type student struct {
			Name string
			score float64
}

func (s student) GetScore() float64 {
		return s.score
}

func GetStudent(name string,score float64) student{
			return student{name,score}
} 