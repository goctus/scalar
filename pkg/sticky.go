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

package scalar

import (
	"context"
	"sync"
)

// Sticky is a scalar that only calculates its value once.
type Sticky[T any] struct {
	origin Scalar[T]

	m     *sync.Mutex
	cache Scalar[T]
}

// NewSticky creates a new Sticky.
func NewSticky[T any](origin Scalar[T]) Sticky[T] {
	return Sticky[T]{origin, &sync.Mutex{}, nil}
}

func (s Sticky[T]) Value(ctx context.Context) (T, error) {
	if s.cache == nil {
		s.m.Lock()
		defer s.m.Unlock()
		origin, err := s.origin.Value(ctx)
		if err == nil {
			s.cache = NewConstant(origin)
		} else {
			s.cache = NewNothing[T](err)
		}
	}
	return s.cache.Value(ctx)
}
