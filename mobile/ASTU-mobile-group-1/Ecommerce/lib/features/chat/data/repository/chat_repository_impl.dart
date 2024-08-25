import 'package:dartz/dartz.dart';
import '../../../../core/error/failure.dart';
import '../../domain/entities/chat_entity.dart';
import '../../domain/entities/message_entity.dart';
import '../../domain/repositories/chat_repository.dart';
import '../data_source/remote_data_source/remote_data_source.dart';


class ChatRepositoryImpl implements ChatRepository {
  final RemoteDataSource _remoteDataSource;

  ChatRepositoryImpl(this._remoteDataSource);

  @override
  Future<Either<Failure, ChatEntity>> initiateChat(String userId) async {
    try {
      final response = await _remoteDataSource.initiateChat(userId);
      if (response != null) {
        return Right((response));
      } else {
        return const Left( UnkownFailure());
      }
    } catch (error) {
        return const Left( UnkownFailure());

    }
  }

  @override
  Future<Either<Failure, List<MessageEntity>>> getChatMessages(String chatId) async {
    try {
      final response = await _remoteDataSource.getChatMessages(chatId); 
      if (response != null) {
        return Right(response);
      } else {
        return const Left(ServerFailure('Failed to fetch chat messages'));
      }
     
    } catch (error) {
      return const Left(ServerFailure('Could not get chat messages'));
    }
  }

  @override
  Future<Either<Failure, ChatEntity>> myChatbyId(String chatId) async {
    try {
      final response = await _remoteDataSource.getChatById(chatId);
      if (response != null) {
        return Right(response);
      } else {
        return const Left(ServerFailure('faied to load chat'));
      }
    } catch (error) {
      return const Left(UnkownFailure());
    }
  }

  @override
  Future<Either<Failure, List<ChatEntity>>> myChat() async {
    try {
      final response = await _remoteDataSource.getAllChats();
      if (response != null) {
        
        return Right(response);
      } else {
        return const Left(ServerFailure('Could not get all'));
      }
    } catch (error) {
        return const Left(ServerFailure('Failed to fetch chat messages'));

    }
  }

  @override
  Future<Either<Failure, bool>> deleteChat(String chatId) async {
    try {
      final response = await _remoteDataSource.deleteChat(chatId);
      return  Right(response);
     
      }
    catch(error) {
      return Left(ServerFailure(error.toString()));
    }
}
  }

