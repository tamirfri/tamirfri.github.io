package lodash

import (
	_ "embed"
	"fmt"

	v8 "rogchap.com/v8go"
)

var (
	//go:embed node_modules/lodash/core.min.js
	lodashSrc string
)

// Find the first object in `objectList` that contains `subObject`
func Find(objectListJSON, subObjectJSON string) (string, error) {
	runner := v8.NewContext()
	defer runner.Close()
	if _, err := runner.RunScript(lodashSrc, "core.min.js"); err != nil {
		return "", err
	}
	_lodash, err := runner.Global().Get("_")
	if err != nil {
		return "", err
	}
	lodash, err := _lodash.AsObject()
	if err != nil {
		return "", err
	}
	objectList, err := v8.JSONParse(runner, objectListJSON)
	if err != nil {
		return "", fmt.Errorf("%w: %v", err, objectListJSON)
	}
	subObject, err := v8.JSONParse(runner, subObjectJSON)
	if err != nil {
		return "", fmt.Errorf("%w: %v", err, subObjectJSON)
	}
	result, err := lodash.MethodCall("find", objectList, subObject)
	if err != nil {
		return "", err
	}
	return v8.JSONStringify(runner, result)
}
