import './App.css';
import {Inputs} from "./components/Inputs";
import {Button} from "./components/Button";
import {useState} from "react";
import axios from 'axios'

function App() {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [token, setToken] = useState('')
    const [file, setFile] = useState(null)

    const handleLogin = async () => {
        try {
            const response = await axios.post('http://localhost:8080/login', { password, username })
            setToken(response.data.token);
        } catch (error) {
            console.error('Login failed', error);
        }
    }

    const handleFileUpload = async () => {
        const formData = new FormData();
        formData.append('file', file);

        try {
            await axios.post('http://localhost:8080/upload', formData, {
                headers: {
                    'Authorization': `Bearer ${token}`,
                }
            });
            alert('File uploaded successfully!');
        } catch (error) {
            console.error('File upload failed', error)
        }
    }

  return (
      <div>
        <h1>File Sharing App</h1>
        <div>
          <Inputs type={"text"} input={username}  onchangeEvent={(e) => setUsername(e.target.value)} placeHolder={"Username"}/>
          <Inputs type={"password"} input={password}  onchangeEvent={(e) => setPassword(e.target.value)} placeHolder={"Password"}/>
          <Button name={"Login"} onClick={handleLogin}/>
        </div>

        <div>
          <Inputs input={file} type={"file"} onchangeEvent={(e) => setFile(e.target.files[0])}/>
          <Button name={"Upload File"} onClick={handleFileUpload}/>
        </div>
      </div>
  );
}

export default App;
