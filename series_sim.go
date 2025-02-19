package main

type SeriesSim struct {
	cntr          float32
	start         float32
	end           float32
	numPoints     float32
	seriesChannel chan float32
}

func NewSeriesSim(start, end float32, numPoints float32) *SeriesSim {
	return &SeriesSim{start, start, end, numPoints, make(chan float32, 1)}
}

func (s *SeriesSim) GetChannel() chan float32 {
	return s.seriesChannel
}

func (s *SeriesSim) Generate() {
	if s.cntr > s.end {
		close(s.seriesChannel)
		return
	}
	s.seriesChannel <- s.cntr
	diff := float32((s.end - s.start) / (s.numPoints))
	s.cntr += diff
}

/*
func main() {
	series := NewSeriesSim(5, 10, 10)

	channel := series.GetChannel()
	series.Generate()
	for val := range channel {
		fmt.Println(val)

		series.Generate()
	}
	fmt.Println(<-channel)
}
*/
