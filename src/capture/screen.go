// Copyright 2020 Donovan Solms github.com/donovansolms twitter.com/donovansolms
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package capture

import (
	"fmt"
	"image"

	"github.com/kbinani/screenshot"
)

// Screen controls capturing screenshots
type Screen struct {
}

// NewScreen creates a new instance of screen capturing
func NewScreen() *Screen {
	screen := &Screen{}
	return screen
}

// CaptureFullscreen captures the entire screen and returns the image
func (screen *Screen) CaptureFullscreen(displayID int) (image.Image, error) {
	numDisplays := screenshot.NumActiveDisplays()
	if displayID > numDisplays {
		return nil, fmt.Errorf("Display ID of '%d' is too large", displayID)
	}
	if displayID < 0 {
		return nil, fmt.Errorf("Display ID of '%d' is too small", displayID)
	}

	bounds := screenshot.GetDisplayBounds(displayID)
	return screen.CaptureBounds(displayID, bounds)
}

// CaptureBounds captures the given rectange and returns the image
func (screen *Screen) CaptureBounds(displayID int, bounds image.Rectangle) (image.Image, error) {
	numDisplays := screenshot.NumActiveDisplays()
	if displayID > numDisplays {
		return nil, fmt.Errorf("Display ID of '%d' is too large", displayID)
	}
	if displayID < 0 {
		return nil, fmt.Errorf("Display ID of '%d' is too small", displayID)
	}

	return screenshot.CaptureRect(bounds)
}
