package phono_test

import (
	"sync"
	"testing"
	"time"

	"github.com/dudk/phono"
	"github.com/stretchr/testify/assert"
)

var sliceTests = []struct {
	in       phono.Buffer
	start    int64
	len      int
	expected phono.Buffer
}{
	{
		in:       phono.Buffer([][]float64{[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}),
		start:    1,
		len:      2,
		expected: phono.Buffer([][]float64{[]float64{1, 2}, []float64{1, 2}}),
	},
	{
		in:       phono.Buffer([][]float64{[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}),
		start:    5,
		len:      2,
		expected: phono.Buffer([][]float64{[]float64{5, 6}, []float64{5, 6}}),
	},
	{
		in:       phono.Buffer([][]float64{[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}),
		start:    7,
		len:      4,
		expected: phono.Buffer([][]float64{[]float64{7, 8, 9}, []float64{7, 8, 9}}),
	},
	{
		in:       phono.Buffer([][]float64{[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}),
		start:    9,
		len:      1,
		expected: phono.Buffer([][]float64{[]float64{9}}),
	},
	{
		in:       phono.Buffer([][]float64{[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}),
		start:    10,
		len:      1,
		expected: nil,
	},
}

func TestBuffer(t *testing.T) {
	var s phono.Buffer
	assert.Equal(t, phono.NumChannels(0), s.NumChannels())
	assert.Equal(t, phono.BufferSize(0), s.Size())
	s = [][]float64{[]float64{}}
	assert.Equal(t, phono.NumChannels(1), s.NumChannels())
	assert.Equal(t, phono.BufferSize(0), s.Size())
	s[0] = make([]float64, 512)
	assert.Equal(t, phono.BufferSize(512), s.Size())

	s2 := [][]float64{make([]float64, 512)}
	s = s.Append(s2)
	assert.Equal(t, phono.BufferSize(1024), s.Size())
	s2[0] = make([]float64, 1024)
	s = s.Append(s2)
	assert.Equal(t, phono.BufferSize(2048), s.Size())
}

func TestSliceBuffer(t *testing.T) {
	for _, test := range sliceTests {
		result := test.in.Slice(test.start, test.len)
		assert.Equal(t, test.expected.Size(), result.Size())
		assert.Equal(t, test.expected.NumChannels(), result.NumChannels())
		for i := range test.expected {
			for j := 0; j < len(test.expected[i]); j++ {
				assert.Equal(t, test.expected[i][j], result[i][j])
			}
		}
	}
}

func TestSampleRate(t *testing.T) {
	sampleRate := phono.SampleRate(44100)
	expected := 500 * time.Millisecond
	result := sampleRate.DurationOf(22050)
	assert.Equal(t, expected, result)
}

func TestSingleUse(t *testing.T) {
	var once sync.Once
	err := phono.SingleUse(&once)
	assert.Nil(t, err)
	err = phono.SingleUse(&once)
	assert.Equal(t, phono.ErrSingleUseReused, err)
}

func TestDuration(t *testing.T) {
	var tests = []struct {
		sampleRate phono.SampleRate
		samples    int64
		expected   time.Duration
	}{
		{
			sampleRate: 44100,
			samples:    44100,
			expected:   1 * time.Second,
		},
		{
			sampleRate: 44100,
			samples:    22050,
			expected:   500 * time.Millisecond,
		},
		{
			sampleRate: 44100,
			samples:    50,
			expected:   1133786 * time.Nanosecond,
		},
	}
	for _, c := range tests {
		assert.Equal(t, c.expected, c.sampleRate.DurationOf(c.samples))
	}
}

func TestReadInts(t *testing.T) {
	var tests = []struct {
		ints []int
		phono.NumChannels
		expected phono.Buffer
	}{}
	for test := range tests {

	}
}
