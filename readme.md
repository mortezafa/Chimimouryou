<h1 style="text-align: center;">Chimimouryou (魑魅魍魎): A Terminal Anime Streaming CLI</h1>


Chimimouryou is a terminal-based application that lets users search for anime, browse episodes, and stream them directly
using the MPV media player. Built with Go and the Bubble Tea framework, Chimimouryou provides a simple and lightweight
solution for anime enthusiasts. 

Chimimouryou is in its **very** early stages. I know there aren't many features yet, but I'll be adding as many as I can in
the coming months as my time permits. :)

## Key Features

- **Search Anime**: Search for your favorite anime by title.
- **Episode Navigation**: Browse through available episodes for any anime.
- **Direct Playback**: Stream episodes directly in MPV with high-quality playback.

## Installation

Homebrew is recommended to have installed on your machine if you are
using MacOS. Follow the download instructions [here](https://brew.sh/).

---

### **1. Local Installation**

#### Prerequisites

- **Operating System**: Linux, macOS, or Windows (with WSL).
- **Dependencies**:
    - Install all dependencies at once on macOS or Linux using Homebrew or a package manager of your choice:
      ```bash
      brew install go mpv git
      ```
    - Alternatively, install them individually:
        - [Go](https://go.dev/dl/) (1.20 or newer)
        - [MPV Media Player](https://mpv.io/)
        - [Git](https://git-scm.com/)

#### Steps

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/mortezafa/Chimimouryou.git
   cd Chimimouryou
   ```

2. **Build the Application**:

   ```bash
   go build -o chimi ./cmd/main.go
   ```

3. **Run the Application**:

   ```bash
   # Move the binary to a directory in your PATH to run it globally
   mv chimi /usr/local/bin/chimi

   # Now run the application using:
   chimi search <anime>
   ```

---

### **2. Docker Installation**

Coming Soon...

---

## Usage

### Example Commands

- **Search for Anime**:

  ```bash
  chimi search <anime>
  ```

  Replace `<anime>` with the title of the anime you want to search for.

- **Navigate Results**:

    - Use the arrow keys to browse the list.
    - Press `Enter` to select an anime or episode.

- **Play an Episode**:

    - Once an episode is selected, it will stream directly in MPV.

---

## Contributing

We welcome contributions from the community! Here's how you can help:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Submit a pull request with a clear description of your changes.

For larger changes, please open an issue first to discuss your ideas.

---

## License

This project is licensed under the MIT License. 

---

## Contact

If you have any questions or issues, feel free to open an issue on GitHub or contact us directly.

Happy streaming!

