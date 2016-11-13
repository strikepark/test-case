import React from 'react'

export function fmtCode(code) {
    let codeStr = code.toString();
    let codeStrLen = codeStr.length;

    return (codeStrLen < 8) ? ('0'.repeat(8 - codeStrLen) + codeStr) : codeStr;
}

export function genStatusSelect(selected) {
    let statusArr = ['Готовится', 'Доставляется', 'Доставлен'];

    let options = [];

    statusArr.map((status, index) => {
        options.push(<option key={index} value={status}>{status}</option>);
    });

    return options;
}

export function checkUser(code, data) {
    if (data === [] || data === null) {
        console.log(1);

        return false
    } else {
        console.log(2);

        data.forEach((val, i) => {
            console.log(code);
            console.log(val.code);

            if (code === val.code) {
                return true
            }
        })

        return false
    }
}
