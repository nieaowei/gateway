package controller_test

import (
	"gateway/lib"
	"testing"
)

func BenchmarkPing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go lib.HttpGET(lib.NewTrace(), "http://127.0.0.1:8880/ping", nil, 1000, nil)
	}
}

func BenchmarkAdminLoginController_AdminLogin(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}
