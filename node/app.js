var messages = require('./article/article_pb');
var services = require('./article/article_grpc_pb');

var grpc = require('@grpc/grpc-js');

function main() {
  require('dotenv').config()

  var client = new services.ArticleServiceClient(process.env.SERVER, grpc.credentials.createInsecure());
  var params = new messages.Params();

  params.setQuery('Test...');
  params.setPage(2);
  params.setPerPage(10);

  client.readArticles(params, function(err, response) {
    if (err) {
      console.error(err);
    }
    console.log('Articles:', response.getArticlesList());
  });
}

main();
