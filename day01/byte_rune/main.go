package main

import "fmt"

func main() {
	s := "Hello沙河"
	//len()求byte字节数量
	n := len(s)
	fmt.Println(n)

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		fmt.Printf("%c\n", s[i]) //%c:字符,中文会乱码
	}

	for _, c := range s { //从字符串中拿出具体字符
		fmt.Printf("%c\n", c)
	}

	//字符修改
	s2 := "白萝卜"
	s3 := []rune(s2) // 把字符串强制转换成了一个rune切片['白' '萝' '卜']
	s3[0] = '红'
	fmt.Println(string(s3)) // 把rune切片强制转换成字符串

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T c2:%T\n", c1, c2) //c1:string c2:int32

	c3 := "H"
	c4 := byte('H')
	fmt.Printf("c3:%T c4:%T\n", c3, c4) //c3:string c4:uint8
	fmt.Printf("%d\n", c4)

}
