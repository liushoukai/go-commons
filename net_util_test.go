package commons

import (
	"fmt"
	"testing"
)

func TestAllConversions(t *testing.T) {
	serverIp := GetLocalIP()
	fmt.Println(serverIp)
}
