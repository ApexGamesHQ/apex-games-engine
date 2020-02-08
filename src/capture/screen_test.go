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

package capture_test

import (
	"image"
	"testing"

	"github.com/ApexGamesHQ/apex-games-engine/src/capture"
)

func TestCaptureTooLargeDisplayID(t *testing.T) {
	displayID := 65000
	bounds := image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: 100,
			Y: 100,
		},
	}

	screenCapture := capture.NewScreen()
	_, err := screenCapture.CaptureFullscreen(displayID)
	if err == nil {
		t.Errorf("Expected display ID to be out of bounds, got nil")
	}

	_, err = screenCapture.CaptureBounds(displayID, bounds)
	if err == nil {
		t.Errorf("Expected display ID to be out of bounds, got nil")
	}
}

func TestCaptureInvalidDisplayID(t *testing.T) {
	displayID := -1
	bounds := image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: 100,
			Y: 100,
		},
	}

	screenCapture := capture.NewScreen()
	_, err := screenCapture.CaptureFullscreen(displayID)
	if err == nil {
		t.Errorf("Expected display ID to be out of bounds, got nil")
	}

	_, err = screenCapture.CaptureBounds(displayID, bounds)
	if err == nil {
		t.Errorf("Expected display ID to be out of bounds, got nil")
	}
}

func TestCaptureFullscreen(t *testing.T) {
	displayID := 0
	screenCapture := capture.NewScreen()
	img, err := screenCapture.CaptureFullscreen(displayID)
	if err != nil {
		t.Errorf("Failed to capture fullscreen: %s", err)
	}
	if img == nil {
		t.Errorf("Fullscreen capture is nil")
	}
}

func TestCaptureBounds(t *testing.T) {
	displayID := 0
	screenCapture := capture.NewScreen()
	bounds := image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: 100,
			Y: 150,
		},
	}
	img, err := screenCapture.CaptureBounds(displayID, bounds)
	if err != nil {
		t.Errorf("Failed to capture fullscreen: %s", err)
	}
	if img == nil {
		t.Errorf("Bounds capture is nil")
	}

	expectedWidth := bounds.Max.X - bounds.Min.X
	if img.Bounds().Dx() != expectedWidth {
		t.Errorf(
			"Bounds capture is too large, expected width of %d, got %d",
			expectedWidth,
			img.Bounds().Dx(),
		)
	}

	expectedHeight := bounds.Max.Y - bounds.Min.Y
	if img.Bounds().Dy() != expectedHeight {
		t.Errorf(
			"Bounds capture is too large, expected width of %d, got %d",
			expectedHeight,
			img.Bounds().Dy(),
		)
	}
}
