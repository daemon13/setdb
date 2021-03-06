package main

import (
	"encoding/binary"

	"github.com/jmhodges/levigo"
)

// Keys stored in LevelDB for strings
//
// MetadataKey | key = StringLengthValue | string length uint32
//
// For each key:
// StringKey | key = value

func Set(args [][]byte, wb *levigo.WriteBatch) interface{} {
	mk := metaKey(args[0])
	res, err := DB.Get(DefaultReadOptions, mk)
	if err != nil {
		return err
	}
	// If there is a non-string key here, let's delete it first
	if len(res) > 0 && res[0] != StringLengthValue {
		del(args[0], res[0], wb)
	}

	setStringLen(mk, len(args[1]), wb)
	wb.Put(stringKey(args[0]), args[1])

	return ReplyOK
}

func Get(args [][]byte, wb *levigo.WriteBatch) interface{} {
	res, err := DB.Get(DefaultReadOptions, stringKey(args[0]))
	if err != nil {
		return err
	}
	return res
}

func DelString(key []byte, wb *levigo.WriteBatch) {
	wb.Delete(stringKey(key))
}

func setStringLen(key []byte, length int, wb *levigo.WriteBatch) {
	meta := make([]byte, 5)
	meta[0] = StringLengthValue
	binary.BigEndian.PutUint32(meta[1:], uint32(length))
	wb.Put(key, meta)
}

func stringKey(k []byte) []byte {
	key := make([]byte, 5+len(k))
	key[0] = StringKey
	binary.BigEndian.PutUint32(key[1:], uint32(len(k)))
	copy(key[5:], k)
	return key
}

// APPEND
// BITCOUNT
// BITOP
// DECR
// DECRBY
// GETBIT
// GETRANGE
// GETSET
// INCR
// INCRBY
// INCRBYFLOAT
// MGET
// MSET
// MSETNX
// PSETEX
// SETBIT
// SETEX
// SETNX
// SETRANGE
// STRLEN
