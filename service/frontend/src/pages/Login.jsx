import {Inputs} from '../components/Inputs';
import {Button} from '../components/Button';

export const Login = ({username, setUsername, password, setPassword, handleLogin}) => (
    <div>
        <Inputs type={"text"} input={username} onchangeEvent={(e) => setUsername(e.target.value)}
                placeHolder={"Username"}/>
        <Inputs type={"password"} input={password} onchangeEvent={(e) => setPassword(e.target.value)}
                placeHolder={"Password"}/>
        <Button name={"Generate Code to login"} onClick={handleLogin}/>
    </div>
)