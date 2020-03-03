package controller_test

/**
 * Public API tests
 * Run with go test -v ./...
 */
import (
	. "randokiak-api/controller"
	"testing"
)

func TestExportedFunction(t *testing.T) {
	t.Logf("Here IController.unexportedFunction is not visible but IController.ExportedFunction is")
	c := Controller{}
	want := 8
	got := c.ExportedFunction()
	if got != want {
		t.Fatalf(`Fail! Wanted '%v', got '%v'`, want, got)
	}
}
