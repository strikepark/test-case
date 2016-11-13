import React, { PropTypes, Component } from 'react'
import { fmtCode } from '../helpers'

export default class UserOrdersItem extends Component {
    constructor(props) {
        super(props);

        this.state = {
            ws: new WebSocket('ws://planadotest.herokuapp.com/ws'),
            active: this.props.order.code === this.props.defaultCode,
            changeHistories: JSON.parse(this.props.order.ChangeHistories),
            curStatus: this.props.order.status
        };
    }

    setupWebsocket() {
        let websocket = this.state.ws

        websocket.onopen = () => {
            console.log('Websocket connected')

            websocket.send(JSON.stringify({
                Code: this.props.order.code
            }))
        }

        websocket.onmessage = (evt) => {
            let curHistory = this.state.changeHistories
            let msg = JSON.parse(evt.data)

            let newHistory = JSON.parse(msg.History)

            curHistory.push(newHistory)

            this.setState({
                changeHistories: curHistory,
                curStatus: newHistory.Status
            })
        }

        websocket.onclose = () => {
            setTimeout(() => {
                this.setupWebsocket();
            }, 3600)
        }
    }

    componentDidMount() {
        this.setupWebsocket()
    }

    componentWillUnmount() {
        let websocket = this.state.ws
        websocket.close()
    }

    onOrderClick() {
        this.setState({
            active: !this.state.active
        });
    }

    render() {
        const order = this.props.order
        const active = this.state.active ? '' : 'hidden'

        const code = fmtCode(order.code)
        const status = this.state.curStatus
        const color = status === 'Готовится' ? 'red' :
            status === 'Доставляется' ? 'orange' : 'green'

        const historyObj = this.state.changeHistories;
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
