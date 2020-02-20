start = Date.now()

var numbers = [1,2,3,4,5];

var result = numbers.every(function(value,index,arrays) {
    return value < 10;
});
console.log(result)

result = numbers.filter(function(value,index,arrays) {
    return (value <= 2);
});
console.log(result)

result = numbers.map(function(value,index,arrays) {
    return value + 2;
});
console.log(result)

result = numbers.reduce(function(prev, cur, index, array) {
    return prev + cur;
});
console.log(result)

var pattern1 = /at/g
console.log(pattern1.test("atb"))
console.log(pattern1.test("atb"))// false, 因为pattern1的索引已经到末尾了，需要新建一个pattern
console.log(pattern1.test("ab"))
console.log(pattern1.global)
console.log(pattern1.ignoreCase)
console.log(pattern1.multiline)
console.log(pattern1.lastIndex)
console.log(pattern1.source)

var pattern2 = /(ab).(cd)/gi
var matchs = pattern2.exec("abocd")
if (matchs != null) {
    console.log(matchs.index)
    console.log(matchs[0])
    console.log(matchs[1])
    console.log(matchs[2])
}

function sum(a1, a2) {
    return a1 + a2;
}

console.log("apply resutl:" + sum.apply(this, [1,2]))
console.log("call resutl:" + sum.call(this, 1,2))

window={}
window.color = "red"
var o ={color : "blue"}
function sayColor() {
    console.log(this.color)
}
sayColor()
sayColor.call(this)
sayColor.call(window)
sayColor.call(o)
var funobj = sayColor.bind(o)
funobj()

var s1 = "abcedgf"
console.log(s1.substring(2))
console.log(s1.substr(2))
 
end = Date.now()
console.log(end - start) //毫秒数