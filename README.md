# Fotos App

![logo-color](https://user-images.githubusercontent.com/71443682/170883787-78cf6584-087f-4e0b-a711-72d3f63fdd1a.svg)

---

# Installation - Client

Clone the repository to your local device and

1. Install the node modules:

```bash
yarn install
```

2. Copy and Configure the `src/utils/config.example` then rename it as `src/utils/config.ts` inside it, paste the following

```typescript
export const config = {
	baseUrl: "http://localhost:8080/",
};
```

3. Start the server in developer mode:

```bash
yarn dev dev
```

The server should now run

# Installation - Server

Clone the repository to your local device and make sure you have [BRA](github.com/unknwon/bra) installed

1. Install the Go Modules:

```bash
go mod download
```

2. Copy and Configure the `config/config.example` then rename it as `config/config.json`. Fill the neccessary info.

3. Copy and Configure the `.env.example` then rename it as `.env` and fill the necessary information

3. Start the server in developer mode:

```bash
bra run
```

The server should now run

---