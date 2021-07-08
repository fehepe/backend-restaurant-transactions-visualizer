package queries

const (
	FindBuyers = `
		query FindBuyers{
			buyers(func: type(Buyer)) {
				id
				name
				age
			}
		}
	`
	FindProducts = `
		query FindProducts{
			products(func: type(Product)) {
				id
				name
				price
			}
		}
	`

	FindTransactions = `
		query FindTransactions{
			transactions(func: type(Transaction)) {
				id
				buyerID
				ip
				device
				productIDs
			}
		}
	`

	BuyerByID = `		
		query BuyerByID($id: string) {
			buyers(func: eq(id, $id)) {
				id
				name
				age
			}
		}
	`
)
