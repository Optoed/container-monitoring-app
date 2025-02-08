import axios from 'axios';

// Для теста на локальном компьютере:
// const API_URL = 'http://localhost:8080/containers';

const API_URL = process.env.REACT_APP_BACKEND_URL;
const FULL_ROUTE_URL = API_URL + "/containers"
//const FULL_ROUTE_URL = 'http://backend:8080/containers';

export const getPingResults = async () => {
    if (!API_URL) {
        throw new Error("Backend URL is not defined in the environment variables.");
    }
    console.log("API_URL = ", API_URL)
    console.log("FULL_ROUTE_URL = ", FULL_ROUTE_URL)
    const response = await axios.get(FULL_ROUTE_URL);
    return response.data;
};