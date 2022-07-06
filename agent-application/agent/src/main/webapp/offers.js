$(document).ready(function()
    {
        $get({
            url: '/listOffers'
        })
    }
);

function displayData(data) {
    preventDefault();
var offerTable = $("<table>")

        for (let offer of data)
        {
            let postTable = $('<tr></tr>');
            let company = $('<tr><td> "Employer: ' + offer.company + '</td></tr><tr></tr>');
            let position = $('<tr><td> "Position: ' + offer.position + '</td></tr><tr></tr>');
            let description = $('<tr rowspan = "5"><td> "Employer: ' + offer.position.description + '</td></tr><tr></tr>');
            postTable.append(company).append(position).append(description)

            var applyButton = $('<tr><td><input type="button" value="Apply"/>');
            var commentButton = $('<input type="button" value="Comment"/></td></tr>');

            postTable.append(applyButton).append(commentButton);

            for (let comment of offer.comments)
            {
                let author = $('<tr><td>'  + comment.author + '</td></tr>');
                let content = $('<tr><td>'  + comment.content + '</td></tr>');
                postTable.append(author);
                postTable.append(content);
            }
            offerTable.append(postTable);
            offerTable.append('<tr><tr><tr></tr></tr></tr>')
        }
	}