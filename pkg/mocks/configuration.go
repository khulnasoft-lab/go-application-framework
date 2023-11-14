// Code generated by MockGen. DO NOT EDIT.
// Source: configuration.go

// Package mocks is a generated GoMock package.
package mocks

import (
	url "net/url"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	configuration "github.com/khulnasoft-lab/go-application-framework/pkg/configuration"
	pflag "github.com/spf13/pflag"
)

// MockConfiguration is a mock of Configuration interface.
type MockConfiguration struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationMockRecorder
}

// MockConfigurationMockRecorder is the mock recorder for MockConfiguration.
type MockConfigurationMockRecorder struct {
	mock *MockConfiguration
}

// NewMockConfiguration creates a new mock instance.
func NewMockConfiguration(ctrl *gomock.Controller) *MockConfiguration {
	mock := &MockConfiguration{ctrl: ctrl}
	mock.recorder = &MockConfigurationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfiguration) EXPECT() *MockConfigurationMockRecorder {
	return m.recorder
}

// AddAlternativeKeys mocks base method.
func (m *MockConfiguration) AddAlternativeKeys(key string, altKeys []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddAlternativeKeys", key, altKeys)
}

// AddAlternativeKeys indicates an expected call of AddAlternativeKeys.
func (mr *MockConfigurationMockRecorder) AddAlternativeKeys(key, altKeys interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAlternativeKeys", reflect.TypeOf((*MockConfiguration)(nil).AddAlternativeKeys), key, altKeys)
}

// AddDefaultValue mocks base method.
func (m *MockConfiguration) AddDefaultValue(key string, defaultValue configuration.DefaultValueFunction) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddDefaultValue", key, defaultValue)
}

// AddDefaultValue indicates an expected call of AddDefaultValue.
func (mr *MockConfigurationMockRecorder) AddDefaultValue(key, defaultValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDefaultValue", reflect.TypeOf((*MockConfiguration)(nil).AddDefaultValue), key, defaultValue)
}

// AddFlagSet mocks base method.
func (m *MockConfiguration) AddFlagSet(flagset *pflag.FlagSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFlagSet", flagset)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFlagSet indicates an expected call of AddFlagSet.
func (mr *MockConfigurationMockRecorder) AddFlagSet(flagset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFlagSet", reflect.TypeOf((*MockConfiguration)(nil).AddFlagSet), flagset)
}

// AllKeys mocks base method.
func (m *MockConfiguration) AllKeys() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllKeys")
	ret0, _ := ret[0].([]string)
	return ret0
}

// AllKeys indicates an expected call of AllKeys.
func (mr *MockConfigurationMockRecorder) AllKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllKeys", reflect.TypeOf((*MockConfiguration)(nil).AllKeys))
}

// Clone mocks base method.
func (m *MockConfiguration) Clone() configuration.Configuration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(configuration.Configuration)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockConfigurationMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockConfiguration)(nil).Clone))
}

// Get mocks base method.
func (m *MockConfiguration) Get(key string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockConfigurationMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConfiguration)(nil).Get), key)
}

// GetAlternativeKeys mocks base method.
func (m *MockConfiguration) GetAlternativeKeys(key string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlternativeKeys", key)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetAlternativeKeys indicates an expected call of GetAlternativeKeys.
func (mr *MockConfigurationMockRecorder) GetAlternativeKeys(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlternativeKeys", reflect.TypeOf((*MockConfiguration)(nil).GetAlternativeKeys), key)
}

// GetAndIsSet mocks base method.
func (m *MockConfiguration) GetAndIsSet(key string) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAndIsSet", key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetAndIsSet indicates an expected call of GetAndIsSet.
func (mr *MockConfigurationMockRecorder) GetAndIsSet(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAndIsSet", reflect.TypeOf((*MockConfiguration)(nil).GetAndIsSet), key)
}

// GetBool mocks base method.
func (m *MockConfiguration) GetBool(key string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBool", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetBool indicates an expected call of GetBool.
func (mr *MockConfigurationMockRecorder) GetBool(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBool", reflect.TypeOf((*MockConfiguration)(nil).GetBool), key)
}

// GetFloat64 mocks base method.
func (m *MockConfiguration) GetFloat64(key string) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFloat64", key)
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetFloat64 indicates an expected call of GetFloat64.
func (mr *MockConfigurationMockRecorder) GetFloat64(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFloat64", reflect.TypeOf((*MockConfiguration)(nil).GetFloat64), key)
}

// GetInt mocks base method.
func (m *MockConfiguration) GetInt(key string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInt", key)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetInt indicates an expected call of GetInt.
func (mr *MockConfigurationMockRecorder) GetInt(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt", reflect.TypeOf((*MockConfiguration)(nil).GetInt), key)
}

// GetStorage mocks base method.
func (m *MockConfiguration) GetStorage() configuration.Storage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorage")
	ret0, _ := ret[0].(configuration.Storage)
	return ret0
}

// GetStorage indicates an expected call of GetStorage.
func (mr *MockConfigurationMockRecorder) GetStorage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorage", reflect.TypeOf((*MockConfiguration)(nil).GetStorage))
}

// GetString mocks base method.
func (m *MockConfiguration) GetString(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetString", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetString indicates an expected call of GetString.
func (mr *MockConfigurationMockRecorder) GetString(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetString", reflect.TypeOf((*MockConfiguration)(nil).GetString), key)
}

// GetStringSlice mocks base method.
func (m *MockConfiguration) GetStringSlice(key string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStringSlice", key)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetStringSlice indicates an expected call of GetStringSlice.
func (mr *MockConfigurationMockRecorder) GetStringSlice(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringSlice", reflect.TypeOf((*MockConfiguration)(nil).GetStringSlice), key)
}

// GetUrl mocks base method.
func (m *MockConfiguration) GetUrl(key string) *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUrl", key)
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// GetUrl indicates an expected call of GetUrl.
func (mr *MockConfigurationMockRecorder) GetUrl(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUrl", reflect.TypeOf((*MockConfiguration)(nil).GetUrl), key)
}

// PersistInStorage mocks base method.
func (m *MockConfiguration) PersistInStorage(key string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PersistInStorage", key)
}

// PersistInStorage indicates an expected call of PersistInStorage.
func (mr *MockConfigurationMockRecorder) PersistInStorage(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersistInStorage", reflect.TypeOf((*MockConfiguration)(nil).PersistInStorage), key)
}

// Set mocks base method.
func (m *MockConfiguration) Set(key string, value interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", key, value)
}

// Set indicates an expected call of Set.
func (mr *MockConfigurationMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockConfiguration)(nil).Set), key, value)
}

// SetStorage mocks base method.
func (m *MockConfiguration) SetStorage(storage configuration.Storage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStorage", storage)
}

// SetStorage indicates an expected call of SetStorage.
func (mr *MockConfigurationMockRecorder) SetStorage(storage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStorage", reflect.TypeOf((*MockConfiguration)(nil).SetStorage), storage)
}