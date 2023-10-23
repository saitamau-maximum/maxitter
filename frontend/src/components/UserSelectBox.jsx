import { 
    FormControl,
    InputLabel,
    Select,
    MenuItem,
} from '@mui/material';


export  const UserSelectBox = ({selectedUser, handleChange}) =>{
  return(
  <>
  <FormControl fullWidth>
    <InputLabel id="user">User</InputLabel>
    <Select
      labelId="user"
      id="user"
      value={selectedUser}
      label="User"
      onChange={handleChange}
    >
    <MenuItem value={10}>Ten</MenuItem>
    <MenuItem value={20}>Twenty</MenuItem>
    <MenuItem value={30}>Thirty</MenuItem>
    </Select>
  </FormControl>
  </>
  )
}