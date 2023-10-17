import { 
    FormControl,
    InputLabel,
    Select,
    MenuItem,
} from '@mui/material';


export  const UserSelectBox = () =>{
  return(
  <FormControl fullWidth>
    <InputLabel id="usertable">User</InputLabel>
    <Select
      labelId="usertable"
      id="usertable"
      value={user}
      label="User"
      onChange={handleChange}
    >
    <MenuItem value={10}>Ten</MenuItem>
    <MenuItem value={20}>Twenty</MenuItem>
    <MenuItem value={30}>Thirty</MenuItem>
    </Select>
  </FormControl>
  )
}