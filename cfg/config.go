// Package cfg contains structs
// that will hold on all needful parameters for our app
// that will be retrieved from  .env or ./cfg/config.yml
package cfg

// Options will keep all needful configs
type Options struct{}

// NewOptions will create instance of Options
// that will be used im main package
func NewOptions() Options {
	return Options{}
}
