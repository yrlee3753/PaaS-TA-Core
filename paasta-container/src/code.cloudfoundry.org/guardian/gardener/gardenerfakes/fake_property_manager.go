// Code generated by counterfeiter. DO NOT EDIT.
package gardenerfakes

import (
	"sync"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/gardener"
)

type FakePropertyManager struct {
	AllStub        func(handle string) (props garden.Properties, err error)
	allMutex       sync.RWMutex
	allArgsForCall []struct {
		handle string
	}
	allReturns struct {
		result1 garden.Properties
		result2 error
	}
	allReturnsOnCall map[int]struct {
		result1 garden.Properties
		result2 error
	}
	SetStub        func(handle string, name string, value string)
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		handle string
		name   string
		value  string
	}
	RemoveStub        func(handle string, name string) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		handle string
		name   string
	}
	removeReturns struct {
		result1 error
	}
	removeReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(handle string, name string) (string, bool)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		handle string
		name   string
	}
	getReturns struct {
		result1 string
		result2 bool
	}
	getReturnsOnCall map[int]struct {
		result1 string
		result2 bool
	}
	MatchesAllStub        func(handle string, props garden.Properties) bool
	matchesAllMutex       sync.RWMutex
	matchesAllArgsForCall []struct {
		handle string
		props  garden.Properties
	}
	matchesAllReturns struct {
		result1 bool
	}
	matchesAllReturnsOnCall map[int]struct {
		result1 bool
	}
	DestroyKeySpaceStub        func(string) error
	destroyKeySpaceMutex       sync.RWMutex
	destroyKeySpaceArgsForCall []struct {
		arg1 string
	}
	destroyKeySpaceReturns struct {
		result1 error
	}
	destroyKeySpaceReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePropertyManager) All(handle string) (props garden.Properties, err error) {
	fake.allMutex.Lock()
	ret, specificReturn := fake.allReturnsOnCall[len(fake.allArgsForCall)]
	fake.allArgsForCall = append(fake.allArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("All", []interface{}{handle})
	fake.allMutex.Unlock()
	if fake.AllStub != nil {
		return fake.AllStub(handle)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.allReturns.result1, fake.allReturns.result2
}

func (fake *FakePropertyManager) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *FakePropertyManager) AllArgsForCall(i int) string {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return fake.allArgsForCall[i].handle
}

func (fake *FakePropertyManager) AllReturns(result1 garden.Properties, result2 error) {
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 garden.Properties
		result2 error
	}{result1, result2}
}

func (fake *FakePropertyManager) AllReturnsOnCall(i int, result1 garden.Properties, result2 error) {
	fake.AllStub = nil
	if fake.allReturnsOnCall == nil {
		fake.allReturnsOnCall = make(map[int]struct {
			result1 garden.Properties
			result2 error
		})
	}
	fake.allReturnsOnCall[i] = struct {
		result1 garden.Properties
		result2 error
	}{result1, result2}
}

func (fake *FakePropertyManager) Set(handle string, name string, value string) {
	fake.setMutex.Lock()
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		handle string
		name   string
		value  string
	}{handle, name, value})
	fake.recordInvocation("Set", []interface{}{handle, name, value})
	fake.setMutex.Unlock()
	if fake.SetStub != nil {
		fake.SetStub(handle, name, value)
	}
}

func (fake *FakePropertyManager) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *FakePropertyManager) SetArgsForCall(i int) (string, string, string) {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return fake.setArgsForCall[i].handle, fake.setArgsForCall[i].name, fake.setArgsForCall[i].value
}

func (fake *FakePropertyManager) Remove(handle string, name string) error {
	fake.removeMutex.Lock()
	ret, specificReturn := fake.removeReturnsOnCall[len(fake.removeArgsForCall)]
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		handle string
		name   string
	}{handle, name})
	fake.recordInvocation("Remove", []interface{}{handle, name})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(handle, name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeReturns.result1
}

func (fake *FakePropertyManager) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakePropertyManager) RemoveArgsForCall(i int) (string, string) {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return fake.removeArgsForCall[i].handle, fake.removeArgsForCall[i].name
}

func (fake *FakePropertyManager) RemoveReturns(result1 error) {
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePropertyManager) RemoveReturnsOnCall(i int, result1 error) {
	fake.RemoveStub = nil
	if fake.removeReturnsOnCall == nil {
		fake.removeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePropertyManager) Get(handle string, name string) (string, bool) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		handle string
		name   string
	}{handle, name})
	fake.recordInvocation("Get", []interface{}{handle, name})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(handle, name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getReturns.result1, fake.getReturns.result2
}

func (fake *FakePropertyManager) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakePropertyManager) GetArgsForCall(i int) (string, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].handle, fake.getArgsForCall[i].name
}

func (fake *FakePropertyManager) GetReturns(result1 string, result2 bool) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 string
		result2 bool
	}{result1, result2}
}

func (fake *FakePropertyManager) GetReturnsOnCall(i int, result1 string, result2 bool) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 string
			result2 bool
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 string
		result2 bool
	}{result1, result2}
}

func (fake *FakePropertyManager) MatchesAll(handle string, props garden.Properties) bool {
	fake.matchesAllMutex.Lock()
	ret, specificReturn := fake.matchesAllReturnsOnCall[len(fake.matchesAllArgsForCall)]
	fake.matchesAllArgsForCall = append(fake.matchesAllArgsForCall, struct {
		handle string
		props  garden.Properties
	}{handle, props})
	fake.recordInvocation("MatchesAll", []interface{}{handle, props})
	fake.matchesAllMutex.Unlock()
	if fake.MatchesAllStub != nil {
		return fake.MatchesAllStub(handle, props)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.matchesAllReturns.result1
}

func (fake *FakePropertyManager) MatchesAllCallCount() int {
	fake.matchesAllMutex.RLock()
	defer fake.matchesAllMutex.RUnlock()
	return len(fake.matchesAllArgsForCall)
}

func (fake *FakePropertyManager) MatchesAllArgsForCall(i int) (string, garden.Properties) {
	fake.matchesAllMutex.RLock()
	defer fake.matchesAllMutex.RUnlock()
	return fake.matchesAllArgsForCall[i].handle, fake.matchesAllArgsForCall[i].props
}

func (fake *FakePropertyManager) MatchesAllReturns(result1 bool) {
	fake.MatchesAllStub = nil
	fake.matchesAllReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakePropertyManager) MatchesAllReturnsOnCall(i int, result1 bool) {
	fake.MatchesAllStub = nil
	if fake.matchesAllReturnsOnCall == nil {
		fake.matchesAllReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.matchesAllReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakePropertyManager) DestroyKeySpace(arg1 string) error {
	fake.destroyKeySpaceMutex.Lock()
	ret, specificReturn := fake.destroyKeySpaceReturnsOnCall[len(fake.destroyKeySpaceArgsForCall)]
	fake.destroyKeySpaceArgsForCall = append(fake.destroyKeySpaceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DestroyKeySpace", []interface{}{arg1})
	fake.destroyKeySpaceMutex.Unlock()
	if fake.DestroyKeySpaceStub != nil {
		return fake.DestroyKeySpaceStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.destroyKeySpaceReturns.result1
}

func (fake *FakePropertyManager) DestroyKeySpaceCallCount() int {
	fake.destroyKeySpaceMutex.RLock()
	defer fake.destroyKeySpaceMutex.RUnlock()
	return len(fake.destroyKeySpaceArgsForCall)
}

func (fake *FakePropertyManager) DestroyKeySpaceArgsForCall(i int) string {
	fake.destroyKeySpaceMutex.RLock()
	defer fake.destroyKeySpaceMutex.RUnlock()
	return fake.destroyKeySpaceArgsForCall[i].arg1
}

func (fake *FakePropertyManager) DestroyKeySpaceReturns(result1 error) {
	fake.DestroyKeySpaceStub = nil
	fake.destroyKeySpaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePropertyManager) DestroyKeySpaceReturnsOnCall(i int, result1 error) {
	fake.DestroyKeySpaceStub = nil
	if fake.destroyKeySpaceReturnsOnCall == nil {
		fake.destroyKeySpaceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyKeySpaceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePropertyManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.matchesAllMutex.RLock()
	defer fake.matchesAllMutex.RUnlock()
	fake.destroyKeySpaceMutex.RLock()
	defer fake.destroyKeySpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePropertyManager) recordInvocation(key string, args []interface{}) {
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

var _ gardener.PropertyManager = new(FakePropertyManager)
