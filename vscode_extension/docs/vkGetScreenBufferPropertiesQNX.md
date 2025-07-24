# vkGetScreenBufferPropertiesQNX(3) Manual Page

## Name

vkGetScreenBufferPropertiesQNX - Get Properties of External Memory QNX Screen Buffers



## [](#_c_specification)C Specification

To determine the memory parameters to use when importing a QNX Screen buffer, call:

```c++
// Provided by VK_QNX_external_memory_screen_buffer
VkResult vkGetScreenBufferPropertiesQNX(
    VkDevice                                    device,
    const struct _screen_buffer*                buffer,
    VkScreenBufferPropertiesQNX*                pProperties);
```

## [](#_parameters)Parameters

- `device` is the logical device that will be importing `buffer`.
- `buffer` is the QNX Screen buffer which will be imported.
- `pProperties` is a pointer to a [VkScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenBufferPropertiesQNX.html) structure in which the properties of `buffer` are returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetScreenBufferPropertiesQNX-buffer-08968)VUID-vkGetScreenBufferPropertiesQNX-buffer-08968  
  `buffer` **must** be a [valid QNX Screen buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-external-screen-buffer-validity)

Valid Usage (Implicit)

- [](#VUID-vkGetScreenBufferPropertiesQNX-device-parameter)VUID-vkGetScreenBufferPropertiesQNX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetScreenBufferPropertiesQNX-buffer-parameter)VUID-vkGetScreenBufferPropertiesQNX-buffer-parameter  
  `buffer` **must** be a valid pointer to a valid `_screen_buffer` value
- [](#VUID-vkGetScreenBufferPropertiesQNX-pProperties-parameter)VUID-vkGetScreenBufferPropertiesQNX-pProperties-parameter  
  `pProperties` **must** be a valid pointer to a [VkScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenBufferPropertiesQNX.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_EXTERNAL_HANDLE_KHR`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_QNX\_external\_memory\_screen\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QNX_external_memory_screen_buffer.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenBufferPropertiesQNX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetScreenBufferPropertiesQNX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0