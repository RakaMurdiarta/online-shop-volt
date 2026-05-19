Required Versioning:
- golang : 1.26

Setup Project:
1. Jalankan di terminal cmd
    ```
    go mod tidy
    ```
2. copy .env.example >> .env
3. untuk menjalnkan program ketikan:
    ```
    air
    ```

Database + Object Storage (Supabase)
1. Masuk ke project
2. pilih connect di bagian atas website
3. pilih direct -> session pooler
4. lihat di bagian bawah terdapat connection string
5. sesuaikan di .env
6. setup object storage : 
    - pilih storage
    - buat bucket dengan nama e-commerce

7. get service role secret:
    - masuk ke settings
    - pilih api key
    - pilih tab ke 2 di atas , legacy
    - copy service role