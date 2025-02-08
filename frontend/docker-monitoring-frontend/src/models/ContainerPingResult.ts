export interface ContainerPingResult {
    id: number;
    ip: string;
    status: string;
    last_ping_time: string;
    ping_duration: number;
}