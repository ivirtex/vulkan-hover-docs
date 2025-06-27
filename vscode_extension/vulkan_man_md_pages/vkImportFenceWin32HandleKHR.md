# vkImportFenceWin32HandleKHR(3) Manual Page

## Name

vkImportFenceWin32HandleKHR - Import a fence from a Windows HANDLE



## <a href="#_c_specification" class="anchor"></a>C Specification

To import a fence payload from a Windows handle, call:

``` c
// Provided by VK_KHR_external_fence_win32
VkResult vkImportFenceWin32HandleKHR(
    VkDevice                                    device,
    const VkImportFenceWin32HandleInfoKHR*      pImportFenceWin32HandleInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that created the fence.

- `pImportFenceWin32HandleInfo` is a pointer to a
  [VkImportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportFenceWin32HandleInfoKHR.html)
  structure specifying the fence and import parameters.

## <a href="#_description" class="anchor"></a>Description

Importing a fence payload from Windows handles does not transfer
ownership of the handle to the Vulkan implementation. For handle types
defined as NT handles, the application **must** release ownership using
the `CloseHandle` system call when the handle is no longer needed.

Applications **can** import the same fence payload into multiple
instances of Vulkan, into the same instance from which it was exported,
and multiple times into a given Vulkan instance.

Valid Usage

- <a href="#VUID-vkImportFenceWin32HandleKHR-fence-04448"
  id="VUID-vkImportFenceWin32HandleKHR-fence-04448"></a>
  VUID-vkImportFenceWin32HandleKHR-fence-04448  
  `fence` **must** not be associated with any queue command that has not
  yet completed execution on that queue

Valid Usage (Implicit)

- <a href="#VUID-vkImportFenceWin32HandleKHR-device-parameter"
  id="VUID-vkImportFenceWin32HandleKHR-device-parameter"></a>
  VUID-vkImportFenceWin32HandleKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkImportFenceWin32HandleKHR-pImportFenceWin32HandleInfo-parameter"
  id="VUID-vkImportFenceWin32HandleKHR-pImportFenceWin32HandleInfo-parameter"></a>
  VUID-vkImportFenceWin32HandleKHR-pImportFenceWin32HandleInfo-parameter  
  `pImportFenceWin32HandleInfo` **must** be a valid pointer to a valid
  [VkImportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportFenceWin32HandleInfoKHR.html)
  structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_fence_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_fence_win32.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkImportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportFenceWin32HandleInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkImportFenceWin32HandleKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
