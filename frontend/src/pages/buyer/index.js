// ** React Imports
import { useState, Fragment } from 'react'
import Grid from '@mui/material/Grid'
import BuyerForm from 'src/pages/buyerForm'
import TableCustomized from 'src/views/tables/TableCustomized'
import FormControl from '@mui/material/FormControl'
import OutlinedInput from '@mui/material/OutlinedInput'
import { styled, useTheme } from '@mui/material/styles'
import MuiCard from '@mui/material/Card'
import MuiFormControlLabel from '@mui/material/FormControlLabel'

// ** Layout Import
import BlankLayout from 'src/@core/layouts/BlankLayout'
import TableCustomizedSales from 'src/views/tables/TablesCustomizedSales'
import { useEffect } from 'react'

// ** Styled Components
const Card = styled(MuiCard)(({ theme }) => ({
  [theme.breakpoints.up('sm')]: { width: '28rem' }
}))

const LinkStyled = styled('a')(({ theme }) => ({
  fontSize: '0.875rem',
  textDecoration: 'none',
  color: theme.palette.primary.main
}))

const FormControlLabel = styled(MuiFormControlLabel)(({ theme }) => ({
  marginTop: theme.spacing(1.5),
  marginBottom: theme.spacing(4),
  '& .MuiFormControlLabel-label': {
    fontSize: '0.875rem',
    color: theme.palette.text.secondary
  }
}))

const BuyerPage = () => {
  // ** States
  const [orderData, setOrderData] = useState()
  const [orderDataSales, setOrderDataSales] = useState()

  // ** Hook
  const theme = useTheme()

  const getData = () => {
    //get table orderbook data
    // const response='data from web3 order book'
    // setOrderData(response)
  }

  const getDataSales = () => {
    //get table orderbook sales data
    // const response='data from web3 order book'
    // setOrderDataSales(response)
  }

  useEffect(() => {
    getData()
    getDataSales()
  }, [])

  return (
    <div>
      <Grid container spacing={6} padding={8}>
        <Grid item xs={12} md={4}>
          <BuyerForm />
        </Grid>
        <Grid item xs={12} md={8}>
          <TableCustomized orderData={orderData} />
          <TableCustomizedSales orderDataSales={orderDataSales} />
        </Grid>
      </Grid>
    </div>
  )
}
BuyerPage.getLayout = page => <BlankLayout>{page}</BlankLayout>

export default BuyerPage
