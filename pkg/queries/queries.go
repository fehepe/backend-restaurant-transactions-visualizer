package queries

const (
	Schema = `
	id:          string   @index(exact) .
	name:        string                 .
	age:         int                    .
	price:       int                    .
	buyerID:     string   @index(exact) .
	ip:          string   @index(exact) .
	device:      string                 .
	productIDs:  [string] @index(exact) .
	products:    [uid]    @reverse      .
	buyer:       uid      @reverse      .
	
	type Buyer {
		id:   string
		name: string 
		age:  int			
	}
	
	type Product {
		id:    string
		name:  string
		price: int			
	}
	
	type Transaction {
		id:         string
		buyerID:    string
		buyer:      Buyer
		ip:         string
		device:     string
		productIDs: [string]
		products:   [Product]
	}`

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
		  
		productsRecomendation(func: uid(uid(otherProds)), orderdesc: val(total), first:1){
			id
		  	name
			price			
		} 
			
	}
	`
	FindBuyerDetailsById = `query BuyerDetails($id: string) {	
		buyer(func: eq(id, $id)) {
		
			id
			name
			age
			
		}	
		
		transactions(func: eq(buyerID,$id),first:5){
			id
			ip as ip
			device
			products: products {
			
			  id
			  name
			  price	
			 		  
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
				
			}
		}
	}
	`
)
