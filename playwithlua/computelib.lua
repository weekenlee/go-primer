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

function temp() 
    local tab={1,2}
    local i=0
    return function()
        i=i+1
        return tab[1]
    end
end

function groupby(c)
    local startline = tonumber(mianshui_quanjia_config.startline)
    local tab = {}
    for i = startline, maxrows-1 do
        value = tonumber(p:get(c..tostring(i)))
        tab[value] = true
    end

    local array={}
    for k, v in pairs(tab) do
        table.insert(array,k)
        print(k, v)
    end

    i = 0
    return function()
        i = i + 1
        print(array[i])
        return array[i]
    end
end

function after(srcc, desc)
    valuesrc = p:get(srcc..tostring(nowrows))
    startline = tonumber(yingshui_quanjia_config.startline)
    for i = startline, maxrows-1 do
        value = p:get(srcc..tostring(i))
        if value == valuesrc then
            return p:get(desc..tostring(i))
        end
    end
    return nil
end

--获取一列,如果srcc列,与nowrows的srcc列值相同
function after_col(srcc, desc)
    valuesrc = tonumber(p:get(srcc..tostring(nowrows)))
    startline = tonumber(yingshui_quanjia_config.startline)
    tab = {}
    for i = startline, maxrows-1 do
        value = tonumber(p:get(srcc..tostring(i)))
        valuedes = tonumber(p:get(desc..tostring(i)))
        if value==valuesrc then 
            table.insert(tab,valuedes)
        end
    end
    return tab
end

--获取一列
function col(C, f)
    startline = tonumber(yingshui_quanjia_config.startline)
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