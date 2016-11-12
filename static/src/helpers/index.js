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
        let s = (index === 0 && selected) ? 'selected' : '';
        options.push(<option key={index} value={status} selected={s}>{status}</option>);
    });

    return options;
}
