package asyncawait

type tuple2[T, W any] struct {
	value1 T
	value2 W
}
type tuple3[T, W, C any] struct {
	value1 T
	value2 W
	value3 C
}
type tuple4[T, W, C, D any] struct {
	value1 T
	value2 W
	value3 C
	value4 D
}
type tuple5[T, W, C, D, E any] struct {
	value1 T
	value2 W
	value3 C
	value4 D
	value5 E
}

type Future[T any] struct {
	output chan T
}
type Future2[T, W any] struct {
	output chan tuple2[T, W]
}
type Future3[T, W, C any] struct {
	output chan tuple3[T, W, C]
}
type Future4[T, W, C, D any] struct {
	output chan tuple4[T, W, C, D]
}
type Future5[T, W, C, D, E any] struct {
	output chan tuple5[T, W, C, D, E]
}

func Async[T any](fn func() T) Future[T] {
	output := make(chan T)
	go func() {
		result := fn()
		output <- result
	}()
	return Future[T]{
		output,
	}
}

func Async2[T, W any](fn func() (T, W)) Future2[T, W] {
	output := make(chan tuple2[T, W])
	go func() {
		res1, res2 := fn()
		resultTuple := tuple2[T, W]{
			value1: res1,
			value2: res2,
		}
		output <- resultTuple
	}()
	return Future2[T, W]{
		output,
	}
}

func Async3[T, W, C any](fn func() (T, W, C)) Future3[T, W, C] {
	output := make(chan tuple3[T, W, C])
	go func() {
		res1, res2, res3 := fn()
		resultTuple := tuple3[T, W, C]{
			value1: res1,
			value2: res2,
			value3: res3,
		}
		output <- resultTuple
	}()
	return Future3[T, W, C]{
		output,
	}
}

func Async4[T, W, C, D any](fn func() (T, W, C, D)) Future4[T, W, C, D] {
	output := make(chan tuple4[T, W, C, D])
	go func() {
		res1, res2, res3, res4 := fn()
		resultTuple := tuple4[T, W, C, D]{
			value1: res1,
			value2: res2,
			value3: res3,
			value4: res4,
		}
		output <- resultTuple
	}()
	return Future4[T, W, C, D]{
		output,
	}
}

func Async5[T, W, C, D, E any](fn func() (T, W, C, D, E)) Future5[T, W, C, D, E] {
	output := make(chan tuple5[T, W, C, D, E])
	go func() {
		res1, res2, res3, res4, res5 := fn()
		resultTuple := tuple5[T, W, C, D, E]{
			value1: res1,
			value2: res2,
			value3: res3,
			value4: res4,
			value5: res5,
		}
		output <- resultTuple
	}()
	return Future5[T, W, C, D, E]{
		output,
	}
}

func Await[T any](future Future[T]) T {
	result, ok := <-future.output
	if !ok {
		panic("you can't await the same future twice!")
	}
	close(future.output)

	return result
}

func Await2[T, W any](future Future2[T, W]) (T, W) {
	result, ok := <-future.output
	if !ok {
		panic("you can't await the same future twice!")
	}
	close(future.output)

	return result.value1, result.value2
}

func Await3[T, W, C any](future Future3[T, W, C]) (T, W, C) {
	result, ok := <-future.output
	if !ok {
		panic("you can't await the same future twice!")
	}
	close(future.output)

	return result.value1, result.value2, result.value3
}

func Await4[T, W, C, D any](future Future4[T, W, C, D]) (T, W, C, D) {
	result, ok := <-future.output
	if !ok {
		panic("you can't await the same future twice!")
	}
	close(future.output)

	return result.value1, result.value2, result.value3, result.value4
}

func Await5[T, W, C, D, E any](future Future5[T, W, C, D, E]) (T, W, C, D, E) {
	result, ok := <-future.output
	if !ok {
		panic("you can't await the same future twice!")
	}
	close(future.output)

	return result.value1, result.value2, result.value3, result.value4, result.value5
}

func (f Future[T]) Await() T {
	return Await(f)
}

func (f Future2[T, W]) Await() (T, W) {
	return Await2(f)
}

func (f Future3[T, W, C]) Await() (T, W, C) {
	return Await3(f)
}

func (f Future4[T, W, C, D]) Await() (T, W, C, D) {
	return Await4(f)
}

func (f Future5[T, W, C, D, E]) Await() (T, W, C, D, E) {
	return Await5(f)
}
