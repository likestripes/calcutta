package calcutta

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _esc_localFS struct{}

var _esc_local _esc_localFS

type _esc_staticFS struct{}

var _esc_static _esc_staticFS

type _esc_file struct {
	compressed string
	size       int64
	local      string
	isDir      bool

	data []byte
	once sync.Once
	name string
}

func (_esc_localFS) Open(name string) (http.File, error) {
	f, present := _esc_data[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_esc_staticFS) Open(name string) (http.File, error) {
	f, present := _esc_data[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		gr, err = gzip.NewReader(bytes.NewBufferString(f.compressed))
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (f *_esc_file) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_esc_file
	}
	return &httpFile{
		Reader:    bytes.NewReader(f.data),
		_esc_file: f,
	}, nil
}

func (f *_esc_file) Close() error {
	return nil
}

func (f *_esc_file) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_esc_file) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_esc_file) Name() string {
	return f.name
}

func (f *_esc_file) Size() int64 {
	return f.size
}

func (f *_esc_file) Mode() os.FileMode {
	return 0
}

func (f *_esc_file) ModTime() time.Time {
	return time.Time{}
}

func (f *_esc_file) IsDir() bool {
	return f.isDir
}

func (f *_esc_file) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _esc_local
	}
	return _esc_static
}

var _esc_data = map[string]*_esc_file{

	"/base.html": {
		local: "templates/base.html",
		size:  1358,
		compressed: "\x1f\x8b\b\x00\x00\tn\x88\x00\xff\x94TAO\xe3<\x10\xbd\xe7W\xcc\a\xfan\xa4\x10\xd2.\xc8DH{\xdb\xc3j\xff\xc2ʱ'\xe9\bǎl\xa7PP\xff\xfb\xdaNB\xb3Э\xc4)\xf1<\xdb\xf3\xe6\xbd\xf1T\xffI#\xfc\xbeG\xd8\xfaN=f\xd5\xf4\xf1\xe4\x15>\xbe\xbd\xc1\xaa\xe7-\xfeNK8\x1c\xaa\xeb\x11\xc8*\xe7" +
			"\xf7\xf1\x9b\xd5F\xee\xe1-\x03蹔\xa4\xdbܛ\x9e\x157\xff?,B\n\x1b?\xc7\x1a\xa3}\xde\xf0\x8eԞ\xc1\xc5\x0fT;\xf4$8\xfc\xa4v\xeb\xe1\x17\x0exq\x05\xef\xe1+\xf8n\x89\xab+p\\\xbbܡ\xa5&^\xa2Hc\xbe\xc5x\x84A\xb1\xba\xdb`\xf7\x90\x1d\xb2\xac1\xb6Kd$\xb9^\xf1=\xab\x95\x11O\xf1" +
			"Dm\xacD;2\x01i\xbcG\tE\xff\x02-'\xed\x02h>\xf1\x85\xdb#_G\xafȊ\xd5:\xa6\x01x&\xe9\xb7\f\xcaXPHJ\xba\x1f|\xca:\x01\xc5\xcdX\xeaL\xb0\\\xf7/q\xddqے\xce\xeb\x90\xddtl3\x06\xa7\x9c\f\xbe\x056\xc5\xed\x18<&\x85b:\xfc\xa1\xe4\xf5\xed\xfd\xe6\xaeX\x97\x11\x12F\x19\xcb" +
			"\xe0r\xb3٤R\xb9xj\xad\x19\xb4\xccg\xa4i\x9a\x0f\bu\xc1V\x06\xdah<\xcaÒ$\xce(\x92p)\x84X\bg\xb9\xa4\xc11Hl\xe6\xa2W\xc2\xf3e\xe1|\xf0\xe6\x9d\xfe\xf3ȵ6J.I\x96ey\x9e\xe4l\x1d\x90N5\xbf[xR\xa9\xbf5\x85\x9b\xaf\xab\xe7\xf1\xc5\xe7\\Q\xab\x19\b\xd4\x1em\xf2xK" +
			"\x1es\xd7s\x91Tz\xb6\xbc\x8f\xe1\x1d\xdaؖj>Б\x94*)\x98w.\xb4\xfe \xb69\x17\x9eLĸ\xa6~P<\xaeR\xa2\xb3\xa8\x18\xac\x8b2\xf4\x86f\x0eA\xc2\xfa\x89|>\x84\xbe\x0f\xbd\xafP\xf8\xa3cyg^\xff\x85\xb8\xd3\xc0\xc9\xe0\x97\x9d\x0f\x9a\xe8\xd3ol\xf1T\u058b>\x9e\x1aA\x87\xb7\xc9U\xbaB\xf1" +
			"\x1aU\xbac\xb2/Ό\xe9=\x9c\xbb\xb5\xfc|\xeb\xdc^K\x83\x8b\xfbq\xdf\a\xb3\xc6\x1eI\x04\xc2\\\xe1lG.x,\xe3o4e\x87\x89\xd1؍\xb5\x1a0\ueb2e\xe79W\xc5A\xf7\x186ę\x98\xe6\xcc!\xc2c\xb4\xba\x1e\xc7\xe6\x9f\x00\x00\x00\xff\xff\x04\xcaChN\x05\x00\x00",
	},

	"/sign_in_form.html": {
		local: "templates/sign_in_form.html",
		size:  518,
		compressed: "\x1f\x8b\b\x00\x00\tn\x88\x00\xffl\x91\xcdn\xc20\x10\x84\xef<\xc5\xca犼\x80\xe3\xde*\xf5T$\xda3\xda$\x1bl\xd5\u007f\xf2O\vB\xbc{\xed@@\x81ޒY\xcf\xe7\x995\x1f]0\x80}Rζ\xacɑB\x839I\x06\x86\x92tC\xcb6\x1f\xdbO&V\x00<z\xb4b\xa3\t#AT{\vʂ\v\xc0" +
			"\x11d\xa0qv\xd7\xc9.{&\xa6#\xd9\xf3\x06Ś7\x93\xb9R\x94\xf59A:zjY\xa2Cb`\xd1\xd4o\xf7M\x96\x81\xd7ؓtz\xa0в\xaf\x02\xac\xd3\x17 \x83J\xd7ۼt\x96\xc0f\xd3Q`O@\x8f1\xfe\xba0\xcc\xd0\xfb\xff\x82\xbb\x99\xe5\t\xa0\xb1#-\x1ej\x94\xbd\xec]\xda\xdd\x00\xe2m\x12\xe0\xe8" +
			"r\tqU_k9\xde\\\x00\x8fY\xa4\x1a\x86\xda\xe8\x92\xe4\x10ø\xbbv\xfcA\x9d\x8bt:\xc1\xfa.\xc3\xf9\xfc\xdcg\xc9\xf0\x18\xd0,\xec\x93\xf2\xaf3\xe6Ψ\xb2\xdc^\x97\xac-\xeb\x13ތ\xdb\xfa0\xef\xb6xx\xadi\xc4\xea/\x00\x00\xff\xff\x8c\x9b\x11l\x06\x02\x00\x00",
	},

	"/sign_up_form.html": {
		local: "templates/sign_up_form.html",
		size:  613,
		compressed: "\x1f\x8b\b\x00\x00\tn\x88\x00\xff\x8c\x92In\xe30\x10E\xf7>E\xa1\x0e`]\x80\"\xba\x17\xbd\xe9M\vp{m\x94\xc5rD\x84\x138$\x1e\u0ec7\x92\x15\x1bv\x12$+I\xc5\xff\xfe\x13J\x12;\x1f-P\x9f\xb5w-6%ql\xfaȔ\x19\xc1r\x1e\xbcj\xb1\xfb\xb7\xfa\x8fr\x01 R ';Ô\x18\x92~" +
			"rP\x02\xf8\b\x82`\x88\xbc{\xe7Ǔ\x8dv(\xa7\x88v\xa2!\xb9\x14\xcd\x04\x8f-\x86\xb6l\xa0\x8a[\xcc\xfe\x99kp]1G\x96E3\x9dM)\xedBɐ\x0f\x81k\x8c\xf7\x19aL\xb4X\xe6,B0\xd4\xf3\xe0\x8d\xe2\xda\xf4\xd7\x0f\xae\xd3\xc7#ᣃ-i\x83\xf2\xcfx\x81\xdfJEN\xe9'\xa2\vwo9\xf8" +
			"\xf2\x8b\xf7d\x83\xe1e\xef\xed\aW\xa0\x94^}T(\xbb\xf9\xee+\xd359\xdbn\xcfw\xc2\xeeZ\xf8X0h\xa5\xea\xeef|\x9f\xe2ns\xd9&\xbc\x90)ut:\xc1\xf26\x86\xf3\xf9\xbb\x8e@\x91\xec\x1d>M>%S\xd9Z]W՛\xfa\x82-\xf6\x99\xae\xe0j\xfc\xea\xebP\x19ь\u007f\x97\\\xbc\x05\x00\x00\xff\xffV" +
			"\x94)oe\x02\x00\x00",
	},

	"/": {
		isDir: true,
		local: "templates",
	},
}
