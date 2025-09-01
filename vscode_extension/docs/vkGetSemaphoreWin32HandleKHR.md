# vkGetSemaphoreWin32HandleKHR(3) Manual Page

## Name

vkGetSemaphoreWin32HandleKHR - Get a Windows HANDLE for a semaphore



## [](#_c_specification)C Specification

To export a Windows handle representing the payload of a semaphore, call:

```c++
// Provided by VK_KHR_external_semaphore_win32
VkResult vkGetSemaphoreWin32HandleKHR(
    VkDevice                                    device,
    const VkSemaphoreGetWin32HandleInfoKHR*     pGetWin32HandleInfo,
    HANDLE*                                     pHandle);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the semaphore being exported.
- `pGetWin32HandleInfo` is a pointer to a [VkSemaphoreGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreGetWin32HandleInfoKHR.html) structure containing parameters of the export operation.
- `pHandle` will return the Windows handle representing the semaphore state.

## [](#_description)Description

For handle types defined as NT handles, the handles returned by `vkGetSemaphoreWin32HandleKHR` are owned by the application. To avoid leaking resources, the application **must** release ownership of them using the `CloseHandle` system call when they are no longer needed.

Exporting a Windows handle from a semaphore **may** have side effects depending on the transference of the specified handle type, as described in [Importing Semaphore Payloads](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-importing).

Valid Usage (Implicit)

- [](#VUID-vkGetSemaphoreWin32HandleKHR-device-parameter)VUID-vkGetSemaphoreWin32HandleKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetSemaphoreWin32HandleKHR-pGetWin32HandleInfo-parameter)VUID-vkGetSemaphoreWin32HandleKHR-pGetWin32HandleInfo-parameter  
  `pGetWin32HandleInfo` **must** be a valid pointer to a valid [VkSemaphoreGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreGetWin32HandleInfoKHR.html) structure
- [](#VUID-vkGetSemaphoreWin32HandleKHR-pHandle-parameter)VUID-vkGetSemaphoreWin32HandleKHR-pHandle-parameter  
  `pHandle` **must** be a valid pointer to a `HANDLE` value

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_TOO_MANY_OBJECTS`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_external\_semaphore\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_semaphore_win32.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSemaphoreGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreGetWin32HandleInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetSemaphoreWin32HandleKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0