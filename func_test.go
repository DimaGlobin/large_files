package main

import (
	"net/http/httptest"
	"testing"

	"github.com/lamoda/gonkey/runner"
)

func test_API(t *testing.T) {

	setup_routes()

	srv := httptest.NewServer(nil)

	runner.RunWithTesting(t, &runner.RunWithTestingParams{
		Server:   srv,
		TestsDir: "test/cases",
	})

}
