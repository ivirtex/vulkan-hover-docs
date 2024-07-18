# VkDescriptorSetVariableDescriptorCountAllocateInfo(3) Manual Page

## Name

VkDescriptorSetVariableDescriptorCountAllocateInfo - Structure
specifying additional allocation parameters for descriptor sets



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of a
[VkDescriptorSetAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetAllocateInfo.html)
structure includes a
`VkDescriptorSetVariableDescriptorCountAllocateInfo` structure, then
that structure includes an array of descriptor counts for variable-sized
descriptor bindings, one for each descriptor set being allocated.

The `VkDescriptorSetVariableDescriptorCountAllocateInfo` structure is
defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkDescriptorSetVariableDescriptorCountAllocateInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           descriptorSetCount;
    const uint32_t*    pDescriptorCounts;
} VkDescriptorSetVariableDescriptorCountAllocateInfo;
```

or the equivalent

``` c
// Provided by VK_EXT_descriptor_indexing
typedef VkDescriptorSetVariableDescriptorCountAllocateInfo VkDescriptorSetVariableDescriptorCountAllocateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `descriptorSetCount` is zero or the number of elements in
  `pDescriptorCounts`.

- `pDescriptorCounts` is a pointer to an array of descriptor counts,
  with each member specifying the number of descriptors in a
  variable-sized descriptor binding in the corresponding descriptor set
  being allocated.

## <a href="#_description" class="anchor"></a>Description

If `descriptorSetCount` is zero or this structure is not included in the
`pNext` chain, then the variable lengths are considered to be zero.
Otherwise, `pDescriptorCounts`\[i\] is the number of descriptors in the
variable-sized descriptor binding in the corresponding descriptor set
layout. If the variable-sized descriptor binding in the corresponding
descriptor set layout has a descriptor type of
`VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` then `pDescriptorCounts`\[i\]
specifies the binding’s capacity in bytes. If
[VkDescriptorSetAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetAllocateInfo.html)::`pSetLayouts`\[i\]
does not include a variable-sized descriptor binding, then
`pDescriptorCounts`\[i\] is ignored.

Valid Usage

- <a
  href="#VUID-VkDescriptorSetVariableDescriptorCountAllocateInfo-descriptorSetCount-03045"
  id="VUID-VkDescriptorSetVariableDescriptorCountAllocateInfo-descriptorSetCount-03045"></a>
  VUID-VkDescriptorSetVariableDescriptorCountAllocateInfo-descriptorSetCount-03045  
  If `descriptorSetCount` is not zero, `descriptorSetCount` **must**
  equal
  [VkDescriptorSetAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetAllocateInfo.html)::`descriptorSetCount`

Valid Usage (Implicit)

- <a
  href="#VUID-VkDescriptorSetVariableDescriptorCountAllocateInfo-sType-sType"
  id="VUID-VkDescriptorSetVariableDescriptorCountAllocateInfo-sType-sType"></a>
  VUID-VkDescriptorSetVariableDescriptorCountAllocateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO`

- <a
  href="#VUID-VkDescriptorSetVariableDescriptorCountAllocateInfo-pDescriptorCounts-parameter"
  id="VUID-VkDescriptorSetVariableDescriptorCountAllocateInfo-pDescriptorCounts-parameter"></a>
  VUID-VkDescriptorSetVariableDescriptorCountAllocateInfo-pDescriptorCounts-parameter  
  If `descriptorSetCount` is not `0`, `pDescriptorCounts` **must** be a
  valid pointer to an array of `descriptorSetCount` `uint32_t` values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_indexing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_indexing.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorSetVariableDescriptorCountAllocateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
