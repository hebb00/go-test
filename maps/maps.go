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


func AddItem(bill, units map[string]int, item, unit string) bool {
	
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
	// To implement this, you'll need to:

    // Return false if the given item is not in the bill
    // Return false if the given unit is not in the units map.
    // Return false if the new quantity would be less than 0.
    // If the new quantity is 0, completely remove the item from the bill then return true.
    // Otherwise, reduce the quantity of the item and return true.
	q:= bill[item]-units[unit]

	_, exists := units[unit]
	_, e := bill[item]
	if e == false{
		return false
	}
	if exists == false{
		return false
	}
	if q < 0{
		
		return false
	
	}
	if q == 0{
		delete(bill, item)
		return true
	}
	bill[item]= q
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
