# vkImportSemaphoreFdKHR(3) Manual Page

## Name

vkImportSemaphoreFdKHR - Import a semaphore from a POSIX file descriptor



## [](#_c_specification)C Specification

To import a semaphore payload from a POSIX file descriptor, call:

```c++
// Provided by VK_KHR_external_semaphore_fd
VkResult vkImportSemaphoreFdKHR(
    VkDevice                                    device,
    const VkImportSemaphoreFdInfoKHR*           pImportSemaphoreFdInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the semaphore.
- `pImportSemaphoreFdInfo` is a pointer to a [VkImportSemaphoreFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreFdInfoKHR.html) structure specifying the semaphore and import parameters.

## [](#_description)Description

Importing a semaphore payload from a file descriptor transfers ownership of the file descriptor from the application to the Vulkan implementation. The application **must** not perform any operations on the file descriptor after a successful import.

Applications **can** import the same semaphore payload into multiple instances of Vulkan, into the same instance from which it was exported, and multiple times into a given Vulkan instance.

Valid Usage

- [](#VUID-vkImportSemaphoreFdKHR-semaphore-01142)VUID-vkImportSemaphoreFdKHR-semaphore-01142  
  `semaphore` **must** not be associated with any queue command that has not yet completed execution on that queue

Valid Usage (Implicit)

- [](#VUID-vkImportSemaphoreFdKHR-device-parameter)VUID-vkImportSemaphoreFdKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkImportSemaphoreFdKHR-pImportSemaphoreFdInfo-parameter)VUID-vkImportSemaphoreFdKHR-pImportSemaphoreFdInfo-parameter  
  `pImportSemaphoreFdInfo` **must** be a valid pointer to a valid [VkImportSemaphoreFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreFdInfoKHR.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_INVALID_EXTERNAL_HANDLE`

## [](#_see_also)See Also

[VK\_KHR\_external\_semaphore\_fd](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_semaphore_fd.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImportSemaphoreFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreFdInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkImportSemaphoreFdKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0