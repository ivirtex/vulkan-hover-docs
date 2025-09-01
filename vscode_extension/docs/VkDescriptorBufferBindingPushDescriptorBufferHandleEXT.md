# VkDescriptorBufferBindingPushDescriptorBufferHandleEXT(3) Manual Page

## Name

VkDescriptorBufferBindingPushDescriptorBufferHandleEXT - Structure specifying push descriptor buffer binding information



## [](#_c_specification)C Specification

When the [`VkPhysicalDeviceDescriptorBufferPropertiesEXT`::`bufferlessPushDescriptors`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-bufferlessPushDescriptors) property is `VK_FALSE`, the `VkBuffer` handle of the buffer for push descriptors is passed in a `VkDescriptorBufferBindingPushDescriptorBufferHandleEXT` structure:

```c++
// Provided by VK_EXT_descriptor_buffer
typedef struct VkDescriptorBufferBindingPushDescriptorBufferHandleEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkBuffer           buffer;
} VkDescriptorBufferBindingPushDescriptorBufferHandleEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `buffer` is the `VkBuffer` handle of the buffer for push descriptors.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDescriptorBufferBindingPushDescriptorBufferHandleEXT-bufferlessPushDescriptors-08059)VUID-VkDescriptorBufferBindingPushDescriptorBufferHandleEXT-bufferlessPushDescriptors-08059  
  [`VkPhysicalDeviceDescriptorBufferPropertiesEXT`::`bufferlessPushDescriptors`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-bufferlessPushDescriptors) **must** be `VK_FALSE`

Valid Usage (Implicit)

- [](#VUID-VkDescriptorBufferBindingPushDescriptorBufferHandleEXT-sType-sType)VUID-VkDescriptorBufferBindingPushDescriptorBufferHandleEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_BUFFER_BINDING_PUSH_DESCRIPTOR_BUFFER_HANDLE_EXT`
- [](#VUID-VkDescriptorBufferBindingPushDescriptorBufferHandleEXT-buffer-parameter)VUID-VkDescriptorBufferBindingPushDescriptorBufferHandleEXT-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorBufferBindingPushDescriptorBufferHandleEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0