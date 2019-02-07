//
// Copyright (c) 2019 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package messagebustrigger

import (
	"github.com/edgexfoundry/app-functions-sdk-go/pkg/configuration"
	"github.com/edgexfoundry/app-functions-sdk-go/pkg/runtime"
)

// MessageBusTrigger implements ITrigger to support MessageBusData
type MessageBusTrigger struct {
	Configuration configuration.Configuration
	Runtime       runtime.GolangRuntime
	outputData    interface{}
}

// Initialize ...
func (mb *MessageBusTrigger) Initialize() error {
	return nil
}

// GetConfiguration ...
func (mb *MessageBusTrigger) GetConfiguration() configuration.Configuration {
	//
	return mb.Configuration
}

// GetData ...
func (mb *MessageBusTrigger) GetData() interface{} {
	return "data"
}

// Complete ...
func (mb *MessageBusTrigger) Complete(outputData interface{}) {
	//
	mb.outputData = outputData
}
