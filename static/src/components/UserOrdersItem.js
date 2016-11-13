import React, { PropTypes, Component } from 'react'
import { fmtCode } from '../helpers'

export default class UserOrdersItem extends Component {
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
            console.log('Получены данные ' + event.data);
        }

        this.state.ws.onerror = function(error) {
            console.log('Ошибка ' + error.message);
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
        const order = this.props.order
        const active = this.state.active ? '' : 'hidden'

        const code = fmtCode(order.code);
        const status = order.status;
        const color = status === 'Готовится' ? 'red' :
            status === 'Доставляется' ? 'orange' : 'green';

        const historyObj = JSON.parse(order.ChangeHistories);
        const history = [];
        historyObj.forEach((val, i) => {
            let color = val.Status === 'Готовится' ? 'red' :
                        val.Status === 'Доставляется' ? 'orange' : 'green';

            let arrow = (i === (historyObj.length - 1)) ? '↳ ' : '↓ ';

            let jsx = (
                <div className='history__item' key={i}>
                    <span className={'history__status history__status_' + color}>{arrow + val.Status} </span>
                    <span className='history__date'>{val.Date}</span>
                </div>
            )

            history.push(jsx)
        })

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
                        <h3 className='history__hl'>История заказа:</h3>
                        {history}
                    </div>
                </div>
            </div>
        )
    }
}
