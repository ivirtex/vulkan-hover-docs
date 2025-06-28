# vkGetMemoryWin32HandlePropertiesKHR(3) Manual Page

## Name

vkGetMemoryWin32HandlePropertiesKHR - Get Properties of External Memory Win32 Handles



## [](#_c_specification)C Specification

Windows memory handles compatible with Vulkan **may** also be created by non-Vulkan APIs using methods beyond the scope of this specification. To determine the correct parameters to use when importing such handles, call:

```c++
// Provided by VK_KHR_external_memory_win32
VkResult vkGetMemoryWin32HandlePropertiesKHR(
    VkDevice                                    device,
    VkExternalMemoryHandleTypeFlagBits          handleType,
    HANDLE                                      handle,
    VkMemoryWin32HandlePropertiesKHR*           pMemoryWin32HandleProperties);
```

## [](#_parameters)Parameters

- `device` is the logical device that will be importing `handle`.
- `handleType` is a [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value specifying the type of the handle `handle`.
- `handle` is the handle which will be imported.
- `pMemoryWin32HandleProperties` is a pointer to a [VkMemoryWin32HandlePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryWin32HandlePropertiesKHR.html) structure in which properties of `handle` are returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetMemoryWin32HandlePropertiesKHR-handle-00665)VUID-vkGetMemoryWin32HandlePropertiesKHR-handle-00665  
  `handle` **must** point to a valid Windows memory handle
- [](#VUID-vkGetMemoryWin32HandlePropertiesKHR-handleType-00666)VUID-vkGetMemoryWin32HandlePropertiesKHR-handleType-00666  
  `handleType` **must** not be one of the handle types defined as opaque

Valid Usage (Implicit)

- [](#VUID-vkGetMemoryWin32HandlePropertiesKHR-device-parameter)VUID-vkGetMemoryWin32HandlePropertiesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetMemoryWin32HandlePropertiesKHR-handleType-parameter)VUID-vkGetMemoryWin32HandlePropertiesKHR-handleType-parameter  
  `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value
- [](#VUID-vkGetMemoryWin32HandlePropertiesKHR-pMemoryWin32HandleProperties-parameter)VUID-vkGetMemoryWin32HandlePropertiesKHR-pMemoryWin32HandleProperties-parameter  
  `pMemoryWin32HandleProperties` **must** be a valid pointer to a [VkMemoryWin32HandlePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryWin32HandlePropertiesKHR.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_INVALID_EXTERNAL_HANDLE`

## [](#_see_also)See Also

[VK\_KHR\_external\_memory\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory_win32.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkMemoryWin32HandlePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryWin32HandlePropertiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetMemoryWin32HandlePropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0