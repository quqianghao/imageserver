package cache

import (
	"crypto/sha256"
	"testing"

	"github.com/pierrre/imageserver"
	imageserver_provider "github.com/pierrre/imageserver/provider"
)

func TestTODO(t *testing.T) {
	t.Log("TODO")
}

func TestProviderInterface(t *testing.T) {
	var _ imageserver_provider.Provider = &Provider{}
}

func TestKeyGeneratorFuncInterface(t *testing.T) {
	var _ KeyGenerator = KeyGeneratorFunc(nil)
}

func TestNewSourceHashKeyGenerator(t *testing.T) {
	g := NewSourceHashKeyGenerator(sha256.New)
	g.GetKey("foobar", imageserver.Parameters{})
}

func TestPrefixKeyGeneratorInterface(t *testing.T) {
	var _ KeyGenerator = &PrefixKeyGenerator{}
}

func TestPrefixKeyGenerator(t *testing.T) {
	kgf := KeyGeneratorFunc(func(source interface{}, parameters imageserver.Parameters) string {
		return "bar"
	})
	pkg := &PrefixKeyGenerator{
		KeyGenerator: kgf,
		Prefix:       "foo",
	}
	key := pkg.GetKey(nil, imageserver.Parameters{})
	if key != "foobar" {
		t.Fatal("not equal")
	}
}
