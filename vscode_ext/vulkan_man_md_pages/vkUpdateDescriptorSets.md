# vkUpdateDescriptorSets(3) Manual Page

## Name

vkUpdateDescriptorSets - Update the contents of a descriptor set object



## <a href="#_c_specification" class="anchor"></a>C Specification

Once allocated, descriptor sets **can** be updated with a combination of
write and copy operations. To update descriptor sets, call:

``` c
// Provided by VK_VERSION_1_0
void vkUpdateDescriptorSets(
    VkDevice                                    device,
    uint32_t                                    descriptorWriteCount,
    const VkWriteDescriptorSet*                 pDescriptorWrites,
    uint32_t                                    descriptorCopyCount,
    const VkCopyDescriptorSet*                  pDescriptorCopies);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that updates the descriptor sets.

- `descriptorWriteCount` is the number of elements in the
  `pDescriptorWrites` array.

- `pDescriptorWrites` is a pointer to an array of
  [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html) structures
  describing the descriptor sets to write to.

- `descriptorCopyCount` is the number of elements in the
  `pDescriptorCopies` array.

- `pDescriptorCopies` is a pointer to an array of
  [VkCopyDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyDescriptorSet.html) structures describing
  the descriptor sets to copy between.

## <a href="#_description" class="anchor"></a>Description

The operations described by `pDescriptorWrites` are performed first,
followed by the operations described by `pDescriptorCopies`. Within each
array, the operations are performed in the order they appear in the
array.

Each element in the `pDescriptorWrites` array describes an operation
updating the descriptor set using descriptors for resources specified in
the structure.

Each element in the `pDescriptorCopies` array is a
[VkCopyDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyDescriptorSet.html) structure describing an
operation copying descriptors between sets.

If the `dstSet` member of any element of `pDescriptorWrites` or
`pDescriptorCopies` is bound, accessed, or modified by any command that
was recorded to a command buffer which is currently in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">recording or executable state</a>, and
any of the descriptor bindings that are updated were not created with
the `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` or
`VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT` bits set, that
command buffer becomes <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">invalid</a>.

Valid Usage

- <a href="#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06236"
  id="VUID-vkUpdateDescriptorSets-pDescriptorWrites-06236"></a>
  VUID-vkUpdateDescriptorSets-pDescriptorWrites-06236  
  For each element i where `pDescriptorWrites`\[i\].`descriptorType` is
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` or
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`, elements of the
  `pTexelBufferView` member of `pDescriptorWrites`\[i\] **must** have
  been created on `device`

- <a href="#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06237"
  id="VUID-vkUpdateDescriptorSets-pDescriptorWrites-06237"></a>
  VUID-vkUpdateDescriptorSets-pDescriptorWrites-06237  
  For each element i where `pDescriptorWrites`\[i\].`descriptorType` is
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`,
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`,
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`, or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`, the `buffer` member of
  any element of the `pBufferInfo` member of `pDescriptorWrites`\[i\]
  **must** have been created on `device`

- <a href="#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06238"
  id="VUID-vkUpdateDescriptorSets-pDescriptorWrites-06238"></a>
  VUID-vkUpdateDescriptorSets-pDescriptorWrites-06238  
  For each element i where `pDescriptorWrites`\[i\].`descriptorType` is
  `VK_DESCRIPTOR_TYPE_SAMPLER` or
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, and `dstSet` was not
  allocated with a layout that included immutable samplers for
  `dstBinding` with `descriptorType`, the `sampler` member of any
  element of the `pImageInfo` member of `pDescriptorWrites`\[i\]
  **must** have been created on `device`

- <a href="#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06239"
  id="VUID-vkUpdateDescriptorSets-pDescriptorWrites-06239"></a>
  VUID-vkUpdateDescriptorSets-pDescriptorWrites-06239  
  For each element i where `pDescriptorWrites`\[i\].`descriptorType` is
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`,
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`,
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, or
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` the `imageView` member of
  any element of `pDescriptorWrites`\[i\] **must** have been created on
  `device`

- <a href="#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06240"
  id="VUID-vkUpdateDescriptorSets-pDescriptorWrites-06240"></a>
  VUID-vkUpdateDescriptorSets-pDescriptorWrites-06240  
  For each element i where `pDescriptorWrites`\[i\].`descriptorType` is
  `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR`, elements of the
  `pAccelerationStructures` member of a
  [VkWriteDescriptorSetAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetAccelerationStructureKHR.html)
  structure in the `pNext` chain of `pDescriptorWrites`\[i\] **must**
  have been created on `device`

- <a href="#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06241"
  id="VUID-vkUpdateDescriptorSets-pDescriptorWrites-06241"></a>
  VUID-vkUpdateDescriptorSets-pDescriptorWrites-06241  
  For each element i where `pDescriptorWrites`\[i\].`descriptorType` is
  `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV`, elements of the
  `pAccelerationStructures` member of a
  [VkWriteDescriptorSetAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetAccelerationStructureNV.html)
  structure in the `pNext` chain of `pDescriptorWrites`\[i\] **must**
  have been created on `device`

- <a href="#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06940"
  id="VUID-vkUpdateDescriptorSets-pDescriptorWrites-06940"></a>
  VUID-vkUpdateDescriptorSets-pDescriptorWrites-06940  
  For each element i where `pDescriptorWrites`\[i\].`descriptorType` is
  `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM` or
  `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM`, the `imageView` member of
  any element of `pDescriptorWrites`\[i\] **must** have been created on
  `device`

- <a href="#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06493"
  id="VUID-vkUpdateDescriptorSets-pDescriptorWrites-06493"></a>
  VUID-vkUpdateDescriptorSets-pDescriptorWrites-06493  
  For each element i where `pDescriptorWrites`\[i\].`descriptorType` is
  `VK_DESCRIPTOR_TYPE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`,
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`,
  `pDescriptorWrites`\[i\].`pImageInfo` **must** be a valid pointer to
  an array of `pDescriptorWrites`\[i\].`descriptorCount` valid
  `VkDescriptorImageInfo` structures

- <a href="#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06941"
  id="VUID-vkUpdateDescriptorSets-pDescriptorWrites-06941"></a>
  VUID-vkUpdateDescriptorSets-pDescriptorWrites-06941  
  For each element i where `pDescriptorWrites`\[i\].`descriptorType` is
  `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM` or
  `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM`,
  `pDescriptorWrites`\[i\].`pImageInfo` **must** be a valid pointer to
  an array of `pDescriptorWrites`\[i\].`descriptorCount` valid
  `VkDescriptorImageInfo` structures

- <a href="#VUID-vkUpdateDescriptorSets-None-03047"
  id="VUID-vkUpdateDescriptorSets-None-03047"></a>
  VUID-vkUpdateDescriptorSets-None-03047  
  The `dstSet` member of each element of `pDescriptorWrites` or
  `pDescriptorCopies` for bindings which were created without the
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` or
  `VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT` bits set
  **must** not be used by any command that was recorded to a command
  buffer which is in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending state</a>

- <a href="#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06993"
  id="VUID-vkUpdateDescriptorSets-pDescriptorWrites-06993"></a>
  VUID-vkUpdateDescriptorSets-pDescriptorWrites-06993  
  Host access to `pDescriptorWrites`\[i\].`dstSet` and
  `pDescriptorCopies`\[i\].`dstSet` **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fundamentals-threadingbehavior"
  target="_blank" rel="noopener">externally synchronized</a> unless
  explicitly denoted otherwise for specific flags

Valid Usage (Implicit)

- <a href="#VUID-vkUpdateDescriptorSets-device-parameter"
  id="VUID-vkUpdateDescriptorSets-device-parameter"></a>
  VUID-vkUpdateDescriptorSets-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkUpdateDescriptorSets-pDescriptorWrites-parameter"
  id="VUID-vkUpdateDescriptorSets-pDescriptorWrites-parameter"></a>
  VUID-vkUpdateDescriptorSets-pDescriptorWrites-parameter  
  If `descriptorWriteCount` is not `0`, `pDescriptorWrites` **must** be
  a valid pointer to an array of `descriptorWriteCount` valid
  [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html) structures

- <a href="#VUID-vkUpdateDescriptorSets-pDescriptorCopies-parameter"
  id="VUID-vkUpdateDescriptorSets-pDescriptorCopies-parameter"></a>
  VUID-vkUpdateDescriptorSets-pDescriptorCopies-parameter  
  If `descriptorCopyCount` is not `0`, `pDescriptorCopies` **must** be a
  valid pointer to an array of `descriptorCopyCount` valid
  [VkCopyDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyDescriptorSet.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCopyDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyDescriptorSet.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkUpdateDescriptorSets"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
