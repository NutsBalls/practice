'use client';

import {useState} from "react";
import {Header} from "../components/Header"
import {Card, Input, Button, Pagination, Link} from "@nextui-org/react"
import axios from "axios"



export default function ParsePage(){

    const FormatSalary = function(salary){
        let from = salary.from
        let to = salary.to
        let currency = salary.currency
        let formattedSalary = ""


        if (from === 0 && to === 0){
            return "Зарплата не указана"
        }
        if (from !== 0){
            formattedSalary += from
        }
        if (to !== 0){
            if (formattedSalary !== ""){
                formattedSalary += " - " + to
            } else{
                formattedSalary += to
            }
        }
        formattedSalary += " " + currency

        return formattedSalary
    }

    const [vacancyName, setVacancyName] = useState("")
    const [city, setCity] = useState("")
    const [loading, setLoading] = useState(false)
    const [parsedVacancies, setParsedVacancies] = useState([])
    const [displayedVacancies, setDisplayedVacancies] = useState([])
    const [maxPages, setMaxPages] = useState(0)
    return <>
    <Header/>
    <div className="container w-96 mx-auto mt-10 px-5">
    <Card className="p-3 mb-10" shadow="lg">
        <Input type="text" placeholder="Vacancy name..." className="mb-3 pt-2" 
        label="Enter vacancy name"
        labelPlacement="outside"
        value={vacancyName}
        onChange={(e) => setVacancyName(e.target.value)}
        >
        </Input>
        <Input type="text" placeholder="Enter city..." className="mb-3 pt-2" 
        label="City to parse"
        labelPlacement="outside"
        value={city}
        onChange={(e) => setCity(e.target.value)}
        >
        </Input>
        <Button color="primary" className="mt-6" isLoading={loading}
        onClick={() => {
            let params = new URLSearchParams()
            params.append("text", vacancyName) 
            params.append("city", city) 
            setLoading(true)
            axios.get("http://localhost:8080/parse", {params}).then(function (resp){
                console.log(resp.data.items)
                setLoading(false)
                setParsedVacancies(resp.data.items)
                setDisplayedVacancies(resp.data.items.slice(0, 20))
                setMaxPages(Math.ceil(resp.data.items.length / 20))

            })
        }}
        >
            Start parsing
        </Button>
        
    </Card>
    </div>
    
    {parsedVacancies.length != 0 && displayedVacancies.map((vacancy) => 
    <div className="w-4/5 mx-auto" key={vacancy.id}>
        <Card className="p-4 mb-10">
            <div><Link isExternal href={vacancy.alternate_url} color="primary" >{vacancy.name}</Link></div>
            <div>{FormatSalary(vacancy.salary)}</div>
            <div>{vacancy.area.name}</div>
            <div> Работодатель: <Link href={vacancy.employer.url}>{vacancy.employer.name}</Link></div>
        </Card>
    </div>)}
    {parsedVacancies.length != 0 && <div className="w-4/5 mx-auto mt-5"><Pagination
    color="primary"
    size="lg"
    className="flex justify-center mr-auto"
    total={maxPages}
    initialPage={1}
    onChange={(page) => {
        let limit = 20
        let offset = (page - 1) * limit
        setDisplayedVacancies(parsedVacancies.slice(offset, offset + limit))
    }}
    /></div>}
    </>
}