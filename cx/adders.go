package base

import (
	// "fmt"
	"github.com/skycoin/skycoin/src/cipher/encoder"
)

func (prgrm *CXProgram) AddPackage(mod *CXPackage) *CXProgram {
	mod.Program = prgrm
	found := false
	for _, md := range prgrm.Packages {
		if md.Name == mod.Name {
			prgrm.CurrentPackage = md
			found = true
			break
		}
	}
	if !found {
		prgrm.Packages = append(prgrm.Packages, mod)
		prgrm.CurrentPackage = mod
	}
	return prgrm
}

func (mod *CXPackage) AddGlobal(def *CXArgument) *CXPackage {
	def.Program = mod.Program
	def.Package = mod
	found := false
	for i, df := range mod.Globals {
		if df.Name == def.Name {
			mod.Globals[i] = def
			found = true
			break
		}
	}
	if !found {
		mod.Globals = append(mod.Globals, def)
	}
	return mod
}

func (mod *CXPackage) AddFunction(fn *CXFunction) *CXPackage {
	fn.Program = mod.Program
	fn.Package = mod

	found := false
	for i, f := range mod.Functions {
		if f.Name == fn.Name {
			mod.Functions[i].Name = fn.Name
			mod.Functions[i].Inputs = fn.Inputs
			mod.Functions[i].Outputs = fn.Outputs
			mod.Functions[i].Expressions = fn.Expressions
			mod.Functions[i].CurrentExpression = fn.CurrentExpression
			mod.Functions[i].Package = fn.Package
			mod.Functions[i].Program = fn.Program
			mod.CurrentFunction = mod.Functions[i]
			found = true
			break
		}
	}
	if !found {
		mod.Functions = append(mod.Functions, fn)
		mod.CurrentFunction = fn
	}

	return mod
}

func (arg *CXArgument) AddValue(val *[]byte) *CXArgument {
	arg.Value = val
	return arg
}

func (arg *CXArgument) AddType(typ string) *CXArgument {
	arg.Typ = typ
	if typCode, found := TypeCodes[typ]; found {
		arg.Type = typCode
		size := GetArgSize(typCode)
		arg.Size = size
		arg.TotalSize = size
	}

	return arg
}

func (mod *CXPackage) AddStruct(strct *CXStruct) *CXPackage {
	prgrm := mod.Program
	strct.Program = prgrm
	strct.Package = mod
	mod.CurrentStruct = strct
	found := false
	for i, s := range mod.Structs {
		if s.Name == strct.Name {
			mod.Structs[i] = strct
			found = true
			break
		}
	}
	if !found {
		mod.Structs = append(mod.Structs, strct)
	}
	return mod
}

func (mod *CXPackage) AddImport(imp *CXPackage) *CXPackage {
	found := false
	for _, im := range mod.Imports {
		if im.Name == imp.Name {
			found = true
			break
		}
	}
	if !found {
		mod.Imports = append(mod.Imports, imp)
	}

	return mod
}

func (strct *CXStruct) AddField(fld *CXArgument) *CXStruct {
	found := false
	for _, fl := range strct.Fields {
		if fl.Name == fld.Name {
			found = true
			break
		}
	}
	if !found {
		fld.CustomType = strct
		strct.Fields = append(strct.Fields, fld)
	}
	return strct
}

func (fn *CXFunction) AddExpression(expr *CXExpression) *CXFunction {
	expr.Program = fn.Program
	expr.Package = fn.Package
	expr.Function = fn
	fn.Expressions = append(fn.Expressions, expr)
	fn.CurrentExpression = expr
	return fn
}

func (fn *CXFunction) AddInput(param *CXArgument) *CXFunction {
	found := false
	for _, inp := range fn.Inputs {
		if inp.Name == param.Name {
			found = true
			break
		}
	}
	if !found {
		fn.Inputs = append(fn.Inputs, param)
	}

	return fn
}

func (fn *CXFunction) AddOutput(param *CXArgument) *CXFunction {
	found := false
	for _, out := range fn.Outputs {
		if out.Name == param.Name {
			found = true
			break
		}
	}
	if !found {
		fn.Outputs = append(fn.Outputs, param)
	}

	return fn
}

func (expr *CXExpression) AddInput(param *CXArgument) *CXExpression {
	// param.Package = expr.Package
	expr.Inputs = append(expr.Inputs, param)
	return expr
}

func (expr *CXExpression) AddOutput(param *CXArgument) *CXExpression {
	// param.Package = expr.Package
	expr.Outputs = append(expr.Outputs, param)
	return expr
}

func (expr *CXExpression) AddLabel(lbl string) *CXExpression {
	expr.Label = lbl
	return expr
}

func (expr *CXExpression) AddOutputName(outName string) *CXExpression {
	if len(expr.Operator.Outputs) > 0 {
		nextOutIdx := len(expr.Outputs)

		var typ string
		if expr.Operator.Name == ID_FN || expr.Operator.Name == INIT_FN {
			var tmp string
			encoder.DeserializeRaw(*expr.Inputs[0].Value, &tmp)

			if expr.Operator.Name == INIT_FN {
				// then tmp is the type (e.g. initDef("i32") to initialize an i32)
				typ = tmp
			} else {
				var err error
				// then tmp is an identifier
				if typ, err = GetIdentType(tmp, expr.FileLine, expr.FileName, expr.Program); err == nil {
				} else {
					panic(err)
				}
			}
		} else {
			typ = expr.Operator.Outputs[nextOutIdx].Typ
		}

		//print(typ + " " + outName + "\n")
		outDef := MakeArgument(outName).AddValue(MakeDefaultValue(expr.Operator.Outputs[nextOutIdx].Typ)).AddType(typ)

		outDef.Package = expr.Package
		outDef.Program = expr.Program

		expr.Outputs = append(expr.Outputs, outDef)
	}

	return expr
}
