/*
Exercise: Stringers
Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
*/

package main

import "fmt"

type ipAddr struct {
	blockOne, blockTwo, blockThree, blockFour int
}

// IPAddress-ify the number
func (ip ipAddr) stringer() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip.blockOne, ip.blockTwo, ip.blockThree, ip.blockFour)
}

func main() {
	myIP := ipAddr{1, 2, 3, 4}

	hosts := map[string]ipAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	fmt.Printf("My IP address is %v [better printed as %v]\n", myIP, myIP.stringer())

	for name, ip := range hosts {
		fmt.Printf("%s >> %v\n", name, ip.stringer())
	}
}
