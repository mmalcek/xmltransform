package main

import "testing"

func TestLuaF(t *testing.T) {
	if luaF("sum", "5", "5") != "10" {
		t.Errorf("luaF unexpected result: %s", luaF("sum", "5", "5"))
	}
}

func TestDateFormat(t *testing.T) {
	if dateFormat("15.03.2021", "02.01.2006", "01022006") != "03152021" {
		t.Errorf("dateFormat unexpected result: %s", dateFormat("15.03.2021", "02.01.2006", "01022006"))
	}
}

func TestAdd(t *testing.T) {
	if add("6", 4) != 10 {
		t.Errorf("dateFormat unexpected result: %v", add("6", 4))
	}
}

func TestToInt(t *testing.T) {
	if toInt("42") != 42 {
		t.Errorf("toInt unexpected result: %d", toInt("42"))
	}
}

func TestRegexMatch(t *testing.T) {
	if !regexMatch(`a.b`, "aaxbb") {
		t.Errorf("regexMatch unexpected result: %v", regexMatch(`^a.b$`, "aaxbb"))
	}
}
