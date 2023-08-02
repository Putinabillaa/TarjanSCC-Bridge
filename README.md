# TarjanSCC-Bridge
> Program ini dapat mencari Strongly Connected Component (SCC) pada graf berarah dan Bridge pada graf tak berarah menggunakan algoritma Tarjan. Program tersedia dalam antarmuka CLI dan web.

## 1. Cara penggunaan
### Untuk CLI:
  <p align="center">
      <img src="https://github.com/Putinabillaa/TarjanSCC-Bridge/assets/109022993/2e7b6c48-a8af-4a2d-a4ab-71c6a611f32b" alt="SCC Bridge Finder - CLI" width="500px" />
    </p>

  - Jalankan command `make run` pada directory `CLI`
  - Setelah program berjalan, masukan metode input: `1. from file` untuk memasukan input graf dari file txt atau `2. from terminal` untuk mengetikan input graf pada terminal. Pada input, setap baris mewakili edge pada graf berarah dengan format: `<nama node asal> <spasi> <nama node tujuan>`. Contoh input file txt:
      
        A B
        B C
        C A
        B D
        D E
        E F
        F E
        
  - Masukan input sesuai instruksi
  - Output akan ditampilkan dalam bentuk teks array dengan visualisasi yang tersimpan dalam folder `output` dalam directory `CLI` dengan nama `graph.png` untuk visualisasi graph, `bridge.png` untuk visualisasi bridge, dan `scc.png` untuk visualisasi SCC. Waktu eksekusi juga akan ditampilkan. Contoh output:
    
    ![image](https://github.com/Putinabillaa/TarjanSCC-Bridge/assets/109022993/d9a5906b-5dc6-4379-a7e8-f89314196856)
    ![image](https://github.com/Putinabillaa/TarjanSCC-Bridge/assets/109022993/d447fb5c-2229-4328-99a7-5aa62d72c29f)
    ![image](https://github.com/Putinabillaa/TarjanSCC-Bridge/assets/109022993/f6bae86b-be31-48e2-a7ba-66ade4906730)

  ### Untuk Web:
  <p align="center">
    <img src="https://github.com/Putinabillaa/TarjanSCC-Bridge/assets/109022993/5c6c8c76-d2aa-4792-adf2-15cea37b3376" alt="SCC Bridge Finder - CLI" width="500px" />
  </p>
  
  - Jalankan command `make run` pada directory `Website/Backend` untuk menjalankan server. Server akan berjalan pada `http://localhost:8080`
  - Jalankan command `npm start` pada directory `Website/Frontend` untuk menjalankan frontend. Frontend akan berjalan pada `http://localhost:3000`
  - Masukan input melalui file txt atau teks dengan format yang sama dengan format input CLI.
  - Klik tombol `Find`
  - Output akan ditampilkan dalam bentuk teks array beserta visualisasinya. Waktu eksekusi juga akan ditampilkan. Contoh output:

    ![Screenshot 2023-07-31 at 11 02 1](https://github.com/Putinabillaa/TarjanSCC-Bridge/assets/109022993/a3b6f3fd-9e11-4a6b-a713-5e6e6d32bce4)
    
## 2. Algoritma Tarjan
### Kompleksitas
- `Kompleksitas Waktu:` Fungsi rekursif tarjan dipanggil sekali untuk setap node dan edge yang belum dikunjungi. Sehingga, kompleksitas waktu algoritma tarjan menjadi linear berdasarkan jumlah node and edge `O(|V|+|E|)`. Dengan catatan, pemeriksaan apakah suatu node berada pada stack dilakukan dalam waktu linear, dalam program ini menggunakan flag OnStack.
- `Kompleksitas Ruang:` O(6 * |V| + |E| + 1) ≈ `O(|V| + |E|)`
  ```
  /* Directed Graph untuk pencarian SCC */
  type DirGraph struct {
	Vertex  []string // V
	Adj     map[string][]string // E
	Index   int // 1
	Disc    map[string]int // V
	LowLink map[string]int // V
	OnStack map[string]bool // V
	Stack   []string // V (worst case)
	SCCs    [][]string // V
  }
  
  /* Undirected Graph untuk pencarian Bridge */
  type UnDirGraph struct {
	Vertex  []string // V
	Adj     map[string][]string // E
	Index   int // 1
	Disc    map[string]int //V
	LowLink map[string]int // V
	Visited map[string]bool // V
	Parent  map[string]string // V
	Bridges [][2]string // V
  }
  ```
### Modifikasi Algoritma Tarjan untuk Mendeteksi Bridge
Pada program ini bridge dideteksi dengan menganggap input graf berarah sebagai graf tidak berarah, sehingga digunakan struktur data yang berbeda untuk menyimpan graf. Selain itu, dilakukan modifikasi pada algoritma dengan mencatat parent dari setiap node (berdasarkan DFS). Pencatatan parent dilakukan karena lowlink dari suatu node hanya akan di-update dengan lowlink minimum dari semua tetangganya kecuali parent-nya (edge yang menghubungkan dengan node dengan parent-nya bukanlah back edge sehingga tidak perlu diperhitungkan). Berbeda dengan algoritma untuk SCC, di mana pengecekan parentnya dilakukan menggunakan stack yang sekaligus digunakan untuk mengidentifikasi nodes-nodes untuk SCC yang sama. Modifikasi juga dilakukan untuk kriteria bridge. Suatu edge (v, w) dikatakan bridge ketika lowlink w > dari index (disc) w, dengan kata lain node w hanya dapat dijangkau melalui v. Karena node w hanya dapat dijangkau melalui v, penghapusan (v, w) menyebabkan isolasi node w, dengan kata lain penghapusan edge tersebut membentuk SCC baru (sesuai definisi bridge).

### Back Edge & Tree Edge
Bakc edge adalah edge yang menghubungkan suatu node ke salah satu pendahulunya pada pohon DFS. Dengan kata lain, back edge menyebabkan cycle dalam graf. Pada iterasi DFS untuk node-node tetangga v, suatu edge (v, w) disebut back edge jika node w sudah pernah dikunjungi (artinya w berada pada level diatas v, dengan kata lain w pendahulu v) dan w bukan parent dari v.
Tree edge adalah edge yang terbentuk selama proses konstruksi pohon DFS. Tree edge menghubungkan suatu simpul dengan salah satu keturunannya pada pohon DFS. Pada iterasi DFS untuk node-node tetangga v, suatu edge (v, w) disebut tree edge jika node w belum pernah dikunjungi sehingga node w menjadi keturunan dari v.

## 3. Framework & Library
Selain library standard Go, berikut library dan framework tambahan yang digunakan.
- `Go graphviz` untuk visualisasi graph.
  ``` github.com/awalterschulze/gographviz ```
- `Color` untuk CLI berwarna.
  ``` github.com/fatih/color ```
- `encoding/json` untuk json encoding & decoding. 
- `Gorilla mux` untuk implementasi HTTP request router. ```github.com/gorilla/mux```
- `CORS` untuk menghandle Cross-Origin Resource Sharing (CORS). ```github.com/rs/cors```
- `React.js` untuk frontend UI.
- `React toastify` untuk implementasi UI notifikasi error.

## Referensi
https://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm

https://www.geeksforgeeks.org/bridge-in-a-graph/
