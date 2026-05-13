Delta Engine API:

quantitative trading platform
    ingest live market data
    generate trading signals
    simulate paper trade execution
    track PnL, exposure, and risk
    evaluate strategy performance
    backtest strategies on historical data
    optionally execute real trades later


Frontend/UI
    -> REST/WebSocket
    -> Go API Layer
        -> Market Data Engine
        -> Strategy Engine
        -> Paper Trading Engine
        -> Risk & Analytics Engine
    -> PostgreSQL DB


Market Data Engine:
GET  /market/quote/:ticker
GET  /market/history/:ticker
GET  /market/watchlist
GET  /market/candles/:ticker

Responsibilities:
    fetch live market data
    cache quote data
    normalize OHLCV candles
    stream price updates
    store historical data


Strategy Engine:
POST /strategy/create
POST /strategy/start
POST /strategy/stop
GET  /strategy/results
GET  /strategy/signals

Responsibilities:
    evaluate trading indicators
    generate BUY/SELL signals
    support pluggable strategies
    run automated strategy loops
    calculate signal confidence

Strategies:
    SMA crossover
    RSI mean reversion
    momentum trading
    volatility breakout


Risk & Analytics Engine:
GET /risk/exposure
GET /risk/drawdown
GET /analytics/sharpe
GET /analytics/performance

Responsibilities:
    calculate Sharpe ratio
    calculate max drawdown
    track win/loss rate
    compute profit factor
    evaluate portfolio volatility


Later Development System Features:
    Scripted Automation for Long Term Trading Simulation
    JWT authentication
    websocket live streaming
    goroutine worker pools
    Redis caching
    structured logging
    Prometheus metrics
    Docker deployment
    Swagger/OpenAPI docs


DB Schema:

users:
    id
    email
    password_hash
    created_at

portfolios:
    id
    user_id
    cash_balance
    created_at

positions:
    id
    portfolio_id
    ticker
    quantity
    avg_cost
    unrealized_pnl
    updated_at

paper_trades:
    id
    portfolio_id
    ticker
    side
    quantity
    price
    fees
    executed_at

market_data:
    id
    ticker
    timestamp
    open
    high
    low
    close
    volume

signals:
    id
    strategy_name
    ticker
    signal_type
    confidence
    timestamp

strategy_results:
    id
    strategy_name
    sharpe_ratio
    max_drawdown
    win_rate
    total_return
    created_at