// Code generated by mockery v2.42.2. DO NOT EDIT.

package auth

import (
	fiber "github.com/gofiber/fiber/v2"
	mock "github.com/stretchr/testify/mock"
)

// MockController is an autogenerated mock type for the Controller type
type MockController struct {
	mock.Mock
}

type MockController_Expecter struct {
	mock *mock.Mock
}

func (_m *MockController) EXPECT() *MockController_Expecter {
	return &MockController_Expecter{mock: &_m.Mock}
}

// AuthMiddleware provides a mock function with given fields: c
func (_m *MockController) AuthMiddleware(c *fiber.Ctx) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for AuthMiddleware")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockController_AuthMiddleware_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthMiddleware'
type MockController_AuthMiddleware_Call struct {
	*mock.Call
}

// AuthMiddleware is a helper method to define mock.On call
//   - c *fiber.Ctx
func (_e *MockController_Expecter) AuthMiddleware(c interface{}) *MockController_AuthMiddleware_Call {
	return &MockController_AuthMiddleware_Call{Call: _e.mock.On("AuthMiddleware", c)}
}

func (_c *MockController_AuthMiddleware_Call) Run(run func(c *fiber.Ctx)) *MockController_AuthMiddleware_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*fiber.Ctx))
	})
	return _c
}

func (_c *MockController_AuthMiddleware_Call) Return(_a0 error) *MockController_AuthMiddleware_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockController_AuthMiddleware_Call) RunAndReturn(run func(*fiber.Ctx) error) *MockController_AuthMiddleware_Call {
	_c.Call.Return(run)
	return _c
}

// Callback provides a mock function with given fields: c
func (_m *MockController) Callback(c *fiber.Ctx) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Callback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockController_Callback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Callback'
type MockController_Callback_Call struct {
	*mock.Call
}

// Callback is a helper method to define mock.On call
//   - c *fiber.Ctx
func (_e *MockController_Expecter) Callback(c interface{}) *MockController_Callback_Call {
	return &MockController_Callback_Call{Call: _e.mock.On("Callback", c)}
}

func (_c *MockController_Callback_Call) Run(run func(c *fiber.Ctx)) *MockController_Callback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*fiber.Ctx))
	})
	return _c
}

func (_c *MockController_Callback_Call) Return(_a0 error) *MockController_Callback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockController_Callback_Call) RunAndReturn(run func(*fiber.Ctx) error) *MockController_Callback_Call {
	_c.Call.Return(run)
	return _c
}

// GetAuth provides a mock function with given fields: c
func (_m *MockController) GetAuth(c *fiber.Ctx) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for GetAuth")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockController_GetAuth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuth'
type MockController_GetAuth_Call struct {
	*mock.Call
}

// GetAuth is a helper method to define mock.On call
//   - c *fiber.Ctx
func (_e *MockController_Expecter) GetAuth(c interface{}) *MockController_GetAuth_Call {
	return &MockController_GetAuth_Call{Call: _e.mock.On("GetAuth", c)}
}

func (_c *MockController_GetAuth_Call) Run(run func(c *fiber.Ctx)) *MockController_GetAuth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*fiber.Ctx))
	})
	return _c
}

func (_c *MockController_GetAuth_Call) Return(_a0 error) *MockController_GetAuth_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockController_GetAuth_Call) RunAndReturn(run func(*fiber.Ctx) error) *MockController_GetAuth_Call {
	_c.Call.Return(run)
	return _c
}

// GetCsrf provides a mock function with given fields: c
func (_m *MockController) GetCsrf(c *fiber.Ctx) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for GetCsrf")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockController_GetCsrf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCsrf'
type MockController_GetCsrf_Call struct {
	*mock.Call
}

// GetCsrf is a helper method to define mock.On call
//   - c *fiber.Ctx
func (_e *MockController_Expecter) GetCsrf(c interface{}) *MockController_GetCsrf_Call {
	return &MockController_GetCsrf_Call{Call: _e.mock.On("GetCsrf", c)}
}

func (_c *MockController_GetCsrf_Call) Run(run func(c *fiber.Ctx)) *MockController_GetCsrf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*fiber.Ctx))
	})
	return _c
}

func (_c *MockController_GetCsrf_Call) Return(_a0 error) *MockController_GetCsrf_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockController_GetCsrf_Call) RunAndReturn(run func(*fiber.Ctx) error) *MockController_GetCsrf_Call {
	_c.Call.Return(run)
	return _c
}

// SignIn provides a mock function with given fields: c
func (_m *MockController) SignIn(c *fiber.Ctx) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for SignIn")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockController_SignIn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SignIn'
type MockController_SignIn_Call struct {
	*mock.Call
}

// SignIn is a helper method to define mock.On call
//   - c *fiber.Ctx
func (_e *MockController_Expecter) SignIn(c interface{}) *MockController_SignIn_Call {
	return &MockController_SignIn_Call{Call: _e.mock.On("SignIn", c)}
}

func (_c *MockController_SignIn_Call) Run(run func(c *fiber.Ctx)) *MockController_SignIn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*fiber.Ctx))
	})
	return _c
}

func (_c *MockController_SignIn_Call) Return(_a0 error) *MockController_SignIn_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockController_SignIn_Call) RunAndReturn(run func(*fiber.Ctx) error) *MockController_SignIn_Call {
	_c.Call.Return(run)
	return _c
}

// SignOut provides a mock function with given fields: c
func (_m *MockController) SignOut(c *fiber.Ctx) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for SignOut")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockController_SignOut_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SignOut'
type MockController_SignOut_Call struct {
	*mock.Call
}

// SignOut is a helper method to define mock.On call
//   - c *fiber.Ctx
func (_e *MockController_Expecter) SignOut(c interface{}) *MockController_SignOut_Call {
	return &MockController_SignOut_Call{Call: _e.mock.On("SignOut", c)}
}

func (_c *MockController_SignOut_Call) Run(run func(c *fiber.Ctx)) *MockController_SignOut_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*fiber.Ctx))
	})
	return _c
}

func (_c *MockController_SignOut_Call) Return(_a0 error) *MockController_SignOut_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockController_SignOut_Call) RunAndReturn(run func(*fiber.Ctx) error) *MockController_SignOut_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockController creates a new instance of MockController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockController(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockController {
	mock := &MockController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
