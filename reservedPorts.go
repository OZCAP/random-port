package main

func combineMaps(maps ...map[int]bool) map[int]bool {
	result := make(map[int]bool)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

var reservedPortsMacos = map[int]bool{
	22:    true,
	25:    true,
	53:    true,
	80:    true,
	88:    true,
	110:   true,
	123:   true,
	137:   true,
	138:   true,
	139:   true,
	143:   true,
	192:   true,
	389:   true,
	443:   true,
	445:   true,
	465:   true,
	500:   true,
	515:   true,
	548:   true,
	554:   true,
	587:   true,
	631:   true,
	636:   true,
	749:   true,
	993:   true,
	995:   true,
	1900:  true,
	2195:  true,
	2196:  true,
	2197:  true,
	3031:  true,
	3283:  true,
	3284:  true,
	3285:  true,
	3478:  true,
	3497:  true,
	3689:  true,
	3690:  true,
	4398:  true,
	4500:  true,
	5100:  true,
	5223:  true,
	5228:  true,
	5297:  true,
	5350:  true,
	5351:  true,
	5353:  true,
	5900:  true,
	8000:  true,
	8999:  true,
	9100:  true,
	9418:  true,
	16384: true,
	16403: true,
	16387: true,
	16393: true,
	16402: true,
	16472: true,
	42000: true,
	42999: true,
	49152: true,
	65535: true,
}

// TODO: add reserved ports for Linux and Windows
var reservedPortsLinux = map[int]bool{}
var reservedPortsWindows = map[int]bool{}

var reservedPorts = combineMaps(reservedPortsMacos, reservedPortsLinux, reservedPortsWindows)
