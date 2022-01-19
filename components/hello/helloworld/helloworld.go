/*
 * Copyright 2021 Layotto Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package helloworld

import (
	"context"
	"fmt"

	"mosn.io/layotto/components/hello"
	"mosn.io/layotto/components/trace"
)

type HelloWorld struct {
	Say string
}

var _ hello.HelloService = &HelloWorld{}

func NewHelloWorld() hello.HelloService {
	return &HelloWorld{}
}

func (hw *HelloWorld) Init(config *hello.HelloConfig) error {
	hw.Say = config.HelloString
	return nil
}

func (hw *HelloWorld) Hello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloReponse, error) {
	trace.SetExtraComponentInfo(ctx, fmt.Sprintf("method: %+v", "hello"))
	greetings := hw.Say
	if req.Name != "" {
		greetings = greetings + ", " + req.Name
	}
	return &hello.HelloReponse{
		HelloString: greetings,
	}, nil
}
