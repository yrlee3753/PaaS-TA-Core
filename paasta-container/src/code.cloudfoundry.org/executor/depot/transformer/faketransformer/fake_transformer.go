// Code generated by counterfeiter. DO NOT EDIT.
package faketransformer

import (
	"sync"

	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/executor"
	"code.cloudfoundry.org/executor/depot/log_streamer"
	"code.cloudfoundry.org/executor/depot/steps"
	"code.cloudfoundry.org/executor/depot/transformer"
	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/lager"
	"github.com/tedsuo/ifrit"
)

type FakeTransformer struct {
	StepForStub        func(log_streamer.LogStreamer, *models.Action, garden.Container, string, string, []executor.PortMapping, bool, bool, lager.Logger) steps.Step
	stepForMutex       sync.RWMutex
	stepForArgsForCall []struct {
		arg1 log_streamer.LogStreamer
		arg2 *models.Action
		arg3 garden.Container
		arg4 string
		arg5 string
		arg6 []executor.PortMapping
		arg7 bool
		arg8 bool
		arg9 lager.Logger
	}
	stepForReturns struct {
		result1 steps.Step
	}
	stepForReturnsOnCall map[int]struct {
		result1 steps.Step
	}
	StepsRunnerStub        func(lager.Logger, executor.Container, garden.Container, log_streamer.LogStreamer) (ifrit.Runner, error)
	stepsRunnerMutex       sync.RWMutex
	stepsRunnerArgsForCall []struct {
		arg1 lager.Logger
		arg2 executor.Container
		arg3 garden.Container
		arg4 log_streamer.LogStreamer
	}
	stepsRunnerReturns struct {
		result1 ifrit.Runner
		result2 error
	}
	stepsRunnerReturnsOnCall map[int]struct {
		result1 ifrit.Runner
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTransformer) StepFor(arg1 log_streamer.LogStreamer, arg2 *models.Action, arg3 garden.Container, arg4 string, arg5 string, arg6 []executor.PortMapping, arg7 bool, arg8 bool, arg9 lager.Logger) steps.Step {
	var arg6Copy []executor.PortMapping
	if arg6 != nil {
		arg6Copy = make([]executor.PortMapping, len(arg6))
		copy(arg6Copy, arg6)
	}
	fake.stepForMutex.Lock()
	ret, specificReturn := fake.stepForReturnsOnCall[len(fake.stepForArgsForCall)]
	fake.stepForArgsForCall = append(fake.stepForArgsForCall, struct {
		arg1 log_streamer.LogStreamer
		arg2 *models.Action
		arg3 garden.Container
		arg4 string
		arg5 string
		arg6 []executor.PortMapping
		arg7 bool
		arg8 bool
		arg9 lager.Logger
	}{arg1, arg2, arg3, arg4, arg5, arg6Copy, arg7, arg8, arg9})
	fake.recordInvocation("StepFor", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6Copy, arg7, arg8, arg9})
	fake.stepForMutex.Unlock()
	if fake.StepForStub != nil {
		return fake.StepForStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stepForReturns.result1
}

func (fake *FakeTransformer) StepForCallCount() int {
	fake.stepForMutex.RLock()
	defer fake.stepForMutex.RUnlock()
	return len(fake.stepForArgsForCall)
}

func (fake *FakeTransformer) StepForArgsForCall(i int) (log_streamer.LogStreamer, *models.Action, garden.Container, string, string, []executor.PortMapping, bool, bool, lager.Logger) {
	fake.stepForMutex.RLock()
	defer fake.stepForMutex.RUnlock()
	return fake.stepForArgsForCall[i].arg1, fake.stepForArgsForCall[i].arg2, fake.stepForArgsForCall[i].arg3, fake.stepForArgsForCall[i].arg4, fake.stepForArgsForCall[i].arg5, fake.stepForArgsForCall[i].arg6, fake.stepForArgsForCall[i].arg7, fake.stepForArgsForCall[i].arg8, fake.stepForArgsForCall[i].arg9
}

func (fake *FakeTransformer) StepForReturns(result1 steps.Step) {
	fake.StepForStub = nil
	fake.stepForReturns = struct {
		result1 steps.Step
	}{result1}
}

func (fake *FakeTransformer) StepForReturnsOnCall(i int, result1 steps.Step) {
	fake.StepForStub = nil
	if fake.stepForReturnsOnCall == nil {
		fake.stepForReturnsOnCall = make(map[int]struct {
			result1 steps.Step
		})
	}
	fake.stepForReturnsOnCall[i] = struct {
		result1 steps.Step
	}{result1}
}

func (fake *FakeTransformer) StepsRunner(arg1 lager.Logger, arg2 executor.Container, arg3 garden.Container, arg4 log_streamer.LogStreamer) (ifrit.Runner, error) {
	fake.stepsRunnerMutex.Lock()
	ret, specificReturn := fake.stepsRunnerReturnsOnCall[len(fake.stepsRunnerArgsForCall)]
	fake.stepsRunnerArgsForCall = append(fake.stepsRunnerArgsForCall, struct {
		arg1 lager.Logger
		arg2 executor.Container
		arg3 garden.Container
		arg4 log_streamer.LogStreamer
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("StepsRunner", []interface{}{arg1, arg2, arg3, arg4})
	fake.stepsRunnerMutex.Unlock()
	if fake.StepsRunnerStub != nil {
		return fake.StepsRunnerStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.stepsRunnerReturns.result1, fake.stepsRunnerReturns.result2
}

func (fake *FakeTransformer) StepsRunnerCallCount() int {
	fake.stepsRunnerMutex.RLock()
	defer fake.stepsRunnerMutex.RUnlock()
	return len(fake.stepsRunnerArgsForCall)
}

func (fake *FakeTransformer) StepsRunnerArgsForCall(i int) (lager.Logger, executor.Container, garden.Container, log_streamer.LogStreamer) {
	fake.stepsRunnerMutex.RLock()
	defer fake.stepsRunnerMutex.RUnlock()
	return fake.stepsRunnerArgsForCall[i].arg1, fake.stepsRunnerArgsForCall[i].arg2, fake.stepsRunnerArgsForCall[i].arg3, fake.stepsRunnerArgsForCall[i].arg4
}

func (fake *FakeTransformer) StepsRunnerReturns(result1 ifrit.Runner, result2 error) {
	fake.StepsRunnerStub = nil
	fake.stepsRunnerReturns = struct {
		result1 ifrit.Runner
		result2 error
	}{result1, result2}
}

func (fake *FakeTransformer) StepsRunnerReturnsOnCall(i int, result1 ifrit.Runner, result2 error) {
	fake.StepsRunnerStub = nil
	if fake.stepsRunnerReturnsOnCall == nil {
		fake.stepsRunnerReturnsOnCall = make(map[int]struct {
			result1 ifrit.Runner
			result2 error
		})
	}
	fake.stepsRunnerReturnsOnCall[i] = struct {
		result1 ifrit.Runner
		result2 error
	}{result1, result2}
}

func (fake *FakeTransformer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.stepForMutex.RLock()
	defer fake.stepForMutex.RUnlock()
	fake.stepsRunnerMutex.RLock()
	defer fake.stepsRunnerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTransformer) recordInvocation(key string, args []interface{}) {
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

var _ transformer.Transformer = new(FakeTransformer)
