package mfa

import (
	"bufio"
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	flagAdd  = flag.Bool("add", false, "add a key")
	flagList = flag.Bool("list", false, "list keys")
	flagHotp = flag.Bool("hotp", false, "add key as HOTP (counter-based) key")
	flag7    = flag.Bool("7", false, "generate 7-digit code")
	flag8    = flag.Bool("8", false, "generate 8-digit code")
	flagClip = flag.Bool("clip", false, "copy code to the clipboard")
)

type Keychain struct {
	file string
	data []byte
	keys map[string]Key
}

type Key struct {
	raw    []byte
	digits int
	offset int // offset of counter
}

const counterLen = 20

func ReadKeychain(file string) *Keychain {
	c := &Keychain{
		file: file,
		keys: make(map[string]Key),
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return c
		}
		log.Fatal(err)
	}
	c.data = data

	lines := bytes.SplitAfter(data, []byte("\n"))
	offset := 0
	for i, line := range lines {
		lineno := i + 1
		offset += len(line)
		f := bytes.Split(bytes.TrimSuffix(line, []byte("\n")), []byte(" "))
		if len(f) == 1 && len(f[0]) == 0 {
			continue
		}
		if len(f) >= 3 && len(f[1]) == 1 && '6' <= f[1][0] && f[1][0] <= '8' {
			var k Key
			name := string(f[0])
			k.digits = int(f[1][0] - '0')
			raw, err := decodeKey(string(f[2]))
			if err == nil {
				k.raw = raw
				if len(f) == 3 {
					c.keys[name] = k
					continue
				}
				if len(f) == 4 && len(f[3]) == counterLen {
					_, err := strconv.ParseUint(string(f[3]), 10, 64)
					if err == nil {
						// Valid counter.
						k.offset = offset - counterLen
						if line[len(line)-1] == '\n' {
							k.offset--
						}
						c.keys[name] = k
						continue
					}
				}
			}
		}
		log.Printf("%s:%d: malformed key", c.file, lineno)
	}
	return c
}

func (c *Keychain) List() (names []string) {
	for name := range c.keys {
		names = append(names, name)
	}
	sort.Strings(names)
	return
}

func noSpace(r rune) rune {
	if unicode.IsSpace(r) {
		return -1
	}
	return r
}

func (c *Keychain) add(name string) {
	size := 6
	if *flag7 {
		size = 7
		if *flag8 {
			log.Fatalf("cannot use -7 and -8 together")
		}
	} else if *flag8 {
		size = 8
	}

	fmt.Fprintf(os.Stderr, "2fa key for %s: ", name)
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatalf("error reading key: %v", err)
	}
	text = strings.Map(noSpace, text)
	text += strings.Repeat("=", -len(text)&7) // pad to 8 bytes
	if _, err := decodeKey(text); err != nil {
		log.Fatalf("invalid key: %v", err)
	}

	line := fmt.Sprintf("%s %d %s", name, size, text)
	if *flagHotp {
		line += " " + strings.Repeat("0", 20)
	}
	line += "\n"

	f, err := os.OpenFile(c.file, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("opening keychain: %v", err)
	}
	f.Chmod(0600)

	if _, err := f.Write([]byte(line)); err != nil {
		log.Fatalf("adding key: %v", err)
	}
	if err := f.Close(); err != nil {
		log.Fatalf("adding key: %v", err)
	}
}

func (c *Keychain) Code(name string) string {
	k, ok := c.keys[name]
	if !ok {
		log.Fatalf("no such key %q", name)
	}
	var code int
	if k.offset != 0 {
		n, err := strconv.ParseUint(string(c.data[k.offset:k.offset+counterLen]), 10, 64)
		if err != nil {
			log.Fatalf("malformed key counter for %q (%q)", name, c.data[k.offset:k.offset+counterLen])
		}
		n++
		code = hotp(k.raw, n, k.digits)
		f, err := os.OpenFile(c.file, os.O_RDWR, 0600)
		if err != nil {
			log.Fatalf("opening keychain: %v", err)
		}
		if _, err := f.WriteAt([]byte(fmt.Sprintf("%0*d", counterLen, n)), int64(k.offset)); err != nil {
			log.Fatalf("updating keychain: %v", err)
		}
		if err := f.Close(); err != nil {
			log.Fatalf("updating keychain: %v", err)
		}
	} else {
		// Time-based key.
		code = totp(k.raw, time.Now(), k.digits)
	}
	return fmt.Sprintf("%0*d", k.digits, code)
}

func (c *Keychain) ShowAll() (data map[string]string) {
	data = make(map[string]string)
	var names []string
	max := 0
	for name, k := range c.keys {
		names = append(names, name)
		if max < k.digits {
			max = k.digits
		}
	}
	sort.Strings(names)
	for _, name := range names {
		k := c.keys[name]
		code := strings.Repeat("-", k.digits)
		if k.offset == 0 {
			code = c.Code(name)
		}
		data[name] = code
	}
	return
}

func decodeKey(key string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(strings.ToUpper(key))
}

func hotp(key []byte, counter uint64, digits int) int {
	h := hmac.New(sha1.New, key)
	binary.Write(h, binary.BigEndian, counter)
	sum := h.Sum(nil)
	v := binary.BigEndian.Uint32(sum[sum[len(sum)-1]&0x0F:]) & 0x7FFFFFFF
	d := uint32(1)
	for i := 0; i < digits && i < 8; i++ {
		d *= 10
	}
	return int(v % d)
}

func totp(key []byte, t time.Time, digits int) int {
	return hotp(key, uint64(t.UnixNano())/30e9, digits)
}
