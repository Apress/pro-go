# Typos for *Pro Go*

This file contains smaller errors that are unlikely to prevent the code examples from working and which I will correct in the next edition of the book. See [this](errata.md) file for mistakes that prevent the code examples from working as expected.

---

**Chapter 4**

On page 70, the title for Listing 4-12 is:

> Using Constants in the main.go File in the basicFeatures Folder

It should be:

> Using ***Variables*** in the main.go File in the basicFeatures Folder

(Thanks to Michael for reporting this problem)

---

**Chapter 5**

Table 5-7 contains the following text in the description for the `Floor` function:

>   The largest integer that is less than 27.1, for example, is 28.

It should be:

>   The largest integer that is less than 27.1, for example, is ***27***.

(Thanks to Francisco Vicedo and Korbai Zoltán for reporting this problem)

---

**Chapter 6**

On page 125, Figure 6-5 uses a semi-colon instead of a colon in the exploding view of the statement.

(Thanks to Saleh for reporting this problem)

---

**Chapter 7**

On page 163, Listing 7-18 contains this statement:

    moreNames := []string { "Hat Gloves"}

It should be:

    moreNames := []string { "Hat", "Gloves"}

(Thanks to Francisco Vicedo for reporting this problem)

---

**Chapter 8**

On page 214, the calcTotalPrice function doesn't calculate a total price. It actually just adds the prices to the minSpend amount.

(Thanks to Francisco Vicedo for reporting this problem)

---

On page 217, the description of the results order doesn't match the actual output.

(Thanks to Francisco Vicedo for reporting this problem)

---

**Chapter 9**

On page 223, listing 9-5 contains this statement in two locations;

     fmt.Println("Function assigned:", calcFunc == nil)

This is confusing and should be:

     fmt.Println("Function unassigned:", calcFunc == nil)


(Thanks to Francisco Vicedo for reporting this problem)

---

**Chapter 10**

On page 272, the summary starts with:

> In this chapter, I describe the Go struct featur

This should be:

>   In this chapter, I ***described*** the Go struct featur

(Thanks to Francisco Vicedo for reporting this problem)

---

**Chapter 11**

On page 289, this text is incorrect:

> This interface has been given the name Expense, and the interface body contains a single method signature. 

The interface defines two methods.

(Thanks to Joseph Kashi and Francisco Vicedo for reporting this problem)

---

**Chapter 12**

On page 319, the wrong chapter is specified for the description of the `FormatFloat` function. It should be Chapter 5.

(Thanks to Francisco Vicedo for reporting this problem)


---

**Chapter 13**

On page 336, the text:

> The *Product field is embedded, which means that its name ***its*** its type

should be:

> The *Product field is embedded, which means that its name ***is*** its type

(Thanks to Joseph Kashi and Francisco Vicedo for reporting this problem)

---

**Chapter 16**

On page 417, Table 16-4 gives the impression that the `Title` function transforms a string value, rather than creating a new string.

(Thanks to Francisco Vicedo for reporting this problem)

---

On Page 421, Listing 16-8 contains two statements that use the `LastIndex` function and omits the `LastIndexAny` function.

(Thanks to Francisco Vicedo for reporting this problem)
