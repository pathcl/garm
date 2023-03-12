// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	github "github.com/google/go-github/v48/github"
	mock "github.com/stretchr/testify/mock"
)

// GithubEnterpriseClient is an autogenerated mock type for the GithubEnterpriseClient type
type GithubEnterpriseClient struct {
	mock.Mock
}

// CreateRegistrationToken provides a mock function with given fields: ctx, enterprise
func (_m *GithubEnterpriseClient) CreateRegistrationToken(ctx context.Context, enterprise string) (*github.RegistrationToken, *github.Response, error) {
	ret := _m.Called(ctx, enterprise)

	var r0 *github.RegistrationToken
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*github.RegistrationToken, *github.Response, error)); ok {
		return rf(ctx, enterprise)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *github.RegistrationToken); ok {
		r0 = rf(ctx, enterprise)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.RegistrationToken)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *github.Response); ok {
		r1 = rf(ctx, enterprise)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, enterprise)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListRunnerApplicationDownloads provides a mock function with given fields: ctx, enterprise
func (_m *GithubEnterpriseClient) ListRunnerApplicationDownloads(ctx context.Context, enterprise string) ([]*github.RunnerApplicationDownload, *github.Response, error) {
	ret := _m.Called(ctx, enterprise)

	var r0 []*github.RunnerApplicationDownload
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*github.RunnerApplicationDownload, *github.Response, error)); ok {
		return rf(ctx, enterprise)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*github.RunnerApplicationDownload); ok {
		r0 = rf(ctx, enterprise)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*github.RunnerApplicationDownload)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *github.Response); ok {
		r1 = rf(ctx, enterprise)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, enterprise)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListRunners provides a mock function with given fields: ctx, enterprise, opts
func (_m *GithubEnterpriseClient) ListRunners(ctx context.Context, enterprise string, opts *github.ListOptions) (*github.Runners, *github.Response, error) {
	ret := _m.Called(ctx, enterprise, opts)

	var r0 *github.Runners
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *github.ListOptions) (*github.Runners, *github.Response, error)); ok {
		return rf(ctx, enterprise, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *github.ListOptions) *github.Runners); ok {
		r0 = rf(ctx, enterprise, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Runners)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *github.ListOptions) *github.Response); ok {
		r1 = rf(ctx, enterprise, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, *github.ListOptions) error); ok {
		r2 = rf(ctx, enterprise, opts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RemoveRunner provides a mock function with given fields: ctx, enterprise, runnerID
func (_m *GithubEnterpriseClient) RemoveRunner(ctx context.Context, enterprise string, runnerID int64) (*github.Response, error) {
	ret := _m.Called(ctx, enterprise, runnerID)

	var r0 *github.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) (*github.Response, error)); ok {
		return rf(ctx, enterprise, runnerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) *github.Response); ok {
		r0 = rf(ctx, enterprise, runnerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, enterprise, runnerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGithubEnterpriseClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewGithubEnterpriseClient creates a new instance of GithubEnterpriseClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGithubEnterpriseClient(t mockConstructorTestingTNewGithubEnterpriseClient) *GithubEnterpriseClient {
	mock := &GithubEnterpriseClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
