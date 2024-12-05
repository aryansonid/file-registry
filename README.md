# FileRegistry

## Overview

**FileRegistry** is a blockchain-based file registry system that allows users to upload files to IPFS (InterPlanetary File System) and store their content identifiers (CID) on the Ethereum blockchain. Users can query the stored CIDs using file paths, ensuring decentralized and tamper-proof file storage.

The project consists of:
- A Solidity smart contract deployed on the Ethereum blockchain.
- A Go-based API to interact with the smart contract and IPFS.
- Docker setup for the IPFS node and Go server.

## Features

- **Upload Files to IPFS**: Upload files and store their CIDs.
- **File Path to CID Mapping**: Store the mapping between a file path and its CID on-chain, can only be done by the contract owner to avoid false entries.
- **Retrieve CID by File Path**: Fetch the CID of a file from the blockchain by providing the file path.

---

## Prerequisites

Ensure you have the following installed on your machine:
- [Go](https://golang.org/doc/install) (1.16 or higher)
- [Docker](https://www.docker.com/get-started)
- [Node.js](https://nodejs.org/en/) and [npm](https://www.npmjs.com/get-npm)
- [Hardhat](https://hardhat.org/getting-started/)
- [IPFS](https://docs.ipfs.io/install/) or run an IPFS node in Docker

## Project Structure

```
.
├── README.md
├── fileregistry
│   └── fileregistry.go
├── contracts
│   └── FileRegistry.sol
├── go.mod
├── go.sum
├── handlers
│   └── fileregistryhttp.go
├── main.go
├── server
│   ├── router.go
│   └── server.go
└── service
    ├── ethclient.go
    └── ipfsupload.go
```

---

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/file-registry.git
cd file-registry
```

### 2. Install Dependencies

#### Smart Contracts - Create ABIs and Binaries

```bash
solc --abi --bin contracts/FileRegistry.sol -o ./build 
```

#### Smart Contracts - Created compile gen-go code for contract interaction using:

```bash
abigen --abi=build/FileRegistry.abi --bin=build/FileRegistry.bin --pkg=fileRegistry --out=FileRegistry.go
```

#### Go Dependencies

```bash
go mod tidy
```

### 3. Set Up Configuration

Update the config.json file with the following details:
- RPC_URL: The RPC URL of your network (e.g., Infura, Alchemy, or a local node).
- CONTRACT_ADDRESS: The address of the deployed smart contract.
- PRIVATE_KEY: The private key of the wallet used for signing transactions (keep this secure and do not share it).

### 4. Running the IPFS Node (Docker)

The project is configured to run an IPFS node using Docker. Start the IPFS node and Go API service with:

```bash
docker-compose up
```

This will start:
- **IPFS node** at `localhost:5001`
- **Go API** at `localhost:8081`

### 6. Running the Go API Locally (Optional)

You can also run the Go API locally without Docker:

```bash
go run main.go
```

---

## API Endpoints

### POST `/v1/files`

Upload a file to IPFS and store the CID in the smart contract.

- **URL**: `/v1/files`
- **Method**: `POST`
- **Form Data**:
  - `file`: The file to be uploaded
  - `filePath`: The path to associate with the file

#### Example `curl`:

```bash
curl -X POST http://localhost:8081/v1/files \
  -F "file=@/path/to/your/file.txt" \
  -F "filePath=file.txt"
```

### GET `/v1/files`

Fetch the CID for a file by its file path.

- **URL**: `/v1/files?filePath=<filePath>`
- **Method**: `GET`
- **Query Param**: `filePath`

#### Example `curl`:

```bash
curl "http://localhost:8081/v1/files?filePath=file.txt"
```

---

## Running Tests

### Testing the API

You can manually test the API using `curl` or [Postman](https://www.postman.com/).

---

## Troubleshooting

- **IPFS Node Issues**: Ensure the IPFS node is running and accessible.
- **Ethereum Issues**: Make sure you have a valid private key and sufficient funds on the network you're deploying to.
- **API Errors**: Check logs in Docker or the Go application for detailed error messages.

---

## License

This project is licensed under the MIT License.

---

## Resources

- [IPFS Documentation](https://docs.ipfs.io/)
- [Hardhat Documentation](https://hardhat.org/getting-started/)
- [Go Documentation](https://golang.org/doc/)
