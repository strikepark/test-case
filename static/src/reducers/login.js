import {
    GET_USER_ORDERS_REQUEST,
    GET_USER_ORDERS_SUCCESS,
    GET_USER_ORDERS_FAIL,

    SHOW_MANAGE,
    SHOW_LOGIN
} from '../constants/Login'

const initialState = {
    user: {},
    error: '',
    isLogin: true,
    isManage: false,
    isUser: false,
    loginFetching: false
}

export default function login(state = initialState, action) {
    switch (action.type) {
        case GET_USER_ORDERS_REQUEST:
            return {
                ...state,
                loginFetching: true,
                error: ''
            }
        case GET_USER_ORDERS_SUCCESS:
            return {
                ...state,

                user: {
                    phoneNumber: action.phoneNumber,
                    defaultCode: action.defaultCode,
                    userOrderList: action.userOrderList
                },
                isLogin: false,
                isUser: true,

                loginFetching: false,
                error: ''
            }
        case GET_USER_ORDERS_FAIL:
            return {
                ...state,
                loginFetching: false,
                error: action.error
            }

        case SHOW_MANAGE:
            return {
                ...state,
                isLogin: false,
                isManage: true
            }

        case SHOW_LOGIN:
            return {
                ...state,
                isLogin: true,
                isManage: false,
                isUser: false
            }

        default:
            return state
    }
}
