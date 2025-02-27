package log

import (
	"context"
	"slices"
)

type Fields []any

type fieldsKey struct{}

func (f Fields) Iterator() *iter {
	// We start from -2 as we iterate over two items per iteration and first iteration will advance iterator to 0.
	return &iter{i: -2, f: f}
}

type iter struct {
	f Fields
	i int
}

func (i *iter) Next() (k string, v any, ok bool) {
	if i.i >= len(i.f) {
		return "", "", false
	}

	i.i += 2
	if i.i >= len(i.f) {
		return "", "", false
	}

	if i.i+1 == len(i.f) {
		// Non even number of elements, add empty string.
		return i.f[i.i].(string), "", true
	}
	return i.f[i.i].(string), i.f[i.i+1], true
}

func (f Fields) Delete(key string) Fields {
	i := f.Iterator()
	for k, _, ok := i.Next(); ok; k, _, ok = i.Next() {
		if k == key {
			return append(f[:i.i], f[i.i+2:]...)
		}
	}
	return f
}

func (f Fields) With(add Fields) Fields {
	if len(add) == 0 {
		return f
	}

	result := make(Fields, len(add), len(f)+len(add))
	copy(result, add)

	if !deduplicationEnabled.Load() {
		return append(result, f...)
	}

	existing := map[string]struct{}{}
	i := add.Iterator()
	for k, _, ok := i.Next(); ok; k, _, ok = i.Next() {
		existing[k] = struct{}{}
	}

	i = f.Iterator()
	for k, v, ok := i.Next(); ok; k, v, ok = i.Next() {
		if _, ok := existing[k]; !ok {
			result = append(result, k, v)
		}
	}

	return slices.Clip(result)
}

func InjectFields(ctx context.Context, fields ...any) context.Context {
	if len(fields) == 0 {
		return ctx
	}

	if f, ok := ctx.Value(fieldsKey{}).(Fields); ok {
		return context.WithValue(ctx, fieldsKey{}, f.With(fields))
	}

	return context.WithValue(ctx, fieldsKey{}, Fields(fields))
}

func ExtractFields(ctx context.Context) Fields {
	if f, ok := ctx.Value(fieldsKey{}).(Fields); ok {
		return f
	}
	return nil
}
