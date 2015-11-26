/**
 * Created by rick on 15/9/5.
 */


var wbSprintf = function (str) {
    var args = arguments,
        i = 1;

    str = str.replace(/%s/g, function () {
        var arg = args[i++];

        if (typeof arg === 'undefined') {
            return '';
        }
        return arg;
    });
    return str;
};

var wbToMoney = function (str){
    a = parseFloat(str)
    a = a.toFixed(2)
    return parseFloat(a)
}

var wbGetMapValue = function (map){
    lst = []
    for (k in map){
        lst.push(map[k])
    }
    return lst
}
