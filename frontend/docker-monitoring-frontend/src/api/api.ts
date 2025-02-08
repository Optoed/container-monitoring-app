import axios from 'axios';

const API_URL = process.env.REACT_APP_BACKEND_URL || 'http://localhost:8080';
const BACKEND_URL = API_URL + "/containers"

export const getPingResults = async () => {
    // if (!API_URL) {
    //     throw new Error("Backend URL is not defined in the environment variables.");
    // }
    console.log("FULL_ROUTE_URL = ", BACKEND_URL)
    const response = await axios.get(BACKEND_URL);
    return response.data;
};