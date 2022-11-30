package internal

import (
	"fmt"
	"testing"
)

func TestCipher_Enc(t *testing.T) {
	v := LoginCipher.Enc([]byte("{\"accompany\":\"1\",\"channelNo\":\"110001\",\"checkAuth\":\"1\",\"face\":\"1\",\"game\":\"wf\",\"media\":\"M311463\",\"mmid\":\"\",\"os\":\"1\",\"password\":\"7add55ab5fd2cdd80c8f93e04accef09\",\"serial\":\"FA:FA:FA:FA:FA:FA|000000000000000|unknown|fa96d544b2a70de3\",\"username\":\"18934771956\",\"versionCode\":1005006,\"versionName\":\"1.5.6\"}"))
	fmt.Println(string(v))
}

func TestCipher_Dec(t *testing.T) {
	v := LoginCipher.Dec([]byte("MSt17GTF7x6ianiYlYRcB78buzpzGkKDB28vOtar/vibcn8OVwAITX1Ny3ID59F1PlmeiI+gsNAJw8vHhT7XqksgrENGuGF5sJXk1YX5DJScmUxsC9Ez4YfcaLX0Ac12XnwSKaNSWwy4dYR32tyvEfYqQOUJoVksZ0lhU/vc8vzzsMBZ3rsVa\nZ6PYqI/LAtLlwdQkSh7btMJ1AiSwjgN5wP+jvlUmod7S6g+yY0zHmIREJIenrz7RDKEiQJSpGFqplTe1EN9odzBTOuBwDeGhYyqT+IFWRcm0k0Ai9I6XrOlbabZlvNianyZ5PwNWmHuT5rHY2+4rV9aa6ZZzHd6q6E4Xwy4FGASW3PuBycvS+YLYmlRFomPNyPb/MiIDJSOP5V0oioEP+hSAd618HOFUg=="))
	fmt.Println(string(v))
}
