# VkBindDescriptorSetsInfo(3) Manual Page

## Name

VkBindDescriptorSetsInfo - Structure specifying a descriptor set binding operation



## [](#_c_specification)C Specification

The `VkBindDescriptorSetsInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkBindDescriptorSetsInfo {
    VkStructureType           sType;
    const void*               pNext;
    VkShaderStageFlags        stageFlags;
    VkPipelineLayout          layout;
    uint32_t                  firstSet;
    uint32_t                  descriptorSetCount;
    const VkDescriptorSet*    pDescriptorSets;
    uint32_t                  dynamicOffsetCount;
    const uint32_t*           pDynamicOffsets;
} VkBindDescriptorSetsInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance6
typedef VkBindDescriptorSetsInfo VkBindDescriptorSetsInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stageFlags` is a bitmask of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) specifying the shader stages the descriptor sets will be bound to.
- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) object used to program the bindings. If the [`dynamicPipelineLayout`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dynamicPipelineLayout) feature is enabled, `layout` **can** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and the layout **must** be specified by chaining the [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html) structure off the `pNext`
- `firstSet` is the set number of the first descriptor set to be bound.
- `descriptorSetCount` is the number of elements in the `pDescriptorSets` array.
- `pDescriptorSets` is a pointer to an array of handles to [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html) objects describing the descriptor sets to bind to.
- `dynamicOffsetCount` is the number of dynamic offsets in the `pDynamicOffsets` array.
- `pDynamicOffsets` is a pointer to an array of `uint32_t` values specifying dynamic offsets.

## [](#_description)Description

If `stageFlags` specifies a subset of all stages corresponding to one or more pipeline bind points, the binding operation still affects all stages corresponding to the given pipeline bind point(s) as if the equivalent original version of this command had been called with the same parameters. For example, specifying a `stageFlags` value of `VK_SHADER_STAGE_VERTEX_BIT` | `VK_SHADER_STAGE_FRAGMENT_BIT` | `VK_SHADER_STAGE_COMPUTE_BIT` is equivalent to calling the original version of this command once with `VK_PIPELINE_BIND_POINT_GRAPHICS` and once with `VK_PIPELINE_BIND_POINT_COMPUTE`.

Valid Usage

- [](#VUID-VkBindDescriptorSetsInfo-pDescriptorSets-00358)VUID-VkBindDescriptorSetsInfo-pDescriptorSets-00358  
  Each element of `pDescriptorSets` that is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) **must** have been allocated with a `VkDescriptorSetLayout` that matches (is the same as, or identically defined as) the `VkDescriptorSetLayout` at set *n* in `layout`, where *n* is the sum of `firstSet` and the index into `pDescriptorSets`
- [](#VUID-VkBindDescriptorSetsInfo-dynamicOffsetCount-00359)VUID-VkBindDescriptorSetsInfo-dynamicOffsetCount-00359  
  `dynamicOffsetCount` **must** be equal to the total number of dynamic descriptors in `pDescriptorSets`
- [](#VUID-VkBindDescriptorSetsInfo-firstSet-00360)VUID-VkBindDescriptorSetsInfo-firstSet-00360  
  The sum of `firstSet` and `descriptorSetCount` **must** be less than or equal to [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount` provided when `layout` was created
- [](#VUID-VkBindDescriptorSetsInfo-pDynamicOffsets-01971)VUID-VkBindDescriptorSetsInfo-pDynamicOffsets-01971  
  Each element of `pDynamicOffsets` which corresponds to a descriptor binding with type `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` **must** be a multiple of `VkPhysicalDeviceLimits`::`minUniformBufferOffsetAlignment`
- [](#VUID-VkBindDescriptorSetsInfo-pDynamicOffsets-01972)VUID-VkBindDescriptorSetsInfo-pDynamicOffsets-01972  
  Each element of `pDynamicOffsets` which corresponds to a descriptor binding with type `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` **must** be a multiple of `VkPhysicalDeviceLimits`::`minStorageBufferOffsetAlignment`
- [](#VUID-VkBindDescriptorSetsInfo-pDescriptorSets-01979)VUID-VkBindDescriptorSetsInfo-pDescriptorSets-01979  
  For each dynamic uniform or storage buffer binding in `pDescriptorSets`, the sum of the [effective offset](#dynamic-effective-offset) and the range of the binding **must** be less than or equal to the size of the buffer
- [](#VUID-VkBindDescriptorSetsInfo-pDescriptorSets-06715)VUID-VkBindDescriptorSetsInfo-pDescriptorSets-06715  
  For each dynamic uniform or storage buffer binding in `pDescriptorSets`, if the range was set with `VK_WHOLE_SIZE` then `pDynamicOffsets` which corresponds to the descriptor binding **must** be 0
- [](#VUID-VkBindDescriptorSetsInfo-pDescriptorSets-04616)VUID-VkBindDescriptorSetsInfo-pDescriptorSets-04616  
  Each element of `pDescriptorSets` **must** not have been allocated from a `VkDescriptorPool` with the `VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT` flag set
- [](#VUID-VkBindDescriptorSetsInfo-pDescriptorSets-06563)VUID-VkBindDescriptorSetsInfo-pDescriptorSets-06563  
  If the [`graphicsPipelineLibrary`](#features-graphicsPipelineLibrary) feature is not enabled, each element of `pDescriptorSets` **must** be a valid [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html)
- [](#VUID-VkBindDescriptorSetsInfo-pDescriptorSets-08010)VUID-VkBindDescriptorSetsInfo-pDescriptorSets-08010  
  Each element of `pDescriptorSets` **must** have been allocated with a `VkDescriptorSetLayout` which was not created with `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-VkBindDescriptorSetsInfo-pDescriptorSets-09914)VUID-VkBindDescriptorSetsInfo-pDescriptorSets-09914  
  If any element of `pDescriptorSets` was allocated from a descriptor pool created with a [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html) structure that had a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure specifying foreign data processing engines in its `pNext` chain, then the command pool from which `commandBuffer` was allocated **must** have been created with a [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) structure that had a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure in its `pNext` chain specifying a superset of all the foreign data processing engines specified when creating the descriptor pools from which the elements of `pDescriptorSets` were allocated
- [](#VUID-VkBindDescriptorSetsInfo-pDescriptorSets-09915)VUID-VkBindDescriptorSetsInfo-pDescriptorSets-09915  
  If none of the elements of `pDescriptorSets` were allocated from a descriptor pool created with a [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html) structure that had a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure specifying foreign data processing engines in its `pNext` chain, then the command pool from which `commandBuffer` was allocated **must** not have been created with a [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) structure that had a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure in its `pNext` chain

<!--THE END-->

- [](#VUID-VkBindDescriptorSetsInfo-None-09495)VUID-VkBindDescriptorSetsInfo-None-09495  
  If the [`dynamicPipelineLayout`](#features-dynamicPipelineLayout) feature is not enabled, `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-VkBindDescriptorSetsInfo-layout-09496)VUID-VkBindDescriptorSetsInfo-layout-09496  
  If `layout` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the `pNext` chain **must** include a valid [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html) structure

Valid Usage (Implicit)

- [](#VUID-VkBindDescriptorSetsInfo-sType-sType)VUID-VkBindDescriptorSetsInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BIND_DESCRIPTOR_SETS_INFO`
- [](#VUID-VkBindDescriptorSetsInfo-pNext-pNext)VUID-VkBindDescriptorSetsInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html)
- [](#VUID-VkBindDescriptorSetsInfo-sType-unique)VUID-VkBindDescriptorSetsInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkBindDescriptorSetsInfo-stageFlags-parameter)VUID-VkBindDescriptorSetsInfo-stageFlags-parameter  
  `stageFlags` **must** be a valid combination of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) values
- [](#VUID-VkBindDescriptorSetsInfo-stageFlags-requiredbitmask)VUID-VkBindDescriptorSetsInfo-stageFlags-requiredbitmask  
  `stageFlags` **must** not be `0`
- [](#VUID-VkBindDescriptorSetsInfo-layout-parameter)VUID-VkBindDescriptorSetsInfo-layout-parameter  
  If `layout` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-VkBindDescriptorSetsInfo-pDescriptorSets-parameter)VUID-VkBindDescriptorSetsInfo-pDescriptorSets-parameter  
  `pDescriptorSets` **must** be a valid pointer to an array of `descriptorSetCount` valid [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html) handles
- [](#VUID-VkBindDescriptorSetsInfo-pDynamicOffsets-parameter)VUID-VkBindDescriptorSetsInfo-pDynamicOffsets-parameter  
  If `dynamicOffsetCount` is not `0`, and `pDynamicOffsets` is not `NULL`, `pDynamicOffsets` **must** be a valid pointer to an array of `dynamicOffsetCount` or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) `uint32_t` values
- [](#VUID-VkBindDescriptorSetsInfo-descriptorSetCount-arraylength)VUID-VkBindDescriptorSetsInfo-descriptorSetCount-arraylength  
  `descriptorSetCount` **must** be greater than `0`
- [](#VUID-VkBindDescriptorSetsInfo-commonparent)VUID-VkBindDescriptorSetsInfo-commonparent  
  Both of `layout`, and the elements of `pDescriptorSets` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdBindDescriptorSets2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorSets2.html), [vkCmdBindDescriptorSets2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorSets2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBindDescriptorSetsInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0