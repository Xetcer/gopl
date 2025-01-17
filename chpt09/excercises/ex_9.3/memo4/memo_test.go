package memo_test

import (
	"testing"

	memo "gopl/chpt09/memo4"
	memotest "gopl/chpt09/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}

func TestConcurrentCancelOK(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.ConcurrentWithCancelOk(t, m)
}

func TestConcurrentCancelAll(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.ConcurrentWitchCancelAll(t, m)
}

func TestConcurrentCancelPartly(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.ConcurrentWitchCancelPartlty(t, m)
}
