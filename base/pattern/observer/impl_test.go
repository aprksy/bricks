package observer_test

import (
	"fmt"
	"testing"

	obs "github.com/aprksy/bricks/base/pattern/observer"
	"github.com/stretchr/testify/assert"
)

func TestNewSimpleSubject(t *testing.T) {
	testCases := []struct {
		name  string
		oid   uint
		key   string
		value int
	}{
		{name: "subject 1", oid: 1, key: "key-1", value: 1},
		{name: "subject 2", oid: 2, key: "key-2", value: 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			instance := obs.NewSimpleSubject[uint, int](tc.oid, tc.key, tc.value)
			assert.NotNilf(t, instance, "instance should not be nil")
		})
	}
}

func TestNewSimpleObserver(t *testing.T) {
	testCases := []struct {
		name  string
		oid   uint
		key   string
		value int
	}{
		{name: "subject 1", oid: 1, key: "key-1", value: 1},
		{name: "subject 2", oid: 2, key: "key-2", value: 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			instance := obs.NewSimpleObserver[uint, int](tc.oid, nil)
			assert.NotNilf(t, instance, "instance should not be nil")
		})
	}
}

func TestNewSubjectManager(t *testing.T) {
	t.Run("NewSubjectManager", func(t *testing.T) {
		instance := obs.NewSubjectManager[uint]()
		assert.NotNilf(t, instance, "instance should not be nil")
	})
}

func TestSubscribe(t *testing.T) {
	chDone := make(chan bool)

	subject := obs.NewSimpleSubject[uint, int](1, "value", 273)
	observer := obs.NewSimpleObserver[uint, int](2, func(key string, value int) {
		chDone <- true
	})

	testCases := []struct {
		name    string
		oid     uint
		key     string
		subject obs.Subject[uint, int]
	}{
		{name: "observer 1", oid: 1, key: "value", subject: nil},
		{name: "observer 2", oid: 2, key: "value", subject: subject},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ptrSubsid, err := observer.Subscribe(tc.subject, tc.key)

			if i == 0 {
				assert.Nil(t, ptrSubsid, "ptrSubsid should be nil")
				assert.NotNilf(t, err, "err should not be nil")
			} else {
				assert.Nil(t, err, "err should be nil")
				assert.NotNilf(t, ptrSubsid, "ptrSubsid should not be nil")
			}
		})
	}
}

func TestSubscribeByKey(t *testing.T) {
	chDone := make(chan bool)

	subjectMgr := obs.NewSubjectManager[uint]()
	subject := obs.NewSimpleSubject[uint, int](1, "value", 273)
	err := obs.AddSubjects(subjectMgr, subject)
	assert.Nil(t, err, "err should equal nil")

	subject1 := obs.NewSimpleSubject[uint, int](2, "value", 273)
	err = obs.AddSubjects(subjectMgr, subject1)
	assert.EqualErrorf(t, err, obs.ErrKeyExists, "err should equal %s", obs.ErrKeyExists)

	observer0 := obs.NewSimpleObserver[uint, int](3, nil)
	observer := obs.NewSimpleObserverWithSubjectManager[uint, int](2, func(key string, value int) {
		chDone <- true
	}, subjectMgr)

	testCases := []struct {
		name     string
		oid      uint
		observer obs.Observer[uint, int]
		key      string
	}{
		{name: "observer 0", oid: 1, observer: observer0, key: "non-existent-key"},
		{name: "observer 1", oid: 2, observer: observer, key: "non-existent-key"},
		{name: "observer 2", oid: 3, observer: observer, key: "value"},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ptrSubsid, mysubject, err := tc.observer.SubscribeByKey(tc.key)

			if i == 0 {
				assert.Nil(t, ptrSubsid, "ptrSubsid should be nil")
				assert.Nil(t, mysubject, "mysubject should be nil")
				assert.EqualErrorf(t, err, obs.ErrSubjectMgrNil, "err should be %s", obs.ErrSubjectMgrNil)
			} else if i == 1 {
				assert.Nil(t, ptrSubsid, "ptrSubsid should be nil")
				assert.Nil(t, mysubject, "mysubject should be nil")
				assert.NotNilf(t, err, "err should not be nil")
			} else {
				assert.Nil(t, err, "err should be nil")
				assert.NotNilf(t, ptrSubsid, "ptrSubsid should not be nil")
				assert.NotNilf(t, mysubject, "mysubject should not be nil")
			}
		})
	}
}

func TestUnsubscribe(t *testing.T) {
	subject := obs.NewSimpleSubject[uint, int](1, "value", 273)
	observer := obs.NewSimpleObserver[uint, int](2, nil)
	ptrSubsid, _ := observer.Subscribe(subject, "value")

	err := observer.Unsubscribe("non-existing")
	assert.NotNilf(t, err, "err should not be nil")

	err = observer.Unsubscribe("non-existent-subsid")
	assert.NotNil(t, err, "err should not be nil")

	err = observer.Unsubscribe(*ptrSubsid)
	assert.Nil(t, err, "err should be nil")

	err = subject.Remove("non-existent-subsid")
	assert.NotNil(t, err, "err should not be nil")
}

func TestInjectExtract(t *testing.T) {
	chDone := make(chan bool)
	subject := obs.NewSimpleSubject[uint, int](1, "value", 0)
	observer := obs.NewSimpleObserver[uint, int](2, func(key string, value int) {
		chDone <- true
	})
	observer.Subscribe(subject, "value")

	correctKey := "value"
	testCases := []struct {
		name  string
		key   string
		value int
	}{
		{name: "extract with wrong key", key: "wrong-key", value: 1004},
		{name: "extract with correct key", key: correctKey, value: 101},
		{name: "extract with correct key", key: correctKey, value: 43},
		{name: "extract with correct key", key: correctKey, value: 567},
		{name: "extract with correct key", key: correctKey, value: 113},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			value := tc.value
			key := tc.key
			subject.Inject(value)

			<-chDone
			outValue := subject.Extract()
			assert.Equal(t, value, outValue, "value should be equal to tc.value")

			ptrValue, err := observer.Extract(key)
			if i == 0 {
				assert.Nil(t, ptrValue, "ptrValue should be nil")
				assert.NotNilf(t, err, "err should not be nil")
			} else {
				fmt.Printf("%d, %d, %d\n", value, outValue, *ptrValue)
				assert.Nil(t, err, "err should be nil")
				assert.NotNilf(t, ptrValue, "ptrValue should not be nil")
				assert.Containsf(t, []int{0, 1004, 101, 43, 567, 113}, *ptrValue, "*ptrValue should be among the values")
			}
		})
	}
}

func TestInject(t *testing.T) {
	subjectMgr := obs.NewSubjectManager[uint]()
	subject := obs.NewSimpleSubject[uint, int](1, "value", 273)
	obs.AddSubjects(subjectMgr, subject)

	err := obs.Inject[uint, int](subjectMgr, "non-existent-key", 100)
	assert.EqualErrorf(t, err, obs.ErrKeyNotFound, "err should equal %s", obs.ErrKeyNotFound)

	err = obs.Inject[uint, int](subjectMgr, "value", 101)
	assert.Nil(t, err, "err should equal nil")
}
