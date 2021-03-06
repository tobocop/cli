// This file was generated by counterfeiter
package fakes

import (
	"sync"

	. "github.com/cloudfoundry/cli/cf/actors"
	"github.com/cloudfoundry/cli/cf/models"
)

type FakeServiceActor struct {
	FilterBrokersStub        func(brokerFlag string, serviceFlag string, orgFlag string) ([]models.ServiceBroker, error)
	filterBrokersMutex       sync.RWMutex
	filterBrokersArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	filterBrokersReturns struct {
		result1 []models.ServiceBroker
		result2 error
	}
	AttachPlansToServiceStub        func(models.ServiceOffering) (models.ServiceOffering, error)
	attachPlansToServiceMutex       sync.RWMutex
	attachPlansToServiceArgsForCall []struct {
		arg1 models.ServiceOffering
	}
	attachPlansToServiceReturns struct {
		result1 models.ServiceOffering
		result2 error
	}
	AttachOrgsToPlansStub        func([]models.ServicePlanFields) ([]models.ServicePlanFields, error)
	attachOrgsToPlansMutex       sync.RWMutex
	attachOrgsToPlansArgsForCall []struct {
		arg1 []models.ServicePlanFields
	}
	attachOrgsToPlansReturns struct {
		result1 []models.ServicePlanFields
		result2 error
	}
}

func (fake *FakeServiceActor) FilterBrokers(arg1 string, arg2 string, arg3 string) ([]models.ServiceBroker, error) {
	fake.filterBrokersMutex.Lock()
	defer fake.filterBrokersMutex.Unlock()
	fake.filterBrokersArgsForCall = append(fake.filterBrokersArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	if fake.FilterBrokersStub != nil {
		return fake.FilterBrokersStub(arg1, arg2, arg3)
	} else {
		return fake.filterBrokersReturns.result1, fake.filterBrokersReturns.result2
	}
}

func (fake *FakeServiceActor) FilterBrokersCallCount() int {
	fake.filterBrokersMutex.RLock()
	defer fake.filterBrokersMutex.RUnlock()
	return len(fake.filterBrokersArgsForCall)
}

func (fake *FakeServiceActor) FilterBrokersArgsForCall(i int) (string, string, string) {
	fake.filterBrokersMutex.RLock()
	defer fake.filterBrokersMutex.RUnlock()
	return fake.filterBrokersArgsForCall[i].arg1, fake.filterBrokersArgsForCall[i].arg2, fake.filterBrokersArgsForCall[i].arg3
}

func (fake *FakeServiceActor) FilterBrokersReturns(result1 []models.ServiceBroker, result2 error) {
	fake.filterBrokersReturns = struct {
		result1 []models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceActor) AttachPlansToService(arg1 models.ServiceOffering) (models.ServiceOffering, error) {
	fake.attachPlansToServiceMutex.Lock()
	defer fake.attachPlansToServiceMutex.Unlock()
	fake.attachPlansToServiceArgsForCall = append(fake.attachPlansToServiceArgsForCall, struct {
		arg1 models.ServiceOffering
	}{arg1})
	if fake.AttachPlansToServiceStub != nil {
		return fake.AttachPlansToServiceStub(arg1)
	} else {
		return fake.attachPlansToServiceReturns.result1, fake.attachPlansToServiceReturns.result2
	}
}

func (fake *FakeServiceActor) AttachPlansToServiceCallCount() int {
	fake.attachPlansToServiceMutex.RLock()
	defer fake.attachPlansToServiceMutex.RUnlock()
	return len(fake.attachPlansToServiceArgsForCall)
}

func (fake *FakeServiceActor) AttachPlansToServiceArgsForCall(i int) models.ServiceOffering {
	fake.attachPlansToServiceMutex.RLock()
	defer fake.attachPlansToServiceMutex.RUnlock()
	return fake.attachPlansToServiceArgsForCall[i].arg1
}

func (fake *FakeServiceActor) AttachPlansToServiceReturns(result1 models.ServiceOffering, result2 error) {
	fake.attachPlansToServiceReturns = struct {
		result1 models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceActor) AttachOrgsToPlans(arg1 []models.ServicePlanFields) ([]models.ServicePlanFields, error) {
	fake.attachOrgsToPlansMutex.Lock()
	defer fake.attachOrgsToPlansMutex.Unlock()
	fake.attachOrgsToPlansArgsForCall = append(fake.attachOrgsToPlansArgsForCall, struct {
		arg1 []models.ServicePlanFields
	}{arg1})
	if fake.AttachOrgsToPlansStub != nil {
		return fake.AttachOrgsToPlansStub(arg1)
	} else {
		return fake.attachOrgsToPlansReturns.result1, fake.attachOrgsToPlansReturns.result2
	}
}

func (fake *FakeServiceActor) AttachOrgsToPlansCallCount() int {
	fake.attachOrgsToPlansMutex.RLock()
	defer fake.attachOrgsToPlansMutex.RUnlock()
	return len(fake.attachOrgsToPlansArgsForCall)
}

func (fake *FakeServiceActor) AttachOrgsToPlansArgsForCall(i int) []models.ServicePlanFields {
	fake.attachOrgsToPlansMutex.RLock()
	defer fake.attachOrgsToPlansMutex.RUnlock()
	return fake.attachOrgsToPlansArgsForCall[i].arg1
}

func (fake *FakeServiceActor) AttachOrgsToPlansReturns(result1 []models.ServicePlanFields, result2 error) {
	fake.attachOrgsToPlansReturns = struct {
		result1 []models.ServicePlanFields
		result2 error
	}{result1, result2}
}

var _ ServiceActor = new(FakeServiceActor)
