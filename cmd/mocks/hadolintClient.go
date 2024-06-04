// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	pkghttp "github.com/SAP/jenkins-library/pkg/http"
	mock "github.com/stretchr/testify/mock"
)

// HadolintClient is an autogenerated mock type for the HadolintClient type
type HadolintClient struct {
	mock.Mock
}

type HadolintClient_Expecter struct {
	mock *mock.Mock
}

func (_m *HadolintClient) EXPECT() *HadolintClient_Expecter {
	return &HadolintClient_Expecter{mock: &_m.Mock}
}

// DownloadFile provides a mock function with given fields: url, filename, header, cookies
func (_m *HadolintClient) DownloadFile(url string, filename string, header http.Header, cookies []*http.Cookie) error {
	ret := _m.Called(url, filename, header, cookies)

	if len(ret) == 0 {
		panic("no return value specified for DownloadFile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, http.Header, []*http.Cookie) error); ok {
		r0 = rf(url, filename, header, cookies)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HadolintClient_DownloadFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadFile'
type HadolintClient_DownloadFile_Call struct {
	*mock.Call
}

// DownloadFile is a helper method to define mock.On call
//   - url string
//   - filename string
//   - header http.Header
//   - cookies []*http.Cookie
func (_e *HadolintClient_Expecter) DownloadFile(url interface{}, filename interface{}, header interface{}, cookies interface{}) *HadolintClient_DownloadFile_Call {
	return &HadolintClient_DownloadFile_Call{Call: _e.mock.On("DownloadFile", url, filename, header, cookies)}
}

func (_c *HadolintClient_DownloadFile_Call) Run(run func(url string, filename string, header http.Header, cookies []*http.Cookie)) *HadolintClient_DownloadFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(http.Header), args[3].([]*http.Cookie))
	})
	return _c
}

func (_c *HadolintClient_DownloadFile_Call) Return(_a0 error) *HadolintClient_DownloadFile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HadolintClient_DownloadFile_Call) RunAndReturn(run func(string, string, http.Header, []*http.Cookie) error) *HadolintClient_DownloadFile_Call {
	_c.Call.Return(run)
	return _c
}

// SetOptions provides a mock function with given fields: options
func (_m *HadolintClient) SetOptions(options pkghttp.ClientOptions) {
	_m.Called(options)
}

// HadolintClient_SetOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetOptions'
type HadolintClient_SetOptions_Call struct {
	*mock.Call
}

// SetOptions is a helper method to define mock.On call
//   - options pkghttp.ClientOptions
func (_e *HadolintClient_Expecter) SetOptions(options interface{}) *HadolintClient_SetOptions_Call {
	return &HadolintClient_SetOptions_Call{Call: _e.mock.On("SetOptions", options)}
}

func (_c *HadolintClient_SetOptions_Call) Run(run func(options pkghttp.ClientOptions)) *HadolintClient_SetOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(pkghttp.ClientOptions))
	})
	return _c
}

func (_c *HadolintClient_SetOptions_Call) Return() *HadolintClient_SetOptions_Call {
	_c.Call.Return()
	return _c
}

func (_c *HadolintClient_SetOptions_Call) RunAndReturn(run func(pkghttp.ClientOptions)) *HadolintClient_SetOptions_Call {
	_c.Call.Return(run)
	return _c
}

// NewHadolintClient creates a new instance of HadolintClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHadolintClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *HadolintClient {
	mock := &HadolintClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}