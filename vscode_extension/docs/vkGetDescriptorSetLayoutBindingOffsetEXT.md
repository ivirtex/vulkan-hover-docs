# vkGetDescriptorSetLayoutBindingOffsetEXT(3) Manual Page

## Name

vkGetDescriptorSetLayoutBindingOffsetEXT - Get the offset of a binding within a descriptor set layout



## [](#_c_specification)C Specification

To get the offset of a binding within a descriptor set layout in memory, call:

```c++
// Provided by VK_EXT_descriptor_buffer
void vkGetDescriptorSetLayoutBindingOffsetEXT(
    VkDevice                                    device,
    VkDescriptorSetLayout                       layout,
    uint32_t                                    binding,
    VkDeviceSize*                               pOffset);
```

## [](#_parameters)Parameters

- `device` is the logical device that gets the offset.
- `layout` is the descriptor set layout being queried.
- `binding` is the binding number being queried.
- `pOffset` is a pointer to [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) where the byte offset of the binding will be written.

## [](#_description)Description

Each binding in a descriptor set layout is assigned an offset in memory by the implementation. When a shader accesses a resource with that binding, it will access the bound descriptor buffer from that offset to look for its descriptor. This command provides an application with that offset, so that descriptors can be placed in the correct locations. The precise location accessed by a shader for a given descriptor is as follows:

location = bufferAddress + setOffset + descriptorOffset + (arrayElement × descriptorSize)

where bufferAddress and setOffset are the base address and offset for the identified descriptor set as specified by [vkCmdBindDescriptorBuffersEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorBuffersEXT.html) and [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html), descriptorOffset is the offset for the binding returned by this command, arrayElement is the index into the array specified in the shader, and descriptorSize is the size of the relevant descriptor as obtained from [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html). Applications are responsible for placing valid descriptors at the expected location in order for a shader to access it. The overall offset added to bufferAddress to calculate location **must** be less than [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`maxSamplerDescriptorBufferRange` for samplers and [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`maxResourceDescriptorBufferRange` for resources.

If any `binding` in `layout` is `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`, that `binding` **must** have the largest offset of any `binding`.

A descriptor `binding` with type `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` **can** be used. Any potential types in [VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMutableDescriptorTypeCreateInfoEXT.html)::`pDescriptorTypes` for `binding` share the same offset. If the size of the [mutable descriptor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-mutable) is larger than the size of a concrete descriptor type being accessed, the padding area is ignored by the implementation.

Valid Usage

- [](#VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-None-08013)VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-None-08013  
  The [`descriptorBuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBuffer) feature **must** be enabled
- [](#VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-layout-08014)VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-layout-08014  
  `layout` **must** have been created with the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT` flag set

Valid Usage (Implicit)

- [](#VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-device-parameter)VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-layout-parameter)VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-layout-parameter  
  `layout` **must** be a valid [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) handle
- [](#VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-pOffset-parameter)VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-pOffset-parameter  
  `pOffset` **must** be a valid pointer to a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) value
- [](#VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-layout-parent)VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-layout-parent  
  `layout` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDescriptorSetLayoutBindingOffsetEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0