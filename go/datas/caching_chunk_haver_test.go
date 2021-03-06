// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package datas

import (
	"testing"

	"github.com/attic-labs/noms/go/chunks"
	"github.com/attic-labs/testify/assert"
)

func TestCachingChunkHaver(t *testing.T) {
	assert := assert.New(t)
	storage := &chunks.TestStorage{}
	ts := storage.NewView()
	ccs := newCachingChunkHaver(ts)
	input := "abc"

	c := chunks.NewChunk([]byte(input))
	assert.False(ccs.Has(c.Hash()))
	assert.Equal(ts.Hases, 1)
	assert.False(ccs.Has(c.Hash()))
	assert.Equal(ts.Hases, 1)

	ts.Put(c)
	assert.True(ts.Commit(ts.Root(), ts.Root()))

	ts = storage.NewView()
	ccs = newCachingChunkHaver(ts)
	assert.True(ccs.Has(c.Hash()))
	assert.Equal(ts.Hases, 1)
	assert.True(ccs.Has(c.Hash()))
	assert.Equal(ts.Hases, 1)
}
