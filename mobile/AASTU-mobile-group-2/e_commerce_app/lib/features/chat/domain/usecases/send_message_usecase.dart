import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/chat/data/data%20repository/data_repository.dart';
import 'package:e_commerce_app/features/chat/domain/entities/message_entity.dart';
import 'package:e_commerce_app/features/chat/domain/repository/chat_repository.dart';

class SendMessageUseCase {
  final ChatRepositoryImpl repository;

  SendMessageUseCase(this.repository);

  Future<Either<Failure,bool>> execute(String chatId, String message, String content) {
    return repository.sendMessage(chatId, message, content);
  }

}