package auth

func contains(a []usersRecord, predicate func(usersRecord) bool) Role {
	for _, v := range a {
		if predicate(v) {
			return v.Role
		}
	}
	return RoleGuest
}
