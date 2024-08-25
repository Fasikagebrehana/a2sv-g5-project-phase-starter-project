import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/message.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

class SocketService {
  late IO.Socket socket;

  void connectToServer() {
    socket = IO.io('https://g5-flutter-learning-path-be.onrender.com/', <String, dynamic>{
      'transports': ['websocket'],
      'autoConnect': false,
    });

    socket.connect();

    // Event listeners
    socket.onConnect((_) {
      var text = 'Connected to server';
      print(text);  
    });

    socket.onDisconnect((_) {
      print('Disconnected from server');
    });

  }

  void sendMessage(Message message) {
    socket.emit('message', message);
  }

  void dispose() {
    socket.dispose();
  }
}
