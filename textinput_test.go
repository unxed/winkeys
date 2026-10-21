package winkeys

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestCommittedTextRetainsPhysicalMetadata(t *testing.T) {
	event := InputEvent{
		Type: KeyEventType, KeyDown: true, TextInput: true,
		VirtualKeyCode: VK_Q, VirtualScanCode: 0x10,
		Char: '@', UnshiftedChar: 'q', RepeatCount: 3,
		ControlKeyState: RightAltPressed | LeftCtrlPressed,
		InputSource:     "win32",
	}
	data, err := json.Marshal(event)
	if err != nil {
		t.Fatal(err)
	}
	var decoded InputEvent
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decoded, event) {
		t.Fatalf("text event changed during serialization: %+v", decoded)
	}
	if !strings.Contains(event.String(), "[TextInput]") {
		t.Fatalf("committed text missing from diagnostics: %s", event)
	}
}

func TestLegacyEventDefaultsToKeyIntent(t *testing.T) {
	var event InputEvent
	if err := json.Unmarshal([]byte(`{"Type":1,"KeyDown":true,"Char":64}`), &event); err != nil {
		t.Fatal(err)
	}
	if event.TextInput || strings.Contains(event.String(), "[TextInput]") {
		t.Fatal("legacy key data unexpectedly acquired text intent")
	}
}
