package utils

import "SecureFileshare/service/backend/models"

func ExistingUsers(username string) *models.User {
	var users = map[string]*models.User{
		"user1": {
			ID:        1,
			Username:  "user1",
			Password:  "user1@123",
			Role:      "user",
			MFASecret: "user1_mfa_secret",
		},
		"john": {
			ID:        2,
			Username:  "john",
			Password:  "john@123",
			Role:      "user",
			MFASecret: "john_mfa_secret",
		},
		"doe": {
			ID:        3,
			Username:  "doe",
			Password:  "doe@123",
			Role:      "user",
			MFASecret: "doe_mfa_secret",
		},
		"jane": {
			ID:        4,
			Username:  "jane",
			Password:  "jane@123",
			Role:      "user",
			MFASecret: "jane_mfa_secret",
		},
		"hal": {
			ID:        5,
			Username:  "hal",
			Password:  "hal@123",
			Role:      "user",
			MFASecret: "hal_mfa_secret",
		},
		"minnie": {
			ID:        6,
			Username:  "minnie",
			Password:  "minnie@123",
			Role:      "user",
			MFASecret: "minnie_mfa_secret",
		},
		"admin1": {
			ID:        7,
			Username:  "admin1",
			Password:  "admin1@123",
			Role:      "admin",
			MFASecret: "admin1_mfa_secret",
		},
		"admin2": {
			ID:        8,
			Username:  "admin2",
			Password:  "admin2@123",
			Role:      "super_admin",
			MFASecret: "admin2_mfa_secret",
		},
	}

	value, ok := users[username]
	if !ok {
		return nil
	}
	return value
}
