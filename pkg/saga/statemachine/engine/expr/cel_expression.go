/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package expr

import (
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/decls"
)

type CELExpression struct {
	env        *cel.Env
	program    cel.Program
	expression string
}

var _ Expression = (*CELExpression)(nil)

func NewCELExpression(expression string) (*CELExpression, error) {
	// Create the standard environment.
	env, err := cel.NewEnv(
		cel.VariableDecls(
			decls.NewVariable("elContext", cel.DynType),
		),
	)

	if err != nil {
		return nil, err
	}

	// Check that the expression compiles and returns a String.
	ast, issues := env.Compile(expression)
	// Report syntax errors, if present.
	if issues != nil && issues.Err() != nil {
		return nil, issues.Err()
	}

	// Type-check the expression ofr correctness.
	checkedAst, issues := env.Check(ast)
	if issues.Err() != nil {
		return nil, issues.Err()
	}

	program, err := env.Program(checkedAst)
	if err != nil {
		return nil, err
	}

	CELExpression := &CELExpression{
		env:        env,
		program:    program,
		expression: expression,
	}

	return CELExpression, nil
}

func (c *CELExpression) Value(elContext any) any {
	result, _, err := c.program.Eval(map[string]any{
		"elContext": elContext,
	})
	if err != nil {
		return err
	}
	return result.Value()
}

func (c *CELExpression) SetValue(val any, elContext any) {
	panic("implement me")
}

func (c *CELExpression) ExpressionString() string {
	return c.expression
}
