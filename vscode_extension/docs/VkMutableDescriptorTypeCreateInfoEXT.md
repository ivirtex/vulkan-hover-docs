# VkMutableDescriptorTypeCreateInfoEXT(3) Manual Page

## Name

VkMutableDescriptorTypeCreateInfoEXT - Structure describing the list of
possible active descriptor types for mutable type descriptors



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of a
[VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html)
or [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateInfo.html)
structure includes a
[VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeCreateInfoEXT.html)
structure, then that structure specifies Information about the possible
descriptor types for mutable descriptor types.

The `VkMutableDescriptorTypeCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_mutable_descriptor_type
typedef struct VkMutableDescriptorTypeCreateInfoEXT {
    VkStructureType                          sType;
    const void*                              pNext;
    uint32_t                                 mutableDescriptorTypeListCount;
    const VkMutableDescriptorTypeListEXT*    pMutableDescriptorTypeLists;
} VkMutableDescriptorTypeCreateInfoEXT;
```

or the equivalent

``` c
// Provided by VK_VALVE_mutable_descriptor_type
typedef VkMutableDescriptorTypeCreateInfoEXT VkMutableDescriptorTypeCreateInfoVALVE;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `mutableDescriptorTypeListCount` is the number of elements in
  `pMutableDescriptorTypeLists`.

- `pMutableDescriptorTypeLists` is a pointer to an array of
  `VkMutableDescriptorTypeListEXT` structures.

## <a href="#_description" class="anchor"></a>Description

If `mutableDescriptorTypeListCount` is zero or if this structure is not
included in the `pNext` chain, the
[VkMutableDescriptorTypeListEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeListEXT.html)
for each element is considered to be zero or `NULL` for each member.
Otherwise, the descriptor set layout binding at
[VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html)::`pBindings`\[i\]
uses the descriptor type lists in
[VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeCreateInfoEXT.html)::`pMutableDescriptorTypeLists`\[i\].

Valid Usage (Implicit)

- <a href="#VUID-VkMutableDescriptorTypeCreateInfoEXT-sType-sType"
  id="VUID-VkMutableDescriptorTypeCreateInfoEXT-sType-sType"></a>
  VUID-VkMutableDescriptorTypeCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_MUTABLE_DESCRIPTOR_TYPE_CREATE_INFO_EXT`

- <a
  href="#VUID-VkMutableDescriptorTypeCreateInfoEXT-pMutableDescriptorTypeLists-parameter"
  id="VUID-VkMutableDescriptorTypeCreateInfoEXT-pMutableDescriptorTypeLists-parameter"></a>
  VUID-VkMutableDescriptorTypeCreateInfoEXT-pMutableDescriptorTypeLists-parameter  
  If `mutableDescriptorTypeListCount` is not `0`,
  `pMutableDescriptorTypeLists` **must** be a valid pointer to an array
  of `mutableDescriptorTypeListCount` valid
  [VkMutableDescriptorTypeListEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeListEXT.html)
  structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_mutable_descriptor_type](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_mutable_descriptor_type.html),
[VK_VALVE_mutable_descriptor_type](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VALVE_mutable_descriptor_type.html),
[VkMutableDescriptorTypeListEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeListEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMutableDescriptorTypeCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
