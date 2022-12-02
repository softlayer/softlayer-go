package session

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestSession_WithDefaultUserAgent(t *testing.T) {
	s := New()
	expected := getDefaultUserAgent()
	actual := s.userAgent
	if actual != expected {
		t.Errorf("UserAgent expected %s, actual %s", expected, actual)
	}
}

func TestSession_WithOSENV(t *testing.T) {

	old, exists := os.LookupEnv("SL_USERAGENT")
	os.Setenv("SL_USERAGENT", "session_test")
	expected := getDefaultUserAgent()

	if !strings.Contains(expected, "session_test") {
		t.Errorf("UserAgent '%s' does not contain %s", expected, "session_test")
	}
	if exists {
		os.Setenv("SL_USERAGENT", old)
	} else {
		os.Unsetenv("SL_USERAGENT")
	}
}

func TestAddToUserAgent(t *testing.T) {
	s := New()
	defaultUA := getDefaultUserAgent()
	ua := "product/v1 ( myproduct )"
	s.AppendUserAgent(ua)
	if !strings.HasSuffix(s.userAgent, ua) {
		t.Errorf("UserAgent expected to end with %s, but ends with %s", ua, s.userAgent)
	}
	if !strings.HasPrefix(s.userAgent, defaultUA) {
		t.Errorf("UserAgent expected to start with %s, but starts with %s", defaultUA, s.userAgent)
	}
	oldUa := s.userAgent
	s.AppendUserAgent("")
	if oldUa != s.userAgent {
		t.Errorf("UserAgent changed. Old: '%s', New: '%s'", oldUa, s.userAgent)
	}
}

func TestResetUserAgent(t *testing.T) {
	s := New()
	ua := "product/v1 ( myproduct )"
	s.AppendUserAgent(ua)
	s.ResetUserAgent()
	if s.userAgent != getDefaultUserAgent() {
		t.Errorf("UserAgent expected to reset to %s, but found to be %s", getDefaultUserAgent(), s.userAgent)
	}
}

func TestSetTimeout(t *testing.T) {
	s := New()
	newTimeout, err := time.ParseDuration("192s")
	if err != nil {
		t.Errorf("Failed to get a time duration: %v", err)
	}
	s = s.SetTimeout(newTimeout)

	if s.Timeout != newTimeout {
		t.Errorf("Session.Timeout (%v) != newTimeout (%v)", s.Timeout, newTimeout)
	}
}

func TestSetRetryWait(t *testing.T) {
	s := New()
	newTimeout, err := time.ParseDuration("192s")
	if err != nil {
		t.Errorf("Failed to get a time duration: %v", err)
	}
	s = s.SetRetryWait(newTimeout)

	if s.RetryWait != newTimeout {
		t.Errorf("Session.RetryWait (%v) != newTimeout (%v)", s.RetryWait, newTimeout)
	}
}

func TestSetRetries(t *testing.T) {
	s := New()
	newVariable := 10

	s = s.SetRetries(newVariable)

	if s.Retries != newVariable {
		t.Errorf("Session.Retries (%v) != newVariable (%v)", s.Retries, newVariable)
	}
}
