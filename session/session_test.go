package session

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/softlayer/softlayer-go/sl"
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

func TestNeedsRefresh(t *testing.T) {
	testError := sl.Error{
		StatusCode: 500,
		Exception:  "SoftLayer_Exception_Account_Authentication_AccessTokenValidation",
	}
	if !NeedsRefresh(testError) {
		t.Errorf("NeedsRefresh() failed to detect refresh error")
	}
	testError = sl.Error{
		StatusCode: 500,
		Exception:  "SoftLayer_Exception_Public",
	}
	if NeedsRefresh(testError) {
		t.Errorf("NeedsRefresh() failed to properly error check")
	}
}

// Tests refreshing a IAM token if it is expired.
func TestRefreshToken(t *testing.T) {
	// setup session and mock environment
	s = New()
	s.Endpoint = restEndpoint
	s.IAMToken = "Bearer TestToken"
	s.IAMRefreshToken = "TestTokenRefresh"
	//s.Debug = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	fmt.Printf("TestRefreshToken [Happy Path]: ")
	expectedError := ""
	// Happy Path
	httpmock.RegisterResponder("POST", IBMCLOUDIAMENDPOINT,
		httpmock.NewStringResponder(200, `{"access_token": "NewToken123", "refresh_token":"NewRefreshToken123", "token_type":"Bearer"}`),
	)
	err := s.RefreshToken()
	if err != nil {
		t.Errorf("Testing Error: %v\n", err.Error())
	}

	if s.IAMToken != "Bearer NewToken123" {
		t.Errorf("(IAMToken) %s != 'Bearer NewToken123', Refresh Failed.", s.IAMToken)
	}
	if s.IAMRefreshToken != "NewRefreshToken123" {
		t.Errorf("(IAMRefreshToken) %s != 'NewRefreshToken123', Refresh Failed.", s.IAMRefreshToken)
	}
	httpmock.Reset()

	// Error returned from IAM API
	fmt.Printf("TestRefreshToken [API error]: ")
	httpmock.RegisterResponder("POST", IBMCLOUDIAMENDPOINT,
		httpmock.NewStringResponder(400, `{"errormessage": "Some Error", "errorcode":"400"}`),
	)
	err = s.RefreshToken()
	if err == nil {
		t.Errorf("Expected an error, none returned\n")
	}
	expectedError = "400: Some Error "
	if err.Error() != expectedError {
		t.Errorf("Expected |%s| == %s", err.Error(), expectedError)
	}
	httpmock.Reset()

	// Junk returned from IAM API
	fmt.Printf("TestRefreshToken [Bad Response]: ")
	httpmock.RegisterResponder("POST", IBMCLOUDIAMENDPOINT,
		httpmock.NewStringResponder(200, ""),
	)
	err = s.RefreshToken()
	if err == nil {
		t.Errorf("Expected an error, none returned\n")
	}
	expectedError = "unexpected end of JSON input"
	if err.Error() != expectedError {
		t.Errorf("Expected %s == %s", err.Error(), expectedError)
	}
	httpmock.Reset()
}
