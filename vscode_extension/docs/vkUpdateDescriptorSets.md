# vkUpdateDescriptorSets(3) Manual Page

## Name

vkUpdateDescriptorSets - Update the contents of a descriptor set object



## [](#_c_specification)C Specification

Once allocated, descriptor sets **can** be updated with a combination of write and copy operations. To update descriptor sets, call:

```c++
// Provided by VK_VERSION_1_0
void vkUpdateDescriptorSets(
    VkDevice                                    device,
    uint32_t                                    descriptorWriteCount,
    const VkWriteDescriptorSet*                 pDescriptorWrites,
    uint32_t                                    descriptorCopyCount,
    const VkCopyDescriptorSet*                  pDescriptorCopies);
```

## [](#_parameters)Parameters

- `device` is the logical device that updates the descriptor sets.
- `descriptorWriteCount` is the number of elements in the `pDescriptorWrites` array.
- `pDescriptorWrites` is a pointer to an array of [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html) structures describing the descriptor sets to write to.
- `descriptorCopyCount` is the number of elements in the `pDescriptorCopies` array.
- `pDescriptorCopies` is a pointer to an array of [VkCopyDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyDescriptorSet.html) structures describing the descriptor sets to copy between.

## [](#_description)Description

The operations described by `pDescriptorWrites` are performed first, followed by the operations described by `pDescriptorCopies`. Within each array, the operations are performed in the order they appear in the array.

Each element in the `pDescriptorWrites` array describes an operation updating the descriptor set using descriptors for resources specified in the structure.

Each element in the `pDescriptorCopies` array is a [VkCopyDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyDescriptorSet.html) structure describing an operation copying descriptors between sets.

If the `dstSet` member of any element of `pDescriptorWrites` or `pDescriptorCopies` is bound, accessed, or modified by any command that was recorded to a command buffer which is currently in the [recording or executable state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle), and any of the descriptor bindings that are updated were not created with the `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` or `VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT` bits set, that command buffer becomes [invalid](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle).

Copying a descriptor from a descriptor set does not constitute a use of the referenced resource or view, as it is the reference itself that is copied. Applications **can** copy a descriptor referencing a destroyed resource, and it **can** copy an undefined descriptor. The destination descriptor becomes undefined in both cases.

Valid Usage

- [](#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06236)VUID-vkUpdateDescriptorSets-pDescriptorWrites-06236  
  For each element i where `pDescriptorWrites`\[i].`descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` or `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`, elements of the `pTexelBufferView` member of `pDescriptorWrites`\[i] **must** have been created on `device`
- [](#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06237)VUID-vkUpdateDescriptorSets-pDescriptorWrites-06237  
  For each element i where `pDescriptorWrites`\[i].`descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`, `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`, `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`, or `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`, the `buffer` member of any element of the `pBufferInfo` member of `pDescriptorWrites`\[i] **must** have been created on `device`
- [](#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06238)VUID-vkUpdateDescriptorSets-pDescriptorWrites-06238  
  For each element i where `pDescriptorWrites`\[i].`descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLER` or `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, and `dstSet` was not allocated with a layout that included immutable samplers for `dstBinding` with `descriptorType`, the `sampler` member of any element of the `pImageInfo` member of `pDescriptorWrites`\[i] **must** have been created on `device`
- [](#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06239)VUID-vkUpdateDescriptorSets-pDescriptorWrites-06239  
  For each element i where `pDescriptorWrites`\[i].`descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, or `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` the `imageView` member of any element of `pDescriptorWrites`\[i] **must** have been created on `device`
- [](#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06240)VUID-vkUpdateDescriptorSets-pDescriptorWrites-06240  
  For each element i where `pDescriptorWrites`\[i].`descriptorType` is `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR`, elements of the `pAccelerationStructures` member of a [VkWriteDescriptorSetAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetAccelerationStructureKHR.html) structure in the `pNext` chain of `pDescriptorWrites`\[i] **must** have been created on `device`
- [](#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06241)VUID-vkUpdateDescriptorSets-pDescriptorWrites-06241  
  For each element i where `pDescriptorWrites`\[i].`descriptorType` is `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV`, elements of the `pAccelerationStructures` member of a [VkWriteDescriptorSetAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetAccelerationStructureNV.html) structure in the `pNext` chain of `pDescriptorWrites`\[i] **must** have been created on `device`
- [](#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06940)VUID-vkUpdateDescriptorSets-pDescriptorWrites-06940  
  For each element i where `pDescriptorWrites`\[i].`descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM` or `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM`, the `imageView` member of any element of `pDescriptorWrites`\[i] **must** have been created on `device`
- [](#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06493)VUID-vkUpdateDescriptorSets-pDescriptorWrites-06493  
  For each element i where `pDescriptorWrites`\[i].`descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLER`, `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, `pDescriptorWrites`\[i].`pImageInfo` **must** be a valid pointer to an array of `pDescriptorWrites`\[i].`descriptorCount` valid `VkDescriptorImageInfo` structures
- [](#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06941)VUID-vkUpdateDescriptorSets-pDescriptorWrites-06941  
  For each element i where `pDescriptorWrites`\[i].`descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM` or `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM`, `pDescriptorWrites`\[i].`pImageInfo` **must** be a valid pointer to an array of `pDescriptorWrites`\[i].`descriptorCount` valid `VkDescriptorImageInfo` structures
- [](#VUID-vkUpdateDescriptorSets-None-03047)VUID-vkUpdateDescriptorSets-None-03047  
  The `dstSet` member of each element of `pDescriptorWrites` or `pDescriptorCopies` for bindings which were created without the `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` or `VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT` bits set **must** not be used by any command that was recorded to a command buffer which is in the [pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)
- [](#VUID-vkUpdateDescriptorSets-pDescriptorWrites-06993)VUID-vkUpdateDescriptorSets-pDescriptorWrites-06993  
  Host access to `pDescriptorWrites`\[i].`dstSet` and `pDescriptorCopies`\[i].`dstSet` **must** be [externally synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-threadingbehavior) unless explicitly denoted otherwise for specific flags

Valid Usage (Implicit)

- [](#VUID-vkUpdateDescriptorSets-device-parameter)VUID-vkUpdateDescriptorSets-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkUpdateDescriptorSets-pDescriptorWrites-parameter)VUID-vkUpdateDescriptorSets-pDescriptorWrites-parameter  
  If `descriptorWriteCount` is not `0`, `pDescriptorWrites` **must** be a valid pointer to an array of `descriptorWriteCount` valid [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html) structures
- [](#VUID-vkUpdateDescriptorSets-pDescriptorCopies-parameter)VUID-vkUpdateDescriptorSets-pDescriptorCopies-parameter  
  If `descriptorCopyCount` is not `0`, `pDescriptorCopies` **must** be a valid pointer to an array of `descriptorCopyCount` valid [VkCopyDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyDescriptorSet.html) structures

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCopyDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyDescriptorSet.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkUpdateDescriptorSets).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0