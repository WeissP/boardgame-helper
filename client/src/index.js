import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import App from './App'
import reportWebVitals from './reportWebVitals'
import {
    BrowserRouter,
    Routes,
    Route
} from 'react-router-dom'
import { DouDiZhuInputDft, DouDiZhuInputEdit } from './routes/doudizhuInput'
import DouDiZhuView from './routes/doudizhuView'

const container = document.getElementById('root')
const root = ReactDOM.createRoot(container)

root.render(<BrowserRouter>
        <Routes>
            <Route path='/' element={<DouDiZhuView />} />
            <Route path='doudizhu-input' element={<DouDiZhuInputDft />} />
            <Route path='doudizhu-edit/:timestamp' element={<DouDiZhuInputEdit />} />
            <Route path='doudizhu-view' element={<DouDiZhuView />} />
        </Routes>
    </BrowserRouter>)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
