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

	FindRecomendationsByProdId = `
	query ProductsRecomendations($id: string) {
		product(func: eq(id,$id)){
			id
			name
			price
		}
		
		var(func: eq(id, $id)) {
		  ~products {
				  products @filter(NOT eq(id, $id)) {
						  otherProds as id
				}
			}
		}
		  
		var(func: uid(uid(otherProds))) {
			  total as count(~products)
		}
		  
		productsRecomendation(func: uid(uid(otherProds)), orderdesc: val(total), first:4){
			id
		  	name
			price			
		} 
			
	}
	`
	FindBuyerDetailsById = `query BuyerDetails($id: string) {	
		buyer(func: eq(id, $id)) {
			uid
			id
			name
			age
			dgraph.type
		}	
		
		transactions(func: eq(buyerID,$id),first:5){
			id
			ip as ip
			device
			products: products {
				uid
			  id
			  name
			  price	
			  dgraph.type		  
			}
		}
				
		buyerEqIp(func: eq(ip, val(ip)),first:5) @filter(NOT uid(ip)) {
			device : device
			ip : ip
			buyer {
				uid
				name: name
				id:  id
			  	age: age
				dgraph.type
			}
		}
	}
	`
)
