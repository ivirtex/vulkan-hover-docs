# VkDescriptorSetAllocateInfo(3) Manual Page

## Name

VkDescriptorSetAllocateInfo - Structure specifying the allocation
parameters for descriptor sets



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDescriptorSetAllocateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkDescriptorSetAllocateInfo {
    VkStructureType                 sType;
    const void*                     pNext;
    VkDescriptorPool                descriptorPool;
    uint32_t                        descriptorSetCount;
    const VkDescriptorSetLayout*    pSetLayouts;
} VkDescriptorSetAllocateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `descriptorPool` is the pool which the sets will be allocated from.

- `descriptorSetCount` determines the number of descriptor sets to be
  allocated from the pool.

- `pSetLayouts` is a pointer to an array of descriptor set layouts, with
  each member specifying how the corresponding descriptor set is
  allocated.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkDescriptorSetAllocateInfo-apiVersion-07895"
  id="VUID-VkDescriptorSetAllocateInfo-apiVersion-07895"></a>
  VUID-VkDescriptorSetAllocateInfo-apiVersion-07895  
  If the [VK_KHR_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html) extension is
  not enabled and
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, `descriptorSetCount` **must** not be greater
  than the number of sets that are currently available for allocation in
  `descriptorPool`

- <a href="#VUID-VkDescriptorSetAllocateInfo-apiVersion-07896"
  id="VUID-VkDescriptorSetAllocateInfo-apiVersion-07896"></a>
  VUID-VkDescriptorSetAllocateInfo-apiVersion-07896  
  If the [VK_KHR_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html) extension is
  not enabled and
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, `descriptorPool` **must** have enough free
  descriptor capacity remaining to allocate the descriptor sets of the
  specified layouts

- <a href="#VUID-VkDescriptorSetAllocateInfo-pSetLayouts-00308"
  id="VUID-VkDescriptorSetAllocateInfo-pSetLayouts-00308"></a>
  VUID-VkDescriptorSetAllocateInfo-pSetLayouts-00308  
  Each element of `pSetLayouts` **must** not have been created with
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR` set

- <a href="#VUID-VkDescriptorSetAllocateInfo-pSetLayouts-03044"
  id="VUID-VkDescriptorSetAllocateInfo-pSetLayouts-03044"></a>
  VUID-VkDescriptorSetAllocateInfo-pSetLayouts-03044  
  If any element of `pSetLayouts` was created with the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set,
  `descriptorPool` **must** have been created with the
  `VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT` flag set

- <a href="#VUID-VkDescriptorSetAllocateInfo-pSetLayouts-09380"
  id="VUID-VkDescriptorSetAllocateInfo-pSetLayouts-09380"></a>
  VUID-VkDescriptorSetAllocateInfo-pSetLayouts-09380  
  If `pSetLayouts`\[i\] was created with an element of `pBindingFlags`
  that includes `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`,
  and
  [VkDescriptorSetVariableDescriptorCountAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetVariableDescriptorCountAllocateInfo.html)
  is included in the `pNext` chain, and
  `VkDescriptorSetVariableDescriptorCountAllocateInfo`::`descriptorSetCount`
  is not zero, then
  [VkDescriptorSetVariableDescriptorCountAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetVariableDescriptorCountAllocateInfo.html)::`pDescriptorCounts`\[i\]
  **must** be less than or equal to
  [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutBinding.html)::`descriptorCount`
  for the corresponding binding used to create `pSetLayouts`\[i\]

- <a href="#VUID-VkDescriptorSetAllocateInfo-pSetLayouts-04610"
  id="VUID-VkDescriptorSetAllocateInfo-pSetLayouts-04610"></a>
  VUID-VkDescriptorSetAllocateInfo-pSetLayouts-04610  
  If any element of `pSetLayouts` was created with the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT` bit set,
  `descriptorPool` **must** have been created with the
  `VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT` flag set

- <a href="#VUID-VkDescriptorSetAllocateInfo-pSetLayouts-08009"
  id="VUID-VkDescriptorSetAllocateInfo-pSetLayouts-08009"></a>
  VUID-VkDescriptorSetAllocateInfo-pSetLayouts-08009  
  Each element of `pSetLayouts` **must** not have been created with the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT` bit set

Valid Usage (Implicit)

- <a href="#VUID-VkDescriptorSetAllocateInfo-sType-sType"
  id="VUID-VkDescriptorSetAllocateInfo-sType-sType"></a>
  VUID-VkDescriptorSetAllocateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO`

- <a href="#VUID-VkDescriptorSetAllocateInfo-pNext-pNext"
  id="VUID-VkDescriptorSetAllocateInfo-pNext-pNext"></a>
  VUID-VkDescriptorSetAllocateInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkDescriptorSetVariableDescriptorCountAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetVariableDescriptorCountAllocateInfo.html)

- <a href="#VUID-VkDescriptorSetAllocateInfo-sType-unique"
  id="VUID-VkDescriptorSetAllocateInfo-sType-unique"></a>
  VUID-VkDescriptorSetAllocateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkDescriptorSetAllocateInfo-descriptorPool-parameter"
  id="VUID-VkDescriptorSetAllocateInfo-descriptorPool-parameter"></a>
  VUID-VkDescriptorSetAllocateInfo-descriptorPool-parameter  
  `descriptorPool` **must** be a valid
  [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html) handle

- <a href="#VUID-VkDescriptorSetAllocateInfo-pSetLayouts-parameter"
  id="VUID-VkDescriptorSetAllocateInfo-pSetLayouts-parameter"></a>
  VUID-VkDescriptorSetAllocateInfo-pSetLayouts-parameter  
  `pSetLayouts` **must** be a valid pointer to an array of
  `descriptorSetCount` valid
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) handles

- <a
  href="#VUID-VkDescriptorSetAllocateInfo-descriptorSetCount-arraylength"
  id="VUID-VkDescriptorSetAllocateInfo-descriptorSetCount-arraylength"></a>
  VUID-VkDescriptorSetAllocateInfo-descriptorSetCount-arraylength  
  `descriptorSetCount` **must** be greater than `0`

- <a href="#VUID-VkDescriptorSetAllocateInfo-commonparent"
  id="VUID-VkDescriptorSetAllocateInfo-commonparent"></a>
  VUID-VkDescriptorSetAllocateInfo-commonparent  
  Both of `descriptorPool`, and the elements of `pSetLayouts` **must**
  have been created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html),
[VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkAllocateDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAllocateDescriptorSets.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorSetAllocateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
