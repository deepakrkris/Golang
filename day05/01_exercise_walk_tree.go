package main

import "fmt"
import "golang.org/x/tour/tree"


func traverse(t *tree.Tree, ch chan int) {
   if t != nil {
     traverse(t.Left, ch) 
	  ch <- t.Value
	  traverse(t.Right, ch)
   }
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
   traverse(t, ch)
   close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
   ch1 := make(chan int, 0)
   ch2 := make(chan int, 0)
   go Walk(t1, ch1)
   go Walk(t2, ch2)
   
   for {
      x1, ok1 := <- ch1
	  x2, ok2 := <- ch2
	  
	  switch {
	     case ok1 != ok2 : 
		     return false
		 case !ok1 && !ok2 :
		     return true
	     case x1 != x2 :
		     return false
	     default :
	  }
   }
}

func main() {
   fmt.Println("Test Walk()")

   ch := make(chan int, 0)
   go Walk(tree.New(1), ch)
   for i := range ch {
      fmt.Print(i, " ")
   }
   
   fmt.Println()
   fmt.Println("Test Same()")
   fmt.Println("expected true; result is", Same(tree.New(1), tree.New(1)))
   fmt.Println("expected false; result is", Same(tree.New(1), tree.New(2)))
}
