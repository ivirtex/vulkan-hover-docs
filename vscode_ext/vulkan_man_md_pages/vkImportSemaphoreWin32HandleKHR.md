# vkImportSemaphoreWin32HandleKHR(3) Manual Page

## Name

vkImportSemaphoreWin32HandleKHR - Import a semaphore from a Windows
HANDLE



## <a href="#_c_specification" class="anchor"></a>C Specification

To import a semaphore payload from a Windows handle, call:

``` c
// Provided by VK_KHR_external_semaphore_win32
VkResult vkImportSemaphoreWin32HandleKHR(
    VkDevice                                    device,
    const VkImportSemaphoreWin32HandleInfoKHR*  pImportSemaphoreWin32HandleInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that created the semaphore.

- `pImportSemaphoreWin32HandleInfo` is a pointer to a
  [VkImportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreWin32HandleInfoKHR.html)
  structure specifying the semaphore and import parameters.

## <a href="#_description" class="anchor"></a>Description

Importing a semaphore payload from Windows handles does not transfer
ownership of the handle to the Vulkan implementation. For handle types
defined as NT handles, the application **must** release ownership using
the `CloseHandle` system call when the handle is no longer needed.

Applications **can** import the same semaphore payload into multiple
instances of Vulkan, into the same instance from which it was exported,
and multiple times into a given Vulkan instance.

Valid Usage (Implicit)

- <a href="#VUID-vkImportSemaphoreWin32HandleKHR-device-parameter"
  id="VUID-vkImportSemaphoreWin32HandleKHR-device-parameter"></a>
  VUID-vkImportSemaphoreWin32HandleKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkImportSemaphoreWin32HandleKHR-pImportSemaphoreWin32HandleInfo-parameter"
  id="VUID-vkImportSemaphoreWin32HandleKHR-pImportSemaphoreWin32HandleInfo-parameter"></a>
  VUID-vkImportSemaphoreWin32HandleKHR-pImportSemaphoreWin32HandleInfo-parameter  
  `pImportSemaphoreWin32HandleInfo` **must** be a valid pointer to a
  valid
  [VkImportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreWin32HandleInfoKHR.html)
  structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_semaphore_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore_win32.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkImportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreWin32HandleInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkImportSemaphoreWin32HandleKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
