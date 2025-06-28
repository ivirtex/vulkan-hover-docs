# vkGetMemoryFdPropertiesKHR(3) Manual Page

## Name

vkGetMemoryFdPropertiesKHR - Get Properties of External Memory File Descriptors



## [](#_c_specification)C Specification

POSIX file descriptor memory handles compatible with Vulkan **may** also be created by non-Vulkan APIs using methods beyond the scope of this specification. To determine the correct parameters to use when importing such handles, call:

```c++
// Provided by VK_KHR_external_memory_fd
VkResult vkGetMemoryFdPropertiesKHR(
    VkDevice                                    device,
    VkExternalMemoryHandleTypeFlagBits          handleType,
    int                                         fd,
    VkMemoryFdPropertiesKHR*                    pMemoryFdProperties);
```

## [](#_parameters)Parameters

- `device` is the logical device that will be importing `fd`.
- `handleType` is a [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value specifying the type of the handle `fd`.
- `fd` is the handle which will be imported.
- `pMemoryFdProperties` is a pointer to a [VkMemoryFdPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryFdPropertiesKHR.html) structure in which the properties of the handle `fd` are returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetMemoryFdPropertiesKHR-fd-00673)VUID-vkGetMemoryFdPropertiesKHR-fd-00673  
  `fd` **must** point to a valid POSIX file descriptor memory handle
- [](#VUID-vkGetMemoryFdPropertiesKHR-handleType-00674)VUID-vkGetMemoryFdPropertiesKHR-handleType-00674  
  `handleType` **must** not be `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT`

Valid Usage (Implicit)

- [](#VUID-vkGetMemoryFdPropertiesKHR-device-parameter)VUID-vkGetMemoryFdPropertiesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetMemoryFdPropertiesKHR-handleType-parameter)VUID-vkGetMemoryFdPropertiesKHR-handleType-parameter  
  `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value
- [](#VUID-vkGetMemoryFdPropertiesKHR-pMemoryFdProperties-parameter)VUID-vkGetMemoryFdPropertiesKHR-pMemoryFdProperties-parameter  
  `pMemoryFdProperties` **must** be a valid pointer to a [VkMemoryFdPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryFdPropertiesKHR.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_INVALID_EXTERNAL_HANDLE`

## [](#_see_also)See Also

[VK\_KHR\_external\_memory\_fd](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory_fd.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkMemoryFdPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryFdPropertiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetMemoryFdPropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0