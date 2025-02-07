/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2022 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package stdlib

import (
	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/cadence/runtime/interpreter"
	"github.com/onflow/cadence/runtime/sema"
)

func NewSignatureAlgorithmCase(inter *interpreter.Interpreter, rawValue uint8) *interpreter.CompositeValue {
	return interpreter.NewEnumCaseValue(
		inter,
		interpreter.ReturnEmptyLocationRange,
		sema.SignatureAlgorithmType,
		interpreter.NewUInt8Value(inter, func() uint8 {
			return rawValue
		}),
		nil,
	)
}

var signatureAlgorithmConstructor = StandardLibraryValue{
	Name: sema.SignatureAlgorithmTypeName,
	Type: cryptoAlgorithmEnumConstructorType(
		sema.SignatureAlgorithmType,
		sema.SignatureAlgorithms,
	),
	ValueFactory: func(inter *interpreter.Interpreter) interpreter.Value {
		return cryptoAlgorithmEnumValue(
			inter,
			sema.SignatureAlgorithmType,
			sema.SignatureAlgorithms,
			NewSignatureAlgorithmCase,
		)
	},
	Kind: common.DeclarationKindEnum,
}
