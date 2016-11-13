import React, { Component } from 'react'

import UserOrdersItem from './UserOrdersList'

export default class User extends Component {
    showLogin() {
        this.props.showLogin()
    }

    render() {
        const { active, orders } = this.props
        const isActive = active ? '' : 'hidden'

        return (
            <div className={'content ' + isActive}>
                <button onClick={::this.showLogin} type='button' className='button-xsmall pure-button content__back'>← Вернуться</button>

                <UserOrdersItem defaultCode={this.props.defaultCode} orders={orders} />
            </div>
        )
    }
}
