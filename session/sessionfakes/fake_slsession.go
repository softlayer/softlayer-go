// Code generated by counterfeiter. DO NOT EDIT.
package sessionfakes

import (
	"sync"
	"time"

	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

type FakeSLSession struct {
	AppendUserAgentStub        func(string)
	appendUserAgentMutex       sync.RWMutex
	appendUserAgentArgsForCall []struct {
		arg1 string
	}
	DoRequestStub        func(string, string, []interface{}, *sl.Options, interface{}) error
	doRequestMutex       sync.RWMutex
	doRequestArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []interface{}
		arg4 *sl.Options
		arg5 interface{}
	}
	doRequestReturns struct {
		result1 error
	}
	doRequestReturnsOnCall map[int]struct {
		result1 error
	}
	ResetUserAgentStub        func()
	resetUserAgentMutex       sync.RWMutex
	resetUserAgentArgsForCall []struct {
	}
	SetRetriesStub        func(int) *session.Session
	setRetriesMutex       sync.RWMutex
	setRetriesArgsForCall []struct {
		arg1 int
	}
	setRetriesReturns struct {
		result1 *session.Session
	}
	setRetriesReturnsOnCall map[int]struct {
		result1 *session.Session
	}
	SetRetryWaitStub        func(time.Duration) *session.Session
	setRetryWaitMutex       sync.RWMutex
	setRetryWaitArgsForCall []struct {
		arg1 time.Duration
	}
	setRetryWaitReturns struct {
		result1 *session.Session
	}
	setRetryWaitReturnsOnCall map[int]struct {
		result1 *session.Session
	}
	SetTimeoutStub        func(time.Duration) *session.Session
	setTimeoutMutex       sync.RWMutex
	setTimeoutArgsForCall []struct {
		arg1 time.Duration
	}
	setTimeoutReturns struct {
		result1 *session.Session
	}
	setTimeoutReturnsOnCall map[int]struct {
		result1 *session.Session
	}
	StringStub        func() string
	stringMutex       sync.RWMutex
	stringArgsForCall []struct {
	}
	stringReturns struct {
		result1 string
	}
	stringReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSLSession) AppendUserAgent(arg1 string) {
	fake.appendUserAgentMutex.Lock()
	fake.appendUserAgentArgsForCall = append(fake.appendUserAgentArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.AppendUserAgentStub
	fake.recordInvocation("AppendUserAgent", []interface{}{arg1})
	fake.appendUserAgentMutex.Unlock()
	if stub != nil {
		fake.AppendUserAgentStub(arg1)
	}
}

func (fake *FakeSLSession) AppendUserAgentCallCount() int {
	fake.appendUserAgentMutex.RLock()
	defer fake.appendUserAgentMutex.RUnlock()
	return len(fake.appendUserAgentArgsForCall)
}

func (fake *FakeSLSession) AppendUserAgentCalls(stub func(string)) {
	fake.appendUserAgentMutex.Lock()
	defer fake.appendUserAgentMutex.Unlock()
	fake.AppendUserAgentStub = stub
}

func (fake *FakeSLSession) AppendUserAgentArgsForCall(i int) string {
	fake.appendUserAgentMutex.RLock()
	defer fake.appendUserAgentMutex.RUnlock()
	argsForCall := fake.appendUserAgentArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSLSession) DoRequest(arg1 string, arg2 string, arg3 []interface{}, arg4 *sl.Options, arg5 interface{}) error {
	var arg3Copy []interface{}
	if arg3 != nil {
		arg3Copy = make([]interface{}, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.doRequestMutex.Lock()
	ret, specificReturn := fake.doRequestReturnsOnCall[len(fake.doRequestArgsForCall)]
	fake.doRequestArgsForCall = append(fake.doRequestArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []interface{}
		arg4 *sl.Options
		arg5 interface{}
	}{arg1, arg2, arg3Copy, arg4, arg5})
	stub := fake.DoRequestStub
	fakeReturns := fake.doRequestReturns
	fake.recordInvocation("DoRequest", []interface{}{arg1, arg2, arg3Copy, arg4, arg5})
	fake.doRequestMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSLSession) DoRequestCallCount() int {
	fake.doRequestMutex.RLock()
	defer fake.doRequestMutex.RUnlock()
	return len(fake.doRequestArgsForCall)
}

func (fake *FakeSLSession) DoRequestCalls(stub func(string, string, []interface{}, *sl.Options, interface{}) error) {
	fake.doRequestMutex.Lock()
	defer fake.doRequestMutex.Unlock()
	fake.DoRequestStub = stub
}

func (fake *FakeSLSession) DoRequestArgsForCall(i int) (string, string, []interface{}, *sl.Options, interface{}) {
	fake.doRequestMutex.RLock()
	defer fake.doRequestMutex.RUnlock()
	argsForCall := fake.doRequestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeSLSession) DoRequestReturns(result1 error) {
	fake.doRequestMutex.Lock()
	defer fake.doRequestMutex.Unlock()
	fake.DoRequestStub = nil
	fake.doRequestReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSLSession) DoRequestReturnsOnCall(i int, result1 error) {
	fake.doRequestMutex.Lock()
	defer fake.doRequestMutex.Unlock()
	fake.DoRequestStub = nil
	if fake.doRequestReturnsOnCall == nil {
		fake.doRequestReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.doRequestReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSLSession) ResetUserAgent() {
	fake.resetUserAgentMutex.Lock()
	fake.resetUserAgentArgsForCall = append(fake.resetUserAgentArgsForCall, struct {
	}{})
	stub := fake.ResetUserAgentStub
	fake.recordInvocation("ResetUserAgent", []interface{}{})
	fake.resetUserAgentMutex.Unlock()
	if stub != nil {
		fake.ResetUserAgentStub()
	}
}

func (fake *FakeSLSession) ResetUserAgentCallCount() int {
	fake.resetUserAgentMutex.RLock()
	defer fake.resetUserAgentMutex.RUnlock()
	return len(fake.resetUserAgentArgsForCall)
}

func (fake *FakeSLSession) ResetUserAgentCalls(stub func()) {
	fake.resetUserAgentMutex.Lock()
	defer fake.resetUserAgentMutex.Unlock()
	fake.ResetUserAgentStub = stub
}

func (fake *FakeSLSession) SetRetries(arg1 int) *session.Session {
	fake.setRetriesMutex.Lock()
	ret, specificReturn := fake.setRetriesReturnsOnCall[len(fake.setRetriesArgsForCall)]
	fake.setRetriesArgsForCall = append(fake.setRetriesArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.SetRetriesStub
	fakeReturns := fake.setRetriesReturns
	fake.recordInvocation("SetRetries", []interface{}{arg1})
	fake.setRetriesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSLSession) SetRetriesCallCount() int {
	fake.setRetriesMutex.RLock()
	defer fake.setRetriesMutex.RUnlock()
	return len(fake.setRetriesArgsForCall)
}

func (fake *FakeSLSession) SetRetriesCalls(stub func(int) *session.Session) {
	fake.setRetriesMutex.Lock()
	defer fake.setRetriesMutex.Unlock()
	fake.SetRetriesStub = stub
}

func (fake *FakeSLSession) SetRetriesArgsForCall(i int) int {
	fake.setRetriesMutex.RLock()
	defer fake.setRetriesMutex.RUnlock()
	argsForCall := fake.setRetriesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSLSession) SetRetriesReturns(result1 *session.Session) {
	fake.setRetriesMutex.Lock()
	defer fake.setRetriesMutex.Unlock()
	fake.SetRetriesStub = nil
	fake.setRetriesReturns = struct {
		result1 *session.Session
	}{result1}
}

func (fake *FakeSLSession) SetRetriesReturnsOnCall(i int, result1 *session.Session) {
	fake.setRetriesMutex.Lock()
	defer fake.setRetriesMutex.Unlock()
	fake.SetRetriesStub = nil
	if fake.setRetriesReturnsOnCall == nil {
		fake.setRetriesReturnsOnCall = make(map[int]struct {
			result1 *session.Session
		})
	}
	fake.setRetriesReturnsOnCall[i] = struct {
		result1 *session.Session
	}{result1}
}

func (fake *FakeSLSession) SetRetryWait(arg1 time.Duration) *session.Session {
	fake.setRetryWaitMutex.Lock()
	ret, specificReturn := fake.setRetryWaitReturnsOnCall[len(fake.setRetryWaitArgsForCall)]
	fake.setRetryWaitArgsForCall = append(fake.setRetryWaitArgsForCall, struct {
		arg1 time.Duration
	}{arg1})
	stub := fake.SetRetryWaitStub
	fakeReturns := fake.setRetryWaitReturns
	fake.recordInvocation("SetRetryWait", []interface{}{arg1})
	fake.setRetryWaitMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSLSession) SetRetryWaitCallCount() int {
	fake.setRetryWaitMutex.RLock()
	defer fake.setRetryWaitMutex.RUnlock()
	return len(fake.setRetryWaitArgsForCall)
}

func (fake *FakeSLSession) SetRetryWaitCalls(stub func(time.Duration) *session.Session) {
	fake.setRetryWaitMutex.Lock()
	defer fake.setRetryWaitMutex.Unlock()
	fake.SetRetryWaitStub = stub
}

func (fake *FakeSLSession) SetRetryWaitArgsForCall(i int) time.Duration {
	fake.setRetryWaitMutex.RLock()
	defer fake.setRetryWaitMutex.RUnlock()
	argsForCall := fake.setRetryWaitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSLSession) SetRetryWaitReturns(result1 *session.Session) {
	fake.setRetryWaitMutex.Lock()
	defer fake.setRetryWaitMutex.Unlock()
	fake.SetRetryWaitStub = nil
	fake.setRetryWaitReturns = struct {
		result1 *session.Session
	}{result1}
}

func (fake *FakeSLSession) SetRetryWaitReturnsOnCall(i int, result1 *session.Session) {
	fake.setRetryWaitMutex.Lock()
	defer fake.setRetryWaitMutex.Unlock()
	fake.SetRetryWaitStub = nil
	if fake.setRetryWaitReturnsOnCall == nil {
		fake.setRetryWaitReturnsOnCall = make(map[int]struct {
			result1 *session.Session
		})
	}
	fake.setRetryWaitReturnsOnCall[i] = struct {
		result1 *session.Session
	}{result1}
}

func (fake *FakeSLSession) SetTimeout(arg1 time.Duration) *session.Session {
	fake.setTimeoutMutex.Lock()
	ret, specificReturn := fake.setTimeoutReturnsOnCall[len(fake.setTimeoutArgsForCall)]
	fake.setTimeoutArgsForCall = append(fake.setTimeoutArgsForCall, struct {
		arg1 time.Duration
	}{arg1})
	stub := fake.SetTimeoutStub
	fakeReturns := fake.setTimeoutReturns
	fake.recordInvocation("SetTimeout", []interface{}{arg1})
	fake.setTimeoutMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSLSession) SetTimeoutCallCount() int {
	fake.setTimeoutMutex.RLock()
	defer fake.setTimeoutMutex.RUnlock()
	return len(fake.setTimeoutArgsForCall)
}

func (fake *FakeSLSession) SetTimeoutCalls(stub func(time.Duration) *session.Session) {
	fake.setTimeoutMutex.Lock()
	defer fake.setTimeoutMutex.Unlock()
	fake.SetTimeoutStub = stub
}

func (fake *FakeSLSession) SetTimeoutArgsForCall(i int) time.Duration {
	fake.setTimeoutMutex.RLock()
	defer fake.setTimeoutMutex.RUnlock()
	argsForCall := fake.setTimeoutArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSLSession) SetTimeoutReturns(result1 *session.Session) {
	fake.setTimeoutMutex.Lock()
	defer fake.setTimeoutMutex.Unlock()
	fake.SetTimeoutStub = nil
	fake.setTimeoutReturns = struct {
		result1 *session.Session
	}{result1}
}

func (fake *FakeSLSession) SetTimeoutReturnsOnCall(i int, result1 *session.Session) {
	fake.setTimeoutMutex.Lock()
	defer fake.setTimeoutMutex.Unlock()
	fake.SetTimeoutStub = nil
	if fake.setTimeoutReturnsOnCall == nil {
		fake.setTimeoutReturnsOnCall = make(map[int]struct {
			result1 *session.Session
		})
	}
	fake.setTimeoutReturnsOnCall[i] = struct {
		result1 *session.Session
	}{result1}
}

func (fake *FakeSLSession) String() string {
	fake.stringMutex.Lock()
	ret, specificReturn := fake.stringReturnsOnCall[len(fake.stringArgsForCall)]
	fake.stringArgsForCall = append(fake.stringArgsForCall, struct {
	}{})
	stub := fake.StringStub
	fakeReturns := fake.stringReturns
	fake.recordInvocation("String", []interface{}{})
	fake.stringMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSLSession) StringCallCount() int {
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	return len(fake.stringArgsForCall)
}

func (fake *FakeSLSession) StringCalls(stub func() string) {
	fake.stringMutex.Lock()
	defer fake.stringMutex.Unlock()
	fake.StringStub = stub
}

func (fake *FakeSLSession) StringReturns(result1 string) {
	fake.stringMutex.Lock()
	defer fake.stringMutex.Unlock()
	fake.StringStub = nil
	fake.stringReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSLSession) StringReturnsOnCall(i int, result1 string) {
	fake.stringMutex.Lock()
	defer fake.stringMutex.Unlock()
	fake.StringStub = nil
	if fake.stringReturnsOnCall == nil {
		fake.stringReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.stringReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSLSession) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appendUserAgentMutex.RLock()
	defer fake.appendUserAgentMutex.RUnlock()
	fake.doRequestMutex.RLock()
	defer fake.doRequestMutex.RUnlock()
	fake.resetUserAgentMutex.RLock()
	defer fake.resetUserAgentMutex.RUnlock()
	fake.setRetriesMutex.RLock()
	defer fake.setRetriesMutex.RUnlock()
	fake.setRetryWaitMutex.RLock()
	defer fake.setRetryWaitMutex.RUnlock()
	fake.setTimeoutMutex.RLock()
	defer fake.setTimeoutMutex.RUnlock()
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSLSession) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ session.SLSession = new(FakeSLSession)
