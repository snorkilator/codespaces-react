import './App.css';

function App() {
  return (
    <div className="App">
      <form action="/login">
        <input type="email" value="enter email"></input>
        <input type="password" value="enter password"></input>
        <button >Log in</ button>
      </form>
    </div>
  );
}

export default App;
