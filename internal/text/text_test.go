package text

import (
	"testing"
)

func TestMsgValidIDs(t *testing.T) {
	// Test valid message IDs
	for id := MsgNum(1); id <= maxMsgNum; id++ {
		result := Msg(id)
		if result == "" {
			t.Errorf("Message ID %d returned empty string", id)
		}
		// Should not be an error message
		if len(result) >= 6 && result[:6] == "<<OOPS" {
			t.Errorf("Message ID %d returned error format: %s", id, result)
		}
	}
}

func TestMsgInvalidIDs(t *testing.T) {
	invalidIDs := []MsgNum{0, -1, maxMsgNum + 1, maxMsgNum + 100}

	for _, id := range invalidIDs {
		result := Msg(id)
		// Should return error format
		if len(result) < 6 || result[:6] != "<<OOPS" {
			t.Errorf("Message ID %d should return error format, got: %s", id, result)
		}
	}
}

func TestMsgFormatting(t *testing.T) {
	// We can't test specific formatting without knowing which IDs have format flags
	// But we can test that calling with args doesn't crash
	for id := MsgNum(1); id <= maxMsgNum; id++ {
		// Try with various argument types
		result1 := Msg(id, "test")
		result2 := Msg(id, 42)
		result3 := Msg(id, "test", 42, true)

		// Should never be empty (even error messages have content)
		if result1 == "" || result2 == "" || result3 == "" {
			t.Errorf("Message ID %d returned empty string with arguments", id)
		}
	}
}

func TestMaxMsgNumIsValid(t *testing.T) {
	if maxMsgNum < 1 {
		t.Error("maxMsgNum should be at least 1")
	}

	// Test that maxMsgNum is actually accessible
	result := Msg(maxMsgNum)
	if len(result) >= 6 && result[:6] == "<<OOPS" {
		t.Errorf("maxMsgNum %d should be valid, got error: %s", maxMsgNum, result)
	}
}

func TestMsgSliceBounds(t *testing.T) {
	// msgs slice should have maxMsgNum+1 elements (including index 0)
	expectedLen := int(maxMsgNum) + 1
	if len(msgs) != expectedLen {
		t.Errorf("msgs slice length %d, expected %d", len(msgs), expectedLen)
	}

	// Index 0 should be empty placeholder
	if msgs[0].t != "" || msgs[0].f != false {
		t.Error("msgs[0] should be empty placeholder")
	}
}
