// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	domain "cleanarch/domain"
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// StickerUseCase is an autogenerated mock type for the StickerUseCase type
type StickerUseCase struct {
	mock.Mock
}

// GetStickerPackStickers provides a mock function with given fields: id
func (_m *StickerUseCase) GetStickerPackStickers(id int) ([]*domain.Sticker, error) {
	ret := _m.Called(id)

	var r0 []*domain.Sticker
	if rf, ok := ret.Get(0).(func(int) []*domain.Sticker); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Sticker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTrendingStickers provides a mock function with given fields:
func (_m *StickerUseCase) GetTrendingStickers() ([]*domain.Sticker, error) {
	ret := _m.Called()

	var r0 []*domain.Sticker
	if rf, ok := ret.Get(0).(func() []*domain.Sticker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Sticker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStickerUseCase creates a new instance of StickerUseCase. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewStickerUseCase(t testing.TB) *StickerUseCase {
	mock := &StickerUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
