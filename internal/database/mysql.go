package database

import (
	"fmt"
	"log"
	"time"

	"github.com/sawalreverr/recything/config"
	aboutus "github.com/sawalreverr/recything/internal/about-us"
	achievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
	adminEntity "github.com/sawalreverr/recything/internal/admin/entity"
	"github.com/sawalreverr/recything/internal/article"
	customDataEntity "github.com/sawalreverr/recything/internal/custom-data"
	faqEntity "github.com/sawalreverr/recything/internal/faq"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/report"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlDatabase struct {
	DB *gorm.DB
}

var (
	dbInstance *mysqlDatabase
)

func NewMySQLDatabase(conf *config.Config) Database {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	log.Println("Successfully connected to database:", conf.DB.DBName)

	dbInstance = &mysqlDatabase{DB: db}
	return dbInstance
}

func (m *mysqlDatabase) InitWasteMaterials() {
	initialWasteMaterials := []report.WasteMaterial{
		{ID: "MTR01", Type: "plastik"},
		{ID: "MTR02", Type: "kaca"},
		{ID: "MTR03", Type: "kayu"},
		{ID: "MTR04", Type: "kertas"},
		{ID: "MTR05", Type: "baterai"},
		{ID: "MTR06", Type: "besi"},
		{ID: "MTR07", Type: "limbah berbahaya"},
		{ID: "MTR08", Type: "limbah beracun"},
		{ID: "MTR09", Type: "sisa makanan"},
		{ID: "MTR10", Type: "tak terdeteksi"},
	}

	for _, material := range initialWasteMaterials {
		m.DB.FirstOrCreate(&material, material)
	}

	log.Println("Dummy Waste material added!")
}

func (m *mysqlDatabase) InitSuperAdmin() {
	hashed, _ := helper.GenerateHash("superadmin@123")

	admin := adminEntity.Admin{
		ID:        "AD0001",
		Name:      "John Doe Senior",
		Email:     "john.doe.sr@gmail.com",
		Password:  hashed,
		Role:      "super admin",
		ImageUrl:  "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	m.GetDB().FirstOrCreate(&admin)
	log.Println("Dummy Super admin added!")
}

func (m *mysqlDatabase) InitFaqs() {
	faqs := []faqEntity.FAQ{
		{ID: "FAQ01", Category: "profil", Question: "Bagaimana cara saya memperbarui informasi profil saya?", Answer: "Anda dapat memperbarui informasi profil Anda melalui menu 'Pengaturan Profil' di aplikasi. Klik ikon profil, pilih 'Pengaturan', dan edit informasi yang diperlukan."},
		{ID: "FAQ02", Category: "profil", Question: "Apakah saya bisa mengubah alamat email yang sudah terdaftar?", Answer: "Ya, Anda bisa mengubah alamat email Anda melalui menu 'Pengaturan Profil'. Namun, Anda mungkin perlu memverifikasi alamat email baru Anda."},
		{ID: "FAQ03", Category: "profil", Question: "Bagaimana cara mengganti foto profil saya?", Answer: "Untuk mengganti foto profil, buka 'Profil Saya', klik pada foto profil Anda saat ini, dan pilih foto baru dari galeri atau ambil foto baru dengan kamera."},

		{ID: "FAQ04", Category: "littering", Question: "Bagaimana cara melaporkan sampah yang tidak pada tempatnya?", Answer: "Anda dapat melaporkan sampah yang tidak pada tempatnya melalui fitur 'Laporkan Sampah' di aplikasi. Ambil foto sampah tersebut, tambahkan deskripsi singkat, dan kirim laporan Anda."},
		{ID: "FAQ05", Category: "littering", Question: "Apakah ada sanksi bagi yang membuang sampah sembarangan?", Answer: "Ya, sesuai dengan peraturan daerah, membuang sampah sembarangan dapat dikenakan denda atau sanksi lainnya. Silakan cek peraturan lokal untuk detailnya."},
		{ID: "FAQ06", Category: "littering", Question: "Apa yang terjadi setelah saya melaporkan sampah?", Answer: "Setelah Anda melaporkan sampah, tim kami akan memverifikasi laporan tersebut dan mengkoordinasikan pembersihan dengan pihak berwenang setempat."},

		{ID: "FAQ07", Category: "rubbish", Question: "Apa saja jenis-jenis sampah yang dapat didaur ulang?", Answer: "Jenis sampah yang dapat didaur ulang termasuk plastik, kertas, kaca, dan logam. Pastikan untuk memisahkan sampah sesuai kategori sebelum mendaur ulang."},
		{ID: "FAQ08", Category: "rubbish", Question: "Bagaimana cara memisahkan sampah dengan benar?", Answer: "Pisahkan sampah berdasarkan jenisnya - organik, anorganik, dan berbahaya. Gunakan tempat sampah yang berbeda untuk setiap kategori untuk mempermudah proses daur ulang."},
		{ID: "FAQ09", Category: "rubbish", Question: "Apa yang dimaksud dengan sampah organik?", Answer: "Sampah organik adalah sampah yang berasal dari bahan-bahan alami yang dapat terurai, seperti sisa makanan, daun, dan potongan kayu."},

		{ID: "FAQ10", Category: "misi", Question: "Bagaimana cara berpartisipasi dalam misi kebersihan?", Answer: "Anda dapat berpartisipasi dalam misi kebersihan dengan mendaftar melalui aplikasi di bagian 'Misi'. Pilih misi yang tersedia dan ikuti instruksi yang diberikan."},
		{ID: "FAQ11", Category: "misi", Question: "Apa saja manfaat mengikuti misi kebersihan?", Answer: "Manfaat mengikuti misi kebersihan termasuk mendapatkan poin dan level, membantu menjaga lingkungan, dan berkesempatan memenangkan penghargaan."},
		{ID: "FAQ12", Category: "misi", Question: "Bagaimana cara menyelesaikan misi dan mendapatkan poin?", Answer: "Untuk menyelesaikan misi, ikuti semua instruksi yang diberikan dan laporkan hasil kerja Anda melalui aplikasi. Poin akan diberikan berdasarkan kontribusi Anda."},

		{ID: "FAQ13", Category: "lokasi sampah", Question: "Bagaimana cara menemukan tempat sampah terdekat?", Answer: "Anda dapat menemukan tempat sampah terdekat menggunakan fitur 'Cari Tempat Sampah' di aplikasi. Aplikasi akan menunjukkan lokasi tempat sampah di peta."},
		{ID: "FAQ14", Category: "lokasi sampah", Question: "Apa yang harus saya lakukan jika tidak menemukan tempat sampah di sekitar saya?", Answer: "Jika Anda tidak menemukan tempat sampah di sekitar Anda, simpan sampah Anda sampai Anda menemukan tempat yang sesuai untuk membuangnya atau laporkan kebutuhan tempat sampah baru melalui aplikasi."},
		{ID: "FAQ15", Category: "lokasi sampah", Question: "Apakah lokasi tempat sampah di aplikasi selalu diperbarui?", Answer: "Ya, kami berusaha untuk selalu memperbarui lokasi tempat sampah di aplikasi berdasarkan laporan pengguna dan data dari pihak berwenang setempat."},

		{ID: "FAQ16", Category: "poin dan level", Question: "Bagaimana cara mendapatkan poin?", Answer: "Anda bisa mendapatkan poin dengan menyelesaikan misi, melaporkan sampah, dan berpartisipasi dalam kegiatan kebersihan. Poin akan otomatis ditambahkan ke akun Anda."},
		{ID: "FAQ17", Category: "poin dan level", Question: "Apa yang bisa saya lakukan dengan poin yang saya kumpulkan?", Answer: "Poin yang Anda kumpulkan bisa ditukar dengan berbagai hadiah, diskon, atau digunakan untuk meningkatkan level akun Anda dalam aplikasi."},
		{ID: "FAQ18", Category: "poin dan level", Question: "Bagaimana cara meningkatkan level saya?", Answer: "Tingkatkan level Anda dengan mengumpulkan poin dari berbagai aktivitas dalam aplikasi. Setiap level baru memberikan akses ke fitur dan penghargaan tambahan."},

		{ID: "FAQ19", Category: "artikel", Question: "Di mana saya bisa membaca artikel terkait daur ulang dan kebersihan?", Answer: "Anda bisa membaca artikel terkait daur ulang dan kebersihan di bagian 'Artikel' dalam aplikasi. Kami menyediakan berbagai artikel informatif untuk membantu Anda lebih peduli terhadap lingkungan."},
		{ID: "FAQ20", Category: "artikel", Question: "Apakah artikel di aplikasi diperbarui secara berkala?", Answer: "Ya, artikel di aplikasi diperbarui secara berkala dengan konten terbaru mengenai daur ulang, tips kebersihan, dan informasi lingkungan lainnya."},
		{ID: "FAQ21", Category: "artikel", Question: "Bisakah saya berkontribusi menulis artikel untuk aplikasi?", Answer: "Tentu saja! Kami menerima kontribusi dari pengguna. Jika Anda tertarik, silakan hubungi kami melalui fitur 'Kontak Kami' di aplikasi untuk informasi lebih lanjut tentang cara berkontribusi."},
	}

	for _, faq := range faqs {
		m.GetDB().FirstOrCreate(&faq, faq)
	}
	log.Println("Dummy FAQs added!")
}

func (m *mysqlDatabase) InitCustomDatas() {
	datas := []customDataEntity.CustomData{
		{ID: "CDT0001", Topic: "Daur Ulang Plastik", Description: "Proses daur ulang plastik melibatkan pengumpulan sampah plastik, pembersihan, penghancuran menjadi serpihan kecil, dan kemudian melelehkannya untuk dibentuk menjadi produk baru. Plastik yang dapat didaur ulang termasuk botol air, wadah makanan, dan kantong belanja tertentu."},
		{ID: "CDT0002", Topic: "Pemanfaatan Sampah Organik", Description: "Sampah organik seperti sisa makanan dan daun dapat diubah menjadi kompos yang berguna sebagai pupuk alami. Proses ini melibatkan penguraian bahan organik oleh mikroorganisme dalam kondisi yang terkontrol."},
		{ID: "CDT0003", Topic: "Pengelolaan Sampah Elektronik", Description: "Sampah elektronik seperti ponsel lama, komputer, dan televisi harus dibawa ke pusat daur ulang elektronik. Komponen-komponen berharga seperti logam mulia bisa diekstraksi dan digunakan kembali, sementara bahan berbahaya dikelola dengan aman."},
		{ID: "CDT0004", Topic: "Kompetisi Pengurangan Sampah", Description: "Kompetisi ini mengajak masyarakat untuk mengurangi sampah yang mereka hasilkan dalam periode tertentu. Pemenang akan ditentukan berdasarkan jumlah sampah yang berhasil dikurangi dan kreativitas dalam mendaur ulang atau mengurangi penggunaan barang sekali pakai."},
		{ID: "CDT0005", Topic: "Melaporkan Sampah yang Tidak pada Tempatnya", Description: "Pengguna aplikasi dapat melaporkan sampah yang ditemukan di tempat umum yang tidak sesuai. Laporan harus mencakup foto, lokasi, dan jenis sampah. Tim kebersihan akan diberitahu untuk mengambil tindakan."},
		{ID: "CDT0006", Topic: "Daur Ulang Kertas", Description: "Kertas dapat didaur ulang menjadi produk baru dengan cara dikumpulkan, dipisahkan berdasarkan jenis, dihancurkan menjadi pulp, kemudian dibersihkan dan diproses menjadi kertas baru. Produk seperti koran, majalah, dan karton sering kali dapat didaur ulang."},
		{ID: "CDT0007", Topic: "Penggunaan Ulang Barang Bekas", Description: "Banyak barang bekas seperti pakaian, furnitur, dan alat rumah tangga masih bisa digunakan kembali. Dengan sedikit kreativitas, barang-barang ini bisa diperbaiki atau dimodifikasi untuk digunakan kembali, mengurangi jumlah sampah yang berakhir di tempat pembuangan akhir."},
		{ID: "CDT0008", Topic: "Pengelolaan Sampah Anorganik", Description: "Sampah anorganik seperti kaca, logam, dan beberapa jenis plastik bisa didaur ulang. Pengelolaan yang tepat melibatkan pemisahan berdasarkan jenis bahan, pembersihan, dan pengiriman ke fasilitas daur ulang yang sesuai."},
		{ID: "CDT0009", Topic: "Dampak Lingkungan dari Sampah Plastik", Description: "Sampah plastik yang tidak terkelola dengan baik bisa mencemari lingkungan, termasuk lautan. Plastik membutuhkan ratusan tahun untuk terurai dan bisa membahayakan kehidupan laut. Mengurangi penggunaan plastik sekali pakai dan mendaur ulang plastik yang ada adalah langkah penting untuk mengatasi masalah ini."},
		{ID: "CDT0010", Topic: "Teknologi Baru dalam Daur Ulang", Description: "Teknologi baru seperti pemanfaatan enzim untuk mendaur ulang plastik dan penggunaan sensor untuk pengelolaan sampah cerdas sedang dikembangkan. Teknologi ini bertujuan untuk meningkatkan efisiensi proses daur ulang dan mengurangi jumlah sampah yang tidak terkelola dengan baik."},
	}

	for _, data := range datas {
		m.GetDB().FirstOrCreate(&data, data)
	}
	log.Println("Dummy Custom Data added!")
}

func (m *mysqlDatabase) InitTasks() {
	dummyData := []task.TaskChallenge{
		{
			ID:          "TM0001",
			Title:       "Challenge 1",
			Description: "Description for Challenge 1",
			Thumbnail:   "thumbnail1.jpg",
			StartDate:   time.Now(),
			EndDate:     time.Now().AddDate(0, 0, 30),
			Point:       1000,
			AdminId:     "AD0001",
		},
		{
			ID:          "TM0002",
			Title:       "Challenge 2",
			Description: "Description for Challenge 2",
			Thumbnail:   "thumbnail2.jpg",
			StartDate:   time.Now(),
			EndDate:     time.Now().AddDate(0, 1, 0),
			Point:       200,
			AdminId:     "AD0001",
		},
	}

	for _, taskChallenge := range dummyData {
		result := m.GetDB().FirstOrCreate(&taskChallenge, task.TaskChallenge{ID: taskChallenge.ID})
		if result.Error != nil {
			log.Printf("Error adding task challenge with ID %s: %v", taskChallenge.ID, result.Error)
		} else if result.RowsAffected == 0 {
			log.Printf("Task challenge with ID %s already exists.", taskChallenge.ID)
		} else {
			log.Printf("Task challenge with ID %s created.", taskChallenge.ID)
		}
	}

	log.Println("Dummy Tasks added!")
}

func (m *mysqlDatabase) InitTaskSteps() {
	dummyData := []task.TaskStep{
		{
			TaskChallengeId: "TM0001",
			Title:           "Step 1",
			Description:     "Description for Step 1",
		},
		{
			TaskChallengeId: "TM0001",
			Title:           "Step 2",
			Description:     "Description for Step 2",
		},
		{
			TaskChallengeId: "TM0001",
			Title:           "Step 3",
			Description:     "Description for Step 3",
		},
		{
			TaskChallengeId: "TM0002",
			Title:           "Step 1",
			Description:     "Description for Step 1",
		},
		{
			TaskChallengeId: "TM0002",
			Title:           "Step 2",
			Description:     "Description for Step 2",
		},
		{
			TaskChallengeId: "TM0002",
			Title:           "Step 3",
			Description:     "Description for Step 3",
		},
	}

	for _, taskStep := range dummyData {
		m.GetDB().FirstOrCreate(&taskStep, taskStep)
	}

	log.Println("Dummy Task Steps added!")
}

func (m *mysqlDatabase) InitAchievements() {
	dumyData := []achievement.Achievement{
		{
			Level:        "classic",
			TargetPoint:  0,
			BadgeUrl:     "https://res.cloudinary.com/dymhvau8n/image/upload/v1717758679/achievement_badge/cq2n246e6twuksnia62t.png",
			BadgeUrlUser: "https://res.cloudinary.com/dymhvau8n/image/upload/v1718189121/user_badge/htaemsjtlhfof7ww01ss.png",
		},
		{
			Level:        "silver",
			TargetPoint:  50000,
			BadgeUrl:     "https://res.cloudinary.com/dymhvau8n/image/upload/v1717758731/achievement_badge/b8igluyain8bwyjusfpk.png",
			BadgeUrlUser: "https://res.cloudinary.com/dymhvau8n/image/upload/v1718189221/user_badge/oespnjdgoynkairlutbk.png",
		},
		{
			Level:        "gold",
			TargetPoint:  150000,
			BadgeUrl:     "https://res.cloudinary.com/dymhvau8n/image/upload/v1717758761/achievement_badge/lazzyh9tytvb4rophbc3.png",
			BadgeUrlUser: "https://res.cloudinary.com/dymhvau8n/image/upload/v1718189184/user_badge/jshs1s2fwevahgtvjkgj.png",
		},
		{
			Level:        "platinum",
			TargetPoint:  300000,
			BadgeUrl:     "https://res.cloudinary.com/dymhvau8n/image/upload/v1717758798/achievement_badge/xc8msr6agowzhfq8ss8a.png",
			BadgeUrlUser: "https://res.cloudinary.com/dymhvau8n/image/upload/v1718188250/user_badge/icureiapdvtzyu5b99zu.png",
		},
	}

	for _, data := range dumyData {
		m.GetDB().FirstOrCreate(&data, data)
	}

	log.Println("Dummy Achievements added!")
}

func (m *mysqlDatabase) InitDataVideos() {
	videos := []video.Video{
		{
			ID:          1,
			Title:       "Daur Ulang",
			Description: "Tips Daur Ulang",
			Link:        "https://www.youtube.com/watch?v=MJd3bo_XRaU",
		},
		{
			ID:          2,
			Title:       "Tutorial Bernapas",
			Description: "Tutorial Bernapas Bagi Pemula",
			Link:        "https://www.youtube.com/watch?v=jp5uhrdhsKI",
		},
	}

	for _, video := range videos {
		m.GetDB().FirstOrCreate(&video, video)
	}

	log.Println("Video data added!")
}

func (m *mysqlDatabase) InitVideoCategories() {
	videoCategories := []video.VideoCategory{
		{Name: "tips", VideoID: 1},
		{Name: "daur ulang", VideoID: 1},
		{Name: "tutorial", VideoID: 1},
		{Name: "edukasi", VideoID: 2},
		{Name: "kampanye", VideoID: 2},
	}
	for _, videoCategory := range videoCategories {
		m.GetDB().FirstOrCreate(&videoCategory, videoCategory)
	}
	log.Println("Video categories data added!")
}

func (m *mysqlDatabase) InitAboutUs() {
	aboutUs := []aboutus.AboutUs{
		{ID: "ABS01", Category: "perusahaan", Title: "Tentang siapa kami", Description: "RecyThing adalah pemimpin di industri daur ulang sampah yang berkomitmen untuk menjaga lingkungan hidup yang lebih bersih dan lebih berkelanjutan."},
		{ID: "ABS02", Category: "perusahaan", Title: "Visi Kami", Description: "Menciptakan masyarakat yang sadar lingkungan di mana setiap individu berperan aktif dalam melestarikan bumi kita."},
		{ID: "ABS03", Category: "perusahaan", Title: "Komitmen Kami", Description: "Prioritaskan penggunaan teknologi terbaru dan praktik terbaik dalam proses daur ulang untuk mengurangi dampak lingkungan."},
		{ID: "ABS04", Category: "perusahaan", Title: "Pelayanan Pelanggan Unggul", Description: "Tim ahli yang berpengalaman memberikan solusi tepat dan responsif sesuai dengan kebutuhan klien."},
		{ID: "ABS05", Category: "perusahaan", Title: "Pendidikan Masyarakat", Description: "Berperan aktif dalam mendidik masyarakat tentang pentingnya daur ulang dan pengelolaan limbah yang berkelanjutan."},

		{ID: "ABS06", Category: "tim", Title: "Tim Manajemen", Description: "Lorem ipsum dolor sit amet consectetur. Faucibus ultricies neque pellentesque tempus eros nulla ultrices laoreet. Posuere placerat cras fames egestas. Turpis odio molestie nec viverra nam justo risus. Suspendisse eget id hac diam faucibus adipiscing."},
		{ID: "ABS07", Category: "tim", Title: "Tim Manajemen", Description: "Lorem ipsum dolor sit amet consectetur. Faucibus ultricies neque pellentesque tempus eros nulla ultrices laoreet. Posuere placerat cras fames egestas. Turpis odio molestie nec viverra nam justo risus. Suspendisse eget id hac diam faucibus adipiscing."},
		{ID: "ABS08", Category: "tim", Title: "Tim Manajemen", Description: "Lorem ipsum dolor sit amet consectetur. Faucibus ultricies neque pellentesque tempus eros nulla ultrices laoreet. Posuere placerat cras fames egestas. Turpis odio molestie nec viverra nam justo risus. Suspendisse eget id hac diam faucibus adipiscing."},
		{ID: "ABS09", Category: "tim", Title: "Tim Manajemen", Description: "Lorem ipsum dolor sit amet consectetur. Faucibus ultricies neque pellentesque tempus eros nulla ultrices laoreet. Posuere placerat cras fames egestas. Turpis odio molestie nec viverra nam justo risus. Suspendisse eget id hac diam faucibus adipiscing."},

		{ID: "ABS10", Category: "contact_us", Title: "Hubungi Kami", Description: "Jika Anda memiliki pertanyaan, masukan, atau ingin bermitra dengan kami, jangan ragu untuk menghubungi tim kami. Kami siap membantu Anda dengan segala kebutuhan terkait daur ulang dan pengelolaan limbah."},
		{ID: "ABS11", Category: "contact_us", Title: "Alamat Kantor", Description: "Recything\nJalan Mangga Dua\nJakarta Pusat, 20012\nIndonesia"},
		{ID: "ABS12", Category: "contact_us", Title: "Jam Operasional", Description: "Senin-Jumat: 08.00 - 17.00 WIB"},
		{ID: "ABS13", Category: "contact_us", Title: "Telepon", Description: "+6289511223344"},
		{ID: "ABS14", Category: "contact_us", Title: "Social Media", Description: "Facebook: https://facebook.com/recything\nTwitter: https://x.com/recything\nInstagram: https://instagram.com/recything\nLinkedin: https://linkedin.com/recything"},
	}

	aboutUsImages := []aboutus.AboutUsImage{
		{ID: "ABSI01", AboutUsID: "ABS01", ImageURL: "https://res.cloudinary.com/dlbbsdd3a/image/upload/v1717758300/recything/about-us/kan9fdnp7h6o4hfclghm.png"},

		{ID: "ABSI02", AboutUsID: "ABS05", ImageURL: "https://res.cloudinary.com/dlbbsdd3a/image/upload/v1717758301/recything/about-us/spgrokvm9un0yq5zsycn.png"},
		{ID: "ABSI03", AboutUsID: "ABS05", ImageURL: "https://res.cloudinary.com/dlbbsdd3a/image/upload/v1717758301/recything/about-us/tynymzulgmkwiqu4a7mb.png"},

		{ID: "ABSI04", AboutUsID: "ABS06", ImageURL: "https://res.cloudinary.com/dlbbsdd3a/image/upload/v1717758300/recything/about-us/qwck5a9jwqht6rdjmxwd.png"},
		{ID: "ABSI05", AboutUsID: "ABS06", ImageURL: "https://res.cloudinary.com/dlbbsdd3a/image/upload/v1717758300/recything/about-us/qwck5a9jwqht6rdjmxwd.png"},
		{ID: "ABSI06", AboutUsID: "ABS06", ImageURL: "https://res.cloudinary.com/dlbbsdd3a/image/upload/v1717758300/recything/about-us/qwck5a9jwqht6rdjmxwd.png"},

		{ID: "ABSI07", AboutUsID: "ABS07", ImageURL: "https://res.cloudinary.com/dlbbsdd3a/image/upload/v1717758300/recything/about-us/qwck5a9jwqht6rdjmxwd.png"},
		{ID: "ABSI08", AboutUsID: "ABS07", ImageURL: "https://res.cloudinary.com/dlbbsdd3a/image/upload/v1717758300/recything/about-us/qwck5a9jwqht6rdjmxwd.png"},
		{ID: "ABSI09", AboutUsID: "ABS07", ImageURL: "https://res.cloudinary.com/dlbbsdd3a/image/upload/v1717758300/recything/about-us/qwck5a9jwqht6rdjmxwd.png"},

		{ID: "ABSI10", AboutUsID: "ABS08", ImageURL: "https://res.cloudinary.com/dlbbsdd3a/image/upload/v1717758300/recything/about-us/qwck5a9jwqht6rdjmxwd.png"},
		{ID: "ABSI11", AboutUsID: "ABS08", ImageURL: "https://res.cloudinary.com/dlbbsdd3a/image/upload/v1717758300/recything/about-us/qwck5a9jwqht6rdjmxwd.png"},
		{ID: "ABSI12", AboutUsID: "ABS08", ImageURL: "https://res.cloudinary.com/dlbbsdd3a/image/upload/v1717758300/recything/about-us/qwck5a9jwqht6rdjmxwd.png"},

		{ID: "ABSI13", AboutUsID: "ABS09", ImageURL: "https://res.cloudinary.com/dlbbsdd3a/image/upload/v1717758300/recything/about-us/qwck5a9jwqht6rdjmxwd.png"},
		{ID: "ABSI14", AboutUsID: "ABS09", ImageURL: "https://res.cloudinary.com/dlbbsdd3a/image/upload/v1717758300/recything/about-us/qwck5a9jwqht6rdjmxwd.png"},
		{ID: "ABSI15", AboutUsID: "ABS09", ImageURL: "https://res.cloudinary.com/dlbbsdd3a/image/upload/v1717758300/recything/about-us/qwck5a9jwqht6rdjmxwd.png"},

		{ID: "ABSI16", AboutUsID: "ABS10", ImageURL: "https://res.cloudinary.com/dlbbsdd3a/image/upload/v1717758300/recything/about-us/mfi5xij2xssmztqwaybz.png"},
	}

	for _, about := range aboutUs {
		m.GetDB().FirstOrCreate(&about, about)
	}

	for _, aboutImage := range aboutUsImages {
		m.GetDB().FirstOrCreate(&aboutImage, aboutImage)
	}

	log.Println("About-us data added!")
}

func (m *mysqlDatabase) InitArticleCategory() {
	categories := []article.WasteCategory{
		{ID: 1, Name: "plastik"},
		{ID: 2, Name: "besi"},
		{ID: 3, Name: "kaca"},
		{ID: 4, Name: "organik"},
		{ID: 5, Name: "kayu"},
		{ID: 6, Name: "kertas"},
		{ID: 7, Name: "baterai"},
		{ID: 8, Name: "kaleng"},
		{ID: 9, Name: "elektronik"},
		{ID: 10, Name: "tekstil"},
		{ID: 11, Name: "minyak"},
		{ID: 12, Name: "bola lampu"},
		{ID: 13, Name: "berbahaya"},
	}

	for _, category := range categories {
		m.GetDB().FirstOrCreate(&category, category)
	}
	log.Println("Article categories data added!")
}

func (m *mysqlDatabase) InitTrashCategoryVideo() {
	trashCategories := []video.TrashCategory{
		{ID: 1, Name: "plastik", VideoID: 1},
		{ID: 2, Name: "besi", VideoID: 1},
		{ID: 3, Name: "kaca", VideoID: 1},
		{ID: 4, Name: "organik", VideoID: 1},
		{ID: 5, Name: "kayu", VideoID: 1},
		{ID: 6, Name: "kertas", VideoID: 1},
		{ID: 7, Name: "baterai", VideoID: 2},
		{ID: 8, Name: "kaleng", VideoID: 2},
		{ID: 9, Name: "elektronik", VideoID: 2},
		{ID: 10, Name: "tekstil", VideoID: 2},
		{ID: 11, Name: "minyak", VideoID: 2},
		{ID: 12, Name: "bola lampu", VideoID: 2},
		{ID: 13, Name: "berbahaya", VideoID: 2},
	}

	for _, category := range trashCategories {
		m.GetDB().FirstOrCreate(&category, category)
	}
	log.Println("Trash categories data added!")
}

func (m *mysqlDatabase) InitArticle() {
	articles := []article.Article{
		{ID: "ART0001", Title: "Cara Mendaur Ulang Botol Plastik", Description: "Panduan langkah demi langkah tentang cara mendaur ulang botol plastik di rumah.", ThumbnailURL: "https://example.com/daur-ulang-plastik.jpg", AuthorID: "AD0001"},
		{ID: "ART0002", Title: "Bahaya Limbah Elektronik dan Cara Pembuangan yang Bertanggung Jawab", Description: "Pelajari tentang bahaya lingkungan dan kesehatan dari limbah elektronik dan temukan metode pembuangan yang aman.", ThumbnailURL: "https://example.com/limbah-elektronik.jpg", AuthorID: "AD0001"},
	}

	articleSections := []article.ArticleSection{
		{ID: 1, ArticleID: "ART0001", Title: "Membersihkan dan Menyortir", Description: "Bilas botol, lepaskan label, dan pisahkan berdasarkan jenisnya (PET, HDPE, dll.).", ImageURL: "https://example.com/membersihkan-botol.jpg"},
		{ID: 2, ArticleID: "ART0001", Title: "Meremukkan dan Menyimpan", Description: "Remukkan botol untuk menghemat ruang dan simpan dalam tempat khusus.", ImageURL: "https://example.com/meremukkan-botol.jpg"},
		{ID: 3, ArticleID: "ART0001", Title: "Mengantar atau Dijemput", Description: "Temukan pusat daur ulang terdekat atau jadwalkan layanan penjemputan.", ImageURL: "https://example.com/pusat-daur-ulang.jpg"},

		{ID: 4, ArticleID: "ART0002", Title: "Apa itu Limbah Elektronik?", Description: "Limbah elektronik mencakup barang elektronik bekas seperti ponsel, komputer, peralatan rumah tangga, dll.", ImageURL: "https://example.com/tumpukan-limbah-elektronik.jpg"},
		{ID: 5, ArticleID: "ART0002", Title: "Komponen Beracun", Description: "Limbah elektronik mengandung zat berbahaya seperti timbal, merkuri, dan kadmium.", ImageURL: "https://example.com/simbol-beracun.jpg"},
		{ID: 6, ArticleID: "ART0002", Title: "Pembuangan yang Benar", Description: "Cari pendaur ulang limbah elektronik bersertifikat atau program pengembalian yang ditawarkan oleh produsen.", ImageURL: "https://example.com/daur-ulang-limbah-elektronik.jpg"},
	}

	articleCategories := []article.ArticleCategories{
		{ID: 1, ArticleID: "ART0001", WasteCategoryID: 1, ContentCategoryID: 3},
		{ID: 2, ArticleID: "ART0002", WasteCategoryID: 9, ContentCategoryID: 4},
		{ID: 3, ArticleID: "ART0002", WasteCategoryID: 13, ContentCategoryID: 5},
	}

	for _, article := range articles {
		m.GetDB().FirstOrCreate(&article, article)
	}

	for _, articleSection := range articleSections {
		m.GetDB().FirstOrCreate(&articleSection, articleSection)
	}

	for _, articleCategory := range articleCategories {
		m.GetDB().FirstOrCreate(&articleCategory, articleCategory)
	}

	log.Println("Article data added!")
}

func (m *mysqlDatabase) GetDB() *gorm.DB {
	return dbInstance.DB
}
