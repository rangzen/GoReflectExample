package service_test

import (
	"context"
	mocks "github.com/rangzen/GoReflectExample/mocks/repository"
	"github.com/rangzen/GoReflectExample/service"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func reflectHelper(t *testing.T, methodName string, dataA string, dataB string, dataC string) string {
	// Indicate to the test framework that any problem here
	// has to be reported one level up.
	t.Helper()

	// Mock initialization
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	r := new(mocks.Repository)
	r.On("DataA", ctx).Return(dataA).Once()
	r.On("DataB", ctx).Return(dataB).Once()
	r.On("DataC", ctx).Return(dataC).Once()
	s := service.Service{Repository: r}

	// Get the reflection interface of the target service
	serviceReflection := reflect.ValueOf(&s)
	// Get the target method by name (string)
	methodReflection := serviceReflection.MethodByName(methodName)
	// Prepare an array of inputs for the target method
	in := []reflect.Value{reflect.ValueOf(ctx)}
	// Execute the method
	values := methodReflection.Call(in)
	// Extract return values
	got := values[0].String()
	return got
}

func TestService_EtlOrder(t *testing.T) {
	dataA := "DataA"
	dataB := "DataB"
	dataC := "DataC"

	got := reflectHelper(t, "EtlOrder", dataA, dataB, dataC)

	want := "DataADataBDataC"
	assert.Equal(t, got, want)
}

func TestService_EtlReverse(t *testing.T) {
	dataA := "DataA"
	dataB := "DataB"
	dataC := "DataC"

	got := reflectHelper(t, "EtlReverse", dataA, dataB, dataC)

	want := "DataCDataBDataA"
	assert.Equal(t, got, want)
}

/*

// Before refactoring

func TestService_EtlOrder(t *testing.T) {
	dataA := "DataA"
	dataB := "DataB"
	dataC := "DataC"
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Minute)
	r := new(mocks.Repository)
	r.On("DataA", ctx).Return(dataA).Once()
	r.On("DataB", ctx).Return(dataB).Once()
	r.On("DataC", ctx).Return(dataC).Once()
	s := service.Service{Repository: r}

	got := s.EtlOrder(ctx)

	want := "DataADataBDataC"
	assert.Equal(t, got, want)
}

func TestService_EtlReverse(t *testing.T) {
	dataA := "DataA"
	dataB := "DataB"
	dataC := "DataC"
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Minute)
	r := new(mocks.Repository)
	r.On("DataA", ctx).Return(dataA).Once()
	r.On("DataB", ctx).Return(dataB).Once()
	r.On("DataC", ctx).Return(dataC).Once()
	s := service.Service{Repository: r}

	got := s.EtlReverse(ctx)

	want := "DataCDataBDataA"
	assert.Equal(t, got, want)
}

*/
