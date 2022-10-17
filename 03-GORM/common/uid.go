package common

import (
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"strconv"
	"strings"
)

// UID is method to generate an virtual unique identifier for whole system
// its structure contains 62 bits: LocalID - ObjectType - SharedID
// 32 bits for localID
// 10 bits for ObjectType
// 18 bits for SharedID

type UID struct {
	localID    uint32
	objectType int
	shardID    uint32
}

func NewUID(localID uint32, objType int, shardID uint32) UID {
	return UID{localID, objType, shardID}
}

// Shard: 1, Objec: 1, ID: 1 => 0001 0001 0001
// 1 << 8 = 0001 0000 0000
// 1 << 4 =         1 0000
// 1 << 0 =              1
// => 0001 0001 0001
func (uid UID) String() string {
	val := uint64(uid.localID)<<28 | uint64(uid.objectType)<<18 | uint64(uid.shardID)<<0
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}

func (uid UID) GetLocalID() uint32 {
	return uid.localID
}

func (uid UID) GetobjectType() int {
	return uid.objectType
}

func (uid UID) GetShardID() uint32 {
	return uid.shardID
}

func DecomposeUID(s string) (UID, error) {
	uid, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		return UID{}, nil
	}

	if (1 << 18) > uid {
		return UID{}, errors.New("wrong uid")
	}

	// x = 1110 1110 0101 => x >> 4 = 1110 1110 & 0000 1111 = 1110
	u := UID{
		localID:    uint32(uid >> 28),
		objectType: int(uid >> 18 & 0x3FF),
		shardID:    uint32(uid >> 0 & 0x3FFFF),
	}
	return u, nil
}

func FromBase58(s string) (UID, error) {
	return DecomposeUID(string(base58.Decode(s)))
}

func (uid UID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", uid.String())), nil
}

func (uid *UID) UnmarshalJSON(data []byte) error {
	decodeUID, err := FromBase58(strings.Replace(string(data), "\"", "", -1))

	if err != nil {
		return err
	}

	uid.localID = decodeUID.localID
	uid.shardID = decodeUID.shardID
	uid.objectType = decodeUID.objectType
	return nil
}
