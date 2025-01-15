import {Inputs} from '../components/Inputs';
import {Button} from '../components/Button';

export const FileUpload = ({handleFileUpload, setFile}) => (
    <div>
        <Inputs type={"file"} onchangeEvent={(e) => setFile(e.target.files[0])}/>
        <Button name={"Upload File"} onClick={handleFileUpload}/>
    </div>
)