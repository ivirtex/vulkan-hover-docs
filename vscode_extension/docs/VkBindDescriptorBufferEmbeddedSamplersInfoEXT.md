# VkBindDescriptorBufferEmbeddedSamplersInfoEXT(3) Manual Page

## Name

VkBindDescriptorBufferEmbeddedSamplersInfoEXT - Structure specifying embedded immutable sampler offsets to set in a command buffer



## [](#_c_specification)C Specification

The `VkBindDescriptorBufferEmbeddedSamplersInfoEXT` structure is defined as:

```c++
// Provided by VK_KHR_maintenance6 with VK_EXT_descriptor_buffer
typedef struct VkBindDescriptorBufferEmbeddedSamplersInfoEXT {
    VkStructureType       sType;
    const void*           pNext;
    VkShaderStageFlags    stageFlags;
    VkPipelineLayout      layout;
    uint32_t              set;
} VkBindDescriptorBufferEmbeddedSamplersInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stageFlags` is a bitmask of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) specifying the shader stages that will use the embedded immutable samplers.
- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) object used to program the bindings. If the [`dynamicPipelineLayout`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dynamicPipelineLayout) feature is enabled, `layout` **can** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and the layout **must** be specified by chaining [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html) structure off the `pNext`
- `set` is the number of the set to be bound.

## [](#_description)Description

If `stageFlags` specifies a subset of all stages corresponding to one or more pipeline bind points, the binding operation still affects all stages corresponding to the given pipeline bind point(s) as if the equivalent original version of this command had been called with the same parameters. For example, specifying a `stageFlags` value of `VK_SHADER_STAGE_VERTEX_BIT` | `VK_SHADER_STAGE_FRAGMENT_BIT` | `VK_SHADER_STAGE_COMPUTE_BIT` is equivalent to calling the original version of this command once with `VK_PIPELINE_BIND_POINT_GRAPHICS` and once with `VK_PIPELINE_BIND_POINT_COMPUTE`.

Valid Usage

- [](#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-set-08070)VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-set-08070  
  The [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) at index `set` when `layout` was created **must** have been created with the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_EMBEDDED_IMMUTABLE_SAMPLERS_BIT_EXT` bit set
- [](#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-set-08071)VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-set-08071  
  `set` **must** be less than or equal to [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount` provided when `layout` was created

<!--THE END-->

- [](#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-None-09495)VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-None-09495  
  If the [`dynamicPipelineLayout`](#features-dynamicPipelineLayout) feature is not enabled, `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-layout-09496)VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-layout-09496  
  If `layout` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the `pNext` chain **must** include a valid [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html) structure

Valid Usage (Implicit)

- [](#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-sType-sType)VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BIND_DESCRIPTOR_BUFFER_EMBEDDED_SAMPLERS_INFO_EXT`
- [](#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-pNext-pNext)VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html)
- [](#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-sType-unique)VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-stageFlags-parameter)VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-stageFlags-parameter  
  `stageFlags` **must** be a valid combination of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) values
- [](#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-stageFlags-requiredbitmask)VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-stageFlags-requiredbitmask  
  `stageFlags` **must** not be `0`
- [](#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-layout-parameter)VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-layout-parameter  
  If `layout` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdBindDescriptorBufferEmbeddedSamplers2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorBufferEmbeddedSamplers2EXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBindDescriptorBufferEmbeddedSamplersInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0