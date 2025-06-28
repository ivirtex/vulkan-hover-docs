# VkDescriptorBufferBindingInfoEXT(3) Manual Page

## Name

VkDescriptorBufferBindingInfoEXT - Structure specifying descriptor buffer binding information



## [](#_c_specification)C Specification

Data describing a descriptor buffer binding is passed in a `VkDescriptorBufferBindingInfoEXT` structure:

```c++
// Provided by VK_EXT_descriptor_buffer
typedef struct VkDescriptorBufferBindingInfoEXT {
    VkStructureType       sType;
    const void*           pNext;
    VkDeviceAddress       address;
    VkBufferUsageFlags    usage;
} VkDescriptorBufferBindingInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `address` is a [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) specifying the device address defining the descriptor buffer to be bound.
- `usage` is a bitmask of [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html) specifying the [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html)::`usage` for the buffer from which `address` was queried.

## [](#_description)Description

If the `pNext` chain includes a [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html) structure, [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html)::`usage` from that structure is used instead of `usage` from this structure.

Valid Usage

- [](#VUID-VkDescriptorBufferBindingInfoEXT-None-09499)VUID-VkDescriptorBufferBindingInfoEXT-None-09499  
  If the `pNext` chain does not include a [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html) structure, `usage` **must** be a valid combination of [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html) values
- [](#VUID-VkDescriptorBufferBindingInfoEXT-None-09500)VUID-VkDescriptorBufferBindingInfoEXT-None-09500  
  If the `pNext` chain does not include a [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html) structure, `usage` **must** not be 0
- [](#VUID-VkDescriptorBufferBindingInfoEXT-bufferlessPushDescriptors-08056)VUID-VkDescriptorBufferBindingInfoEXT-bufferlessPushDescriptors-08056  
  If [`VkPhysicalDeviceDescriptorBufferPropertiesEXT`::`bufferlessPushDescriptors`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-bufferlessPushDescriptors) is `VK_FALSE`, and `usage` contains `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT`, then the `pNext` chain **must** include a [VkDescriptorBufferBindingPushDescriptorBufferHandleEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferBindingPushDescriptorBufferHandleEXT.html) structure
- [](#VUID-VkDescriptorBufferBindingInfoEXT-address-08057)VUID-VkDescriptorBufferBindingInfoEXT-address-08057  
  `address` **must** be aligned to [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`descriptorBufferOffsetAlignment`
- [](#VUID-VkDescriptorBufferBindingInfoEXT-usage-08122)VUID-VkDescriptorBufferBindingInfoEXT-usage-08122  
  If `usage` includes `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT`, `address` **must** be an address within a valid buffer that was created with `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-VkDescriptorBufferBindingInfoEXT-usage-08123)VUID-VkDescriptorBufferBindingInfoEXT-usage-08123  
  If `usage` includes `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`, `address` **must** be an address within a valid buffer that was created with `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-VkDescriptorBufferBindingInfoEXT-usage-08124)VUID-VkDescriptorBufferBindingInfoEXT-usage-08124  
  If `usage` includes `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT`, `address` **must** be an address within a valid buffer that was created with `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT`

Valid Usage (Implicit)

- [](#VUID-VkDescriptorBufferBindingInfoEXT-sType-sType)VUID-VkDescriptorBufferBindingInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_BUFFER_BINDING_INFO_EXT`
- [](#VUID-VkDescriptorBufferBindingInfoEXT-pNext-pNext)VUID-VkDescriptorBufferBindingInfoEXT-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html) or [VkDescriptorBufferBindingPushDescriptorBufferHandleEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferBindingPushDescriptorBufferHandleEXT.html)
- [](#VUID-VkDescriptorBufferBindingInfoEXT-sType-unique)VUID-VkDescriptorBufferBindingInfoEXT-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkBufferUsageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdBindDescriptorBuffersEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorBuffersEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorBufferBindingInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0