package asyncawait_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/AvnonIdo/asyncawait"
)

func addAndPrintSleepFiveSeconds(num1, num2 int) {
	time.Sleep(5 * time.Second)
	fmt.Println(num1 + num2)
}

func TestAsync0(t *testing.T) {
	startime := time.Now()
	future1 := asyncawait.Async0(func() { addAndPrintSleepFiveSeconds(1, 2) })
	future2 := asyncawait.Async0(func() { addAndPrintSleepFiveSeconds(3, 4) })

	asyncawait.Await0(future1)
	asyncawait.Await0(future2)

	if time.Since(startime) >= 6*time.Second {
		t.Errorf("Test took longer than expected. Expected runtime = ~5 seconds. Actual runtime = %v seconds", time.Since(startime).Seconds())
	}
}

func TestAsync0Method(t *testing.T) {
	startime := time.Now()
	future1 := asyncawait.Async0(func() { addAndPrintSleepFiveSeconds(1, 2) })
	future2 := asyncawait.Async0(func() { addAndPrintSleepFiveSeconds(3, 4) })

	future1.Await()
	future2.Await()

	if time.Since(startime) >= 6*time.Second {
		t.Errorf("Test took longer than expected. Expected runtime = ~5 seconds. Actual runtime = %v seconds", time.Since(startime).Seconds())
	}
}

func addAndSleepFiveSeconds(num1, num2 int) int {
	time.Sleep(5 * time.Second)
	return num1 + num2
}

func TestAsync(t *testing.T) {
	startime := time.Now()
	future1 := asyncawait.Async(func() int { return addAndSleepFiveSeconds(1, 2) })
	future2 := asyncawait.Async(func() int { return addAndSleepFiveSeconds(3, 4) })

	res1, res2 := asyncawait.Await(future1), asyncawait.Await(future2)
	if res1 != 3 || res2 != 7 {
		t.Errorf("Expected results different from actual results. Expected: res1 = 3, res2 = 7. Recieved: res1 = %v, res2 = %v", res1, res2)
	}

	if time.Since(startime) >= 6*time.Second {
		t.Errorf("Test took longer than expected. Expected runtime = ~5 seconds. Actual runtime = %v seconds", time.Since(startime).Seconds())
	}
}
func TestAsyncMethod(t *testing.T) {
	startime := time.Now()
	future1 := asyncawait.Async(func() int { return addAndSleepFiveSeconds(1, 2) })
	future2 := asyncawait.Async(func() int { return addAndSleepFiveSeconds(3, 4) })

	res1, res2 := future1.Await(), future2.Await()
	if res1 != 3 || res2 != 7 {
		t.Errorf("Expected results different from actual results. Expected: res1 = 3, res2 = 7. Recieved: res1 = %v, res2 = %v", res1, res2)
	}

	if time.Since(startime) >= 6*time.Second {
		t.Errorf("Test took longer than expected. Expected runtime = ~5 seconds. Actual runtime = %v seconds", time.Since(startime).Seconds())
	}
}

func addSubAndSleepFiveSeconds(num1, num2 int) (int, int) {
	time.Sleep(5 * time.Second)
	return num1 + num2, num1 - num2
}
func TestAsync2(t *testing.T) {
	startime := time.Now()
	future1 := asyncawait.Async2(func() (int, int) { return addSubAndSleepFiveSeconds(1, 2) })
	future2 := asyncawait.Async2(func() (int, int) { return addSubAndSleepFiveSeconds(3, 5) })

	add1, sub1 := asyncawait.Await2(future1)
	add2, sub2 := asyncawait.Await2(future2)

	if add1 != 3 || sub1 != -1 || add2 != 8 || sub2 != -2 {
		t.Errorf("Expected results different from actual results. Expected: add1 = 3, sub1 = -1, add2 = 8, sub2 = -2. Recieved: add1 = %v, sub1 = %v, add2 = %v, sub2 = %v", add1, sub1, add2, sub2)
	}

	if time.Since(startime) >= 6*time.Second {
		t.Errorf("Test took longer than expected. Expected runtime = ~5 seconds. Actual runtime = %v seconds", time.Since(startime).Seconds())
	}
}

func TestAsync2Method(t *testing.T) {
	startime := time.Now()
	future1 := asyncawait.Async2(func() (int, int) { return addSubAndSleepFiveSeconds(1, 2) })
	future2 := asyncawait.Async2(func() (int, int) { return addSubAndSleepFiveSeconds(3, 5) })

	add1, sub1 := future1.Await()
	add2, sub2 := future2.Await()

	if add1 != 3 || sub1 != -1 || add2 != 8 || sub2 != -2 {
		t.Errorf("Expected results different from actual results. Expected: add1 = 3, sub1 = -1, add2 = 8, sub2 = -2. Recieved: add1 = %v, sub1 = %v, add2 = %v, sub2 = %v", add1, sub1, add2, sub2)
	}

	if time.Since(startime) >= 6*time.Second {
		t.Errorf("Test took longer than expected. Expected runtime = ~5 seconds. Actual runtime = %v seconds", time.Since(startime).Seconds())
	}
}

func addSubMulAndSleepFiveSeconds(num1, num2 int) (int, int, int) {
	time.Sleep(5 * time.Second)
	return num1 + num2, num1 - num2, num1 * num2
}
func TestAsync3(t *testing.T) {
	startime := time.Now()
	future1 := asyncawait.Async3(func() (int, int, int) { return addSubMulAndSleepFiveSeconds(1, 2) })
	future2 := asyncawait.Async3(func() (int, int, int) { return addSubMulAndSleepFiveSeconds(3, 5) })

	add1, sub1, mul1 := asyncawait.Await3(future1)
	add2, sub2, mul2 := asyncawait.Await3(future2)

	if add1 != 3 || sub1 != -1 || mul1 != 2 || add2 != 8 || sub2 != -2 || mul2 != 15 {
		t.Errorf("Expected results different from actual results. Expected: add1 = 3, sub1 = -1, mul1 = 2, add2 = 8, sub2 = -2, mul3 = 15. Recieved: add1 = %v, sub1 = %v, mul1 = %v, add2 = %v, sub2 = %v, mul3 = %v", add1, sub1, mul1, add2, sub2, mul2)
	}

	if time.Since(startime) >= 6*time.Second {
		t.Errorf("Test took longer than expected. Expected runtime = ~5 seconds. Actual runtime = %v seconds", time.Since(startime).Seconds())
	}
}

func TestAsync3Method(t *testing.T) {
	startime := time.Now()
	future1 := asyncawait.Async3(func() (int, int, int) { return addSubMulAndSleepFiveSeconds(1, 2) })
	future2 := asyncawait.Async3(func() (int, int, int) { return addSubMulAndSleepFiveSeconds(3, 5) })

	add1, sub1, mul1 := future1.Await()
	add2, sub2, mul2 := future2.Await()

	if add1 != 3 || sub1 != -1 || mul1 != 2 || add2 != 8 || sub2 != -2 || mul2 != 15 {
		t.Errorf("Expected results different from actual results. Expected: add1 = 3, sub1 = -1, mul1 = 2, add2 = 8, sub2 = -2, mul3 = 15. Recieved: add1 = %v, sub1 = %v, mul1 = %v, add2 = %v, sub2 = %v, mul3 = %v", add1, sub1, mul1, add2, sub2, mul2)
	}

	if time.Since(startime) >= 6*time.Second {
		t.Errorf("Test took longer than expected. Expected runtime = ~5 seconds. Actual runtime = %v seconds", time.Since(startime).Seconds())
	}
}

func addSubMulDivAndSleepFiveSeconds(num1, num2 float64) (float64, float64, float64, float64) {
	time.Sleep(5 * time.Second)
	return num1 + num2, num1 - num2, num1 * num2, num1 / num2
}
func TestAsync4(t *testing.T) {
	startime := time.Now()
	future1 := asyncawait.Async4(func() (float64, float64, float64, float64) { return addSubMulDivAndSleepFiveSeconds(1, 2) })
	future2 := asyncawait.Async4(func() (float64, float64, float64, float64) { return addSubMulDivAndSleepFiveSeconds(3, 5) })

	add1, sub1, mul1, div1 := asyncawait.Await4(future1)
	add2, sub2, mul2, div2 := asyncawait.Await4(future2)

	if add1 != 3 || sub1 != -1 || mul1 != 2 || div1 != 1.0/2 || add2 != 8 || sub2 != -2 || mul2 != 15 || div2 != 3.0/5 {
		t.Errorf("Expected results different from actual results. Expected: add1 = 3, sub1 = -1, mul1 = 2, div1 = 0.5, add2 = 8, sub2 = -2, mul2 = 15, div2 = 0.6. Recieved: add1 = %v, sub1 = %v, mul1 = %v, div1 = %v, add2 = %v, sub2 = %v, mul2 = %v, div2 = %v", add1, sub1, mul1, div1, add2, sub2, mul2, div2)
	}

	if time.Since(startime) >= 6*time.Second {
		t.Errorf("Test took longer than expected. Expected runtime = ~5 seconds. Actual runtime = %v seconds", time.Since(startime).Seconds())
	}
}

func TestAsync4Method(t *testing.T) {
	startime := time.Now()
	future1 := asyncawait.Async4(func() (float64, float64, float64, float64) { return addSubMulDivAndSleepFiveSeconds(1, 2) })
	future2 := asyncawait.Async4(func() (float64, float64, float64, float64) { return addSubMulDivAndSleepFiveSeconds(3, 5) })

	add1, sub1, mul1, div1 := future1.Await()
	add2, sub2, mul2, div2 := future2.Await()

	if add1 != 3 || sub1 != -1 || mul1 != 2 || div1 != 1.0/2 || add2 != 8 || sub2 != -2 || mul2 != 15 || div2 != 3.0/5 {
		t.Errorf("Expected results different from actual results. Expected: add1 = 3, sub1 = -1, mul1 = 2, div1 = 0.5, add2 = 8, sub2 = -2, mul2 = 15, div2 = 0.6. Recieved: add1 = %v, sub1 = %v, mul1 = %v, div1 = %v, add2 = %v, sub2 = %v, mul2 = %v, div2 = %v", add1, sub1, mul1, div1, add2, sub2, mul2, div2)
	}

	if time.Since(startime) >= 6*time.Second {
		t.Errorf("Test took longer than expected. Expected runtime = ~5 seconds. Actual runtime = %v seconds", time.Since(startime).Seconds())
	}
}

func addSubMulDivPowAndSleepFiveSeconds(num1, num2 float64) (float64, float64, float64, float64, float64) {
	time.Sleep(5 * time.Second)
	return num1 + num2, num1 - num2, num1 * num2, num1 / num2, math.Pow(num1, num2)
}
func TestAsync5(t *testing.T) {
	startime := time.Now()
	future1 := asyncawait.Async5(func() (float64, float64, float64, float64, float64) { return addSubMulDivPowAndSleepFiveSeconds(1, 2) })
	future2 := asyncawait.Async5(func() (float64, float64, float64, float64, float64) { return addSubMulDivPowAndSleepFiveSeconds(3, 5) })

	add1, sub1, mul1, div1, pow1 := asyncawait.Await5(future1)
	add2, sub2, mul2, div2, pow2 := asyncawait.Await5(future2)

	if add1 != 3 || sub1 != -1 || mul1 != 2 || div1 != 1.0/2 || pow1 != 1 || add2 != 8 || sub2 != -2 || mul2 != 15 || div2 != 3.0/5 || pow2 != 243 {
		t.Errorf("Expected results different from actual results. Expected: add1 = 3, sub1 = -1, mul1 = 2, div1 = 0.5, pow1 = 1, add2 = 8, sub2 = -2, mul2 = 15, div2 = 0.6, pow2 = 243. Recieved: add1 = %v, sub1 = %v, mul1 = %v, div1 = %v, pow1 = %v, add2 = %v, sub2 = %v, mul2 = %v, div2 = %v, pow2 = %v", add1, sub1, mul1, div1, pow1, add2, sub2, mul2, div2, pow2)
	}

	if time.Since(startime) >= 6*time.Second {
		t.Errorf("Test took longer than expected. Expected runtime = ~5 seconds. Actual runtime = %v seconds", time.Since(startime).Seconds())
	}
}

func TestAsync5Method(t *testing.T) {
	startime := time.Now()
	future1 := asyncawait.Async5(func() (float64, float64, float64, float64, float64) { return addSubMulDivPowAndSleepFiveSeconds(1, 2) })
	future2 := asyncawait.Async5(func() (float64, float64, float64, float64, float64) { return addSubMulDivPowAndSleepFiveSeconds(3, 5) })

	add1, sub1, mul1, div1, pow1 := future1.Await()
	add2, sub2, mul2, div2, pow2 := future2.Await()

	if add1 != 3 || sub1 != -1 || mul1 != 2 || div1 != 1.0/2 || pow1 != 1 || add2 != 8 || sub2 != -2 || mul2 != 15 || div2 != 3.0/5 || pow2 != 243 {
		t.Errorf("Expected results different from actual results. Expected: add1 = 3, sub1 = -1, mul1 = 2, div1 = 0.5, pow1 = 1, add2 = 8, sub2 = -2, mul2 = 15, div2 = 0.6, pow2 = 243. Recieved: add1 = %v, sub1 = %v, mul1 = %v, div1 = %v, pow1 = %v, add2 = %v, sub2 = %v, mul2 = %v, div2 = %v, pow2 = %v", add1, sub1, mul1, div1, pow1, add2, sub2, mul2, div2, pow2)
	}

	if time.Since(startime) >= 6*time.Second {
		t.Errorf("Test took longer than expected. Expected runtime = ~5 seconds. Actual runtime = %v seconds", time.Since(startime).Seconds())
	}
}
