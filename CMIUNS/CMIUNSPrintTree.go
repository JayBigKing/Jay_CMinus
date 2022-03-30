package CMIUNS

import (
	"container/list"
	"io"
	"os"
	"strconv"
)

var spaceNode Node = Node{nil, nil, 0, 0, -1, "", -1}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func PrintTree(treeFileName string) {
	//4层
	var f *os.File
	var err1 error
	if checkFileIsExist(treeFileName + ".txt") { //如果文件存在
		f, err1 = os.OpenFile(treeFileName+".txt", os.O_RDWR|os.O_TRUNC, 0666) //打开文件
	} else {
		f, err1 = os.Create(treeFileName + ".txt") //创建文件
	}
	check(err1)

	l := list.New()
	l.PushBack(&ParseTree)
	layerCount := 1
	nextLayerCount := 0
	stoneSpace := 2
	theSpace := 0
	for i := layersOfTree; i >= 0; i-- {
		layerStr := "Layer " + strconv.Itoa(layersOfTree-i+1)
		_, err1 = io.WriteString(f, layerStr) //写入文件(字符串)
		check(err1)
		//		theSpace = i*stoneSpace + 1
		theSpace = stoneSpace + 1
		nextLayerCount = 0
		for j := layerCount; j > 0; j-- {
			l1 := l.Front()
			//nowNode = (*Node)(unsafe.Pointer(&l1.Value))
			var a interface{} = l1.Value
			nowNode = a.(*Node)
			l.Remove(l1)
			for t := 0; t < len(nowNode.child); t++ {
				nextLayerCount++
				l.PushBack(nowNode.child[t])
			}

			for k := 0; k < theSpace; k++ {
				//				fmt.Printf("\t")
				_, err1 = io.WriteString(f, "\t") //写入文件(字符串)
				check(err1)
			}
			if nowNode.kind != -1 {
				if nowNode.layer != 1 {
					//nowNodeParent := (*nowNode).parent
					_, err1 = io.WriteString(f, nowNode.lexeme+"("+strconv.Itoa((*nowNode).parent.NodeNo)+","+strconv.Itoa(nowNode.NodeNo)+")") //写入文件(字符串)
					check(err1)
				} else {
					_, err1 = io.WriteString(f, nowNode.lexeme) //写入文件(字符串)
					check(err1)
				}

				check(err1)
			} else {
				//				fmt.Printf(" ")
				_, err1 = io.WriteString(f, "") //写入文件(字符串)
				check(err1)
			}
		}
		layerCount = nextLayerCount
		//		fmt.Println()
		_, err1 := io.WriteString(f, "\r\n") //写入文件(字符串)
		check(err1)
	}
	f.Close()
}

/*
func PrintTree(treeFileName string) {
	//4层
	var f *os.File
	var err1 error
	if checkFileIsExist(treeFileName + ".txt") { //如果文件存在
		f, err1 = os.OpenFile(treeFileName+".txt", os.O_RDWR|os.O_TRUNC, 0666) //打开文件
	} else {
		f, err1 = os.Create(treeFileName + ".txt") //创建文件
	}
	check(err1)

	l := list.New()
	l.PushBack(&ParseTree)
	layerCount := 1
	nextLayerCount := 0
	stoneSpace := 2
	theSpace := 0
	for i := layersOfTree; i >= 0; i-- {

		//		theSpace = i*stoneSpace + 1
		theSpace = stoneSpace + 1
		nextLayerCount = 0
		for j := layerCount; j > 0; j-- {
			l1 := l.Front()
			//nowNode = (*Node)(unsafe.Pointer(&l1.Value))
			var a interface{} = l1.Value
			nowNode = a.(*Node)
			l.Remove(l1)
			if nowNode.kind != -1 {
				for t := 0; t < len(nowNode.child); t++ {
					nextLayerCount++
					l.PushBack(nowNode.child[t])
				}
				if len(nowNode.child) < 3 {
					for t := 1; t <= 3-len(nowNode.child); t++ {
						nextLayerCount++
						l.PushBack(&spaceNode)
					}
				}
			}

			for k := 0; k < theSpace; k++ {
				//				fmt.Printf("\t")
				_, err1 := io.WriteString(f, "\t") //写入文件(字符串)
				check(err1)
			}
			if nowNode.kind != -1 {
				//				fmt.Print(nowNode.lexeme)
				_, err1 := io.WriteString(f, nowNode.lexeme) //写入文件(字符串)
				check(err1)
			} else {
				//				fmt.Printf(" ")
				_, err1 := io.WriteString(f, "") //写入文件(字符串)
				check(err1)
			}
		}
		layerCount = nextLayerCount
		//		fmt.Println()
		_, err1 := io.WriteString(f, "\r\n") //写入文件(字符串)
		check(err1)
	}
	f.Close()
}
*/
