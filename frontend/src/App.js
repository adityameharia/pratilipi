import SignUp from './components/SignUp';
import Books from './components/Books';
import Navbar from './components/Navbar';
import { useEffect,useState } from 'react';
function App() {
  let[id,setId]=useState(localStorage.getItem('userid'))
  useEffect(()=>{
    setId(localStorage.getItem('userid'))
  },[id])
  return (
    <>
     {localStorage.getItem('userid')?<Books/>:<SignUp/>}
    </>
   
  );
}

export default App;
