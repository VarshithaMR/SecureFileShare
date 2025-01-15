import './App.css';
import {Button} from "./components/Button";
import {useEffect, useState} from "react";
import {Login} from "./pages/Login";
import {MFA} from "./pages/MFA";
import {FileUpload} from "./pages/Upload";
import {ShowFiles} from "./pages/Show";
import {Logout} from "./pages/Logout";
import {generateMfaCode, showUploadedFiles, uploadFile, verifyMfaCodeAndLogin} from "./services/service";

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
            const response = await generateMfaCode(username, password)
            console.log(response)
            setMFACode(response.code)
            setIsCodeGenerated(true)
            localStorage.setItem("username", username)
            localStorage.setItem("password", password)
            localStorage.setItem("mfacode", response.code)
        } catch (error) {
            console.error("MFA Code generation failed", error)
        }
    }

    const handleMfaSubmit = async () => {
        try {
            const response = await verifyMfaCodeAndLogin(mfaCode, username, password)
            setToken(response.token)
            setIsLoggedIn(true)
            localStorage.setItem("authToken", response.token)
            if (response.role === "admin") {
                setAdminRole(true)
            }
            console.log("Login successful, JWT token : ", response.token)
        } catch (error) {
            console.error("MFA Validation failed", error)
        }
    }

    const handleFileUpload = async () => {
        const formData = new FormData();
        formData.append("file", file)

        try {
            await uploadFile(file, token, username)
            console.log("File uploaded successfully!")
        } catch (error) {
            console.error("File upload failed", error)
        }
    }

    const handleShowUploadedFiles = async () => {
        try {
            const response = await showUploadedFiles()
            setFilesUploaded(response)
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
                <Login username={username} setUsername={setUsername} password={password} setPassword={setPassword}
                       handleLogin={handleLogin}/>
            )}

            {!isLoggedIn && isCodeGenerated && (
                <MFA mfaCode={mfaCode} handleMfaSubmit={handleMfaSubmit}/>
            )}

            {isLoggedIn && isCodeGenerated && (
                <div>
                    <FileUpload handleFileUpload={handleFileUpload} setFile={setFile}/>
                    {isAdminRole && (
                        <>
                            <Button name={"Show Files Uploaded"} onClick={handleShowUploadedFiles}/>
                            {isShowFiles && <ShowFiles filesUploaded={filesUploaded}/>}
                        </>
                    )}
                </div>
            )}

            {isLoggedIn && <Logout handleLogout={handleLogout}/>}

        </div>
    )
}

export default App;
