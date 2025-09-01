# vkGetMemoryFdKHR(3) Manual Page

## Name

vkGetMemoryFdKHR - Get a POSIX file descriptor for a memory object



## [](#_c_specification)C Specification

To export a POSIX file descriptor referencing the payload of a Vulkan device memory object, call:

```c++
// Provided by VK_KHR_external_memory_fd
VkResult vkGetMemoryFdKHR(
    VkDevice                                    device,
    const VkMemoryGetFdInfoKHR*                 pGetFdInfo,
    int*                                        pFd);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the device memory being exported.
- `pGetFdInfo` is a pointer to a [VkMemoryGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetFdInfoKHR.html) structure containing parameters of the export operation.
- `pFd` will return a file descriptor referencing the payload of the device memory object.

## [](#_description)Description

Each call to `vkGetMemoryFdKHR` **must** create a new file descriptor holding a reference to the memory object’s payload and transfer ownership of the file descriptor to the application. To avoid leaking resources, the application **must** release ownership of the file descriptor using the `close` system call when it is no longer needed, or by importing a Vulkan memory object from it. Where supported by the operating system, the implementation **must** set the file descriptor to be closed automatically when an `execve` system call is made.

Valid Usage (Implicit)

- [](#VUID-vkGetMemoryFdKHR-device-parameter)VUID-vkGetMemoryFdKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetMemoryFdKHR-pGetFdInfo-parameter)VUID-vkGetMemoryFdKHR-pGetFdInfo-parameter  
  `pGetFdInfo` **must** be a valid pointer to a valid [VkMemoryGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetFdInfoKHR.html) structure
- [](#VUID-vkGetMemoryFdKHR-pFd-parameter)VUID-vkGetMemoryFdKHR-pFd-parameter  
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

[VK\_KHR\_external\_memory\_fd](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory_fd.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMemoryGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetFdInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetMemoryFdKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0