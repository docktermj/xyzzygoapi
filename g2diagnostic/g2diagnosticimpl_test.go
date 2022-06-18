package g2diagnostic

import (
	"context"
	"testing"

	"github.com/docktermj/xyzzygoapi/g2helper"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getTestObject() (G2diagnostic, error) {
	var err error = nil
	g2diagnostic := G2diagnosticImpl{}
	ctx := context.TODO()

	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, _ := g2helper.GetSimpleSystemConfigurationJson(ctx)

	err = g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
	return &g2diagnostic, err
}

func testError(test *testing.T, ctx context.Context, g2diagnostic G2diagnostic, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		lastException, _ := g2diagnostic.GetLastException(ctx)
		assert.FailNow(test, lastException)
	}
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

// Reference: https://medium.com/nerd-for-tech/setup-and-teardown-unit-test-in-go-bd6fa1b785cd
func setupSuite(test testing.TB) func(test testing.TB) {
	test.Log("setup suite")

	// Return a function to teardown the test
	return func(test testing.TB) {
		test.Log("teardown suite")
	}
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestCheckDBPerf(test *testing.T) {
	teardownSuite := setupSuite(test)
	defer teardownSuite(test)

	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	secondsToRun := 1
	actual, err := g2diagnostic.CheckDBPerf(ctx, secondsToRun)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Actual:", actual)
}

func TestClearLastException(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	g2diagnostic.ClearLastException(ctx)
}

// FAIL:
func TestEntityListBySize(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	aSize := 10

	aHandle, err := g2diagnostic.GetEntityListBySize(ctx, aSize)
	testError(test, ctx, g2diagnostic, err)

	anEntity, err := g2diagnostic.FetchNextEntityBySize(ctx, aHandle)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Entity:", anEntity)

	err = g2diagnostic.CloseEntityListBySize(ctx, aHandle)
	testError(test, ctx, g2diagnostic, err)
}

func TestDestroy(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	err := g2diagnostic.Destroy(ctx)
	testError(test, ctx, g2diagnostic, err)
}

func TestFindEntitiesByFeatureIDs(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	features := "{\"ENTITY_ID\":1,\"LIB_FEAT_IDS\":[1,3,4]}"
	actual, err := g2diagnostic.FindEntitiesByFeatureIDs(ctx, features)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Actual:", actual)
}

func TestGetAvailableMemory(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, err := g2diagnostic.GetAvailableMemory(ctx)
	testError(test, ctx, g2diagnostic, err)
	assert.Greater(test, actual, int64(0))
	test.Log("Actual:", actual)
}

func TestGetDataSourceCounts(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, err := g2diagnostic.GetDataSourceCounts(ctx)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Data Source counts:", actual)
}

func TestGetDBInfo(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, err := g2diagnostic.GetDBInfo(ctx)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Actual:", actual)
}

func TestGetEntityDetails(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	entityID := int64(1)
	includeInternalFeatures := 1
	actual, err := g2diagnostic.GetEntityDetails(ctx, entityID, includeInternalFeatures)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Actual:", actual)
}

func TestGetEntityResume(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	entityID := int64(1)
	actual, err := g2diagnostic.GetEntityResume(ctx, entityID)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Actual:", actual)
}

func TestGetEntitySizeBreakdown(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	minimumEntitySize := 1
	includeInternalFeatures := 1
	actual, err := g2diagnostic.GetEntitySizeBreakdown(ctx, minimumEntitySize, includeInternalFeatures)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Actual:", actual)
}

func TestGetFeature(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	libFeatID := int64(1)
	actual, err := g2diagnostic.GetFeature(ctx, libFeatID)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Actual:", actual)
}

func TestGetGenericFeatures(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	featureType := "PHONE"
	maximumEstimatedCount := 10
	actual, err := g2diagnostic.GetGenericFeatures(ctx, featureType, maximumEstimatedCount)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Actual:", actual)
}

func TestGetLastException(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, err := g2diagnostic.GetLastException(ctx)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Actual:", actual)
}

func TestGetLastExceptionCode(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, err := g2diagnostic.GetLastExceptionCode(ctx)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Actual:", actual)
}

func TestGetLogicalCores(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, err := g2diagnostic.GetLogicalCores(ctx)
	testError(test, ctx, g2diagnostic, err)
	assert.Greater(test, actual, 0)
	test.Log("Actual:", actual)
}

func TestGetMappingStatistics(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	includeInternalFeatures := 1
	actual, err := g2diagnostic.GetMappingStatistics(ctx, includeInternalFeatures)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Actual:", actual)
}

func TestGetPhysicalCores(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, err := g2diagnostic.GetPhysicalCores(ctx)
	testError(test, ctx, g2diagnostic, err)
	assert.Greater(test, actual, 0)
	test.Log("Actual:", actual)
}

func TestGetRelationshipDetails(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	relationshipID := int64(1)
	includeInternalFeatures := 1
	actual, err := g2diagnostic.GetRelationshipDetails(ctx, relationshipID, includeInternalFeatures)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Actual:", actual)
}

func TestGetResolutionStatistics(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, err := g2diagnostic.GetResolutionStatistics(ctx)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Actual:", actual)
}

func TestGetTotalSystemMemory(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, err := g2diagnostic.GetTotalSystemMemory(ctx)
	testError(test, ctx, g2diagnostic, err)
	assert.Greater(test, actual, int64(0))
	test.Log("Actual:", actual)
}

func TestInit(test *testing.T) {
	g2diagnostic := &G2diagnosticImpl{}
	ctx := context.TODO()
	moduleName := "Test module name"
	verboseLogging := 0
	iniParams, _ := g2helper.GetSimpleSystemConfigurationJson(ctx)
	err := g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2diagnostic, err)
}

func TestInitWithConfigID(test *testing.T) {
	g2diagnostic := &G2diagnosticImpl{}
	ctx := context.TODO()
	moduleName := "Test module name"
	initConfigID := int64(1)
	verboseLogging := 0
	iniParams, _ := g2helper.GetSimpleSystemConfigurationJson(ctx)
	err := g2diagnostic.InitWithConfigID(ctx, moduleName, iniParams, initConfigID, verboseLogging)
	testError(test, ctx, g2diagnostic, err)
}

func TestReinit(test *testing.T) {
	g2diagnostic := &G2diagnosticImpl{}
	ctx := context.TODO()
	initConfigID := int64(4019066234)
	err := g2diagnostic.Reinit(ctx, initConfigID)
	testError(test, ctx, g2diagnostic, err)
}
