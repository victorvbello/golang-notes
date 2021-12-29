// Package mystring provide a custom string funcions
package mystring

// MyStringJoin function get a slice of string an join using (,)
func MyStringJoin(xs ...string) string {
	result := xs[0]
	for _, v := range xs[1:] {
		result += "," + v
	}
	return result
}
