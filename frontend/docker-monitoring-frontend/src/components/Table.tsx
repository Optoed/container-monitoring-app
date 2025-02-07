import React from 'react';
import { Table } from 'antd';
import { ContainerPingResult } from '../models/ContainerPingResult';

interface TableContainersProps {
    data: ContainerPingResult[];
}

const columns = [
    {title: 'ID', dataIndex: 'id', key: 'id'},
    { title: 'IP', dataIndex: 'ip', key: 'ip' },
    { title: 'Status', dataIndex: 'status', key: 'status' },
    { title: 'Last Ping Time', dataIndex: 'lastPingTime', key: 'lastPingTime' },
    { title: 'Ping Duration', dataIndex: 'pingDuration', key: 'pingDuration'},
];

const ContainersPingTable: React.FC<TableContainersProps> = ({ data }) => {
    return <Table dataSource={data} columns={columns} />;
};

export default ContainersPingTable;