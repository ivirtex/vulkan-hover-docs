# vkImportSemaphoreFdKHR(3) Manual Page

## Name

vkImportSemaphoreFdKHR - Import a semaphore from a POSIX file descriptor



## <a href="#_c_specification" class="anchor"></a>C Specification

To import a semaphore payload from a POSIX file descriptor, call:

``` c
// Provided by VK_KHR_external_semaphore_fd
VkResult vkImportSemaphoreFdKHR(
    VkDevice                                    device,
    const VkImportSemaphoreFdInfoKHR*           pImportSemaphoreFdInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that created the semaphore.

- `pImportSemaphoreFdInfo` is a pointer to a
  [VkImportSemaphoreFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreFdInfoKHR.html)
  structure specifying the semaphore and import parameters.

## <a href="#_description" class="anchor"></a>Description

Importing a semaphore payload from a file descriptor transfers ownership
of the file descriptor from the application to the Vulkan
implementation. The application **must** not perform any operations on
the file descriptor after a successful import.

Applications **can** import the same semaphore payload into multiple
instances of Vulkan, into the same instance from which it was exported,
and multiple times into a given Vulkan instance.

Valid Usage

- <a href="#VUID-vkImportSemaphoreFdKHR-semaphore-01142"
  id="VUID-vkImportSemaphoreFdKHR-semaphore-01142"></a>
  VUID-vkImportSemaphoreFdKHR-semaphore-01142  
  `semaphore` **must** not be associated with any queue command that has
  not yet completed execution on that queue

Valid Usage (Implicit)

- <a href="#VUID-vkImportSemaphoreFdKHR-device-parameter"
  id="VUID-vkImportSemaphoreFdKHR-device-parameter"></a>
  VUID-vkImportSemaphoreFdKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkImportSemaphoreFdKHR-pImportSemaphoreFdInfo-parameter"
  id="VUID-vkImportSemaphoreFdKHR-pImportSemaphoreFdInfo-parameter"></a>
  VUID-vkImportSemaphoreFdKHR-pImportSemaphoreFdInfo-parameter  
  `pImportSemaphoreFdInfo` **must** be a valid pointer to a valid
  [VkImportSemaphoreFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreFdInfoKHR.html)
  structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_semaphore_fd](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore_fd.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkImportSemaphoreFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreFdInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkImportSemaphoreFdKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
