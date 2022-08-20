# Errata for *Pro Go*

This file contains errors that are likely to prevent the code examples from working as described in ths book. See [this](typos.md) file for small mistakes that I intend to correct in the next edition.

---

**Chapter 33**

On Page 909, the getAnonymousFieldMethods function is shown with the wrong parameters and result. This statement:

    func getAnonymousFieldMethods(target []reflect.Type) reflect.Method {

should be:

    func getAnonymousFieldMethods(target reflect.Type) []reflect.Method {


(Thanks to Yaroslav Lazarev for reporting this problem)

---