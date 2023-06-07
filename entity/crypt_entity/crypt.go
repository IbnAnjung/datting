package crypt_entity

type Crypt interface {
	HashString(str string) (string, error)
	VerifyHash(strHash, plainText string) bool
}
