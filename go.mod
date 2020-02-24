module github.com/iavian/quic-send

go 1.13

//github.com/lucas-clemente/quic-go v0.7.1-0.20200224094014-ca469eb0b6de - very bad 20+
//github.com/lucas-clemente/quic-go v0.7.1-0.20200221100903-c1cb3e64dc00 - 3+
//github.com/lucas-clemente/quic-go v0.7.1-0.20200220095333-73937e87539f - 2.5+
//github.com/lucas-clemente/quic-go v0.7.1-0.20200220092722-244e1ae8e750 - 2+(best)
//github.com/lucas-clemente/quic-go v0.7.1-0.20200220092450-88fc6b9a8714 - 5+
//require github.com/lucas-clemente/quic-go v0.7.1-0.20200218105105-d08c2145a4d9 5+

require (
	github.com/lucas-clemente/quic-go v0.7.1-0.20200220092722-244e1ae8e750
	golang.org/x/crypto v0.0.0-20200221231518-2aa609cf4a9d // indirect
)
