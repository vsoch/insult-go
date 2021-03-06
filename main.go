// Copyright 2020 Vanessa Sochat. All rights reserved.
// Use of this source code is governed by the Polyform Strict license
// that can be found in the LICENSE file and available at
// https://polyformproject.org/licenses/noncommercial/1.0.0

package main

import "syscall/js"

func main() {

	c := make(chan struct{}, 0)
	js.Global().Set("generateInsult", js.FuncOf(generateInsult))
	js.Global().Set("generateRandomInsult", js.FuncOf(generateRandomInsult))
	<-c
}
