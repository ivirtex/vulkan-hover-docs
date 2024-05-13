# vkImportFenceFdKHR(3) Manual Page

## Name

vkImportFenceFdKHR - Import a fence from a POSIX file descriptor



## <a href="#_c_specification" class="anchor"></a>C Specification

To import a fence payload from a POSIX file descriptor, call:

``` c
// Provided by VK_KHR_external_fence_fd
VkResult vkImportFenceFdKHR(
    VkDevice                                    device,
    const VkImportFenceFdInfoKHR*               pImportFenceFdInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that created the fence.

- `pImportFenceFdInfo` is a pointer to a
  [VkImportFenceFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportFenceFdInfoKHR.html) structure
  specifying the fence and import parameters.

## <a href="#_description" class="anchor"></a>Description

Importing a fence payload from a file descriptor transfers ownership of
the file descriptor from the application to the Vulkan implementation.
The application **must** not perform any operations on the file
descriptor after a successful import.

Applications **can** import the same fence payload into multiple
instances of Vulkan, into the same instance from which it was exported,
and multiple times into a given Vulkan instance.

Valid Usage

- <a href="#VUID-vkImportFenceFdKHR-fence-01463"
  id="VUID-vkImportFenceFdKHR-fence-01463"></a>
  VUID-vkImportFenceFdKHR-fence-01463  
  `fence` **must** not be associated with any queue command that has not
  yet completed execution on that queue

Valid Usage (Implicit)

- <a href="#VUID-vkImportFenceFdKHR-device-parameter"
  id="VUID-vkImportFenceFdKHR-device-parameter"></a>
  VUID-vkImportFenceFdKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkImportFenceFdKHR-pImportFenceFdInfo-parameter"
  id="VUID-vkImportFenceFdKHR-pImportFenceFdInfo-parameter"></a>
  VUID-vkImportFenceFdKHR-pImportFenceFdInfo-parameter  
  `pImportFenceFdInfo` **must** be a valid pointer to a valid
  [VkImportFenceFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportFenceFdInfoKHR.html) structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_fence_fd](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_fence_fd.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkImportFenceFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportFenceFdInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkImportFenceFdKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
