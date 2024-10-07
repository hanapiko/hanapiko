# GROUPIE TRACKER
![](https://www.globallogic.com/wp-content/uploads/2021/11/1024-Develop-Restful-API-using-Go-and-Gin.jpg)
Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information.

## API Structure
The API, that consists of four parts:
* Artists, containing information about some bands and artists including:
   * Name(s)
   * Image
   * In which year they began their activity
   * Date of their first album
   * Members.
* Locations: consists in their last and/or upcoming concert locations.
* Dates, consists in their last and/or upcoming concert dates.
* Relation links all the other parts, artists, dates and locations.

## Features
User-friendly website to display artist information and concert details
Client-server communication for real-time data fetching
Error handling to ensure stability across all pages

To run the project locally follow these steps:
1. Clone the repository
```bash
git clone https://learn.zone01kisumu.ke/git/ferdyodhiambo/groupie-tracker.git
```

2. Run the server
- navigate to the directory
```bash
cd groupie-tracker
``` 
then use the following command to run
```bash
go run .
```

3. Access the website: open a web browser and navigate to  http://localhost:3000

## Technologies Used

    Go (Golang)
    HTML/CSS
    JavaScript for frontend

## Demo project
![Demo GIF](images/art.gif)


## License

This project is licensed under the terms of the [MIT License](./LICENSE).
Contact Information

For any questions or inquiries, feel free to reach out:
    Email: apikojuma94@gmail.com

## Contribution
* If you want to contribute to this project, feel free to open issues or pull requests. Any improvements, bug fixes, or features are welcome!
