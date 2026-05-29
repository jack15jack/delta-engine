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
        routes.go (RegisterAnalyticsRoutes, no URLS to call yet)
        handler.go (empty)
        service.go (empty)
    /db
        /data
            delta.db
        /repos
            signals.go (NewSignalRepo, Insert)
        sqlite.go (NewPostgres)
    /backtest
        engine.go (empty)
        replay.go (empty)
        simulator.go (empty)

/scripts
/tests
    