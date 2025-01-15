import {Inputs} from '../components/Inputs';
import {Button} from '../components/Button';

export const MFA = ({mfaCode, handleMfaSubmit}) => (
    <div>
        <p>Your MFA code is: <strong>{mfaCode}</strong></p>
        <Inputs type={"text"} id={"mfaCodeInput"} placeHolder={"Enter code"}/>
        <Button name={"Submit"} onClick={handleMfaSubmit}/>
    </div>
)