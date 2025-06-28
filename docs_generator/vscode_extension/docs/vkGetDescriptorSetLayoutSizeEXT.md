# vkGetDescriptorSetLayoutSizeEXT(3) Manual Page

## Name

vkGetDescriptorSetLayoutSizeEXT - Get the size of a descriptor set layout in memory



## [](#_c_specification)C Specification

To determine the amount of memory needed to store all descriptors with a given layout, call:

```c++
// Provided by VK_EXT_descriptor_buffer
void vkGetDescriptorSetLayoutSizeEXT(
    VkDevice                                    device,
    VkDescriptorSetLayout                       layout,
    VkDeviceSize*                               pLayoutSizeInBytes);
```

## [](#_parameters)Parameters

- `device` is the logical device that gets the size.
- `layout` is the descriptor set layout being queried.
- `pLayoutSizeInBytes` is a pointer to [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) where the size in bytes will be written.

## [](#_description)Description

The size of a descriptor set layout will be at least as large as the sum total of the size of all descriptors in the layout, and **may** be larger. This size represents the amount of memory that will be required to store all of the descriptors for this layout in memory, when placed according to the layout’s offsets as obtained by [vkGetDescriptorSetLayoutBindingOffsetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorSetLayoutBindingOffsetEXT.html).

If any `binding` in `layout` is `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`, the returned size includes space for the maximum `descriptorCount` descriptors as declared for that `binding`. To compute the required size of a descriptor set with a `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`:

size = offset + descriptorSize × variableDescriptorCount

where offset is obtained by [vkGetDescriptorSetLayoutBindingOffsetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorSetLayoutBindingOffsetEXT.html) and descriptorSize is the size of the relevant descriptor as obtained from [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html), and variableDescriptorCount is the equivalent of [VkDescriptorSetVariableDescriptorCountAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetVariableDescriptorCountAllocateInfo.html)::`pDescriptorCounts`. For `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, variableDescriptorCount is the size in bytes for the inline uniform block, and descriptorSize is 1.

If [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`combinedImageSamplerDescriptorSingleArray` is `VK_FALSE` and the variable descriptor type is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, variableDescriptorCount is always considered to be the upper bound.

Valid Usage

- [](#VUID-vkGetDescriptorSetLayoutSizeEXT-None-08011)VUID-vkGetDescriptorSetLayoutSizeEXT-None-08011  
  The [`descriptorBuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBuffer) feature **must** be enabled
- [](#VUID-vkGetDescriptorSetLayoutSizeEXT-layout-08012)VUID-vkGetDescriptorSetLayoutSizeEXT-layout-08012  
  `layout` **must** have been created with the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT` flag set

Valid Usage (Implicit)

- [](#VUID-vkGetDescriptorSetLayoutSizeEXT-device-parameter)VUID-vkGetDescriptorSetLayoutSizeEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDescriptorSetLayoutSizeEXT-layout-parameter)VUID-vkGetDescriptorSetLayoutSizeEXT-layout-parameter  
  `layout` **must** be a valid [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) handle
- [](#VUID-vkGetDescriptorSetLayoutSizeEXT-pLayoutSizeInBytes-parameter)VUID-vkGetDescriptorSetLayoutSizeEXT-pLayoutSizeInBytes-parameter  
  `pLayoutSizeInBytes` **must** be a valid pointer to a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) value
- [](#VUID-vkGetDescriptorSetLayoutSizeEXT-layout-parent)VUID-vkGetDescriptorSetLayoutSizeEXT-layout-parent  
  `layout` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDescriptorSetLayoutSizeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0