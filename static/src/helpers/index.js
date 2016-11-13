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

        return false
    } else {
        let flag = false
        data.forEach((val, i) => {
            if (code === val.code) {
                flag = true
            }
        })

        return flag
    }
}
