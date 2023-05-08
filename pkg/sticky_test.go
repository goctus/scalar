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
	"errors"
	scalar "github.com/goctus/scalar/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSticky(t *testing.T) {
	t.Run("retries on error", func(t *testing.T) {
		subject := scalar.NewSticky[bool](newErrorOnFirst())
		firstValue, firstError := subject.Value(context.TODO())
		assert.NotNil(t, firstError)
		assert.False(t, firstValue)
		secondValue, secondError := subject.Value(context.TODO())
		assert.Nil(t, secondError)
		assert.True(t, secondValue)
	})
	t.Run("caches", func(t *testing.T) {
		counter := new(int)
		subject := scalar.NewSticky[bool](newCounting(counter))
		for i := 0; i < 10; i++ {
			value, err := subject.Value(context.TODO())
			assert.True(t, value)
			assert.Nil(t, err)
		}
		assert.Equal(t, 1, *counter, "Expected sticky to only call origin once")
	})
}

type errorOnFirst struct {
	counter *int
}

func newErrorOnFirst() errorOnFirst {
	return errorOnFirst{new(int)}
}

func (eof errorOnFirst) Value(context.Context) (bool, error) {
	if *eof.counter == 0 {
		*eof.counter = 1
		return false, errors.New("first try")
	}
	return true, nil
}

type counting struct {
	counter *int
}

func newCounting(counter *int) counting {
	return counting{counter}
}

func (c counting) Value(context.Context) (bool, error) {
	*c.counter++
	return true, nil
}
