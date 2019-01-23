require "computelib"

--应税全价
yingshui_quanjia_cofig = {
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
        {M = "math.abs(division('L',1.03) * 0.03)"},
        {N = "get_str('附加税=应缴增值税*附加税率')"},
    },
    fenbieruls={
        {ABC = "get_str('分别申报合计:')"},
        {D = "sum(col('D'))"},
        {E = "sum(col('E'))"},
        {F = "sum(col('F'))"},
        {G = "sum(col('G'))"},
        {H = "sum(col('H'))"},
        {I = "sum(col('I'))"},
        {J = "sum(col('J'))"},
        {K = "sum(col('K'))"},
        {L = "sum(col('L'))"},
        {M = "sum(col('M'))"},
        {N = "sum(col('N'))"},
    }
}

--应税差价
yingshui_chajia_cofig = {
	sheetname = "平台展示",  			 --每个需要处理的文件的sheet名称
    startline = "6", 					--每个需要处理的文件的数据记录开始行
    huizongruls={
        {ABC = "get_str('汇总申报合计:')"},
        {D = "sum(col('D'))"},
        {E = "sum(col('E'))"},
        {F = "sum(col('F'))"},
        {G = "if_zhengshu_0('E')"},
        {H = "sum({'F','G'})"},
        {I = "max_0(min('E','H'))"},
        {J = "minus('H','I')"},
        {K = "sum({'D', max_0('E')})/sum({1 + 0.03})"},
        {L = "minus(sum({'D' , max_0('E')}), 'I')"},
        {M = "math.abs(division('L',1.03) * 0.03)"},
        {N = "get_str('附加税=应缴增值税*附加税率')"},
    },
    fenbieruls={
        {ABC = "get_str('分别申报合计:')"},
        {D = "sum(col('D'))"},
        {E = "sum(col('E', function(a) return a>0 end))"},
        {F = "sum(col('F'))"},
        {G = "sum(col('G'))"},
        {H = "sum(col('H'))"},
        {I = "sum(col('I'))"},
        {J = "sum(col('J'))"},
        {K = "sum(col('K'))"},
        {L = "sum(col('L'))"},
        {M = "sum(col('M'))"},
        {N = "sum(col('N'))"},
    }
}

--免税全价
mianshui_quanjia_cofig = {
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
        {K = "sum({'D', max_0('E')})"},
        {L = "minus('K', 'I')"},
    },
    fenbieruls={
        {ABC = "get_str('分别申报合计:')"},
        {D = "sum(col('D'))"},
        {E = "sum(col('E'))"},
        {F = "sum(col('F'))"},
        {G = "sum(col('G'))"},
        {H = "sum(col('H'))"},
        {I = "sum(col('I'))"},
        {J = "sum(col('J'))"},
        {K = "sum(col('K'))"},
        {L = "sum(col('L'))"},
    }
}

--免税差价
mianshui_chajia_cofig = {
	sheetname = "平台展示",  			 --每个需要处理的文件的sheet名称
    startline = "6", 					--每个需要处理的文件的数据记录开始行
    huizongruls={
        {ABC = "get_str('汇总申报合计:')"},
        {D = "sum(col('D'))"},
        {E = "sum(col('E'))"},
        {F = "sum(col('F'))"},
        {G = "sum(col('E'))"},
        {H = "sum({'F','G'})"},
        {I = "max_0(min('E','H'))"},
        {J = "minus('H','I')"},
        {K = "sum({'D', max_0('E')})"},
        {L = "minus('K', 'I')"},
    },
    fenbieruls={
        {ABC = "get_str('分别申报合计:')"},
        {D = "sum(col('D'))"},
        {E = "sum(col('E', function(a) return a>0 end))"},
        {F = "sum(col('F'))"},
        {G = "sum(col('G'))"},
        {H = "sum(col('H'))"},
        {I = "sum(col('I'))"},
        {J = "sum(col('J'))"},
        {K = "sum(col('K'))"},
        {L = "sum(col('L'))"},
    }
}

--免税全价明细
mianshui_quanjia_cofig = {
	sheetname = "平台展示",  			 --每个需要处理的文件的sheet名称
    startline = "6", 					--每个需要处理的文件的数据记录开始行
    computeTimes = "gettable_groupby('E')",
    huizongruls={
        {ABC = "get_str('汇总申报合计:')"},
        {E = "i"},
        {D = "after('E','D'))"},
        {F = "after('E','F'))"},
        {G = "sum(after_col('E','G')))"},
        {H = "sum(after_col('E','H')))"},
        {I = "min('G','H')"},
        {J = "minus('G','I')"},
        {K = "max_0('J')*0.03"},
    },
    fenbieruls={
        {ABC = "get_str('分别申报合计:')"},
        {E = "i"},
        {D = "after('E','D'))"},
        {F = "after('E','F'))"},
        {G = "sum(after_col('E','G')))"},
        {H = "sum(after_col('E','H')))"},
        {I = "sum(after_col('E','I'))"},
        {J = "sum(after_col('E','J'))"},
        {K = "sum(after_col('E','K'))"},
    }
}


p = "" 
maxrows=0
nowrows=0


function computeruls(ruls,p, computeTimesF)
    if computeTimes ~= nil then
        f = dostring(computeTimes)
        times = f()
    else
        times = function () 
            local t={1}
            local i=0
            return function()
                i=i+1
                return t[1]
            end
        end
    end

    for i in times() do 
        for k, v in pairs(ruls) do 
            for kk, vv in pairs(v) do 
                f = loadstring("result = "..vv)
                f()
                if type(result) ~= "string" then
                    p:set(kk..tostring(nowrows), string.format("%0.2f", result))
                else
                    p:set(kk..tostring(nowrows), result)
                end
            end
        end
        nowrows = nowrows + 1
    end
end

function computefile(filename,sheetname)
    print(filename)
    p = "" 
    maxrows=0
    nowrows=0

    p = datarows.new(filename, sheetname)

    if string.find(filename,"全价_应税台账汇总表") ~= nil then 
        config = yingshui_quanjia_cofig
    elseif string.find(filename, "全价_免税台账汇总表") ~= nil then
        config = yingshui_chajia_cofig
    elseif string.find(filename, "差价_应税台账汇总表") ~= nil then
        config = mianshui_chajia_cofig
    elseif string.find(filename, "差价_应税台账汇总表") ~= nil then
        config = mianshui_chajia_cofig
    end

    maxrows = get_endline(config)
    nowrows = maxrows + 2 --结果隔开两行
    computeruls(config.huizongruls, p, config.computeTimes)
    computeruls(config.fenbieruls, p, config.computeTimes)

    p:save()
end