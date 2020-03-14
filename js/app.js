'use strict';
const morgan = require("./MorganAndaString");
//const morgan = new MorganAndaString;
var obj = new morgan();
console.log(obj.getMinLex("ABCDEFG", "JGFAFGF"));