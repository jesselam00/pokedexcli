# 🧭 Pokedex CLI

A command-line Pokédex built in Go. Explore the Pokémon world, catch Pokémon, and inspect your collection — all from your terminal. This project uses the [PokéAPI](https://pokeapi.co) to fetch Pokémon data in real time.

---

## Overview

A REPL (Read–Eval–Print Loop) CLI that makes HTTP requests to the PokéAPI, parses JSON responses in Go, and provides a short, snappy interface for exploring location areas and managing a local Pokédex.

---

## ⚙️ Installation

### Prerequisites
- [Go](https://go.dev/dl/) (latest version)

### Clone and Run
```bash
# Clone the repository
git clone https://github.com/<your-username>/pokedexcli.git
cd pokedexcli

# Build the binary
go build -o pokedex

# Run the Pokedex CLI
./pokedex
```

---

## 💬 Commands

Once inside the CLI, try the following commands:

### 📍 Map Navigation
```bash
map      # Displays the next 20 location areas
mapb     # Go back to the previous 20 location areas
```

### 🔍 Explore
```bash
explore <location-name>
# Lists all Pokémon found in a specific location area
```

### 🎯 Catch Pokémon
```bash
catch <pokemon-name>
# Attempts to catch a Pokémon (catch chance varies by base experience)
```

### 🧾 Inspect Pokémon
```bash
inspect <pokemon-name>
# View stats, height, weight, and types — only works if you've caught the Pokémon
```

### 📖 View Your Pokedex
```bash
pokedex
# Lists all Pokémon you've successfully caught
```

### ❓ Help & Exit
```bash
help
exit
```

---

## 🚀 Caching System

To make the CLI faster, responses from the PokéAPI are cached in memory.

**Cache design (summary):**
- `Cache` uses a `map[string]cacheEntry` protected by a `sync.Mutex` for concurrency safety.
- Each `cacheEntry` holds:
  - `createdAt` — when the entry was added.
  - `val` — raw cached data as `[]byte`.
- A background `reapLoop()` driven by a `time.Ticker` periodically removes entries older than the configured interval.

This reduces redundant API calls and makes repeated navigation feel instant.

---
