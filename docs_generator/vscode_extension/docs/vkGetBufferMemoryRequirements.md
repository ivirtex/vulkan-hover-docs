# vkGetBufferMemoryRequirements(3) Manual Page

## Name

vkGetBufferMemoryRequirements - Returns the memory requirements for specified Vulkan object



## [](#_c_specification)C Specification

To determine the memory requirements for a buffer resource, call:

```c++
// Provided by VK_VERSION_1_0
void vkGetBufferMemoryRequirements(
    VkDevice                                    device,
    VkBuffer                                    buffer,
    VkMemoryRequirements*                       pMemoryRequirements);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the buffer.
- `buffer` is the buffer to query.
- `pMemoryRequirements` is a pointer to a [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure in which the memory requirements of the buffer object are returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetBufferMemoryRequirements-device-parameter)VUID-vkGetBufferMemoryRequirements-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetBufferMemoryRequirements-buffer-parameter)VUID-vkGetBufferMemoryRequirements-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-vkGetBufferMemoryRequirements-pMemoryRequirements-parameter)VUID-vkGetBufferMemoryRequirements-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure
- [](#VUID-vkGetBufferMemoryRequirements-buffer-parent)VUID-vkGetBufferMemoryRequirements-buffer-parent  
  `buffer` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetBufferMemoryRequirements)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0