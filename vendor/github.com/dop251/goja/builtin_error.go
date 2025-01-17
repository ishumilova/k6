package goja

func (r *Runtime) builtin_Error(args []Value, proto *Object) *Object {
	obj := r.newBaseObject(proto, classError)
	if len(args) > 0 && args[0] != _undefined {
		obj._putProp("message", args[0], true, false, true)
	}
	return obj.val
}

func (r *Runtime) builtin_AggregateError(args []Value, proto *Object) *Object {
	obj := r.newBaseObject(proto, classAggError)
	if len(args) > 1 && args[1] != nil && args[1] != _undefined {
		obj._putProp("message", args[1].toString(), true, false, true)
	}
	var errors []Value
	if len(args) > 0 {
		errors = r.iterableToList(args[0], nil)
	}
	obj._putProp("errors", r.newArrayValues(errors), true, false, true)

	return obj.val
}

func (r *Runtime) createErrorPrototype(name valueString) *Object {
	o := r.newBaseObject(r.global.ErrorPrototype, classObject)
	o._putProp("message", stringEmpty, true, false, true)
	o._putProp("name", name, true, false, true)
	return o.val
}

func (r *Runtime) initErrors() {
	r.global.ErrorPrototype = r.NewObject()
	o := r.global.ErrorPrototype.self
	o._putProp("message", stringEmpty, true, false, true)
	o._putProp("name", stringError, true, false, true)
	o._putProp("toString", r.newNativeFunc(r.error_toString, nil, "toString", nil, 0), true, false, true)

	r.global.Error = r.newNativeFuncConstruct(r.builtin_Error, "Error", r.global.ErrorPrototype, 1)
	r.addToGlobal("Error", r.global.Error)

	r.global.AggregateErrorPrototype = r.createErrorPrototype(stringAggregateError)
	r.global.AggregateError = r.newNativeFuncConstructProto(r.builtin_AggregateError, "AggregateError", r.global.AggregateErrorPrototype, r.global.Error, 2)
	r.addToGlobal("AggregateError", r.global.AggregateError)

	r.global.TypeErrorPrototype = r.createErrorPrototype(stringTypeError)

	r.global.TypeError = r.newNativeFuncConstructProto(r.builtin_Error, "TypeError", r.global.TypeErrorPrototype, r.global.Error, 1)
	r.addToGlobal("TypeError", r.global.TypeError)

	r.global.ReferenceErrorPrototype = r.createErrorPrototype(stringReferenceError)

	r.global.ReferenceError = r.newNativeFuncConstructProto(r.builtin_Error, "ReferenceError", r.global.ReferenceErrorPrototype, r.global.Error, 1)
	r.addToGlobal("ReferenceError", r.global.ReferenceError)

	r.global.SyntaxErrorPrototype = r.createErrorPrototype(stringSyntaxError)

	r.global.SyntaxError = r.newNativeFuncConstructProto(r.builtin_Error, "SyntaxError", r.global.SyntaxErrorPrototype, r.global.Error, 1)
	r.addToGlobal("SyntaxError", r.global.SyntaxError)

	r.global.RangeErrorPrototype = r.createErrorPrototype(stringRangeError)

	r.global.RangeError = r.newNativeFuncConstructProto(r.builtin_Error, "RangeError", r.global.RangeErrorPrototype, r.global.Error, 1)
	r.addToGlobal("RangeError", r.global.RangeError)

	r.global.EvalErrorPrototype = r.createErrorPrototype(stringEvalError)
	o = r.global.EvalErrorPrototype.self
	o._putProp("name", stringEvalError, true, false, true)

	r.global.EvalError = r.newNativeFuncConstructProto(r.builtin_Error, "EvalError", r.global.EvalErrorPrototype, r.global.Error, 1)
	r.addToGlobal("EvalError", r.global.EvalError)

	r.global.URIErrorPrototype = r.createErrorPrototype(stringURIError)

	r.global.URIError = r.newNativeFuncConstructProto(r.builtin_Error, "URIError", r.global.URIErrorPrototype, r.global.Error, 1)
	r.addToGlobal("URIError", r.global.URIError)

	r.global.GoErrorPrototype = r.createErrorPrototype(stringGoError)

	r.global.GoError = r.newNativeFuncConstructProto(r.builtin_Error, "GoError", r.global.GoErrorPrototype, r.global.Error, 1)
	r.addToGlobal("GoError", r.global.GoError)
}
