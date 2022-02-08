import{Spinner,Center} from '@chakra-ui/react';

export default function Spin() {
    return (
        <Center><Spinner
            m={'5'}
            thickness='4px'
            speed='0.65s'
            emptyColor='gray.200'
            color='blue.500'
            size='xl'
        /></Center>
    )
}