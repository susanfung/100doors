package main

import (
	"testing"
	approvals "github.com/approvals/go-approval-tests"
)

func TestInitialState(t *testing.T) {
	approvals.VerifyArray(t, Answer())
}