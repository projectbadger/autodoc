package config_test

import (
	"fmt"

	"github.com/projectbadger/autodoc/config"
)

func Example() {
	// config package is initialized immediately,
	// so values can be called directly
	fmt.Println(config.Cfg.PackageDir)
	// Output: .
}
