// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package goai_test

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/test/gtest"
)

// Test_ValidationRules_ArrayItems asserts that the length related rules of an array
// field are reported as `minItems` and `maxItems`.
//
// `minLength` and `maxLength` constrain the number of characters of a string and say
// nothing about an array, so an array field used to carry its bound in a keyword that
// no client could apply to it, while `minItems` and `maxItems` stayed empty.
func Test_ValidationRules_ArrayItems(t *testing.T) {
	type Req struct {
		g.Meta `path:"/array-items" method:"POST" tags:"Rules" summary:"Array item rules."`
		Tags   []string `v:"required|max-length:20" dc:"Tags"`
		Ids    []uint64 `v:"length:1,500" dc:"Ids"`
		Codes  []int    `v:"min-length:2" dc:"Codes"`
		Name   string   `v:"min-length:3|max-length:32" dc:"Name"`
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
			req = new(Req)
		)
		err = oai.Add(goai.AddInput{Object: req})
		t.AssertNil(err)

		schema := oai.Components.Schemas.Get("github.com.gogf.gf.v2.net.goai_test.Req").Value

		// An array carries its bound in the item keywords.
		tags := schema.Properties.Get("Tags").Value
		t.Assert(tags.Type, goai.TypeArray)
		t.Assert(tags.MaxItems, 20)
		t.AssertNil(tags.MaxLength)

		ids := schema.Properties.Get("Ids").Value
		t.Assert(ids.MinItems, 1)
		t.Assert(ids.MaxItems, 500)
		t.Assert(ids.MinLength, 0)
		t.AssertNil(ids.MaxLength)

		codes := schema.Properties.Get("Codes").Value
		t.Assert(codes.MinItems, 2)
		t.Assert(codes.MinLength, 0)

		// A string is untouched and keeps carrying its bound in the length keywords.
		name := schema.Properties.Get("Name").Value
		t.Assert(name.Type, goai.TypeString)
		t.Assert(name.MinLength, 3)
		t.Assert(name.MaxLength, 32)
		t.Assert(name.MinItems, 0)
		t.AssertNil(name.MaxItems)
	})
}
