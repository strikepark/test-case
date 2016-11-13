import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import * as orderActions from '../actions/OrderActions'
import * as loginActions from '../actions/LoginActions'
import Manage from '../components/Manage'
import Login from '../components/Login'
import User from '../components/User'
import $ from 'jquery';

class App extends Component {
    render() {
        const { fetching, orderList } = this.props.order
        const { updateOrder, createOrder } = this.props.orderActions

        const { loginFetching, user, isManage, isLogin, isUser } = this.props.login
        const { showManage, getUserOrders, showLogin } = this.props.loginActions

        if (fetching || loginFetching) {
            $('body').addClass('fetching')
        } else {
            $('body').removeClass('fetching')
        }

        return (
            <div>
                <Manage active={isManage} showLogin={showLogin} createOrder={createOrder} orders={orderList} updateOrder={updateOrder} />
                <Login showManage={showManage} getUserOrders={getUserOrders} active={isLogin} />
                <User orders={user.userOrderList} active={isUser} />
            </div>
        );
    }
}

function mapStateToProps(state) {
    return {
        order: state.order,
        login: state.login
    }
}

function mapDispatchToProps(dispatch) {
    return {
        orderActions: bindActionCreators(orderActions, dispatch),
        loginActions: bindActionCreators(loginActions, dispatch)
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(App)
