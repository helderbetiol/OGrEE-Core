import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:ogree_app/common/api_backend.dart';
import 'package:ogree_app/common/definitions.dart';
import 'package:ogree_app/common/snackbar.dart';
import 'package:ogree_app/common/theme.dart';
import 'package:ogree_app/pages/projects_page.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:ogree_app/widgets/language_toggle.dart';

import 'reset_page.dart';

class LoginPage extends StatefulWidget {
  static String tag = 'login-page';

  const LoginPage({super.key});
  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final _formKey = GlobalKey<FormState>();
  bool _isChecked = false;
  static const inputStyle = OutlineInputBorder(
    borderSide: BorderSide(
      color: Colors.grey,
      width: 1,
    ),
  );

  String? _email;
  String? _password;
  String _apiUrl = "";
  BackendType? apiType;
  bool forgot = false;
  bool _tappedInside = false;

  @override
  Widget build(BuildContext context) {
    final localeMsg = AppLocalizations.of(context)!;
    bool isSmallDisplay = IsSmallDisplay(MediaQuery.of(context).size.width);
    return Scaffold(
      body: Container(
        // height: MediaQuery.of(context).size.height,
        decoration: const BoxDecoration(
          image: DecorationImage(
            image: AssetImage("assets/server_background.png"),
            fit: BoxFit.cover,
          ),
        ),
        child: CustomScrollView(slivers: [
          SliverFillRemaining(
            hasScrollBody: false,
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                Align(
                  alignment: Alignment.topCenter,
                  child: LanguageToggle(),
                ),
                const SizedBox(height: 5),
                Card(
                  child: Form(
                    key: _formKey,
                    child: Container(
                      constraints:
                          const BoxConstraints(maxWidth: 550, maxHeight: 515),
                      padding: EdgeInsets.only(
                          right: isSmallDisplay ? 45 : 100,
                          left: isSmallDisplay ? 45 : 100,
                          top: 50,
                          bottom: 30),
                      child: SingleChildScrollView(
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.stretch,
                          children: [
                            forgot
                                ? Row(
                                    children: [
                                      IconButton(
                                          padding: const EdgeInsets.all(0),
                                          constraints: const BoxConstraints(),
                                          onPressed: () => Navigator.of(context)
                                              .push(MaterialPageRoute(
                                                  builder: (context) =>
                                                      LoginPage())),
                                          icon: Icon(
                                            Icons.arrow_back,
                                            color: Colors.blue.shade900,
                                          )),
                                      SizedBox(width: isSmallDisplay ? 0 : 5),
                                      Text(
                                        "Request password reset",
                                        style: Theme.of(context)
                                            .textTheme
                                            .headlineLarge,
                                      ),
                                    ],
                                  )
                                : Center(
                                    child: Text(localeMsg.welcome,
                                        style: Theme.of(context)
                                            .textTheme
                                            .headlineLarge)),
                            const SizedBox(height: 8),
                            forgot
                                ? SizedBox(height: 10)
                                : Center(
                                    child: Text(
                                      localeMsg.welcomeConnect,
                                      style: Theme.of(context)
                                          .textTheme
                                          .headlineSmall,
                                    ),
                                  ),
                            forgot ? Container() : const SizedBox(height: 20),
                            dotenv.env['ALLOW_SET_BACK'] == "true"
                                ? backendInput()
                                : Center(
                                    child: Image.asset(
                                      "assets/custom/logo.png",
                                      height: 40,
                                    ),
                                  ),
                            dotenv.env['ALLOW_SET_BACK'] == "true"
                                ? Align(
                                    child: Padding(
                                      padding: const EdgeInsets.symmetric(
                                          vertical: 10),
                                      child: Badge(
                                        backgroundColor: Colors.white,
                                        label: Text(
                                          getBackendTypeText(),
                                          style: TextStyle(color: Colors.black),
                                        ),
                                      ),
                                    ),
                                  )
                                : const SizedBox(height: 30),
                            TextFormField(
                              onSaved: (newValue) => _email = newValue,
                              validator: (text) {
                                if (text == null || text.isEmpty) {
                                  return localeMsg.mandatoryField;
                                }
                                return null;
                              },
                              decoration: InputDecoration(
                                contentPadding: isSmallDisplay
                                    ? EdgeInsets.symmetric(
                                        horizontal: 12, vertical: 16)
                                    : null,
                                labelText: 'E-mail',
                                hintText: 'abc@example.com',
                                labelStyle: GoogleFonts.inter(
                                  fontSize: 11,
                                  color: Colors.black,
                                ),
                                border: inputStyle,
                              ),
                            ),
                            SizedBox(height: isSmallDisplay ? 10 : 20),
                            forgot
                                ? Container()
                                : TextFormField(
                                    obscureText: true,
                                    onSaved: (newValue) => _password = newValue,
                                    onEditingComplete: () => tryLogin(),
                                    validator: (text) {
                                      if (!forgot &&
                                          (text == null || text.isEmpty)) {
                                        return localeMsg.mandatoryField;
                                      }
                                      return null;
                                    },
                                    decoration: InputDecoration(
                                      contentPadding: isSmallDisplay
                                          ? EdgeInsets.symmetric(
                                              horizontal: 12, vertical: 16)
                                          : null,
                                      labelText: localeMsg.password,
                                      hintText: '********',
                                      labelStyle: GoogleFonts.inter(
                                        fontSize: 11,
                                        color: Colors.black,
                                      ),
                                      border: inputStyle,
                                    ),
                                  ),
                            !forgot
                                ? SizedBox(height: isSmallDisplay ? 15 : 25)
                                : Container(),
                            forgot
                                ? TextButton(
                                    onPressed: () => Navigator.of(context).push(
                                      MaterialPageRoute(
                                        builder: (context) => ResetPage(
                                          token: '',
                                        ),
                                      ),
                                    ),
                                    child: Text(
                                      "I have a reset token",
                                      style: TextStyle(
                                        fontSize: 14,
                                        color: const Color.fromARGB(
                                            255, 0, 84, 152),
                                      ),
                                    ),
                                  )
                                : Wrap(
                                    alignment: WrapAlignment.spaceBetween,
                                    crossAxisAlignment:
                                        WrapCrossAlignment.center,
                                    children: [
                                      !isSmallDisplay
                                          ? Wrap(
                                              crossAxisAlignment:
                                                  WrapCrossAlignment.center,
                                              children: [
                                                SizedBox(
                                                  height: 24,
                                                  width: 24,
                                                  child: Checkbox(
                                                    value: _isChecked,
                                                    onChanged: (bool? value) =>
                                                        setState(() =>
                                                            _isChecked =
                                                                value!),
                                                  ),
                                                ),
                                                const SizedBox(width: 8),
                                                Text(
                                                  localeMsg.stayLogged,
                                                  style: TextStyle(
                                                    fontSize: 14,
                                                    color: Colors.black,
                                                  ),
                                                ),
                                              ],
                                            )
                                          : Container(),
                                      TextButton(
                                        onPressed: () => setState(() {
                                          forgot = !forgot;
                                        }),
                                        child: Text(
                                          localeMsg.forgotPassword,
                                          style: TextStyle(
                                            fontSize: 14,
                                            color: const Color.fromARGB(
                                                255, 0, 84, 152),
                                          ),
                                        ),
                                      ),
                                    ],
                                  ),
                            SizedBox(
                                height:
                                    forgot ? 20 : (isSmallDisplay ? 15 : 30)),
                            Align(
                              child: ElevatedButton(
                                onPressed: () =>
                                    forgot ? resetPassword() : tryLogin(),
                                style: ElevatedButton.styleFrom(
                                  padding: const EdgeInsets.symmetric(
                                    vertical: 20,
                                    horizontal: 20,
                                  ),
                                ),
                                child: Text(
                                  forgot ? "Request Reset" : localeMsg.login,
                                  style: TextStyle(
                                    fontSize: 14,
                                    fontWeight: FontWeight.w600,
                                  ),
                                ),
                              ),
                            ),
                            const SizedBox(height: 15),
                          ],
                        ),
                      ),
                    ),
                  ),
                ),
              ],
            ),
          )
        ]),
      ),
    );
  }

  tryLogin() async {
    if (_formKey.currentState!.validate()) {
      _formKey.currentState!.save();
      final result = await loginAPI(_email!, _password!, userUrl: _apiUrl);
      switch (result) {
        case Success(value: final loginData):
          if (apiType == BackendType.tenant) {
            await fetchApiVersion(_apiUrl);
          }
          Navigator.of(context).push(
            MaterialPageRoute(
              builder: (context) => ProjectsPage(
                userEmail: loginData.first,
                isTenantMode: loginData[1] == "true",
              ),
            ),
          );
        case Failure(exception: final exception):
          String errorMsg = exception.toString() == "Exception"
              ? AppLocalizations.of(context)!.invalidLogin
              : exception.toString();
          showSnackBar(context, errorMsg, isError: true);
      }
    }
  }

  resetPassword() async {
    if (_formKey.currentState!.validate()) {
      _formKey.currentState!.save();
      final result = await userForgotPassword(_email!, userUrl: _apiUrl);
      switch (result) {
        case Success():
          showSnackBar(context, "Reset request sent", isSuccess: true);
        case Failure(exception: final exception):
          showSnackBar(context, exception.toString().trim(), isError: true);
      }
    }
  }

  backendInput() {
    List<String> options = [];
    if (dotenv.env['BACK_URLS'] != null) {
      options = dotenv.env['BACK_URLS']!.split(",");
    }
    final localeMsg = AppLocalizations.of(context)!;
    return RawAutocomplete<String>(
      optionsBuilder: (TextEditingValue textEditingValue) {
        return options.where((String option) {
          return option.contains(textEditingValue.text);
        });
      },
      fieldViewBuilder: (BuildContext context,
          TextEditingController textEditingController,
          FocusNode focusNode,
          VoidCallback onFieldSubmitted) {
        //textEditingController.text = options.first;
        return TextFormField(
          controller: textEditingController,
          focusNode: focusNode,
          onSaved: (newValue) => _apiUrl = newValue!,
          validator: (text) {
            if (text == null || text.trim().isEmpty) {
              return localeMsg.mandatoryField;
            }
            return null;
          },
          decoration: InputDecoration(
              isDense: true,
              labelText: localeMsg.selectServer,
              labelStyle: TextStyle(fontSize: 14)),
          onTap: () {
            setState(() {
              apiType = null;
            });
          },
          onEditingComplete: () => getBackendType(textEditingController.text),
        );
      },
      optionsViewBuilder: (BuildContext context,
          AutocompleteOnSelected<String> onSelected, Iterable<String> options) {
        return Align(
          alignment: Alignment.topLeft,
          child: Material(
            elevation: 4.0,
            child: SizedBox(
              height: options.length > 2 ? 171.0 : 57.0 * options.length,
              width: 350,
              child: ListView.builder(
                padding: const EdgeInsets.all(8.0),
                itemCount: options.length,
                itemBuilder: (BuildContext context, int index) {
                  final String option = options.elementAt(index);
                  return GestureDetector(
                    onTap: () async {
                      getBackendType(option);
                      onSelected(option);
                    },
                    child: ListTile(
                      title: Text(option, style: const TextStyle(fontSize: 14)),
                    ),
                  );
                },
              ),
            ),
          ),
        );
      },
    );
  }

  getBackendType(inputUrl) async {
    final result = await fetchApiVersion(inputUrl);
    switch (result) {
      case Success(value: final type):
        setState(() {
          apiType = type;
        });
      case Failure(exception: final exception):
        print(exception);
        setState(() {
          apiType = BackendType.unavailable;
        });
    }
  }

  getBackendTypeText() {
    if (apiType == null) {
      return "";
    } else if (apiType == BackendType.unavailable) {
      return AppLocalizations.of(context)!.unavailable.toUpperCase();
    } else {
      return "${apiType!.name.toUpperCase()} SERVER";
    }
  }
}