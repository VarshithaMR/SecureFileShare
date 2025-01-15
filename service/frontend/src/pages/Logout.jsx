import {Button} from "../components/Button";

export const Logout = ({handleLogout}) => (
    <div>
        <Button name={"Logout"} onClick={handleLogout}/>
    </div>
)