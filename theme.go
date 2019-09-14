package ext

import "os"

// SetThemeLight set light theme
func SetThemeLight() {
	os.Setenv("FYNE_THEME", "light")
}
