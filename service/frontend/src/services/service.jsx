import axios from 'axios';

// Function to generate MFA code
export const generateMfaCode = async (username, password) => {
    try {
        const response = await axios.post("/login", {username, password})
        return response.data
    } catch (error) {
        throw new Error("MFA Code generation failed: " + error.message)
    }
}

// Function to verify MFA and login
export const verifyMfaCodeAndLogin = async (mfaCode, username, password) => {
    try {
        const response = await axios.post("/verify", {mfaCode, username, password})
        return response.data
    } catch (error) {
        throw new Error("MFA Validation failed: " + error.message)
    }
}

// Function to upload a file
export const uploadFile = async (file, token, username) => {
    const formData = new FormData()
    formData.append("file", file)

    try {
        await axios.post("/upload", formData, {
            headers: {
                'Authorization': `Bearer ${token}`,
                'UploadedBy': username
            }
        })
        return {success: true, message: "File uploaded successfully!"}
    } catch (error) {
        throw new Error("File upload failed: " + error.message)
    }
}

// Function to show uploaded files
export const showUploadedFiles = async () => {
    try {
        const response = await axios.get("/showfiles")
        return response.data
    } catch (error) {
        throw new Error("Show files upload failed: " + error.message)
    }
}
