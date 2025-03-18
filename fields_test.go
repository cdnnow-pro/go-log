// SPDX-License-Identifier: MIT

package log_test

import (
	"context"
	"testing"

	. "github.com/cdnnow-pro/logger-go"
	"github.com/stretchr/testify/assert"
)

func TestFieldsContext(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx := context.Background()

		// Act
		f := ExtractFields(ctx)

		// Assert
		assert.Empty(t, f)
	})

	t.Run("one", func(t *testing.T) {
		t.Parallel()

		// Arrange
		ctx := context.Background()

		// Act
		ctx = InjectFields(ctx, "qey", 42)
		f := ExtractFields(ctx)
		ExtractFields(ctx).Delete("qey")

		// Assert
		assert.Equal(t, Fields{"qey", 42}, f)
	})

	t.Run("few with override", func(t *testing.T) {
		// Arrange
		ctx := context.Background()

		// Act
		SetDeduplicationEnabled(true)
		ctx = InjectFields(ctx,
			"qey", 42,
			"str", "anything",
			"bool", true,
		)
		ctx = InjectFields(ctx,
			"str", "second anything",
		)
		f := ExtractFields(ctx)

		// Assert
		assert.Equal(t, Fields{"str", "second anything", "qey", 42, "bool", true}, f)
	})
}
