export const Inputs = ({type, input, onchangeEvent, placeHolder}) => (
    <input
        type={type}
        value={input}
        onChange={onchangeEvent}
        placeholder={placeHolder}
    />
)