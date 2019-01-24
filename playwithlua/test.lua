function Set(list)
    local set = {}
    for _, l in ipairs(list) do
        set[l] = true
    end
    return set
end

s = Set({'a', 'a', 'b'})

for l in pairs(s) do
    print(l)
end

for i, l in ipairs(s) do
    print(i, l)
end
