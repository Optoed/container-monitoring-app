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
                const last_ping_time = new Date(result.last_ping_time).toLocaleString();
                const ping_duration = result.ping_duration / 1e6; //перевод в ms из наносекунд
                return {
                    ...result,
                    last_ping_time,
                    ping_duration,
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