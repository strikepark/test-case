import React, { PropTypes, Component } from 'react'
import { fmtCode } from '../helpers'

export default class UserOrdersItem extends Component {
    constructor(props) {
        super(props);

        this.state = {
            active: false
        };
    }

    onOrderClick() {
        this.setState({
            active: !this.state.active
        });
    }

    render() {
        const order = this.props.order
        const active = this.state.active ? '' : 'hidden'

        const code = fmtCode(order.code);
        const status = order.status;
        const color = status === 'Готовится' ? 'red' :
            status === 'Доставляется' ? 'orange' : 'green';

        const history = order.ChangeHistories;

        return (
            <div className='list__item'>
                <div className='list__name' onClick={::this.onOrderClick}>
                    Заказ № {code}
                </div>

                <div className='list__status' onClick={::this.onOrderClick}>
                    Статус заказа: <span className={'list__status_' + color}>{status}</span>
                </div>

                <div className={active}>
                    <p><b>Адрес получателя:</b> {order.recipientAddress}</p>
                    <p><b>Адрес отправителя:</b> {order.sendAddress}</p>

                    <div className='history'>
                        <b>История заказа:</b> {history}
                    </div>
                </div>
            </div>
        )
    }
}