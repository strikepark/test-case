import {
    GET_ORDERS_REQUEST,
    GET_ORDERS_SUCCESS,
    GET_ORDERS_FAIL,

    UPDATE_ORDER_REQUEST,
    UPDATE_ORDER_SUCCESS,
    UPDATE_ORDER_FAIL,

    NEW_ORDER_REQUEST,
    NEW_ORDER_SUCCESS,
    NEW_ORDER_FAIL
} from '../constants/Order'

import $ from 'jquery'

export function getOrders() {
    return (dispatch) => {
        dispatch({
            type: GET_ORDERS_REQUEST
        })

        $.get('http://planadotest.herokuapp.com/api/orders')
            .done(function(data) {
                dispatch({
                    type: GET_ORDERS_SUCCESS,
                    orderList: data
                });
            })
            .fail(function(data) {
                dispatch({
                    type: GET_ORDERS_FAIL,
                    error: data
                });
            });
    }
}

export function updateOrder(order, id) {
    return (dispatch) => {
        dispatch({
            type: UPDATE_ORDER_REQUEST
        })

        let url = 'http://planadotest.herokuapp.com/api/orders/' + id;

        $.ajax({
            beforeSend: function(xhr) {
                xhr.setRequestHeader('Access-Control-Allow-Origin', '*');
                xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
            },
            crossDomain: true,
            async: true,
            type: 'PUT',
            url: url,
            data: JSON.stringify(order)
        }).done(function(data) {
            alert('Заказ обновлен!');

            dispatch({
                type: UPDATE_ORDER_SUCCESS,
                order: data
            });
        })
        .fail(function(data) {
            alert('При обновлении заказа произошла ошибка: ' + data.statusText);
            console.log(data);

            dispatch({
                type: UPDATE_ORDER_FAIL,
                error: data
            });
        });
    }
}

export function createOrder(order) {
    return (dispatch) => {
        dispatch({
            type: NEW_ORDER_REQUEST
        })

        let url = 'http://planadotest.herokuapp.com/api/orders';

        $.ajax({
            beforeSend: function(xhr) {
                xhr.setRequestHeader('Access-Control-Allow-Origin', '*');
                xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
            },
            crossDomain: true,
            async: true,
            type: 'PUT',
            url: url,
            data: JSON.stringify(order)
        }).done(function(data) {
            alert('Заказ добавлен!');

            dispatch({
                type: NEW_ORDER_SUCCESS,
                order: data
            });
        })
        .fail(function(data) {
            alert('При обновлении заказа произошла ошибка: ' + data.statusText);
            console.log(data);

            dispatch({
                type: NEW_ORDER_FAIL,
                error: data
            });
        });
    }
}
