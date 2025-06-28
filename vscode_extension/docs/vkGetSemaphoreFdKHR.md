# vkGetSemaphoreFdKHR(3) Manual Page

## Name

vkGetSemaphoreFdKHR - Get a POSIX file descriptor handle for a semaphore



## [](#_c_specification)C Specification

To export a POSIX file descriptor representing the payload of a semaphore, call:

```c++
// Provided by VK_KHR_external_semaphore_fd
VkResult vkGetSemaphoreFdKHR(
    VkDevice                                    device,
    const VkSemaphoreGetFdInfoKHR*              pGetFdInfo,
    int*                                        pFd);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the semaphore being exported.
- `pGetFdInfo` is a pointer to a [VkSemaphoreGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreGetFdInfoKHR.html) structure containing parameters of the export operation.
- `pFd` will return the file descriptor representing the semaphore payload.

## [](#_description)Description

Each call to `vkGetSemaphoreFdKHR` **must** create a new file descriptor and transfer ownership of it to the application. To avoid leaking resources, the application **must** release ownership of the file descriptor when it is no longer needed.

Note

Ownership can be released in many ways. For example, the application can call `close`() on the file descriptor, or transfer ownership back to Vulkan by using the file descriptor to import a semaphore payload.

Where supported by the operating system, the implementation **must** set the file descriptor to be closed automatically when an `execve` system call is made.

Exporting a file descriptor from a semaphore **may** have side effects depending on the transference of the specified handle type, as described in [Importing Semaphore State](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-importing).

Valid Usage (Implicit)

- [](#VUID-vkGetSemaphoreFdKHR-device-parameter)VUID-vkGetSemaphoreFdKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetSemaphoreFdKHR-pGetFdInfo-parameter)VUID-vkGetSemaphoreFdKHR-pGetFdInfo-parameter  
  `pGetFdInfo` **must** be a valid pointer to a valid [VkSemaphoreGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreGetFdInfoKHR.html) structure
- [](#VUID-vkGetSemaphoreFdKHR-pFd-parameter)VUID-vkGetSemaphoreFdKHR-pFd-parameter  
  `pFd` **must** be a valid pointer to an `int` value

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_TOO_MANY_OBJECTS`
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## [](#_see_also)See Also

[VK\_KHR\_external\_semaphore\_fd](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_semaphore_fd.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSemaphoreGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreGetFdInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetSemaphoreFdKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0