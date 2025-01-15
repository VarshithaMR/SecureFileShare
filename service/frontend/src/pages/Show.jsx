export const ShowFiles = ({filesUploaded}) => (
    <div>
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
    </div>
)