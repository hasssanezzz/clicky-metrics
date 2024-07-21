import '../styles/navbar.css';
const Navbar = () => {
  return (
    <nav className='navbar'>
      <div className='logo-container'>
        <p className='logo-text'> Clicky metrics</p>
      </div>
      <ul className='buttons-container'>
        <li>
          <a href='#' className='btn-sign'>
            Sign up
          </a>
        </li>
        <li>
          <a href='#' className='btn-login'>
            Login
          </a>
        </li>
      </ul>
    </nav>
  );
};
export default Navbar;
