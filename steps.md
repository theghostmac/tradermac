You've made significant progress in building the foundational elements of your prop trading firm's codebase, covering options modeling, trade execution, risk management, and market data integration. To progress towards deploying your system in a live environment, consider the following steps:

1. **Implement a Live Data Feed**:
    - Integrate a live market data feed. This is crucial for real-time trading. You've already set up AlphaVantage for historical data, but you'll need a reliable source for live data. Look for APIs that offer real-time stock, options, and commodities data.
    - Ensure your system can process this data efficiently and update your trading algorithms in real time.

2. **Develop and Test Trading Algorithms**:
    - Use the infrastructure you've built to develop specific trading algorithms. These could be based on various strategies like statistical arbitrage, mean reversion, momentum trading, etc.
    - Backtest these algorithms with historical data to evaluate their performance. Ensure your backtesting is as realistic as possible, including factors like transaction costs and market impact.

3. **Complete the Order Execution System**:
    - Your code for executing trades should be robust and capable of handling different market conditions.
    - Implement logic to handle partial fills, slippage, and rejections.
    - Ensure the system can manage multiple orders simultaneously and efficiently.

4. **Risk Management Enhancements**:
    - Expand your risk management system to include real-time monitoring of market risk, credit risk, operational risk, etc.
    - Implement automated triggers to reduce positions or halt trading if certain risk thresholds are exceeded.

5. **Compliance and Reporting**:
    - Ensure your system is compliant with relevant trading regulations. This includes implementing proper trade reporting, maintaining logs, and ensuring data privacy and security.
    - Develop a reporting module that can generate performance reports, risk exposure reports, and other relevant documentation.

6. **User Interface (UI) Development**:
    - If you haven’t already, develop a user-friendly interface for monitoring and controlling the system. This UI should provide real-time insights into trading performance, risk metrics, and system status.

7. **Simulated Live Trading (Paper Trading)**:
    - Before going live, simulate trading with real-time data without using actual capital. This will help you identify any issues in a risk-free environment.

8. **Infrastructure and Deployment Readiness**:
    - Ensure your infrastructure is capable of handling the computational load. This includes server capabilities, data storage, and network latency optimizations.
    - Plan for deployment, which includes setting up a production environment, ensuring robustness, and having a rollback plan in case of failures.

9. **Monitoring and Maintenance**:
    - Develop tools for monitoring system performance and health in real-time.
    - Plan for regular maintenance, updates, and bug fixes.

10. **Legal and Financial Consulting**:
- Before going live, consult with legal and financial experts to ensure all regulatory requirements are met, and your business model is sound.

By following these steps, you'll move closer to a live deployment. Remember, trading involves significant risk, especially with a newly developed system. Start with a small capital base and scale gradually as you gain confidence in your system's performance.