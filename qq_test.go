package qq

import "testing"

func TestNew(t *testing.T) {
    qq := New("/some/path", "", LstdFlags)

    if qq.path != "/some/path" {
        t.Error("Expected \"/some/path\", got ", qq.path)
    }
    if qq.prefix != "" {
        t.Error("Expected \"\", got ", qq.prefix)
    }
    if qq.flag != LstdFlags {
        t.Error("Expected \"\", got ", qq.flag)
    }
}

func TestFlags(t *testing.T) {
    qq := New("/some/path", "", LstdFlags)

    if qq.Flags() != LstdFlags {
        t.Error(
            "expected", LstdFlags,
            "got", qq.Flags(),
        )
    }
}
