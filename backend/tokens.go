package backend

import (
	"crypto"
	"encoding/base64"
	"encoding/hex"
	"strings"

	"github.com/pkg/errors"
	"github.com/zumak/zumo/datatypes"
)

func (b *backend) CreateToken(username, unhashedKey string) (*datatypes.Token, error) {
	// TODO Check user exist....
	if len(unhashedKey) < 8 {
		return nil, errors.Errorf("Token too short")
	}

	token := &datatypes.Token{
		Username:  username,
		HashedKey: hash(unhashedKey),
	}

	token, err := b.Store.PutToken(token)
	if err != nil {
		return nil, errors.Wrap(err, "token create failed")
	}

	return token, nil
}
func (b *backend) Token(tokenStr string) (*datatypes.Token, error) {
	if tokenStr[:6] != "Basic " {
		return nil, errors.Errorf("Not collect token type")
	}

	buf, err := base64.StdEncoding.DecodeString(tokenStr[6:])
	if err != nil {
		return nil, errors.Errorf("token decode fail")
	}

	arr := strings.SplitN(string(buf), ":", 2)
	if len(arr) < 2 {
		return nil, errors.Errorf("Not enought argument")
	}

	return b.Store.GetToken(arr[0], hash(arr[1]))
}

func hash(str string) string {
	hashed := crypto.SHA512.New().Sum([]byte(str))
	return hex.EncodeToString(hashed)
}
