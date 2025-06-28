# VkDescriptorSetAllocateInfo(3) Manual Page

## Name

VkDescriptorSetAllocateInfo - Structure specifying the allocation parameters for descriptor sets



## [](#_c_specification)C Specification

The `VkDescriptorSetAllocateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkDescriptorSetAllocateInfo {
    VkStructureType                 sType;
    const void*                     pNext;
    VkDescriptorPool                descriptorPool;
    uint32_t                        descriptorSetCount;
    const VkDescriptorSetLayout*    pSetLayouts;
} VkDescriptorSetAllocateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `descriptorPool` is the pool which the sets will be allocated from.
- `descriptorSetCount` determines the number of descriptor sets to be allocated from the pool.
- `pSetLayouts` is a pointer to an array of descriptor set layouts, with each member specifying how the corresponding descriptor set is allocated.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDescriptorSetAllocateInfo-apiVersion-07895)VUID-VkDescriptorSetAllocateInfo-apiVersion-07895  
  If the [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html) extension is not enabled and [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is less than Vulkan 1.1, `descriptorSetCount` **must** not be greater than the number of sets that are currently available for allocation in `descriptorPool`
- [](#VUID-VkDescriptorSetAllocateInfo-apiVersion-07896)VUID-VkDescriptorSetAllocateInfo-apiVersion-07896  
  If the [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html) extension is not enabled and [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is less than Vulkan 1.1, `descriptorPool` **must** have enough free descriptor capacity remaining to allocate the descriptor sets of the specified layouts
- [](#VUID-VkDescriptorSetAllocateInfo-pSetLayouts-00308)VUID-VkDescriptorSetAllocateInfo-pSetLayouts-00308  
  Each element of `pSetLayouts` **must** not have been created with `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT` set
- [](#VUID-VkDescriptorSetAllocateInfo-pSetLayouts-03044)VUID-VkDescriptorSetAllocateInfo-pSetLayouts-03044  
  If any element of `pSetLayouts` was created with the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set, `descriptorPool` **must** have been created with the `VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT` flag set
- [](#VUID-VkDescriptorSetAllocateInfo-pSetLayouts-09380)VUID-VkDescriptorSetAllocateInfo-pSetLayouts-09380  
  If `pSetLayouts`\[i] was created with an element of `pBindingFlags` that includes `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`, and [VkDescriptorSetVariableDescriptorCountAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetVariableDescriptorCountAllocateInfo.html) is included in the `pNext` chain, and `VkDescriptorSetVariableDescriptorCountAllocateInfo`::`descriptorSetCount` is not zero, then [VkDescriptorSetVariableDescriptorCountAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetVariableDescriptorCountAllocateInfo.html)::`pDescriptorCounts`\[i] **must** be less than or equal to [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBinding.html)::`descriptorCount` for the corresponding binding used to create `pSetLayouts`\[i]
- [](#VUID-VkDescriptorSetAllocateInfo-pSetLayouts-04610)VUID-VkDescriptorSetAllocateInfo-pSetLayouts-04610  
  If any element of `pSetLayouts` was created with the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT` bit set, `descriptorPool` **must** have been created with the `VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT` flag set
- [](#VUID-VkDescriptorSetAllocateInfo-pSetLayouts-08009)VUID-VkDescriptorSetAllocateInfo-pSetLayouts-08009  
  Each element of `pSetLayouts` **must** not have been created with the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT` bit set

Valid Usage (Implicit)

- [](#VUID-VkDescriptorSetAllocateInfo-sType-sType)VUID-VkDescriptorSetAllocateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO`
- [](#VUID-VkDescriptorSetAllocateInfo-pNext-pNext)VUID-VkDescriptorSetAllocateInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkDescriptorSetVariableDescriptorCountAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetVariableDescriptorCountAllocateInfo.html)
- [](#VUID-VkDescriptorSetAllocateInfo-sType-unique)VUID-VkDescriptorSetAllocateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkDescriptorSetAllocateInfo-descriptorPool-parameter)VUID-VkDescriptorSetAllocateInfo-descriptorPool-parameter  
  `descriptorPool` **must** be a valid [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPool.html) handle
- [](#VUID-VkDescriptorSetAllocateInfo-pSetLayouts-parameter)VUID-VkDescriptorSetAllocateInfo-pSetLayouts-parameter  
  `pSetLayouts` **must** be a valid pointer to an array of `descriptorSetCount` valid [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) handles
- [](#VUID-VkDescriptorSetAllocateInfo-descriptorSetCount-arraylength)VUID-VkDescriptorSetAllocateInfo-descriptorSetCount-arraylength  
  `descriptorSetCount` **must** be greater than `0`
- [](#VUID-VkDescriptorSetAllocateInfo-commonparent)VUID-VkDescriptorSetAllocateInfo-commonparent  
  Both of `descriptorPool`, and the elements of `pSetLayouts` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `descriptorPool` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPool.html), [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkAllocateDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAllocateDescriptorSets.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorSetAllocateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0