# vkImportSemaphoreWin32HandleKHR(3) Manual Page

## Name

vkImportSemaphoreWin32HandleKHR - Import a semaphore from a Windows HANDLE



## [](#_c_specification)C Specification

To import a semaphore payload from a Windows handle, call:

```c++
// Provided by VK_KHR_external_semaphore_win32
VkResult vkImportSemaphoreWin32HandleKHR(
    VkDevice                                    device,
    const VkImportSemaphoreWin32HandleInfoKHR*  pImportSemaphoreWin32HandleInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the semaphore.
- `pImportSemaphoreWin32HandleInfo` is a pointer to a [VkImportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreWin32HandleInfoKHR.html) structure specifying the semaphore and import parameters.

## [](#_description)Description

Importing a semaphore payload from Windows handles does not transfer ownership of the handle to the Vulkan implementation. For handle types defined as NT handles, the application **must** release ownership using the `CloseHandle` system call when the handle is no longer needed.

Applications **can** import the same semaphore payload into multiple instances of Vulkan, into the same instance from which it was exported, and multiple times into a given Vulkan instance.

Valid Usage (Implicit)

- [](#VUID-vkImportSemaphoreWin32HandleKHR-device-parameter)VUID-vkImportSemaphoreWin32HandleKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkImportSemaphoreWin32HandleKHR-pImportSemaphoreWin32HandleInfo-parameter)VUID-vkImportSemaphoreWin32HandleKHR-pImportSemaphoreWin32HandleInfo-parameter  
  `pImportSemaphoreWin32HandleInfo` **must** be a valid pointer to a valid [VkImportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreWin32HandleInfoKHR.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_external\_semaphore\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_semaphore_win32.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreWin32HandleInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkImportSemaphoreWin32HandleKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0