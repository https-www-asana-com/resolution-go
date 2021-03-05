package resolution

import (
	"encoding/json"
)

// SupportedKeys struct of supported keys
type SupportedKeys map[string]struct {
	DeprecatedKeyName string
	Deprecated        bool
	ValidationRegex   string
}

// NewSupportedKeys returns SupportedKeys
func NewSupportedKeys() (SupportedKeys, error) {
	var keysObject struct {
		Keys SupportedKeys
	}
	err := json.Unmarshal(supportedKeysJSON, &keysObject)
	if err != nil {
		return nil, err
	}
	return keysObject.Keys, nil
}

var supportedKeysJSON = []byte(`
{
    "version": "1.1.0",
    "keys": {
        "crypto.BTC.address": {
            "deprecatedKeyName": "BTC",
            "deprecated": false,
            "validationRegex": "^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$"
        },
        "crypto.ETH.address": {
            "deprecatedKeyName": "ETH",
            "deprecated": false,
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.ZIL.address": {
            "deprecatedKeyName": "ZIL",
            "deprecated": false,
            "validationRegex": "^0x[a-fA-F0-9]{40}$|^zil1[qpzry9x8gf2tvdw0s3jn54khce6mua7l]{38}$"
        },
        "crypto.LTC.address": {
            "deprecatedKeyName": "LTC",
            "deprecated": false,
            "validationRegex": "^[LM3][a-km-zA-HJ-NP-Z1-9]{26,33}$"
        },
        "crypto.ETC.address": {
            "deprecatedKeyName": "ETC",
            "deprecated": false,
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.EQL.address": {
            "deprecatedKeyName": "EQL",
            "deprecated": false,
            "validationRegex": "^bnb[0-9a-z]{39}$"
        },
        "crypto.LINK.address": {
            "deprecatedKeyName": "LINK",
            "deprecated": false,
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.USDC.address": {
            "deprecatedKeyName": "USDC",
            "deprecated": false,
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.BAT.address": {
            "deprecatedKeyName": "BAT",
            "deprecated": false,
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.REP.address": {
            "deprecatedKeyName": "REP",
            "deprecated": false,
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.ZRX.address": {
            "deprecatedKeyName": "ZRX",
            "deprecated": false,
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.DAI.address": {
            "deprecatedKeyName": "DAI",
            "deprecated": false,
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.BCH.address": {
            "deprecatedKeyName": "BCH",
            "deprecated": false,
            "validationRegex": "^[13][a-km-zA-HJ-NP-Z1-9]{33}$|^((bitcoincash|bchreg|bchtest):)?(q|p)[a-z0-9]{41}$|^((BITCOINCASH:)?(Q|P)[A-Z0-9]{41})$"
        },
        "crypto.XMR.address": {
            "deprecatedKeyName": "XMR",
            "deprecated": false,
            "validationRegex": "^[48]{1}[0-9AB][1-9A-HJ-NP-Za-km-z]{93}$"
        },
        "crypto.DASH.address": {
            "deprecatedKeyName": "DASH",
            "deprecated": false,
            "validationRegex": "^X[1-9A-HJ-NP-Za-km-z]{33}$"
        },
        "crypto.NEO.address": {
            "deprecatedKeyName": "NEO",
            "deprecated": false,
            "validationRegex": "^A[0-9a-zA-Z]{33}$"
        },
        "crypto.SWTH.address": {
            "deprecatedKeyName": "SWTH",
            "deprecated": false,
            "validationRegex": "^A[0-9a-zA-Z]{33}$"
        },
        "crypto.DOGE.address": {
            "deprecatedKeyName": "DOGE",
            "deprecated": false,
            "validationRegex": "^D[5-9A-HJ-NP-U]{1}[1-9A-HJ-NP-Za-km-z]{32}$"
        },
        "crypto.XRP.address": {
            "deprecatedKeyName": "XRP",
            "deprecated": false,
            "validationRegex": "^r[1-9a-km-zA-HJ-NP-Z]{24,34}$"
        },
        "crypto.ZEC.address": {
            "deprecatedKeyName": "ZEC",
            "deprecated": false,
            "validationRegex": "^z([a-zA-Z0-9]){94}$|^zs([a-zA-Z0-9]){75}$|^t([a-zA-Z0-9]){34}$"
        },
        "crypto.ADA.address": {
            "deprecatedKeyName": "ADA",
            "deprecated": false,
            "validationRegex": "^[1-9a-km-zA-HJ-NP-Z]{104}$|^A[1-9A-HJ-NP-Za-km-z]{58}$|^addr[0-9a-zA-Z]{99}$"
        },
        "crypto.EOS.address": {
            "deprecatedKeyName": "EOS",
            "deprecated": false,
            "validationRegex": "^[a-z][a-z1-5.]{10}[a-z1-5]$"
        },
        "crypto.XLM.address": {
            "deprecatedKeyName": "XLM",
            "deprecated": false,
            "validationRegex": "^G[A-Z2-7]{55}$"
        },
        "crypto.BNB.address": {
            "deprecatedKeyName": "BNB",
            "deprecated": false,
            "validationRegex": "^bnb[0-9a-z]{39}$"
        },
        "crypto.BTG.address": {
            "deprecatedKeyName": "BTG",
            "deprecated": false,
            "validationRegex": "^[GA][a-km-zA-HJ-NP-Z1-9]{33}$"
        },
        "crypto.NANO.address": {
            "deprecatedKeyName": "NANO",
            "deprecated": false,
            "validationRegex": "^nano_[1-9a-z]{60}$"
        },
        "crypto.WAVES.address": {
            "deprecatedKeyName": "WAVES",
            "deprecated": false,
            "validationRegex": "^3[a-km-zA-HJ-NP-Z1-9]{34}$"
        },
        "crypto.KMD.address": {
            "deprecatedKeyName": "KMD",
            "deprecated": false,
            "validationRegex": "^R[a-km-zA-Z1-9]{33}$"
        },
        "crypto.AE.address": {
            "deprecatedKeyName": "AE",
            "deprecated": false,
            "validationRegex": "^ak_[a-km-zA-Z1-9]{48,52}$"
        },
        "crypto.RSK.address": {
            "deprecatedKeyName": "RSK",
            "deprecated": false,
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.WAN.address": {
            "deprecatedKeyName": "WAN",
            "deprecated": false,
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.STRAT.address": {
            "deprecatedKeyName": "STRAT",
            "deprecated": false,
            "validationRegex": "^S[a-km-zA-HJ-NP-Z1-9]{33}$"
        },
        "crypto.UBQ.address": {
            "deprecatedKeyName": "UBQ",
            "deprecated": false,
            "validationRegex": "^0x[a-km-zA-HJ-NP-Z0-9]{40}$"
        },
        "crypto.XTZ.address": {
            "deprecatedKeyName": "XTZ",
            "deprecated": false,
            "validationRegex": "^(tz|KT)[a-km-zA-HJ-NP-Z1-9]{34}$"
        },
        "crypto.IOTA.address": {
            "deprecatedKeyName": "IOTA",
            "deprecated": false,
            "validationRegex": "^[A-Z0-9]{90}$"
        },
        "crypto.VET.address": {
            "deprecatedKeyName": "VET",
            "deprecated": false,
            "validationRegex": "^0x[a-km-zA-HJ-NP-Z0-9]{40}$"
        },
        "crypto.QTUM.address": {
            "deprecatedKeyName": "QTUM",
            "deprecated": false,
            "validationRegex": "^Q[a-km-zA-HJ-NP-Z1-9]{33}$"
        },
        "crypto.ICX.address": {
            "deprecatedKeyName": "ICX",
            "deprecated": false,
            "validationRegex": "^[a-km-zA-HJ-NP-Z0-9]{42}$"
        },
        "crypto.DGB.address": {
            "deprecatedKeyName": "DGB",
            "deprecated": false,
            "validationRegex": "^[a-km-zA-HJ-NP-Z1-9]{34}$|^[a-zA-Z1-9]{42}$"
        },
        "crypto.XZC.address": {
            "deprecatedKeyName": "XZC",
            "deprecated": false,
            "validationRegex": "^[a-km-zA-HJ-NP-Z1-9]{34}$"
        },
        "crypto.BURST.address": {
            "deprecatedKeyName": "BURST",
            "deprecated": false,
            "validationRegex": "^BURST-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{5}"
        },
        "crypto.DCR.address": {
            "deprecatedKeyName": "DCR",
            "deprecated": false,
            "validationRegex": null
        },
        "crypto.XEM.address": {
            "deprecatedKeyName": "XEM",
            "deprecated": false,
            "validationRegex": "^N[ABCDEFGHIJKLMNOPQRSTUVWXYZ234567]{39}$"
        },
        "crypto.LSK.address": {
            "deprecatedKeyName": "LSK",
            "deprecated": false,
            "validationRegex": "^\\d{1,21}[L]$"
        },
        "crypto.ATOM.address": {
            "deprecatedKeyName": "ATOM",
            "deprecated": false,
            "validationRegex": "^(cosmos)1([qpzry9x8gf2tvdw0s3jn54khce6mua7l]+)$"
        },
        "crypto.ONG.address": {
            "deprecatedKeyName": "ONG",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.ONT.address": {
            "deprecatedKeyName": "ONT",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.SMART.address": {
            "deprecatedKeyName": "SMART",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.TPAY.address": {
            "deprecatedKeyName": "TPAY",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.GRS.address": {
            "deprecatedKeyName": "GRS",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.BSV.address": {
            "deprecatedKeyName": "BSV",
            "deprecated": false,
            "validationRegex": "^bitcoincash:[a-zA-Z0-9]{42}$"
        },
        "crypto.GAS.address": {
            "deprecatedKeyName": "GAS",
            "deprecated": false,
            "validationRegex": null
        },
        "crypto.TRX.address": {
            "deprecatedKeyName": "TRX",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.VTHO.address": {
            "deprecatedKeyName": "VTHO",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{42}$"
        },
        "crypto.BCD.address": {
            "deprecatedKeyName": "BCD",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.BTT.address": {
            "deprecatedKeyName": "BTT",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.KIN.address": {
            "deprecatedKeyName": "KIN",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{56}$"
        },
        "crypto.RVN.address": {
            "deprecatedKeyName": "RVN",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.ARK.address": {
            "deprecatedKeyName": "ARK",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.XVG.address": {
            "deprecatedKeyName": "XVG",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.ALGO.address": {
            "deprecatedKeyName": "ALGO",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{58}$"
        },
        "crypto.NEBL.address": {
            "deprecatedKeyName": "NEBL",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.XPM.address": {
            "deprecatedKeyName": "XPM",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.ONE.address": {
            "deprecatedKeyName": "ONE",
            "deprecated": false,
            "validationRegex": "^one[a-zA-Z0-9]{39}$"
        },
        "crypto.BNTY.address": {
            "deprecatedKeyName": "BNTY",
            "deprecated": false,
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.CRO.address": {
            "deprecatedKeyName": "CRO",
            "deprecated": false,
            "validationRegex": "^0x[a-fA-F0-9]{40}$"
        },
        "crypto.TWT.address": {
            "deprecatedKeyName": "TWT",
            "deprecated": false,
            "validationRegex": "^bnb[0-9a-z]{39}$"
        },
        "crypto.SIERRA.address": {
            "deprecatedKeyName": "SIERRA",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{34}$"
        },
        "crypto.VSYS.address": {
            "deprecatedKeyName": "VSYS",
            "deprecated": false,
            "validationRegex": "^[a-zA-Z0-9]{35}$"
        },
        "crypto.HIVE.address": {
            "deprecatedKeyName": "HIVE",
            "validationRegex": "^(?!s*$).+",
            "deprecated": false
        },
        "crypto.HT.address": {
            "deprecatedKeyName": "HT",
            "validationRegex": "^0x[a-fA-F0-9]{40}$",
            "deprecated": false
        },
        "crypto.ENJ.address": {
            "deprecatedKeyName": "ENJ",
            "validationRegex": "^0x[a-fA-F0-9]{40}$",
            "deprecated": false
        },
        "crypto.YFI.address": {
            "deprecatedKeyName": "YFI",
            "validationRegex": "^0x[a-fA-F0-9]{40}$",
            "deprecated": false
        },
        "crypto.MTA.address": {
            "deprecatedKeyName": "MTA",
            "validationRegex": "^0x[a-fA-F0-9]{40}$",
            "deprecated": false
        },
        "crypto.COMP.address": {
            "deprecatedKeyName": "COMP",
            "validationRegex": "^0x[a-fA-F0-9]{40}$",
            "deprecated": false
        },
        "crypto.BAL.address": {
            "deprecatedKeyName": "BAL",
            "validationRegex": "^0x[a-fA-F0-9]{40}$",
            "deprecated": false
        },
        "crypto.AMPL.address": {
            "deprecatedKeyName": "AMPL",
            "validationRegex": "^0x[a-fA-F0-9]{40}$",
            "deprecated": false
        },
        "crypto.LEND.address": {
            "deprecatedKeyName": "LEND",
            "validationRegex": "^0x[a-fA-F0-9]{40}$",
            "deprecated": false
        },
        "crypto.TLOS.address": {
            "deprecatedKeyName": "TLOS",
            "validationRegex": "^[a-z][a-z1-5.]{10}[a-z1-5]$",
            "deprecated": false
        },
        "crypto.USDT.version.ERC20.address": {
            "deprecatedKeyName": "USDT_ERC20",
            "validationRegex": "^0x[a-fA-F0-9]{40}$",
            "deprecated": false
        },
        "crypto.USDT.version.TRON.address": {
            "deprecatedKeyName": "USDT_TRON",
            "validationRegex": "^[T][a-zA-HJ-NP-Z0-9]{33}$",
            "deprecated": false
        },
        "crypto.USDT.version.EOS.address": {
            "deprecatedKeyName": "USDT_EOS",
            "validationRegex": "^[a-z][a-z1-5.]{10}[a-z1-5]$",
            "deprecated": false
        },
        "crypto.USDT.version.OMNI.address": {
            "deprecatedKeyName": "USDT_OMNI",
            "validationRegex": "^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$",
            "deprecated": false
        },
        "social.payid.name": {
            "deprecatedKeyName": "payid",
            "validationRegex": "^[0-9a-zA-Z]+\\$[0-9a-zA-Z]+\\.[0-9a-zA-Z]+$",
            "deprecated": false
        },
        "whois.email.value": {
            "deprecatedKeyName": "email",
            "validationRegex": "^[^@]+@[^\\.]+\\..+$",
            "deprecated": false
        },
        "whois.for_sale.value": {
            "deprecatedKeyName": "for_sale",
            "validationRegex": "(true)|(false)",
            "deprecated": false
        },
        "ipfs.html.value": {
            "deprecatedKeyName": "html",
            "validationRegex": ".{0,100}",
            "deprecated": false
        },
        "ipfs.redirect_domain.value": {
            "deprecatedKeyName": "redirect_domain",
            "validationRegex": ".{0,253}",
            "deprecated": false
        },
        "gundb.username.value": {
            "deprecatedKeyName": "gundb_username",
            "validationRegex": null,
            "deprecated": false
        },
        "gundb.public_key.value": {
            "deprecatedKeyName": "gundb_public_key",
            "validationRegex": null,
            "deprecated": false
        },
        "social.image.value": {
            "deprecatedKeyName": "image",
            "validationRegex": null,
            "deprecated": false
        },
        "social.twitter.username": {
            "deprecatedKeyName": "twitter_username",
            "validationRegex": null,
            "deprecated": false
        },
        "validation.social.twitter.username": {
            "deprecatedKeyName": "validation_twitter_username",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.ttl": {
            "deprecatedKeyName": "ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.A": {
            "deprecatedKeyName": "A",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.A.ttl": {
            "deprecatedKeyName": "A_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.AAAA": {
            "deprecatedKeyName": "AAAA",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.AAAA.ttl": {
            "deprecatedKeyName": "AAAA_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.AFSDB": {
            "deprecatedKeyName": "AFSDB",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.AFSDB.ttl": {
            "deprecatedKeyName": "AFSDB_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.APL": {
            "deprecatedKeyName": "APL",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.APL.ttl": {
            "deprecatedKeyName": "APL_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.CAA": {
            "deprecatedKeyName": "CAA",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.CAA.ttl": {
            "deprecatedKeyName": "CAA_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.CDNSKEY": {
            "deprecatedKeyName": "CDNSKEY",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.CDNSKEY.ttl": {
            "deprecatedKeyName": "CDNSKEY_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.CDS": {
            "deprecatedKeyName": "CDS",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.CDS.ttl": {
            "deprecatedKeyName": "CDS_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.CERT": {
            "deprecatedKeyName": "CERT",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.CERT.ttl": {
            "deprecatedKeyName": "CERT_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.CNAME": {
            "deprecatedKeyName": "CNAME",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.CNAME.ttl": {
            "deprecatedKeyName": "CNAME_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.CSYNC": {
            "deprecatedKeyName": "CSYNC",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.CSYNC.ttl": {
            "deprecatedKeyName": "CSYNC_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.DHCID": {
            "deprecatedKeyName": "DHCID",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.DHCID.ttl": {
            "deprecatedKeyName": "DHCID_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.DLV": {
            "deprecatedKeyName": "DLV",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.DLV.ttl": {
            "deprecatedKeyName": "DLV_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.DNAME": {
            "deprecatedKeyName": "DNAME",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.DNAME.ttl": {
            "deprecatedKeyName": "DNAME_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.DNSKEY": {
            "deprecatedKeyName": "DNSKEY",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.DNSKEY.ttl": {
            "deprecatedKeyName": "DNSKEY_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.DS": {
            "deprecatedKeyName": "DS",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.DS.ttl": {
            "deprecatedKeyName": "DS_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.EUI48": {
            "deprecatedKeyName": "EUI48",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.EUI48.ttl": {
            "deprecatedKeyName": "EUI48_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.EUI64": {
            "deprecatedKeyName": "EUI64",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.EUI64.ttl": {
            "deprecatedKeyName": "EUI64_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.HINFO": {
            "deprecatedKeyName": "HINFO",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.HINFO.ttl": {
            "deprecatedKeyName": "HINFO_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.HIP": {
            "deprecatedKeyName": "HIP",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.HIP.ttl": {
            "deprecatedKeyName": "HIP_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.HTTPS": {
            "deprecatedKeyName": "HTTPS",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.HTTPS.ttl": {
            "deprecatedKeyName": "HTTPS_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.IPSECKEY": {
            "deprecatedKeyName": "IPSECKEY",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.IPSECKEY.ttl": {
            "deprecatedKeyName": "IPSECKEY_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.KEY": {
            "deprecatedKeyName": "KEY",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.KEY.ttl": {
            "deprecatedKeyName": "KEY_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.KX": {
            "deprecatedKeyName": "KX",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.KX.ttl": {
            "deprecatedKeyName": "KX_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.LOC": {
            "deprecatedKeyName": "LOC",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.LOC.ttl": {
            "deprecatedKeyName": "LOC_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.MX": {
            "deprecatedKeyName": "MX",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.MX.ttl": {
            "deprecatedKeyName": "MX_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.NAPTR": {
            "deprecatedKeyName": "NAPTR",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.NAPTR.ttl": {
            "deprecatedKeyName": "NAPTR_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.NS": {
            "deprecatedKeyName": "NS",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.NS.ttl": {
            "deprecatedKeyName": "NS_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.NSEC": {
            "deprecatedKeyName": "NSEC",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.NSEC.ttl": {
            "deprecatedKeyName": "NSEC_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.NSEC3": {
            "deprecatedKeyName": "NSEC3",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.NSEC3.ttl": {
            "deprecatedKeyName": "NSEC3_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.NSEC3PARAM": {
            "deprecatedKeyName": "NSEC3PARAM",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.NSEC3PARAM.ttl": {
            "deprecatedKeyName": "NSEC3PARAM_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.OPENPGPKEY": {
            "deprecatedKeyName": "OPENPGPKEY",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.OPENPGPKEY.ttl": {
            "deprecatedKeyName": "OPENPGPKEY_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.PTR": {
            "deprecatedKeyName": "PTR",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.PTR.ttl": {
            "deprecatedKeyName": "PTR_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.RP": {
            "deprecatedKeyName": "RP",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.RP.ttl": {
            "deprecatedKeyName": "RP_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.RRSIG": {
            "deprecatedKeyName": "RRSIG",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.RRSIG.ttl": {
            "deprecatedKeyName": "RRSIG_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.SIG": {
            "deprecatedKeyName": "SIG",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.SIG.ttl": {
            "deprecatedKeyName": "SIG_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.SMIMEA": {
            "deprecatedKeyName": "SMIMEA",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.SMIMEA.ttl": {
            "deprecatedKeyName": "SMIMEA_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.SOA": {
            "deprecatedKeyName": "SOA",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.SOA.ttl": {
            "deprecatedKeyName": "SOA_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.SRV": {
            "deprecatedKeyName": "SRV",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.SRV.ttl": {
            "deprecatedKeyName": "SRV_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.SSHFP": {
            "deprecatedKeyName": "SSHFP",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.SSHFP.ttl": {
            "deprecatedKeyName": "SSHFP_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.SVCB": {
            "deprecatedKeyName": "SVCB",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.SVCB.ttl": {
            "deprecatedKeyName": "SVCB_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.TA": {
            "deprecatedKeyName": "TA",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.TA.ttl": {
            "deprecatedKeyName": "TA_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.TKEY": {
            "deprecatedKeyName": "TKEY",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.TKEY.ttl": {
            "deprecatedKeyName": "TKEY_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.TLSA": {
            "deprecatedKeyName": "TLSA",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.TLSA.ttl": {
            "deprecatedKeyName": "TLSA_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.TSIG": {
            "deprecatedKeyName": "TSIG",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.TSIG.ttl": {
            "deprecatedKeyName": "TSIG_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.TXT": {
            "deprecatedKeyName": "TXT",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.TXT.ttl": {
            "deprecatedKeyName": "TXT_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.URI": {
            "deprecatedKeyName": "URI",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.URI.ttl": {
            "deprecatedKeyName": "URI_ttl",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.ZONEMD": {
            "deprecatedKeyName": "ZONEMD",
            "validationRegex": null,
            "deprecated": false
        },
        "dns.ZONEMD.ttl": {
            "deprecatedKeyName": "ZONEMD_ttl",
            "validationRegex": null,
            "deprecated": false
        }
    }
}
`)
