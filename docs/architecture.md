/cmd
    main.go

/internal
    /api
        routes.go (RegisterRoutes)
        health.go (HealthCheck)
    /config
        config.go (LoadConfig)
    /middleware
        auth.go (Auth)
        logging.go (Logger)
        recovery.go (Recovery)
    /models
        candle.go (Candle struct)
        portfolio.go (Portfolio struct)
        position.go (Position struct)
        signal.go (Signal struct)
        trade.go (Trade struct)
    /market
        routes.go (RegisterMarketRoutes)
        handler.go (NewMarketHandler, GetQuote)
        candles.go (empty)
        finnhub.go (NewFinnhubProvider, GetQuote, GetHistory)
        provider.go (Provider interface)
        service.go (NewService, GetQuote, GetHistory)
        websocket.go (empty)
    /strategy
        /indicators
            sma.go (SMA)
            rsi.go (empty)
            macd.go (empty)
            ema.go (empty)
        /strategies
            momentum.go (empty)
            rsi.go (empty)
            sma_cross.go (Name, Evaluate)
        engine.go (NewEngine, Run)
        routes.go (RegisterStrategyRoutes)
        handler.go (NewStrategyHandler, Run)
        strategy.go (Strategy interface)
    /portfolio
        service.go (NewService, CreatePortfolio, GetPortfolio, GetPostions, ExecuteTrade)
        routes.go (RegisterPortfolioRoutes)
        handler.go (NewPortfolioHandler, CreatePortfolio, GetPortfolio, Postitions, Trade)
    /analytics
        routes.go (RegisterAnalyticsRoutes)
        handler.go (NewHandler, GetPortfolioPerformance)
        service.go (BuildSnapshot, GetPortfolioPerformance, GetExposure, GetRiskMetrics, GetTradeMetrics, GetPositionMetrics)
        analytics_models.go (PortfolioPerformance, ExposureMetrics, EquitySnapshot, PortfolioSnapshot, RiskMetrics, TradeMetrics, PositionMetrics)
    /backtest
        engine.go (NewEngine, Run, Results)
        replay.go (NewReplay, Next)
        simulator.go (NewSimulator, Buy, Sell, Snapshot)
        backtest_models.go (EquitySnapshot, SimulatedPosition, Results)
    /db
        /data
            delta.db
        /repos
            signals.go (NewSignalRepo, Insert)
        migrate.go (AutoMigrate)
        sqlite.go (NewSQLite)
    

/scripts
/tests
    