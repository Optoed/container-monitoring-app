// src/App.tsx
import React, { useEffect, useState } from 'react';
import { getPingResults } from './api/api';
import ContainersPingTable from './components/Table';

const App: React.FC = () => {
    const [data, setData] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            const results = await getPingResults();
            console.log("results = ", results)

            const transformedResults = results.map((result: any) => {
                const lastPingTime = new Date(result.lastPingTime).toLocaleString();
                return {
                    ...result,
                    lastPingTime,
                };
            });
            setData(transformedResults);
            console.log("transformedResults = ", transformedResults)
        };
        fetchData();

        const interval = setInterval(fetchData, 10000); // Обновление каждые 10 секунд
        return () => clearInterval(interval);
    }, []);

    return (
        <div style={{ padding: '20px' }}>
            <h1>Ping Results</h1>
            <ContainersPingTable data={data} />
        </div>
    );
};

export default App;