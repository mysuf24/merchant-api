# merchant-api
ctx.BindJSON(&input) akan membaca body JSON dari request. error jika pada proses penulisan gagal akan menghasilkan error, dan sistem langsung merespons dengan status 400 dan pesan error.

config.DB.Create(&input) mencoba menyimpan data input (merchant) ke dalam database.