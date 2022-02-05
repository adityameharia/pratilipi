import {Box,Spacer,Text,Flex,Button} from '@chakra-ui/react'
import {useState } from 'react';
export default function Navbar({callback}){
    const [topcon, setTopcon] = useState(false);
    return(
        <Box bg="gray.100" p={2}>
            <Flex>
            <Text marginLeft="5vw" fontSize="3xl" fontWeight="extrabold">Books</Text>
            <Spacer/>
            <Button marginRight="5vw" colorScheme="blue" variant={topcon?'solid':'outline'} onClick={()=>{
                setTopcon(!topcon)
                callback(!topcon)
            }}>Top contents</Button>
            </Flex>
        </Box>
    )
}