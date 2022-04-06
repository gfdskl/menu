package main

import (
	"fmt"
    "unsafe"
    "os"
)

const CMD_MAX_LEN int = 128
const DESC_LEN int = 1024
const CMD_NUM int = 10

type DataNode struct{
    pNext *LinkTableNode
    cmd string
    desc string
    handler func()
}

var head *LinkTable = nil

func main() {
    InitMenuData(&head)
    for {
        var cmd string;
        fmt.Println("Input a cmd number > ")
        fmt.Scan(&cmd)
        p := FindCmd(head,cmd)
        if p == nil {
            fmt.Println("This is a wrong cmd!")
            continue
        }
        fmt.Printf("%s - %s\n",p.cmd,p.desc)
        if p.handler != nil {
            p.handler()
        }
    }
}

func FindCmd(head *LinkTable,cmd string) *DataNode{
    node := (*DataNode)(unsafe.Pointer(GetLinkTableHead(head)))
    for ;node != nil;node = (*DataNode)(unsafe.Pointer(GetNextLinkTableNode(head,(*LinkTableNode)(unsafe.Pointer(node))))) {
        if node.cmd == cmd {
            return node
        }
    }
    return nil
}

func ShowAllCmd(head *LinkTable) int{
    pNode := (*DataNode)(unsafe.Pointer(GetLinkTableHead(head)))
    for ;pNode != nil;pNode = (*DataNode)(unsafe.Pointer(GetNextLinkTableNode(head,(*LinkTableNode)(unsafe.Pointer(pNode))))) {
        fmt.Printf("%s - %s\n",pNode.cmd,pNode.desc)
    }
    return 1
}

func Help() {
    ShowAllCmd(head)
}

func Quit() {
    os.Exit(0)
}

func InitMenuData(ppLinkTable **LinkTable) int {
    *ppLinkTable = CreateLinkTable()
    pNode := new(DataNode)
    pNode.cmd = "help"
    pNode.desc = "Menu List:"
    pNode.handler = Help;
    AddLinkTableNode(*ppLinkTable,(*LinkTableNode)(unsafe.Pointer(pNode)))
    pNode1 := new(DataNode)
    pNode1.cmd = "version"
    pNode1.desc = "Menu Program V1.0"
    pNode1.handler = nil;
    AddLinkTableNode(*ppLinkTable,(*LinkTableNode)(unsafe.Pointer(pNode1)))
    pNode2 := new(DataNode)
    pNode2.cmd = "quit"
    pNode2.desc = "Quit from Menu Program V1.0"
    pNode2.handler = Quit;
    AddLinkTableNode(*ppLinkTable,(*LinkTableNode)(unsafe.Pointer(pNode2)))
    return 1
}