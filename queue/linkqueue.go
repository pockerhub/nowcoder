package queue

type LinkNode struct {
    Next  *LinkNode
    Pre   *LinkNode
    Value interface{}
}

type LinkQueue struct {
    Values *LinkNode
    Length int
    Size   int
}

