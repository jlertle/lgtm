package mock

import "github.com/stretchr/testify/mock"

import "net/http"
import "github.com/lgtmco/lgtm/model"

type Remote struct {
	mock.Mock
}

func (_m *Remote) GetUser(_a0 http.ResponseWriter, _a1 *http.Request) (*model.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(http.ResponseWriter, *http.Request) *model.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(http.ResponseWriter, *http.Request) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Remote) GetUserToken(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Remote) GetTeams(_a0 *model.User) ([]*model.Team, error) {
	ret := _m.Called(_a0)

	var r0 []*model.Team
	if rf, ok := ret.Get(0).(func(*model.User) []*model.Team); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Remote) GetMembers(_a0 *model.User, _a1 string) ([]*model.Member, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*model.Member
	if rf, ok := ret.Get(0).(func(*model.User, string) []*model.Member); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Member)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Remote) GetRepo(_a0 *model.User, _a1 string, _a2 string) (*model.Repo, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *model.Repo
	if rf, ok := ret.Get(0).(func(*model.User, string, string) *model.Repo); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Repo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Remote) GetPerm(_a0 *model.User, _a1 string, _a2 string) (*model.Perm, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *model.Perm
	if rf, ok := ret.Get(0).(func(*model.User, string, string) *model.Perm); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Perm)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Remote) GetRepos(_a0 *model.User) ([]*model.Repo, error) {
	ret := _m.Called(_a0)

	var r0 []*model.Repo
	if rf, ok := ret.Get(0).(func(*model.User) []*model.Repo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Repo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Remote) SetHook(_a0 *model.User, _a1 *model.Repo, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User, *model.Repo, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *Remote) DelHook(_a0 *model.User, _a1 *model.Repo, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User, *model.Repo, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *Remote) GetComments(_a0 *model.User, _a1 *model.Repo, _a2 int) ([]*model.Comment, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []*model.Comment
	if rf, ok := ret.Get(0).(func(*model.User, *model.Repo, int) []*model.Comment); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User, *model.Repo, int) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Remote) GetContents(_a0 *model.User, _a1 *model.Repo, _a2 string) ([]byte, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*model.User, *model.Repo, string) []byte); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User, *model.Repo, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Remote) SetStatus(_a0 *model.User, _a1 *model.Repo, _a2 int, _a3 bool) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User, *model.Repo, int, bool) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *Remote) GetHook(r *http.Request) (*model.Hook, error) {
	ret := _m.Called(r)

	var r0 *model.Hook
	if rf, ok := ret.Get(0).(func(*http.Request) *model.Hook); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Hook)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
