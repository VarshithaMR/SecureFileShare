import './App.css';
import {Inputs} from "./components/Inputs";
import {Button} from "./components/Button";
import {useEffect, useState} from "react";
import axios from 'axios'

function App() {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [token, setToken] = useState('')
    const [isCodeGenerated, setIsCodeGenerated] = useState(false);
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [file, setFile] = useState(null)
    const [mfaCode, setMFACode] = useState('')
    const [isAdminRole, setAdminRole] = useState(false)
    const [filesUploaded, setFilesUploaded] = useState('')
    const [isShowFiles, setIsShowFiles] = useState(false)

    useEffect(() => {
        window.localStorage.clear();
    }, [])

    const handleLogin = async () => {
        try {
            const response = await axios.post("/login", {password, username})
            console.log(response)
            setMFACode(response.data.code)
            setIsCodeGenerated(true)
            localStorage.setItem("username", username)
            localStorage.setItem("password", password)
            localStorage.setItem("mfacode", response.data.code)
        } catch (error) {
            console.error("MFA Code generation failed", error)
        }
    }

    const handleMfaSubmit = async() => {
        try {
            const response = await axios.post("/verify", {
                 mfaCode, username, password})
            setToken(response.data.token)
            setIsLoggedIn(true)
            localStorage.setItem("authToken", response.data.token)
            if (response.data.role === "admin") {
                setAdminRole(true)
            }
            console.log("Login successful, JWT token : ", response.data.token)
        } catch (error) {
            console.error("MFA Validation failed", error)
        }
    }

    const handleFileUpload = async () => {
        const formData = new FormData();
        formData.append("file", file)

        try {
            await axios.post("/upload", formData, {
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'UploadedBy': username
                }
            })
            console.log("File uploaded successfully!")
        } catch (error) {
            console.error("File upload failed", error)
        }
    }

    const handleShowUploadedFiles = async () => {
        try {
            const response = await axios.get("/showfiles")
            setFilesUploaded(response.data)
            setIsShowFiles(true)
            console.log("Showing all files")
        } catch (error) {
            console.error("Show Files uploaded failed", error)
        }
    }

    const handleLogout = () => {
        localStorage.clear()
        setUsername('')
        setPassword('')
        setToken('')
        setIsCodeGenerated(false)
        setIsLoggedIn(false)
        setFile(null)
        setMFACode('')
        setAdminRole(false)
        setFilesUploaded('')
        setIsShowFiles(false)
    }

    return (
        <div>
            <h1>File Sharing App</h1>
            {!isLoggedIn && !isCodeGenerated && (
                <div>
                    <Inputs type={"text"} input={username} onchangeEvent={(e) => setUsername(e.target.value)}
                            placeHolder={"Username"}/>
                    <Inputs type={"password"} input={password} onchangeEvent={(e) => setPassword(e.target.value)}
                            placeHolder={"Password"}/>
                    <Button name={"Generate Code to login"} onClick={handleLogin}/>
                </div>
            )}

            {!isLoggedIn && isCodeGenerated && (
                <div>
                    <p>Your MFA code is: <strong>{mfaCode}</strong></p>
                    <Inputs type={"text"} id={"mfaCodeInput"} placeHolder={"Enter code"}/>
                    <Button name={"Submit"} onClick={handleMfaSubmit}/>
                </div>
            )}

            {isLoggedIn && isCodeGenerated  &&(
                <div>
                    <Inputs type={"file"} onchangeEvent={(e) => setFile(e.target.files[0])}/>
                    <Button name={"Upload File"} onClick={handleFileUpload}/>
                    {isAdminRole && (
                        <>
                            <Button name={"Show Files Uploaded"} onClick={handleShowUploadedFiles}/>
                            {isShowFiles && (<div>
                                <h3>Uploaded Files</h3>
                                {filesUploaded.length === 0 ? (
                                    <p>No files uploaded yet.</p>
                                ) : (
                                    <ul>
                                        {filesUploaded.map((filename, index) => (
                                            <li key={index}>
                                                <strong>{filename.filename}:{filename.username}</strong>
                                            </li>
                                        ))}
                                    </ul>
                                )}
                            </div>)}
                        </>
                    )}
                </div>
            )}

            {isLoggedIn && (
                <div>
                    <button onClick={handleLogout}>Logout</button>
                </div>
            )}

        </div>
    )
}

export default App;
