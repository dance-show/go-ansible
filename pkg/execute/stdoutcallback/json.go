package stdoutcallback

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	jsonresults "github.com/apenella/go-ansible/v2/pkg/execute/result/json"
)

const (
	// JSONStdoutCallback ansible screen output as json
	JSONStdoutCallback = "json"
)

// JSONStdoutCallbackExecute defines an executor to run an ansible command with a json stdout callback
type JSONStdoutCallbackExecute struct {
	executor ExecutorQuietStdoutCallbackSetter
}

// NewJSONStdoutCallbackExecute creates a JSONStdoutCallbackExecute
func NewJSONStdoutCallbackExecute(executor ExecutorQuietStdoutCallbackSetter) *JSONStdoutCallbackExecute {
	return &JSONStdoutCallbackExecute{executor: executor}
}

// WithExecutor sets the executor for the JSONStdoutCallbackExecute
func (e *JSONStdoutCallbackExecute) WithExecutor(exec ExecutorQuietStdoutCallbackSetter) *JSONStdoutCallbackExecute {
	e.executor = exec
	return e
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *JSONStdoutCallbackExecute) Execute(ctx context.Context) error {

	if e.executor == nil {
		return fmt.Errorf("JSONStdoutCallbackExecute executor requires an executor")
	}

	e.executor.Quiet()
	e.executor.WithOutput(jsonresults.NewJSONStdoutCallbackResults())

	return configuration.NewAnsibleWithConfigurationSettingsExecute(e.executor,
		configuration.WithAnsibleStdoutCallback(JSONStdoutCallback),
	).Execute(ctx)
}
