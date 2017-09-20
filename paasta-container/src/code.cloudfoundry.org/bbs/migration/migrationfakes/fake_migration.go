// This file was generated by counterfeiter
package migrationfakes

import (
	"database/sql"
	"sync"

	"code.cloudfoundry.org/bbs/db/etcd"
	"code.cloudfoundry.org/bbs/encryption"
	"code.cloudfoundry.org/bbs/migration"
	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"
)

type FakeMigration struct {
	VersionStub        func() int64
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns     struct {
		result1 int64
	}
	UpStub        func(logger lager.Logger) error
	upMutex       sync.RWMutex
	upArgsForCall []struct {
		logger lager.Logger
	}
	upReturns struct {
		result1 error
	}
	DownStub        func(logger lager.Logger) error
	downMutex       sync.RWMutex
	downArgsForCall []struct {
		logger lager.Logger
	}
	downReturns struct {
		result1 error
	}
	SetStoreClientStub        func(storeClient etcd.StoreClient)
	setStoreClientMutex       sync.RWMutex
	setStoreClientArgsForCall []struct {
		storeClient etcd.StoreClient
	}
	SetCryptorStub        func(cryptor encryption.Cryptor)
	setCryptorMutex       sync.RWMutex
	setCryptorArgsForCall []struct {
		cryptor encryption.Cryptor
	}
	SetClockStub        func(c clock.Clock)
	setClockMutex       sync.RWMutex
	setClockArgsForCall []struct {
		c clock.Clock
	}
	SetRawSQLDBStub        func(rawSQLDB *sql.DB)
	setRawSQLDBMutex       sync.RWMutex
	setRawSQLDBArgsForCall []struct {
		rawSQLDB *sql.DB
	}
	SetDBFlavorStub        func(flavor string)
	setDBFlavorMutex       sync.RWMutex
	setDBFlavorArgsForCall []struct {
		flavor string
	}
	RequiresSQLStub        func() bool
	requiresSQLMutex       sync.RWMutex
	requiresSQLArgsForCall []struct{}
	requiresSQLReturns     struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMigration) Version() int64 {
	fake.versionMutex.Lock()
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	} else {
		return fake.versionReturns.result1
	}
}

func (fake *FakeMigration) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeMigration) VersionReturns(result1 int64) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 int64
	}{result1}
}

func (fake *FakeMigration) Up(logger lager.Logger) error {
	fake.upMutex.Lock()
	fake.upArgsForCall = append(fake.upArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.recordInvocation("Up", []interface{}{logger})
	fake.upMutex.Unlock()
	if fake.UpStub != nil {
		return fake.UpStub(logger)
	} else {
		return fake.upReturns.result1
	}
}

func (fake *FakeMigration) UpCallCount() int {
	fake.upMutex.RLock()
	defer fake.upMutex.RUnlock()
	return len(fake.upArgsForCall)
}

func (fake *FakeMigration) UpArgsForCall(i int) lager.Logger {
	fake.upMutex.RLock()
	defer fake.upMutex.RUnlock()
	return fake.upArgsForCall[i].logger
}

func (fake *FakeMigration) UpReturns(result1 error) {
	fake.UpStub = nil
	fake.upReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMigration) Down(logger lager.Logger) error {
	fake.downMutex.Lock()
	fake.downArgsForCall = append(fake.downArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.recordInvocation("Down", []interface{}{logger})
	fake.downMutex.Unlock()
	if fake.DownStub != nil {
		return fake.DownStub(logger)
	} else {
		return fake.downReturns.result1
	}
}

func (fake *FakeMigration) DownCallCount() int {
	fake.downMutex.RLock()
	defer fake.downMutex.RUnlock()
	return len(fake.downArgsForCall)
}

func (fake *FakeMigration) DownArgsForCall(i int) lager.Logger {
	fake.downMutex.RLock()
	defer fake.downMutex.RUnlock()
	return fake.downArgsForCall[i].logger
}

func (fake *FakeMigration) DownReturns(result1 error) {
	fake.DownStub = nil
	fake.downReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMigration) SetStoreClient(storeClient etcd.StoreClient) {
	fake.setStoreClientMutex.Lock()
	fake.setStoreClientArgsForCall = append(fake.setStoreClientArgsForCall, struct {
		storeClient etcd.StoreClient
	}{storeClient})
	fake.recordInvocation("SetStoreClient", []interface{}{storeClient})
	fake.setStoreClientMutex.Unlock()
	if fake.SetStoreClientStub != nil {
		fake.SetStoreClientStub(storeClient)
	}
}

func (fake *FakeMigration) SetStoreClientCallCount() int {
	fake.setStoreClientMutex.RLock()
	defer fake.setStoreClientMutex.RUnlock()
	return len(fake.setStoreClientArgsForCall)
}

func (fake *FakeMigration) SetStoreClientArgsForCall(i int) etcd.StoreClient {
	fake.setStoreClientMutex.RLock()
	defer fake.setStoreClientMutex.RUnlock()
	return fake.setStoreClientArgsForCall[i].storeClient
}

func (fake *FakeMigration) SetCryptor(cryptor encryption.Cryptor) {
	fake.setCryptorMutex.Lock()
	fake.setCryptorArgsForCall = append(fake.setCryptorArgsForCall, struct {
		cryptor encryption.Cryptor
	}{cryptor})
	fake.recordInvocation("SetCryptor", []interface{}{cryptor})
	fake.setCryptorMutex.Unlock()
	if fake.SetCryptorStub != nil {
		fake.SetCryptorStub(cryptor)
	}
}

func (fake *FakeMigration) SetCryptorCallCount() int {
	fake.setCryptorMutex.RLock()
	defer fake.setCryptorMutex.RUnlock()
	return len(fake.setCryptorArgsForCall)
}

func (fake *FakeMigration) SetCryptorArgsForCall(i int) encryption.Cryptor {
	fake.setCryptorMutex.RLock()
	defer fake.setCryptorMutex.RUnlock()
	return fake.setCryptorArgsForCall[i].cryptor
}

func (fake *FakeMigration) SetClock(c clock.Clock) {
	fake.setClockMutex.Lock()
	fake.setClockArgsForCall = append(fake.setClockArgsForCall, struct {
		c clock.Clock
	}{c})
	fake.recordInvocation("SetClock", []interface{}{c})
	fake.setClockMutex.Unlock()
	if fake.SetClockStub != nil {
		fake.SetClockStub(c)
	}
}

func (fake *FakeMigration) SetClockCallCount() int {
	fake.setClockMutex.RLock()
	defer fake.setClockMutex.RUnlock()
	return len(fake.setClockArgsForCall)
}

func (fake *FakeMigration) SetClockArgsForCall(i int) clock.Clock {
	fake.setClockMutex.RLock()
	defer fake.setClockMutex.RUnlock()
	return fake.setClockArgsForCall[i].c
}

func (fake *FakeMigration) SetRawSQLDB(rawSQLDB *sql.DB) {
	fake.setRawSQLDBMutex.Lock()
	fake.setRawSQLDBArgsForCall = append(fake.setRawSQLDBArgsForCall, struct {
		rawSQLDB *sql.DB
	}{rawSQLDB})
	fake.recordInvocation("SetRawSQLDB", []interface{}{rawSQLDB})
	fake.setRawSQLDBMutex.Unlock()
	if fake.SetRawSQLDBStub != nil {
		fake.SetRawSQLDBStub(rawSQLDB)
	}
}

func (fake *FakeMigration) SetRawSQLDBCallCount() int {
	fake.setRawSQLDBMutex.RLock()
	defer fake.setRawSQLDBMutex.RUnlock()
	return len(fake.setRawSQLDBArgsForCall)
}

func (fake *FakeMigration) SetRawSQLDBArgsForCall(i int) *sql.DB {
	fake.setRawSQLDBMutex.RLock()
	defer fake.setRawSQLDBMutex.RUnlock()
	return fake.setRawSQLDBArgsForCall[i].rawSQLDB
}

func (fake *FakeMigration) SetDBFlavor(flavor string) {
	fake.setDBFlavorMutex.Lock()
	fake.setDBFlavorArgsForCall = append(fake.setDBFlavorArgsForCall, struct {
		flavor string
	}{flavor})
	fake.recordInvocation("SetDBFlavor", []interface{}{flavor})
	fake.setDBFlavorMutex.Unlock()
	if fake.SetDBFlavorStub != nil {
		fake.SetDBFlavorStub(flavor)
	}
}

func (fake *FakeMigration) SetDBFlavorCallCount() int {
	fake.setDBFlavorMutex.RLock()
	defer fake.setDBFlavorMutex.RUnlock()
	return len(fake.setDBFlavorArgsForCall)
}

func (fake *FakeMigration) SetDBFlavorArgsForCall(i int) string {
	fake.setDBFlavorMutex.RLock()
	defer fake.setDBFlavorMutex.RUnlock()
	return fake.setDBFlavorArgsForCall[i].flavor
}

func (fake *FakeMigration) RequiresSQL() bool {
	fake.requiresSQLMutex.Lock()
	fake.requiresSQLArgsForCall = append(fake.requiresSQLArgsForCall, struct{}{})
	fake.recordInvocation("RequiresSQL", []interface{}{})
	fake.requiresSQLMutex.Unlock()
	if fake.RequiresSQLStub != nil {
		return fake.RequiresSQLStub()
	} else {
		return fake.requiresSQLReturns.result1
	}
}

func (fake *FakeMigration) RequiresSQLCallCount() int {
	fake.requiresSQLMutex.RLock()
	defer fake.requiresSQLMutex.RUnlock()
	return len(fake.requiresSQLArgsForCall)
}

func (fake *FakeMigration) RequiresSQLReturns(result1 bool) {
	fake.RequiresSQLStub = nil
	fake.requiresSQLReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeMigration) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	fake.upMutex.RLock()
	defer fake.upMutex.RUnlock()
	fake.downMutex.RLock()
	defer fake.downMutex.RUnlock()
	fake.setStoreClientMutex.RLock()
	defer fake.setStoreClientMutex.RUnlock()
	fake.setCryptorMutex.RLock()
	defer fake.setCryptorMutex.RUnlock()
	fake.setClockMutex.RLock()
	defer fake.setClockMutex.RUnlock()
	fake.setRawSQLDBMutex.RLock()
	defer fake.setRawSQLDBMutex.RUnlock()
	fake.setDBFlavorMutex.RLock()
	defer fake.setDBFlavorMutex.RUnlock()
	fake.requiresSQLMutex.RLock()
	defer fake.requiresSQLMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeMigration) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ migration.Migration = new(FakeMigration)
