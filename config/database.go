package config

import "be-assignment/prisma/db"

func ConnectDB() {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
}
