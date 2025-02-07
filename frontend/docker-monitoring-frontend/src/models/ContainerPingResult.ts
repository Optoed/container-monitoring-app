export interface ContainerPingResult {
    id: number;
    ip: string;
    status: string;
    lastPingTime: string;
    pingDuration?: number;
}