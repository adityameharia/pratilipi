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
import { ChangeEvent, useCallback, useEffect, useState } from 'react';
import Navbar from './Navbar';
import {
    Pagination,
    usePagination,
    PaginationNext,
    PaginationPage,
    PaginationPrevious,
    PaginationContainer,
    PaginationPageGroup,
} from "@ajna/pagination";


export default function Books() {

    const [topcon, setTopcon] = useState(false)
    const {
        currentPage,
        setCurrentPage,
        pagesCount,
        pages
    } = usePagination({
        // pagesCount: 200,
        total: 60,
        initialState: { currentPage: 1, pageSize: 5 },
        limits: {
            outer: 2,
            inner: 2,
        },
    });

    useEffect(() => {
        console.log(currentPage)
    }, [currentPage])

    const callback = useCallback((topcon) => {
        console.log(currentPage)
        setTopcon(topcon)
        console.log(topcon)
    }, [])

    const [books, setBooks] = useState([{ title: "i dont know what i ", story:"Sit nulla est ex deserunt exercitation anim occaecat. Nostrud ullamco deserunt aute id consequat veniam incididunt duis in sint irure nisi. Mollit officia cillum Lorem ullamco minim nostrud elit officia tempor esse quis.Sunt ad dolore quis aute consequat. Magna exercitation reprehenderit magna aute tempor cupidatat consequat elit dolor adipisicing. Mollit dolor eiusmod sunt ex incididunt cillum quis. Velit duis sit officia eiusmod Lorem aliqua enim laboris do dolor eiusmod. Et mollit incididunt nisi consectetur esse laborum eiusmod pariatur proident Lorem eiusmod et. Culpa deserunt nostrud ad veniam.",likes: 10, date: "09/09/2001" }, 
    { title: "i dont know what i ", likes: 10, date: "09/09/2001" }, { title: "i dont know what i ", likes: 10, date: "09/09/2001" }, { title: "i dont know what i ", likes: 10, date: "09/09/2001" }, { title: "i dont know what i ", likes: 10, date: "09/09/2001" }, { title: "i dont know what i am doing", likes: 10, date: "09/09/2001" }, { title: "i dont know what i am doing", likes: 10, date: "09/09/2001" }, { title: "i dont know what i am doing", likes: 10, date: "09/09/2001" }, { title: "i dont know what i am doing", likes: 10, date: "09/09/2001" }])

    return (
        <>
            <Navbar callback={callback} />
            <Box marginTop="5vh"
             marginLeft="15vw" 
             marginRight="15vw">
                <Box>
                    {!topcon ?
                        <Center>
                            <Heading as="h1" marginLeft="3vw">
                                Some Books By Us
                            </Heading>
                        </Center> :
                        <Center>
                            <Heading as="h1" marginLeft="3vw">
                                Top Books by us
                            </Heading>
                        </Center>}
                    <br />
                    <SimpleGrid
                     minChildWidth='300px'
                      spacingX='50px' 
                      spacingY='5vh' 
                      justifyContent="center">
                        {books.map((b) => (
                            <Card key={b} 
                            title={b.title} 
                            likes={b.likes} 
                            date={b.date} 
                            story={b.story}
                            liked={false}
                             />
                        ))}
                    </SimpleGrid>
                </Box>
                <br />
                <br />
                <Center>
                    <Pagination
                        pagesCount={pagesCount}
                        currentPage={currentPage}
                        onPageChange={setCurrentPage}
                    >
                        <PaginationContainer>
                            <PaginationPrevious p={4} mx={2}>Previous</PaginationPrevious>
                            <PaginationPageGroup>
                                {pages.map((page) => (
                                    <PaginationPage
                                        p={4}
                                        mx={1}
                                        _hover={{
                                            bg: "blue.500"
                                        }}
                                        _current={{
                                            bg: "blue.100"
                                        }}
                                        key={`pagination_page_${page}`}
                                        page={page}
                                    />
                                ))}
                            </PaginationPageGroup>
                            <PaginationNext p={4} mx={2}>Next</PaginationNext>
                        </PaginationContainer>
                    </Pagination>
                </Center>
            </Box>

        </>
    )

}