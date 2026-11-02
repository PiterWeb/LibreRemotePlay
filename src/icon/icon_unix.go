//go:build linux || darwin

package icon

import (
	_ "embed"
)

//go:embed lrp.png
var Data []byte