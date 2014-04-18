package results_test

import "github.com/zaiuz/testutil"
import z "github.com/zaiuz/zaiuz"

type contextHolder struct {
	context       *z.Context
	calledContext *z.Context
}

func newContextHolder() *contextHolder {
	context := testutil.NewTestContext()
	return &contextHolder{context, nil}
}
