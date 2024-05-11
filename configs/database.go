package config

import (
	"be-assignment/prisma/db"
)

func ConnectDB() *db.PrismaClient {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil
	}

	return client
}
