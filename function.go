// return type declaration
func functionName() int {
	return 42
	}
	// Can return multiple values at once
	func returnMulti() (int, string) {
	return 42, "foobar"
	}
	var x, str = returnMulti()
	// Return multiple named results simply by return
	func returnMulti2() (n int, s string) {
	n = 42
	s = "foobar"
	// n and s will be returned
	return
	}
	var x, str = returnMulti2()
	Functions As Values And Closures
	func main() {
	// assign a function to a name
	add := func(a, b int) int {
	return a + b
	}
	// use the name to call the function
	fmt.Println(add(3, 4))
	}
	// Closures, lexically scoped: Functions can access values that were
	// in scope when defining the function
	func scope() func() int{
	outer_var := 2
	foo := func() int { return outer_var}
	return foo
	}
	func another_scope() func() int{
	// won't compile - outer_var and foo not defined in this scope
	outer_var = 444
	return foo
	}
	// Closures: don't mutate outer vars, instead redefine them!
	func outer() (func() int, int) {
	outer_var := 2
	// NOTE outer_var is outside inner's scope
	inner := func() int {
	outer_var += 99 // attempt to mutate outer_var
	return outer_var // => 101 (but outer_var is a newly redefined
	// variable visible only inside inner)
	}
	return inner, outer_var // => 101, 2 (still!)
	}