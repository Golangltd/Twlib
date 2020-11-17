
第一步
git config --global url."git@192.168.2.199:root/twlib".insteadof "https://github.com/Golangltd/Twlib"

第二步
go env -w GOPRIVATE=github.com/Golangltd/Twlib

第三步
拷贝以来mod到本地