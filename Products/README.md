## Product Service
<br> </br>

ALERT : Since the frontend is not there I have integrated grpc with the REST API To Take The Input which is not the correct way of doing so

<br> </br>

<b>Product Service</b>

| Feature  | Description  |
|----------|:-------------|
| Add a Product | To Add a Product on the System For Sale |
| List Products | To List Products |
| Edit a Product | Allowing Seller To Edit The Deatils Of The Already Selling Porduct |
| Delete a Product | Allowing Seller To Delete Product |
| SellerView | Ability To View The Quantity Sold and Quantity Left Of The Product For The Seller |


<br> </br>

<b>End Points</b>

| Request  | Description  | Url |
|----------|:-------------|:-------------|
| Post | To Sell New Item | host:8086/items/sellitem |
| Get | To Get The Details Of The Particular Item |host:8086/items/:id |
| Delete | To Delete The Item |host:8086/items/:id|
| Post | To Buy The Particular Item |host:8086/items/buy/:itemsID/:addressID |
| Patch | To Update The Item Details |host:8086/items/:id |
| Post | Seller View For The Parituclar Item They Are Selling |host:8086/seller/items/:id |
| Get | To Search All The Item With Given Name |host:8086/item/search/:itemName |

<br></br>
