# Configuration Management and CLI

## Command Line
Command line adalah antarmuka berbasis teks yang cepat dan kuat yang digunakan pengembang untuk berkomunikasi secara lebih efektif dan efisien dengan komputer guna menyelesaikan serangkaian tugas yang lebih luas. 

Kenapa menggunakan command line?
1. Kontrol granular dari suatu OS atau aplikasi 
2. Manajemen lebih cepat dari sejumlah besar sistem operasi
3. Kemampuan untuk menyimpan skrip untuk diotomatisasi tugas rutin
4. Menghubungkan pengetahuan untuk penyelesaian masalah

Command line interface :
1. UNIX Shell
2. Command Prompt (MSDOS)
3. Python REPL
4. MySQL CLI client
5. Mongo Shell
6. redis-cli, etc.

## UNIX Shell
Normal User vs Root User
- Normal User = Diizinkan memanipulasi dir/home/username dir only
- Root User = Diizinkan memanipulasi / dir (semua directory)
- Normal User +  Sudoers = Diizinkan bertindak sebagai root sementara. 

Directory :
- pwd 
- ls
- mkdir
- cd
- rm
- cp
- mv
- ln

Files :
- create : touch
- view : head, cat, tail, less
- editor : vim, nano
- permission : chown, chmod
- different : diff

Network :
- ping
- ssh
- netstat
- nmap
- ip addr (ifconfig)
- wget
- curl
- telnet
- netcat

Utility :
- man
- env
- echo
- date
- which
- watch
- sudo
- history
- grep
- locate

Shell, Program yang berfungsi sebagai jembatan antara user dan kernel. Shell script adalah bahasa pemrograman yang dikompilasi berdasarkan perintah shell.