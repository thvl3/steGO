# steGO - Simple LSB Steganography Tool

### Written by Ethan Hulse
![GitHub release (latest by date)](https://img.shields.io/github/v/release/thvl3/steGO)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/thvl3/steGO/build.yml?branch=main)
![GitHub license](https://img.shields.io/github/license/thvl3/steGO)

**steGO** is a lightweight steganography tool written in **Go** that allows you to **hide and extract secret messages** inside PNG images using **Least Significant Bit (LSB) encoding**.

## 🚀 Features
✅ **Embed messages into PNG images** using LSB.  
✅ **Extract hidden messages** from PNG images.  
✅ **Cross-platform support** (Linux, macOS, Windows).  
✅ **Lightweight & Fast** – written in Go, no dependencies required.  
✅ **Command-line interface** for easy use.  

---

## 📥 Download & Install
You can download precompiled binaries from the **[Releases](https://github.com/thvl3/steGO/releases)** page.

| Platform | Binary Name | Download Link |
|----------|------------|---------------|
| 🐧 **Linux** | `steGo` | [Download](https://github.com/thvl3/steGO/releases/latest) |
| 🍏 **macOS** | `steGoMac` | [Download](https://github.com/thvl3/steGO/releases/latest) |
| 🖥 **Windows** | `steGo.exe` | [Download](https://github.com/thvl3/steGO/releases/latest) |

### **Installation Steps**
1. Download the correct binary for your OS.
2. Make it executable (for Linux/macOS):
   ```sh
   chmod +x steGo  # or steGoMac on macOS
   ```
3. Move it to a directory in your `$PATH`:
   ```sh
   sudo mv steGo /usr/local/bin/  # or steGoMac on macOS
   ```

For **Windows**, simply place `steGo.exe` in any folder and run it from **Command Prompt or PowerShell**.

---

## ⚡ Usage
### **1️⃣ Hide a Message in an Image**
```sh
steGo encode image.png secret.png -file texttohide.txt
```
### **2️⃣ Extract a Hidden Message**
```sh
steGo decode secret.png
```

---

## 🛠 Building from Source
If you prefer to build `steGO` yourself, ensure you have **Go installed** and run:
```sh
git clone https://github.com/thvl3/steGO.git
cd steGO
go build -o steGo steGo.go  # Linux
GOOS=darwin GOARCH=amd64 go build -o steGoMac steGo.go  # macOS
GOOS=windows GOARCH=amd64 go build -o steGo.exe steGo.go  # Windows
```

---

## ❓ FAQ
### **1️⃣ What image formats does steGO support?**
Currently, steGO only supports **PNG images** for encoding and decoding.

### **2️⃣ Can I hide files instead of just text?**
Right now, steGO supports **text messages or .txt files**.

### **3️⃣ Is there a graphical user interface (GUI)?**
No, steGO is a **command-line tool**

---

## 🎯 Planned Features
🔹 **Support for additional image formats (JPEG, BMP, etc.)**  
🔹 **Encryption for added security**  
🔹 **Support for hiding entire files inside images**  
🔹 **More advanced steganography techniques (e.g., DCT-based hiding)**  

---

## 📜 License
This project is licensed under the **MIT License**. See [LICENSE](LICENSE) for details.

---

## ❤️ Contributing
Contributions are welcome! Feel free to submit **issues**, suggest **features**, or open **pull requests**.

### **Steps to Contribute**
1. **Fork** the repo and create a new branch.
2. Implement your changes.
3. Submit a **Pull Request** for review.

---

## 📧 Contact
For any questions, reach out via:
📬 **Email:** `ethanhulse.work@gmail.com`  
🐙 **GitHub Issues:** [Create an Issue](https://github.com/thvl3/steGO/issues)
