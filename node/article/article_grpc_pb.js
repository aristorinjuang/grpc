// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var article_article_pb = require('../article/article_pb.js');

function serialize_Articles(arg) {
  if (!(arg instanceof article_article_pb.Articles)) {
    throw new Error('Expected argument of type Articles');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_Articles(buffer_arg) {
  return article_article_pb.Articles.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_Params(arg) {
  if (!(arg instanceof article_article_pb.Params)) {
    throw new Error('Expected argument of type Params');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_Params(buffer_arg) {
  return article_article_pb.Params.deserializeBinary(new Uint8Array(buffer_arg));
}


var ArticleServiceService = exports.ArticleServiceService = {
  readArticles: {
    path: '/ArticleService/ReadArticles',
    requestStream: false,
    responseStream: false,
    requestType: article_article_pb.Params,
    responseType: article_article_pb.Articles,
    requestSerialize: serialize_Params,
    requestDeserialize: deserialize_Params,
    responseSerialize: serialize_Articles,
    responseDeserialize: deserialize_Articles,
  },
};

exports.ArticleServiceClient = grpc.makeGenericClientConstructor(ArticleServiceService);
