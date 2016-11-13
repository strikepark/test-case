import {
    GET_USER_ORDERS_REQUEST,
    GET_USER_ORDERS_SUCCESS,
    GET_USER_ORDERS_FAIL,

    SHOW_MANAGE,
    SHOW_LOGIN
} from '../constants/Login'

import { checkUser } from '../helpers'

import $ from 'jquery'

export function getUserOrders(userInfo) {
    return (dispatch) => {
        dispatch({
            type: GET_USER_ORDERS_REQUEST
        })

        let url = 'http://planadotest.herokuapp.com/api/orders/costumer/' + userInfo.PhoneNumber;

        $.get(url)
            .done(function(data) {
                if (!checkUser(parseInt(userInfo.Code), data)) {
                    alert('Покупателя с такими данными нет');

                    dispatch({
                        type: GET_USER_ORDERS_FAIL,
                        error: 'Покупателя с такими данными нет'
                    });
                } else {
                    alert('Вы опознаны');

                    dispatch({
                        type: GET_USER_ORDERS_SUCCESS,
                        phoneNumber: userInfo.PhoneNumber,
                        defaultCode: userInfo.Code,
                        userOrderList: data
                    });
                }
            })
            .fail(function(data) {
                alert('Покупателя с такими данными нет');
                console.log(data);

                dispatch({
                    type: GET_USER_ORDERS_FAIL,
                    error: 'Покупателя с такими данными нет'
                });
            });
    }
}


export function showManage() {
    return (dispatch) => {
        dispatch({
            type: SHOW_MANAGE
        })
    }
}

export function showLogin() {
    return (dispatch) => {
        dispatch({
            type: SHOW_LOGIN
        })
    }
}
