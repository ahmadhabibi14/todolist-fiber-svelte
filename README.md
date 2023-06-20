## JavaScript, HTML, CSS, Svelte

Buatlah aplikasi frontend Svelte sederhana menggunakan svelte-mpa https://github.com/kokizzu/svelte-mpa
dimana terdapat 3 halaman:

- register/login (username, password) apabila belum login, ganti password dan logout apabila sudah login
- list todos, dengan 3 action: add todo, edit todo, delete todo, tiap baris todo ditampilkan username dan isi todo dan 2 tombol: edit apbila yang membuat todo adalah akun yang sedang login, apabila belum login, hanya bisa melihat

tiap halaman memiliki link menuju kedua halaman tersebut, dan tulisan "(c) [username email anda] [tahun]", misal: "(c) tejo 2022" di bagian bawah halaman

penilaian: kesesuaian dengan spesifikasi, common state bug (refresh/force reload, clear cookie)

5. [Ops] deploy kedua aplikasi nomor 3 dan 4 ke salah satu PaaS gratis, misal: fly.io, railway, koyeb atau provider yang lain.

penilaian: berhasil up dan berjalan dengan baik, kemampuan mendokumentasikan proses deployment.


## Golang, Fiber

Buatlah API server sederhana (tidak perlu menggunakan database sama sekali) dengan http://gofiber.io dengan 7 buah API route:

-  `/login`
   -  Input: username, password
   -  Output: sessionID (set di cookie), error | otomatis membuat akun apabila username belum ada
-  `/logout`
   -  Input: sessionID
   -  Output: error
-  `/todo/add`  
   -  Input: sessionID, text
   -  Output: todo object, error
-  `/todo/overwrite`  
   -  Input: sessionID, id, text
   -  Output: todo object, error | hanya boleh overwrite todo milik sendiri
-  `/todo/delete`  
   -  Input: sessionID: id
   -  Output: todo object, error
-  `/todo/list`  
   -  Input: page, limit
   -  Output: array of todo (username, text, created_at, updated_at) | urut berdasarkan updated_at descending
-  `/stats`  
   -  Output: total user logged in, total user registered, total todo created, total todo deleted

*`sessionID dikirimkan lewat cookie header.`*

**`Penilaian`**:  
Kesesuaian dengan spesifikasi, kelengkapan unit test, security, struktur code (pecah business logic, serialization/transport, dan data model)