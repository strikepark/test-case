import { combineReducers } from 'redux'
import order from './order'
import login from './login'

export default combineReducers({
    order,
    login
})
