/*
 * Copyright 2019 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package aeternity

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	address := "ak_qrKKpCkCYej3MakQqgxZZNCfEwV1T6vYvKuEUr9j3hDY3aXYK"
	//pubkey, _ := hex.DecodeString("6eeba7851c2ddb4dd4d47588f1e1a204f88d6e0eed3026caf3b8ec3e432f9d89")
	message, _ := hex.DecodeString("61655f6d61696e6e6574f85b0c01a1016eeba7851c2ddb4dd4d47588f1e1a204f88d6e0eed3026caf3b8ec3e432f9d89a1016e6490ba9ffa3ed276048e23c52f09a7622e02111124e9c770d1a6ac11a723c6872386f26fc100008612309ce540008301dfcb0180")
	signature, _ := hex.DecodeString("6d3987133430a65b834052e3ca22282ca5b7741ca46d8a5c295338a8b890e340c42aaffa1f887d2016b25e2bb0b652c310a02f03c3f3c11abe9e87cc9a0e9f08")
	pub, _ := Decode(address)
	fmt.Printf("pubkey: %s\n", hex.EncodeToString(pub))
	valid, err := Verify(address, message, signature)
	if err != nil {
		t.Errorf("verify failed, error: %v", err)
		return
	}
	fmt.Printf("valid: %v\n", valid)
}
