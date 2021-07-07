package queries

const (
	AllBuyers = `
		query AllBuyers{
			buyers(func: type(Buyer)) {
				id
				name
				age
			}
		}
	`
)
