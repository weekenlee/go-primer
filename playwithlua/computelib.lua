function get_str(str)
    return str 
end 

function get_endline(config)
    endline = tonumber(config.startline)
    value = p:get('A'..tostring(endline))
    while (string.len(value) > 0)  do 
        endline = endline+1
        value = p:get("A"..endline)
    end
    return endline 
end


function groupby(c)
    startline = tonumber(yingshui_quanjia_cofig.startline)
    tab = {}
    i = 0
    for i = startline, maxrows-1 do
        value = tonumber(p:get(C..tostring(i)))
        table.insert(tab,value)
    end

    return function()
        i = i + 1
        return tab[i]
    end
end

--获取一列
function col(C, f)
    startline = tonumber(yingshui_quanjia_cofig.startline)
    tab = {}
    for i = startline, maxrows-1 do
        value = tonumber(p:get(C..tostring(i)))
        if f ~= nil then
            r = f(value)
        else
            r =  true
        end
        if r then 
            table.insert(tab,value)
        end
    end
    return tab
end

--获取一个值
function get_real_value(a)
    if type(a) == "string" then 
        return tonumber(p:get(a..tostring(nowrows)))
    end
    return a
end

function sum(t)
    local arg=t
    local s = 0
    for k, v in ipairs(arg) do
        s = s + get_real_value(v)
    end
    return s
end

function division(a, b)
    avalue = get_real_value(a)
    bvalue = get_real_value(b)
    return avalue/bvalue
end

function minus(a, b)
    avalue = get_real_value(a) 
    bvalue = get_real_value(b) 
    return avalue-bvalue
end

function max_0(C)
    cvalue = get_real_value(C)
    return math.max(cvalue, 0)
end

function min(a, b)
    avalue = get_real_value(a)
    bvalue = get_real_value(b)
    return math.min(avalue,bvalue)
end

function if_zhengshu_0(C) 
    value = get_real_value(C)
    if value > 0 then 
        return 0 
    else 
        return math.abs(value) 
    end 
end