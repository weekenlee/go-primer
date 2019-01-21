function add(a , b)
    return a + b
end

function min(a , b)
    return a - b
end

function sum(...)
    s = 0
    local arg = {...}
    for k, v in ipairs(arg) do
        s = s + v 
    end
    return s
end


computestr = "return sum(add(1,2), min(6, sum(1,2,1)))"
print(sum(add(1,2), min(6, sum(1,2,1))))
print(loadstring(computestr)())
