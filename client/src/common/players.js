import players from './players.json'

const m = new Map(Object.entries(players))

export function playersArray() {
    return Array.from(m, ([name, value]) => ([name, value]))
}

export function playerNames() {
    return Object.values(players)
}

export function playerIDs() {
    return Object.keys(players)
}

export function playerValue(id) {
    if (m.has(id)) {
        return m.get(id)
    }
    return "unknown"
}
