package point

import "fmt"

func zeroval(ival int) {
    ival = 0
}
func zeroptr(iptr *int) {
    *iptr = 0
}
func poi() {
	i := 1
    fmt.Println("initial:", i)

	

    zeroval(i)
    fmt.Println("zeroval:", i)
	v, j := 42, 2701

	p := &v         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(v)  // see the new value of i
	zeroptr(&i)
    fmt.Println("zeroptr:", i)
	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
	x := 0xFF
     
    // Displaying the values
    fmt.Printf("Type of variable x is %T\n", x)
    fmt.Printf("Value of x in hexadecimal is %X\n", x)
    fmt.Printf("Value of x in decimal is %v\n", x)
	for i := 0; i < 4; i++{ 
		fmt.Printf("test\n")   
	  } 
	  ii:= 0 
	  for ii < 3 { 
		 ii += 2 
	  } 
	fmt.Println(ii)  
     
}