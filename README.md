# Bitcoin SV Transaction Creator & Fetcher (Go)

> A comprehensive example demonstrating Bitcoin SV transaction creation and retrieval using Go

[![Go](https://img.shields.io/badge/Go-1.14+-blue.svg)](https://golang.org)
[![BSV](https://img.shields.io/badge/Bitcoin%20SV-Compatible-orange.svg)](https://bitcoinsv.com)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## 🚀 Overview

This Go application demonstrates how to create and send Bitcoin SV transactions programmatically. It provides a practical example of working with the Tokenized SDK for transaction building and RPC communication.

## ✨ Features

- **Transaction Creation**: Build Bitcoin SV transactions with inputs and outputs
- **Transaction Retrieval**: Fetch transaction details via RPC
- **RPC Integration**: Connect to Bitcoin SV node for blockchain operations
- **Key Management**: Handle private keys and address generation
- **Fee Calculation**: Automatic fee calculation based on transaction size

## 🛠️ Prerequisites

- Go 1.14 or higher
- Access to a Bitcoin SV RPC node
- Basic understanding of Bitcoin transactions

## 📦 Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd create-and-send-tx-go
   ```

2. **Install dependencies**
   ```bash
   go get ./...
   ```

3. **Build the application**
   ```bash
   go build
   ```

## ⚙️ Configuration

Create or edit the `.env` file in the project root:

```env
# Bitcoin SV RPC Configuration
RPC_HOST="your-rpc-hostname:port"
RPC_USERNAME="your-rpc-username"
RPC_PASSWORD="your-rpc-password"

# Transaction to query with 'get' command
TX_ID="transaction-id-to-fetch"
```

### Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `RPC_HOST` | Bitcoin SV node RPC endpoint | `localhost:8332` |
| `RPC_USERNAME` | RPC authentication username | `bitcoin` |
| `RPC_PASSWORD` | RPC authentication password | `your-password` |
| `TX_ID` | Transaction ID for retrieval | `699eba6bd9f9fd6a...` |

## 🎯 Usage

### Available Commands

**Create a new transaction:**
```bash
./create-and-send-tx-go create
```

**Fetch transaction details:**
```bash
./create-and-send-tx-go get
```

### Example Output

**Transaction Creation:**
```
Building transaction with:
- Input: 699eba6bd9f9fd6a82dcb4ae69d756003a6b1356e4d2a2eefa5d8ca916b71588:0
- Output: 1000 satoshis to mmdaHdaiSqzu9AL1DjnBR7NgUcBKZMzwGZ
- Change: Remaining balance to mmsio7EbL2ZwJConqaUovnSYFEc8vxnbuf
```

**Transaction Retrieval:**
```
Connect to localhost:8332 as bitcoin
Get Tx: 699eba6bd9f9fd6a82dcb4ae69d756003a6b1356e4d2a2eefa5d8ca916b71588
Tx: 699eba6bd9f9fd6a82dcb4ae69d756003a6b1356e4d2a2eefa5d8ca916b71588
[Transaction Details...]
```

## 🏗️ Architecture

### Core Components

```
create-and-send-tx-go/
├── main.go           # CLI entry point and command routing
├── createTx.go       # Transaction building and signing
├── getTX.go          # Transaction retrieval via RPC
├── .env             # Configuration file
└── README.md        # This documentation
```

### Dependencies

- **[tokenized/pkg/bitcoin](https://github.com/tokenized/pkg)**: Bitcoin SV transaction handling
- **[tokenized/pkg/txbuilder](https://github.com/tokenized/pkg)**: Transaction builder utilities
- **[tokenized/pkg/rpcnode](https://github.com/tokenized/pkg)**: RPC node communication
- **[joho/godotenv](https://github.com/joho/godotenv)**: Environment variable loading

## 🔧 Customization

### Modifying Transaction Parameters

Edit `createTx.go` to customize:

```go
// Change dust limit and fee rate
builder := txbuilder.NewTxBuilder(512, 1.1)

// Modify input parameters
value := uint64(5000000000) // 50 BSV in satoshis

// Adjust output amount
_ = builder.AddPaymentOutput(paymentAddress, 1000, false) // 1000 satoshis
```

### Adding Multiple Outputs

```go
// Add multiple payment outputs
_ = builder.AddPaymentOutput(address1, 1000, false)
_ = builder.AddPaymentOutput(address2, 2000, false)
_ = builder.AddPaymentOutput(address3, 3000, false)
```

## 🛡️ Security Considerations

- **Private Keys**: Never commit private keys to version control
- **RPC Credentials**: Store RPC credentials securely
- **Environment Variables**: Use `.env` files for sensitive configuration
- **Network**: Ensure RPC connections are properly secured

## 📚 Learning Resources

- **Bitcoin SV Documentation**: [BitcoinSV.com](https://bitcoinsv.com)
- **Tokenized SDK**: [GitHub Repository](https://github.com/tokenized/pkg)
- **Go Bitcoin Resources**: [Go Bitcoin Development](https://github.com/btcsuite)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is available under the MIT License. See LICENSE file for details.

---

**Note**: This is an educational example for Bitcoin SV development. Always test thoroughly on testnet before using with real funds.

