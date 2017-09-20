// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/gorouter/metrics/reporter"
)

type FakeRouteRegistryReporter struct {
	CaptureRouteStatsStub        func(totalRoutes int, msSinceLastUpdate uint64)
	captureRouteStatsMutex       sync.RWMutex
	captureRouteStatsArgsForCall []struct {
		totalRoutes       int
		msSinceLastUpdate uint64
	}
	CaptureLookupTimeStub        func(t time.Duration)
	captureLookupTimeMutex       sync.RWMutex
	captureLookupTimeArgsForCall []struct {
		t time.Duration
	}
	CaptureRegistryMessageStub        func(msg reporter.ComponentTagged)
	captureRegistryMessageMutex       sync.RWMutex
	captureRegistryMessageArgsForCall []struct {
		msg reporter.ComponentTagged
	}
}

func (fake *FakeRouteRegistryReporter) CaptureRouteStats(totalRoutes int, msSinceLastUpdate uint64) {
	fake.captureRouteStatsMutex.Lock()
	fake.captureRouteStatsArgsForCall = append(fake.captureRouteStatsArgsForCall, struct {
		totalRoutes       int
		msSinceLastUpdate uint64
	}{totalRoutes, msSinceLastUpdate})
	fake.captureRouteStatsMutex.Unlock()
	if fake.CaptureRouteStatsStub != nil {
		fake.CaptureRouteStatsStub(totalRoutes, msSinceLastUpdate)
	}
}

func (fake *FakeRouteRegistryReporter) CaptureRouteStatsCallCount() int {
	fake.captureRouteStatsMutex.RLock()
	defer fake.captureRouteStatsMutex.RUnlock()
	return len(fake.captureRouteStatsArgsForCall)
}

func (fake *FakeRouteRegistryReporter) CaptureRouteStatsArgsForCall(i int) (int, uint64) {
	fake.captureRouteStatsMutex.RLock()
	defer fake.captureRouteStatsMutex.RUnlock()
	return fake.captureRouteStatsArgsForCall[i].totalRoutes, fake.captureRouteStatsArgsForCall[i].msSinceLastUpdate
}

func (fake *FakeRouteRegistryReporter) CaptureLookupTime(t time.Duration) {
	fake.captureLookupTimeMutex.Lock()
	fake.captureLookupTimeArgsForCall = append(fake.captureLookupTimeArgsForCall, struct {
		t time.Duration
	}{t})
	fake.captureLookupTimeMutex.Unlock()
	if fake.CaptureLookupTimeStub != nil {
		fake.CaptureLookupTimeStub(t)
	}
}

func (fake *FakeRouteRegistryReporter) CaptureLookupTimeCallCount() int {
	fake.captureLookupTimeMutex.RLock()
	defer fake.captureLookupTimeMutex.RUnlock()
	return len(fake.captureLookupTimeArgsForCall)
}

func (fake *FakeRouteRegistryReporter) CaptureLookupTimeArgsForCall(i int) time.Duration {
	fake.captureLookupTimeMutex.RLock()
	defer fake.captureLookupTimeMutex.RUnlock()
	return fake.captureLookupTimeArgsForCall[i].t
}

func (fake *FakeRouteRegistryReporter) CaptureRegistryMessage(msg reporter.ComponentTagged) {
	fake.captureRegistryMessageMutex.Lock()
	fake.captureRegistryMessageArgsForCall = append(fake.captureRegistryMessageArgsForCall, struct {
		msg reporter.ComponentTagged
	}{msg})
	fake.captureRegistryMessageMutex.Unlock()
	if fake.CaptureRegistryMessageStub != nil {
		fake.CaptureRegistryMessageStub(msg)
	}
}

func (fake *FakeRouteRegistryReporter) CaptureRegistryMessageCallCount() int {
	fake.captureRegistryMessageMutex.RLock()
	defer fake.captureRegistryMessageMutex.RUnlock()
	return len(fake.captureRegistryMessageArgsForCall)
}

func (fake *FakeRouteRegistryReporter) CaptureRegistryMessageArgsForCall(i int) reporter.ComponentTagged {
	fake.captureRegistryMessageMutex.RLock()
	defer fake.captureRegistryMessageMutex.RUnlock()
	return fake.captureRegistryMessageArgsForCall[i].msg
}

var _ reporter.RouteRegistryReporter = new(FakeRouteRegistryReporter)
