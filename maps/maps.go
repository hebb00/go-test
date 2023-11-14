package gross

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	var myUnits = make(map[string]int, 0)
	myUnits["quarter_of_a_dozen"] = 3
	myUnits["half_of_a_dozen"] = 6
	myUnits["dozen"] = 12
	myUnits["small_gross"] = 120
	myUnits["gross"] = 144
	myUnits["great_gross"] = 1728
	return myUnits
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	var myBill = make(map[string]int, 0)
	return myBill
}

// AddItem adds an item to customer bill.
//bill := NewBill()
// units := Units()
// ok := AddItem(bill, units, "carrot", "dozen")
// fmt.Println(ok)
// Output: true (since dozen is a valid unit)
func AddItem(bill, units map[string]int, item, unit string) bool {
  value, exists := units[unit]
  _, e := bill[item]
  
  if exists == true{
	if e == true{
		delete(bill, item)
		bill[item]= value
		fmt.Print(bill)
		return true
	}

	bill[item]= value
	return true
  }

  return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	v, e := bill[item]
	if e == false{
		return false
	}
	if exists == false{
		return false
	}
	if v == 0 || value == 0{
		delete(bill, item)
		return true
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	v, e := bill[item]
	if e == true{
		return v, true
	}
	return v,false
}
