# VkPushDescriptorSetInfoKHR(3) Manual Page

## Name

VkPushDescriptorSetInfoKHR - Structure specifying a descriptor set push
operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPushDescriptorSetInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_maintenance6 with VK_KHR_push_descriptor
typedef struct VkPushDescriptorSetInfoKHR {
    VkStructureType                sType;
    const void*                    pNext;
    VkShaderStageFlags             stageFlags;
    VkPipelineLayout               layout;
    uint32_t                       set;
    uint32_t                       descriptorWriteCount;
    const VkWriteDescriptorSet*    pDescriptorWrites;
} VkPushDescriptorSetInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stageFlags` is a bitmask of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) specifying the
  shader stages that will use the descriptors.

- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) object used to
  program the bindings. If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dynamicPipelineLayout"
  target="_blank" rel="noopener"><code>dynamicPipelineLayout</code></a>
  feature is enabled, `layout` **can** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and the layout **must** be
  specified by chaining
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)
  structure off the `pNext`

- `set` is the set number of the descriptor set in the pipeline layout
  that will be updated.

- `descriptorWriteCount` is the number of elements in the
  `pDescriptorWrites` array.

- `pDescriptorWrites` is a pointer to an array of
  [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html) structures
  describing the descriptors to be updated.

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

- <a href="#VUID-VkPushDescriptorSetInfoKHR-set-00364"
  id="VUID-VkPushDescriptorSetInfoKHR-set-00364"></a>
  VUID-VkPushDescriptorSetInfoKHR-set-00364  
  `set` **must** be less than
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount`
  provided when `layout` was created

- <a href="#VUID-VkPushDescriptorSetInfoKHR-set-00365"
  id="VUID-VkPushDescriptorSetInfoKHR-set-00365"></a>
  VUID-VkPushDescriptorSetInfoKHR-set-00365  
  `set` **must** be the unique set number in the pipeline layout that
  uses a descriptor set layout that was created with
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR`

- <a href="#VUID-VkPushDescriptorSetInfoKHR-pDescriptorWrites-06494"
  id="VUID-VkPushDescriptorSetInfoKHR-pDescriptorWrites-06494"></a>
  VUID-VkPushDescriptorSetInfoKHR-pDescriptorWrites-06494  
  For each element i where `pDescriptorWrites`\[i\].`descriptorType` is
  `VK_DESCRIPTOR_TYPE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`,
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`,
  `pDescriptorWrites`\[i\].`pImageInfo` **must** be a valid pointer to
  an array of `pDescriptorWrites`\[i\].`descriptorCount` valid
  `VkDescriptorImageInfo` structures

<!-- -->

- <a href="#VUID-VkPushDescriptorSetInfoKHR-None-09495"
  id="VUID-VkPushDescriptorSetInfoKHR-None-09495"></a>
  VUID-VkPushDescriptorSetInfoKHR-None-09495  
  If the [`dynamicPipelineLayout`](#features-dynamicPipelineLayout)
  feature is not enabled, `layout` **must** be a valid
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) handle

- <a href="#VUID-VkPushDescriptorSetInfoKHR-layout-09496"
  id="VUID-VkPushDescriptorSetInfoKHR-layout-09496"></a>
  VUID-VkPushDescriptorSetInfoKHR-layout-09496  
  If `layout` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the `pNext`
  chain **must** include a valid
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)
  structure

Valid Usage (Implicit)

- <a href="#VUID-VkPushDescriptorSetInfoKHR-sType-sType"
  id="VUID-VkPushDescriptorSetInfoKHR-sType-sType"></a>
  VUID-VkPushDescriptorSetInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PUSH_DESCRIPTOR_SET_INFO_KHR`

- <a href="#VUID-VkPushDescriptorSetInfoKHR-pNext-pNext"
  id="VUID-VkPushDescriptorSetInfoKHR-pNext-pNext"></a>
  VUID-VkPushDescriptorSetInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)

- <a href="#VUID-VkPushDescriptorSetInfoKHR-sType-unique"
  id="VUID-VkPushDescriptorSetInfoKHR-sType-unique"></a>
  VUID-VkPushDescriptorSetInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkPushDescriptorSetInfoKHR-stageFlags-parameter"
  id="VUID-VkPushDescriptorSetInfoKHR-stageFlags-parameter"></a>
  VUID-VkPushDescriptorSetInfoKHR-stageFlags-parameter  
  `stageFlags` **must** be a valid combination of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) values

- <a href="#VUID-VkPushDescriptorSetInfoKHR-stageFlags-requiredbitmask"
  id="VUID-VkPushDescriptorSetInfoKHR-stageFlags-requiredbitmask"></a>
  VUID-VkPushDescriptorSetInfoKHR-stageFlags-requiredbitmask  
  `stageFlags` **must** not be `0`

- <a href="#VUID-VkPushDescriptorSetInfoKHR-layout-parameter"
  id="VUID-VkPushDescriptorSetInfoKHR-layout-parameter"></a>
  VUID-VkPushDescriptorSetInfoKHR-layout-parameter  
  If `layout` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `layout`
  **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) handle

- <a href="#VUID-VkPushDescriptorSetInfoKHR-pDescriptorWrites-parameter"
  id="VUID-VkPushDescriptorSetInfoKHR-pDescriptorWrites-parameter"></a>
  VUID-VkPushDescriptorSetInfoKHR-pDescriptorWrites-parameter  
  `pDescriptorWrites` **must** be a valid pointer to an array of
  `descriptorWriteCount` valid
  [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html) structures

- <a
  href="#VUID-VkPushDescriptorSetInfoKHR-descriptorWriteCount-arraylength"
  id="VUID-VkPushDescriptorSetInfoKHR-descriptorWriteCount-arraylength"></a>
  VUID-VkPushDescriptorSetInfoKHR-descriptorWriteCount-arraylength  
  `descriptorWriteCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance6](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance6.html),
[VK_KHR_push_descriptor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_push_descriptor.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html),
[VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html),
[vkCmdPushDescriptorSet2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSet2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPushDescriptorSetInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
