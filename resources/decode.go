package delorian

import (
	"github.com/pelletier/go-toml"
	"os"
)

// PackDecoder is an unmarshaller of resources packs, it decodes it to a file
type PackDecoder struct {
	// Pack is the chosen pack
	Pack ResourcePack
}

func (p *PackDecoder) Decode(pack ResourcePack) ResourcePack {
	b, err := toml.Marshal(pack)
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(b, &p.Pack)
	if err != nil {
		panic(err)
	}
	for h := range b {
		v := toml.Unmarshal(b, &p.Pack)
		if v.Error != nil {
			panic(v)
		}
		t, err := toml.Marshal(h)
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(t)
		os.Stdout.Write(b)
	}
	go func() {
		defer func() {
			os.Stdout.Close()
		}()
	}()
	return p.Pack
}
