# Civil Servant Warrior Backend

## Description
Civil Servant Warrior Backend adalah sebuah aplikasi backend yang ditulis menggunakan bahasa pemrograman Go (Golang) dengan tema Learning Management System. Aplikasi ini bertujuan untuk menyediakan backend yang handal, efisien, dan berkelanjutan untuk mendukung aplikasi web  yang terkait dengan Pendidikan. Aplikasi ini dibangun dengan menggunakan framework Gin, yang merupakan salah satu framework populer dalam pengembangan aplikasi menggunakan Go.

## Fitur Utama
-   Routing HTTP: Aplikasi ini menyediakan routing HTTP yang kuat dan fleksibel untuk menangani permintaan dari aplikasi klien. Pengguna dapat menentukan rute dan penanganan permintaan sesuai kebutuhan.
-   Koneksi Database: Aplikasi ini dapat terhubung ke database, seperti MySQL Hal ini memungkinkan pengguna untuk menyimpan dan mengambil data dari database sesuai kebutuhan.
-   Keamanan: Aplikasi ini memperhatikan aspek keamanan dengan menggunakan fitur-fitur seperti autentikasi dan otorisasi. Pengguna dapat mengimplementasikan lapisan keamanan sesuai dengan kebutuhan aplikasi.
-   Middleware: Aplikasi ini mendukung penggunaan middleware untuk memanipulasi dan memproses permintaan sebelum dan setelah rute dijalankan. Middleware dapat digunakan untuk melakukan validasi data, logging, dan lain sebagainya.
-   Pengaturan Konfigurasi: Aplikasi ini memungkinkan pengguna untuk mengatur konfigurasi melalui file atau variabel lingkungan (environment variables). Ini mempermudah pengguna untuk menyesuaikan aplikasi sesuai dengan lingkungan yang digunakan.

## Requirement & Stack

-   Golang >=1.18
-   Postgres >= 16.1

## Local Installation

1. Clone this project
    ```
    git clone https://github.com/Marvit-Solutions/csw-golang.git
    ```

2. Copy `.env.example` to `.env`
    ```
    cp .env.example .env
    ```
3. Configure environment variables for database connection
    ```
   APP_PORT="${APP_PORT}"
   DB_CONNECTION="${DB_CONNECTION}"
   DB_HOST="${DB_HOST}"
   DB_PORT="${DB_PORT}"
   DB_NAME="${DB_NAME}"
   DB_USERNAME="${DB_USERNAME}"
   DB_PASSWORD="${DB_PASSWORD}"
   SECRET_KEY="${SECRET_KEY}"
    ```

4.  Run the application
    ```
    go run main.go
    ```