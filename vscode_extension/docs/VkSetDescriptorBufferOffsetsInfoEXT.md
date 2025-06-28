# VkSetDescriptorBufferOffsetsInfoEXT(3) Manual Page

## Name

VkSetDescriptorBufferOffsetsInfoEXT - Structure specifying descriptor buffer offsets to set in a command buffer



## [](#_c_specification)C Specification

The `VkSetDescriptorBufferOffsetsInfoEXT` structure is defined as:

```c++
// Provided by VK_KHR_maintenance6 with VK_EXT_descriptor_buffer
typedef struct VkSetDescriptorBufferOffsetsInfoEXT {
    VkStructureType        sType;
    const void*            pNext;
    VkShaderStageFlags     stageFlags;
    VkPipelineLayout       layout;
    uint32_t               firstSet;
    uint32_t               setCount;
    const uint32_t*        pBufferIndices;
    const VkDeviceSize*    pOffsets;
} VkSetDescriptorBufferOffsetsInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stageFlags` is a bitmask of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) specifying the shader stages the descriptor sets will be bound to
- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) object used to program the bindings. If the [`dynamicPipelineLayout`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dynamicPipelineLayout) feature is enabled, `layout` **can** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and the layout **must** be specified by chaining [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html) structure off the `pNext`
- `firstSet` is the number of the first set to be bound.
- `setCount` is the number of elements in the `pBufferIndices` and `pOffsets` arrays.
- `pBufferIndices` is a pointer to an array of indices into the descriptor buffer binding points set by [vkCmdBindDescriptorBuffersEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorBuffersEXT.html).
- `pOffsets` is a pointer to an array of [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) offsets to apply to the bound descriptor buffers.

## [](#_description)Description

If `stageFlags` specifies a subset of all stages corresponding to one or more pipeline bind points, the binding operation still affects all stages corresponding to the given pipeline bind point(s) as if the equivalent original version of this command had been called with the same parameters. For example, specifying a `stageFlags` value of `VK_SHADER_STAGE_VERTEX_BIT` | `VK_SHADER_STAGE_FRAGMENT_BIT` | `VK_SHADER_STAGE_COMPUTE_BIT` is equivalent to calling the original version of this command once with `VK_PIPELINE_BIND_POINT_GRAPHICS` and once with `VK_PIPELINE_BIND_POINT_COMPUTE`.

Valid Usage

- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-pOffsets-08061)VUID-VkSetDescriptorBufferOffsetsInfoEXT-pOffsets-08061  
  The offsets in `pOffsets` **must** be aligned to [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`descriptorBufferOffsetAlignment`
- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-pOffsets-08063)VUID-VkSetDescriptorBufferOffsetsInfoEXT-pOffsets-08063  
  The offsets in `pOffsets` **must** be small enough such that any descriptor binding referenced by `layout` without the `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT` flag computes a valid address inside the underlying [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html)
- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-pOffsets-08126)VUID-VkSetDescriptorBufferOffsetsInfoEXT-pOffsets-08126  
  The offsets in `pOffsets` **must** be small enough such that any location accessed by a shader as a sampler descriptor **must** be within [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`maxSamplerDescriptorBufferRange` of the sampler descriptor buffer binding
- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-pOffsets-08127)VUID-VkSetDescriptorBufferOffsetsInfoEXT-pOffsets-08127  
  The offsets in `pOffsets` **must** be small enough such that any location accessed by a shader as a resource descriptor **must** be within [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`maxResourceDescriptorBufferRange` of the resource descriptor buffer binding
- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-pBufferIndices-08064)VUID-VkSetDescriptorBufferOffsetsInfoEXT-pBufferIndices-08064  
  Each element of `pBufferIndices` **must** be less than [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`maxDescriptorBufferBindings`
- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-pBufferIndices-08065)VUID-VkSetDescriptorBufferOffsetsInfoEXT-pBufferIndices-08065  
  Each element of `pBufferIndices` **must** reference a valid descriptor buffer binding set by a previous call to [vkCmdBindDescriptorBuffersEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorBuffersEXT.html) in `commandBuffer`
- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-firstSet-08066)VUID-VkSetDescriptorBufferOffsetsInfoEXT-firstSet-08066  
  The sum of `firstSet` and `setCount` **must** be less than or equal to [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount` provided when `layout` was created
- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-firstSet-09006)VUID-VkSetDescriptorBufferOffsetsInfoEXT-firstSet-09006  
  The [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) for each set from `firstSet` to `firstSet` + `setCount` when `layout` was created **must** have been created with the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT` bit set

<!--THE END-->

- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-None-09495)VUID-VkSetDescriptorBufferOffsetsInfoEXT-None-09495  
  If the [`dynamicPipelineLayout`](#features-dynamicPipelineLayout) feature is not enabled, `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-layout-09496)VUID-VkSetDescriptorBufferOffsetsInfoEXT-layout-09496  
  If `layout` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the `pNext` chain **must** include a valid [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html) structure

Valid Usage (Implicit)

- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-sType-sType)VUID-VkSetDescriptorBufferOffsetsInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SET_DESCRIPTOR_BUFFER_OFFSETS_INFO_EXT`
- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-pNext-pNext)VUID-VkSetDescriptorBufferOffsetsInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html)
- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-sType-unique)VUID-VkSetDescriptorBufferOffsetsInfoEXT-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-stageFlags-parameter)VUID-VkSetDescriptorBufferOffsetsInfoEXT-stageFlags-parameter  
  `stageFlags` **must** be a valid combination of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) values
- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-stageFlags-requiredbitmask)VUID-VkSetDescriptorBufferOffsetsInfoEXT-stageFlags-requiredbitmask  
  `stageFlags` **must** not be `0`
- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-layout-parameter)VUID-VkSetDescriptorBufferOffsetsInfoEXT-layout-parameter  
  If `layout` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-pBufferIndices-parameter)VUID-VkSetDescriptorBufferOffsetsInfoEXT-pBufferIndices-parameter  
  `pBufferIndices` **must** be a valid pointer to an array of `setCount` `uint32_t` values
- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-pOffsets-parameter)VUID-VkSetDescriptorBufferOffsetsInfoEXT-pOffsets-parameter  
  `pOffsets` **must** be a valid pointer to an array of `setCount` [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) values
- [](#VUID-VkSetDescriptorBufferOffsetsInfoEXT-setCount-arraylength)VUID-VkSetDescriptorBufferOffsetsInfoEXT-setCount-arraylength  
  `setCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdSetDescriptorBufferOffsets2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsets2EXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSetDescriptorBufferOffsetsInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0