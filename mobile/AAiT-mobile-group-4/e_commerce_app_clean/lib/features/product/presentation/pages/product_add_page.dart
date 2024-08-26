import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:image_picker/image_picker.dart';

import '../../domain/entities/product_entity.dart';
import '../bloc/product_bloc.dart';
import '../widgets/components/styles/custom_button.dart';
import '../widgets/components/styles/snack_bar_style.dart';
import '../widgets/components/styles/text_field_styles.dart';
import '../widgets/components/styles/text_style.dart';

class AddProudctPage extends StatefulWidget {
  const AddProudctPage({
    super.key,
  });

  @override
  State<AddProudctPage> createState() => _AddProudctPageState();
}

class _AddProudctPageState extends State<AddProudctPage> {
  File? _selectedImage;
  final TextEditingController _nameController = TextEditingController();
  final TextEditingController _categoryController = TextEditingController();
  final TextEditingController _priceController = TextEditingController();
  final TextEditingController _descriptionController = TextEditingController();

  Future _pickImageFromGallery() async {
    final returnedImage =
        await ImagePicker().pickImage(source: ImageSource.gallery);
    setState(() {
      if (returnedImage != null) {
        _selectedImage = File(returnedImage.path);
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: BlocConsumer<ProductBloc, ProductState>(
        listener: (context, state) {
          if (state is ProductCreatedState) {
            ScaffoldMessenger.of(context)
                .showSnackBar(customSnackBar('Product Added Successfully', Theme.of(context).primaryColor));
            context.read<ProductBloc>().add(LoadAllProductEvent());
            Navigator.pushNamed(context, '/home_page');
          } else if (state is ProductErrorState) {
            ScaffoldMessenger.of(context)
                .showSnackBar(SnackBar(content: Text(state.message)));
          }
        },
        builder: (BuildContext context, ProductState state) {
          if (state is ProductLoading) {
            return const Center(
              child: CircularProgressIndicator(),
            );
          } else {
            return Container(
              margin: const EdgeInsets.all(32),
              child: SingleChildScrollView(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Row(
                      crossAxisAlignment: CrossAxisAlignment.center,
                      children: [
                        IconButton(
                          onPressed: () {
                            Navigator.pop(context);
                          },
                          icon: Icon(
                            Icons.arrow_back_ios_rounded,
                            color: Theme.of(context).primaryColor
                          ),
                        ),
                        const SizedBox(width: 80),
                        const CustomTextStyle(
                            name: 'Add Product',
                            weight: FontWeight.w500,
                            size: 16),
                      ],
                    ),
                    const SizedBox(height: 23),
                    GestureDetector(
                      onTap: () {
                        _pickImageFromGallery();
                      },
                      child: Container(
                        width: 366,
                        height: 190,
                        decoration: BoxDecoration(
                          borderRadius: BorderRadius.circular(16),
                          color: const Color.fromRGBO(243, 243, 243, 1),
                        ),
                        child: _selectedImage == null
                            ? const Center(
                                child: Column(
                                  mainAxisAlignment: MainAxisAlignment.center,
                                  children: [
                                    Icon(
                                      Icons.image_outlined,
                                      size: 48,
                                    ),
                                    SizedBox(
                                      height: 16,
                                    ),
                                    CustomTextStyle(
                                        name: 'upload image',
                                        weight: FontWeight.w500,
                                        size: 14)
                                  ],
                                ),
                              )
                            : Image.file(_selectedImage!),
                      ),
                    ),
                    const SizedBox(height: 16),
                    const CustomTextStyle(
                      name: 'name',
                      weight: FontWeight.w500,
                      size: 14,
                    ),
                    const SizedBox(height: 8),
                    CustomTextField(controller: _nameController),
                    const SizedBox(height: 16),
                    const CustomTextStyle(
                      name: 'category',
                      weight: FontWeight.w500,
                      size: 14,
                    ),
                    const SizedBox(height: 8),
                    CustomTextField(
                      controller: _categoryController,
                    ),
                    const SizedBox(height: 16),
                    const CustomTextStyle(
                      name: 'price',
                      weight: FontWeight.w500,
                      size: 14,
                    ),
                    const SizedBox(height: 8),
                    Stack(
                      children: [
                        CustomTextField(
                          controller: _priceController,
                          keyboardType: TextInputType.number,
                        ),
                        const Positioned(
                          left: 290,
                          top: 16,
                          child: Icon(Icons.attach_money),
                        )
                      ],
                    ),
                    const SizedBox(height: 16),
                    const CustomTextStyle(
                      name: 'description',
                      weight: FontWeight.w500,
                      size: 14,
                    ),
                    const SizedBox(height: 8),
                    Container(
                      decoration: BoxDecoration(
                        color: const Color.fromARGB(195, 238, 238, 238),
                        borderRadius: BorderRadius.circular(6),
                      ),
                      height: 140,
                      child: CustomTextField(
                          controller: _descriptionController,
                          textInputAction: TextInputAction.newline,
                          keyboardType: TextInputType.multiline,
                          ),
                    ),
                    const SizedBox(height: 32),
                    CustomButton(
                      pressed: () {
                        // Ensure that all required fields are filled
                        if (_nameController.text.isEmpty ||
                            _categoryController.text.isEmpty ||
                            _priceController.text.isEmpty ||
                            _descriptionController.text.isEmpty ||
                            _selectedImage == null) {
                          ScaffoldMessenger.of(context).showSnackBar(
                            customSnackBar('All fields are required', Colors.redAccent),
                          );
                          return;
                        }
                        ProductEntity newProduct = ProductEntity(
                          id: '',
                          name: _nameController.text,
                          imageUrl: _selectedImage!.path,
                          description: _descriptionController.text,
                          price: double.parse(_priceController.text),
                        );
                        context
                            .read<ProductBloc>()
                            .add(CreateProductEvent(product:  newProduct));
                      },
                      name: 'ADD',
                      width: double.infinity,
                      height: 50,
                      textBgColor: Colors.white,
                      fgcolor: Theme.of(context).primaryColor,
                      bgcolor: Theme.of(context).primaryColor
                    ),
                    const SizedBox(height: 16),
                    CustomButton(
                      pressed: (){
                        _nameController.text = '';
                        _categoryController.text = '';
                        _priceController.text = '';
                        _descriptionController.text = '';
                        setState(() {
                          _selectedImage = null;
                        });
                      },
                      name: 'CLEAR',
                      width: double.infinity,
                      height: 50,
                      textBgColor: Theme.of(context).secondaryHeaderColor,
                      fgcolor: Theme.of(context).secondaryHeaderColor,
                      bgcolor: Colors.white ,
                    ),
                  ],
                ),
              ),
            );
          }
        },
      ),
    );
  }
}
