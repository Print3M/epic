package cli

import "os"

func DisableColors() {
	os.Setenv("NO_COLOR", "true")
}
