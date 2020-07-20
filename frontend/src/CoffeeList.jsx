import React from "react";
import Table from "react-bootstrap/Table";
import axios from "axios";
import { useState } from "react";
import { useEffect } from "react";

const CoffeeList = () => {
  const [state, setState] = useState({ products: [] });

  useEffect(() => {
    readData();
  });

  const readData = () => {
    axios
      .get(window.global.api_location + "/products")
      .then((response) => {
        console.log(response.data);

        setState({ products: response.data });
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const getProducts = () => {
    let table = [];

    for (let i = 0; i < state.products.length; i++) {
      table.push(
        <tr key={i}>
          <td>{state.products[i].name}</td>
          <td>{state.products[i].price}</td>
          <td>{state.products[i].sku}</td>
        </tr>
      );
    }

    return table;
  };

  return (
    <div>
      <h1 style={{ marginBottom: "40px" }}>Menu</h1>
      <Table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Price</th>
            <th>SKU</th>
          </tr>
        </thead>
        <tbody>{getProducts()}</tbody>
      </Table>
    </div>
  );
};

export default CoffeeList;
