module github.com/iavian/quic-send

go 1.14

//github.com/lucas-clemente/quic-go v0.7.1-0.20200224065754-ebe3c1cca40a - very bad
//github.com/lucas-clemente/quic-go v0.7.1-0.20200224065754-24b840f56d29 - ok

//replace github.com/lucas-clemente/quic-go => ../quic-go

require github.com/lucas-clemente/quic-go v0.7.1-0.20200224065754-ebe3c1cca40a
