package queries

const (
	FindBuyers = `
		query FindBuyers{
			buyers(func: type(Buyer)) {
				uid
				id
				name
				age
			}
		}
	`
	FindProducts = `
		query FindProducts{
			products(func: type(Product)) {
				uid
				id
				name
				price
			}
		}
	`

	FindTransactions = `
		query FindTransactions{
			transactions(func: type(Transaction)) {
				uid
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
