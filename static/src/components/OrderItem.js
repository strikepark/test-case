import React, { PropTypes, Component } from 'react'
import { fmtCode } from '../helpers'

import OrderEdit from './OrderEdit'

export default class OrderItem extends Component {
    constructor(props) {
        super(props);

        this.state = {
            active: false,
            ws: new WebSocket('ws://planadotest.herokuapp.com/ws')
        };

        this.state.ws.onopen = () => {
            console.log('Соединение установлено')

            this.state.ws.send(JSON.stringify({
                Code: this.props.order.Code
            }))
        }

        this.state.ws.onclose = (event) => {
            if (event.wasClean) {
                console.log('Соединение закрыто чисто');
            } else {
                console.log('Обрыв соединения');
            }

            console.log('Код: ' + event.code + ' причина: ' + event.reason);
        }

        this.state.ws.onmessage = function(event) {
            console.log("Получены данные " + event.data);
        }

        this.state.ws.onerror = function(error) {
            console.log("Ошибка " + error.message);
        }
    }

    componentWillUnmount() {
        this.state.ws.close()
    }

    onOrderClick() {
        this.setState({
            active: !this.state.active
        });
    }

    render() {
        const order = this.props.order;
        const updateOrder= this.props.updateOrder;
        let code = fmtCode(order.code);
        let status = order.status;
        let color = status === 'Готовится' ? 'red' :
            status === 'Доставляется' ? 'orange' : 'green';

        return (
            <div className='list__item'>
                <div className='list__name' onClick={::this.onOrderClick}>
                    Заказ № {code}
                </div>

                <div className='list__status' onClick={::this.onOrderClick}>
                    Статус заказа: <span className={'list__status_' + color}>{status}</span>
                </div>

                <OrderEdit order={order} active={this.state.active} updateOrder={updateOrder} />
            </div>
        )
    }
}

OrderItem.propTypes = {
    order: PropTypes.object.isRequired
}
