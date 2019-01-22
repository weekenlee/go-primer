require "computelib"

--应税全价
yingshui_quanjia_cofig = {
 	filedir="./应税申报表_全价收入_生成xlsxs/",	  --需要处理的文件所在目录
	sheetname = "平台展示",  			 --每个需要处理的文件的sheet名称
    startline = "6", 					--每个需要处理的文件的数据记录开始行
    huizongruls={
        {ABC = "get_str('汇总申报合计:')"},
        {D = "sum(col('D'))"},
        {E = "sum(col('E'))"},
        {F = "sum(col('F'))"},
        {G = "sum(col('G'))"},
        {H = "sum({'F','G'})"},
        {I = "max_0(min('E','H'))"},
        {J = "minus('H','I')"},
        {K = "sum({'D', max_0('E')})/sum({1 + 0.03})"},
        {L = "minus(sum({'D' , max_0('E')}), 'I')"},
        {M = "sum(col('M'))"},
        {N = "get_str('附加税=应缴增值税*附加税率')"},
    },
    fenbieruls={
        {ABC = "get_str('分别申报合计:')"},
        {D = "sum(col('D'))"},
        {E = "sum(col('E'))"},
        {F = "sum(col('F'))"},
        {G = "sum(col('G'))"},
        {H = "sum({'F','G'})"},
        {I = "max_0(min('E','H'))"},
        {J = "minus('H','I')"},
        {K = "sum({'D', max_0('E')})/sum({1 + 0.03})"},
        {L = "minus(sum({'D' , max_0('E')}), 'I')"},
        {M = "sum(col('M'))"},
        {N = "sum(col('N'))"},
    }
}

p = "" 
maxrows=0
nowrows=0

function get_str(str)
    return str 
end 

function get_endline(ruls)
    endline = tonumber(ruls)
    value = p:get('A'..tostring(endline))
    while (string.len(value) > 0)  do 
        endline = endline+1
        value = p:get("A"..endline)
    end
    return endline 
end

--获取一列
function col(C)
    startline = tonumber(yingshui_quanjia_cofig.startline)
    tab = {}
    for i = startline, maxrows-1 do
        table.insert(tab,tonumber(p:get(C..tostring(i))))
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

function computeruls(ruls,p)
    for k, v in pairs(ruls) do 
        for kk, vv in pairs(v) do 
            f = loadstring("result = "..vv)
            f()
            if type(result) ~= "string" then
                p:set(kk..tostring(nowrows), string.format("%0.2f", result))
            else
                print(kk..tostring(nowrows), result)
                p:set(kk..tostring(nowrows), result)
            end
        end
    end
    nowrows = nowrows + 1
end

function computefile(filename,sheetname)
    p = "" 
    maxrows=0
    nowrows=0

    p = datarows.new(filename, sheetname)
    maxrows = get_endline(yingshui_quanjia_cofig.huizongruls)

    nowrows = maxrows + 2
    computeruls(yingshui_quanjia_cofig.huizongruls, p)
    computeruls(yingshui_quanjia_cofig.fenbieruls, p)

    p:save()
end