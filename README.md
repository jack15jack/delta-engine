Delta Engine API:

quantitative trading platform
    ingest live market data
    generate trading signals
    simulate trade execution
    track PnL, exposure, and risk
    evaluate strategy performance
    backtest strategies on historical data
    


Frontend/UI
    -> REST/WebSocket
    -> Go API Layer
        -> Market Data Engine
        -> Strategy Engine
        -> Trading Engine
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


How it will work:
    Market data comes in through GET market/quote/:ticker or a database of historical market data
    Strategy engine reads the market data
    Indicators update
    BUY/SELL/HOLD signal created
    Signal stored in DB
    Trading Engine executes BUY/SELL
    Risk Engine updates metrics
    Over the course of a backtest, PnL will be recorded to view profitability of the system    

DB Schema:
    quotes
    candles
    strategies
    signals
    trades
    positions
    performance_metrics
    users


For me:
    routes are what URL calls what code
    middleware is filters and utils for logging/auth
    handlers are HTTP request translators
    market/ is for market functionality
    strategy/ is for strategy functionality
    analytics/ is for risk, volatility, and winrate calculations
    models/ is for shared data structures
    db/ is for database interaction
    
