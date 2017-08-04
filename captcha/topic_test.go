package captcha

import (
	"captcha-zh/config"

	"strconv"
	"testing"
	"time"
)

func TestNum2CN(t *testing.T) {
	a := Num2Cn(5)
	if a.Size != 1 || a.Cn != "五" {
		t.Log(a)
		t.Fatalf("Parse num 5 with err")
	}
	b := Num2Cn(15)
	if b.Size != 2 || b.Cn != "十 五" {
		t.Log(b)
		t.Fatalf("Parse num 15 with err")
	}
	d := Num2Cn(20)
	if d.Size != 2 || d.Cn != "二 十" {
		t.Log(d)
		t.Fatalf("Parse num 15 with err")
	}
	c := Num2Cn(25)
	if c.Size != 3 || c.Cn != "二 十 五" {
		t.Log(c)
		t.Fatalf("Parse num 25 with err")
	}
}

func TestTopicParse(t *testing.T) {
	a := TopicParse(NumSt{3, "二 十 一"}, NumSt{3, "二 十 一"}, "+", 1)
	if a != "二 十 一 加 二 十 一" {
		t.Fatalf("Topic unexpected: %s", a)
	}
	b := TopicParse(NumSt{3, "二 十 一"}, NumSt{3, "二 十 一"}, "+", 2)
	if b != "二 十 一 加 二 十 一" {
		t.Fatalf("Topic unexpected: %s", b)
	}
	c := TopicParse(NumSt{1, "12"}, NumSt{3, "二 十 一"}, "-", 2)
	if c != "12 减 掉 二 十 一 是" {
		t.Fatalf("Topic unexpected: %s", c)
	}
	d := TopicParse(NumSt{1, "12"}, NumSt{1, "12"}, "-", 1)
	if d != "12 减 12 是 多 少 呢" {
		t.Fatalf("Topic unexpected: %s", d)
	}
}

func BenchmarkTopic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tp := RandTopic()
		Draw(tp.Subject, config.PATH_ROOT + "bin/tmp/"+strconv.Itoa(time.Now().Nanosecond())+"|"+tp.Result+".gif")
	}
}
