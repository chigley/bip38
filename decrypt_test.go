package bip38

import "testing"

func TestECMultiplyNoLotSequenceCorrect(t *testing.T) {
	const key = "6PfQu77ygVyJLZjfvMLyhLMQbYnu5uguoJJ4kMCLqWwPEdfpwANVS76gTX"
	const pass = "TestingOneTwoThree"
	const out = "a43a940577f4e97f5c4d39eb14ff083a98187c64ea7c99ef7ce460833959a519"
	if x := DecryptWithPassphrase(key, pass); x != out {
		t.Errorf("DecryptWithPassphrase(%v, %v) = %v, want %v", key, pass, x, out)
	}
}

func TestECMultiplyNoLotSequenceIncorrect(t *testing.T) {
	const key = "6PfQu77ygVyJLZjfvMLyhLMQbYnu5uguoJJ4kMCLqWwPEdfpwANVS76gTX"
	const pass = "IncorrectPassphrase"
	const out = ""
	if x := DecryptWithPassphrase(key, pass); x != out {
		t.Errorf("DecryptWithPassphrase(%v, %v) = %v, want %v", key, pass, x, out)
	}
}

func TestECMultiplyLotSequenceCorrect(t *testing.T) {
	const key = "6PgNBNNzDkKdhkT6uJntUXwwzQV8Rr2tZcbkDcuC9DZRsS6AtHts4Ypo1j"
	const pass = "MOLON LABE"
	const out = "44ea95afbf138356a05ea32110dfd627232d0f2991ad221187be356f19fa8190"
	if x := DecryptWithPassphrase(key, pass); x != out {
		t.Errorf("DecryptWithPassphrase(%v, %v) = %v, want %v", key, pass, x, out)
	}
}

func TestECMultiplyLotSequenceIncorrect(t *testing.T) {
	const key = "6PgNBNNzDkKdhkT6uJntUXwwzQV8Rr2tZcbkDcuC9DZRsS6AtHts4Ypo1j"
	const pass = "MOLON_LABE"
	const out = ""
	if x := DecryptWithPassphrase(key, pass); x != out {
		t.Errorf("DecryptWithPassphrase(%v, %v) = %v, want %v", key, pass, x, out)
	}
}

func TestUTF8Correct(t *testing.T) {
	const key = "6PgGWtx25kUg8QWvwuJAgorN6k9FbE25rv5dMRwu5SKMnfpfVe5mar2ngH"
	const pass = "ΜΟΛΩΝ ΛΑΒΕ"
	const out = "ca2759aa4adb0f96c414f36abeb8db59342985be9fa50faac228c8e7d90e3006"
	if x := DecryptWithPassphrase(key, pass); x != out {
		t.Errorf("DecryptWithPassphrase(%v, %v) = %v, want %v", key, pass, x, out)
	}
}

func TestUTF8Incorrect(t *testing.T) {
	const key = "6PgGWtx25kUg8QWvwuJAgorN6k9FbE25rv5dMRwu5SKMnfpfVe5mar2ngH"
	const pass = "ΜΟΛΩΝ ΛΑΒΩ"
	const out = ""
	if x := DecryptWithPassphrase(key, pass); x != out {
		t.Errorf("DecryptWithPassphrase(%v, %v) = %v, want %v", key, pass, x, out)
	}
}
