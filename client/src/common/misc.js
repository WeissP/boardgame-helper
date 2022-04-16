const objectMap = (obj, fn) =>
    Object.fromEntries(
        Object.entries(obj).map(
            ([k, v]) => [k, fn(v, k)]
        )
    )

export { objectMap }
