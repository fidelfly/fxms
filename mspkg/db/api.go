package db

import (
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

func Create(data interface{}) (int64, error) {
	if Engine == nil {
		panic("database engine is not initialized")
	}
	return Engine.Insert(data)
}

func Update(data interface{}, opts ...QueryOption) (int64, error) {
	if Engine == nil {
		panic("database engine is not initialized")
	}
	session := attachOption(Engine.NewSession(), opts...)
	return session.Update(data)
}

func Read(data interface{}, opts ...QueryOption) (bool, error) {
	if Engine == nil {
		panic("database engine is not initialized")
	}
	session := attachOption(Engine.NewSession(), opts...)
	return session.Get(data)
}

func Delete(data interface{}, opts ...QueryOption) (int64, error) {
	if Engine == nil {
		panic("database engine is not initialized")
	}
	session := attachOption(Engine.NewSession(), opts...)
	return session.Delete(data)
}

func attachOption(session *xorm.Session, opts ...QueryOption) *xorm.Session {
	if len(opts) > 0 {
		for _, opt := range opts {
			session = opt(session)
		}
	}
	return session
}

type QueryOption func(session *xorm.Session) *xorm.Session

func Cols(cols ...string) QueryOption {
	return func(session *xorm.Session) *xorm.Session {
		return session.Cols(cols...)
	}
}

func Tablename(name string) QueryOption {
	return func(session *xorm.Session) *xorm.Session {
		return session.Table(name)
	}
}

func ID(id interface{}) QueryOption {
	return func(session *xorm.Session) *xorm.Session {
		return session.ID(id)
	}
}
