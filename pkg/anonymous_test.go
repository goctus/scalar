/*
 * MIT License
 *
 * Copyright (c) 2023-present goctus
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package scalar_test

import (
	"context"
	scalar "github.com/goctus/scalar/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnonymous(t *testing.T) {
	t.Run("uses provided func", func(t *testing.T) {
		want := "Hello, World!"
		subject := scalar.NewAnonymous(func(context.Context) (string, error) {
			return want, nil
		})
		actual, err := subject.Value(context.TODO())
		if err != nil {
			t.Errorf("failed to retrieve actual value: %v", err)
		}
		assert.Equal(t, want, actual)
	})
	t.Run("panics without a func", func(t *testing.T) {
		assert.Panics(t, func() {
			scalar.NewAnonymous[any](nil).Value(context.TODO())
		})
	})
}
