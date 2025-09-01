# VkPushDescriptorSetInfo(3) Manual Page

## Name

VkPushDescriptorSetInfo - Structure specifying a descriptor set push operation



## [](#_c_specification)C Specification

The `VkPushDescriptorSetInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPushDescriptorSetInfo {
    VkStructureType                sType;
    const void*                    pNext;
    VkShaderStageFlags             stageFlags;
    VkPipelineLayout               layout;
    uint32_t                       set;
    uint32_t                       descriptorWriteCount;
    const VkWriteDescriptorSet*    pDescriptorWrites;
} VkPushDescriptorSetInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance6 with VK_KHR_push_descriptor
typedef VkPushDescriptorSetInfo VkPushDescriptorSetInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stageFlags` is a bitmask of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) specifying the shader stages that will use the descriptors.
- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) object used to program the bindings. If the [`dynamicPipelineLayout`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dynamicPipelineLayout) feature is enabled, `layout` **can** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and the layout **must** be specified by chaining [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html) structure off the `pNext`
- `set` is the set number of the descriptor set in the pipeline layout that will be updated.
- `descriptorWriteCount` is the number of elements in the `pDescriptorWrites` array.
- `pDescriptorWrites` is a pointer to an array of [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html) structures describing the descriptors to be updated.

## [](#_description)Description

If `stageFlags` specifies a subset of all stages corresponding to one or more pipeline bind points, the binding operation still affects all stages corresponding to the given pipeline bind point(s) as if the equivalent original version of this command had been called with the same parameters. For example, specifying a `stageFlags` value of `VK_SHADER_STAGE_VERTEX_BIT` | `VK_SHADER_STAGE_FRAGMENT_BIT` | `VK_SHADER_STAGE_COMPUTE_BIT` is equivalent to calling the original version of this command once with `VK_PIPELINE_BIND_POINT_GRAPHICS` and once with `VK_PIPELINE_BIND_POINT_COMPUTE`.

Valid Usage

- [](#VUID-VkPushDescriptorSetInfo-set-00364)VUID-VkPushDescriptorSetInfo-set-00364  
  `set` **must** be less than [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount` provided when `layout` was created
- [](#VUID-VkPushDescriptorSetInfo-set-00365)VUID-VkPushDescriptorSetInfo-set-00365  
  `set` **must** be the unique set number in the pipeline layout that uses a descriptor set layout that was created with `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT`
- [](#VUID-VkPushDescriptorSetInfo-pDescriptorWrites-06494)VUID-VkPushDescriptorSetInfo-pDescriptorWrites-06494  
  For each element i where `pDescriptorWrites`\[i].`descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLER`, `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, `pDescriptorWrites`\[i].`pImageInfo` **must** be a valid pointer to an array of `pDescriptorWrites`\[i].`descriptorCount` valid `VkDescriptorImageInfo` structures

<!--THE END-->

- [](#VUID-VkPushDescriptorSetInfo-None-09495)VUID-VkPushDescriptorSetInfo-None-09495  
  If the [`dynamicPipelineLayout`](#features-dynamicPipelineLayout) feature is not enabled, `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-VkPushDescriptorSetInfo-layout-09496)VUID-VkPushDescriptorSetInfo-layout-09496  
  If `layout` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the `pNext` chain **must** include a valid [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html) structure

Valid Usage (Implicit)

- [](#VUID-VkPushDescriptorSetInfo-sType-sType)VUID-VkPushDescriptorSetInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PUSH_DESCRIPTOR_SET_INFO`
- [](#VUID-VkPushDescriptorSetInfo-pNext-pNext)VUID-VkPushDescriptorSetInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html)
- [](#VUID-VkPushDescriptorSetInfo-sType-unique)VUID-VkPushDescriptorSetInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkPushDescriptorSetInfo-stageFlags-parameter)VUID-VkPushDescriptorSetInfo-stageFlags-parameter  
  `stageFlags` **must** be a valid combination of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) values
- [](#VUID-VkPushDescriptorSetInfo-stageFlags-requiredbitmask)VUID-VkPushDescriptorSetInfo-stageFlags-requiredbitmask  
  `stageFlags` **must** not be `0`
- [](#VUID-VkPushDescriptorSetInfo-layout-parameter)VUID-VkPushDescriptorSetInfo-layout-parameter  
  If `layout` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-VkPushDescriptorSetInfo-pDescriptorWrites-parameter)VUID-VkPushDescriptorSetInfo-pDescriptorWrites-parameter  
  `pDescriptorWrites` **must** be a valid pointer to an array of `descriptorWriteCount` valid [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html) structures
- [](#VUID-VkPushDescriptorSetInfo-descriptorWriteCount-arraylength)VUID-VkPushDescriptorSetInfo-descriptorWriteCount-arraylength  
  `descriptorWriteCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html), [VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html), [vkCmdPushDescriptorSet2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSet2.html), [vkCmdPushDescriptorSet2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSet2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPushDescriptorSetInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0