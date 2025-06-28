# VkMutableDescriptorTypeCreateInfoEXT(3) Manual Page

## Name

VkMutableDescriptorTypeCreateInfoEXT - Structure describing the list of possible active descriptor types for mutable type descriptors



## [](#_c_specification)C Specification

If the `pNext` chain of a [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html) or [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html) structure includes a [VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMutableDescriptorTypeCreateInfoEXT.html) structure, then that structure specifies Information about the possible descriptor types for mutable descriptor types.

The `VkMutableDescriptorTypeCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_mutable_descriptor_type
typedef struct VkMutableDescriptorTypeCreateInfoEXT {
    VkStructureType                          sType;
    const void*                              pNext;
    uint32_t                                 mutableDescriptorTypeListCount;
    const VkMutableDescriptorTypeListEXT*    pMutableDescriptorTypeLists;
} VkMutableDescriptorTypeCreateInfoEXT;
```

or the equivalent

```c++
// Provided by VK_VALVE_mutable_descriptor_type
typedef VkMutableDescriptorTypeCreateInfoEXT VkMutableDescriptorTypeCreateInfoVALVE;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `mutableDescriptorTypeListCount` is the number of elements in `pMutableDescriptorTypeLists`.
- `pMutableDescriptorTypeLists` is a pointer to an array of `VkMutableDescriptorTypeListEXT` structures.

## [](#_description)Description

If `mutableDescriptorTypeListCount` is zero or if this structure is not included in the `pNext` chain, the [VkMutableDescriptorTypeListEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMutableDescriptorTypeListEXT.html) for each element is considered to be zero or `NULL` for each member. Otherwise, the descriptor set layout binding at [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html)::`pBindings`\[i] uses the descriptor type lists in [VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMutableDescriptorTypeCreateInfoEXT.html)::`pMutableDescriptorTypeLists`\[i].

Valid Usage (Implicit)

- [](#VUID-VkMutableDescriptorTypeCreateInfoEXT-sType-sType)VUID-VkMutableDescriptorTypeCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MUTABLE_DESCRIPTOR_TYPE_CREATE_INFO_EXT`
- [](#VUID-VkMutableDescriptorTypeCreateInfoEXT-pMutableDescriptorTypeLists-parameter)VUID-VkMutableDescriptorTypeCreateInfoEXT-pMutableDescriptorTypeLists-parameter  
  If `mutableDescriptorTypeListCount` is not `0`, `pMutableDescriptorTypeLists` **must** be a valid pointer to an array of `mutableDescriptorTypeListCount` valid [VkMutableDescriptorTypeListEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMutableDescriptorTypeListEXT.html) structures

## [](#_see_also)See Also

[VK\_EXT\_mutable\_descriptor\_type](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_mutable_descriptor_type.html), [VK\_VALVE\_mutable\_descriptor\_type](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VALVE_mutable_descriptor_type.html), [VkMutableDescriptorTypeListEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMutableDescriptorTypeListEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMutableDescriptorTypeCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0