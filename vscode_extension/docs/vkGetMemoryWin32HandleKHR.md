# vkGetMemoryWin32HandleKHR(3) Manual Page

## Name

vkGetMemoryWin32HandleKHR - Get a Windows HANDLE for a memory object



## [](#_c_specification)C Specification

To export a Windows handle representing the payload of a Vulkan device memory object, call:

```c++
// Provided by VK_KHR_external_memory_win32
VkResult vkGetMemoryWin32HandleKHR(
    VkDevice                                    device,
    const VkMemoryGetWin32HandleInfoKHR*        pGetWin32HandleInfo,
    HANDLE*                                     pHandle);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the device memory being exported.
- `pGetWin32HandleInfo` is a pointer to a [VkMemoryGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetWin32HandleInfoKHR.html) structure containing parameters of the export operation.
- `pHandle` will return the Windows handle representing the payload of the device memory object.

## [](#_description)Description

For handle types defined as NT handles, the handles returned by `vkGetMemoryWin32HandleKHR` are owned by the application and hold a reference to their payload. To avoid leaking resources, the application **must** release ownership of them using the `CloseHandle` system call when they are no longer needed.

Note

Non-NT handle types do not add a reference to their associated payload. If the original object owning the payload is destroyed, all resources and handles sharing that payload will become invalid.

Valid Usage (Implicit)

- [](#VUID-vkGetMemoryWin32HandleKHR-device-parameter)VUID-vkGetMemoryWin32HandleKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetMemoryWin32HandleKHR-pGetWin32HandleInfo-parameter)VUID-vkGetMemoryWin32HandleKHR-pGetWin32HandleInfo-parameter  
  `pGetWin32HandleInfo` **must** be a valid pointer to a valid [VkMemoryGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetWin32HandleInfoKHR.html) structure
- [](#VUID-vkGetMemoryWin32HandleKHR-pHandle-parameter)VUID-vkGetMemoryWin32HandleKHR-pHandle-parameter  
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

[VK\_KHR\_external\_memory\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory_win32.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMemoryGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetWin32HandleInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetMemoryWin32HandleKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0