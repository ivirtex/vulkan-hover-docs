# vkGetFenceFdKHR(3) Manual Page

## Name

vkGetFenceFdKHR - Get a POSIX file descriptor handle for a fence



## [](#_c_specification)C Specification

To export a POSIX file descriptor representing the payload of a fence, call:

```c++
// Provided by VK_KHR_external_fence_fd
VkResult vkGetFenceFdKHR(
    VkDevice                                    device,
    const VkFenceGetFdInfoKHR*                  pGetFdInfo,
    int*                                        pFd);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the fence being exported.
- `pGetFdInfo` is a pointer to a [VkFenceGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceGetFdInfoKHR.html) structure containing parameters of the export operation.
- `pFd` will return the file descriptor representing the fence payload.

## [](#_description)Description

Each call to `vkGetFenceFdKHR` **must** create a new file descriptor and transfer ownership of it to the application. To avoid leaking resources, the application **must** release ownership of the file descriptor when it is no longer needed.

Note

Ownership can be released in many ways. For example, the application can call `close`() on the file descriptor, or transfer ownership back to Vulkan by using the file descriptor to import a fence payload.

If `pGetFdInfo->handleType` is `VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT` and the fence is signaled at the time `vkGetFenceFdKHR` is called, `pFd` **may** return the value `-1` instead of a valid file descriptor.

Where supported by the operating system, the implementation **must** set the file descriptor to be closed automatically when an `execve` system call is made.

Exporting a file descriptor from a fence **may** have side effects depending on the transference of the specified handle type, as described in [Importing Fence State](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-fences-importing).

Valid Usage (Implicit)

- [](#VUID-vkGetFenceFdKHR-device-parameter)VUID-vkGetFenceFdKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetFenceFdKHR-pGetFdInfo-parameter)VUID-vkGetFenceFdKHR-pGetFdInfo-parameter  
  `pGetFdInfo` **must** be a valid pointer to a valid [VkFenceGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceGetFdInfoKHR.html) structure
- [](#VUID-vkGetFenceFdKHR-pFd-parameter)VUID-vkGetFenceFdKHR-pFd-parameter  
  `pFd` **must** be a valid pointer to an `int` value

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_TOO_MANY_OBJECTS`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_external\_fence\_fd](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_fence_fd.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkFenceGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceGetFdInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetFenceFdKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0