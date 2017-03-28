package statistics

//Counter counter
type Counter struct {
	sample []float64
	data   map[float64]int
}

//NewCounter initializae a Counter
func NewCounter(sample []float64) Counter {
	counter := Counter{}
	counter.sample = sample
	counter.data = map[float64]int{}
	make(counter)

	return counter
}

func make(counter Counter) {
	for _, key := range counter.sample {
		if _, ok := counter.data[key]; ok {
			counter.data[key]++
		} else {
			counter.data[key] = 1
		}
	}
}

//Values Return the Values counted
func (counter Counter) Values() []int {
	values := []int{}
	for _, value := range counter.data {
		values = append(values, value)
	}
	return values

}
