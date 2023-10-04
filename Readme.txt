Notes : Pertama kali run akan muncul Error dibawah, tetapi program akan terus berjalan baik kedepannya,
        Itu dikarenakan ketika program dijalankan pertama kali di database yang data di tabelnya masi kosong, maka First() dinyatakan record not found / masi kosong.

2023/10/04 11:35:20 C:/Users/mihsa/go/src/assignment3-019/controllers/weatherController.go:12 record not found
[0.739ms] [rows:0] SELECT * FROM "weathers" WHERE "weathers"."deleted_at" IS NULL ORDER BY "weathers"."id" LIMIT 1
2023/10/04 11:35:23 Water: 53.72 m (bahaya), Wind: 8.88 m/s (siaga)