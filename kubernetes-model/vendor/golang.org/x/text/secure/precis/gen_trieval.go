/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

// entry is the entry of a trie table
// 7..6   property (unassigned, disallowed, maybe, valid)
// 5..0   category
type entry uint8

const (
	propShift = 6
	propMask  = 0xc0
	catMask   = 0x3f
)

func (e entry) property() property { return property(e & propMask) }
func (e entry) category() category { return category(e & catMask) }

type property uint8

// The order of these constants matter. A Profile may consider runes to be
// allowed either from pValid or idDisOrFreePVal.
const (
	unassigned property = iota << propShift
	disallowed
	idDisOrFreePVal // disallowed for Identifier, pValid for FreeForm
	pValid
)

// compute permutations of all properties and specialCategories.
type category uint8

const (
	other category = iota

	// Special rune types
	joiningL
	joiningD
	joiningT
	joiningR
	viramaModifier
	viramaJoinT // Virama + JoiningT
	latinSmallL // U+006c
	greek
	greekJoinT // Greek + JoiningT
	hebrew
	hebrewJoinT // Hebrew + JoiningT
	japanese    // hirigana, katakana, han

	// Special rune types associated with contextual rules defined in
	// https://tools.ietf.org/html/rfc5892#appendix-A.
	// ContextO
	zeroWidthNonJoiner // rule 1
	zeroWidthJoiner    // rule 2
	// ContextJ
	middleDot                // rule 3
	greekLowerNumeralSign    // rule 4
	hebrewPreceding          // rule 5 and 6
	katakanaMiddleDot        // rule 7
	arabicIndicDigit         // rule 8
	extendedArabicIndicDigit // rule 9

	numCategories
)
