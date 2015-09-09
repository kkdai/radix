package radix

import "testing"

func TeststringSubsetPrefix(t *testing.T) {
	sub, _ := stringSubsetPrefix([]byte("playground"), []byte("yield"))
	if string(sub) != "" {
		t.Errorf("None string subset failed, expect:%s but get:%s\n", "play", sub)
	}

	sub, _ = stringSubsetPrefix([]byte("playground"), []byte("playground"))
	if string(sub) != "playground" {
		t.Errorf("full subset failed, expect:%s but get:%s\n", "playground", sub)
	}

	sub, _ = stringSubsetPrefix([]byte("playground"), []byte("playboy"))
	if string(sub) != "play" {
		t.Errorf("Sub string subset failed, expect:%s but get:%s\n", "play", sub)
	}

}
