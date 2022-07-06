package util

import (
	"github.com/duke-git/lancet/slice"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

var (
	// 可以去掉一些容易混淆的字符，比如q o k
	chars = []rune{
		'4', '1', '7', '2', '3', '6', '5', '8', '9',
		'q', 'n', 'c', 'r', 't', 'm', 'a', 'd', 'e', 'h', 'j', 'l', 'f', 'b', 'g',
		'Q', 'N', 'C', 'R', 'T', 'M', 'A', 'D', 'E', 'H', 'J', 'L', 'F', 'B', 'G',
	}
	divider      = 'i' // 分割标识(区分补位，应该是chars里面没出现的字符)
	charLen      = uint64(len(chars))
	charIndexMap = make(map[rune]int, len(chars))
)

func init() {
	for i, char := range chars {
		charIndexMap[char] = i
	}
}

// Id2Code 把id转为兑换码
func Id2Code(id uint64, codeMixLength int, isRandomFix bool) string {
	code := make([]rune, 0, codeMixLength)
	for id/charLen > 0 {
		code = append(code, chars[id%charLen])
		id /= charLen
	}
	code = append(code, chars[id%charLen]) // 处理未除尽的余数
	slice.ReverseSlice(code)
	fixLen := codeMixLength - len(code) // 需要补码的长度
	if fixLen > 0 {
		rand.Seed(time.Now().UnixNano())
		code = append(code, divider)
		for i := 0; i < fixLen-1; i++ {
			// 每次固定，如果需要变的话，后面补码的内容可以改变
			if isRandomFix {
				code = append(code, chars[rand.Intn(int(charLen))])
			} else {
				code = append(code, chars[i])
			}
		}
	}
	return string(code)
}

func Code2Id(code string) (uint64, error) {
	if len(code) == 0 {
		return 0, nil
	}
	var id uint64 = 0
	codeRuneList := []rune(code)
	for i := range codeRuneList {
		// 如果是补码标志直接退出
		if codeRuneList[i] == divider {
			break
		}
		charIndex, ok := charIndexMap[codeRuneList[i]]
		if !ok {
			return 0, errors.New("code有误，解码失败")
		}
		if i > 0 {
			id = id*charLen + uint64(charIndex)
		} else {
			id = uint64(charIndex)
		}
	}
	return id, nil
}
