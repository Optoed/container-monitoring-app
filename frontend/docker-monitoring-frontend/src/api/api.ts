import axios from 'axios';

// TODO: замени на 'http://backend:8080/containers' когда поместишь frontend в docker контейнер
// const API_URL = 'http://backend:8080/containers';
// Для теста на локальном компьютере:
const API_URL = 'http://localhost:8080/containers';

export const getPingResults = async () => {
    const response = await axios.get(API_URL);
    return response.data;
};