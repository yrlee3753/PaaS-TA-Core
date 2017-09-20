// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/consuladapter"
	"github.com/hashicorp/consul/api"
)

type FakeAgent struct {
	ChecksStub        func() (map[string]*api.AgentCheck, error)
	checksMutex       sync.RWMutex
	checksArgsForCall []struct{}
	checksReturns     struct {
		result1 map[string]*api.AgentCheck
		result2 error
	}
	ServicesStub        func() (map[string]*api.AgentService, error)
	servicesMutex       sync.RWMutex
	servicesArgsForCall []struct{}
	servicesReturns     struct {
		result1 map[string]*api.AgentService
		result2 error
	}
	ServiceRegisterStub        func(service *api.AgentServiceRegistration) error
	serviceRegisterMutex       sync.RWMutex
	serviceRegisterArgsForCall []struct {
		service *api.AgentServiceRegistration
	}
	serviceRegisterReturns struct {
		result1 error
	}
	ServiceDeregisterStub        func(serviceID string) error
	serviceDeregisterMutex       sync.RWMutex
	serviceDeregisterArgsForCall []struct {
		serviceID string
	}
	serviceDeregisterReturns struct {
		result1 error
	}
	PassTTLStub        func(checkID, note string) error
	passTTLMutex       sync.RWMutex
	passTTLArgsForCall []struct {
		checkID string
		note    string
	}
	passTTLReturns struct {
		result1 error
	}
	WarnTTLStub        func(checkID, note string) error
	warnTTLMutex       sync.RWMutex
	warnTTLArgsForCall []struct {
		checkID string
		note    string
	}
	warnTTLReturns struct {
		result1 error
	}
	FailTTLStub        func(checkID, note string) error
	failTTLMutex       sync.RWMutex
	failTTLArgsForCall []struct {
		checkID string
		note    string
	}
	failTTLReturns struct {
		result1 error
	}
	NodeNameStub        func() (string, error)
	nodeNameMutex       sync.RWMutex
	nodeNameArgsForCall []struct{}
	nodeNameReturns     struct {
		result1 string
		result2 error
	}
	CheckDeregisterStub        func(checkID string) error
	checkDeregisterMutex       sync.RWMutex
	checkDeregisterArgsForCall []struct {
		checkID string
	}
	checkDeregisterReturns struct {
		result1 error
	}
}

func (fake *FakeAgent) Checks() (map[string]*api.AgentCheck, error) {
	fake.checksMutex.Lock()
	fake.checksArgsForCall = append(fake.checksArgsForCall, struct{}{})
	fake.checksMutex.Unlock()
	if fake.ChecksStub != nil {
		return fake.ChecksStub()
	} else {
		return fake.checksReturns.result1, fake.checksReturns.result2
	}
}

func (fake *FakeAgent) ChecksCallCount() int {
	fake.checksMutex.RLock()
	defer fake.checksMutex.RUnlock()
	return len(fake.checksArgsForCall)
}

func (fake *FakeAgent) ChecksReturns(result1 map[string]*api.AgentCheck, result2 error) {
	fake.ChecksStub = nil
	fake.checksReturns = struct {
		result1 map[string]*api.AgentCheck
		result2 error
	}{result1, result2}
}

func (fake *FakeAgent) Services() (map[string]*api.AgentService, error) {
	fake.servicesMutex.Lock()
	fake.servicesArgsForCall = append(fake.servicesArgsForCall, struct{}{})
	fake.servicesMutex.Unlock()
	if fake.ServicesStub != nil {
		return fake.ServicesStub()
	} else {
		return fake.servicesReturns.result1, fake.servicesReturns.result2
	}
}

func (fake *FakeAgent) ServicesCallCount() int {
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	return len(fake.servicesArgsForCall)
}

func (fake *FakeAgent) ServicesReturns(result1 map[string]*api.AgentService, result2 error) {
	fake.ServicesStub = nil
	fake.servicesReturns = struct {
		result1 map[string]*api.AgentService
		result2 error
	}{result1, result2}
}

func (fake *FakeAgent) ServiceRegister(service *api.AgentServiceRegistration) error {
	fake.serviceRegisterMutex.Lock()
	fake.serviceRegisterArgsForCall = append(fake.serviceRegisterArgsForCall, struct {
		service *api.AgentServiceRegistration
	}{service})
	fake.serviceRegisterMutex.Unlock()
	if fake.ServiceRegisterStub != nil {
		return fake.ServiceRegisterStub(service)
	} else {
		return fake.serviceRegisterReturns.result1
	}
}

func (fake *FakeAgent) ServiceRegisterCallCount() int {
	fake.serviceRegisterMutex.RLock()
	defer fake.serviceRegisterMutex.RUnlock()
	return len(fake.serviceRegisterArgsForCall)
}

func (fake *FakeAgent) ServiceRegisterArgsForCall(i int) *api.AgentServiceRegistration {
	fake.serviceRegisterMutex.RLock()
	defer fake.serviceRegisterMutex.RUnlock()
	return fake.serviceRegisterArgsForCall[i].service
}

func (fake *FakeAgent) ServiceRegisterReturns(result1 error) {
	fake.ServiceRegisterStub = nil
	fake.serviceRegisterReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgent) ServiceDeregister(serviceID string) error {
	fake.serviceDeregisterMutex.Lock()
	fake.serviceDeregisterArgsForCall = append(fake.serviceDeregisterArgsForCall, struct {
		serviceID string
	}{serviceID})
	fake.serviceDeregisterMutex.Unlock()
	if fake.ServiceDeregisterStub != nil {
		return fake.ServiceDeregisterStub(serviceID)
	} else {
		return fake.serviceDeregisterReturns.result1
	}
}

func (fake *FakeAgent) ServiceDeregisterCallCount() int {
	fake.serviceDeregisterMutex.RLock()
	defer fake.serviceDeregisterMutex.RUnlock()
	return len(fake.serviceDeregisterArgsForCall)
}

func (fake *FakeAgent) ServiceDeregisterArgsForCall(i int) string {
	fake.serviceDeregisterMutex.RLock()
	defer fake.serviceDeregisterMutex.RUnlock()
	return fake.serviceDeregisterArgsForCall[i].serviceID
}

func (fake *FakeAgent) ServiceDeregisterReturns(result1 error) {
	fake.ServiceDeregisterStub = nil
	fake.serviceDeregisterReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgent) PassTTL(checkID string, note string) error {
	fake.passTTLMutex.Lock()
	fake.passTTLArgsForCall = append(fake.passTTLArgsForCall, struct {
		checkID string
		note    string
	}{checkID, note})
	fake.passTTLMutex.Unlock()
	if fake.PassTTLStub != nil {
		return fake.PassTTLStub(checkID, note)
	} else {
		return fake.passTTLReturns.result1
	}
}

func (fake *FakeAgent) PassTTLCallCount() int {
	fake.passTTLMutex.RLock()
	defer fake.passTTLMutex.RUnlock()
	return len(fake.passTTLArgsForCall)
}

func (fake *FakeAgent) PassTTLArgsForCall(i int) (string, string) {
	fake.passTTLMutex.RLock()
	defer fake.passTTLMutex.RUnlock()
	return fake.passTTLArgsForCall[i].checkID, fake.passTTLArgsForCall[i].note
}

func (fake *FakeAgent) PassTTLReturns(result1 error) {
	fake.PassTTLStub = nil
	fake.passTTLReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgent) WarnTTL(checkID string, note string) error {
	fake.warnTTLMutex.Lock()
	fake.warnTTLArgsForCall = append(fake.warnTTLArgsForCall, struct {
		checkID string
		note    string
	}{checkID, note})
	fake.warnTTLMutex.Unlock()
	if fake.WarnTTLStub != nil {
		return fake.WarnTTLStub(checkID, note)
	} else {
		return fake.warnTTLReturns.result1
	}
}

func (fake *FakeAgent) WarnTTLCallCount() int {
	fake.warnTTLMutex.RLock()
	defer fake.warnTTLMutex.RUnlock()
	return len(fake.warnTTLArgsForCall)
}

func (fake *FakeAgent) WarnTTLArgsForCall(i int) (string, string) {
	fake.warnTTLMutex.RLock()
	defer fake.warnTTLMutex.RUnlock()
	return fake.warnTTLArgsForCall[i].checkID, fake.warnTTLArgsForCall[i].note
}

func (fake *FakeAgent) WarnTTLReturns(result1 error) {
	fake.WarnTTLStub = nil
	fake.warnTTLReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgent) FailTTL(checkID string, note string) error {
	fake.failTTLMutex.Lock()
	fake.failTTLArgsForCall = append(fake.failTTLArgsForCall, struct {
		checkID string
		note    string
	}{checkID, note})
	fake.failTTLMutex.Unlock()
	if fake.FailTTLStub != nil {
		return fake.FailTTLStub(checkID, note)
	} else {
		return fake.failTTLReturns.result1
	}
}

func (fake *FakeAgent) FailTTLCallCount() int {
	fake.failTTLMutex.RLock()
	defer fake.failTTLMutex.RUnlock()
	return len(fake.failTTLArgsForCall)
}

func (fake *FakeAgent) FailTTLArgsForCall(i int) (string, string) {
	fake.failTTLMutex.RLock()
	defer fake.failTTLMutex.RUnlock()
	return fake.failTTLArgsForCall[i].checkID, fake.failTTLArgsForCall[i].note
}

func (fake *FakeAgent) FailTTLReturns(result1 error) {
	fake.FailTTLStub = nil
	fake.failTTLReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgent) NodeName() (string, error) {
	fake.nodeNameMutex.Lock()
	fake.nodeNameArgsForCall = append(fake.nodeNameArgsForCall, struct{}{})
	fake.nodeNameMutex.Unlock()
	if fake.NodeNameStub != nil {
		return fake.NodeNameStub()
	} else {
		return fake.nodeNameReturns.result1, fake.nodeNameReturns.result2
	}
}

func (fake *FakeAgent) NodeNameCallCount() int {
	fake.nodeNameMutex.RLock()
	defer fake.nodeNameMutex.RUnlock()
	return len(fake.nodeNameArgsForCall)
}

func (fake *FakeAgent) NodeNameReturns(result1 string, result2 error) {
	fake.NodeNameStub = nil
	fake.nodeNameReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAgent) CheckDeregister(checkID string) error {
	fake.checkDeregisterMutex.Lock()
	fake.checkDeregisterArgsForCall = append(fake.checkDeregisterArgsForCall, struct {
		checkID string
	}{checkID})
	fake.checkDeregisterMutex.Unlock()
	if fake.CheckDeregisterStub != nil {
		return fake.CheckDeregisterStub(checkID)
	} else {
		return fake.checkDeregisterReturns.result1
	}
}

func (fake *FakeAgent) CheckDeregisterCallCount() int {
	fake.checkDeregisterMutex.RLock()
	defer fake.checkDeregisterMutex.RUnlock()
	return len(fake.checkDeregisterArgsForCall)
}

func (fake *FakeAgent) CheckDeregisterArgsForCall(i int) string {
	fake.checkDeregisterMutex.RLock()
	defer fake.checkDeregisterMutex.RUnlock()
	return fake.checkDeregisterArgsForCall[i].checkID
}

func (fake *FakeAgent) CheckDeregisterReturns(result1 error) {
	fake.CheckDeregisterStub = nil
	fake.checkDeregisterReturns = struct {
		result1 error
	}{result1}
}

var _ consuladapter.Agent = new(FakeAgent)
