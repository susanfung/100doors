package main

import (
	"testing"
	approvals "github.com/approvals/go-approval-tests"
)

func TestInitialState(t *testing.T) {
	approvals.VerifyArray(t, Answer(0))
}

func TestFirstPass(t *testing.T) {
	approvals.VerifyArray(t, Answer(1))
}

func TestSecondPass(t *testing.T) {
	approvals.VerifyArray(t, Answer(2))
}