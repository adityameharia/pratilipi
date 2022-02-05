import {
    Box,
    Heading,
    Link,
    Text,
    Flex,
    SimpleGrid,
    Center,
    Button,
    Spacer
} from '@chakra-ui/react';
import Card from './Card'
import { useCallback, useState } from 'react';
import Pagination from "react-js-pagination";
import Navbar from './Navbar';


export default function Books() {

    const [topcon, setTopcon] = useState(false)
    const [activePage, setActivePage] = useState(1)

    const callback = useCallback((topcon) => {
        setTopcon(topcon)
        console.log(topcon)
    }, [])

    const changeActivePage = (page) => {
        setActivePage(page)
    }

    const [books, setBooks] = useState([{ title: "i dont know what i ", likes: 10, date: "09/09/2001" }, { title: "i dont know what i am doing", likes: 10, date: "09/09/2001" }, { title: "i dont know what i am doing", likes: 10, date: "09/09/2001" }, { title: "i dont know what i am doing", likes: 10, date: "09/09/2001" }, { title: "i dont know what i am doing", likes: 10, date: "09/09/2001" }])

    return (
        <>
            <Navbar callback={callback} />
            <Box marginTop="5vh" marginLeft="15vw" marginRight="15vw">
                {!topcon ?
                    <Box>
                        <Center><Heading as="h1" marginLeft="3vw">Some Books By Us</Heading></Center>

                        <SimpleGrid minChildWidth='300px' spacingX='50px' spacingY='5vh' justifyContent="center">
                            {books.map((b) => (
                                <Card k={b} title={b.title} likes={b.likes} date={b.date} />
                            ))}
                        </SimpleGrid>
                    </Box> :
                    <Box>
                        <Center><Heading as="h1" marginLeft="3vw">Top Books by us</Heading></Center>
                        <br />
                        <SimpleGrid minChildWidth='300px' spacingX='50px' spacingY='5vh' justifyContent="center">
                            {books.map((b) => (
                                <Card k={b} title={b.title} likes={b.likes} date={b.date} />
                            ))}
                        </SimpleGrid>
                    </Box>}
                <br />
                <Pagination
                    activePage={activePage}
                    itemsCountPerPage={10}
                    totalItemsCount={450}
                    pageRangeDisplayed={5}
                    onChange={changeActivePage}
                />
            </Box>

        </>
    )

}