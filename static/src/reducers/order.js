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

const initialState = {
    orderList: [],
    fetching: true,
    error: ''
}

export default function order(state = initialState, action) {
    switch (action.type) {
        case GET_ORDERS_REQUEST:
            return {
                ...state,
                fetching: true,
                error: ''
            }
        case GET_ORDERS_SUCCESS:
            return {
                ...state,
                orderList: action.orderList,
                fetching: false,
                error: ''
            }
        case GET_ORDERS_FAIL:
            return {
                ...state,
                fetching: false,
                error: action.error
            }


        case UPDATE_ORDER_REQUEST:
            return {
                ...state,
                fetching: true,
                error: ''
            }
        case UPDATE_ORDER_SUCCESS:
            return {
                ...state,
                fetching: false,
                error: ''
            }
        case UPDATE_ORDER_FAIL:
            return {
                ...state,
                fetching: false,
                error: action.error
            }


        case NEW_ORDER_REQUEST:
            return {
                ...state,
                fetching: true,
                error: ''
            }
        case NEW_ORDER_SUCCESS:
            return {
                ...state,
                orderList: [action.order, ...state.orderList],
                fetching: false,
                error: ''
            }

        case NEW_ORDER_FAIL:
            return {
                ...state,
                fetching: false,
                error: action.error
            }


        default:
            return state;
    }
}
