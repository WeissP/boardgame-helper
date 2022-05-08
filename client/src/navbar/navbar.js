import React, { Component } from 'react'
import { Navbar, Nav, Dropdown } from 'rsuite'
import { useNavigate } from 'react-router-dom'
import EditIcon from '@rsuite/icons/Edit'
import DashboardIcon from '@rsuite/icons/Dashboard'

const NavBarInstance = ({ onSelect, activeKey, ...props }) => {
    const navigate = useNavigate()
    const routeChange = (path) => {
        navigate(path)
    }

    return (
        <Navbar {...props}>
            <Navbar.Brand href='/doudizhu-view'>Boardgame Helper</Navbar.Brand>
            <Nav onSelect={onSelect} activeKey={activeKey}>
                <Dropdown menuStyle={{ minWidth: 95 }} class='z-50' title='斗地主'>
                    <Dropdown.Item icon={<EditIcon />} eventKey='/doudizhu-input' onSelect={(key, _) => routeChange(key)}>记分</Dropdown.Item>
                    <Dropdown.Item icon={<DashboardIcon />} eventKey='/doudizhu-view' onSelect={(key, _) => routeChange(key)}>查分</Dropdown.Item>
                </Dropdown>
                <Dropdown menuStyle={{ minWidth: 95 }} class='z-50' title='日麻'>
                    <Dropdown.Item icon={<EditIcon />} eventKey='/jp-mahjong' onSelect={(key, _) => routeChange(key)}>记分</Dropdown.Item>
                </Dropdown>
            </Nav>
        </Navbar>
    )
}

const Header = () => {
    const [activeKey, setActiveKey] = React.useState(null)
    return (
        <div className='nav-wrapper' class='sticky'>
            <NavBarInstance appearance='inverse' activeKey={activeKey} onSelect={setActiveKey} />
        </div>
    )
}

export default Header
