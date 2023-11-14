package gross

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
	
    // Return false if the given unit is not in the units map.
    // Otherwise add the item to the customer bill, indexed by the item name, then return true.
    // If the item is already present in the bill, increase its quantity by the amount that belongs to the provided unit.

  value, exists := units[unit]
  v, e := bill[item]
  if exists == false{
	return false

  }else{
	if e == true{
		bill[item] = v+value
		return true
	  }
	bill[item] = value
	return true
  }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
	v, e := bill[item]
	if e == false{
		return false
	}
	if exists == false{
		return false
	}
	bill[item]= v-1
	if bill[item] < 0{
		return false
	
	}
	if bill[item] == 0{
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
