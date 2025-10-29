# 🔢 İkilik ↔ Onluq Ədəd Çeviricisi (Go)

Bu layihə Go proqramlaşdırma dilində yazılmış sadə, lakin funksional konsol tətbiqidir.  
Proqram ikilik (binary) və onluq (decimal) say sistemləri arasında çevirmə aparır və daxil edilən məlumatları yoxlayaraq xətalı girişlərin qarşısını alır.


## ✨ Xüsusiyyətlər

✅ İkilik → Onluq çevirmə  
✅ Onluq → İkilik çevirmə  
✅ Daxil edilən məlumatın yoxlanması:
- İkilik ədədlər yalnız `0` və `1` simvollarından ibarət olmalıdır  
- Onluq ədədlər **0 ilə 1000 arasında** olmalıdır  
✅ Aydın və istifadəsi asan interfeys  
✅ Xəta mesajları aydın və izahlıdır  


## 🧠 Necə işləyir?

1. İstifadəçi çevirmə istiqamətini seçir:
   - `1` → Binary → Decimal  
   - `2` → Decimal → Binary  
2. Proqram konsoldan məlumatı oxuyur.  
3. Girişin düzgünlüyü yoxlanılır.  
4. Nəticə ekranda göstərilir.  


## 💻 İstifadə nümunələri

### ▶ Binary → Decimal

=== Say Sistemləri Çeviricisi (Binary <-> Decimal) ===
Seçimlər:
1 - Binary -> Decimal
2 - Decimal -> Binary
Seçiminizi edin (1 və ya 2): 1
Binary ədədi daxil edin (yalnız 0 və 1): 1100100
Binary 1100100  ->  Decimal 100

### ▶ Decimal → Binary

=== Say Sistemləri Çeviricisi (Binary <-> Decimal) ===
Seçimlər:
1 - Binary -> Decimal
2 - Decimal -> Binary
Seçiminizi edin (1 və ya 2): 2
Decimal ədədi daxil edin (maksimum 1000): 25
Decimal 25 -> Binary 11001


## ⚙️ Qurulum və İşə Salma

### 1. Layihəni klonlayın
```bash
git clone https://github.com/<istifadəçi-adı>/binary-decimal-converter.git
cd binary-decimal-converter
````

### 2. Proqramı işə salın

```bash
go run main.go
```

### 3. (İstəyə görə) icra edilə bilən fayl yaradın

```bash
go build -o cevirici main.go
./cevirici
```

---

## 🚨 Xəta hallarının izahı

| Ssenari               | Təsvir                               | Nümunə Giriş | Nəticə                                   |
| --------------------- | ------------------------------------ | ------------ | ---------------------------------------- |
| ❌ Yanlış binary giriş | 0 və 1-dən başqa simvol daxil edilib | `1021`       | `Xəta: yalnız 0 və 1 daxil edilə bilər`  |
| ⚠ Limit aşımı         | Decimal ədəd 1000-dən böyükdür       | `1500`       | `Xəta: nəticə 1000-dən böyük ola bilməz` |
| ⚠ Mənfi ədəd          | Decimal ədəd mənfidir                | `-5`         | `Xəta: mənfi ədədlər qəbul edilmir`      |

---



