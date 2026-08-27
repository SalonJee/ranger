# Ranger

Ranger is a command-line tool that automatically organizes files inside a folder into categorized sub-folders based on their extensions (such as `images/`, `videos/`, `documents/`, `archives/`, etc.). 

If it encounters unknown extensions, it dynamically creates dedicated folders for them (e.g. `csv_files/`). It also includes a safe **undo feature** to revert your files to their original locations.

---

## 🏃 Quick Start

### 1. Clone the repository to your device
First, clone this repository and navigate into the project directory:
```bash
git clone https://github.com/<your-username>/ranger.git
cd ranger
```

### 2. Organize a Folder
Run the `main.go` file directly and pass the path of the folder you want to organize:
```bash
go run main.go /path/to/folder
```

### 3. Revert (Undo)
If you want to move all organized files back to their original state:
```bash
go run main.go --revert /path/to/folder
```

---

## 🚀 Optional: Install as a System Command (`ranger`)

If you want to run `ranger` from any directory without needing `go run main.go`:

1. Build the executable:
   ```bash
   go build -o ranger main.go
   ```

2. Move it to your local binary folder:
   ```bash
   mkdir -p ~/.local/bin
   mv ranger ~/.local/bin/
   ```

Now you can run it from anywhere in your terminal:
```bash
ranger /path/to/folder
ranger --revert /path/to/folder
```

---

## ⚙️ Features
- **Smart Categorization:** Automatically groups common formats (`.png`, `.jpg` $\rightarrow$ `images/`, `.pdf`, `.docx` $\rightarrow$ `documents/`, etc.).
- **Dynamic Folder Creation:** Unknown extensions get their own folder (e.g., `.blend` $\rightarrow$ `blend_files/`).
- **Safe Undo System:** Keeps track of file movements in a temporary `.ranger_undo.json` file so you can safely revert at any time.

## Implementation 

# Before
![Before ](/src_images/before.png)
# After
![After ](/src_images/after.png)
