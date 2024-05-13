# VkBindDescriptorSetsInfoKHR(3) Manual Page

## Name

VkBindDescriptorSetsInfoKHR - Structure specifying a descriptor set
binding operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBindDescriptorSetsInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_maintenance6
typedef struct VkBindDescriptorSetsInfoKHR {
    VkStructureType           sType;
    const void*               pNext;
    VkShaderStageFlags        stageFlags;
    VkPipelineLayout          layout;
    uint32_t                  firstSet;
    uint32_t                  descriptorSetCount;
    const VkDescriptorSet*    pDescriptorSets;
    uint32_t                  dynamicOffsetCount;
    const uint32_t*           pDynamicOffsets;
} VkBindDescriptorSetsInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stageFlags` is a bitmask of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) specifying the
  shader stages the descriptor sets will be bound to.

- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) object used to
  program the bindings. If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dynamicPipelineLayout"
  target="_blank" rel="noopener"><code>dynamicPipelineLayout</code></a>
  feature is enabled, `layout` **can** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and the layout **must** be
  specified by chaining the
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)
  structure off the `pNext`

- `firstSet` is the set number of the first descriptor set to be bound.

- `descriptorSetCount` is the number of elements in the
  `pDescriptorSets` array.

- `pDescriptorSets` is a pointer to an array of handles to
  [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html) objects describing the
  descriptor sets to bind to.

- `dynamicOffsetCount` is the number of dynamic offsets in the
  `pDynamicOffsets` array.

- `pDynamicOffsets` is a pointer to an array of `uint32_t` values
  specifying dynamic offsets.

## <a href="#_description" class="anchor"></a>Description

If `stageFlags` specifies a subset of all stages corresponding to one or
more pipeline bind points, the binding operation still affects all
stages corresponding to the given pipeline bind point(s) as if the
equivalent original version of this command had been called with the
same parameters. For example, specifying a `stageFlags` value of
`VK_SHADER_STAGE_VERTEX_BIT` \| `VK_SHADER_STAGE_FRAGMENT_BIT` \|
`VK_SHADER_STAGE_COMPUTE_BIT` is equivalent to calling the original
version of this command once with `VK_PIPELINE_BIND_POINT_GRAPHICS` and
once with `VK_PIPELINE_BIND_POINT_COMPUTE`.

Valid Usage

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-00358"
  id="VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-00358"></a>
  VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-00358  
  Each element of `pDescriptorSets` **must** have been allocated with a
  `VkDescriptorSetLayout` that matches (is the same as, or identically
  defined as) the `VkDescriptorSetLayout` at set *n* in `layout`, where
  *n* is the sum of `firstSet` and the index into `pDescriptorSets`

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-dynamicOffsetCount-00359"
  id="VUID-VkBindDescriptorSetsInfoKHR-dynamicOffsetCount-00359"></a>
  VUID-VkBindDescriptorSetsInfoKHR-dynamicOffsetCount-00359  
  `dynamicOffsetCount` **must** be equal to the total number of dynamic
  descriptors in `pDescriptorSets`

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-firstSet-00360"
  id="VUID-VkBindDescriptorSetsInfoKHR-firstSet-00360"></a>
  VUID-VkBindDescriptorSetsInfoKHR-firstSet-00360  
  The sum of `firstSet` and `descriptorSetCount` **must** be less than
  or equal to
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount`
  provided when `layout` was created

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-pDynamicOffsets-01971"
  id="VUID-VkBindDescriptorSetsInfoKHR-pDynamicOffsets-01971"></a>
  VUID-VkBindDescriptorSetsInfoKHR-pDynamicOffsets-01971  
  Each element of `pDynamicOffsets` which corresponds to a descriptor
  binding with type `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` **must**
  be a multiple of
  `VkPhysicalDeviceLimits`::`minUniformBufferOffsetAlignment`

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-pDynamicOffsets-01972"
  id="VUID-VkBindDescriptorSetsInfoKHR-pDynamicOffsets-01972"></a>
  VUID-VkBindDescriptorSetsInfoKHR-pDynamicOffsets-01972  
  Each element of `pDynamicOffsets` which corresponds to a descriptor
  binding with type `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` **must**
  be a multiple of
  `VkPhysicalDeviceLimits`::`minStorageBufferOffsetAlignment`

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-01979"
  id="VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-01979"></a>
  VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-01979  
  For each dynamic uniform or storage buffer binding in
  `pDescriptorSets`, the sum of the [effective
  offset](#dynamic-effective-offset) and the range of the binding
  **must** be less than or equal to the size of the buffer

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-06715"
  id="VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-06715"></a>
  VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-06715  
  For each dynamic uniform or storage buffer binding in
  `pDescriptorSets`, if the range was set with `VK_WHOLE_SIZE` then
  `pDynamicOffsets` which corresponds to the descriptor binding **must**
  be 0

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-04616"
  id="VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-04616"></a>
  VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-04616  
  Each element of `pDescriptorSets` **must** not have been allocated
  from a `VkDescriptorPool` with the
  `VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT` flag set

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-06563"
  id="VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-06563"></a>
  VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-06563  
  If [`graphicsPipelineLibrary`](#features-graphicsPipelineLibrary) is
  not enabled, each element of `pDescriptorSets` **must** be a valid
  [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html)

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-08010"
  id="VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-08010"></a>
  VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-08010  
  Each element of `pDescriptorSets` **must** have been allocated with a
  `VkDescriptorSetLayout` which was not created with
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`

<!-- -->

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-None-09495"
  id="VUID-VkBindDescriptorSetsInfoKHR-None-09495"></a>
  VUID-VkBindDescriptorSetsInfoKHR-None-09495  
  If the [`dynamicPipelineLayout`](#features-dynamicPipelineLayout)
  feature is not enabled, `layout` **must** be a valid
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) handle

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-layout-09496"
  id="VUID-VkBindDescriptorSetsInfoKHR-layout-09496"></a>
  VUID-VkBindDescriptorSetsInfoKHR-layout-09496  
  If `layout` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the `pNext`
  chain **must** include a valid
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)
  structure

Valid Usage (Implicit)

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-sType-sType"
  id="VUID-VkBindDescriptorSetsInfoKHR-sType-sType"></a>
  VUID-VkBindDescriptorSetsInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BIND_DESCRIPTOR_SETS_INFO_KHR`

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-pNext-pNext"
  id="VUID-VkBindDescriptorSetsInfoKHR-pNext-pNext"></a>
  VUID-VkBindDescriptorSetsInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-sType-unique"
  id="VUID-VkBindDescriptorSetsInfoKHR-sType-unique"></a>
  VUID-VkBindDescriptorSetsInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-stageFlags-parameter"
  id="VUID-VkBindDescriptorSetsInfoKHR-stageFlags-parameter"></a>
  VUID-VkBindDescriptorSetsInfoKHR-stageFlags-parameter  
  `stageFlags` **must** be a valid combination of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) values

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-stageFlags-requiredbitmask"
  id="VUID-VkBindDescriptorSetsInfoKHR-stageFlags-requiredbitmask"></a>
  VUID-VkBindDescriptorSetsInfoKHR-stageFlags-requiredbitmask  
  `stageFlags` **must** not be `0`

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-layout-parameter"
  id="VUID-VkBindDescriptorSetsInfoKHR-layout-parameter"></a>
  VUID-VkBindDescriptorSetsInfoKHR-layout-parameter  
  If `layout` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `layout`
  **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) handle

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-parameter"
  id="VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-parameter"></a>
  VUID-VkBindDescriptorSetsInfoKHR-pDescriptorSets-parameter  
  `pDescriptorSets` **must** be a valid pointer to an array of
  `descriptorSetCount` valid [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html)
  handles

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-pDynamicOffsets-parameter"
  id="VUID-VkBindDescriptorSetsInfoKHR-pDynamicOffsets-parameter"></a>
  VUID-VkBindDescriptorSetsInfoKHR-pDynamicOffsets-parameter  
  If `dynamicOffsetCount` is not `0`, and `pDynamicOffsets` is not
  `NULL`, `pDynamicOffsets` **must** be a valid pointer to an array of
  `dynamicOffsetCount` or [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  `uint32_t` values

- <a
  href="#VUID-VkBindDescriptorSetsInfoKHR-descriptorSetCount-arraylength"
  id="VUID-VkBindDescriptorSetsInfoKHR-descriptorSetCount-arraylength"></a>
  VUID-VkBindDescriptorSetsInfoKHR-descriptorSetCount-arraylength  
  `descriptorSetCount` **must** be greater than `0`

- <a href="#VUID-VkBindDescriptorSetsInfoKHR-commonparent"
  id="VUID-VkBindDescriptorSetsInfoKHR-commonparent"></a>
  VUID-VkBindDescriptorSetsInfoKHR-commonparent  
  Both of `layout`, and the elements of `pDescriptorSets` that are valid
  handles of non-ignored parameters **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance6](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance6.html),
[VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html),
[VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdBindDescriptorSets2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorSets2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBindDescriptorSetsInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
