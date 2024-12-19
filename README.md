# st0pcha/godo

## 📒 About

**godo** is a CLI tool that simplifies the execution of scripts by managing them in a configuration file.

## 💻 Installation

1. Clone the repository:

```bash
git clone https://github.com/st0pcha/godo.git
cd godo
```

2. Build and install using Golang:

```bash
go build -o godo ./cmd
```

3. Add `./godo` to your environment's PATH
4. Enjoy!

## ⚙ Example .godo configuration

The .godo file should be placed in root of your project and contain a list of commands you need. For example:

```
run=node index.js
echo=echo Hello World!
python=python main.py
```

## 🔨 Usage

### Initialize a configuration file

To initialize .godo (configuration) file in your project, run:

```bash
godo --init
```

### Show help

To view all available godo commands, run:

```bash
godo --help
```

### Running a command

To execute a command defined in .godo file, run:

```bash
godo <command>
```

like a:

```bash
godo run
godo dev
godo test
godo python
```
