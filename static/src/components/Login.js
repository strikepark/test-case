import React, { Component } from 'react'

import { Link } from 'react-router';

export default class Login extends Component {
    render() {
        return (
            <div className='login'>
                <h1>Аутентификация</h1>
                <form  className='pure-form'>
                    <fieldset className='pure-group'>
                        <input
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
                            type='button'
                            className='pure-button pure-input-1 pure-button-primary'
                        >
                            Посмотреть заказ
                        </button>
                    </fieldset>
                </form>

                <Link to='/manage' className='login__manage'>Операторский интерфейса →</Link>
            </div>
        );
    }
}
