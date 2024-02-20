import React from 'react';
import { Link } from 'react-router-dom';

export default function Header() {
  return (
    <header>
        <div>
            <span className='logo'><a href='/'>ПАО НПО "Алмаз"</a></span>
            <ul className='nav'>
                <li><Link to="/docs">Документы</Link></li>
                <li><Link to="/profile">Профиль</Link></li>
            </ul>
        </div>
        <div className='presentation'></div>
    </header>
  )
}
