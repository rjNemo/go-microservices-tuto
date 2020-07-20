import React, { useState, useEffect } from "react";
import Table from "react-bootstrap/Table";
import axios from "axios";

const Products = ({ products }) =>
  products.map((product) => (
    <tr key={product.id}>
      <td>{product.name}</td>
      <td>{product.price}</td>
      <td>{product.sku}</td>
    </tr>
  ));

const CoffeeList = () => {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    const readData = () => {
      axios
        .get(window.global.api_location + "/products")
        .then((response) => {
          setProducts(response.data);
        })
        .catch((error) => {
          console.error(error);
        });
    };
    readData();
  }, []);

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
        <tbody>
          <Products products={products} />
        </tbody>
      </Table>
    </div>
  );
};

export default CoffeeList;
