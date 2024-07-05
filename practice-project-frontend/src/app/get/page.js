'use client'

import {Header} from "../components/Header"

import {useState} from "react";
import {Card, Input, Button, Autocomplete, AutocompleteItem, Link, Pagination} from "@nextui-org/react"
import axios from "axios";

const experiences = [
  {label:"No experience", value:0},
  {label:"From 1 to 3 years", value:1},
  {label:"From 3 years to 6 years", value:2},
  {label:"Over 6 years", value:3}
]

export default function Parse(){

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

  const [loading, setLoading] = useState(false)
  const [minSalary, setMinSalary] = useState(0)
  const [experience, setExperience] = useState({label:"No experience", value:0})
  const [parsedVacancies, setParsedVacancies] = useState([])
  const [displayedVacancies, setDisplayedVacancies] = useState([])
  const [maxPages, setMaxPages] = useState(0)
  const [city, setCity] = useState("")
  const [name, setName] = useState("")
  return <>
  <Header/>
  <div className="container w-96 mx-auto mt-10 px-5">
  <Card className="p-3 mb-10" shadow="lg">
      <Input type="text" placeholder="Vacancy name..." className="mb-3 pt-2" 
      label="Enter vacancy name"
      labelPlacement="outside"
      onChange={(input) => setName(input.target.value)}
      ></Input>
      <Input type="text" placeholder="Enter city..." className="mb-3 pt-2" 
      label="City to parse"
      labelPlacement="outside"
      onChange={(input) => setCity(input.target.value)}
      ></Input>
      <Input
        label="Minimal Salary"
        placeholder="0"
        value={minSalary}
        labelPlacement="outside"
        className="mt-5"
        onChange={(event) => {
          if (parseInt(event.target.value) < 0){
              setMinSalary(0)
              return
          }
          setMinSalary(parseInt(event.target.value))
        }}
        startContent={
          <div className="pointer-events-none flex items-center">
            <span className="text-default-400 text-small">₽</span>
          </div>
        }
        type="number"
      />
      <Autocomplete
          defaultItems={experiences}
          label="Select your experience"
          placeholder="Select experience..."
          className="max-w-xs mt-6"
          onSelectionChange={(changed) => setExperience(changed)}
      >
          {(experience) => <AutocompleteItem key={experience.value}>{experience.label}</AutocompleteItem>}
      </Autocomplete>
      <Button color="primary" className="mt-6" isLoading={loading}
      onClick={() =>{
        let params = new URLSearchParams()
        params.append("experience", experience)
        params.append("minSalary", minSalary)
        params.append("city", city)
        params.append("currency", "RUR")
        params.append("name", name)
        setLoading(true)
        axios.get("http://localhost:8080/jobs", {params}).then((resp) => {
          setLoading(false)
          if (resp.data.items !== null){
            setParsedVacancies(resp.data.items)
            setDisplayedVacancies(resp.data.items.slice(0, 20))
            setMaxPages(Math.ceil(resp.data.pages))
          } else {
            setParsedVacancies([])
          }
        })
      }
      }
      >Get vacancies</Button>
      
  </Card></div>
  {parsedVacancies.length != 0 && displayedVacancies.map((vacancy) => 
    <div className="w-4/5 mx-auto" key={vacancy.id}>
        <Card className="p-4 mb-10">
            <div><Link isExternal href={vacancy.alternate_url} color="primary">{vacancy.name}</Link></div>
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