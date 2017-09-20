// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/bbs/models"
)

type FakeValidator struct {
	ValidateStub        func() error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct{}
	validateReturns     struct {
		result1 error
	}
}

func (fake *FakeValidator) Validate() error {
	fake.validateMutex.Lock()
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct{}{})
	fake.validateMutex.Unlock()
	if fake.ValidateStub != nil {
		return fake.ValidateStub()
	} else {
		return fake.validateReturns.result1
	}
}

func (fake *FakeValidator) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *FakeValidator) ValidateReturns(result1 error) {
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 error
	}{result1}
}

var _ models.Validator = new(FakeValidator)
