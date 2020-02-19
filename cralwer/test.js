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