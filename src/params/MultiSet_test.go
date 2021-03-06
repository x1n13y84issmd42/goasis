package params_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/oasis/src/errors"
	"github.com/x1n13y84issmd42/oasis/src/params"
)

func Test_MultiSet(T *testing.T) {
	T.Run("Load&Iterate", func(T *testing.T) {
		src1 := params.NewMemorySource("test_1")
		src1.Add("A", "The aye")
		src1.Add("B", "The bee")
		src1.Add("C", "The sea")
		src1.Add("D", "The D")

		src2 := params.NewMemorySource("TEST 2")
		src2.Add("A", "The aye aye")
		src2.Add("D", "The D #2")

		set := params.NewMultiSet("test")
		set.Load(src1)
		set.Load(src2)

		expected := []string{
			"A test_1 The aye",
			"A TEST 2 The aye aye",
			"B test_1 The bee",
			"C test_1 The sea",
			"D test_1 The D",
			"D TEST 2 The D #2",
		}

		actual := []string{}

		for p := range set.Iterate() {
			actual = append(actual, p.N+" "+p.Source+" "+p.V())
		}

		assert.Equal(T, expected, actual)
	})

	T.Run("Require&Validate Fail", func(T *testing.T) {
		src := params.NewMemorySource("test")
		src.Add("A", "The aye")
		src.Add("B", "The bee")

		set := params.NewMultiSet("7357")
		set.Load(src)

		set.Require("A")
		set.Require("C")
		set.Require("D")
		set.Require("D")
		set.Require("D")

		expected := errors.NoParameters([]string{"C", "D"}, "7357", nil)
		// Otherwise TheCaller points to this ^ place.
		expected.TheCaller = "github.com/x1n13y84issmd42/oasis/src/params/MultiSet.go:90"

		assert.Equal(T, expected, set.Validate())
	})

	T.Run("Require&Validate Success", func(T *testing.T) {
		src := params.NewMemorySource("test")
		src.Add("A", "The aye")
		src.Add("B", "The bee")

		set := params.NewMultiSet("test")
		set.Load(src)

		set.Require("A")
		set.Require("B")
		set.Require("B")
		set.Require("B")
		set.Require("B")

		assert.Nil(T, set.Validate())
	})
}
