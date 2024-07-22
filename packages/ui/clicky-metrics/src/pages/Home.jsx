import { IoCopy } from 'react-icons/io5';
import Navbar from '../components/Navbar';
import '../styles/home.css';
const Home = () => {
  return (
    <>
      <Navbar />
      <section className='home-container'>
        <h1 className='heading-primary'>URL Shortener</h1>
        <form className='form'>
          <input
            type='text'
            placeholder='Enter a long URL to shorten'
            className='user-input'
          ></input>
          <button type='submit' onClick={(e) => e.preventDefault()}>
            Shorten URL
          </button>
        </form>
        <p className='example'>Example: https://google.com</p>
        <h2 className='heading-secondary'>Shortened URL:</h2>
        <div className='shortened__info'>
          <a
            className='shortened-url'
            href='https://t.ly/WVus7'
            target='_blank'
          >
            https://t.ly/WVus7
          </a>
          <IoCopy className='icon-copy' />
        </div>
      </section>
    </>
  );
};

export default Home;
