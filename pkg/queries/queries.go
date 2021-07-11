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

	FindBuyerById = `
		query BuyerDetails($id: string) {
			buyer(func: eq(id, $id)){
				id
				name
				age
			}
				
			transactions(func: eq(id, $id), first: 10){
				id
				ipAddress as ipAddress
				device
				products: bought {
					id
					name
					price
				}

			}
				
			buyersWithSameIp(func: eq(ipAddress, val(ipAddress)), first: 10) @filter(NOT uid(ipAddress)) 
			{
				device
				ipAddress
				buyer: was_made_by  {
					id
					name
					age
				}
			}

			var(func: eq(id, $id)){
				made {
					bought {
						productsBought as id
					}
				}
			} 
		
			var(func: eq(id, val(productsBought))){
				id
				name
				price
				was_bought {
					id
					bought @filter(NOT uid(productsBought)) {
						productsToBeRecommended as id
					}
				}
			}
				
			var(func: eq(id, val(productsToBeRecommended))){
				id
				total as count(was_bought)
			}
				
			top10Products(func: uid(total), orderdesc: val(total), first: 10){
					id
					name
					price
				}
		}
	`
)
