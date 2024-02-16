import (
    "math"
    "testing"
)

func TestLinearRegression(t *testing.T) {
    x := []float64{1, 2, 3, 4, 5}
    y := []float64{2, 4, 6, 8, 10}

    beta1, beta0 := LinearRegression(x, y)

    // Check beta1
    expectedBeta1 := 2.0
    if math.Abs(beta1-expectedBeta1) > 1e-6 {
        t.Errorf("Expected beta1 to be %f, but got %f", expectedBeta1, beta1)
    }

    // Check beta0
    expectedBeta0 := 0.0
    if math.Abs(beta0-expectedBeta0) > 1e-6 {
        t.Errorf("Expected beta0 to be %f, but got %f", expectedBeta0, beta0)
    }
}






func TestMean(t *testing.T) {
    values := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
    expected := 3.0

    result := mean(values)
    if result != expected {
        t.Errorf("Mean calculation incorrect. Expected %f, got %f", expected, result)
    }
}




func TestMean(t *testing.T) {
    values := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
    expected := 3.0

    result := mean(values)
    if result != expected {
        t.Errorf("Mean calculation incorrect. Expected %f, got %f", expected, result)
    }
}

