# Penggunaan AI di Golang
## Kenapa memilih Go?
Beberapa alasan memilih Go:
- Dikembangkan oleh google untuk pengembangan aplikasi server-side dan berbasis cloud
- Go adalah compiled language sehingga lebih cepat dibanding interpreted language (contoh: Python)
- Mendukung concurrency melalui fitur goroutines sehingga dapat menjalankan banyak task secara simultan tanpa membebani memori
- Memiliki standar libraries yang sangat kaya, didukung oleh komunitas yang kuat

## Library Go untuk machine learning
1. GoLearn -> bersifat open source, simplicity & customizability.
2. Gorgonia -> Memudahkan menulis dan mengevaluasi persamaan matematika yang melibatkan array multidimensi, Low-level library (seperti Theano, namun lebih tinggi dibanding Tensorflow).
3. Gonum -> set packages yang didesain untuk menulis numerical & algoritma sains menjadi produktif, memiliki performa tinggi, dan scalable, Mirip seperti numpy dan scipy.
4. Gomind

## A.I. as a Service (AlaaS)
Sebuah tools AI yang dapat segera digunakan (pre-built or off-the-shelf A.I tools). Beberapa perusahaan penyedia AlaaS, contohnya: AWS, GCP, Microsoft Azure, IBM, OpenAi.

### Tipe-tipe Alaas
Bots/Chatbots
- Amazon lex (AWS)
- Azure Bot Service (Microsft Azure)
- DialogFlow(GCP) 

APIs
- Amazon Rekognition
- Amazon Comprehend
- Azure Cognitive Service
- Azure Speech Services
- Google Cloud Vision
- Cloud Natural language 

Machine learning
- Amazon SageMaker
- Azure Machine learning
- Google Cloud AI

### Keuntungan
1. Pengurangan Cost
2. Low-mode
3. Proses deployment cepat
4. Flexibility
5. Usability
6. Scalability
7. Customization

### Kelemahan
1. Cost yang berlebih dalam jangka panjang
2. Privasi dapat
3. Keamanan
4. Transparasi
5. Vendor lock-in

## OpenAI API
Target :
- Mendesain prompt yang tepat (prompt engineering)
- Membuat aplikasi berbasis AI menggunakan API OpenAI

OpenAI API Pricing
- Start for free , mendapatkan free credit senilai $5 yang dapat digunakan selama 3 bulan pertama semenjak register
- Pay as you go, setiap model memiliki tarif yang berbeda berbeda

Langkah mendapatkan API key OpenAI API
1. Buka situs web OpenAI
2. Lakukan signup jika belum memiliki akun, atau login jika sudah memiliki
3. Masuk ke bagian API keys atau melalui halaman berikut API Keys
4. Setelah masuk ke halaman API keys klik tombol create new secret key