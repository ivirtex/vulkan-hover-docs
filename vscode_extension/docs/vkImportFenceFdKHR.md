# vkImportFenceFdKHR(3) Manual Page

## Name

vkImportFenceFdKHR - Import a fence from a POSIX file descriptor



## [](#_c_specification)C Specification

To import a fence payload from a POSIX file descriptor, call:

```c++
// Provided by VK_KHR_external_fence_fd
VkResult vkImportFenceFdKHR(
    VkDevice                                    device,
    const VkImportFenceFdInfoKHR*               pImportFenceFdInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the fence.
- `pImportFenceFdInfo` is a pointer to a [VkImportFenceFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportFenceFdInfoKHR.html) structure specifying the fence and import parameters.

## [](#_description)Description

Importing a fence payload from a file descriptor transfers ownership of the file descriptor from the application to the Vulkan implementation. The application **must** not perform any operations on the file descriptor after a successful import.

Applications **can** import the same fence payload into multiple instances of Vulkan, into the same instance from which it was exported, and multiple times into a given Vulkan instance.

Valid Usage

- [](#VUID-vkImportFenceFdKHR-fence-01463)VUID-vkImportFenceFdKHR-fence-01463  
  `fence` **must** not be associated with any queue command that has not yet completed execution on that queue

Valid Usage (Implicit)

- [](#VUID-vkImportFenceFdKHR-device-parameter)VUID-vkImportFenceFdKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkImportFenceFdKHR-pImportFenceFdInfo-parameter)VUID-vkImportFenceFdKHR-pImportFenceFdInfo-parameter  
  `pImportFenceFdInfo` **must** be a valid pointer to a valid [VkImportFenceFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportFenceFdInfoKHR.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_external\_fence\_fd](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_fence_fd.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImportFenceFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportFenceFdInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkImportFenceFdKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0