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

print(#s)

--打印不出来
for i, l in ipairs(s) do
    print(i, l)
end

--table.remove(s) 不行

ss={{1,2,3},1,2,3}
print(#ss)
for i,l in pairs(ss) do
    print(i,l)
end
--有序可以
for i, l in ipairs(ss) do
    print(i,l)
end
