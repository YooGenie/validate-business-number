# validate-business-number
사업자 유효성 검사하는 모듈

# Installation
```
go get -u github.com/YooGenie/validate-business-number
```

# Description

* 사업자 번호 입력만 해서 유효성 검사(bisNo 숫자 10자리 입력)
```
func BusinessNumber(bisNo string) bool {
    return bool
	}
```
사업자 번호넣어서 호출하면 false, true 값으로 옵니다

# 사용법
```
validate.BusinessNumber("1234567890")

```

# UPDATE
2022-01-25 v0.0.2 #하이픈 제거하기
2022-02-17 v0.0.3 #유효한 사업자 번호 끝이 0인 경우 true 나오게 한다
