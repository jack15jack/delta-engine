# Delta Engine API

**Delta Engine** is a quantitative trading backend designed to ingest market data, generate algorithmic trading signals, simulate trade execution, and evaluate strategy performance.

The system provides a REST API for market data, strategy management, simulated trading, risk analysis, and performance analytics. It is designed around a modular architecture so new strategies, data sources, and analytics can be added without restructuring the core system.

---

## Architecture

```text
                         ┌──────────────────────┐
                         │    Frontend / UI     │
                         └──────────┬───────────┘
                                    │
                         REST / WebSocket API
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │      Go API Layer    │
                         └──────────┬───────────┘
                                    │
              ┌─────────────────────┼─────────────────────┐
              │                     │                     │
              ▼                     ▼                     ▼
      ┌───────────────┐     ┌───────────────┐     ┌───────────────┐
      │ Market Data   │     │   Strategy    │     │     Risk &    │
      │    Engine     │────▶│    Engine     │────▶│   Analytics   │
      └───────────────┘     └───────┬───────┘     └───────────────┘
              │                     │                     ▲
              │                     ▼                     │
              │              ┌───────────────┐            │
              └─────────────▶│    Trading    │────────────┘
                             │    Engine      │
                             └───────┬───────┘
                                     │
                                     ▼
                            ┌─────────────────┐
                            │   PostgreSQL    │
                            └─────────────────┘
```

### Core Pipeline

```text
Market Data
     │
     ▼
Market Data Engine
     │
     ▼
Strategy Engine
     │
     ├── Indicators
     │
     ├── Signal Generation
     │
     └── Signal Confidence
     │
     ▼
Trading Engine
     │
     ├── BUY
     ├── SELL
     └── HOLD
     │
     ▼
Risk & Analytics
     │
     ├── PnL
     ├── Exposure
     ├── Drawdown
     ├── Sharpe Ratio
     └── Performance Metrics
```

---

## Features

### Market Data Engine

The Market Data Engine handles ingestion, normalization, storage, and retrieval of market information.

**Responsibilities:**

* Fetch live market data
* Retrieve historical market data
* Normalize OHLCV candle data
* Cache frequently accessed quotes
* Store historical market data
* Provide market data to strategy workers

**Endpoints:**

| Method | Endpoint                  | Description                     |
| ------ | ------------------------- | ------------------------------- |
| `GET`  | `/market/quote/:ticker`   | Retrieve the latest quote       |
| `GET`  | `/market/history/:ticker` | Retrieve historical market data |
| `GET`  | `/market/watchlist`       | Retrieve configured watchlist   |
| `GET`  | `/market/candles/:ticker` | Retrieve OHLCV candles          |

---

### Strategy Engine

The Strategy Engine evaluates market data and converts indicator conditions into trading signals.

**Responsibilities:**

* Calculate technical indicators
* Generate `BUY`, `SELL`, and `HOLD` signals
* Calculate signal confidence
* Run automated strategy loops
* Store generated signals
* Support pluggable trading strategies
* Evaluate strategy performance

**Endpoints:**

| Method | Endpoint            | Description                |
| ------ | ------------------- | -------------------------- |
| `POST` | `/strategy/create`  | Create a trading strategy  |
| `POST` | `/strategy/start`   | Start a strategy           |
| `POST` | `/strategy/stop`    | Stop a strategy            |
| `GET`  | `/strategy/results` | Retrieve strategy results  |
| `GET`  | `/strategy/signals` | Retrieve generated signals |

### Supported Strategies

The architecture is designed to support interchangeable strategy implementations, including:

* **SMA Crossover**
* **RSI Mean Reversion**
* **Momentum Trading**
* **Volatility Breakout**

---

### Trading Engine

The Trading Engine converts strategy signals into simulated trades and maintains the state of the trading system.

**Responsibilities:**

* Process trading signals
* Simulate order execution
* Track open positions
* Record completed trades
* Calculate realized and unrealized PnL
* Maintain portfolio state

This provides the execution layer between strategy generation and portfolio analytics.

---

### Risk & Analytics Engine

The Risk & Analytics Engine evaluates the performance and risk characteristics of a strategy or simulated portfolio.

**Endpoints:**

| Method | Endpoint                 | Description                  |
| ------ | ------------------------ | ---------------------------- |
| `GET`  | `/risk/exposure`         | Calculate portfolio exposure |
| `GET`  | `/risk/drawdown`         | Calculate maximum drawdown   |
| `GET`  | `/analytics/sharpe`      | Calculate Sharpe ratio       |
| `GET`  | `/analytics/performance` | Retrieve performance metrics |

**Metrics include:**

* Profit & Loss
* Portfolio exposure
* Maximum drawdown
* Sharpe ratio
* Win/loss rate
* Profit factor
* Portfolio volatility
* Strategy performance

---

## Backtesting

Delta Engine is designed to support historical strategy evaluation through a simulated trading pipeline.

```text
Historical Market Data
          │
          ▼
   Strategy Evaluation
          │
          ▼
    Signal Generation
          │
          ▼
   Simulated Execution
          │
          ▼
      Trade History
          │
          ▼
     PnL Calculation
          │
          ▼
   Risk & Performance
       Analytics
```

During a backtest, historical market data is processed chronologically. Strategies generate signals based on the available data, simulated trades are executed, and the resulting portfolio state is recorded for analysis.

This allows strategies to be compared using both **return and risk metrics** rather than profitability alone.

---

## Database

PostgreSQL is used for persistent storage.

### Core Tables

```text
quotes
candles
strategies
signals
trades
positions
performance_metrics
users
```

The database separates raw market data from strategy state, generated signals, trade execution, and performance results.

---

## Project Structure

The backend is organized by responsibility:

```text
/
├── routes/
│   └── API route definitions
│
├── middleware/
│   └── Request filters and middleware
│
├── handlers/
│   └── HTTP request/response translation
│
├── market/
│   └── Market data functionality
│
├── strategy/
│   └── Strategy execution and signal generation
│
├── analytics/
│   └── Risk, volatility, and performance calculations
│
├── models/
│   └── Shared data structures
│
├── db/
│   └── Database interaction
│
└── utils/
    └── Shared utilities such as logging and authentication
```

This separation keeps HTTP handling, business logic, database access, and analytics independent from one another.

---

## Technology Stack

| Component               | Technology        |
| ----------------------- | ----------------- |
| API                     | Go                |
| Database                | PostgreSQL        |
| API Protocol            | REST              |
| Real-Time Communication | WebSocket         |
| Authentication          | JWT               |
| Caching                 | Redis             |
| Metrics                 | Prometheus        |
| Deployment              | Docker            |
| Documentation           | Swagger / OpenAPI |

---

## Design Goals

Delta Engine is built around several core design goals:

**Modularity**
Trading strategies, market-data providers, and analytics should be replaceable without changing unrelated components.

**Separation of Concerns**
API handling, trading logic, data persistence, and analytics are isolated into separate layers.

**Extensibility**
The strategy engine is designed to allow new algorithmic strategies to be added without modifying the core trading pipeline.

**Measurable Performance**
Strategies are evaluated using both profitability and risk metrics, allowing performance to be analyzed beyond simple returns.

**Realistic Simulation**
The trading engine provides a foundation for simulating the lifecycle of market data → signal → execution → position → performance.

---

## Project Status

Delta Engine is an actively developed quantitative trading platform. The current implementation focuses on establishing the backend architecture, API layer, market-data pipeline, strategy engine, and analytics infrastructure.

Future development will expand the system toward automated backtesting, real-time market streaming, portfolio management, and production-oriented observability.
