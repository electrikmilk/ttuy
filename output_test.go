/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"testing"
)

func TestOutput(t *testing.T) {
	// Success
	Success("Done!")
	var success = "Done!"
	Successf("%s", success)
	// Info
	Info("FYI!")
	var info = "FYI!"
	Infof("%s", info)
	// Warning
	Warn("Unable to do something!")
	var warning = "Unable to do something!"
	Warnf("%s", warning)
	// Error
	// Fail("Unable to do something!")
	// var _, err = os.ReadDir("doesntExist")
	// FailErr("Label", err)
	// var message = "Unable to do something!"
	// Failf("%s", message)
}
