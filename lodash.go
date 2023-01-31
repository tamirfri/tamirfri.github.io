package lodash

import (
	_ "embed"
	"fmt"

	v8 "rogchap.com/v8go"
)

var (
	//go:embed node_modules/lodash/core.min.js
	lodashCoreSource string
)

// Find the first object in `objects` that contains `partialObject`
func Find(objects string, partialObject string) (string, error) {
	runner := v8.NewContext()
	defer runner.Close()
	if _, err := runner.RunScript(lodashCoreSource, "core.min.js"); err != nil {
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
	_objects, err := v8.JSONParse(runner, objects)
	if err != nil {
		return "", fmt.Errorf("%w: %v", err, objects)
	}
	_partialObject, err := v8.JSONParse(runner, partialObject)
	if err != nil {
		return "", fmt.Errorf("%w: %v", err, partialObject)
	}
	result, err := lodash.MethodCall("find", _objects, _partialObject)
	if err != nil {
		return "", err
	}
	return v8.JSONStringify(runner, result)
}
