package sub

// Food is a type to show how type documentation works.  Note that
// conventionally, documentation on a type, function, or method starts with the
// name of that thing.
//
// Formatting
//
// Yep, you get formatting on type, function, and method docs, too.
//
// Below is the auto-embedded example.
type Food struct {
	// Name holds the common name of the food.
	Name string

	// weight is an unexported field that won't show up in documentation.
	weight int
}

// NewFood creates a new Food with the given name and a calculated weight.
//
// Note that you do not need to (and should not) make separate lines describing
// each parameter to the function.  Simply refer to the parameters by their name
// when you need to refernce them.
func NewFood(name string) Food {
	return Food{name, len(name)}
}

// Weight reports the weight of the food in letters.
//
// When decribing a function without parameters that simply returns data, it is
// traditional to write the doc saying the function reports foo, rather than
// *returns* foo.
//
// The example below intentionally fails during go test, to show you what that
// looks like.  The test will report:
//
//   --- FAIL: ExampleFood_Weight (4.933us)
//   got:
//   9
//   want:
//   10
//   FAIL
//   exit status 1
func (f Food) Weight() int {
	return f.weight
}
