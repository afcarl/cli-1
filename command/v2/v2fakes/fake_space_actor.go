// Code generated by counterfeiter. DO NOT EDIT.
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/v2"
)

type FakeSpaceActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	GetSpaceByOrganizationAndNameStub        func(orgGUID string, spaceName string) (v2action.Space, v2action.Warnings, error)
	getSpaceByOrganizationAndNameMutex       sync.RWMutex
	getSpaceByOrganizationAndNameArgsForCall []struct {
		orgGUID   string
		spaceName string
	}
	getSpaceByOrganizationAndNameReturns struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	getSpaceByOrganizationAndNameReturnsOnCall map[int]struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	GetSpaceSummaryByOrganizationAndNameStub        func(orgGUID string, spaceName string) (v2action.SpaceSummary, v2action.Warnings, error)
	getSpaceSummaryByOrganizationAndNameMutex       sync.RWMutex
	getSpaceSummaryByOrganizationAndNameArgsForCall []struct {
		orgGUID   string
		spaceName string
	}
	getSpaceSummaryByOrganizationAndNameReturns struct {
		result1 v2action.SpaceSummary
		result2 v2action.Warnings
		result3 error
	}
	getSpaceSummaryByOrganizationAndNameReturnsOnCall map[int]struct {
		result1 v2action.SpaceSummary
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpaceActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloudControllerAPIVersionReturns.result1
}

func (fake *FakeSpaceActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeSpaceActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSpaceActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSpaceActor) GetSpaceByOrganizationAndName(orgGUID string, spaceName string) (v2action.Space, v2action.Warnings, error) {
	fake.getSpaceByOrganizationAndNameMutex.Lock()
	ret, specificReturn := fake.getSpaceByOrganizationAndNameReturnsOnCall[len(fake.getSpaceByOrganizationAndNameArgsForCall)]
	fake.getSpaceByOrganizationAndNameArgsForCall = append(fake.getSpaceByOrganizationAndNameArgsForCall, struct {
		orgGUID   string
		spaceName string
	}{orgGUID, spaceName})
	fake.recordInvocation("GetSpaceByOrganizationAndName", []interface{}{orgGUID, spaceName})
	fake.getSpaceByOrganizationAndNameMutex.Unlock()
	if fake.GetSpaceByOrganizationAndNameStub != nil {
		return fake.GetSpaceByOrganizationAndNameStub(orgGUID, spaceName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getSpaceByOrganizationAndNameReturns.result1, fake.getSpaceByOrganizationAndNameReturns.result2, fake.getSpaceByOrganizationAndNameReturns.result3
}

func (fake *FakeSpaceActor) GetSpaceByOrganizationAndNameCallCount() int {
	fake.getSpaceByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceByOrganizationAndNameMutex.RUnlock()
	return len(fake.getSpaceByOrganizationAndNameArgsForCall)
}

func (fake *FakeSpaceActor) GetSpaceByOrganizationAndNameArgsForCall(i int) (string, string) {
	fake.getSpaceByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceByOrganizationAndNameMutex.RUnlock()
	return fake.getSpaceByOrganizationAndNameArgsForCall[i].orgGUID, fake.getSpaceByOrganizationAndNameArgsForCall[i].spaceName
}

func (fake *FakeSpaceActor) GetSpaceByOrganizationAndNameReturns(result1 v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.GetSpaceByOrganizationAndNameStub = nil
	fake.getSpaceByOrganizationAndNameReturns = struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceActor) GetSpaceByOrganizationAndNameReturnsOnCall(i int, result1 v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.GetSpaceByOrganizationAndNameStub = nil
	if fake.getSpaceByOrganizationAndNameReturnsOnCall == nil {
		fake.getSpaceByOrganizationAndNameReturnsOnCall = make(map[int]struct {
			result1 v2action.Space
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getSpaceByOrganizationAndNameReturnsOnCall[i] = struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceActor) GetSpaceSummaryByOrganizationAndName(orgGUID string, spaceName string) (v2action.SpaceSummary, v2action.Warnings, error) {
	fake.getSpaceSummaryByOrganizationAndNameMutex.Lock()
	ret, specificReturn := fake.getSpaceSummaryByOrganizationAndNameReturnsOnCall[len(fake.getSpaceSummaryByOrganizationAndNameArgsForCall)]
	fake.getSpaceSummaryByOrganizationAndNameArgsForCall = append(fake.getSpaceSummaryByOrganizationAndNameArgsForCall, struct {
		orgGUID   string
		spaceName string
	}{orgGUID, spaceName})
	fake.recordInvocation("GetSpaceSummaryByOrganizationAndName", []interface{}{orgGUID, spaceName})
	fake.getSpaceSummaryByOrganizationAndNameMutex.Unlock()
	if fake.GetSpaceSummaryByOrganizationAndNameStub != nil {
		return fake.GetSpaceSummaryByOrganizationAndNameStub(orgGUID, spaceName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getSpaceSummaryByOrganizationAndNameReturns.result1, fake.getSpaceSummaryByOrganizationAndNameReturns.result2, fake.getSpaceSummaryByOrganizationAndNameReturns.result3
}

func (fake *FakeSpaceActor) GetSpaceSummaryByOrganizationAndNameCallCount() int {
	fake.getSpaceSummaryByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceSummaryByOrganizationAndNameMutex.RUnlock()
	return len(fake.getSpaceSummaryByOrganizationAndNameArgsForCall)
}

func (fake *FakeSpaceActor) GetSpaceSummaryByOrganizationAndNameArgsForCall(i int) (string, string) {
	fake.getSpaceSummaryByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceSummaryByOrganizationAndNameMutex.RUnlock()
	return fake.getSpaceSummaryByOrganizationAndNameArgsForCall[i].orgGUID, fake.getSpaceSummaryByOrganizationAndNameArgsForCall[i].spaceName
}

func (fake *FakeSpaceActor) GetSpaceSummaryByOrganizationAndNameReturns(result1 v2action.SpaceSummary, result2 v2action.Warnings, result3 error) {
	fake.GetSpaceSummaryByOrganizationAndNameStub = nil
	fake.getSpaceSummaryByOrganizationAndNameReturns = struct {
		result1 v2action.SpaceSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceActor) GetSpaceSummaryByOrganizationAndNameReturnsOnCall(i int, result1 v2action.SpaceSummary, result2 v2action.Warnings, result3 error) {
	fake.GetSpaceSummaryByOrganizationAndNameStub = nil
	if fake.getSpaceSummaryByOrganizationAndNameReturnsOnCall == nil {
		fake.getSpaceSummaryByOrganizationAndNameReturnsOnCall = make(map[int]struct {
			result1 v2action.SpaceSummary
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getSpaceSummaryByOrganizationAndNameReturnsOnCall[i] = struct {
		result1 v2action.SpaceSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.getSpaceByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceByOrganizationAndNameMutex.RUnlock()
	fake.getSpaceSummaryByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceSummaryByOrganizationAndNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSpaceActor) recordInvocation(key string, args []interface{}) {
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

var _ v2.SpaceActor = new(FakeSpaceActor)
