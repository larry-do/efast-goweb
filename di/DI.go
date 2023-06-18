package DI

import (
	"github.com/rs/zerolog/log"
	"reflect"
)

type contextValue[V any] struct {
	value V
	scope Scope
	constructor func() V
}

var context = make(map[reflect.Type]map[any]contextValue[any])

func Inject(typeOf reflect.Type, qualifier any, value any, scope Scope, constructor func() any) {
	if qualifier == nil {
		qualifier = typeOf.String()
	}

	ctxValue, ok := context[typeOf]
	if !ok {
		ctxValue = make(map[any]contextValue[any])
		context[typeOf] = ctxValue
	}

	ctxValue[qualifier] = contextValue[any]{
		value: value,
		scope: scope,
		constructor: constructor,
	}

	log.Debug().
		Str("type", typeOf.String()).
		Any("value", value).
		Any("scope", scope).
		Any("qualifier", qualifier).
		Msg("Injected to DI")
}

func InjectSingleton(typeOf reflect.Type, value any) {
	Inject(typeOf, nil, value, SINGLETON, nil)
}

func InjectSingletonWithQualifier(typeOf reflect.Type, qualifier any, value any) {
	Inject(typeOf, qualifier, value, SINGLETON, nil)
}

func InjectPrototype(typeOf reflect.Type, constructor func() any) {
	Inject(typeOf, nil, nil, PROTOTYPE, constructor)
}

func GetWithQualifier[I any](qualifier any) I {
	typeOf := reflect.TypeOf(*new(I))
	if qualifier == nil {
		qualifier = typeOf.String()
	}

	ctxValue, ok := context[typeOf]
	if !ok {
		var i I
		return i
	}
	value, ok := ctxValue[qualifier]
	if !ok {
		var i I
		return i
	}

	if value.scope == SINGLETON {
		return value.value.(I)
	}

	if value.constructor != nil {
		return value.constructor().(I)
	}

	return *new(I)
}

func Get[I any]() I {
	return GetWithQualifier[I](nil)
}

const (
	PROTOTYPE = "PROTOTYPE"
	SINGLETON = "SINGLETON"
)

type Scope string