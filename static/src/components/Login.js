import React, { Component } from 'react'

import serialize from 'form-serialize'
import $ from 'jquery'

export default class Login extends Component {
    constructor(props) {
        super(props)
    }

    showManage() {
        this.props.showManage()
    }

    authHandle(e) {
        let form = $(e.target).closest('.form')

        if (form[0].checkValidity()) {
            let user = serialize(form[0], {hash: true})

            this.props.getUserOrders(user)
        } else {
            alert('Ошибки в полях формы')
        }
    }

    render() {
        const { active } = this.props
        const isActive = active ? '' : 'hidden'

        return (
            <div className={'login ' + isActive}>
                <h1>Аутентификация</h1>
                <form className='form pure-form'>
                    <fieldset className='pure-group'>
                        <input
                            name='Code'
                            type='number'
                            className='pure-input-1'
                            placeholder='Идентификатор заказа'
                            max='99999999'
                            min='0'
                            required
                        />

                        <input
                            name='PhoneNumber'
                            type='number'
                            className='pure-input-1'
                            placeholder='Телефон получателя'
                            min='70000000000'
                            max='89999999999'
                            required
                        />

                        <button
                            onClick={::this.authHandle}
                            type='button'
                            className='pure-button pure-input-1 pure-button-primary'
                        >
                            Посмотреть заказ
                        </button>
                    </fieldset>
                </form>

                <button onClick={::this.showManage} type='button' className='pure-button'>Операторский интерфейса →</button>
            </div>
        );
    }
}
