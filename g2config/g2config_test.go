package g2config

import (
	"context"
	"github.com/docktermj/go-xyzzy-helpers/g2configuration"
	"github.com/docktermj/go-xyzzy-helpers/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	g2config G2config
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context) G2config {

	if g2config == nil {
		g2config = &G2configImpl{}

		moduleName := "Test module name"
		verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
		iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			logger.Fatalf("Cannot construct system configuration: %v", jsonErr)
		}

		initErr := g2config.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			logger.Fatalf("Cannot Init: %v", initErr)
		}
	}
	return g2config
}

func testError(test *testing.T, ctx context.Context, g2config G2config, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		lastException, _ := g2config.GetLastException(ctx)
		assert.FailNow(test, lastException)
	}
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestBuildSimpleSystemConfigurationJson(test *testing.T) {
	actual, err := g2configuration.BuildSimpleSystemConfigurationJson("")
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, actual)
	}
	test.Log("Actual:", actual)
}

func TestGetObject(test *testing.T) {
	ctx := context.TODO()
	getTestObject(ctx)
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestAddDataSource(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	configHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)
	inputJson := `{"DSRC_CODE": "GO_TEST"}`
	actual, err := g2config.AddDataSource(ctx, configHandle, inputJson)
	testError(test, ctx, g2config, err)
	test.Log("Actual:", actual)
	err = g2config.Close(ctx, configHandle)
	testError(test, ctx, g2config, err)
}

func TestClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	err := g2config.ClearLastException(ctx)
	testError(test, ctx, g2config, err)
}

func TestClose(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	configHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)
	err = g2config.Close(ctx, configHandle)
	testError(test, ctx, g2config, err)
}

func TestCreate(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	actual, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)
	test.Log("Actual:", actual)
}

func TestDeleteDataSource(test *testing.T) {

	ctx := context.TODO()
	g2config := getTestObject(ctx)
	configHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)

	actual, err := g2config.ListDataSources(ctx, configHandle)
	testError(test, ctx, g2config, err)
	test.Log("Original:", actual)

	inputJson := `{"DSRC_CODE": "GO_TEST"}`
	actual, err = g2config.AddDataSource(ctx, configHandle, inputJson)
	testError(test, ctx, g2config, err)

	actual, err = g2config.ListDataSources(ctx, configHandle)
	testError(test, ctx, g2config, err)
	test.Log("     Add:", actual)

	err = g2config.DeleteDataSource(ctx, configHandle, inputJson)
	testError(test, ctx, g2config, err)

	actual, err = g2config.ListDataSources(ctx, configHandle)
	testError(test, ctx, g2config, err)
	test.Log("  Delete:", actual)

	err = g2config.Close(ctx, configHandle)
	testError(test, ctx, g2config, err)
}

func TestGetLastException(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	actual, err := g2config.GetLastException(ctx)
	testError(test, ctx, g2config, err)
	test.Log("Actual:", actual)
}

func TestGetLastExceptionCode(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	actual, err := g2config.GetLastExceptionCode(ctx)
	testError(test, ctx, g2config, err)
	test.Log("Actual:", actual)
}

func TestInit(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
	if jsonErr != nil {
		logger.Fatalf("Cannot construct system configuration: %v", jsonErr)
	}
	err := g2config.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2config, err)
}

func TestListDataSources(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	configHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)
	actual, err := g2config.ListDataSources(ctx, configHandle)
	testError(test, ctx, g2config, err)
	test.Log("Actual:", actual)
	err = g2config.Close(ctx, configHandle)
	testError(test, ctx, g2config, err)
}

func TestLoad(test *testing.T) {
}

func TestSave(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	configHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)
	actual, err := g2config.Save(ctx, configHandle)
	testError(test, ctx, g2config, err)
	test.Log("Actual:", actual)
}

func TestDestroy(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	err := g2config.Destroy(ctx)
	testError(test, ctx, g2config, err)
}
