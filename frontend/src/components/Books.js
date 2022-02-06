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
import axios from 'axios'


export default function Books() {

    const [topcon, setTopcon] = useState(false)
    const [books, setBooks] = useState([])
    const [totalPage, setTotalPage] = useState(0)
    const {
        currentPage,
        setCurrentPage,
        pagesCount,
        pages
    } = usePagination({
        total: totalPage,
        initialState: { currentPage: 1, pageSize: 9 },
        limits: {
            outer: 2,
            inner: 2,
        },
    });

    useEffect(() => {
        async function fetchData() {
            try {
                if (topcon == false) {
                    let resp = await axios.get("http://localhost:8001/books/" + localStorage.getItem('userid') + "/" + currentPage)
                    setBooks(resp.data.books.data)
                    setTotalPage(resp.data.books.count)
                } else {
                    let resp = await axios.get("http://localhost:8001/getmostliked/" + localStorage.getItem('userid'))
                    setBooks(resp.data.mostLiked)
                }
            } catch (err) {
                console.log("error screen")
                if (!err.response) {
                    alert("Network Error")
                } else {
                    if (err.response.status === 400) {
                        alert("Invalid userid")
                        localStorage.removeItem("userid");
                        window.location.reload();
                    } else {
                        alert("Internal Server Error")
                    }
                }
            }

        }
        fetchData()
    }, [currentPage, topcon])

    const callback = useCallback((topcon) => {
        console.log(currentPage)
        setTopcon(topcon)
        console.log(topcon)
    }, [])



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
                        {books?.map((b) => (
                            <Card key={b.id}
                                title={b.title}
                                likes={b.likes}
                                date={b.date}
                                story={b.story}
                                liked={b.liked}
                            />
                        ))}
                    </SimpleGrid>
                </Box>
                <br />
                <br />
                {!topcon && <Center>
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
                </Center>}
            </Box>

        </>
    )

}