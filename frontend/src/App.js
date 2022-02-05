import SignUp from './components/SignUp';
import Books from './components/Books';
import Navbar from './components/Navbar';
function App() {
 
  return (
    <>
     {localStorage.getItem('userid')?<Books/>:<SignUp/>}
    </>
   
  );
}

export default App;
