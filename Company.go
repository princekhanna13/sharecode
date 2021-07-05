func InitCompanies() []Company {
	companies := []Company{}
	company := Company{
		Name:     "Atoss",
		Location: Location{city: "Munich"},
		Employees: []Employee{
			{
				Name:   "Prince",
				Age:    29,
				Gender: "Male",
				Location: Location{
					city: "Munich",
				},
			},
			{
				Name:   "Andreas",
				Age:    20,
				Gender: "Male",
				Location: Location{
					city: "Timisoara",
				},
			},
			{
				Name:   "Alex",
				Gender: "Male",
				Location: Location{
					city: "Timisoara",
				},
			},
			{
				Name:   "Tom",
				Gender: "Male",
				Location: Location{
					city: "Timisoara",
				},
			},
			{
				Name:   "Robert",
				Gender: "Male",
				Location: Location{
					city: "Munich",
				},
			},
		}}
	company2 := Company{
		Name:     "Atoss",
		Location: Location{city: "Munich"},
		Employees: []Employee{
			{
				Name:   "Prince",
				Age:    29,
				Gender: "Male",
				Location: Location{
					city: "Munich",
				},
			},
			{
				Name:   "Andreas",
				Age:    20,
				Gender: "Male",
				Location: Location{
					city: "Timisoara",
				},
			},
			{
				Name:   "Alex",
				Gender: "Male",
				Location: Location{
					city: "Timisoara",
				},
			},
			{
				Name:   "Tom",
				Gender: "Male",
				Location: Location{
					city: "Timisoara",
				},
			},
			{
				Name:   "Robert",
				Gender: "Male",
				Location: Location{
					city: "Munich",
				},
			},
		}}
	companies = append(companies, company, company2)
	return companies
}
