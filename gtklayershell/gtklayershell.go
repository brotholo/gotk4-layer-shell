package gtklayershell

import (
	_ "runtime/cgo"
)

// #cgo pkg-config: gtk4-layer-shell-0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gtk4-layer-shell/gtk4-layer-shell.h>
import "C"
