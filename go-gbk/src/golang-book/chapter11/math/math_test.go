package math
import "testing"

type testpair struct {
    values []float64
    average float64
}

var tests = []testpair{
    { []float64{1,2}, 1.5 },
    { []float64{3,4}, 3.5 },
    { []float64{5,6}, 5.5 }
}

func TestAverage(t *testing.T) {
    // using our structure
    for _, pair := range tests {
        v := Average(pair.values)
        if v != pair.average {
          t.Error("For", pair.values,
          "expected", pair.average,
          "got", v)
        }
    }
}
