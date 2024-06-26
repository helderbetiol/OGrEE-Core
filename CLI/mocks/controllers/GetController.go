// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// GetController is an autogenerated mock type for the GetController type
type GetController struct {
	mock.Mock
}

// GetObject provides a mock function with given fields: path
func (_m *GetController) GetObject(path string) (map[string]interface{}, error) {
	ret := _m.Called(path)

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (map[string]interface{}, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) map[string]interface{}); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObjectWithChildren provides a mock function with given fields: path, depth
func (_m *GetController) GetObjectWithChildren(path string, depth int) (map[string]interface{}, error) {
	ret := _m.Called(path, depth)

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int) (map[string]interface{}, error)); ok {
		return rf(path, depth)
	}
	if rf, ok := ret.Get(0).(func(string, int) map[string]interface{}); ok {
		r0 = rf(path, depth)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(path, depth)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObjectsWildcard provides a mock function with given fields: path
func (_m *GetController) GetObjectsWildcard(path string) ([]map[string]interface{}, []string, error) {
	ret := _m.Called(path)

	var r0 []map[string]interface{}
	var r1 []string
	var r2 error
	if rf, ok := ret.Get(0).(func(string) ([]map[string]interface{}, []string, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) []map[string]interface{}); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) []string); ok {
		r1 = rf(path)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(path)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PollObjectWithChildren provides a mock function with given fields: path, depth
func (_m *GetController) PollObjectWithChildren(path string, depth int) (map[string]interface{}, error) {
	ret := _m.Called(path, depth)

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int) (map[string]interface{}, error)); ok {
		return rf(path, depth)
	}
	if rf, ok := ret.Get(0).(func(string, int) map[string]interface{}); ok {
		r0 = rf(path, depth)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(path, depth)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGetController interface {
	mock.TestingT
	Cleanup(func())
}

// NewGetController creates a new instance of GetController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGetController(t mockConstructorTestingTNewGetController) *GetController {
	mock := &GetController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
