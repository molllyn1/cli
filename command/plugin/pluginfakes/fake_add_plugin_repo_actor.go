// Code generated by counterfeiter. DO NOT EDIT.
package pluginfakes

import (
	sync "sync"

	plugin "code.cloudfoundry.org/cli/command/plugin"
)

type FakeAddPluginRepoActor struct {
	AddPluginRepositoryStub        func(string, string) error
	addPluginRepositoryMutex       sync.RWMutex
	addPluginRepositoryArgsForCall []struct {
		arg1 string
		arg2 string
	}
	addPluginRepositoryReturns struct {
		result1 error
	}
	addPluginRepositoryReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAddPluginRepoActor) AddPluginRepository(arg1 string, arg2 string) error {
	fake.addPluginRepositoryMutex.Lock()
	ret, specificReturn := fake.addPluginRepositoryReturnsOnCall[len(fake.addPluginRepositoryArgsForCall)]
	fake.addPluginRepositoryArgsForCall = append(fake.addPluginRepositoryArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("AddPluginRepository", []interface{}{arg1, arg2})
	fake.addPluginRepositoryMutex.Unlock()
	if fake.AddPluginRepositoryStub != nil {
		return fake.AddPluginRepositoryStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addPluginRepositoryReturns
	return fakeReturns.result1
}

func (fake *FakeAddPluginRepoActor) AddPluginRepositoryCallCount() int {
	fake.addPluginRepositoryMutex.RLock()
	defer fake.addPluginRepositoryMutex.RUnlock()
	return len(fake.addPluginRepositoryArgsForCall)
}

func (fake *FakeAddPluginRepoActor) AddPluginRepositoryCalls(stub func(string, string) error) {
	fake.addPluginRepositoryMutex.Lock()
	defer fake.addPluginRepositoryMutex.Unlock()
	fake.AddPluginRepositoryStub = stub
}

func (fake *FakeAddPluginRepoActor) AddPluginRepositoryArgsForCall(i int) (string, string) {
	fake.addPluginRepositoryMutex.RLock()
	defer fake.addPluginRepositoryMutex.RUnlock()
	argsForCall := fake.addPluginRepositoryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAddPluginRepoActor) AddPluginRepositoryReturns(result1 error) {
	fake.addPluginRepositoryMutex.Lock()
	defer fake.addPluginRepositoryMutex.Unlock()
	fake.AddPluginRepositoryStub = nil
	fake.addPluginRepositoryReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAddPluginRepoActor) AddPluginRepositoryReturnsOnCall(i int, result1 error) {
	fake.addPluginRepositoryMutex.Lock()
	defer fake.addPluginRepositoryMutex.Unlock()
	fake.AddPluginRepositoryStub = nil
	if fake.addPluginRepositoryReturnsOnCall == nil {
		fake.addPluginRepositoryReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addPluginRepositoryReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAddPluginRepoActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addPluginRepositoryMutex.RLock()
	defer fake.addPluginRepositoryMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAddPluginRepoActor) recordInvocation(key string, args []interface{}) {
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

var _ plugin.AddPluginRepoActor = new(FakeAddPluginRepoActor)