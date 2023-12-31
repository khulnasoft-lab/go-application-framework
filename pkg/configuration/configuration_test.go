package configuration

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_FILENAME      string = "test"
	TEST_FILENAME_JSON string = "test.json"
)

func prepareConfigstore(content string) error {
	file, err := CreateConfigurationFile(TEST_FILENAME_JSON)
	if err != nil {
		return err
	}

	// write content to file
	err = ioutil.WriteFile(file, []byte(content), 0755)
	return err
}

func cleanupConfigstore() {
	file, _ := CreateConfigurationFile(TEST_FILENAME_JSON)
	os.RemoveAll(file)
}

func cleanUpEnvVars() {
	os.Unsetenv("VULNMAP_TOKEN")
	os.Unsetenv("VULNMAP_OAUTH_TOKEN")
	os.Unsetenv("VULNMAP_DOCKER_TOKEN")
	os.Unsetenv("VULNMAP_DISABLE_ANALYTICS")
}

func Test_ConfigurationGet_AUTHENTICATION_TOKEN(t *testing.T) {
	os.Unsetenv("VULNMAP_TOKEN")
	expectedValue := "mytoken"
	expectedValue2 := "123456"
	assert.Nil(t, prepareConfigstore(`{"api": "mytoken", "somethingElse": 12}`))

	config := NewFromFiles("test")
	config.AddAlternativeKeys(AUTHENTICATION_TOKEN, []string{"vulnmap_token", "vulnmap_cfg_api", "api"})

	actualValue := config.GetString(AUTHENTICATION_TOKEN)
	assert.Equal(t, expectedValue, actualValue)

	os.Setenv("VULNMAP_TOKEN", expectedValue2)
	actualValue = config.GetString(AUTHENTICATION_TOKEN)
	assert.Equal(t, expectedValue2, actualValue)

	cleanupConfigstore()
	cleanUpEnvVars()
}

func Test_ConfigurationGet_AUTHENTICATION_BEARER_TOKEN(t *testing.T) {
	expectedValue := "anotherToken"
	expectedValueDocker := "dockerTocken"
	assert.Nil(t, prepareConfigstore(`{"api": "mytoken", "somethingElse": 12}`))

	config := NewFromFiles(TEST_FILENAME)
	config.AddAlternativeKeys(AUTHENTICATION_BEARER_TOKEN, []string{"vulnmap_oauth_token", "vulnmap_docker_token"})

	os.Setenv("VULNMAP_OAUTH_TOKEN", expectedValue)
	actualValue := config.GetString(AUTHENTICATION_BEARER_TOKEN)
	assert.Equal(t, expectedValue, actualValue)

	os.Unsetenv("VULNMAP_OAUTH_TOKEN")
	os.Setenv("VULNMAP_DOCKER_TOKEN", expectedValueDocker)
	actualValue = config.GetString(AUTHENTICATION_BEARER_TOKEN)
	assert.Equal(t, expectedValueDocker, actualValue)

	cleanupConfigstore()
	cleanUpEnvVars()
}

func Test_ConfigurationGet_ANALYTICS_DISABLED(t *testing.T) {
	assert.Nil(t, prepareConfigstore(`{"vulnmap_oauth_token": "mytoken", "somethingElse": 12}`))

	config := NewFromFiles(TEST_FILENAME)

	os.Setenv("VULNMAP_DISABLE_ANALYTICS", "1")
	actualValue := config.GetBool(ANALYTICS_DISABLED)
	assert.True(t, actualValue)

	os.Setenv("VULNMAP_DISABLE_ANALYTICS", "0")
	actualValue = config.GetBool(ANALYTICS_DISABLED)
	assert.False(t, actualValue)

	cleanupConfigstore()
	cleanUpEnvVars()
}

func Test_ConfigurationGet_ALTERNATE_KEYS(t *testing.T) {
	alternateKeys := []string{"vulnmap_token", "vulnmap_cfg_api", "api"}

	config := NewFromFiles(TEST_FILENAME)
	config.AddAlternativeKeys(AUTHENTICATION_TOKEN, alternateKeys)

	actualValue := config.GetAlternativeKeys(AUTHENTICATION_TOKEN)
	assert.Equal(t, alternateKeys, actualValue)
}

func Test_ConfigurationGet_unset(t *testing.T) {
	assert.Nil(t, prepareConfigstore(`{"api": "mytoken", "somethingElse": 12}`))

	config := NewFromFiles(TEST_FILENAME)
	actualValue := config.Get("notthere")
	assert.Nil(t, actualValue)

	actualValueString := config.GetString("notthere")
	assert.Empty(t, actualValueString)

	actualValueBool := config.GetBool("notthere")
	assert.False(t, actualValueBool)

	cleanupConfigstore()
}

func Test_ConfigurationSet_differentCases(t *testing.T) {
	assert.Nil(t, prepareConfigstore(`{"api": "mytoken", "somethingElse": 12, "number": 74}`))

	config := NewFromFiles(TEST_FILENAME)

	actualValueString := config.GetString("api")
	assert.Equal(t, "mytoken", actualValueString)

	actualValueInt := config.GetInt("somethingElse")
	assert.Equal(t, 12, actualValueInt)

	actualValueFloat := config.GetFloat64("number")
	assert.Equal(t, 74.0, actualValueFloat)

	config.Set("api", "newToken")
	config.Set("somethingElse", 798)
	config.Set("number", "798.36")

	actualValueString = config.GetString("api")
	assert.Equal(t, "newToken", actualValueString)

	actualValueInt = config.GetInt("somethingElse")
	assert.Equal(t, 798, actualValueInt)

	actualValueFloat = config.GetFloat64("number")
	assert.Equal(t, 798.36, actualValueFloat)

	cleanupConfigstore()
}

func Test_ConfigurationGet_Url(t *testing.T) {
	assert.Nil(t, prepareConfigstore(`{"validUrl": "https://www.khulnasoft.com", "invalidUrl": "something"}`))

	config := NewFromFiles(TEST_FILENAME)

	validUrl := config.GetUrl("validUrl")
	assert.NotNil(t, validUrl)

	invalidUrl := config.GetUrl("invalidUrl")
	assert.NotNil(t, invalidUrl)
}

func Test_ConfigurationGet_StringSlice(t *testing.T) {
	config := New()

	expectedDefault := []string{}

	actual := config.GetStringSlice(RAW_CMD_ARGS)
	assert.Equal(t, expectedDefault, actual)

	expectedNew := []string{"1", "2", "go"}
	config.Set(RAW_CMD_ARGS, expectedNew)

	actualNew := config.GetStringSlice(RAW_CMD_ARGS)
	assert.Equal(t, expectedNew, actualNew)

	actualEmpty := config.GetStringSlice(API_URL)
	assert.Empty(t, actualEmpty)
}

func Test_ConfigurationClone(t *testing.T) {
	assert.Nil(t, prepareConfigstore(`{"api": "mytoken", "somethingElse": 12, "number": 74}`))
	flagset := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flagset.Bool("debug", true, "debugging")
	flagset.Float64("size", 10, "size")

	config := NewFromFiles(TEST_FILENAME)
	config.AddFlagSet(flagset)

	// test cloning of default values
	defaultValueKey := "MyDefault"
	expectedDefaultValue := "-my-default-"
	config.AddDefaultValue(defaultValueKey, StandardDefaultValueFunction(expectedDefaultValue))

	// test cloning of alternate keys
	expectedAlternateValue := "bla"
	notExistingKey := "notExisting"
	alternateValueKey := "AlternateMyDefault"
	config.Set(alternateValueKey, expectedAlternateValue)
	config.AddAlternativeKeys(notExistingKey, []string{alternateValueKey})

	actualValueString := config.GetString("api")
	assert.Equal(t, "mytoken", actualValueString)

	// create the clone
	clonedConfig := config.Clone()

	// manipulate the token
	clonedConfig.Set("api", "10987654321")

	// ensure that the token isn't changed in the original instance
	actualValueString = config.GetString("api")
	assert.Equal(t, "mytoken", actualValueString)

	actualValueString = clonedConfig.GetString("api")
	assert.Equal(t, "10987654321", actualValueString)

	originalKeys := config.AllKeys()
	clonedKeys := clonedConfig.AllKeys()
	sort.Strings(originalKeys)
	sort.Strings(clonedKeys)

	assert.Equal(t, originalKeys, clonedKeys)

	for i := range originalKeys {
		key := originalKeys[i]
		fmt.Println("- key:", key)
		if key != "api" {
			originalValue := config.Get(key)
			clonedValue := clonedConfig.Get(key)
			assert.Equal(t, originalValue, clonedValue)
		}
	}

	actualDefaultValue := clonedConfig.GetString(defaultValueKey)
	assert.Equal(t, expectedDefaultValue, actualDefaultValue)

	actualAlternateValue := clonedConfig.GetString(notExistingKey)
	assert.Equal(t, expectedAlternateValue, actualAlternateValue)

	// we assume that a cloned configuration uses the same storage object. Just the pointer is cloned.
	assert.Equal(t, config.(*extendedViper).storage, clonedConfig.(*extendedViper).storage)

	cleanupConfigstore()
}

func TestNewInMemory(t *testing.T) {
	config := NewInMemory()
	ev := config.(*extendedViper)
	assert.Nil(t, ev.storage)
	assert.NotNil(t, config)
	assert.Equal(t, inMemory, ev.getConfigType())
}

func TestNewFromFiles(t *testing.T) {
	config := NewFromFiles(filepath.Join(t.TempDir(), t.Name()))
	ev := config.(*extendedViper)
	assert.NotNil(t, config)
	assert.Equal(t, jsonFile, ev.getConfigType())
}

func TestNewInMemory_shouldNotBreakWhenTryingToPersist(t *testing.T) {
	config := NewInMemory()

	assert.Nil(t, config.(*extendedViper).storage)
	assert.NotNil(t, config)

	const key = "test"
	const keyValue = "keyValue"
	config.PersistInStorage(key)
	config.Set(key, keyValue)

	assert.Equal(t, config.Get(key), keyValue)
}

func Test_DefaultValuehandling(t *testing.T) {
	keyNoDefault := "name"
	keyWithDefault := "last name"
	valueWithDefault := "default"
	valueExplicitlySet := "explicitly set value"

	config := NewInMemory()
	config.AddDefaultValue(keyWithDefault, func(existingValue interface{}) interface{} {
		if existingValue != nil {
			return existingValue
		}

		return valueWithDefault
	})

	// access value that has a default value
	actualValue, actualWasSet := config.GetAndIsSet(keyWithDefault)
	assert.Equal(t, valueWithDefault, actualValue.(string))
	assert.False(t, actualWasSet)

	// access value that has a default value but is explicitly set
	config.Set(keyWithDefault, valueExplicitlySet)
	actualValue, actualWasSet = config.GetAndIsSet(keyWithDefault)
	assert.Equal(t, valueExplicitlySet, actualValue.(string))
	assert.True(t, actualWasSet)

	// access value that has NO default value
	actualValue, actualWasSet = config.GetAndIsSet(keyNoDefault)
	assert.Nil(t, actualValue)
	assert.False(t, actualWasSet)

}
