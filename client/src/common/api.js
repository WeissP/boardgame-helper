import axios from 'axios'

export function post(endpoint, data) {
    return axios.post('/api/' + endpoint, data)
}

export function get(endpoint, data) {
    return axios.get('/api/' + endpoint, { params: data })
}

