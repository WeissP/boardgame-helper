
const axios = require('axios')

axios.get('http://192.168.8.143:8888/api/save').then(resp => {
    console.log(resp.data)
})
