package gocolor

// Prepare is a convenience method that calls EnableConsole
// and returns an enabled color if there's no error and startEnabled is True.
func Prepare(startEnabled bool) Color {
	err := EnableConsole()
	if err != nil || !startEnabled {
		return Color{}
	}
	col := Color{}
	col.Enable()
	return col
}

// EnableConsole enables color printing on Windows and is a no-op
// on non-Windows platforms (which enable color printing by default).
func EnableConsole() error {
	return enableConsole()
}

// Add colors a message and resets the color afterwards
func (c *Color) Add(color Code, message string) string {
	return string(color) + message + string(c.Default)
}

// ColorFunc colors a message and resets the color afterwards. Use Func() to generate a ColorFunc with a specific color
type ColorFunc = func(message string) string

// Func generates a ColorFunc from a color code. It uses the address of the code instead
// of the value because the code's value might change (for example to "" when Disable() is called)
func (c *Color) Func(codePtr *Code) ColorFunc {
	return func(message string) string {
		return c.Add(*codePtr, message)
	}
}

// Code is a color code string
type Code string

// Empty represents the absence of a color code. Used for Disable
const Empty = Code("")
